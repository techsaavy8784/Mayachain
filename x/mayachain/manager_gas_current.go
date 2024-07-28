package mayachain

import (
	"fmt"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// GasMgrV109 implement GasManager interface which will store the gas related events happened in thorchain to memory
// emit GasEvent per block if there are any
type GasMgrV109 struct {
	gasEvent          *EventGas
	gas               common.Gas
	gasCount          map[common.Asset]int64
	constantsAccessor constants.ConstantValues
	keeper            keeper.Keeper
	mgr               Manager
}

// newGasMgrVCUR create a new instance of GasMgrV109
func newGasMgrVCUR(constantsAccessor constants.ConstantValues, k keeper.Keeper) *GasMgrV109 {
	return &GasMgrV109{
		gasEvent:          NewEventGas(),
		gas:               common.Gas{},
		gasCount:          make(map[common.Asset]int64),
		constantsAccessor: constantsAccessor,
		keeper:            k,
	}
}

func (gm *GasMgrV109) reset() {
	gm.gasEvent = NewEventGas()
	gm.gas = common.Gas{}
	gm.gasCount = make(map[common.Asset]int64)
}

// BeginBlock need to be called when a new block get created , update the internal EventGas to new one
func (gm *GasMgrV109) BeginBlock(mgr Manager) {
	gm.mgr = mgr
	gm.reset()
}

// AddGasAsset to the EventGas
func (gm *GasMgrV109) AddGasAsset(gas common.Gas, increaseTxCount bool) {
	gm.gas = gm.gas.Add(gas)
	if !increaseTxCount {
		return
	}
	for _, coin := range gas {
		gm.gasCount[coin.Asset]++
	}
}

// GetGas return gas
func (gm *GasMgrV109) GetGas() common.Gas {
	return gm.gas
}

// GetFee retrieve the network fee information from kv store, and calculate the dynamic fee customer should pay
// the return value is the amount of fee in RUNE
func (gm *GasMgrV109) GetFee(ctx cosmos.Context, chain common.Chain, asset common.Asset) cosmos.Uint {
	outboundTxFee, err := gm.keeper.GetMimir(ctx, constants.OutboundTransactionFee.String())
	if outboundTxFee < 0 || err != nil {
		outboundTxFee = gm.constantsAccessor.GetInt64Value(constants.OutboundTransactionFee)
	}
	transactionFee := cosmos.NewUint(uint64(outboundTxFee))
	// if the asset is Native CACAO, then we could just return the transaction Fee
	// because transaction fee is always in native CACAO. This is called from both
	// ends when a tx is sent out to BASEChain so we need to check the MAYA.CACAO case
	if asset.IsBase() && chain.Equals(common.BASEChain) {
		return transactionFee
	}

	// if the asset is synthetic asset , it need to get the layer 1 asset pool and convert it
	// synthetic asset live on BASEChain , thus it doesn't need to get the layer1 network fee
	if asset.IsSyntheticAsset() {
		return gm.getRuneInAssetValue(ctx, transactionFee, asset)
	}

	pool, err := gm.keeper.GetPool(ctx, chain.GetGasAsset())
	if err != nil {
		ctx.Logger().Error("fail to get pool", "asset", asset, "error", err)
		return transactionFee
	}

	var fee cosmos.Uint

	networkFee, err := gm.keeper.GetNetworkFee(ctx, chain)
	if err != nil {
		ctx.Logger().Error("fail to get network fee", "error", err)
		return transactionFee
	}

	if err = networkFee.Valid(); err != nil {
		ctx.Logger().Error("network fee is invalid", "error", err, "chain", chain)
		return transactionFee
	}

	minOutboundUSD, err := gm.keeper.GetMimir(ctx, constants.MinimumL1OutboundFeeUSD.String())
	if minOutboundUSD < 0 || err != nil {
		minOutboundUSD = gm.constantsAccessor.GetInt64Value(constants.MinimumL1OutboundFeeUSD)
	}
	oneDollarRune := cosmos.ZeroUint()
	// since gm.mgr get set at BeginBlock , so here add a safeguard in case gm.mgr is nil
	if gm.mgr != nil {
		oneDollarRune = DollarInRune(ctx, gm.mgr)
	}
	minAsset := cosmos.ZeroUint()
	if !oneDollarRune.IsZero() {
		// since MinOutboundUSD is in USD value , thus need to figure out how much RUNE
		// here use GetShare instead GetSafeShare it is because minOutboundUSD can set to more than $1
		minOutboundInRune := common.GetUncappedShare(cosmos.NewUint(uint64(minOutboundUSD)),
			cosmos.NewUint(common.One),
			oneDollarRune)

		minAsset = pool.RuneValueInAsset(minOutboundInRune)
	}
	network, err := gm.keeper.GetNetwork(ctx)
	if err != nil {
		ctx.Logger().Error("fail to get network data", "error", err)
	}

	targetOutboundFeeSurplus := gm.mgr.GetConfigInt64(ctx, constants.TargetOutboundFeeSurplusRune)
	maxMultiplierBasisPoints := gm.mgr.GetConfigInt64(ctx, constants.MaxOutboundFeeMultiplierBasisPoints)
	minMultiplierBasisPoints := gm.mgr.GetConfigInt64(ctx, constants.MinOutboundFeeMultiplierBasisPoints)

	// Calculate outbound fee based on current fee multiplier
	chainBaseFee := networkFee.TransactionSize * networkFee.TransactionFeeRate
	if chain.Equals(common.ARBChain) {
		chainBaseFee /= 1000 // convert from 1e11 to 1e8
	}
	feeMultiplierBps := gm.CalcOutboundFeeMultiplier(ctx, cosmos.NewUint(uint64(targetOutboundFeeSurplus)), cosmos.NewUint(network.OutboundGasSpentCacao), cosmos.NewUint(network.OutboundGasWithheldCacao), cosmos.NewUint(uint64(maxMultiplierBasisPoints)), cosmos.NewUint(uint64(minMultiplierBasisPoints)))
	finalFee := common.GetUncappedShare(cosmos.NewUint(chainBaseFee), cosmos.NewUint(10_000), feeMultiplierBps)

	fee = cosmos.RoundToDecimal(
		finalFee,
		pool.Decimals,
	)

	// Ensure fee is always more than minAsset
	if fee.LT(minAsset) {
		fee = minAsset
	}

	if asset.Equals(asset.GetChain().GetGasAsset()) && chain.Equals(asset.GetChain()) {
		return fee
	}

	// convert gas asset value into cacao
	if pool.BalanceAsset.Equal(cosmos.ZeroUint()) || pool.BalanceCacao.Equal(cosmos.ZeroUint()) {
		// hardcode value to previous transactionFee value
		return cosmos.NewUint(2_000000)
	}

	fee = pool.AssetValueInRune(fee)
	if asset.IsBase() {
		return fee
	}

	// convert rune value into non-gas asset value
	pool, err = gm.keeper.GetPool(ctx, asset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "asset", asset, "error", err)
		return transactionFee
	}
	if pool.BalanceAsset.Equal(cosmos.ZeroUint()) || pool.BalanceCacao.Equal(cosmos.ZeroUint()) {
		// hardcode value to previous transactionFee value
		return cosmos.NewUint(2_000000)
	}
	return pool.RuneValueInAsset(fee)
}

// CalcOutboundFeeMultiplier returns the current outbound fee multiplier based on current and target outbound fee surplus
func (gm *GasMgrV109) CalcOutboundFeeMultiplier(ctx cosmos.Context, targetSurplusRune, gasSpentRune, gasWithheldRune, maxMultiplier, minMultiplier cosmos.Uint) cosmos.Uint {
	// Sanity check
	if targetSurplusRune.Equal(cosmos.ZeroUint()) {
		ctx.Logger().Error("target gas surplus is zero")
		return maxMultiplier
	}
	if minMultiplier.GT(maxMultiplier) {
		ctx.Logger().Error("min multiplier greater than max multiplier", "minMultiplier", minMultiplier, "maxMultiplier", maxMultiplier)
		return cosmos.NewUint(30_000) // should never happen, return old default
	}

	// Find current surplus (gas withheld from user - gas spent by the reserve)
	surplusRune := common.SafeSub(gasWithheldRune, gasSpentRune)

	// How many BPs to reduce the multiplier
	multiplierReducedBps := common.GetSafeShare(surplusRune, targetSurplusRune, common.SafeSub(maxMultiplier, minMultiplier))
	return common.SafeSub(maxMultiplier, multiplierReducedBps)
}

// getRuneInAssetValue convert the transaction fee to asset value , when the given asset is synthetic , it will need to get
// the layer1 asset first , and then use the pool to convert
func (gm *GasMgrV109) getRuneInAssetValue(ctx cosmos.Context, transactionFee cosmos.Uint, asset common.Asset) cosmos.Uint {
	if asset.IsSyntheticAsset() {
		asset = asset.GetLayer1Asset()
	}
	pool, err := gm.keeper.GetPool(ctx, asset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "asset", asset, "error", err)
		return transactionFee
	}
	if pool.BalanceAsset.Equal(cosmos.ZeroUint()) || pool.BalanceCacao.Equal(cosmos.ZeroUint()) {
		return transactionFee
	}

	return pool.RuneValueInAsset(transactionFee)
}

// GetGasRate return the gas rate
func (gm *GasMgrV109) GetGasRate(ctx cosmos.Context, chain common.Chain) cosmos.Uint {
	outboundTxFee, err := gm.keeper.GetMimir(ctx, constants.OutboundTransactionFee.String())
	if outboundTxFee < 0 || err != nil {
		outboundTxFee = gm.constantsAccessor.GetInt64Value(constants.OutboundTransactionFee)
	}
	transactionFee := cosmos.NewUint(uint64(outboundTxFee))
	if chain.Equals(common.BASEChain) {
		return transactionFee
	}

	networkFee, err := gm.keeper.GetNetworkFee(ctx, chain)
	if err != nil {
		ctx.Logger().Error("fail to get network fee", "error", err)
		return transactionFee
	}
	if err := networkFee.Valid(); err != nil {
		ctx.Logger().Error("network fee is invalid", "error", err, "chain", chain)
		return transactionFee
	}
	if chain.Equals(common.THORChain) {
		return cosmos.NewUint(networkFee.TransactionFeeRate)
	}

	return cosmos.RoundToDecimal(
		cosmos.NewUint(networkFee.TransactionFeeRate*3/2),
		chain.GetGasAssetDecimal(),
	)
}

func (gm *GasMgrV109) GetNetworkFee(ctx cosmos.Context, chain common.Chain) (NetworkFee, error) {
	outboundTxFee := fetchConfigInt64(ctx, gm.mgr, constants.OutboundTransactionFee)

	if chain.Equals(common.BASEChain) {
		return types.NewNetworkFee(chain, 1, uint64(outboundTxFee)), nil
	}

	return gm.keeper.GetNetworkFee(ctx, chain)
}

// GetMaxGas will calculate the maximum gas fee a tx can use
func (gm *GasMgrV109) GetMaxGas(ctx cosmos.Context, chain common.Chain) (common.Coin, error) {
	gasAsset := chain.GetGasAsset()
	var amount cosmos.Uint

	nf, err := gm.keeper.GetNetworkFee(ctx, chain)
	if err != nil {
		return common.NoCoin, fmt.Errorf("fail to get network fee for chain(%s): %w", chain, err)
	}
	if chain.IsBNB() || chain.Equals(common.THORChain) {
		amount = cosmos.NewUint(nf.TransactionSize * nf.TransactionFeeRate)
	} else {
		amount = cosmos.NewUint(nf.TransactionSize * nf.TransactionFeeRate).MulUint64(3).QuoUint64(2)
	}

	if chain.Equals(common.ARBChain) {
		amount = amount.QuoUint64(1000) // convert from 1e11 to 1e8
	}

	gasCoin := common.NewCoin(gasAsset, amount)
	chainGasAssetPrecision := chain.GetGasAssetDecimal()
	gasCoin.Amount = cosmos.RoundToDecimal(amount, chainGasAssetPrecision)
	gasCoin.Decimals = chainGasAssetPrecision
	return gasCoin, nil
}

// SubGas will subtract the gas from the gas manager
func (gm *GasMgrV109) SubGas(gas common.Gas) {
	gm.gas = gm.gas.Sub(gas)
}

// EndBlock emit the events
func (gm *GasMgrV109) EndBlock(ctx cosmos.Context, keeper keeper.Keeper, eventManager EventManager) {
	gm.ProcessGas(ctx, keeper)

	blocksPerDay := gm.constantsAccessor.GetInt64Value(constants.BlocksPerDay)
	if IsPeriodLastBlock(ctx, uint64(blocksPerDay)) {
		keeper.DistributeMayaFund(ctx, gm.constantsAccessor)
	}

	if len(gm.gasEvent.Pools) == 0 {
		return
	}
	if err := eventManager.EmitGasEvent(ctx, gm.gasEvent); nil != err {
		ctx.Logger().Error("fail to emit gas event", "error", err)
	}
	gm.reset() // do not remove, will cause consensus failures
}

// ProcessGas to subsidise the pool with CACAO for the gas they have spent
func (gm *GasMgrV109) ProcessGas(ctx cosmos.Context, keeper keeper.Keeper) {
	if keeper.RagnarokInProgress(ctx) {
		// ragnarok is in progress , stop
		return
	}

	network, err := keeper.GetNetwork(ctx)
	if err != nil {
		ctx.Logger().Error("fail to get network data", "error", err)
		return
	}
	for _, gas := range gm.gas {
		// if the coin is zero amount, don't need to do anything
		if gas.Amount.IsZero() {
			continue
		}

		pool, err := keeper.GetPool(ctx, gas.Asset)
		if err != nil {
			ctx.Logger().Error("fail to get pool", "pool", gas.Asset, "error", err)
			continue
		}
		if err := pool.Valid(); err != nil {
			ctx.Logger().Error("invalid pool", "pool", gas.Asset, "error", err)
			continue
		}
		cacaoGas := pool.AssetValueInRune(gas.Amount) // Convert to CACAO (gas will never be CACAO)
		if cacaoGas.IsZero() {
			continue
		}
		// If Cacao owed now exceeds the Total Reserve, return it all
		if cacaoGas.LT(keeper.GetRuneBalanceOfModule(ctx, ReserveName)) {
			coin := common.NewCoin(common.BaseNative, cacaoGas)
			if err := keeper.SendFromModuleToModule(ctx, ReserveName, AsgardName, common.NewCoins(coin)); err != nil {
				ctx.Logger().Error("fail to transfer funds from reserve to asgard", "pool", gas.Asset, "error", err)
				continue
			}
			pool.BalanceCacao = pool.BalanceCacao.Add(cacaoGas) // Add to the pool
			network.OutboundGasSpentCacao += cacaoGas.Uint64()  // Add $CACAO spent on gas by the reserve
		} else {
			// since we don't have enough in the reserve to cover the gas used,
			// no cacao is added to the pool, sorry LPs!
			cacaoGas = cosmos.ZeroUint()
		}
		pool.BalanceAsset = common.SafeSub(pool.BalanceAsset, gas.Amount)

		if err := keeper.SetPool(ctx, pool); err != nil {
			ctx.Logger().Error("fail to set pool", "pool", gas.Asset, "error", err)
			continue
		}

		gasPool := GasPool{
			Asset:    gas.Asset,
			AssetAmt: gas.Amount,
			CacaoAmt: cacaoGas,
			Count:    gm.gasCount[gas.Asset],
		}
		gm.gasEvent.UpsertGasPool(gasPool)
	}

	if err := keeper.SetNetwork(ctx, network); err != nil {
		ctx.Logger().Error("fail to set network data", "error", err)
	}
}
