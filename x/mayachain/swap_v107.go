package mayachain

import (
	"errors"
	"fmt"

	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
)

type SwapperV107 struct{}

func newSwapperV107() *SwapperV107 {
	return &SwapperV107{}
}

// validateMessage is trying to validate the legitimacy of the incoming message and decide whether THORNode can handle it
func (s *SwapperV107) validateMessage(tx common.Tx, target common.Asset, destination common.Address) error {
	if err := tx.Valid(); err != nil {
		return err
	}
	if target.IsEmpty() {
		return errors.New("target is empty")
	}
	if destination.IsEmpty() {
		return errors.New("destination is empty")
	}

	return nil
}

func (s *SwapperV107) Swap(ctx cosmos.Context,
	keeper keeper.Keeper,
	tx common.Tx,
	target common.Asset,
	destination common.Address,
	swapTarget cosmos.Uint,
	dexAgg string,
	dexAggTargetAsset string,
	dexAggLimit *cosmos.Uint,
	_ StreamingSwap,
	transactionFee cosmos.Uint, synthVirtualDepthMult int64, mgr Manager,
) (cosmos.Uint, []*EventSwap, error) {
	var swapEvents []*EventSwap

	if err := s.validateMessage(tx, target, destination); err != nil {
		return cosmos.ZeroUint(), swapEvents, err
	}
	source := tx.Coins[0].Asset

	if source.IsSyntheticAsset() {
		burnHeight, _ := keeper.GetMimir(ctx, "BurnSynths")
		if burnHeight > 0 && ctx.BlockHeight() > burnHeight {
			return cosmos.ZeroUint(), swapEvents, fmt.Errorf("burning synthetics has been disabled")
		}
	}
	if target.IsSyntheticAsset() {
		mintHeight, _ := keeper.GetMimir(ctx, "MintSynths")
		if mintHeight > 0 && ctx.BlockHeight() > mintHeight {
			return cosmos.ZeroUint(), swapEvents, fmt.Errorf("minting synthetics has been disabled")
		}
	}

	if !destination.IsNoop() && !destination.IsChain(target.GetChain(), keeper.GetVersion()) {
		return cosmos.ZeroUint(), swapEvents, fmt.Errorf("destination address is not a valid %s address", target.GetChain())
	}
	if source.Equals(target) {
		return cosmos.ZeroUint(), swapEvents, fmt.Errorf("cannot swap from %s --> %s, assets match", source, target)
	}

	isDoubleSwap := !source.IsBase() && !target.IsBase()
	if isDoubleSwap {
		var swapErr error
		var swapEvt *EventSwap
		var amt cosmos.Uint
		// Here we use a swapTarget of 0 because the target is for the next swap asset in a double swap
		amt, swapEvt, swapErr = s.swapOne(ctx, keeper, tx, common.BaseAsset(), destination, cosmos.ZeroUint(), transactionFee, synthVirtualDepthMult, mgr)
		if swapErr != nil {
			return cosmos.ZeroUint(), swapEvents, swapErr
		}
		tx.Coins = common.Coins{common.NewCoin(common.BaseAsset(), amt)}
		tx.Gas = nil
		swapEvt.OutTxs = common.NewTx(common.BlankTxID, tx.FromAddress, tx.ToAddress, tx.Coins, tx.Gas, tx.Memo)
		swapEvents = append(swapEvents, swapEvt)
	}
	assetAmount, swapEvt, swapErr := s.swapOne(ctx, keeper, tx, target, destination, swapTarget, transactionFee, synthVirtualDepthMult, mgr)
	if swapErr != nil {
		return cosmos.ZeroUint(), swapEvents, swapErr
	}
	swapEvents = append(swapEvents, swapEvt)
	if !swapTarget.IsZero() && assetAmount.LT(swapTarget) {
		// **NOTE** this error string is utilized by the order book manager to
		// catch the error. DO NOT change this error string without updating
		// the order book manager as well
		return cosmos.ZeroUint(), swapEvents, fmt.Errorf("emit asset %s less than price limit %s", assetAmount, swapTarget)
	}
	if target.IsBase() {
		if assetAmount.LTE(transactionFee) {
			return cosmos.ZeroUint(), swapEvents, fmt.Errorf("output CACAO (%s) is not enough to pay transaction fee", assetAmount)
		}
	}
	// emit asset is zero
	if assetAmount.IsZero() {
		return cosmos.ZeroUint(), swapEvents, errors.New("zero emit asset")
	}

	// Thanks to CacheContext, the swap event can be emitted before handling outbounds,
	// since if there's a later error the event emission will not take place.
	for _, evt := range swapEvents {
		if err := mgr.EventMgr().EmitEvent(ctx, evt); err != nil {
			ctx.Logger().Error("fail to emit swap event", "error", err)
		}
		if !evt.OutTxs.IsEmpty() {
			outboundEvt := NewEventOutbound(evt.InTx.ID, evt.OutTxs)
			if err := mgr.EventMgr().EmitEvent(ctx, outboundEvt); err != nil {
				ctx.Logger().Error("fail to emit an outbound event for double swap", "error", err)
			}
		}
		if err := keeper.AddToLiquidityFees(ctx, evt.Pool, evt.LiquidityFeeInCacao); err != nil {
			return assetAmount, swapEvents, fmt.Errorf("fail to add to liquidity fees: %w", err)
		}
		telemetry.IncrCounterWithLabels(
			[]string{"mayanode", "swap", "count"},
			float32(1),
			[]metrics.Label{telemetry.NewLabel("pool", evt.Pool.String())},
		)
		telemetry.IncrCounterWithLabels(
			[]string{"mayanode", "swap", "slip"},
			telem(evt.SwapSlip),
			[]metrics.Label{telemetry.NewLabel("pool", evt.Pool.String())},
		)
		telemetry.IncrCounterWithLabels(
			[]string{"mayanode", "swap", "liquidity_fee"},
			telem(evt.LiquidityFeeInCacao),
			[]metrics.Label{telemetry.NewLabel("pool", evt.Pool.String())},
		)
	}

	if !destination.IsNoop() {
		toi := TxOutItem{
			Chain:                 target.GetChain(),
			InHash:                tx.ID,
			ToAddress:             destination,
			Coin:                  common.NewCoin(target, assetAmount),
			Aggregator:            dexAgg,
			AggregatorTargetAsset: dexAggTargetAsset,
			AggregatorTargetLimit: dexAggLimit,
		}
		// let the txout manager mint our outbound asset if it is a synthetic asset
		if toi.Chain.IsBASEChain() && toi.Coin.Asset.IsSyntheticAsset() {
			toi.ModuleName = ModuleName
		}

		ok, err := mgr.TxOutStore().TryAddTxOutItem(ctx, mgr, toi, swapTarget)
		if err != nil {
			return assetAmount, swapEvents, ErrInternal(err, "fail to add outbound tx")
		}
		if !ok {
			return assetAmount, swapEvents, errFailAddOutboundTx
		}
	}

	return assetAmount, swapEvents, nil
}

func (s *SwapperV107) burnCoins(ctx cosmos.Context, keeper keeper.Keeper, coins common.Coins) error {
	err := keeper.SendFromModuleToModule(ctx, AsgardName, ModuleName, coins)
	if err != nil {
		ctx.Logger().Error("fail to move coins during swap", "error", err)
		return err
	}
	for _, coin := range coins {
		err := keeper.BurnFromModule(ctx, ModuleName, coin)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SwapperV107) swapOne(ctx cosmos.Context,
	keeper keeper.Keeper, tx common.Tx,
	target common.Asset,
	destination common.Address,
	swapTarget cosmos.Uint,
	transactionFee cosmos.Uint,
	synthVirtualDepthMult int64,
	mgr Manager,
) (amt cosmos.Uint, evt *EventSwap, swapErr error) {
	source := tx.Coins[0].Asset
	amount := tx.Coins[0].Amount

	ctx.Logger().Info("swapping", "from", tx.FromAddress, "coins", tx.Coins[0], "target", target, "to", destination, "fee", transactionFee)

	var X, x, Y, liquidityFee, emitAssets cosmos.Uint
	var swapSlip cosmos.Uint
	var pool Pool
	var err error

	// Set asset to our non-rune asset
	asset := source
	if source.IsBase() {
		asset = target
		if amount.LTE(transactionFee) {
			// stop swap , because the output will not enough to pay for transaction fee
			return cosmos.ZeroUint(), evt, errSwapFailNotEnoughFee
		}
	}
	if asset.IsSyntheticAsset() {
		asset = asset.GetLayer1Asset()
	}

	swapEvt := NewEventSwap(
		asset,
		swapTarget,
		cosmos.ZeroUint(),
		cosmos.ZeroUint(),
		cosmos.ZeroUint(),
		tx,
		common.NoCoin,
		cosmos.ZeroUint(),
	)

	// Check if pool exists
	if !keeper.PoolExist(ctx, asset.GetLayer1Asset()) {
		err = fmt.Errorf("pool %s doesn't exist", asset)
		return cosmos.ZeroUint(), evt, err
	}

	pool, err = keeper.GetPool(ctx, asset.GetLayer1Asset())
	if err != nil {
		return cosmos.ZeroUint(), evt, ErrInternal(err, fmt.Sprintf("fail to get pool(%s)", asset))
	}
	// sanity check: ensure we're never swapping with the vault
	// (technically is actually the yield bearing synth vault)
	if pool.Asset.IsVaultAsset() {
		return cosmos.ZeroUint(), evt, ErrInternal(err, fmt.Sprintf("dev error: swapping with a vault(%s) is not allowed", asset))
	}
	synthSupply := keeper.GetTotalSupply(ctx, pool.Asset.GetSyntheticAsset())
	pool.CalcUnits(keeper.GetVersion(), synthSupply)

	// pool must be available unless source is synthetic
	// synths may be redeemed regardless of pool status
	if !source.IsSyntheticAsset() && !pool.IsAvailable() {
		return cosmos.ZeroUint(), evt, fmt.Errorf("pool(%s) is not available", asset)
	}

	// Get our X, x, Y values
	if source.IsBase() {
		X = pool.BalanceCacao
		Y = pool.BalanceAsset
	} else {
		Y = pool.BalanceCacao
		X = pool.BalanceAsset
	}
	x = amount

	// give virtual pool depth if we're swapping with a synthetic asset
	if source.IsSyntheticAsset() || target.IsSyntheticAsset() {
		X = common.GetUncappedShare(cosmos.NewUint(uint64(synthVirtualDepthMult)), cosmos.NewUint(10_000), X)
		Y = common.GetUncappedShare(cosmos.NewUint(uint64(synthVirtualDepthMult)), cosmos.NewUint(10_000), Y)
	}

	// check our X,x,Y values are valid
	if x.IsZero() {
		return cosmos.ZeroUint(), evt, errSwapFailInvalidAmount
	}
	if X.IsZero() || Y.IsZero() {
		return cosmos.ZeroUint(), evt, errSwapFailInvalidBalance
	}

	liquidityFee = s.CalcLiquidityFee(X, x, Y)

	slipFeeAddedBasisPoints := getSlipFeeAddedBasisPoints(ctx, mgr)

	swapSlip = s.CalcSwapSlip(X, x, cosmos.NewUint(slipFeeAddedBasisPoints))

	emitAssets = s.CalcAssetEmission(X, x, Y)
	emitAssets = cosmos.RoundToDecimal(emitAssets, pool.Decimals)
	swapEvt.LiquidityFee = liquidityFee

	if source.IsBase() {
		swapEvt.LiquidityFeeInCacao = pool.AssetValueInRune(liquidityFee)
	} else {
		// because the output asset is RUNE , so liqualidtyFee is already in RUNE
		swapEvt.LiquidityFeeInCacao = liquidityFee
	}
	swapEvt.SwapSlip = swapSlip
	swapEvt.EmitAsset = common.NewCoin(target, emitAssets)

	// do THORNode have enough balance to swap?
	if emitAssets.GTE(Y) {
		return cosmos.ZeroUint(), evt, errSwapFailNotEnoughBalance
	}

	ctx.Logger().Info("pre swap", "pool", pool.Asset, "rune", pool.BalanceCacao, "asset", pool.BalanceAsset, "lp units", pool.LPUnits, "synth units", pool.SynthUnits)

	if source.IsSyntheticAsset() || target.IsSyntheticAsset() {
		// we're doing a synth swap
		if source.IsSyntheticAsset() {
			// our source is a pegged asset, burn it all
			pool.BalanceCacao = common.SafeSub(pool.BalanceCacao, emitAssets)
			if err := s.burnCoins(ctx, keeper, tx.Coins); err != nil {
				return cosmos.ZeroUint(), evt, err
			}
		} else {
			pool.BalanceCacao = pool.BalanceCacao.Add(x)
		}
	} else {
		if source.IsBase() {
			pool.BalanceCacao = X.Add(x)
			pool.BalanceAsset = common.SafeSub(Y, emitAssets)
		} else {
			pool.BalanceAsset = X.Add(x)
			pool.BalanceCacao = common.SafeSub(Y, emitAssets)
		}
	}
	ctx.Logger().Info("post swap", "pool", pool.Asset, "rune", pool.BalanceCacao, "asset", pool.BalanceAsset, "lp units", pool.LPUnits, "synth units", pool.SynthUnits, "emit asset", emitAssets)

	if err := keeper.SetPool(ctx, pool); err != nil {
		return cosmos.ZeroUint(), evt, fmt.Errorf("fail to set pool")
	}

	return emitAssets, swapEvt, nil
}

// calculate the number of assets sent to the address (includes liquidity fee)
// nolint
func (s *SwapperV107) CalcAssetEmission(X, x, Y cosmos.Uint) cosmos.Uint {
	// ( x * X * Y ) / ( x + X )^2
	numerator := x.Mul(X).Mul(Y)
	denominator := x.Add(X).Mul(x.Add(X))
	if denominator.IsZero() {
		return cosmos.ZeroUint()
	}
	return numerator.Quo(denominator)
}

// CalculateLiquidityFee the fee of the swap
// nolint
func (s *SwapperV107) CalcLiquidityFee(X, x, Y cosmos.Uint) cosmos.Uint {
	// ( x^2 *  Y ) / ( x + X )^2
	numerator := x.Mul(x).Mul(Y)
	denominator := x.Add(X).Mul(x.Add(X))
	if denominator.IsZero() {
		return cosmos.ZeroUint()
	}
	return numerator.Quo(denominator)
}

// CalcSwapSlip - calculate the swap slip, expressed in basis points (10000)
// nolint
func (s *SwapperV107) CalcSwapSlip(Xi, xi cosmos.Uint, slipFeeAddedBasisPoints cosmos.Uint) cosmos.Uint {
	// Cast to DECs
	xD := cosmos.NewDecFromBigInt(xi.BigInt())
	XD := cosmos.NewDecFromBigInt(Xi.BigInt())
	dec10k := cosmos.NewDec(10000)
	// x / (x + X)
	denD := xD.Add(XD)
	if denD.IsZero() {
		return cosmos.ZeroUint()
	}
	swapSlipD := xD.Quo(denD)                                     // Division with DECs
	swapSlip := swapSlipD.Mul(dec10k)                             // Adds 5 0's
	swapSlipUint := cosmos.NewUint(uint64(swapSlip.RoundInt64())) // Casts back to Uint as Basis Points

	swapSlipUint = swapSlipUint.Add(slipFeeAddedBasisPoints) // Add slip fee basis point. default 10 = 0.1%

	return swapSlipUint
}
