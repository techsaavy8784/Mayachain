package mayachain

import (
	"context"
	"fmt"
	"strings"

	"github.com/armon/go-metrics"
	"github.com/blang/semver"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
)

// CommonOutboundTxHandler is the place where those common logic can be shared
// between multiple different kind of outbound tx handler
// at the moment, handler_refund, and handler_outbound_tx are largely the same
// , only some small difference
type CommonOutboundTxHandler struct {
	mgr Manager
}

// NewCommonOutboundTxHandler create a new instance of the CommonOutboundTxHandler
func NewCommonOutboundTxHandler(mgr Manager) CommonOutboundTxHandler {
	return CommonOutboundTxHandler{
		mgr: mgr,
	}
}

func (h CommonOutboundTxHandler) slashV96(ctx cosmos.Context, tx ObservedTx) error {
	toSlash := make(common.Coins, len(tx.Tx.Coins))
	copy(toSlash, tx.Tx.Coins)
	toSlash = toSlash.Adds_deprecated(tx.Tx.Gas.ToCoins())

	ctx = ctx.WithContext(context.WithValue(ctx.Context(), constants.CtxMetricLabels, []metrics.Label{ // nolint
		telemetry.NewLabel("reason", "failed_outbound"),
		telemetry.NewLabel("chain", string(tx.Tx.Chain)),
	}))

	return h.mgr.Slasher().SlashVaultToLP(ctx, tx.ObservedPubKey, toSlash, h.mgr, true)
}

func (h CommonOutboundTxHandler) handle(ctx cosmos.Context, tx ObservedTx, inTxID common.TxID) (*cosmos.Result, error) {
	version := h.mgr.GetVersion()
	switch {
	case version.GTE(semver.MustParse("1.96.0")):
		return h.handleV96(ctx, tx, inTxID)
	default:
		return nil, errBadVersion
	}
}

func (h CommonOutboundTxHandler) handleV96(ctx cosmos.Context, tx ObservedTx, inTxID common.TxID) (*cosmos.Result, error) {
	// note: Outbound tx usually it is related to an inbound tx except migration
	// thus here try to get the ObservedTxInVoter,  and set the tx out hash accordingly
	voter, err := h.mgr.Keeper().GetObservedTxInVoter(ctx, inTxID)
	if err != nil {
		return nil, ErrInternal(err, "fail to get observed tx voter")
	}
	if voter.AddOutTx(tx.Tx) {
		if err = h.mgr.EventMgr().EmitEvent(ctx, NewEventOutbound(inTxID, tx.Tx)); err != nil {
			return nil, ErrInternal(err, "fail to emit outbound event")
		}
	}
	h.mgr.Keeper().SetObservedTxInVoter(ctx, voter)

	if tx.Tx.Chain.Equals(common.BASEChain) {
		return &cosmos.Result{}, nil
	}

	shouldSlash := true
	signingTransPeriod := h.mgr.GetConstants().GetInt64Value(constants.SigningTransactionPeriod)
	// every Signing Transaction Period , THORNode will check whether a
	// TxOutItem had been sent by signer or not
	// if a txout item that is older than SigningTransactionPeriod, but has not
	// been sent out by signer , then it will create a new TxOutItem
	// here THORNode will have to mark all the TxOutItem to complete one the tx
	// get processed
	outHeight := voter.OutboundHeight
	if outHeight == 0 {
		outHeight = voter.FinalisedHeight
	}
	for height := outHeight; height <= ctx.BlockHeight(); height += signingTransPeriod {

		if height < ctx.BlockHeight()-signingTransPeriod {
			ctx.Logger().Info("Expired outbound transaction, should slash")
			continue
		}

		// update txOut record with our TxID that sent funds out of the pool
		var txOut *TxOut
		txOut, err = h.mgr.Keeper().GetTxOut(ctx, height)
		if err != nil {
			ctx.Logger().Error("unable to get txOut record", "error", err)
			return nil, cosmos.ErrUnknownRequest(err.Error())
		}

		// Save TxOut back with the TxID only when the TxOut on the block height is
		// not empty
		for i, txOutItem := range txOut.TxArray {
			// withdraw , refund etc, one inbound tx might result two outbound
			// txes, THORNode have to correlate outbound tx back to the
			// inbound, and also txitem , thus THORNode could record both
			// outbound tx hash correctly given every tx item will only have
			// one coin in it , THORNode could use that to identify which tx it
			// is

			if txOutItem.InHash.Equals(inTxID) &&
				txOutItem.OutHash.IsEmpty() &&
				tx.Tx.Chain.Equals(txOutItem.Chain) &&
				tx.Tx.ToAddress.Equals(txOutItem.ToAddress) &&
				strings.EqualFold(tx.Aggregator, txOutItem.Aggregator) &&
				strings.EqualFold(tx.AggregatorTarget, txOutItem.AggregatorTargetAsset) &&
				tx.ObservedPubKey.Equals(txOutItem.VaultPubKey) {

				matchCoin := tx.Tx.Coins.EqualsEx(common.Coins{txOutItem.Coin})
				if !matchCoin {
					// In case the mismatch is caused by decimals , round the tx out item's amount , and compare it again
					p, err := h.mgr.Keeper().GetPool(ctx, txOutItem.Coin.Asset)
					if err != nil {
						ctx.Logger().Error("fail to get pool", "error", err)
					}
					if !p.IsEmpty() {
						matchCoin = tx.Tx.Coins.EqualsEx(common.Coins{
							common.NewCoin(txOutItem.Coin.Asset, cosmos.RoundToDecimal(txOutItem.Coin.Amount, p.Decimals)),
						})
					}
				}
				// when outbound is gas asset
				if !matchCoin && txOutItem.Coin.Asset.Equals(txOutItem.Chain.GetGasAsset()) {
					asset := txOutItem.Chain.GetGasAsset()
					intendToSpend := txOutItem.Coin.Amount.Add(txOutItem.MaxGas.ToCoins().GetCoin(asset).Amount)
					actualSpend := tx.Tx.Coins.GetCoin(asset).Amount.Add(tx.Tx.Gas.ToCoins().GetCoin(asset).Amount)
					if intendToSpend.Equal(actualSpend) {
						matchCoin = true
						maxGasAmt := txOutItem.MaxGas.ToCoins().GetCoin(asset).Amount
						realGasAmt := tx.Tx.Gas.ToCoins().GetCoin(asset).Amount
						ctx.Logger().Info("override match coin", "intend to spend", intendToSpend, "actual spend", actualSpend, "max_gas", maxGasAmt, "actual gas", realGasAmt)
						if maxGasAmt.GT(realGasAmt) {
							// the outbound spend less than MaxGas
							diffGas := maxGasAmt.Sub(realGasAmt)
							h.mgr.GasMgr().AddGasAsset(common.Gas{
								common.NewCoin(asset, diffGas),
							}, false)
						} else if maxGasAmt.LT(realGasAmt) {
							// signer spend more than the maximum gas prescribed by BASEChain , slash it
							ctx.Logger().Info("slash node", "max gas", maxGasAmt, "real gas spend", realGasAmt, "gap", common.SafeSub(realGasAmt, maxGasAmt).String())
							matchCoin = false
						}
					}
				}
				if txOutItem.Chain.Equals(common.ETHChain) {
					maxGasAmount := txOutItem.MaxGas.ToCoins().GetCoin(common.ETHAsset).Amount
					gasAmount := tx.Tx.Gas.ToCoins().GetCoin(common.ETHAsset).Amount
					// If thornode instruct bifrost to spend more than MaxETHGas , then it should not be slashed.
					if gasAmount.GTE(cosmos.NewUint(constants.MaxETHGas)) && maxGasAmount.LT(cosmos.NewUint(constants.MaxETHGas)) {
						ctx.Logger().Info("ETH chain transaction spend more than MaxETHGas should be slashed", "gas", gasAmount.String(), "max gas", constants.MaxETHGas)
						matchCoin = false
					}
				}

				if !matchCoin {
					continue
				}
				txOut.TxArray[i].OutHash = tx.Tx.ID
				shouldSlash = false
				if err := h.mgr.Keeper().SetTxOut(ctx, txOut); err != nil {
					ctx.Logger().Error("fail to save tx out", "error", err)
				}
				break

			}
		}
	}

	if shouldSlash {
		ctx.Logger().Info("slash node account, no matched tx out item", "inbound txid", inTxID, "outbound tx", tx.Tx)

		// send security alert for events that are not eth burn
		if len(tx.Tx.Coins) != 1 || !tx.Tx.Coins[0].Amount.Equal(types.NewUint(1)) || !tx.Tx.Coins[0].Asset.Equals(common.ETHAsset) {
			msg := fmt.Sprintf("missing tx out in=%s", inTxID)
			if err := h.mgr.EventMgr().EmitEvent(ctx, NewEventSecurity(tx.Tx, msg)); err != nil {
				ctx.Logger().Error("fail to emit security event", "error", err)
			}
		}

		if err := h.slashV96(ctx, tx); err != nil {
			return nil, ErrInternal(err, "fail to slash account")
		}
	}

	if err := h.mgr.Keeper().SetLastSignedHeight(ctx, voter.FinalisedHeight); err != nil {
		ctx.Logger().Info("fail to update last signed height", "error", err)
	}

	return &cosmos.Result{}, nil
}
