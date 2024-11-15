package mayachain

import (
	"errors"
	"fmt"

	"github.com/blang/semver"
	se "github.com/cosmos/cosmos-sdk/types/errors"
	tmtypes "github.com/tendermint/tendermint/types"
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
)

func (h DepositHandler) handleV90(ctx cosmos.Context, msg MsgDeposit) (*cosmos.Result, error) {
	version := h.mgr.GetVersion()
	var haltHeight int64
	var err error

	// FIXME Remove when a more critical change in handler_deposit.go is made
	switch {
	case version.GTE(semver.MustParse(("1.105.0"))):
		haltHeight, err = h.mgr.Keeper().GetMimir(ctx, "HaltMAYAChain")
		if err != nil {
			return nil, fmt.Errorf("failed to get mimir setting: %w", err)
		}
	default:
		haltHeight, err = h.mgr.Keeper().GetMimir(ctx, "HaltTHORChain")
		if err != nil {
			return nil, fmt.Errorf("failed to get mimir setting: %w", err)
		}
	}

	if haltHeight > 0 && ctx.BlockHeight() > haltHeight {
		return nil, fmt.Errorf("mimir has halted MAYAChain transactions")
	}

	nativeTxFee, err := h.mgr.Keeper().GetMimir(ctx, constants.NativeTransactionFee.String())
	if err != nil || nativeTxFee < 0 {
		nativeTxFee = h.mgr.GetConstants().GetInt64Value(constants.NativeTransactionFee)
	}
	gas := common.NewCoin(common.BaseNative, cosmos.NewUint(uint64(nativeTxFee)))
	gasFee, err := gas.Native()
	if err != nil {
		return nil, fmt.Errorf("fail to get gas fee: %w", err)
	}

	coins, err := msg.Coins.Native()
	if err != nil {
		return nil, ErrInternal(err, "coins are native to BASEChain")
	}

	totalCoins := cosmos.NewCoins(gasFee).Add(coins...)
	if !h.mgr.Keeper().HasCoins(ctx, msg.GetSigners()[0], totalCoins) {
		return nil, cosmos.ErrInsufficientCoins(err, "insufficient funds")
	}

	memo, _ := ParseMemoWithMAYANames(ctx, h.mgr.Keeper(), msg.Memo) // ignore err
	if memo.IsOutbound() || memo.IsInternal() {
		return nil, fmt.Errorf("cannot send inbound an outbound or internal transaction")
	}

	// Only calculate percentages if the msg is not bond, otherwise send 100% to Reserve
	switch memo.GetType() {
	case TxBond, TxUnBond, TxLeave:
		// send gas to reserve
		sdkErr := h.mgr.Keeper().SendFromAccountToModule(ctx, msg.GetSigners()[0], ReserveName, common.NewCoins(gas))
		if sdkErr != nil {
			return nil, fmt.Errorf("unable to send gas to reserve: %w", sdkErr)
		}
	default:
		// Calculate Maya Fund -->  gasFee = 90%, Maya Fund = 10%
		newGas, mayaGas := CalculateMayaFundPercentage(gas, h.mgr)

		// send gas to reserve
		sdkErr := h.mgr.Keeper().SendFromAccountToModule(ctx, msg.GetSigners()[0], ReserveName, common.NewCoins(newGas))
		if sdkErr != nil {
			return nil, fmt.Errorf("unable to send gas to reserve: %w", sdkErr)
		}

		// send corresponding fees to Maya Fund
		sdkErr = h.mgr.Keeper().SendFromAccountToModule(ctx, msg.GetSigners()[0], MayaFund, common.NewCoins(mayaGas))
		if sdkErr != nil {
			return nil, fmt.Errorf("unable to send gas to maya fund: %w", sdkErr)
		}

	}

	hash := tmtypes.Tx(ctx.TxBytes()).Hash()
	txID, err := common.NewTxID(fmt.Sprintf("%X", hash))
	if err != nil {
		return nil, fmt.Errorf("fail to get tx hash: %w", err)
	}
	from, err := common.NewAddress(msg.GetSigners()[0].String())
	if err != nil {
		return nil, fmt.Errorf("fail to get from address: %w", err)
	}

	handler := NewInternalHandler(h.mgr)

	var targetModule string
	switch memo.GetType() {
	case TxBond, TxUnBond, TxLeave:
		targetModule = BondName
	case TxReserve, TxMAYAName:
		targetModule = ReserveName
	default:
		targetModule = AsgardName
	}
	coinsInMsg := msg.Coins
	if !coinsInMsg.IsEmpty() {
		// send funds to target module
		sdkErr := h.mgr.Keeper().SendFromAccountToModule(ctx, msg.GetSigners()[0], targetModule, msg.Coins)
		if sdkErr != nil {
			return nil, sdkErr
		}
	}

	to, err := h.mgr.Keeper().GetModuleAddress(targetModule)
	if err != nil {
		return nil, fmt.Errorf("fail to get to address: %w", err)
	}

	tx := common.NewTx(txID, from, to, coinsInMsg, common.Gas{gas}, msg.Memo)
	tx.Chain = common.BASEChain

	// construct msg from memo
	txIn := ObservedTx{Tx: tx}
	txInVoter := NewObservedTxVoter(txIn.Tx.ID, []ObservedTx{txIn})
	activeNodes, err := h.mgr.Keeper().ListActiveValidators(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to get all active nodes: %w", err)
	}
	for _, node := range activeNodes {
		txInVoter.Add(txIn, node.NodeAddress)
	}
	txInVoter.FinalisedHeight = ctx.BlockHeight()
	txInVoter.Tx = txInVoter.GetTx(activeNodes)
	h.mgr.Keeper().SetObservedTxInVoter(ctx, txInVoter)
	m, txErr := processOneTxIn(ctx, h.mgr.GetVersion(), h.mgr.Keeper(), txIn, msg.Signer)
	if txErr != nil {
		ctx.Logger().Error("fail to process native inbound tx", "error", txErr.Error(), "tx hash", tx.ID.String())
		if txIn.Tx.Coins.IsEmpty() {
			return &cosmos.Result{}, nil
		}
		if newErr := refundTx(ctx, txIn, h.mgr, CodeInvalidMemo, txErr.Error(), targetModule); nil != newErr {
			return nil, newErr
		}

		return &cosmos.Result{}, nil
	}

	// check if we've halted trading
	_, isSwap := m.(*MsgSwap)
	_, isAddLiquidity := m.(*MsgAddLiquidity)
	if isSwap || isAddLiquidity {
		if isSwap && isLiquidityAuction(ctx, h.mgr.Keeper()) {
			if newErr := refundTx(ctx, txIn, h.mgr, se.ErrUnauthorized.ABCICode(), "cannot swap, liquidity auction enabled", targetModule); nil != newErr {
				return nil, ErrInternal(newErr, "liquidity auction enabled, fail to refund")
			}
			return &cosmos.Result{}, nil
		}

		if isTradingHalt(ctx, m, h.mgr) || h.mgr.Keeper().RagnarokInProgress(ctx) {
			if txIn.Tx.Coins.IsEmpty() {
				return &cosmos.Result{}, nil
			}
			if newErr := refundTx(ctx, txIn, h.mgr, se.ErrUnauthorized.ABCICode(), "trading halted", targetModule); nil != newErr {
				return nil, ErrInternal(newErr, "trading is halted, fail to refund")
			}
			return &cosmos.Result{}, nil
		}
	}

	// if its a swap, send it to our queue for processing later
	if isSwap {
		msg, ok := m.(*MsgSwap)
		if ok {
			h.addSwap(ctx, *msg)
		}
		return &cosmos.Result{}, nil
	}

	result, err := handler(ctx, m)
	if err != nil {
		code := uint32(1)
		var e se.Error
		if errors.As(err, &e) {
			code = e.ABCICode()
		}
		if txIn.Tx.Coins.IsEmpty() {
			return &cosmos.Result{}, nil
		}
		if err = refundTx(ctx, txIn, h.mgr, code, err.Error(), targetModule); err != nil {
			return nil, fmt.Errorf("fail to refund tx: %w", err)
		}
		return &cosmos.Result{}, nil
	}
	// for those Memo that will not have outbound at all , set the observedTx to done
	if !memo.GetType().HasOutbound() {
		txInVoter.SetDone()
		h.mgr.Keeper().SetObservedTxInVoter(ctx, txInVoter)
	}
	return result, nil
}

func (h DepositHandler) validateV1(ctx cosmos.Context, msg MsgDeposit) error {
	return msg.ValidateBasic()
}
