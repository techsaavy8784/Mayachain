package mayachain

import (
	"errors"
	"fmt"

	"github.com/blang/semver"
	"github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
)

// UnBondHandler a handler to process unbond request
type UnBondHandler struct {
	mgr Manager
}

// NewUnBondHandler create new UnBondHandler
func NewUnBondHandler(mgr Manager) UnBondHandler {
	return UnBondHandler{
		mgr: mgr,
	}
}

// Run execute the handler
func (h UnBondHandler) Run(ctx cosmos.Context, m cosmos.Msg) (*cosmos.Result, error) {
	msg, ok := m.(*MsgUnBond)
	if !ok {
		return nil, errInvalidMessage
	}
	ctx.Logger().Info("receive MsgUnBond",
		"node address", msg.NodeAddress,
		"request hash", msg.TxIn.ID)
	if err := h.validate(ctx, *msg); err != nil {
		ctx.Logger().Error("msg unbond fail validation", "error", err)
		return nil, err
	}
	if err := h.handle(ctx, *msg); err != nil {
		ctx.Logger().Error("msg unbond fail handler", "error", err)
		return nil, err
	}

	return &cosmos.Result{}, nil
}

func (h UnBondHandler) validate(ctx cosmos.Context, msg MsgUnBond) error {
	version := h.mgr.GetVersion()
	switch {
	case version.GTE(semver.MustParse("1.105.0")):
		return h.validateV105(ctx, msg)
	case version.GTE(semver.MustParse("1.88.0")):
		return h.validateV88(ctx, msg)
	default:
		return errBadVersion
	}
}

func (h UnBondHandler) validateV105(ctx cosmos.Context, msg MsgUnBond) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	na, err := h.mgr.Keeper().GetNodeAccount(ctx, msg.NodeAddress)
	if err != nil {
		return ErrInternal(err, fmt.Sprintf("fail to get node account(%s)", msg.NodeAddress))
	}

	if na.Status == NodeActive || na.Status == NodeReady {
		return cosmos.ErrUnknownRequest("cannot unbond while node is in active or ready status")
	}

	if h.mgr.GetVersion().GTE(semver.MustParse("1.88.1")) {
		if fetchConfigInt64(ctx, h.mgr, constants.PauseUnbond) > 0 {
			return ErrInternal(err, "unbonding has been paused")
		}
	}

	if h.mgr.Keeper().VaultExists(ctx, na.PubKeySet.Secp256k1) {
		var ygg Vault
		ygg, err = h.mgr.Keeper().GetVault(ctx, na.PubKeySet.Secp256k1)
		if err != nil {
			return err
		}
		if !ygg.IsYggdrasil() {
			return errors.New("this is not a Yggdrasil vault")
		}
	}

	jail, err := h.mgr.Keeper().GetNodeAccountJail(ctx, msg.NodeAddress)
	if err != nil {
		// ignore this error and carry on. Don't want a jail bug causing node
		// accounts to not be able to get their funds out
		ctx.Logger().Error("fail to get node account jail", "error", err)
	}
	if jail.IsJailed(ctx) {
		return fmt.Errorf("failed to unbond due to jail status: (release height %d) %s", jail.ReleaseHeight, jail.Reason)
	}

	bp, err := h.mgr.Keeper().GetBondProviders(ctx, msg.NodeAddress)
	if err != nil {
		return ErrInternal(err, fmt.Sprintf("fail to get bond providers(%s)", msg.NodeAddress))
	}
	from, err := msg.BondAddress.AccAddress()
	if err != nil {
		return ErrInternal(err, fmt.Sprintf("fail to parse bond address(%s)", msg.BondAddress))
	}
	if !bp.Has(from) && !na.BondAddress.Equals(msg.BondAddress) {
		return cosmos.ErrUnauthorized(fmt.Sprintf("%s are not authorized to manage %s", msg.BondAddress, msg.NodeAddress))
	}

	if !msg.Asset.IsEmpty() {
		var lp LiquidityProvider
		lp, err = h.mgr.Keeper().GetLiquidityProvider(ctx, msg.Asset, msg.BondAddress)
		if err != nil {
			return ErrInternal(err, fmt.Sprintf("fail to get liquidity provider(%s,%s)", msg.BondAddress, msg.NodeAddress))
		}

		bonded := lp.GetUnitsBondedToNode(msg.NodeAddress)
		if bonded.IsZero() {
			return cosmos.ErrUnknownRequest(fmt.Sprintf("address has not bonded to node %s with asset %s", msg.NodeAddress, msg.Asset))
		}
		if bonded.LT(msg.Units) {
			return cosmos.ErrUnknownRequest(fmt.Sprintf("request unbond units %s is more than bonded units(%s)", msg.Units, bonded))
		}

		liquidityPools := GetLiquidityPools(h.mgr.GetVersion())
		if found := common.ContainsAsset(msg.Asset, liquidityPools); !found {
			return cosmos.ErrUnknownRequest("asset is not in valid liquidity pools list")
		}
	}

	return nil
}

func (h UnBondHandler) handle(ctx cosmos.Context, msg MsgUnBond) error {
	version := h.mgr.GetVersion()
	switch {
	case version.GTE(semver.MustParse("1.108.0")):
		return h.handleV108(ctx, msg)
	case version.GTE(semver.MustParse("1.107.0")):
		return h.handleV107(ctx, msg)
	case version.GTE(semver.MustParse("1.105.0")):
		return h.handleV105(ctx, msg)
	case version.GTE(semver.MustParse("1.92.0")):
		return h.handleV92(ctx, msg)
	default:
		return errBadVersion
	}
}

func (h UnBondHandler) handleV108(ctx cosmos.Context, msg MsgUnBond) error {
	na, err := h.mgr.Keeper().GetNodeAccount(ctx, msg.NodeAddress)
	if err != nil {
		return ErrInternal(err, fmt.Sprintf("fail to get node account(%s)", msg.NodeAddress))
	}

	var ygg Vault
	if h.mgr.Keeper().VaultExists(ctx, na.PubKeySet.Secp256k1) {
		ygg, err = h.mgr.Keeper().GetVault(ctx, na.PubKeySet.Secp256k1)
		if err != nil {
			return err
		}
	}

	if ygg.HasFunds() {
		canUnbond := true
		totalRuneValue := cosmos.ZeroUint()
		for _, c := range ygg.Coins {
			if c.Amount.IsZero() {
				continue
			}
			if !c.Asset.IsGasAsset() {
				// None gas asset has not been sent back to asgard in full
				canUnbond = false
				break
			}
			chain := c.Asset.GetChain()
			var maxGas common.Coin
			maxGas, err = h.mgr.GasMgr().GetMaxGas(ctx, chain)
			if err != nil {
				ctx.Logger().Error("fail to get max gas", "chain", chain, "error", err)
				canUnbond = false
				break
			}
			// 10x the maxGas , if the amount of gas asset left in the yggdrasil vault is larger than 10x of the MaxGas , then we don't allow node to unbond
			if c.Amount.GT(maxGas.Amount.MulUint64(10)) {
				canUnbond = false
			}
			var pool Pool
			pool, err = h.mgr.Keeper().GetPool(ctx, c.Asset)
			if err != nil {
				ctx.Logger().Error("fail to get pool", "asset", c.Asset, "error", err)
				canUnbond = false
				break
			}
			totalRuneValue = totalRuneValue.Add(pool.AssetValueInRune(c.Amount))
		}
		if !canUnbond {
			ctx.Logger().Error("cannot unbond while yggdrasil vault still has funds")
			if err = h.mgr.ValidatorMgr().RequestYggReturn(ctx, na, h.mgr); err != nil {
				return ErrInternal(err, "fail to request yggdrasil return fund")
			}
			return nil
		}

		penaltyPts := fetchConfigInt64(ctx, h.mgr, constants.SlashPenalty)
		totalRuneValue = common.GetUncappedShare(cosmos.NewUint(uint64(penaltyPts)), cosmos.NewUint(10_000), totalRuneValue)
		_, _, err = h.mgr.Slasher().SlashNodeAccountLP(ctx, na, totalRuneValue)
		if err != nil {
			return ErrInternal(err, "fail to slash node account")
		}
	}

	bondLockPeriod, err := h.mgr.Keeper().GetMimir(ctx, constants.BondLockupPeriod.String())
	if err != nil || bondLockPeriod < 0 {
		bondLockPeriod = h.mgr.GetConstants().GetInt64Value(constants.BondLockupPeriod)
	}
	if ctx.BlockHeight()-na.StatusSince < bondLockPeriod {
		return fmt.Errorf("node can not unbond before %d", na.StatusSince+bondLockPeriod)
	}
	vaults, err := h.mgr.Keeper().GetAsgardVaultsByStatus(ctx, RetiringVault)
	if err != nil {
		return ErrInternal(err, "fail to get retiring vault")
	}
	isMemberOfRetiringVault := false
	for _, v := range vaults {
		if v.GetMembership().Contains(na.PubKeySet.Secp256k1) {
			isMemberOfRetiringVault = true
			ctx.Logger().Info("node account is still part of the retiring vault,can't return bond yet")
			break
		}
	}
	if isMemberOfRetiringVault {
		return ErrInternal(err, "fail to unbond, still part of the retiring vault")
	}

	from, err := cosmos.AccAddressFromBech32(msg.BondAddress.String())
	if err != nil {
		return ErrInternal(err, "fail to parse from address")
	}

	// remove/unbonding bond provider
	// check that 1) requester is node operator, 2) references
	var bondProviderAddress types.AccAddress
	if msg.BondAddress.Equals(na.BondAddress) && !msg.BondProviderAddress.Empty() {
		bondProviderAddress = msg.BondProviderAddress

		// either the asset is empty and we won't take the units into account (take all assets out) or specific asset is withdrawn (some part or all of it)
		if err = refundBond(ctx, msg.TxIn, msg.BondProviderAddress, msg.Asset, msg.Units, &na, h.mgr); err != nil {
			return ErrInternal(err, "fail to unbond")
		}
	} else {
		if err = refundBond(ctx, msg.TxIn, from, msg.Asset, msg.Units, &na, h.mgr); err != nil {
			return ErrInternal(err, "fail to unbond")
		}
		bondProviderAddress, err = msg.BondAddress.AccAddress()
		if err != nil {
			return ErrInternal(err, "fail to unbond, cannot get msg BondAddress")
		}
	}

	if !bondProviderAddress.Empty() {
		// remove bond provider (if bond is now zero)
		var bp BondProviders
		bp, err = h.mgr.Keeper().GetBondProviders(ctx, na.NodeAddress)
		if err != nil {
			return ErrInternal(err, fmt.Sprintf("fail to get bond providers(%s)", na.NodeAddress))
		}
		provider := bp.Get(bondProviderAddress)
		var providerBond cosmos.Uint
		providerBond, err = h.mgr.Keeper().CalcLPLiquidityBond(ctx, common.Address(provider.BondAddress.String()), na.NodeAddress)
		if err != nil {
			return ErrInternal(err, "fail to get bond provider liquidity")
		}
		if !provider.IsEmpty() && providerBond.IsZero() {
			if ok := bp.Remove(bondProviderAddress); ok {
				if err = h.mgr.Keeper().SetBondProviders(ctx, bp); err != nil {
					return ErrInternal(err, fmt.Sprintf("fail to save bond providers(%s)", bp.NodeAddress.String()))
				}
			}
		}
	}

	coin := msg.TxIn.Coins.GetCoin(common.BaseAsset())
	if !coin.IsEmpty() {
		na.Reward = na.Reward.Add(coin.Amount)
		if err = h.mgr.Keeper().SetNodeAccount(ctx, na); err != nil {
			return ErrInternal(err, "fail to save node account to key value store")
		}
	}

	return nil
}
