package mayachain

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// withdrawV108 all the asset
// it returns runeAmt,assetAmount,protectionRuneAmt,units, lastWithdraw,err
func withdrawV108(ctx cosmos.Context, msg MsgWithdrawLiquidity, mgr Manager) (cosmos.Uint, cosmos.Uint, cosmos.Uint, cosmos.Uint, cosmos.Uint, error) {
	if err := validateWithdrawV105(ctx, mgr.Keeper(), msg); err != nil {
		ctx.Logger().Error("msg withdraw fail validation", "error", err)
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), err
	}

	pool, err := mgr.Keeper().GetPool(ctx, msg.Asset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "error", err)
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), err
	}
	synthSupply := mgr.Keeper().GetTotalSupply(ctx, pool.Asset.GetSyntheticAsset())
	pool.CalcUnits(mgr.GetVersion(), synthSupply)

	lp, err := mgr.Keeper().GetLiquidityProvider(ctx, msg.Asset, msg.WithdrawAddress)
	if err != nil {
		ctx.Logger().Error("can't find liquidity provider", "error", err)
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), err

	}

	poolCacao := pool.BalanceCacao
	poolAsset := pool.BalanceAsset
	originalLiquidityProviderUnits := lp.Units
	fLiquidityProviderUnit := lp.Units
	if lp.Units.IsZero() {
		if !lp.PendingCacao.IsZero() || !lp.PendingAsset.IsZero() {
			if isLiquidityAuction(ctx, mgr.Keeper()) {
				tier1 := mgr.GetConstants().GetInt64Value(constants.WithdrawTier1)
				var tier int64
				tier, err = mgr.Keeper().GetLiquidityAuctionTier(ctx, lp.CacaoAddress)
				if err != nil {
					ctx.Logger().Error("fail to get liquidity auction tier", "error", err)
				}

				if tier == tier1 {
					return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), errors.New("tier1 cannot withdraw during liquidity auction")
				}
			}

			mgr.Keeper().RemoveLiquidityProvider(ctx, lp)
			pool.PendingInboundCacao = common.SafeSub(pool.PendingInboundCacao, lp.PendingCacao)
			pool.PendingInboundAsset = common.SafeSub(pool.PendingInboundAsset, lp.PendingAsset)
			if err = mgr.Keeper().SetPool(ctx, pool); err != nil {
				ctx.Logger().Error("fail to save pool pending inbound funds", "error", err)
			}

			return lp.PendingCacao, cosmos.RoundToDecimal(lp.PendingAsset, pool.Decimals), cosmos.ZeroUint(), lp.Units, cosmos.ZeroUint(), nil
		}
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), errNoLiquidityUnitLeft
	}

	cv := mgr.GetConstants()
	height := ctx.BlockHeight()
	if height < (lp.LastAddHeight + cv.GetInt64Value(constants.LiquidityLockUpBlocks)) {
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), errWithdrawWithin24Hours
	}

	ctx.Logger().Info("pool before withdraw", "pool units", pool.GetPoolUnits(), "balance RUNE", poolCacao, "balance asset", poolAsset)
	ctx.Logger().Info("liquidity provider before withdraw", "liquidity provider unit", fLiquidityProviderUnit)

	pauseAsym, _ := mgr.Keeper().GetMimir(ctx, fmt.Sprintf("PauseAsymWithdrawal-%s", pool.Asset.GetChain()))
	assetToWithdraw := assetToWithdrawV89(msg, lp, pauseAsym)

	if pool.Status == PoolAvailable && lp.CacaoDepositValue.IsZero() && lp.AssetDepositValue.IsZero() {
		lp.CacaoDepositValue = lp.CacaoDepositValue.Add(common.GetSafeShare(lp.Units, pool.GetPoolUnits(), pool.BalanceCacao))
		lp.AssetDepositValue = lp.AssetDepositValue.Add(common.GetSafeShare(lp.Units, pool.GetPoolUnits(), pool.BalanceAsset))
	}

	// calculate any impermanent loss protection or not
	protectionCacaoAmount := cosmos.ZeroUint()
	extraUnits := cosmos.ZeroUint()
	fullProtectionLine, err := mgr.Keeper().GetMimir(ctx, constants.FullImpLossProtectionBlocks.String())
	if fullProtectionLine < 0 || err != nil {
		fullProtectionLine = cv.GetInt64Value(constants.FullImpLossProtectionBlocks)
	}
	ilpPoolMimirKey := fmt.Sprintf("ILP-DISABLED-%s", pool.Asset)
	ilpDisabled, err := mgr.Keeper().GetMimir(ctx, ilpPoolMimirKey)
	if err != nil {
		ctx.Logger().Error("fail to get ILP-DISABLED mimir", "error", err, "key", ilpPoolMimirKey)
		ilpDisabled = 0
	}
	// only when Pool is in Available status will apply impermanent loss protection
	if fullProtectionLine > 0 && pool.Status == PoolAvailable && !(ilpDisabled > 0 && !pool.Asset.IsVaultAsset()) { // if protection line is zero, no imp loss protection is given
		lastAddHeight := lp.LastAddHeight
		if lastAddHeight < pool.StatusSince {
			lastAddHeight = pool.StatusSince
		}
		implProtectionCacaoAmount, depositValue, redeemValue := calcImpLossV102(ctx, mgr, lastAddHeight, lp, msg.BasisPoints, fullProtectionLine, pool)
		ctx.Logger().Info("imp loss calculation", "deposit value", depositValue, "redeem value", redeemValue, "protection", implProtectionCacaoAmount)
		if !implProtectionCacaoAmount.IsZero() {
			protectionCacaoAmount = implProtectionCacaoAmount
			_, extraUnits, err = calculatePoolUnitsV1(pool.GetPoolUnits(), poolCacao, poolAsset, implProtectionCacaoAmount, cosmos.ZeroUint())
			if err != nil {
				return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), err
			}
			ctx.Logger().Info("liquidity provider granted imp loss protection", "extra provider units", extraUnits, "extra rune", implProtectionCacaoAmount)
			poolCacao = poolCacao.Add(implProtectionCacaoAmount)
			fLiquidityProviderUnit = fLiquidityProviderUnit.Add(extraUnits)
			pool.LPUnits = pool.LPUnits.Add(extraUnits)
		}
	}

	var withdrawCacao, withDrawAsset, unitAfter cosmos.Uint
	if pool.Asset.IsVaultAsset() {
		withdrawCacao, withDrawAsset, unitAfter = calculateVaultWithdrawV1(pool.GetPoolUnits(), poolAsset, originalLiquidityProviderUnits, msg.BasisPoints)
	} else {
		withdrawCacao, withDrawAsset, unitAfter, err = calculateWithdrawV91(pool.GetPoolUnits(), poolCacao, poolAsset, originalLiquidityProviderUnits, extraUnits, msg.BasisPoints, assetToWithdraw)
		if err != nil {
			ctx.Logger().Error("fail to withdraw", "error", err)
			return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), errWithdrawFail
		}
	}
	if !pool.Asset.IsVaultAsset() {
		if (withdrawCacao.Equal(poolCacao) && !withDrawAsset.Equal(poolAsset)) || (!withdrawCacao.Equal(poolCacao) && withDrawAsset.Equal(poolAsset)) {
			ctx.Logger().Error("fail to withdraw: cannot withdraw 100% of only one side of the pool")
			return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), errWithdrawFail
		}
	}

	lp, err = checkWithdrawLimit(ctx, mgr, msg, cv, lp)
	if err != nil {
		ctx.Logger().Error("fail to withdraw", "error", err)
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), err
	}

	withDrawAsset = cosmos.RoundToDecimal(withDrawAsset, pool.Decimals)
	gasAsset := cosmos.ZeroUint()
	// If the pool is empty, and there is a gas asset, subtract required gas
	if common.SafeSub(pool.GetPoolUnits(), fLiquidityProviderUnit).Add(unitAfter).IsZero() {
		var maxGas common.Coin
		maxGas, err = mgr.GasMgr().GetMaxGas(ctx, pool.Asset.GetChain())
		if err != nil {
			ctx.Logger().Error("fail to get gas for asset", "asset", pool.Asset, "error", err)
			return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), errWithdrawFail
		}
		// minus gas costs for our transactions
		// TODO: chain specific logic should be in a single location
		if pool.Asset.IsBNB() && !common.BaseAsset().Chain.Equals(common.BASEChain) {
			originalAsset := withDrawAsset
			withDrawAsset = common.SafeSub(
				withDrawAsset,
				maxGas.Amount.MulUint64(2), // RUNE asset is on binance chain
			)
			gasAsset = originalAsset.Sub(withDrawAsset)
		} else if pool.Asset.GetChain().GetGasAsset().Equals(pool.Asset) {
			gasAsset = maxGas.Amount
			if gasAsset.GT(withDrawAsset) {
				gasAsset = withDrawAsset
			}
			withDrawAsset = common.SafeSub(withDrawAsset, gasAsset)
		}
	}

	ctx.Logger().Info("client withdraw", "RUNE", withdrawCacao, "asset", withDrawAsset, "units left", unitAfter)
	// update pool
	pool.LPUnits = common.SafeSub(pool.LPUnits, common.SafeSub(fLiquidityProviderUnit, unitAfter))
	pool.BalanceCacao = common.SafeSub(poolCacao, withdrawCacao)
	pool.BalanceAsset = common.SafeSub(poolAsset, withDrawAsset)

	ctx.Logger().Info("pool after withdraw", "pool unit", pool.GetPoolUnits(), "balance RUNE", pool.BalanceCacao, "balance asset", pool.BalanceAsset)

	lp.LastWithdrawHeight = ctx.BlockHeight()
	maxPts := cosmos.NewUint(uint64(MaxWithdrawBasisPoints))
	lp.CacaoDepositValue = common.SafeSub(lp.CacaoDepositValue, common.GetSafeShare(msg.BasisPoints, maxPts, lp.CacaoDepositValue))
	lp.AssetDepositValue = common.SafeSub(lp.AssetDepositValue, common.GetSafeShare(msg.BasisPoints, maxPts, lp.AssetDepositValue))
	lp.Units = unitAfter

	// sanity check, we don't increase LP units
	if unitAfter.GTE(originalLiquidityProviderUnits) {
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), ErrInternal(err, fmt.Sprintf("sanity check: LP units cannot increase during a withdrawal: %d --> %d", originalLiquidityProviderUnits.Uint64(), unitAfter.Uint64()))
	}

	// Create a pool event if THORNode have no rune or assets
	if (pool.BalanceAsset.IsZero() || pool.BalanceCacao.IsZero()) && !pool.Asset.IsVaultAsset() {
		poolEvt := NewEventPool(pool.Asset, PoolStaged)
		if err := mgr.EventMgr().EmitEvent(ctx, poolEvt); nil != err {
			ctx.Logger().Error("fail to emit pool event", "error", err)
		}
		pool.Status = PoolStaged
	}

	if err := mgr.Keeper().SetPool(ctx, pool); err != nil {
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), ErrInternal(err, "fail to save pool")
	}
	if mgr.Keeper().RagnarokInProgress(ctx) {
		mgr.Keeper().SetLiquidityProvider(ctx, lp)
	} else {
		if !lp.Units.Add(lp.PendingAsset).Add(lp.PendingCacao).IsZero() {
			mgr.Keeper().SetLiquidityProvider(ctx, lp)
		} else {
			mgr.Keeper().RemoveLiquidityProvider(ctx, lp)
		}
	}
	// add rune from the reserve to the asgard module, to cover imp loss protection
	if !protectionCacaoAmount.IsZero() {
		err := mgr.Keeper().SendFromModuleToModule(ctx, ReserveName, AsgardName, common.NewCoins(common.NewCoin(common.BaseAsset(), protectionCacaoAmount)))
		if err != nil {
			return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint(), ErrInternal(err, "fail to move imp loss protection rune from the reserve to asgard")
		}
	}
	return withdrawCacao, withDrawAsset, protectionCacaoAmount, common.SafeSub(originalLiquidityProviderUnits, unitAfter), gasAsset, nil
}

func validateWithdrawV105(ctx cosmos.Context, keeper keeper.Keeper, msg MsgWithdrawLiquidity) error {
	if msg.WithdrawAddress.IsEmpty() {
		return errors.New("empty withdraw address")
	}
	if msg.Tx.ID.IsEmpty() {
		return errors.New("request tx hash is empty")
	}
	if msg.Asset.IsEmpty() {
		return errors.New("empty asset")
	}
	withdrawBasisPoints := msg.BasisPoints
	// when BasisPoints is zero, it will be override in parse memo, so if a message can get here
	// the witdrawBasisPoints must between 1~MaxWithdrawBasisPoints
	if !withdrawBasisPoints.GT(cosmos.ZeroUint()) || withdrawBasisPoints.GT(cosmos.NewUint(MaxWithdrawBasisPoints)) {
		return fmt.Errorf("withdraw basis points %s is invalid", msg.BasisPoints)
	}
	if !keeper.PoolExist(ctx, msg.Asset) {
		// pool doesn't exist
		return fmt.Errorf("pool-%s doesn't exist", msg.Asset)
	}

	lp, err := keeper.GetLiquidityProvider(ctx, msg.Asset, msg.WithdrawAddress)
	if err != nil {
		return fmt.Errorf("fail to get liquidity provider: %w", err)
	}
	unitsToClaim := common.GetSafeShare(withdrawBasisPoints, cosmos.NewUint(MaxWithdrawBasisPoints), lp.Units)
	remainingUnits := lp.GetRemainingUnits()
	if unitsToClaim.GT(remainingUnits) {
		return fmt.Errorf("some units are bonded, withdrawing %s basis points exceeds remaining %s units", msg.BasisPoints, remainingUnits)
	}

	return nil
}

// calcImpLossV102 if there needs to add some imp loss protection, in rune
func calcImpLossV102(ctx sdk.Context, mgr Manager, lastAddHeight int64, lp LiquidityProvider, withdrawBasisPoints cosmos.Uint, fullProtectionLine int64, pool Pool) (cosmos.Uint, cosmos.Uint, cosmos.Uint) {
	/*
		A0 = assetDepositValue; R0 = runeDepositValue;

		liquidityUnits = units the member wishes to redeem after applying withdrawBasisPoints
		A1 = GetUncappedShare(liquidityUnits, lpUnits, assetDepth);
		R1 = GetUncappedShare(liquidityUnits, lpUnits, runeDepth);
		P1 = R1/A1
		coverage = ((A0 * P1) + R0) - ((A1 * P1) + R1) => ((A0 * R1/A1) + R0) - (R1 + R1)
	*/
	A0 := lp.AssetDepositValue
	R0 := lp.CacaoDepositValue
	poolUnits := pool.GetPoolUnits()
	A1 := common.GetSafeShare(lp.Units, poolUnits, pool.BalanceAsset)
	R1 := common.GetSafeShare(lp.Units, poolUnits, pool.BalanceCacao)
	if A1.IsZero() {
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint()
	}
	depositValue := A0.Mul(R1).Quo(A1).Add(R0)
	redeemValue := R1.Add(R1)
	coverage := common.SafeSub(depositValue, redeemValue)

	// taking withdrawBasisPoints, calculate how much of the coverage the user should receives
	coverage = common.GetSafeShare(withdrawBasisPoints, cosmos.NewUint(10000), coverage)

	/* Check current ratio vs the one when add was done. If current is bigger, meaning:
		        ASSET decreases in price relative to CACAO
						              OR
		        CACAO increases in price relative to ASSET
	   the fullProtectionLine is multiply by 4 so it goes from 100 to 400 days
	*/
	prevAssetDec := cosmos.NewDec(A0.BigInt().Int64())
	prevCACAODec := cosmos.NewDec(R0.BigInt().Int64())
	currentAssetDec := cosmos.NewDec(A1.BigInt().Int64())
	currentCACAODec := cosmos.NewDec(R1.BigInt().Int64())

	if prevAssetDec.IsZero() || currentAssetDec.IsZero() {
		ctx.Logger().Info("calcImpLossV102: prevAssetInt or currentAssetInt is zero")
		return cosmos.ZeroUint(), cosmos.ZeroUint(), cosmos.ZeroUint()
	}

	prevRatio := prevCACAODec.Quo(prevAssetDec)
	currRatio := currentCACAODec.Quo(currentAssetDec)

	if currRatio.LTE(prevRatio) {
		fullProtectionLineTimes4, err := mgr.Keeper().GetMimir(ctx, constants.FullImpLossProtectionBlocksTimes4.String())
		if err != nil || fullProtectionLineTimes4 < 0 {
			fullProtectionLineTimes4 = mgr.GetConstants().GetInt64Value(constants.FullImpLossProtectionBlocksTimes4)
		}
		fullProtectionLine = fullProtectionLineTimes4
	}

	protectionBasisPoints := calcImpLossProtectionAmtV2(ctx, mgr, lastAddHeight, fullProtectionLine)

	// taking protection basis points, calculate how much of the coverage the user actually receives
	result := coverage.MulUint64(uint64(protectionBasisPoints)).QuoUint64(10000)
	return result, depositValue, redeemValue
}

// calculate percentage (in basis points) of the amount of impermanent loss protection
func calcImpLossProtectionAmtV2(ctx cosmos.Context, mgr Manager, lastDepositHeight, target int64) int64 {
	age := ctx.BlockHeight() - lastDepositHeight
	zeroILPBlocks, err := mgr.Keeper().GetMimir(ctx, constants.ZeroImpLossProtectionBlocks.String())
	if err != nil || zeroILPBlocks < 0 {
		zeroILPBlocks = mgr.GetConstants().GetInt64Value(constants.ZeroImpLossProtectionBlocks)
	}

	if age < zeroILPBlocks { // set minimum age to ZeroImpLossProtectionBlocks
		return 0
	}
	if age >= target {
		return 10000
	}
	return ((age - zeroILPBlocks) * 10000) / (target - zeroILPBlocks)
}

func checkWithdrawLimit(ctx cosmos.Context, mgr Manager, msg MsgWithdrawLiquidity, cv constants.ConstantValues, lp types.LiquidityProvider) (types.LiquidityProvider, error) {
	// This function will only be check if the lp is on the withdraw limit days or is not Ragnarök
	if !isWithinWithdrawDaysLimit(ctx, mgr, cv, lp.CacaoAddress) || mgr.Keeper().RagnarokInProgress(ctx) || msg.Tx.Memo == "Ragnarok" {
		return lp, nil
	}

	// You can only withdraw certain percentage of your total every 24 hours
	blockThreshold := cosmos.NewUint((uint64)(lp.LastWithdrawCounterHeight + cv.GetInt64Value(constants.BlocksPerDay)))
	currentBlock := cosmos.NewUint((uint64)(ctx.BlockHeight()))
	addRune := lp.WithdrawCounter.Add(msg.BasisPoints)

	// Set LastWithdrawCounterHeight/WithdrawCounter
	if lp.WithdrawCounter.IsZero() {
		lp.LastWithdrawCounterHeight = ctx.BlockHeight()
	} else if !currentBlock.LT(blockThreshold) {
		lp.WithdrawCounter = cosmos.ZeroUint()
		lp.LastWithdrawCounterHeight = ctx.BlockHeight()
		addRune = msg.BasisPoints
	}

	// Get withdraw limit depending on the lp tier
	withdrawLimit, err := getWithdrawLimit(ctx, mgr, cv, lp.CacaoAddress)
	if err != nil {
		return lp, err
	}

	// Check limit
	if lp.WithdrawCounter.GTE(sdk.NewUint((uint64)(withdrawLimit))) {
		return lp, errMaxWithdrawReach
	}
	if addRune.GT(cosmos.NewUint((uint64)(withdrawLimit))) {
		return lp, errMaxWithdrawWillBeReach
	} else {
		lp.WithdrawCounter = addRune
	}

	return lp, nil
}

func assetToWithdrawV89(msg MsgWithdrawLiquidity, lp LiquidityProvider, pauseAsym int64) common.Asset {
	if lp.CacaoAddress.IsEmpty() {
		return msg.Asset
	}
	if lp.AssetAddress.IsEmpty() {
		return common.BaseAsset()
	}
	if pauseAsym > 0 {
		return common.EmptyAsset
	}

	return msg.WithdrawalAsset
}

func calcAsymWithdrawalV1(s, t, a cosmos.Uint) cosmos.Uint {
	// share = (s * A * (2 * T^2 - 2 * T * s + s^2))/T^3
	// s = liquidity provider units for member (after factoring in withdrawBasisPoints)
	// T = totalPoolUnits for pool
	// A = assetDepth to be withdrawn
	// (part1 * (part2 - part3 + part4)) / part5
	part1 := s.Mul(a)
	part2 := t.Mul(t).MulUint64(2)
	part3 := t.Mul(s).MulUint64(2)
	part4 := s.Mul(s)
	numerator := part1.Mul(common.SafeSub(part2, part3).Add(part4))
	part5 := t.Mul(t).Mul(t)
	return numerator.Quo(part5)
}
