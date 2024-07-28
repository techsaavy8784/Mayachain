package evm

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ecommon "github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gitlab.com/mayachain/mayanode/bifrost/blockscanner"
	"gitlab.com/mayachain/mayanode/bifrost/mayaclient"
	stypes "gitlab.com/mayachain/mayanode/bifrost/mayaclient/types"
	"gitlab.com/mayachain/mayanode/bifrost/metrics"
	"gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/shared/evm"
	"gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/shared/runners"
	"gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/shared/signercache"
	"gitlab.com/mayachain/mayanode/bifrost/pubkeymanager"
	"gitlab.com/mayachain/mayanode/bifrost/tss"
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/config"
	"gitlab.com/mayachain/mayanode/constants"
	mem "gitlab.com/mayachain/mayanode/x/mayachain/memo"
	tssp "gitlab.com/thorchain/tss/go-tss/tss"
)

////////////////////////////////////////////////////////////////////////////////////////
// EVMClient
////////////////////////////////////////////////////////////////////////////////////////

const (
	MaxContractGas = 4_000_000
)

// EVMClient is a generic client for interacting with EVM chains.
type EVMClient struct {
	logger                  zerolog.Logger
	cfg                     config.BifrostChainConfiguration
	localPubKey             common.PubKey
	kw                      *evm.KeySignWrapper
	ethClient               *ethclient.Client
	evmScanner              *EVMScanner
	bridge                  mayaclient.MayachainBridge
	blockScanner            *blockscanner.BlockScanner
	vaultABI                *abi.ABI
	pubkeyMgr               pubkeymanager.PubKeyValidator
	poolMgr                 mayaclient.PoolManager
	tssKeySigner            *tss.KeySign
	wg                      *sync.WaitGroup
	stopchan                chan struct{}
	globalSolvencyQueue     chan stypes.Solvency
	signerCacheManager      *signercache.CacheManager
	lastSolvencyCheckHeight int64
}

// NewEVMClient creates a new EVMClient.
func NewEVMClient(
	thorKeys *mayaclient.Keys,
	cfg config.BifrostChainConfiguration,
	server *tssp.TssServer,
	bridge mayaclient.MayachainBridge,
	m *metrics.Metrics,
	pubkeyMgr pubkeymanager.PubKeyValidator,
	poolMgr mayaclient.PoolManager,
) (*EVMClient, error) {
	// check required arguments
	if thorKeys == nil {
		return nil, fmt.Errorf("failed to create EVM client, maya keys empty")
	}
	if bridge == nil {
		return nil, errors.New("mayachain bridge is nil")
	}
	if pubkeyMgr == nil {
		return nil, errors.New("pubkey manager is nil")
	}
	if poolMgr == nil {
		return nil, errors.New("pool manager is nil")

		// create keys
	}
	tssKm, err := tss.NewKeySign(server, bridge)
	if err != nil {
		return nil, fmt.Errorf("failed to create tss signer: %w", err)
	}

	priv, err := thorKeys.GetPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get private key: %w", err)
	}

	temp, err := codec.ToTmPubKeyInterface(priv.PubKey())
	if err != nil {
		return nil, fmt.Errorf("failed to get tm pub key: %w", err)
	}
	pk, err := common.NewPubKeyFromCrypto(temp)
	if err != nil {
		return nil, fmt.Errorf("failed to get pub key: %w", err)
	}
	evmPrivateKey, err := evm.GetPrivateKey(priv)
	if err != nil {
		return nil, err
	}

	// create rpc clients
	rpcClient, err := evm.NewEthRPC(cfg.RPCHost, cfg.BlockScanner.HTTPRequestTimeout, cfg.ChainID.String())
	if err != nil {
		return nil, fmt.Errorf("fail to create ETH rpc host(%s): %w", cfg.RPCHost, err)
	}

	ethClient, err := ethclient.Dial(cfg.RPCHost)
	if err != nil {
		return nil, fmt.Errorf("fail to dial ETH rpc host(%s): %w", cfg.RPCHost, err)
	}

	// get chain id
	chainID, err := getChainID(ethClient, cfg.BlockScanner.HTTPRequestTimeout)
	if err != nil {
		return nil, err
	}
	if chainID.Uint64() == 0 {
		return nil, fmt.Errorf("chain id is: %d , invalid", chainID.Uint64())
	}

	// create keysign wrapper
	keysignWrapper, err := evm.NewKeySignWrapper(evmPrivateKey, pk, tssKm, chainID, cfg.ChainID.String())
	if err != nil {
		return nil, fmt.Errorf("fail to create %s key sign wrapper: %w", cfg.ChainID, err)
	}

	// load vault abi
	vaultABI, _, err := evm.GetContractABI(routerContractABI, erc20ContractABI)
	if err != nil {
		return nil, fmt.Errorf("fail to get contract abi: %w", err)
	}

	// TODO: Do we need to call this?
	pubkeyMgr.GetPubKeys()

	c := &EVMClient{
		logger:       log.With().Str("module", "evm").Stringer("chain", cfg.ChainID).Logger(),
		cfg:          cfg,
		ethClient:    ethClient,
		localPubKey:  pk,
		kw:           keysignWrapper,
		bridge:       bridge,
		vaultABI:     vaultABI,
		pubkeyMgr:    pubkeyMgr,
		poolMgr:      poolMgr,
		tssKeySigner: tssKm,
		wg:           &sync.WaitGroup{},
		stopchan:     make(chan struct{}),
	}

	// initialize storage
	var path string // if not set later, will in memory storage
	if len(c.cfg.BlockScanner.DBPath) > 0 {
		path = fmt.Sprintf("%s/%s", c.cfg.BlockScanner.DBPath, c.cfg.BlockScanner.ChainID)
	}
	storage, err := blockscanner.NewBlockScannerStorage(path)
	if err != nil {
		return c, fmt.Errorf("fail to create blockscanner storage: %w", err)
	}
	signerCacheManager, err := signercache.NewSignerCacheManager(storage.GetInternalDb())
	if err != nil {
		return nil, fmt.Errorf("fail to create signer cache manager")
	}
	c.signerCacheManager = signerCacheManager

	// create block scanner
	c.evmScanner, err = NewEVMScanner(
		c.cfg.BlockScanner,
		storage,
		chainID,
		ethClient,
		rpcClient,
		c.bridge,
		m,
		pubkeyMgr,
		c.ReportSolvency,
		signerCacheManager,
	)
	if err != nil {
		return c, fmt.Errorf("fail to create evm block scanner: %w", err)
	}

	// initialize block scanner
	c.blockScanner, err = blockscanner.NewBlockScanner(
		c.cfg.BlockScanner, storage, m, c.bridge, c.evmScanner,
	)
	if err != nil {
		return c, fmt.Errorf("fail to create block scanner: %w", err)
	}

	// TODO: Is this necessary?
	localNodeAddress, err := c.localPubKey.GetAddress(cfg.ChainID)
	if err != nil {
		c.logger.Err(err).Stringer("chain", cfg.ChainID).Msg("failed to get local node address")
	}
	c.logger.Info().
		Stringer("chain", cfg.ChainID).
		Stringer("address", localNodeAddress).
		Msg("local node address")

	return c, nil
}

// Start starts the chain client with the given queues.
func (c *EVMClient) Start(globalTxsQueue chan stypes.TxIn, globalErrataQueue chan stypes.ErrataBlock, globalSolvencyQueue chan stypes.Solvency) {
	c.globalSolvencyQueue = globalSolvencyQueue
	c.tssKeySigner.Start()
	c.blockScanner.Start(globalTxsQueue)
	c.wg.Add(1)
	go c.unstuck()
	c.wg.Add(1)
	go runners.SolvencyCheckRunner(c.GetChain(), c, c.bridge, c.stopchan, c.wg, constants.MayachainBlockTime)
}

// Stop stops the chain client.
func (c *EVMClient) Stop() {
	c.tssKeySigner.Stop()
	c.blockScanner.Stop()
	close(c.stopchan)
	c.wg.Wait()
}

// IsBlockScannerHealthy returns true if the block scanner is healthy.
func (c *EVMClient) IsBlockScannerHealthy() bool {
	return c.blockScanner.IsHealthy()
}

// --------------------------------- config ---------------------------------

// GetConfig returns the chain configuration.
func (c *EVMClient) GetConfig() config.BifrostChainConfiguration {
	return c.cfg
}

// GetChain returns the chain.
func (c *EVMClient) GetChain() common.Chain {
	return c.cfg.ChainID
}

// --------------------------------- status ---------------------------------

// GetHeight returns the current height of the chain.
func (c *EVMClient) GetHeight() (int64, error) {
	return c.evmScanner.GetHeight()
}

// GetBlockScannerHeight returns blockscanner height
func (c *EVMClient) GetBlockScannerHeight() (int64, error) {
	return c.blockScanner.PreviousHeight(), nil
}

func (c *EVMClient) GetLatestTxForVault(vault string) (string, string, error) {
	lastObserved, err := c.signerCacheManager.GetLatestRecordedTx(stypes.InboundCacheKey(vault, c.GetChain().String()))
	if err != nil {
		return "", "", err
	}
	lastBroadCasted, err := c.signerCacheManager.GetLatestRecordedTx(stypes.BroadcastCacheKey(vault, c.GetChain().String()))
	return lastObserved, lastBroadCasted, err
}

// GetAddress returns the address for the given public key.
func (c *EVMClient) GetAddress(poolPubKey common.PubKey) string {
	addr, err := poolPubKey.GetAddress(c.cfg.ChainID)
	if err != nil {
		c.logger.Error().Err(err).Str("pool_pub_key", poolPubKey.String()).Msg("fail to get pool address")
		return ""
	}
	return addr.String()
}

// GetAccount returns the account for the given public key.
func (c *EVMClient) GetAccount(pk common.PubKey, height *big.Int) (common.Account, error) {
	addr := c.GetAddress(pk)
	nonce, err := c.evmScanner.GetNonce(addr)
	if err != nil {
		return common.Account{}, err
	}
	coins, err := c.GetBalances(addr, height)
	if err != nil {
		return common.Account{}, err
	}
	account := common.NewAccount(int64(nonce), 0, coins, false)
	return account, nil
}

// GetAccountByAddress returns the account for the given address.
func (c *EVMClient) GetAccountByAddress(address string, height *big.Int) (common.Account, error) {
	nonce, err := c.evmScanner.GetNonce(address)
	if err != nil {
		return common.Account{}, err
	}
	coins, err := c.GetBalances(address, height)
	if err != nil {
		return common.Account{}, err
	}
	account := common.NewAccount(int64(nonce), 0, coins, false)
	return account, nil
}

/* Gas-related methods */

func (c *EVMClient) getSmartContractAddr(pubkey common.PubKey) common.Address {
	return c.pubkeyMgr.GetContract(c.cfg.ChainID, pubkey)
}

func (c *EVMClient) getSmartContractByAddress(addr common.Address) common.Address {
	for _, pk := range c.pubkeyMgr.GetPubKeys() {
		evmAddr, err := pk.GetAddress(c.cfg.ChainID)
		if err != nil {
			return common.NoAddress
		}
		if evmAddr.Equals(addr) {
			return c.pubkeyMgr.GetContract(c.cfg.ChainID, pk)
		}
	}
	return common.NoAddress
}

func (c *EVMClient) getTokenAddressFromAsset(asset common.Asset) string {
	if asset.Equals(c.cfg.ChainID.GetGasAsset()) {
		return evm.NativeTokenAddr
	}
	allParts := strings.Split(asset.Symbol.String(), "-")
	return allParts[len(allParts)-1]
}

// --------------------------------- balances ---------------------------------

// GetBalance returns the balance of the provided address.
func (c *EVMClient) GetBalance(addr, token string, height *big.Int) (*big.Int, error) {
	contractAddresses := c.pubkeyMgr.GetContracts(c.cfg.ChainID)
	c.logger.Debug().Interface("contractAddresses", contractAddresses).Msg("got contracts")
	if len(contractAddresses) == 0 {
		return nil, fmt.Errorf("fail to get contract address")
	}

	return c.evmScanner.tokenManager.GetBalance(addr, token, height, contractAddresses[0].String())
}

// GetBalances returns the balances of the provided address.
func (c *EVMClient) GetBalances(addr string, height *big.Int) (common.Coins, error) {
	// for all the tokens the chain client has dealt with before
	tokens, err := c.evmScanner.GetTokens()
	if err != nil {
		return nil, fmt.Errorf("fail to get all the tokens: %w", err)
	}
	coins := common.Coins{}
	for _, token := range tokens {
		balance, err := c.GetBalance(addr, token.Address, height)
		if err != nil {
			c.logger.Err(err).Str("token", token.Address).Msg("fail to get balance for token")
			continue
		}
		asset := c.cfg.ChainID.GetGasAsset()
		if !strings.EqualFold(token.Address, evm.NativeTokenAddr) {
			asset, err = common.NewAsset(fmt.Sprintf("EVM.%s-%s", token.Symbol, token.Address))
			if err != nil {
				return nil, err
			}
		}
		bal := c.evmScanner.tokenManager.ConvertAmount(token.Address, balance)
		coins = append(coins, common.NewCoin(asset, bal))
	}

	return coins.Distinct(), nil
}

// --------------------------------- gas ---------------------------------

// GetGasFee returns the gas fee based on the current gas price.
func (c *EVMClient) GetGasFee(gas uint64) common.Gas {
	return common.GetEVMGasFee(c.cfg.ChainID, c.GetGasPrice(), gas)
}

// GetGasPrice returns the current gas price.
func (c *EVMClient) GetGasPrice() *big.Int {
	gasPrice := c.evmScanner.GetGasPrice()
	return gasPrice
}

// --------------------------------- build transaction ---------------------------------

// getOutboundTxData generates the tx data and tx value of the outbound Router Contract call, and checks if the router contract has been updated
func (c *EVMClient) getOutboundTxData(txOutItem stypes.TxOutItem, memo mem.Memo, contractAddr common.Address) ([]byte, bool, *big.Int, error) {
	var data []byte
	var err error
	var tokenAddr string
	value := big.NewInt(0)
	evmValue := big.NewInt(0)
	hasRouterUpdated := false

	if len(txOutItem.Coins) == 1 {
		coin := txOutItem.Coins[0]
		tokenAddr = c.getTokenAddressFromAsset(coin.Asset)
		value = value.Add(value, coin.Amount.BigInt())
		value = c.evmScanner.tokenManager.ConvertSigningAmount(value, tokenAddr)
		if strings.EqualFold(tokenAddr, evm.NativeTokenAddr) {
			evmValue = value
		}
	}

	toAddr := ecommon.HexToAddress(txOutItem.ToAddress.String())

	switch memo.GetType() {
	case mem.TxOutbound, mem.TxRefund, mem.TxRagnarok:
		if txOutItem.Aggregator == "" {
			data, err = c.vaultABI.Pack("transferOut", toAddr, ecommon.HexToAddress(tokenAddr), value, txOutItem.Memo)
			if err != nil {
				return nil, hasRouterUpdated, nil, fmt.Errorf("fail to create data to call smart contract(transferOut): %w", err)
			}
		} else {
			memoType := memo.GetType()
			if memoType == mem.TxRefund || memoType == mem.TxRagnarok {
				return nil, hasRouterUpdated, nil, fmt.Errorf("%s can't use transferOutAndCall", memoType)
			}
			c.logger.Info().Msgf("aggregator target asset address: %s", txOutItem.AggregatorTargetAsset)
			if evmValue.Uint64() == 0 {
				return nil, hasRouterUpdated, nil, fmt.Errorf("transferOutAndCall can only be used when outbound asset is native")
			}
			targetLimit := txOutItem.AggregatorTargetLimit
			if targetLimit == nil {
				zeroLimit := cosmos.ZeroUint()
				targetLimit = &zeroLimit
			}
			aggAddr := ecommon.HexToAddress(txOutItem.Aggregator)
			targetAddr := ecommon.HexToAddress(txOutItem.AggregatorTargetAsset)
			// when address can't be round trip , the tx out item will be dropped
			if !strings.EqualFold(aggAddr.String(), txOutItem.Aggregator) {
				c.logger.Error().Msgf("aggregator address can't roundtrip , ignore tx (%s != %s)", txOutItem.Aggregator, aggAddr.String())
				return nil, hasRouterUpdated, nil, nil
			}
			if !strings.EqualFold(targetAddr.String(), txOutItem.AggregatorTargetAsset) {
				c.logger.Error().Msgf("aggregator target asset address can't roundtrip , ignore tx (%s != %s)", txOutItem.AggregatorTargetAsset, targetAddr.String())
				return nil, hasRouterUpdated, nil, nil
			}
			data, err = c.vaultABI.Pack("transferOutAndCall", aggAddr, targetAddr, toAddr, targetLimit.BigInt(), txOutItem.Memo)
			if err != nil {
				return nil, hasRouterUpdated, nil, fmt.Errorf("fail to create data to call smart contract(transferOutAndCall): %w", err)
			}
		}
	case mem.TxMigrate, mem.TxYggdrasilFund:
		if txOutItem.Aggregator != "" || txOutItem.AggregatorTargetAsset != "" {
			return nil, hasRouterUpdated, nil, fmt.Errorf("migration / yggdrasil+ can't use aggregator")
		}
		if strings.EqualFold(tokenAddr, evm.NativeTokenAddr) {
			data, err = c.vaultABI.Pack("transferOut", toAddr, ecommon.HexToAddress(tokenAddr), value, txOutItem.Memo)
			if err != nil {
				return nil, hasRouterUpdated, nil, fmt.Errorf("fail to create data to call smart contract(transferOut): %w", err)
			}
		} else {
			newSmartContractAddr := c.getSmartContractByAddress(txOutItem.ToAddress)
			if newSmartContractAddr.IsEmpty() {
				return nil, hasRouterUpdated, nil, fmt.Errorf("fail to get new smart contract address")
			}
			data, err = c.vaultABI.Pack("transferAllowance", ecommon.HexToAddress(newSmartContractAddr.String()), toAddr, ecommon.HexToAddress(tokenAddr), value, txOutItem.Memo)
			if err != nil {
				return nil, hasRouterUpdated, nil, fmt.Errorf("fail to create data to call smart contract(transferAllowance): %w", err)
			}
		}
	case mem.TxYggdrasilReturn:
		if txOutItem.Aggregator != "" || txOutItem.AggregatorTargetAsset != "" {
			return nil, hasRouterUpdated, nil, fmt.Errorf("yggdrasil- can't use aggregator")
		}
		newSmartContractAddr := c.getSmartContractByAddress(txOutItem.ToAddress)
		if newSmartContractAddr.IsEmpty() {
			return nil, hasRouterUpdated, nil, fmt.Errorf("fail to get new smart contract address")
		}
		hasRouterUpdated = !newSmartContractAddr.Equals(contractAddr)

		var coins []evm.RouterCoin
		for _, item := range txOutItem.Coins {
			assetAddr := c.getTokenAddressFromAsset(item.Asset)
			assetAmt := c.evmScanner.tokenManager.ConvertSigningAmount(item.Amount.BigInt(), assetAddr)
			if strings.EqualFold(assetAddr, evm.NativeTokenAddr) {
				evmValue = assetAmt
				continue
			}
			coins = append(coins, evm.RouterCoin{
				Asset:  ecommon.HexToAddress(assetAddr),
				Amount: assetAmt,
			})
		}
		data, err = c.vaultABI.Pack("returnVaultAssets", ecommon.HexToAddress(newSmartContractAddr.String()), toAddr, coins, txOutItem.Memo)
		if err != nil {
			return nil, hasRouterUpdated, nil, fmt.Errorf("fail to create data to call smart contract(transferVaultAssets): %w", err)
		}
	}
	return data, hasRouterUpdated, evmValue, nil
}

func (c *EVMClient) buildOutboundTx(txOutItem stypes.TxOutItem, memo mem.Memo, nonce uint64) (*etypes.Transaction, error) {
	contractAddr := c.getSmartContractAddr(txOutItem.VaultPubKey)
	if contractAddr.IsEmpty() {
		// we may be churning from a vault that does not have a contract
		// try getting the toAddress (new vault) contract instead
		if memo.GetType() == mem.TxMigrate {
			contractAddr = c.getSmartContractByAddress(txOutItem.ToAddress)
		}
		if contractAddr.IsEmpty() {
			if c.cfg.ChainID != common.ARBChain {
				return nil, fmt.Errorf("can't sign tx, fail to get smart contract address")
			}

			// If chain is ARB we will not get a contract address, we will use the to_address from the tx
			contractAddr = txOutItem.ToAddress
		}
	}

	fromAddr, err := txOutItem.VaultPubKey.GetAddress(c.cfg.ChainID)
	if err != nil {
		return nil, fmt.Errorf("fail to get EVM address for pub key(%s): %w", txOutItem.VaultPubKey, err)
	}

	txData, _, evmValue, err := c.getOutboundTxData(txOutItem, memo, contractAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get outbound tx data %w", err)
	}
	if evmValue == nil {
		evmValue = cosmos.ZeroUint().BigInt()
	}

	// compare the gas rate prescribed by MAYAChain against the price it can get from the chain
	// ensure signer always pay enough higher gas price
	// GasRate from mayachain is in 1e8, need to convert to Wei/CentiWei
	var gasRate *big.Int
	if c.cfg.ChainID.Equals(common.ARBChain) {
		gasRate = convertThorchainAmountToCentiWei(big.NewInt(txOutItem.GasRate))
	} else {
		gasRate = convertThorchainAmountToWei(big.NewInt(txOutItem.GasRate))
	}

	if gasRate.Cmp(c.GetGasPrice()) < 0 {
		gasRate = c.GetGasPrice()
	}
	// outbound tx always send to smart contract address
	estimatedEVMValue := big.NewInt(0)
	if evmValue.Uint64() > 0 {
		// when the EVM value is non-zero, here override it with a fixed value to estimate gas
		// when EVM value is non-zero, if we send the real value for estimate gas, sometimes it will fail, for many reasons, a few I saw during test
		// 1. insufficient fund
		// 2. gas required exceeds allowance
		// as long as we pass in an EVM value , which we almost guarantee it will not exceed the EVM balance , so we can avoid the above two errors
		estimatedEVMValue = estimatedEVMValue.SetInt64(21000)
	}
	createdTx := etypes.NewTransaction(nonce, ecommon.HexToAddress(contractAddr.String()), estimatedEVMValue, c.cfg.BlockScanner.MaxGasLimit, gasRate, txData)
	estimatedGas, err := c.evmScanner.ethRpc.EstimateGas(fromAddr.String(), createdTx)
	if err != nil {
		// in an edge case that vault doesn't have enough fund to fulfill an outbound transaction , it will fail to estimate gas
		// the returned error is `execution reverted`
		// when this fail , chain client should skip the outbound and move on to the next. The network will reschedule the outbound
		// after 300 blocks
		c.logger.Err(err).Msg("fail to estimate gas")
		return nil, nil
	}

	scheduledMaxFee := big.NewInt(0)
	for _, coin := range txOutItem.MaxGas {
		scheduledMaxFee.Add(scheduledMaxFee, convertThorchainAmountToWei(coin.Amount.BigInt()))
	}

	if txOutItem.Aggregator != "" {
		var gasLimitForAggregator uint64
		// if the estimate gas is over the max, abort and let thornode reschedule for now
		if estimatedGas > MaxContractGas {
			c.logger.Warn().
				Stringer("in_hash", txOutItem.InHash).
				Uint64("estimated_gas", estimatedGas).
				Uint64("aggregator_gas_limit", gasLimitForAggregator).
				Msg("swap out gas limit exceeded, aborting to let thornode reschedule")
			return nil, nil
		}

		// set limit to aggregator gas limit
		estimatedGas = MaxContractGas

		// aggregator swap outs currently ignore max gas, but abort if 10x over for safety
		//
		// TODO: Update thornode to take aggregator gas limit into consideration and set a
		// max gas that should be respected.
		scheduledMaxFee = scheduledMaxFee.Mul(scheduledMaxFee, big.NewInt(10))
	}

	// if over max scheduled gas, abort and let thornode reschedule
	estimatedFee := big.NewInt(int64(estimatedGas) * gasRate.Int64())
	if scheduledMaxFee.Cmp(estimatedFee) < 0 {
		c.logger.Warn().
			Stringer("in_hash", txOutItem.InHash).
			Str("estimated_fee", estimatedFee.String()).
			Str("scheduled_max_fee", scheduledMaxFee.String()).
			Msg("max gas exceeded, aborting to let mayanode reschedule")
		return nil, nil
	}

	createdTx = etypes.NewTransaction(
		nonce, ecommon.HexToAddress(contractAddr.String()), evmValue, estimatedGas, gasRate, txData,
	)

	return createdTx, nil
}

// --------------------------------- sign ---------------------------------

// SignTx returns the signed transaction.
func (c *EVMClient) SignTx(tx stypes.TxOutItem, height int64) ([]byte, []byte, error) {
	if !tx.Chain.Equals(c.cfg.ChainID) {
		return nil, nil, fmt.Errorf("chain %s is not support by evm chain client", tx.Chain)
	}

	if c.signerCacheManager.HasSigned(tx.CacheHash()) {
		c.logger.Info().Interface("tx", tx).Msg("transaction signed before, ignore")
		return nil, nil, nil
	}

	if tx.ToAddress.IsEmpty() {
		return nil, nil, fmt.Errorf("to address is empty")
	}
	if tx.VaultPubKey.IsEmpty() {
		return nil, nil, fmt.Errorf("vault public key is empty")
	}

	memo, err := mem.ParseMemo(common.LatestVersion, tx.Memo)
	if err != nil {
		return nil, nil, fmt.Errorf("fail to parse memo(%s):%w", tx.Memo, err)
	}

	if memo.IsInbound() {
		return nil, nil, fmt.Errorf("inbound memo should not be used for outbound tx")
	}

	if len(tx.Memo) == 0 {
		return nil, nil, fmt.Errorf("can't sign tx when it doesn't have memo")
	}

	// the nonce is stored as the transaction checkpoint, if it is set deserialize it
	// so we only retry with the same nonce to avoid double spend
	var nonce uint64
	if tx.Checkpoint != nil {
		if err = json.Unmarshal(tx.Checkpoint, &nonce); err != nil {
			return nil, nil, fmt.Errorf("fail to unmarshal checkpoint: %w", err)
		}
	} else {
		var fromAddr common.Address
		fromAddr, err = tx.VaultPubKey.GetAddress(c.cfg.ChainID)
		if err != nil {
			return nil, nil, fmt.Errorf("fail to get EVM address for pub key(%s): %w", tx.VaultPubKey, err)
		}
		nonce, err = c.evmScanner.GetNonce(fromAddr.String())
		if err != nil {
			return nil, nil, fmt.Errorf("fail to fetch account(%s) nonce : %w", fromAddr, err)
		}
	}

	// serialize nonce for later
	nonceBytes, err := json.Marshal(nonce)
	if err != nil {
		return nil, nil, fmt.Errorf("fail to marshal nonce: %w", err)
	}

	outboundTx, err := c.buildOutboundTx(tx, memo, nonce)
	if err != nil {
		c.logger.Err(err).Msg("Failed to build outbound tx")
		return nil, nil, err
	}

	// if transaction is nil, abort to allow thornode reschedule
	if outboundTx == nil {
		return nil, nil, nil
	}

	rawTx, err := c.sign(outboundTx, tx.VaultPubKey, height, tx)
	if err != nil || len(rawTx) == 0 {
		return nil, nonceBytes, fmt.Errorf("fail to sign message: %w", err)
	}

	return rawTx, nil, nil
}

// sign is design to sign a given message with keysign party and keysign wrapper
func (c *EVMClient) sign(tx *etypes.Transaction, poolPubKey common.PubKey, height int64, txOutItem stypes.TxOutItem) ([]byte, error) {
	rawBytes, err := c.kw.Sign(tx, poolPubKey)
	if err == nil && rawBytes != nil {
		return rawBytes, nil
	}
	var keysignError tss.KeysignError
	if errors.As(err, &keysignError) {
		if len(keysignError.Blame.BlameNodes) == 0 {
			// TSS doesn't know which node to blame
			return nil, fmt.Errorf("fail to sign tx: %w", err)
		}
		// key sign error forward the keysign blame to mayachain
		txID, errPostKeysignFail := c.bridge.PostKeysignFailure(keysignError.Blame, height, txOutItem.Memo, txOutItem.Coins, txOutItem.VaultPubKey)
		if errPostKeysignFail != nil {
			return nil, multierror.Append(err, errPostKeysignFail)
		}
		c.logger.Info().Str("tx_id", txID.String()).Msg("post keysign failure to mayachain")
	}
	return nil, fmt.Errorf("fail to sign tx: %w", err)
}

// --------------------------------- broadcast ---------------------------------

// BroadcastTx broadcasts the transaction and returns the transaction hash.
func (c *EVMClient) BroadcastTx(txOutItem stypes.TxOutItem, hexTx []byte) (string, error) {
	// decode the transaction
	tx := &etypes.Transaction{}
	if err := tx.UnmarshalJSON(hexTx); err != nil {
		return "", err
	}
	txID := tx.Hash().String()

	// get context with default timeout
	ctx, cancel := c.getTimeoutContext()
	defer cancel()

	// send the transaction
	if err := c.ethClient.SendTransaction(ctx, tx); err != nil {
		c.logger.Error().Str("txid", txID).Err(err).Msg("failed to send transaction")
		return "", nil
	}
	if err := c.signerCacheManager.SetSigned(txOutItem.CacheHash(), txOutItem.CacheVault(c.GetChain()), txID); err != nil {
		c.logger.Err(err).Msgf("fail to mark tx out item (%+v) as signed", txOutItem)
	}

	blockHeight, err := c.bridge.GetBlockHeight()
	if err != nil {
		c.logger.Err(err).Msg("fail to get current MAYAChain block height")
		// at this point , the tx already broadcast successfully , don't return an error
		// otherwise will cause the same tx to retry
	} else if err := c.AddSignedTxItem(txID, blockHeight, txOutItem.VaultPubKey.String()); err != nil {
		c.logger.Err(err).Str("hash", txID).Msg("fail to add signed tx item")
	}

	return txID, nil
}

// --------------------------------- observe ---------------------------------

// OnObservedTxIn is called when a new observed tx is received.
func (c *EVMClient) OnObservedTxIn(txIn stypes.TxInItem, blockHeight int64) {
	m, err := mem.ParseMemo(common.LatestVersion, txIn.Memo)
	if err != nil {
		// Debug log only as ParseMemo error is expected for MAYAName inbounds.
		c.logger.Debug().Err(err).Str("memo", txIn.Memo).Msg("fail to parse memo")
		return
	}
	if !m.IsOutbound() {
		return
	}
	if m.GetTxID().IsEmpty() {
		return
	}
	if err = c.signerCacheManager.SetSigned(txIn.CacheHash(c.GetChain(), m.GetTxID().String()), txIn.CacheVault(c.GetChain()), txIn.Tx); err != nil {
		c.logger.Err(err).Msg("fail to update signer cache")
	}
}

// GetConfirmationCount returns the confirmation count for the given tx.
func (c *EVMClient) GetConfirmationCount(txIn stypes.TxIn) int64 {
	switch c.cfg.ChainID {
	case common.AVAXChain, common.ARBChain: // instant finality
		return 0
	default:
		c.logger.Fatal().Msgf("unsupported chain: %s", c.cfg.ChainID)
		return 0
	}
}

// ConfirmationCountReady returns true if the confirmation count is ready.
func (c *EVMClient) ConfirmationCountReady(txIn stypes.TxIn) bool {
	switch c.cfg.ChainID {
	case common.AVAXChain, common.ARBChain: // instant finality
		return true
	default:
		c.logger.Fatal().Msgf("unsupported chain: %s", c.cfg.ChainID)
		return false
	}
}

// --------------------------------- solvency ---------------------------------

// ReportSolvency reports solvency once per configured solvency blocks.
func (c *EVMClient) ReportSolvency(height int64) error {
	if !c.ShouldReportSolvency(height) {
		return nil
	}

	// skip reporting solvency if the block scanner is unhealthy and we are synced
	if !c.IsBlockScannerHealthy() && height == c.evmScanner.currentBlockHeight {
		return nil
	}

	asgardVaults, err := c.bridge.GetAsgards()
	if err != nil {
		return fmt.Errorf("fail to get asgards, err: %w", err)
	}

	currentGasFee := cosmos.NewUint(3 * c.cfg.BlockScanner.MaxGasLimit * c.evmScanner.lastReportedGasPrice)

	for _, asgard := range asgardVaults {
		acct, err := c.GetAccount(asgard.PubKey, new(big.Int).SetInt64(height))
		if err != nil {
			c.logger.Err(err).Msg("fail to get account balance")
			continue
		}

		// skip reporting solvency if the account is solvent and block scanner is healthy
		solvent := runners.IsVaultSolvent(acct, asgard, currentGasFee)
		if solvent && c.IsBlockScannerHealthy() {
			continue
		}
		c.logger.Info().
			Stringer("asgard", asgard.PubKey).
			Interface("coins", acct.Coins).
			Bool("solvent", solvent).
			Msg("reporting solvency")

		select {
		case c.globalSolvencyQueue <- stypes.Solvency{
			Height: height,
			Chain:  c.cfg.ChainID,
			PubKey: asgard.PubKey,
			Coins:  acct.Coins,
		}:
		case <-time.After(constants.MayachainBlockTime):
			c.logger.Info().Msg("fail to send solvency info to mayachain, timeout")
		}
	}
	c.lastSolvencyCheckHeight = height
	return nil
}

// ShouldReportSolvency returns true if the given height is a solvency report height.
func (c *EVMClient) ShouldReportSolvency(height int64) bool {
	return height%c.cfg.SolvencyBlocks == 0
}

// --------------------------------- helpers ---------------------------------

func (c *EVMClient) getTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), c.cfg.BlockScanner.HTTPRequestTimeout)
}
