package signer

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tssp "gitlab.com/thorchain/tss/go-tss/tss"

	"gitlab.com/mayachain/mayanode/bifrost/blockscanner"
	"gitlab.com/mayachain/mayanode/bifrost/mayaclient"
	"gitlab.com/mayachain/mayanode/bifrost/mayaclient/types"
	"gitlab.com/mayachain/mayanode/bifrost/metrics"
	"gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients"
	"gitlab.com/mayachain/mayanode/bifrost/pubkeymanager"
	"gitlab.com/mayachain/mayanode/bifrost/tss"
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/config"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain"
	ttypes "gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// Signer will pull the tx out from thorchain and then forward it to chain
type Signer struct {
	logger                zerolog.Logger
	cfg                   config.BifrostSignerConfiguration
	wg                    *sync.WaitGroup
	mayachainBridge       mayaclient.MayachainBridge
	stopChan              chan struct{}
	blockScanner          *blockscanner.BlockScanner
	thorchainBlockScanner *ThorchainBlockScan
	chains                map[common.Chain]chainclients.ChainClient
	storage               SignerStorage
	m                     *metrics.Metrics
	errCounter            *prometheus.CounterVec
	tssKeygen             *tss.KeyGen
	pubkeyMgr             pubkeymanager.PubKeyValidator
	constantsProvider     *ConstantsProvider
	localPubKey           common.PubKey
	tssKeysignMetricMgr   *metrics.TssKeysignMetricMgr
}

// NewSigner create a new instance of signer
func NewSigner(cfg config.BifrostSignerConfiguration,
	mayachainBridge mayaclient.MayachainBridge,
	thorKeys *mayaclient.Keys,
	pubkeyMgr pubkeymanager.PubKeyValidator,
	tssServer *tssp.TssServer,
	chains map[common.Chain]chainclients.ChainClient,
	m *metrics.Metrics,
	tssKeysignMetricMgr *metrics.TssKeysignMetricMgr,
) (*Signer, error) {
	storage, err := NewSignerStore(cfg.SignerDbPath, mayachainBridge.GetConfig().SignerPasswd)
	if err != nil {
		return nil, fmt.Errorf("fail to create thorchain scan storage: %w", err)
	}
	if tssKeysignMetricMgr == nil {
		return nil, fmt.Errorf("fail to create signer , tss keysign metric manager is nil")
	}
	var na *ttypes.NodeAccount
	for i := 0; i < 300; i++ { // wait for 5 min before timing out
		na, err = mayachainBridge.GetNodeAccount(thorKeys.GetSignerInfo().GetAddress().String())
		if err != nil {
			return nil, fmt.Errorf("fail to get node account from thorchain,err:%w", err)
		}

		if !na.PubKeySet.Secp256k1.IsEmpty() {
			break
		}
		time.Sleep(constants.MayachainBlockTime)
		fmt.Println("Waiting for node account to be registered...")
	}
	for _, item := range na.GetSignerMembership() {
		pubkeyMgr.AddPubKey(item, true)
	}
	if na.PubKeySet.Secp256k1.IsEmpty() {
		return nil, fmt.Errorf("unable to find pubkey for this node account. exiting... ")
	}
	pubkeyMgr.AddNodePubKey(na.PubKeySet.Secp256k1)

	cfg.BlockScanner.ChainID = common.BASEChain // hard code to basechain

	// Create pubkey manager and add our private key (Yggdrasil pubkey)
	thorchainBlockScanner, err := NewThorchainBlockScan(cfg.BlockScanner, storage, mayachainBridge, m, pubkeyMgr)
	if err != nil {
		return nil, fmt.Errorf("fail to create thorchain block scan: %w", err)
	}

	blockScanner, err := blockscanner.NewBlockScanner(cfg.BlockScanner, storage, m, mayachainBridge, thorchainBlockScanner)
	if err != nil {
		return nil, fmt.Errorf("fail to create block scanner: %w", err)
	}

	kg, err := tss.NewTssKeyGen(thorKeys, tssServer, mayachainBridge)
	if err != nil {
		return nil, fmt.Errorf("fail to create Tss Key gen,err:%w", err)
	}
	constantProvider := NewConstantsProvider(mayachainBridge)
	return &Signer{
		logger:                log.With().Str("module", "signer").Logger(),
		cfg:                   cfg,
		wg:                    &sync.WaitGroup{},
		stopChan:              make(chan struct{}),
		blockScanner:          blockScanner,
		thorchainBlockScanner: thorchainBlockScanner,
		chains:                chains,
		m:                     m,
		storage:               storage,
		errCounter:            m.GetCounterVec(metrics.SignerError),
		pubkeyMgr:             pubkeyMgr,
		mayachainBridge:       mayachainBridge,
		tssKeygen:             kg,
		constantsProvider:     constantProvider,
		localPubKey:           na.PubKeySet.Secp256k1,
		tssKeysignMetricMgr:   tssKeysignMetricMgr,
	}, nil
}

func (s *Signer) getChain(chainID common.Chain) (chainclients.ChainClient, error) {
	chain, ok := s.chains[chainID]
	if !ok {
		s.logger.Debug().Str("chain", chainID.String()).Msg("is not supported yet")
		return nil, errors.New("not supported")
	}
	return chain, nil
}

// Start signer process
func (s *Signer) Start() error {
	s.wg.Add(1)
	go s.processTxnOut(s.thorchainBlockScanner.GetTxOutMessages(), 1)

	s.wg.Add(1)
	go s.processKeygen(s.thorchainBlockScanner.GetKeygenMessages())

	s.wg.Add(1)
	go s.signTransactions()

	s.blockScanner.Start(nil)
	return nil
}

func (s *Signer) shouldSign(tx types.TxOutItem) bool {
	return s.pubkeyMgr.HasPubKey(tx.VaultPubKey)
}

// signTransactions - looks for work to do by getting a list of all unsigned
// transactions stored in the storage
func (s *Signer) signTransactions() {
	s.logger.Info().Msg("start to sign transactions")
	defer s.logger.Info().Msg("stop to sign transactions")
	defer s.wg.Done()
	for {
		select {
		case <-s.stopChan:
			return
		default:
			// When BASEChain is catching up , bifrost might get stale data from mayanode , thus it shall pause signing
			catchingUp, err := s.mayachainBridge.IsCatchingUp()
			if err != nil {
				s.logger.Error().Err(err).Msg("fail to get mayachain sync status")
				time.Sleep(constants.MayachainBlockTime)
				break // this will break select
			}
			if !catchingUp {
				s.processTransactions()
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func runWithContext(ctx context.Context, fn func() ([]byte, error)) ([]byte, error) {
	ch := make(chan error, 1)
	var checkpoint []byte
	go func() {
		var err error
		checkpoint, err = fn()
		ch <- err
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-ch:
		return checkpoint, err
	}
}

func (s *Signer) processTransactions() {
	wg := &sync.WaitGroup{}
	for _, items := range s.storage.OrderedLists() {
		wg.Add(1)

		go func(items []TxOutStoreItem) {
			defer wg.Done()

			// if any tx out items are in broadcast or round 7 failure retry, only proceed with those
			retryItems := []TxOutStoreItem{}
			for _, item := range items {
				if item.Round7Retry || len(item.SignedTx) > 0 {
					retryItems = append(retryItems, item)
				}
			}
			if len(retryItems) > 0 {
				s.logger.Info().Msgf("found %d retry items", len(retryItems))
				items = retryItems
			}
			if len(retryItems) > 1 {
				s.logger.Error().Msgf("found %d retry items, there should only be one", len(retryItems))
			}

			for i, item := range items {
				select {
				case <-s.stopChan:
					return
				default:
					if item.Status == TxSpent { // don't rebroadcast spent transactions
						continue
					}

					s.logger.Info().Int("num", i).Int64("height", item.Height).Int("status", int(item.Status)).Interface("tx", item.TxOutItem).Msgf("Signing transaction")
					// a single keysign should not take longer than 5 minutes , regardless TSS or local
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
					if checkpoint, err := runWithContext(ctx, func() ([]byte, error) {
						return s.signAndBroadcast(item)
					}); err != nil {
						// mark the txout on round 7 failure to block other txs for the chain / pubkey
						ksErr := tss.KeysignError{}
						if errors.As(err, &ksErr) && ksErr.IsRound7() {
							s.logger.Error().Err(err).Interface("tx", item.TxOutItem).Msg("round 7 signing error")
							item.Round7Retry = true
							item.Checkpoint = checkpoint
							if storeErr := s.storage.Set(item); storeErr != nil {
								s.logger.Error().Err(storeErr).Msg("fail to update tx out store item with round 7 retry")
							}
						}

						if errors.Is(err, context.DeadlineExceeded) {
							panic(fmt.Errorf("tx out item: %+v , keysign timeout : %w", item.TxOutItem, err))
						}
						s.logger.Error().Err(err).Msg("fail to sign and broadcast tx out store item")
						cancel()
						return
						// The 'item' for loop should not be items[0],
						// because problems which return 'nil, nil' should be skipped over instead of blocking others.
						// When signAndBroadcast returns an error (such as from a keysign timeout),
						// a 'return' and not a 'continue' should be used so that nodes can all restart the list,
						// for when the keysign failure was from a loss of list synchrony.
						// Otherwise, out-of-sync lists would cycle one timeout at a time, maybe never resynchronising.
					}
					cancel()

					// We have a successful broadcast! Remove the item from our store
					if err := s.storage.Remove(item); err != nil {
						s.logger.Error().Err(err).Msg("fail to update tx out store item")
					}
				}
			}
		}(items)
	}
	wg.Wait()
}

// processTxnOut processes outbound TxOuts and save them to storage
func (s *Signer) processTxnOut(ch <-chan types.TxOut, idx int) {
	s.logger.Info().Int("idx", idx).Msg("start to process tx out")
	defer s.logger.Info().Int("idx", idx).Msg("stop to process tx out")
	defer s.wg.Done()
	for {
		select {
		case <-s.stopChan:
			return
		case txOut, more := <-ch:
			if !more {
				return
			}
			s.logger.Info().Msgf("Received a TxOut Array of %v from the Thorchain", txOut)
			items := make([]TxOutStoreItem, 0, len(txOut.TxArray))

			for i, tx := range txOut.TxArray {
				items = append(items, NewTxOutStoreItem(txOut.Height, tx.TxOutItem(), int64(i)))
			}
			if err := s.storage.Batch(items); err != nil {
				s.logger.Error().Err(err).Msg("fail to save tx out items to storage")
			}
		}
	}
}

func (s *Signer) processKeygen(ch <-chan ttypes.KeygenBlock) {
	s.logger.Info().Msg("start to process keygen")
	defer s.logger.Info().Msg("stop to process keygen")
	defer s.wg.Done()
	for {
		select {
		case <-s.stopChan:
			return
		case keygenBlock, more := <-ch:
			if !more {
				return
			}
			s.logger.Info().Interface("keygenBlock", keygenBlock).Msg("received a keygen block from mayachain")
			s.processKeygenBlock(keygenBlock)
		}
	}
}

func (s *Signer) scheduleRetry(keygenBlock ttypes.KeygenBlock) bool {
	churnRetryInterval, err := s.mayachainBridge.GetMimir(constants.ChurnRetryInterval.String())
	if err != nil {
		s.logger.Error().Err(err).Msg("fail to get churn retry mimir")
		return false
	}
	if churnRetryInterval <= 0 {
		churnRetryInterval = constants.NewConstantValue010().GetInt64Value(constants.ChurnRetryInterval)
	}
	keygenRetryInterval, err := s.mayachainBridge.GetMimir(constants.KeygenRetryInterval.String())
	if err != nil {
		s.logger.Error().Err(err).Msg("fail to get keygen retries mimir")
		return false
	}
	if keygenRetryInterval <= 0 {
		return false
	}

	// sanity check the retry interval is at least 1.5x the timeout
	retryIntervalDuration := time.Duration(keygenRetryInterval) * constants.MayachainBlockTime
	if retryIntervalDuration <= s.cfg.KeygenTimeout*3/2 {
		s.logger.Error().
			Stringer("retryInterval", retryIntervalDuration).
			Stringer("keygenTimeout", s.cfg.KeygenTimeout).
			Msg("retry interval too short")
		return false
	}

	height, err := s.mayachainBridge.GetBlockHeight()
	if err != nil {
		s.logger.Error().Err(err).Msg("fail to get last chain height")
		return false
	}

	// target retry height is the next keygen retry interval over the keygen block height
	targetRetryHeight := (keygenRetryInterval - ((height - keygenBlock.Height) % keygenRetryInterval)) + height

	// skip trying close to churn retry
	if targetRetryHeight > keygenBlock.Height+churnRetryInterval-keygenRetryInterval {
		return false
	}

	go func() {
		// every block, try to start processing again
		for {
			time.Sleep(constants.MayachainBlockTime)
			height, err := s.mayachainBridge.GetBlockHeight()
			if err != nil {
				s.logger.Error().Err(err).Msg("fail to get last chain height")
			}
			if height >= targetRetryHeight {
				s.logger.Info().
					Interface("keygenBlock", keygenBlock).
					Int64("currentHeight", height).
					Msg("retrying keygen")
				s.processKeygenBlock(keygenBlock)
				return
			}
		}
	}()

	s.logger.Info().
		Interface("keygenBlock", keygenBlock).
		Int64("retryHeight", targetRetryHeight).
		Msg("scheduled keygen retry")

	return true
}

func (s *Signer) processKeygenBlock(keygenBlock ttypes.KeygenBlock) {
	s.logger.Info().Interface("keygenBlock", keygenBlock).Msg("processing keygen block")

	// NOTE: in practice there is only one keygen in the keygen block
	for _, keygenReq := range keygenBlock.Keygens {
		var err error
		keygenStart := time.Now()
		pubKey, blame, err := s.tssKeygen.GenerateNewKey(keygenBlock.Height, keygenReq.GetMembers())
		if !blame.IsEmpty() {
			err = fmt.Errorf("reason: %s, nodes %+v", blame.FailReason, blame.BlameNodes)
			s.logger.Error().Err(err).Msg("Blame")
		}
		keygenTime := time.Since(keygenStart).Milliseconds()

		if err != nil {
			s.errCounter.WithLabelValues("fail_to_keygen_pubkey", "").Inc()
			s.logger.Error().Err(err).Msg("fail to generate new pubkey")
		}

		// re-enqueue the keygen block to retry if we failed to generate a key
		if pubKey.Secp256k1.IsEmpty() {
			if s.scheduleRetry(keygenBlock) {
				return
			}
			s.logger.Error().Interface("keygenBlock", keygenBlock).Msg("done with keygen retries")
		}

		if err := s.sendKeygenToMayachain(keygenBlock.Height, pubKey.Secp256k1, blame, keygenReq.GetMembers(), keygenReq.Type, keygenTime); err != nil {
			s.errCounter.WithLabelValues("fail_to_broadcast_keygen", "").Inc()
			s.logger.Error().Err(err).Msg("fail to broadcast keygen")
		}

		// monitor the new pubkey and any new members
		if !pubKey.Secp256k1.IsEmpty() {
			s.pubkeyMgr.AddPubKey(pubKey.Secp256k1, true)
		}
		for _, pk := range keygenReq.GetMembers() {
			s.pubkeyMgr.AddPubKey(pk, false)
		}
	}
}

func (s *Signer) sendKeygenToMayachain(height int64, poolPk common.PubKey, blame ttypes.Blame, input common.PubKeys, keygenType ttypes.KeygenType, keygenTime int64) error {
	// collect supported chains in the configuration
	chains := common.Chains{
		common.BASEChain,
	}
	for name, chain := range s.chains {
		if !chain.GetConfig().OptToRetire {
			chains = append(chains, name)
		}
	}

	keygenMsg, err := s.mayachainBridge.GetKeygenStdTx(poolPk, blame, input, keygenType, chains, height, keygenTime)
	if err != nil {
		return fmt.Errorf("fail to get keygen id: %w", err)
	}
	strHeight := strconv.FormatInt(height, 10)

	bf := backoff.NewExponentialBackOff()
	bf.MaxElapsedTime = constants.MayachainBlockTime
	return backoff.Retry(func() error {
		txID, err := s.mayachainBridge.Broadcast(keygenMsg)
		if err != nil {
			s.logger.Warn().Err(err).Msg("fail to send keygen tx to mayachain")
			s.errCounter.WithLabelValues("fail_to_send_to_mayachain", strHeight).Inc()
			return fmt.Errorf("fail to send the tx to mayachain: %w", err)
		}
		s.logger.Info().Stringer("txid", txID).Int64("block", height).Msg("sent keygen tx to mayachain")
		return nil
	}, bf)
}

// signAndBroadcast will sign the tx and broadcast it to the corresponding chain. On
// SignTx error for the chain client, if we receive checkpoint bytes we also return them
// with the error so they can be set on the TxOutStoreItem and re-used on a subsequent
// retry to avoid double spend.
func (s *Signer) signAndBroadcast(item TxOutStoreItem) ([]byte, error) {
	height := item.Height
	tx := item.TxOutItem

	// set the checkpoint on the tx out item if it was stored
	if item.Checkpoint != nil {
		tx.Checkpoint = item.Checkpoint
	}

	blockHeight, err := s.mayachainBridge.GetBlockHeight()
	if err != nil {
		s.logger.Error().Err(err).Msgf("fail to get block height")
		return nil, err
	}
	var signingTransactionPeriod int64
	signingTransactionPeriod, err = s.constantsProvider.GetInt64Value(blockHeight, constants.SigningTransactionPeriod)
	s.logger.Debug().Msgf("signing transaction period:%d", signingTransactionPeriod)
	if err != nil {
		s.logger.Error().Err(err).Msgf("fail to get constant value for(%s)", constants.SigningTransactionPeriod)
		return nil, err
	}
	// rounds up to nearth 100th, then minuses signingTxPeriod. This is in an
	// effort for multi-bifrost nodes to get deterministic consensus on which
	// transaction to sign next. If we didn't round up, which transaction to
	// sign would change every 5 seconds. And with 20 sec party timeouts, luck
	// of execution time will determine if consensus is reached. Instead, we
	// have the same transaction selected for a longer period of time, making
	// it easier for the nodes to all select the same transaction, even if they
	// don't execute at the same time.
	if ((blockHeight/100*100)+100)-(signingTransactionPeriod) > height {
		s.logger.Error().Msgf("tx was created at block height(%d), now it is (%d), it is older than (%d) blocks , skip it ", height, blockHeight, signingTransactionPeriod)
		return nil, nil
	}
	var chain chainclients.ChainClient
	chain, err = s.getChain(tx.Chain)
	if err != nil {
		s.logger.Error().Err(err).Msgf("not supported %s", tx.Chain.String())
		return nil, err
	}
	mimirKey := "HALTSIGNING"
	haltSigningGlobalMimir, err := s.mayachainBridge.GetMimir(mimirKey)
	if err != nil {
		s.logger.Err(err).Msgf("fail to get %s", mimirKey)
		return nil, err
	}
	if haltSigningGlobalMimir > 0 {
		s.logger.Info().Msg("signing has been halted globally")
		return nil, nil
	}
	mimirKey = fmt.Sprintf("HALTSIGNING%s", tx.Chain)
	haltSigningMimir, err := s.mayachainBridge.GetMimir(mimirKey)
	if err != nil {
		s.logger.Err(err).Msgf("fail to get %s", mimirKey)
		return nil, err
	}
	if haltSigningMimir > 0 {
		s.logger.Info().Msgf("signing for %s is halted", tx.Chain)
		return nil, nil
	}
	if !s.shouldSign(tx) {
		s.logger.Info().Str("signer_address", chain.GetAddress(tx.VaultPubKey)).Msg("different pool address, ignore")
		return nil, nil
	}

	if len(tx.ToAddress) == 0 {
		s.logger.Info().Msg("To address is empty, MAYANode don't know where to send the fund , ignore")
		return nil, nil // return nil and discard item
	}

	// don't sign if the block scanner is unhealthy. This is because the
	// network may not be able to detect the outbound transaction, and
	// therefore reschedule the transaction to another signer. In a disaster
	// scenario, the network could broadcast a transaction several times,
	// bleeding funds.
	if !chain.IsBlockScannerHealthy() {
		return nil, fmt.Errorf("the block scanner for chain %s is unhealthy, not signing transactions due to it", chain.GetChain())
	}

	// Check if we're sending all funds back , given we don't have memo in txoutitem anymore, so it rely on the coins field to be empty
	// In this scenario, we should chose the coins to send ourselves
	if tx.Coins.IsEmpty() {
		tx, err = s.handleYggReturn(height, tx)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed to handle yggdrasil return")
			return nil, err
		}
	}

	start := time.Now()
	defer func() {
		s.m.GetHistograms(metrics.SignAndBroadcastDuration(chain.GetChain())).Observe(time.Since(start).Seconds())
	}()

	if !tx.OutHash.IsEmpty() {
		s.logger.Info().Str("OutHash", tx.OutHash.String()).Msg("tx had been sent out before")
		return nil, nil // return nil and discard item
	}

	// We get the keysign object from thorchain again to ensure it hasn't
	// been signed already, and we can skip. This helps us not get stuck on
	// a task that we'll never sign, because 2/3rds already has and will
	// never be available to sign again.
	txOut, err := s.mayachainBridge.GetKeysign(height, tx.VaultPubKey.String())
	if err != nil {
		s.logger.Error().Err(err).Msg("fail to get keysign items")
		return nil, err
	}
	for _, txArray := range txOut.TxArray {
		if txArray.TxOutItem().Equals(tx) && !txArray.OutHash.IsEmpty() {
			// already been signed, we can skip it
			s.logger.Info().Str("tx_id", tx.OutHash.String()).Msgf("already signed. skipping...")
			return nil, nil
		}
	}

	// If SignedTx is set, we already signed and should only retry broadcast.
	var signedTx, checkpoint []byte
	var elapse time.Duration
	if len(item.SignedTx) > 0 {
		s.logger.Info().Str("memo", tx.Memo).Msg("retrying broadcast of already signed tx")
		signedTx = item.SignedTx
	} else {
		startKeySign := time.Now()
		signedTx, checkpoint, err = chain.SignTx(tx, height)
		if err != nil {
			s.logger.Error().Err(err).Msg("fail to sign tx")
			return checkpoint, err
		}
		elapse = time.Since(startKeySign)
	}

	// looks like the transaction is already signed
	if len(signedTx) == 0 {
		s.logger.Warn().Msgf("signed transaction is empty")
		return nil, nil
	}
	var hash string
	hash, err = chain.BroadcastTx(tx, signedTx)
	if err != nil {
		s.logger.Error().Err(err).Str("memo", tx.Memo).Msg("fail to broadcast tx to chain")

		// store the signed tx for the next retry
		item.SignedTx = signedTx
		if storeErr := s.storage.Set(item); storeErr != nil {
			s.logger.Error().Err(storeErr).Msg("fail to update tx out store item with signed tx")
		}

		return nil, err
	}

	if s.isTssKeysign(tx.VaultPubKey) {
		s.tssKeysignMetricMgr.SetTssKeysignMetric(hash, elapse.Milliseconds())
	}

	return nil, nil
}

func (s *Signer) handleYggReturn(height int64, tx types.TxOutItem) (types.TxOutItem, error) {
	chain, err := s.getChain(tx.Chain)
	if err != nil {
		s.logger.Error().Err(err).Msgf("not supported %s", tx.Chain.String())
		return tx, err
	}
	isValid, _ := s.pubkeyMgr.IsValidPoolAddress(tx.ToAddress.String(), tx.Chain)
	if !isValid {
		errInvalidPool := fmt.Errorf("yggdrasil return should return to a valid pool address,%s is not valid", tx.ToAddress.String())
		s.logger.Error().Err(errInvalidPool).Msg("invalid yggdrasil return address")
		return tx, errInvalidPool
	}
	// it is important to set the memo field to `yggdrasil-` , thus chain client can use it to decide leave some gas coin behind to pay the fees
	tx.Memo = mayachain.NewYggdrasilReturn(height).String()
	var acct common.Account
	acct, err = chain.GetAccount(tx.VaultPubKey, nil)
	if err != nil {
		return tx, fmt.Errorf("fail to get chain account info: %w", err)
	}
	tx.Coins = make(common.Coins, 0)
	for _, coin := range acct.Coins {
		asset, err := common.NewAsset(coin.Asset.String())
		asset.Chain = tx.Chain
		if err != nil {
			return tx, fmt.Errorf("fail to parse asset: %w", err)
		}
		if coin.Amount.Uint64() > 0 {
			amount := coin.Amount
			tx.Coins = append(tx.Coins, common.NewCoin(asset, amount))
		}
	}
	// Yggdrasil return should pay whatever gas is necessary
	tx.MaxGas = common.Gas{}
	return tx, nil
}

func (s *Signer) isTssKeysign(pubKey common.PubKey) bool {
	return !s.localPubKey.Equals(pubKey)
}

// Stop the signer process
func (s *Signer) Stop() error {
	s.logger.Info().Msg("receive request to stop signer")
	defer s.logger.Info().Msg("signer stopped successfully")
	close(s.stopChan)
	s.wg.Wait()
	if err := s.m.Stop(); err != nil {
		s.logger.Error().Err(err).Msg("fail to stop metric server")
	}
	s.blockScanner.Stop()
	return s.storage.Close()
}
