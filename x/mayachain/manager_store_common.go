package mayachain

import (
	"fmt"
	"strings"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
)

// removeTransactions is a method used to remove a tx out item in the queue
func removeTransactions(ctx cosmos.Context, mgr Manager, hashes ...string) {
	for _, txID := range hashes {
		inTxID, err := common.NewTxID(txID)
		if err != nil {
			ctx.Logger().Error("fail to parse tx id", "error", err, "tx_id", inTxID)
			continue
		}
		voter, err := mgr.Keeper().GetObservedTxInVoter(ctx, inTxID)
		if err != nil {
			ctx.Logger().Error("fail to get observed tx voter", "error", err)
			continue
		}
		// all outbound action get removed
		voter.Actions = []TxOutItem{}
		if voter.Tx.IsEmpty() {
			continue
		}
		voter.Tx.SetDone(common.BlankTxID, 0)
		// set the tx outbound with a blank txid will mark it as down , and will be skipped in the reschedule logic
		for idx := range voter.Txs {
			voter.Txs[idx].SetDone(common.BlankTxID, 0)
		}
		mgr.Keeper().SetObservedTxInVoter(ctx, voter)
	}
}

// nolint
type adhocRefundTx struct {
	inboundHash string
	toAddr      string
	amount      uint64
	asset       string
}

// refundTransactions is design to use store migration to refund adhoc transactions
// nolint
func refundTransactions(ctx cosmos.Context, mgr *Mgrs, pubKey string, adhocRefundTxes ...adhocRefundTx) {
	asgardPubKey, err := common.NewPubKey(pubKey)
	if err != nil {
		ctx.Logger().Error("fail to parse pub key", "error", err, "pubkey", pubKey)
		return
	}
	for _, item := range adhocRefundTxes {
		hash, err := common.NewTxID(item.inboundHash)
		if err != nil {
			ctx.Logger().Error("fail to parse hash", "hash", item.inboundHash, "error", err)
			continue
		}
		addr, err := common.NewAddress(item.toAddr)
		if err != nil {
			ctx.Logger().Error("fail to parse address", "address", item.toAddr, "error", err)
			continue
		}
		asset, err := common.NewAsset(item.asset)
		if err != nil {
			ctx.Logger().Error("fail to parse asset", "asset", item.asset, "error", err)
			continue
		}
		coin := common.NewCoin(asset, cosmos.NewUint(item.amount))
		maxGas, err := mgr.GasMgr().GetMaxGas(ctx, coin.Asset.GetChain())
		if err != nil {
			ctx.Logger().Error("fail to get max gas", "error", err)
			continue
		}
		toi := TxOutItem{
			Chain:       coin.Asset.GetChain(),
			InHash:      hash,
			ToAddress:   addr,
			Coin:        coin,
			Memo:        NewRefundMemo(hash).String(),
			MaxGas:      common.Gas{maxGas},
			GasRate:     int64(mgr.GasMgr().GetGasRate(ctx, coin.Asset.GetChain()).Uint64()),
			VaultPubKey: asgardPubKey,
		}

		voter, err := mgr.Keeper().GetObservedTxInVoter(ctx, toi.InHash)
		if err != nil {
			ctx.Logger().Error("fail to get observe tx in voter", "error", err)
			continue
		}
		voter.OutboundHeight = ctx.BlockHeight()
		voter.Actions = append(voter.Actions, toi)
		mgr.Keeper().SetObservedTxInVoter(ctx, voter)

		if err := mgr.TxOutStore().UnSafeAddTxOutItem(ctx, mgr, toi, ctx.BlockHeight()); err != nil {
			ctx.Logger().Error("fail to send manual refund", "address", item.toAddr, "error", err)
		}
	}
}

// When an ObservedTxInVoter has dangling Actions items swallowed by the vaults, requeue them.
func requeueDanglingActions(ctx cosmos.Context, mgr *Mgrs, txIDs []common.TxID) {
	// Select the least secure ActiveVault Asgard for all outbounds.
	// Even if it fails (as in if the version changed upon the keygens-complete block of a churn),
	// updating the voter's FinalisedHeight allows another MaxOutboundAttempts for LackSigning vault selection.
	activeAsgards, err := mgr.Keeper().GetAsgardVaultsByStatus(ctx, ActiveVault)
	if err != nil || len(activeAsgards) == 0 {
		ctx.Logger().Error("fail to get active asgard vaults", "error", err)
		return
	}
	if len(activeAsgards) > 1 {
		signingTransactionPeriod := mgr.GetConstants().GetInt64Value(constants.SigningTransactionPeriod)
		activeAsgards = mgr.Keeper().SortBySecurity(ctx, activeAsgards, signingTransactionPeriod)
	}
	vaultPubKey := activeAsgards[0].PubKey

	for _, txID := range txIDs {
		voter, err := mgr.Keeper().GetObservedTxInVoter(ctx, txID)
		if err != nil {
			ctx.Logger().Error("fail to get observed tx voter", "error", err)
			continue
		}

		if len(voter.OutTxs) >= len(voter.Actions) {
			log := fmt.Sprintf("(%d) OutTxs present for (%s), despite expecting fewer than the (%d) Actions.", len(voter.OutTxs), txID.String(), len(voter.Actions))
			ctx.Logger().Debug(log)
			continue
		}

		var indices []int
		for i := range voter.Actions {
			if isActionsItemDangling(voter, i) {
				indices = append(indices, i)
			}
		}
		if len(indices) == 0 {
			log := fmt.Sprintf("No dangling Actions item found for (%s).", txID.String())
			ctx.Logger().Debug(log)
			continue
		}

		if len(voter.Actions)-len(voter.OutTxs) != len(indices) {
			log := fmt.Sprintf("(%d) Actions and (%d) OutTxs present for (%s), yet there appeared to be (%d) dangling Actions.", len(voter.Actions), len(voter.OutTxs), txID.String(), len(indices))
			ctx.Logger().Debug(log)
			continue
		}

		// Update the voter's FinalisedHeight to give another MaxOutboundAttempts.
		voter.FinalisedHeight = ctx.BlockHeight()
		voter.OutboundHeight = ctx.BlockHeight()

		for _, index := range indices {
			// Use a pointer to update the voter as well.
			actionItem := &voter.Actions[index]

			// Update the vault pubkey.
			actionItem.VaultPubKey = vaultPubKey

			// Update the Actions item's MaxGas and GasRate.
			// Note that nothing in this function should require a GasManager BeginBlock.
			gasCoin, err := mgr.GasMgr().GetMaxGas(ctx, actionItem.Chain)
			if err != nil {
				ctx.Logger().Error("fail to get max gas", "chain", actionItem.Chain, "error", err)
				continue
			}
			actionItem.MaxGas = common.Gas{gasCoin}
			actionItem.GasRate = int64(mgr.GasMgr().GetGasRate(ctx, actionItem.Chain).Uint64())

			// UnSafeAddTxOutItem is used to queue the txout item directly, without for instance deducting another fee.
			err = mgr.TxOutStore().UnSafeAddTxOutItem(ctx, mgr, *actionItem, ctx.BlockHeight())
			if err != nil {
				ctx.Logger().Error("fail to add outbound tx", "error", err)
				continue
			}
		}

		// Having requeued all dangling Actions items, set the updated voter.
		mgr.Keeper().SetObservedTxInVoter(ctx, voter)
	}
}

func isActionsItemDangling(voter ObservedTxVoter, i int) bool {
	if i < 0 || i > len(voter.Actions)-1 {
		// No such Actions item exists in the voter.
		return false
	}

	toi := voter.Actions[i]

	// If any OutTxs item matches an Actions item, deem it to be not dangling.
	for _, outboundTx := range voter.OutTxs {
		// The comparison code is based on matchActionItem, as matchActionItem is unimportable.
		// note: Coins.Contains will match amount as well
		matchCoin := outboundTx.Coins.Contains(toi.Coin)
		if !matchCoin && toi.Coin.Asset.Equals(toi.Chain.GetGasAsset()) {
			asset := toi.Chain.GetGasAsset()
			intendToSpend := toi.Coin.Amount.Add(toi.MaxGas.ToCoins().GetCoin(asset).Amount)
			actualSpend := outboundTx.Coins.GetCoin(asset).Amount.Add(outboundTx.Gas.ToCoins().GetCoin(asset).Amount)
			if intendToSpend.Equal(actualSpend) {
				matchCoin = true
			}
		}
		if strings.EqualFold(toi.Memo, outboundTx.Memo) &&
			toi.ToAddress.Equals(outboundTx.ToAddress) &&
			toi.Chain.Equals(outboundTx.Chain) &&
			matchCoin {
			return false
		}
	}
	return true
}

// nolint
type unbondBondProvider struct {
	bondProviderAddress string
	nodeAccountAddress  string
}

func unbondBondProviders(ctx cosmos.Context, mgr *Mgrs, unbondBPAddresses []unbondBondProvider) {
	for _, unbondAddress := range unbondBPAddresses {
		nodeAcc, err := cosmos.AccAddressFromBech32(unbondAddress.nodeAccountAddress)
		if err != nil {
			ctx.Logger().Error("fail to parse address: %s", unbondAddress.nodeAccountAddress, "error", err)
		}

		bps, err := mgr.Keeper().GetBondProviders(ctx, nodeAcc)
		if err != nil {
			ctx.Logger().Error("fail to get bond providers(%s)", nodeAcc)
		}

		bpAcc, err := cosmos.AccAddressFromBech32(unbondAddress.bondProviderAddress)
		if err != nil {
			ctx.Logger().Error("fail to parse address: %s", unbondAddress.bondProviderAddress, "error", err)
		}

		provider := bps.Get(bpAcc)
		providerBond, err := mgr.Keeper().CalcLPLiquidityBond(ctx, common.Address(bpAcc.String()), nodeAcc)
		if err != nil {
			ctx.Logger().Error("fail to get bond provider liquidity: %s", err)
		}
		if !provider.IsEmpty() && providerBond.IsZero() {
			bps.Unbond(bpAcc)
			if ok := bps.Remove(bpAcc); ok {
				if err := mgr.Keeper().SetBondProviders(ctx, bps); err != nil {
					ctx.Logger().Error("fail to save bond providers(%s)", bpAcc, "error", err)
				}
			}
		}

	}
}

type RefundTxCACAO struct {
	sendAddress string
	amount      cosmos.Uint
}

func refundTxsCACAO(ctx cosmos.Context, mgr *Mgrs, refunds []RefundTxCACAO) {
	for _, refund := range refunds {
		if refund.amount.IsZero() {
			continue
		}

		acc, err := cosmos.AccAddressFromBech32(refund.sendAddress)
		if err != nil {
			ctx.Logger().Error("fail to parse address: %s", refund.sendAddress, "error", err)
			continue
		}

		if err := mgr.Keeper().SendFromModuleToAccount(ctx, ReserveName, acc, common.NewCoins(common.NewCoin(common.BaseNative, refund.amount))); err != nil {
			ctx.Logger().Error("fail to send provider reward: %s", refund.sendAddress, "error", err)
		}
	}
}
