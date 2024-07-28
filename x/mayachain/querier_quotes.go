package mayachain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog"
	abci "github.com/tendermint/tendermint/abci/types"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/log"
	openapi "gitlab.com/mayachain/mayanode/openapi/gen"
	mem "gitlab.com/mayachain/mayanode/x/mayachain/memo"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// -------------------------------------------------------------------------------------
// Config
// -------------------------------------------------------------------------------------

const (
	// heightParam    = "height"
	fromAssetParam = "from_asset"
	toAssetParam   = "to_asset"
	assetParam     = "asset"
	addressParam   = "address"
	// loanOwnerParam           = "loan_owner"
	withdrawBasisPointsParam = "withdraw_bps"
	amountParam              = "amount"
	// repayBpsParam             = "repay_bps"
	destinationParam          = "destination"
	toleranceBasisPointsParam = "tolerance_bps"
	affiliateParam            = "affiliate"
	affiliateBpsParam         = "affiliate_bps"
	minOutParam               = "min_out"
	intervalParam             = "streaming_interval"
	quantityParam             = "streaming_quantity"
	// refundAddressParam = "refund_address"

	quoteWarning    = "Do not cache this response. Do not send funds after the expiry."
	quoteExpiration = 15 * time.Minute
)

var nullLogger = &log.TendermintLogWrapper{Logger: zerolog.New(io.Discard)}

// -------------------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------------------

func quoteErrorResponse(err error) ([]byte, error) {
	return json.Marshal(map[string]string{"error": err.Error()})
}

func quoteParseParams(data []byte) (params url.Values, err error) {
	// parse the query parameters
	u, err := url.ParseRequestURI(string(data))
	if err != nil {
		return nil, fmt.Errorf("bad params: %w", err)
	}

	// error if parameters were not provided
	if len(u.Query()) == 0 {
		return nil, fmt.Errorf("no parameters provided")
	}

	return u.Query(), nil
}

func quoteParseAddress(ctx cosmos.Context, mgr *Mgrs, addrString string, chain common.Chain) (common.Address, error) {
	if addrString == "" {
		return common.NoAddress, nil
	}

	// attempt to parse a raw address
	addr, err := common.NewAddress(addrString)
	if err == nil {
		return addr, nil
	}

	// attempt to lookup a mayaname address
	name, err := mgr.Keeper().GetMAYAName(ctx, addrString)
	if err != nil {
		return common.NoAddress, fmt.Errorf("unable to parse address: %w", err)
	}

	// find the address for the correct chain
	for _, alias := range name.Aliases {
		if alias.Chain.Equals(chain) {
			return alias.Address, nil
		}
	}

	return common.NoAddress, fmt.Errorf("no mayaname alias for chain %s", chain)
}

func quoteHandleAffiliate(ctx cosmos.Context, mgr *Mgrs, params url.Values, amount sdk.Uint) (affiliate common.Address, memo string, bps, newAmount, affiliateAmt sdk.Uint, err error) {
	// parse affiliate
	affAmt := cosmos.ZeroUint()
	memo = "" // do not resolve mayaname for the memo
	if len(params[affiliateParam]) > 0 {
		affiliate, err = quoteParseAddress(ctx, mgr, params[affiliateParam][0], common.BASEChain)
		if err != nil {
			err = fmt.Errorf("bad affiliate address: %w", err)
			return
		}
		memo = params[affiliateParam][0]
	}

	// parse affiliate fee
	bps = sdk.NewUint(0)
	if len(params[affiliateBpsParam]) > 0 {
		bps, err = sdk.ParseUint(params[affiliateBpsParam][0])
		if err != nil {
			err = fmt.Errorf("bad affiliate fee: %w", err)
			return
		}
	}

	// verify affiliate fee
	if bps.GT(sdk.NewUint(10000)) {
		err = fmt.Errorf("affiliate fee must be less than 10000 bps")
		return
	}

	// compute the new swap amount if an affiliate fee will be taken first
	if affiliate != common.NoAddress && !bps.IsZero() {
		// calculate the affiliate amount
		affAmt = common.GetSafeShare(
			bps,
			cosmos.NewUint(10000),
			amount,
		)

		// affiliate fee modifies amount at observation before the swap
		amount = amount.Sub(affAmt)
	}

	return affiliate, memo, bps, amount, affAmt, nil
}

func quoteReverseFuzzyAsset(ctx cosmos.Context, mgr *Mgrs, asset common.Asset) (common.Asset, error) {
	// get all pools
	pools, err := mgr.Keeper().GetPools(ctx)
	if err != nil {
		return asset, fmt.Errorf("failed to get pools: %w", err)
	}

	// return the asset if no symbol to shorten
	aSplit := strings.Split(asset.Symbol.String(), "-")
	if len(aSplit) == 1 {
		return asset, nil
	}

	// find all other assets that match the chain and ticker
	// (without exactly matching the symbol)
	addressMatches := []string{}
	for _, p := range pools {
		if p.IsAvailable() && !p.IsEmpty() && !p.Asset.IsVaultAsset() &&
			!p.Asset.Symbol.Equals(asset.Symbol) &&
			p.Asset.Chain.Equals(asset.Chain) && p.Asset.Ticker.Equals(asset.Ticker) {
			pSplit := strings.Split(p.Asset.Symbol.String(), "-")
			if len(pSplit) != 2 {
				return asset, fmt.Errorf("ambiguous match: %s", p.Asset.Symbol)
			}
			addressMatches = append(addressMatches, pSplit[1])
		}
	}

	if len(addressMatches) == 0 { // if only one match, drop the address
		asset.Symbol = common.Symbol(asset.Ticker)
	} else { // find the shortest unique suffix of the asset symbol
		address := aSplit[1]

		for i := len(address) - 1; i > 0; i-- {
			if !hasSuffixMatch(address[i:], addressMatches) {
				asset.Symbol = common.Symbol(
					fmt.Sprintf("%s-%s", asset.Ticker, address[i:]),
				)
				break
			}
		}
	}

	return asset, nil
}

func hasSuffixMatch(suffix string, values []string) bool {
	for _, value := range values {
		if strings.HasSuffix(value, suffix) {
			return true
		}
	}
	return false
}

// NOTE: streamingQuantity > 0 is a precondition.
func quoteSimulateSwap(ctx cosmos.Context, mgr *Mgrs, amount sdk.Uint, msg *MsgSwap, streamingQuantity uint64) (
	res *openapi.QuoteSwapResponse, emitAmount, outboundFeeAmount sdk.Uint, err error,
) {
	// should be unreachable
	if streamingQuantity == 0 {
		return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("streaming quantity must be greater than zero")
	}

	msg.Tx.Coins[0].Amount = msg.Tx.Coins[0].Amount.QuoUint64(streamingQuantity)

	// if the generated memo is too long for the source chain send error
	maxMemoLength := msg.Tx.Coins[0].Asset.Chain.MaxMemoLength()
	if !msg.Tx.Coins[0].Asset.Synth && len(msg.Tx.Memo) > maxMemoLength {
		return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("generated memo too long for source chain")
	}

	// use the first active node account as the signer
	nodeAccounts, err := mgr.Keeper().ListActiveValidators(ctx)
	if err != nil {
		return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("no active node accounts: %w", err)
	}
	msg.Signer = nodeAccounts[0].NodeAddress

	// simulate the swap
	events, err := simulateInternal(ctx, mgr, msg)
	if err != nil {
		return nil, sdk.ZeroUint(), sdk.ZeroUint(), err
	}

	// extract events
	var swaps []map[string]string
	var fee map[string]string
	for _, e := range events {
		switch e.Type {
		case "swap":
			swaps = append(swaps, eventMap(e))
		case "fee":
			fee = eventMap(e)
		}
	}
	finalSwap := swaps[len(swaps)-1]

	// parse outbound fee from event (except on trade assets with no outbound fee)
	outboundFeeCoin, err := common.ParseCoin(fee["coins"])
	if err != nil {
		return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("unable to parse outbound fee coin: %w", err)
	}
	outboundFeeAmount = outboundFeeCoin.Amount

	// parse outbound amount from event
	emitCoin, err := common.ParseCoin(finalSwap["emit_asset"])
	if err != nil {
		return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("unable to parse emit coin: %w", err)
	}
	emitAmount = emitCoin.Amount.MulUint64(streamingQuantity)

	// sum the liquidity fees and convert to target asset
	liquidityFee := sdk.ZeroUint()
	for _, s := range swaps {
		liquidityFee = liquidityFee.Add(sdk.NewUintFromString(s["liquidity_fee_in_cacao"]))
	}
	var targetPool types.Pool
	if !msg.TargetAsset.IsNativeBase() {
		targetPool, err = mgr.Keeper().GetPool(ctx, msg.TargetAsset.GetLayer1Asset())
		if err != nil {
			return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("unable to get pool: %w", err)
		}
		liquidityFee = targetPool.RuneValueInAsset(liquidityFee)
	}
	liquidityFee = liquidityFee.MulUint64(streamingQuantity)

	// approximate the affiliate fee in the target asset
	affiliateFee := sdk.ZeroUint()
	if msg.AffiliateAddress != common.NoAddress && !msg.AffiliateBasisPoints.IsZero() {
		inAsset := msg.Tx.Coins[0].Asset.GetLayer1Asset()
		if !inAsset.IsNativeBase() {
			pool, err := mgr.Keeper().GetPool(ctx, msg.Tx.Coins[0].Asset.GetLayer1Asset())
			if err != nil {
				return nil, sdk.ZeroUint(), sdk.ZeroUint(), fmt.Errorf("unable to get pool: %w", err)
			}
			amount = pool.AssetValueInRune(amount)
		}
		affiliateFee = common.GetUncappedShare(msg.AffiliateBasisPoints, cosmos.NewUint(10_000), amount)
		if !msg.TargetAsset.IsNativeBase() {
			affiliateFee = targetPool.RuneValueInAsset(affiliateFee)
		}
	}

	slipFeeAddedBasisPoints := fetchConfigInt64(ctx, mgr, constants.SlipFeeAddedBasisPoints)

	// compute slip based on emit amount instead of slip in event to handle double swap
	slippageBps := liquidityFee.MulUint64(10000).Quo(emitAmount.Add(liquidityFee))
	slippageBps = slippageBps.AddUint64(uint64(slipFeeAddedBasisPoints))

	// build fees
	totalFees := affiliateFee.Add(liquidityFee).Add(outboundFeeAmount)
	fees := openapi.QuoteFees{
		Asset:       msg.TargetAsset.String(),
		Affiliate:   wrapString(affiliateFee.String()),
		Liquidity:   liquidityFee.String(),
		Outbound:    wrapString(outboundFeeAmount.String()),
		Total:       totalFees.String(),
		SlippageBps: slippageBps.BigInt().Int64(),
		TotalBps:    totalFees.MulUint64(10000).Quo(emitAmount.Add(totalFees)).BigInt().Int64(),
	}

	// build response from simulation result events
	return &openapi.QuoteSwapResponse{
		ExpectedAmountOut: emitAmount.String(),
		Fees:              fees,
	}, emitAmount, outboundFeeAmount, nil
}

func quoteInboundInfo(ctx cosmos.Context, mgr *Mgrs, amount sdk.Uint, chain common.Chain, asset common.Asset) (address, router common.Address, confirmations int64, err error) {
	// If inbound chain is BASEChain there is no inbound address
	if chain.IsBASEChain() {
		address = common.NoAddress
		router = common.NoAddress
	} else {
		// get the most secure vault for inbound
		active, err := mgr.Keeper().GetAsgardVaultsByStatus(ctx, ActiveVault)
		if err != nil {
			return common.NoAddress, common.NoAddress, 0, err
		}
		constAccessor := mgr.GetConstants()
		signingTransactionPeriod := constAccessor.GetInt64Value(constants.SigningTransactionPeriod)
		vault := mgr.Keeper().GetMostSecure(ctx, active, signingTransactionPeriod)
		address, err = vault.PubKey.GetAddress(chain)
		if err != nil {
			return common.NoAddress, common.NoAddress, 0, err
		}

		router = common.NoAddress
		if chain.IsEVM() {
			router = vault.GetContract(chain).Router
		}
	}

	// estimate the inbound confirmation count blocks: ceil(amount/coinbase)
	if chain.DefaultCoinbase() > 0 {
		coinbase := cosmos.NewUint(uint64(chain.DefaultCoinbase()) * common.One)
		confirmations = amount.Quo(coinbase).BigInt().Int64()
		if !amount.Mod(coinbase).IsZero() {
			confirmations++
		}
	}

	return address, router, confirmations, nil
}

func quoteOutboundInfo(ctx cosmos.Context, mgr *Mgrs, coin common.Coin) (int64, error) {
	toi := TxOutItem{
		Memo: "OUT:-",
		Coin: coin,
	}
	outboundHeight, err := mgr.txOutStore.CalcTxOutHeight(ctx, mgr.GetVersion(), toi)
	if err != nil {
		return 0, err
	}
	return outboundHeight - ctx.BlockHeight(), nil
}

// func convertMayachainAmountToWei(amt *big.Int) *big.Int {
// 	return big.NewInt(0).Mul(amt, big.NewInt(common.One*100))
// }

// quoteConvertAsset - converts amount to target asset using MAYAChain pools
func quoteConvertAsset(ctx cosmos.Context, mgr *Mgrs, fromAsset common.Asset, amount sdk.Uint, toAsset common.Asset) (sdk.Uint, error) {
	// no conversion necessary
	if fromAsset.Equals(toAsset) {
		return amount, nil
	}

	// convert to rune
	if !fromAsset.IsBase() {
		// get the fromPool for the from asset
		fromPool, err := mgr.Keeper().GetPool(ctx, fromAsset.GetLayer1Asset())
		if err != nil {
			return sdk.ZeroUint(), fmt.Errorf("failed to get pool: %w", err)
		}

		// ensure pool exists
		if fromPool.IsEmpty() {
			return sdk.ZeroUint(), fmt.Errorf("pool does not exist")
		}

		amount = fromPool.AssetValueInRune(amount)
	}

	// convert to target asset
	if !toAsset.IsBase() {

		toPool, err := mgr.Keeper().GetPool(ctx, toAsset.GetLayer1Asset())
		if err != nil {
			return sdk.ZeroUint(), fmt.Errorf("failed to get pool: %w", err)
		}

		// ensure pool exists
		if toPool.IsEmpty() {
			return sdk.ZeroUint(), fmt.Errorf("pool does not exist")
		}

		amount = toPool.RuneValueInAsset(amount)
	}

	return amount, nil
}

// -------------------------------------------------------------------------------------
// Swap
// -------------------------------------------------------------------------------------

// calculateMinSwapAmount returns the recommended minimum swap amount
// The recommended min swap amount is:
// - MAX(outbound_fee(src_chain), outbound_fee(dest_chain)) * 4 (priced in the inbound asset)
//
// The reason the base value is the MAX of the outbound fees of each chain is because if
// the swap is refunded the input amount will need to cover the outbound fee of the
// source chain. A 4x buffer is applied because outbound fees can spike quickly, meaning
// the original input amount could be less than the new outbound fee. If this happens
// and the swap is refunded, the refund will fail, and the user will lose the entire
// input amount. The min amount could also be determined by the affiliate bps of the
// swap. The affiliate bps of the input amount needs to be enough to cover the native tx fee for the
// affiliate swap to RUNE. In this case, we give a 2x buffer on the native_tx_fee so the
// affiliate receives some amount after the fee is deducted.
func calculateMinSwapAmount(ctx cosmos.Context, mgr *Mgrs, fromAsset, toAsset common.Asset, affiliateBps cosmos.Uint) (cosmos.Uint, error) {
	srcOutboundFee := mgr.GasMgr().GetFee(ctx, fromAsset.GetChain(), fromAsset)
	destOutboundFee := mgr.GasMgr().GetFee(ctx, toAsset.GetChain(), toAsset)

	if fromAsset.GetChain().IsBASEChain() && toAsset.GetChain().IsBASEChain() {
		// If this is a purely THORChain swap, no need to give a 4x buffer since outbound fees do not change
		// 2x buffer should suffice
		return srcOutboundFee.Mul(cosmos.NewUint(2)), nil
	}

	destInSrcAsset, err := quoteConvertAsset(ctx, mgr, toAsset, destOutboundFee, fromAsset)
	if err != nil {
		return cosmos.ZeroUint(), fmt.Errorf("fail to convert dest fee to src asset %w", err)
	}

	minSwapAmount := srcOutboundFee
	if destInSrcAsset.GT(srcOutboundFee) {
		minSwapAmount = destInSrcAsset
	}

	minSwapAmount = minSwapAmount.Mul(cosmos.NewUint(4))

	if affiliateBps.GT(cosmos.ZeroUint()) {
		nativeTxFeeRune := mgr.GasMgr().GetFee(ctx, common.THORChain, common.BaseNative)
		affSwapAmountRune := nativeTxFeeRune.Mul(cosmos.NewUint(2))
		mainSwapAmountRune := affSwapAmountRune.Mul(cosmos.NewUint(10_000)).Quo(affiliateBps)

		mainSwapAmount, err := quoteConvertAsset(ctx, mgr, common.BaseAsset(), mainSwapAmountRune, fromAsset)
		if err != nil {
			return cosmos.ZeroUint(), fmt.Errorf("fail to convert main swap amount to src asset %w", err)
		}

		if mainSwapAmount.GT(minSwapAmount) {
			minSwapAmount = mainSwapAmount
		}
	}

	return minSwapAmount, nil
}

func queryQuoteSwap(ctx cosmos.Context, path []string, req abci.RequestQuery, mgr *Mgrs) ([]byte, error) {
	// extract parameters
	params, err := quoteParseParams(req.Data)
	if err != nil {
		return quoteErrorResponse(err)
	}

	// validate required parameters
	for _, p := range []string{fromAssetParam, toAssetParam, amountParam} {
		if len(params[p]) == 0 {
			return quoteErrorResponse(fmt.Errorf("missing required parameter %s", p))
		}
	}

	// parse assets
	fromAsset, err := common.NewAssetWithShortCodes(mgr.GetVersion(), params[fromAssetParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad from asset: %w", err))
	}
	fromAsset = fuzzyAssetMatch(ctx, mgr.Keeper(), fromAsset)
	toAsset, err := common.NewAssetWithShortCodes(mgr.GetVersion(), params[toAssetParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad to asset: %w", err))
	}
	toAsset = fuzzyAssetMatch(ctx, mgr.Keeper(), toAsset)

	// parse amount
	amount, err := cosmos.ParseUint(params[amountParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad amount: %w", err))
	}

	// parse streaming interval
	streamingInterval := uint64(0) // default value
	if len(params[intervalParam]) > 0 {
		streamingInterval, err = strconv.ParseUint(params[intervalParam][0], 10, 64)
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("bad streaming interval amount: %w", err))
		}
	}
	streamingQuantity := uint64(0) // default value
	if len(params[quantityParam]) > 0 {
		streamingQuantity, err = strconv.ParseUint(params[quantityParam][0], 10, 64)
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("bad streaming quantity amount: %w", err))
		}
	}
	swp := StreamingSwap{
		Interval: streamingInterval,
		Deposit:  amount,
	}
	maxSwapQuantity, err := getMaxSwapQuantity(ctx, mgr, fromAsset, toAsset, swp)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("failed to calculate max streaming swap quantity: %w", err))
	}

	// cap the streaming quantity to the max swap quantity
	if streamingQuantity > maxSwapQuantity {
		streamingQuantity = maxSwapQuantity
	}

	// if from asset is a synth, transfer asset to asgard module
	if fromAsset.IsSyntheticAsset() {
		// mint required coins to asgard so swap can be simulated
		err = mgr.Keeper().MintToModule(ctx, ModuleName, common.NewCoin(fromAsset, amount))
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("failed to mint coins to module: %w", err))
		}

		err = mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(common.NewCoin(fromAsset, amount)))
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("failed to send coins to asgard: %w", err))
		}
	}

	// parse affiliate
	affiliate, affiliateMemo, affiliateBps, swapAmount, _, err := quoteHandleAffiliate(ctx, mgr, params, amount)
	if err != nil {
		return quoteErrorResponse(err)
	}

	// simulate/validate the affiliate swap
	// if affAmt.GT(sdk.ZeroUint()) {
	// 	if fromAsset.IsNativeBase() {
	// 		fee := cosmos.NewUint(uint64(fetchConfigInt64(ctx, mgr, constants.NativeTransactionFee)))
	// 		if affAmt.LTE(fee) {
	// 			return quoteErrorResponse(fmt.Errorf("affiliate amount must be greater than native fee %s", fee))
	// 		}
	// 	} else {
	// 		// validate affiliate address
	// 		affiliateSwapMsg := &types.MsgSwap{
	// 			Tx: common.Tx{
	// 				ID:          common.BlankTxID,
	// 				Chain:       fromAsset.Chain,
	// 				FromAddress: common.NoopAddress,
	// 				ToAddress:   common.NoopAddress,
	// 				Coins: []common.Coin{
	// 					{
	// 						Asset:  fromAsset,
	// 						Amount: affAmt,
	// 					},
	// 				},
	// 				Gas: []common.Coin{{
	// 					Asset:  common.BaseAsset(),
	// 					Amount: sdk.NewUint(1),
	// 				}},
	// 				Memo: "",
	// 			},
	// 			TargetAsset:          common.BaseAsset(),
	// 			TradeTarget:          cosmos.ZeroUint(),
	// 			Destination:          affiliate,
	// 			AffiliateAddress:     common.NoAddress,
	// 			AffiliateBasisPoints: cosmos.ZeroUint(),
	// 		}

	// 		nodeAccounts, err := mgr.Keeper().ListActiveValidators(ctx)
	// 		if err != nil {
	// 			return nil, fmt.Errorf("no active node accounts: %w", err)
	// 		}
	// 		affiliateSwapMsg.Signer = nodeAccounts[0].NodeAddress

	// 		// simulate the swap
	// 		_, err = simulateInternal(ctx, mgr, affiliateSwapMsg)

	// 		if err != nil {
	// 			return quoteErrorResponse(fmt.Errorf("affiliate swap failed: %w", err))
	// 		}
	// 	}
	// }

	// parse destination address or generate a random one
	sendMemo := true
	var destination common.Address
	if len(params[destinationParam]) > 0 {
		destination, err = quoteParseAddress(ctx, mgr, params[destinationParam][0], toAsset.Chain)
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("bad destination address: %w", err))
		}

	} else {
		chain := common.BASEChain
		if !toAsset.IsSyntheticAsset() {
			chain = toAsset.Chain
		}
		destination, err = types.GetRandomPubKey().GetAddress(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to generate address: %w", err)
		}
		sendMemo = false // do not send memo if destination was random
	}

	// parse tolerance basis points
	limit := sdk.ZeroUint()
	if len(params[toleranceBasisPointsParam]) > 0 {
		// validate tolerance basis points
		var toleranceBasisPoints sdk.Uint
		toleranceBasisPoints, err = sdk.ParseUint(params[toleranceBasisPointsParam][0])
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("bad tolerance basis points: %w", err))
		}
		if toleranceBasisPoints.GT(sdk.NewUint(10000)) {
			return quoteErrorResponse(fmt.Errorf("tolerance basis points must be less than 10000"))
		}

		// convert to a limit of target asset amount assuming zero fees and slip
		var feelessEmit sdk.Uint
		feelessEmit, err = quoteConvertAsset(ctx, mgr, fromAsset, swapAmount, toAsset)
		if err != nil {
			return quoteErrorResponse(err)
		}

		limit = feelessEmit.MulUint64(10000 - toleranceBasisPoints.Uint64()).QuoUint64(10000)
	}

	// custom refund addr
	// refundAddress := common.NoAddress
	// if len(params[refundAddressParam]) > 0 {
	// 	refundAddress, err = quoteParseAddress(ctx, mgr, params[refundAddressParam][0], fromAsset.Chain)
	// 	if err != nil {
	// 		return quoteErrorResponse(fmt.Errorf("bad refund address: %w", err))
	// 	}
	// }

	// create the memo
	memo := &SwapMemo{
		MemoBase: mem.MemoBase{
			TxType: TxSwap,
			Asset:  toAsset,
		},
		Destination:          destination,
		SlipLimit:            limit,
		AffiliateAddress:     common.Address(affiliateMemo),
		AffiliateBasisPoints: affiliateBps,
		StreamInterval:       streamingInterval,
		StreamQuantity:       streamingQuantity,
	}

	// if from asset chain has memo length restrictions use a prefix
	memoString := memo.String()
	if !fromAsset.Synth && len(memoString) > fromAsset.Chain.MaxMemoLength() {
		if len(memo.ShortString()) < len(memoString) { // use short codes if available
			memoString = memo.ShortString()
		} else {
			// attempt to shorten
			var fuzzyAsset common.Asset
			fuzzyAsset, err = quoteReverseFuzzyAsset(ctx, mgr, toAsset)
			if err == nil {
				memo.Asset = fuzzyAsset
				memoString = memo.String()
			}
		}

		// this is the shortest we can make it
		if len(memoString) > fromAsset.Chain.MaxMemoLength() {
			return quoteErrorResponse(fmt.Errorf("generated memo too long for source chain"))
		}
	}

	// create the swap message
	msg := &types.MsgSwap{
		Tx: common.Tx{
			ID:          common.BlankTxID,
			Chain:       fromAsset.Chain,
			FromAddress: common.NoopAddress,
			ToAddress:   common.NoopAddress,
			Coins: []common.Coin{
				{
					Asset:  fromAsset,
					Amount: swapAmount,
				},
			},
			Gas: []common.Coin{{
				Asset:  common.BaseAsset(),
				Amount: sdk.NewUint(1),
			}},
			Memo: memoString,
		},
		TargetAsset:          toAsset,
		TradeTarget:          limit,
		Destination:          destination,
		AffiliateAddress:     affiliate,
		AffiliateBasisPoints: affiliateBps,
	}

	// simulate the swap
	res, emitAmount, outboundFeeAmount, err := quoteSimulateSwap(ctx, mgr, amount, msg, 1)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("failed to simulate swap: %w", err))
	}

	// if we're using a streaming swap, calculate emit amount by a sub-swap amount instead
	// of the full amount, then multiply the result by the swap count
	if streamingInterval > 0 && streamingQuantity == 0 {
		streamingQuantity = maxSwapQuantity
	}
	if streamingInterval > 0 && streamingQuantity > 0 {
		msg.TradeTarget = msg.TradeTarget.QuoUint64(streamingQuantity)
		// simulate the swap
		var streamRes *openapi.QuoteSwapResponse
		streamRes, emitAmount, _, err = quoteSimulateSwap(ctx, mgr, amount, msg, streamingQuantity)
		if err != nil {
			return quoteErrorResponse(fmt.Errorf("failed to simulate swap: %w", err))
		}
		res.Fees = streamRes.Fees
	}

	// TODO: After UIs have transitioned everything below the message definition above
	// should reduce to the following:
	//
	// if streamingInterval > 0 && streamingQuantity == 0 {
	//   streamingQuantity = maxSwapQuantity
	// }
	// if streamingInterval > 0 && streamingQuantity > 0 {
	//   msg.TradeTarget = msg.TradeTarget.QuoUint64(streamingQuantity)
	// }
	// res, emitAmount, outboundFeeAmount, err := quoteSimulateSwap(ctx, mgr, amount, msg, streamingQuantity)
	// if err != nil {
	//   return quoteErrorResponse(fmt.Errorf("failed to simulate swap: %w", err))
	// }

	// check invariant
	if emitAmount.LT(outboundFeeAmount) {
		return quoteErrorResponse(fmt.Errorf("invariant broken: emit %s less than outbound fee %s", emitAmount, outboundFeeAmount))
	}

	// the amount out will deduct the outbound fee
	res.ExpectedAmountOut = emitAmount.Sub(outboundFeeAmount).String()

	maxQ := int64(maxSwapQuantity)
	res.MaxStreamingQuantity = &maxQ
	var streamSwapBlocks int64
	if streamingQuantity > 0 {
		streamSwapBlocks = int64(streamingInterval) * int64(streamingQuantity-1)
	}
	res.StreamingSwapBlocks = &streamSwapBlocks
	res.StreamingSwapSeconds = wrapInt64(streamSwapBlocks * common.THORChain.ApproximateBlockMilliseconds() / 1000)

	// estimate the inbound info
	inboundAddress, routerAddress, inboundConfirmations, err := quoteInboundInfo(ctx, mgr, amount, fromAsset.GetChain(), fromAsset)
	if err != nil {
		return quoteErrorResponse(err)
	}
	res.InboundAddress = wrapString(inboundAddress.String())
	if inboundConfirmations > 0 {
		res.InboundConfirmationBlocks = wrapInt64(inboundConfirmations)
		res.InboundConfirmationSeconds = wrapInt64(inboundConfirmations * msg.Tx.Chain.ApproximateBlockMilliseconds() / 1000)
	}

	res.OutboundDelayBlocks = 0
	res.OutboundDelaySeconds = 0
	if !toAsset.Chain.IsBASEChain() {
		// estimate the outbound info
		var outboundDelay int64
		outboundDelay, err = quoteOutboundInfo(ctx, mgr, common.Coin{Asset: toAsset, Amount: emitAmount})
		if err != nil {
			return quoteErrorResponse(err)
		}
		res.OutboundDelayBlocks = outboundDelay
		res.OutboundDelaySeconds = outboundDelay * common.BASEChain.ApproximateBlockMilliseconds() / 1000
	}

	totalSeconds := res.OutboundDelaySeconds
	if res.StreamingSwapSeconds != nil && res.OutboundDelaySeconds < *res.StreamingSwapSeconds {
		totalSeconds = *res.StreamingSwapSeconds
	}
	if inboundConfirmations > 0 {
		totalSeconds += *res.InboundConfirmationSeconds
	}
	res.TotalSwapSeconds = wrapInt64(totalSeconds)

	// send memo if the destination was provided
	if sendMemo {
		res.Memo = wrapString(memoString)
	}

	// set info fields
	if fromAsset.Chain.IsEVM() {
		res.Router = wrapString(routerAddress.String())
	}
	if !fromAsset.Chain.DustThreshold().IsZero() {
		res.DustThreshold = wrapString(fromAsset.Chain.DustThreshold().String())
	}

	res.Notes = fromAsset.GetChain().InboundNotes()
	res.Warning = quoteWarning
	res.Expiry = time.Now().Add(quoteExpiration).Unix()
	minSwapAmount, err := calculateMinSwapAmount(ctx, mgr, fromAsset, toAsset, affiliateBps)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("Failed to calculate min amount in: %s", err.Error()))
	}
	res.RecommendedMinAmountIn = wrapString(minSwapAmount.String())

	// set inbound recommended gas for non-native swaps
	if !fromAsset.Chain.IsBASEChain() {
		inboundGas := mgr.GasMgr().GetGasRate(ctx, fromAsset.Chain)
		res.RecommendedGasRate = wrapString(inboundGas.String())
		res.GasRateUnits = wrapString(fromAsset.Chain.GetGasUnits())
	}

	return json.MarshalIndent(res, "", "  ")
}

// -------------------------------------------------------------------------------------
// Saver Deposit
// -------------------------------------------------------------------------------------

func queryQuoteSaverDeposit(ctx cosmos.Context, path []string, req abci.RequestQuery, mgr *Mgrs) ([]byte, error) {
	// extract parameters
	params, err := quoteParseParams(req.Data)
	if err != nil {
		return quoteErrorResponse(err)
	}

	// validate required parameters
	for _, p := range []string{assetParam, amountParam} {
		if len(params[p]) == 0 {
			return quoteErrorResponse(fmt.Errorf("missing required parameter %s", p))
		}
	}

	// parse asset
	asset, err := common.NewAssetWithShortCodes(mgr.GetVersion(), params[assetParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad asset: %w", err))
	}
	asset = fuzzyAssetMatch(ctx, mgr.Keeper(), asset)

	// parse amount
	amount, err := cosmos.ParseUint(params[amountParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad amount: %w", err))
	}

	// parse affiliate
	affiliate, affiliateMemo, affiliateBps, depositAmount, _, err := quoteHandleAffiliate(ctx, mgr, params, amount)
	if err != nil {
		return quoteErrorResponse(err)
	}

	// generate deposit memo
	depositMemoComponents := []string{
		"+",
		asset.GetSyntheticAsset().String(),
		"",
		affiliateMemo,
		affiliateBps.String(),
	}
	depositMemo := strings.Join(depositMemoComponents[:2], ":")
	if affiliate != common.NoAddress && !affiliateBps.IsZero() {
		depositMemo = strings.Join(depositMemoComponents, ":")
	}

	q := url.Values{}
	q.Add("from_asset", asset.String())
	q.Add("to_asset", asset.GetSyntheticAsset().String())
	q.Add("amount", depositAmount.String())
	q.Add("destination", string(GetRandomBaseAddress())) // required param, not actually used, spoof it

	// ssInterval := mgr.Keeper().GetConfigInt64(ctx, constants.SaversStreamingSwapsInterval)
	// if ssInterval > 0 {
	// 	q.Add("streaming_interval", fmt.Sprintf("%d", ssInterval))
	// 	q.Add("streaming_quantity", fmt.Sprintf("%d", 0))
	// }

	swapReq := abci.RequestQuery{Data: []byte("/mayachain/quote/swap?" + q.Encode())}
	swapResRaw, err := queryQuoteSwap(ctx, nil, swapReq, mgr)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("unable to queryQuoteSwap: %w", err))
	}

	var swapRes *openapi.QuoteSwapResponse
	err = json.Unmarshal(swapResRaw, &swapRes)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("unable to unmarshal swapRes: %w", err))
	}

	expectedAmountOut, _ := sdk.ParseUint(swapRes.ExpectedAmountOut)
	outboundFee, _ := sdk.ParseUint(*swapRes.Fees.Outbound)
	depositAmount = expectedAmountOut.Add(outboundFee)

	// use the swap result info to generate the deposit quote
	res := &openapi.QuoteSaverDepositResponse{
		// TODO: deprecate ExpectedAmountOut in future version
		ExpectedAmountOut:          wrapString(depositAmount.String()),
		ExpectedAmountDeposit:      depositAmount.String(),
		Fees:                       swapRes.Fees,
		InboundConfirmationBlocks:  swapRes.InboundConfirmationBlocks,
		InboundConfirmationSeconds: swapRes.InboundConfirmationSeconds,
		Memo:                       depositMemo,
	}

	// estimate the inbound info
	inboundAddress, _, inboundConfirmations, err := quoteInboundInfo(ctx, mgr, amount, asset.GetLayer1Asset().Chain, asset)
	if err != nil {
		return quoteErrorResponse(err)
	}
	res.InboundAddress = inboundAddress.String()
	res.InboundConfirmationBlocks = wrapInt64(inboundConfirmations)

	// set info fields
	chain := asset.GetLayer1Asset().Chain
	if !chain.DustThreshold().IsZero() {
		res.DustThreshold = wrapString(chain.DustThreshold().String())
		res.RecommendedMinAmountIn = res.DustThreshold
	}
	res.Notes = chain.InboundNotes()
	res.Warning = quoteWarning
	res.Expiry = time.Now().Add(quoteExpiration).Unix()

	// set inbound recommended gas
	inboundGas := mgr.GasMgr().GetGasRate(ctx, chain)
	res.RecommendedGasRate = inboundGas.String()
	res.GasRateUnits = chain.GetGasUnits()

	return json.MarshalIndent(res, "", "  ")
}

// -------------------------------------------------------------------------------------
// Saver Withdraw
// -------------------------------------------------------------------------------------

func queryQuoteSaverWithdraw(ctx cosmos.Context, path []string, req abci.RequestQuery, mgr *Mgrs) ([]byte, error) {
	// extract parameters
	params, err := quoteParseParams(req.Data)
	if err != nil {
		return quoteErrorResponse(err)
	}

	// validate required parameters
	for _, p := range []string{assetParam, addressParam, withdrawBasisPointsParam} {
		if len(params[p]) == 0 {
			return quoteErrorResponse(fmt.Errorf("missing required parameter %s", p))
		}
	}

	// parse asset
	asset, err := common.NewAssetWithShortCodes(mgr.GetVersion(), params[assetParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad asset: %w", err))
	}
	asset = fuzzyAssetMatch(ctx, mgr.Keeper(), asset)
	asset = asset.GetSyntheticAsset() // always use the vault asset

	// parse address
	address, err := common.NewAddress(params[addressParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad address: %w", err))
	}

	// parse basis points
	basisPoints, err := cosmos.ParseUint(params[withdrawBasisPointsParam][0])
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("bad basis points: %w", err))
	}

	// validate basis points
	if basisPoints.GT(sdk.NewUint(10_000)) {
		return quoteErrorResponse(fmt.Errorf("basis points must be less than 10000"))
	}

	// get liquidity provider
	lp, err := mgr.Keeper().GetLiquidityProvider(ctx, asset, address)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("failed to get liquidity provider: %w", err))
	}

	// get the pool
	pool, err := mgr.Keeper().GetPool(ctx, asset)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("failed to get pool: %w", err))
	}

	// get the liquidity provider share of the pool
	lpShare := lp.GetSaversAssetRedeemValue(pool)

	// calculate the withdraw amount
	amount := common.GetSafeShare(basisPoints, sdk.NewUint(10_000), lpShare)

	q := url.Values{}
	q.Add("from_asset", asset.String())
	q.Add("to_asset", asset.GetLayer1Asset().String())
	q.Add("amount", amount.String())
	q.Add("destination", address.String()) // required param, not actually used, spoof it

	swapReq := abci.RequestQuery{Data: []byte("/mayachain/quote/swap?" + q.Encode())}
	swapResRaw, err := queryQuoteSwap(ctx, nil, swapReq, mgr)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("unable to queryQuoteSwap: %w", err))
	}

	var swapRes *openapi.QuoteSwapResponse
	err = json.Unmarshal(swapResRaw, &swapRes)
	if err != nil {
		return quoteErrorResponse(fmt.Errorf("unable to unmarshal swapRes: %w", err))
	}

	// use the swap result info to generate the withdraw quote
	res := &openapi.QuoteSaverWithdrawResponse{
		ExpectedAmountOut: swapRes.ExpectedAmountOut,
		Fees:              swapRes.Fees,
		Memo:              fmt.Sprintf("-:%s:%s", asset.String(), basisPoints.String()),
		DustAmount:        asset.GetLayer1Asset().Chain.DustThreshold().Add(basisPoints).String(),
	}

	// estimate the inbound info
	inboundAddress, _, _, err := quoteInboundInfo(ctx, mgr, amount, asset.GetLayer1Asset().Chain, asset)
	if err != nil {
		return quoteErrorResponse(err)
	}
	res.InboundAddress = inboundAddress.String()

	// estimate the outbound info
	expectedAmountOut, _ := sdk.ParseUint(swapRes.ExpectedAmountOut)
	outboundCoin := common.Coin{Asset: asset.GetLayer1Asset(), Amount: expectedAmountOut}
	outboundDelay, err := quoteOutboundInfo(ctx, mgr, outboundCoin)
	if err != nil {
		return quoteErrorResponse(err)
	}
	res.OutboundDelayBlocks = outboundDelay
	res.OutboundDelaySeconds = outboundDelay * common.BASEChain.ApproximateBlockMilliseconds() / 1000

	// set info fields
	chain := asset.GetLayer1Asset().Chain
	if !chain.DustThreshold().IsZero() {
		res.DustThreshold = wrapString(chain.DustThreshold().String())
	}
	res.Notes = chain.InboundNotes()
	res.Warning = quoteWarning
	res.Expiry = time.Now().Add(quoteExpiration).Unix()

	// set inbound recommended gas
	inboundGas := mgr.GasMgr().GetGasRate(ctx, chain)
	res.RecommendedGasRate = inboundGas.String()
	res.GasRateUnits = chain.GetGasUnits()

	return json.MarshalIndent(res, "", "  ")
}
