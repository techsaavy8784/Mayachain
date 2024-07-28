package mayachain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// SlasherVCUR is v88 implementation of slasher
type SlasherVCUR struct {
	keeper   keeper.Keeper
	eventMgr EventManager
}

// newSlasherVCUR create a new instance of Slasher
func newSlasherVCUR(keeper keeper.Keeper, eventMgr EventManager) *SlasherVCUR {
	return &SlasherVCUR{keeper: keeper, eventMgr: eventMgr}
}

// BeginBlock called when a new block get proposed to detect whether there are duplicate vote
func (s *SlasherVCUR) BeginBlock(ctx cosmos.Context, req abci.RequestBeginBlock, constAccessor constants.ConstantValues) {
	// Iterate through any newly discovered evidence of infraction
	// Slash any validators (and since-unbonded liquidity within the unbonding period)
	// who contributed to valid infractions
	for _, evidence := range req.ByzantineValidators {
		switch evidence.Type {
		case abci.EvidenceType_DUPLICATE_VOTE:
			if err := s.HandleDoubleSign(ctx, evidence.Validator.Address, evidence.Height, constAccessor); err != nil {
				ctx.Logger().Error("fail to slash for double signing a block", "error", err)
			}
		default:
			ctx.Logger().Error("ignored unknown evidence type", "type", evidence.Type)
		}
	}
}

// HandleDoubleSign - slashes a validator for signing two blocks at the same
// block height
// https://blog.cosmos.network/consensus-compare-casper-vs-tendermint-6df154ad56ae
func (s *SlasherVCUR) HandleDoubleSign(ctx cosmos.Context, addr crypto.Address, infractionHeight int64, constAccessor constants.ConstantValues) error {
	// check if we're recent enough to slash for this behavior
	maxAge := constAccessor.GetInt64Value(constants.DoubleSignMaxAge)
	if (ctx.BlockHeight() - infractionHeight) > maxAge {
		ctx.Logger().Info("double sign detected but too old to be slashed", "infraction height", fmt.Sprintf("%d", infractionHeight), "address", addr.String())
		return nil
	}
	nas, err := s.keeper.ListActiveValidators(ctx)
	if err != nil {
		return err
	}

	for _, na := range nas {
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeConsPub, na.ValidatorConsPubKey)
		if err != nil {
			return err
		}

		if addr.String() == pk.Address().String() {
			naBond, err := s.keeper.CalcNodeLiquidityBond(ctx, na)
			if err != nil {
				return ErrInternal(err, "fail to get node account bond")
			}
			if naBond.IsZero() {
				return fmt.Errorf("found account to slash for double signing, but did not have any bond to slash: %s", addr)
			}
			// take 5% of the minimum bond, and put it into the reserve
			minBond, err := s.keeper.GetMimir(ctx, constants.MinimumBondInCacao.String())
			if minBond < 0 || err != nil {
				minBond = constAccessor.GetInt64Value(constants.MinimumBondInCacao)
			}
			slashAmount := cosmos.NewUint(uint64(minBond)).MulUint64(5).QuoUint64(100)
			if slashAmount.GT(naBond) {
				slashAmount = naBond
			}

			slashFloat, _ := new(big.Float).SetInt(slashAmount.BigInt()).Float32()
			telemetry.IncrCounterWithLabels(
				[]string{"mayanode", "bond_slash"},
				slashFloat,
				[]metrics.Label{
					telemetry.NewLabel("address", addr.String()),
					telemetry.NewLabel("reason", "double_sign"),
				},
			)

			slashedAmount, _, err := s.SlashNodeAccountLP(ctx, na, slashAmount)

			if slashedAmount.LT(slashAmount) {
				ctx.Logger().Error("slashed less than slash amount", "slashed", slashedAmount, "slash", slashAmount)
			}

			return err
		}
	}

	return fmt.Errorf("could not find node account with validator address: %s", addr)
}

// LackObserving Slash node accounts that didn't observe a single inbound txn
func (s *SlasherVCUR) LackObserving(ctx cosmos.Context, constAccessor constants.ConstantValues) error {
	signingTransPeriod := constAccessor.GetInt64Value(constants.SigningTransactionPeriod)
	height := ctx.BlockHeight()
	if height < signingTransPeriod {
		return nil
	}
	heightToCheck := height - signingTransPeriod
	tx, err := s.keeper.GetTxOut(ctx, heightToCheck)
	if err != nil {
		return fmt.Errorf("fail to get txout for block height(%d): %w", heightToCheck, err)
	}
	// no txout , return
	if tx == nil || tx.IsEmpty() {
		return nil
	}
	for _, item := range tx.TxArray {
		if item.InHash.IsEmpty() {
			continue
		}
		if item.InHash.Equals(common.BlankTxID) {
			continue
		}
		if err := s.slashNotObserving(ctx, item.InHash, constAccessor); err != nil {
			ctx.Logger().Error("fail to slash not observing", "error", err)
		}
	}

	return nil
}

func (s *SlasherVCUR) slashNotObserving(ctx cosmos.Context, txHash common.TxID, constAccessor constants.ConstantValues) error {
	voter, err := s.keeper.GetObservedTxInVoter(ctx, txHash)
	if err != nil {
		return fmt.Errorf("fail to get observe txin voter (%s): %w", txHash.String(), err)
	}

	if len(voter.Txs) == 0 {
		return nil
	}

	nodes, err := s.keeper.ListActiveValidators(ctx)
	if err != nil {
		return fmt.Errorf("unable to get list of active accounts: %w", err)
	}
	if len(voter.Txs) > 0 {
		tx := voter.Tx
		if !tx.IsEmpty() && len(tx.Signers) > 0 {
			height := voter.Height
			if tx.IsFinal() {
				height = voter.FinalisedHeight
			}
			// as long as the node has voted one of the tx , regardless finalised or not , it should not be slashed
			var allSigners []cosmos.AccAddress
			for _, item := range voter.Txs {
				allSigners = append(allSigners, item.GetSigners()...)
			}
			s.checkSignerAndSlash(ctx, nodes, height, allSigners, constAccessor)
		}
	}
	return nil
}

func (s *SlasherVCUR) checkSignerAndSlash(ctx cosmos.Context, nodes NodeAccounts, blockHeight int64, signers []cosmos.AccAddress, constAccessor constants.ConstantValues) {
	for _, na := range nodes {
		// the node is active after the tx finalised
		if na.ActiveBlockHeight > blockHeight {
			continue
		}
		found := false
		for _, addr := range signers {
			if na.NodeAddress.Equals(addr) {
				found = true
				break
			}
		}
		// this na is not found, therefore it should be slashed
		if !found {
			lackOfObservationPenalty := constAccessor.GetInt64Value(constants.LackOfObservationPenalty)
			slashCtx := ctx.WithContext(context.WithValue(ctx.Context(), constants.CtxMetricLabels, []metrics.Label{
				telemetry.NewLabel("reason", "not_observing"),
			}))
			if err := s.keeper.IncNodeAccountSlashPoints(slashCtx, na.NodeAddress, lackOfObservationPenalty); err != nil {
				ctx.Logger().Error("fail to inc slash points", "error", err)
			}
		}
	}
}

// LackSigning slash account that fail to sign tx
func (s *SlasherVCUR) LackSigning(ctx cosmos.Context, mgr Manager) error {
	var resultErr error
	signingTransPeriod := mgr.GetConstants().GetInt64Value(constants.SigningTransactionPeriod)
	if ctx.BlockHeight() < signingTransPeriod {
		return nil
	}
	height := ctx.BlockHeight() - signingTransPeriod
	txs, err := s.keeper.GetTxOut(ctx, height)
	if err != nil {
		return fmt.Errorf("fail to get txout from block height(%d): %w", height, err)
	}
	for i, tx := range txs.TxArray {
		if !common.CurrentChainNetwork.SoftEquals(tx.ToAddress.GetNetwork(mgr.GetVersion(), tx.Chain)) {
			continue // skip this transaction
		}
		if tx.OutHash.IsEmpty() {
			// Slash node account for not sending funds
			vault, err := s.keeper.GetVault(ctx, tx.VaultPubKey)
			if err != nil {
				// in some edge cases, when a txout item had been schedule to be send out by an yggdrasil vault
				// however the node operator decide to quit by sending a leave command, which will result in the vault get removed
				// if that happen , txout item should be scheduled to send out using asgard, thus when if fail to get vault , just
				// log the error, and continue
				ctx.Logger().Error("Unable to get vault", "error", err, "vault pub key", tx.VaultPubKey.String())
			}

			// don't reschedule transactions on frozen vaults. This will cause
			// txns to be trapped in a specific asgard forever, which is the
			// expected result. This is here to protect the network from a
			// round7 attack
			if len(vault.Frozen) > 0 {
				var chains common.Chains
				chains, err = common.NewChains(vault.Frozen)
				if err != nil {
					ctx.Logger().Error("failed to convert chains", "error", err)
				}
				if chains.Has(tx.Coin.Asset.GetChain()) {
					etx := common.Tx{
						ID:        tx.InHash,
						Chain:     tx.Chain,
						ToAddress: tx.ToAddress,
						Coins:     []common.Coin{tx.Coin},
						Gas:       tx.MaxGas,
						Memo:      tx.Memo,
					}
					eve := NewEventSecurity(etx, "skipping reschedule on frozen vault")
					if err = mgr.EventMgr().EmitEvent(ctx, eve); err != nil {
						ctx.Logger().Error("fail to emit security event", "error", err)
					}
					continue // skip this transaction
				}
			}

			// slash if its a yggdrasil vault, and the chain isn't halted
			if vault.IsYggdrasil() && !isChainHalted(ctx, mgr, tx.Chain) {
				var na NodeAccount
				na, err = s.keeper.GetNodeAccountByPubKey(ctx, tx.VaultPubKey)
				if err != nil {
					ctx.Logger().Error("Unable to get node account", "error", err, "vault pub key", tx.VaultPubKey.String())
					continue
				}
				slashPoints := signingTransPeriod * 2

				slashCtx := ctx.WithContext(context.WithValue(ctx.Context(), constants.CtxMetricLabels, []metrics.Label{
					telemetry.NewLabel("reason", "not_signing"),
				}))
				if err = s.keeper.IncNodeAccountSlashPoints(slashCtx, na.NodeAddress, slashPoints); err != nil {
					ctx.Logger().Error("fail to inc slash points", "error", err, "node addr", na.NodeAddress.String())
				}
				if err = mgr.EventMgr().EmitEvent(ctx, NewEventSlashPoint(na.NodeAddress, slashPoints, fmt.Sprintf("fail to sign out tx after %d blocks", signingTransPeriod))); err != nil {
					ctx.Logger().Error("fail to emit slash point event")
				}
				releaseHeight := ctx.BlockHeight() + (signingTransPeriod * 2)
				reason := "fail to send yggdrasil transaction"
				if err = s.keeper.SetNodeAccountJail(ctx, na.NodeAddress, releaseHeight, reason); err != nil {
					ctx.Logger().Error("fail to set node account jail", "node address", na.NodeAddress, "reason", reason, "error", err)
				}
			}

			memo, _ := ParseMemoWithMAYANames(ctx, s.keeper, tx.Memo) // ignore err
			if memo.IsInternal() {
				// there is a different mechanism for rescheduling outbound
				// transactions for migration transactions
				continue
			}
			var voter ObservedTxVoter
			if !memo.IsType(TxRagnarok) {
				voter, err = s.keeper.GetObservedTxInVoter(ctx, tx.InHash)
				if err != nil {
					ctx.Logger().Error("fail to get observed tx voter", "error", err)
					resultErr = fmt.Errorf("failed to get observed tx voter: %w", err)
					continue
				}
			}

			maxOutboundAttempts := fetchConfigInt64(ctx, mgr, constants.MaxOutboundAttempts)
			if maxOutboundAttempts > 0 {
				age := ctx.BlockHeight() - voter.FinalisedHeight
				attempts := age / signingTransPeriod
				if attempts >= maxOutboundAttempts {
					ctx.Logger().Info("txn dropped, too many attempts", "hash", tx.InHash)
					continue
				}
			}

			nas, err := s.keeper.ListActiveValidators(ctx)
			if err != nil {
				ctx.Logger().Error("fail to get all active validators", "error", err)
			}
			if s.needsNewVault(ctx, mgr, len(nas), signingTransPeriod, voter.FinalisedHeight, tx.InHash, tx.VaultPubKey) {
				var active Vaults
				active, err = s.keeper.GetAsgardVaultsByStatus(ctx, ActiveVault)
				if err != nil {
					return fmt.Errorf("fail to get active asgard vaults: %w", err)
				}
				available := active.Has(tx.Coin).SortBy(tx.Coin.Asset)
				if len(available) == 0 {
					// we need to give it somewhere to send from, even if that
					// asgard doesn't have enough funds. This is because if we
					// don't the transaction will just be dropped on the floor,
					// which is bad. Instead it may try to send from an asgard that
					// doesn't have enough funds, fail, and then get rescheduled
					// again later. Maybe by then the network will have enough
					// funds to satisfy.
					// TODO add split logic to send it out from multiple asgards in
					// this edge case.
					ctx.Logger().Error("unable to determine asgard vault to send funds, trying first asgard")
					if len(active) > 0 {
						vault = active[0]
					}
				} else {
					// each time we reschedule a transaction, we take the age of
					// the transaction, and move it to an vault that has less funds
					// than last time. This is here to ensure that if an asgard
					// vault becomes unavailable, the network will reschedule the
					// transaction on a different asgard vault.
					age := ctx.BlockHeight() - voter.FinalisedHeight
					if vault.IsYggdrasil() {
						// since the last attempt was a yggdrasil vault, lets
						// artificially inflate the age to ensure that the first
						// attempt is the largest asgard vault with funds
						age -= signingTransPeriod
						if age < 0 {
							age = 0
						}
					}
					rep := int(age / signingTransPeriod)
					if vault.PubKey.Equals(available[rep%len(available)].PubKey) {
						// looks like the new vault is going to be the same as the
						// old vault, increment rep to ensure a differ asgard is
						// chosen (if there is more than one option)
						rep++
					}
					vault = available[rep%len(available)]
				}
				if !memo.IsType(TxRagnarok) {
					// update original tx action in observed tx
					// check observedTx has done status. Skip if it does already.
					voterTx := voter.GetTx(NodeAccounts{})
					if voterTx.IsDone(len(voter.Actions)) {
						if len(voterTx.OutHashes) > 0 && len(voterTx.GetOutHashes()) > 0 {
							txs.TxArray[i].OutHash = voterTx.GetOutHashes()[0]
						}
						continue
					}

					// update the actions in the voter with the new vault pubkey
					for i, action := range voter.Actions {
						if action.Equals(tx) {
							voter.Actions[i].VaultPubKey = vault.PubKey
						}
					}
					s.keeper.SetObservedTxInVoter(ctx, voter)

				}
				// Save the tx to as a new tx, select Asgard to send it this time.
				tx.VaultPubKey = vault.PubKey
			}

			// update max gas
			maxGas, err := mgr.GasMgr().GetMaxGas(ctx, tx.Chain)
			if err != nil {
				ctx.Logger().Error("fail to get max gas", "error", err)
			} else {
				tx.MaxGas = common.Gas{maxGas}
				// Update MaxGas in ObservedTxVoter action as well
				if err = updateTxOutGas(ctx, s.keeper, tx, common.Gas{maxGas}); err != nil {
					ctx.Logger().Error("Failed to update MaxGas of action in ObservedTxVoter", "hash", tx.InHash, "error", err)
				}
			}
			// Equals checks GasRate so update actions GasRate too (before updating in the queue item)
			// for future updates of MaxGas, which must match for matchActionItem in AddOutTx.
			gasRate := int64(mgr.GasMgr().GetGasRate(ctx, tx.Chain).Uint64())
			if err = updateTxOutGasRate(ctx, s.keeper, tx, gasRate); err != nil {
				ctx.Logger().Error("Failed to update GasRate of action in ObservedTxVoter", "hash", tx.InHash, "error", err)
			}
			tx.GasRate = gasRate

			// if a pool with the asset name doesn't exist, skip rescheduling
			if !tx.Coin.Asset.IsBase() && !s.keeper.PoolExist(ctx, tx.Coin.Asset) {
				ctx.Logger().Error("fail to add outbound tx", "error", "coin is not rune and does not have an associated pool")
				continue
			}

			// round up to next coalesce block
			rescheduleHeight := ctx.BlockHeight()
			rescheduleCoalesceBlocks := mgr.GetConfigInt64(ctx, constants.RescheduleCoalesceBlocks)
			if rescheduleCoalesceBlocks > 1 {
				rescheduleHeight += rescheduleCoalesceBlocks - (rescheduleHeight % rescheduleCoalesceBlocks)
			}

			err = mgr.TxOutStore().UnSafeAddTxOutItem(ctx, mgr, tx, rescheduleHeight)
			if err != nil {
				ctx.Logger().Error("fail to add outbound tx", "error", err)
				resultErr = fmt.Errorf("failed to add outbound tx: %w", err)
				continue
			}
			// because the txout item has been rescheduled, thus mark the replaced tx out item as already send out, even it is not
			// in this way bifrost will not send it out again cause node to be slashed
			txs.TxArray[i].OutHash = common.BlankTxID
		}
	}
	if !txs.IsEmpty() {
		if err := s.keeper.SetTxOut(ctx, txs); err != nil {
			return fmt.Errorf("fail to save tx out : %w", err)
		}
	}

	return resultErr
}

// IncSlashPoints will increase the given account's slash points
func (s *SlasherVCUR) IncSlashPoints(ctx cosmos.Context, point int64, addresses ...cosmos.AccAddress) {
	for _, addr := range addresses {
		if err := s.keeper.IncNodeAccountSlashPoints(ctx, addr, point); err != nil {
			ctx.Logger().Error("fail to increase node account slash point", "error", err, "address", addr.String())
		}
	}
}

// DecSlashPoints will decrease the given account's slash points
func (s *SlasherVCUR) DecSlashPoints(ctx cosmos.Context, point int64, addresses ...cosmos.AccAddress) {
	for _, addr := range addresses {
		if err := s.keeper.DecNodeAccountSlashPoints(ctx, addr, point); err != nil {
			ctx.Logger().Error("fail to decrease node account slash point", "error", err, "address", addr.String())
		}
	}
}

// SlashVaultToLP slashes a vault the membership's LPUnits that are bonded.
func (s *SlasherVCUR) SlashVaultToLP(ctx cosmos.Context, vaultPK common.PubKey, coins common.Coins, mgr Manager, subsidize bool) error {
	if coins.IsEmpty() {
		return nil
	}

	vault, err := s.keeper.GetVault(ctx, vaultPK)
	if err != nil {
		return fmt.Errorf("fail to get slash vault (pubkey %s), %w", vaultPK, err)
	}

	// Get total bond of membership of the vault.
	membership := vault.GetMembership()
	totalBond := cosmos.ZeroUint()
	for _, member := range membership {
		na, err := s.keeper.GetNodeAccountByPubKey(ctx, member)
		if err != nil {
			ctx.Logger().Error("fail to get node account bond", "pk", member, "error", err)
			continue
		}
		naBond, err := s.keeper.CalcNodeLiquidityBond(ctx, na)
		if err != nil {
			ctx.Logger().Error("fail to get node account bond", "pk", member, "error", err)
			continue
		}

		totalBond = totalBond.Add(naBond)
	}

	totalBaseToSlash := cosmos.ZeroUint()
	totalBaseStolen := cosmos.ZeroUint()
	for _, coin := range coins {
		if coin.IsEmpty() {
			continue
		}
		pool, err := s.keeper.GetPool(ctx, coin.Asset)
		if err != nil {
			ctx.Logger().Error("fail to get pool for slash", "asset", coin.Asset, "error", err)
			continue
		}
		// BASEChain doesn't have a pool for the asset
		if pool.IsEmpty() {
			ctx.Logger().Error("cannot slash for an empty pool", "asset", coin.Asset)
			continue
		}

		stolenAssetValue := coin.Amount
		vaultAmount := vault.GetCoin(coin.Asset).Amount
		if stolenAssetValue.GT(vaultAmount) {
			stolenAssetValue = vaultAmount
		}
		if stolenAssetValue.GT(pool.BalanceAsset) {
			stolenAssetValue = pool.BalanceAsset
		}

		// stolenBaseValue is the value in RUNE of the missing funds
		stolenBaseValue := pool.AssetValueInRune(stolenAssetValue)
		totalBaseStolen = totalBaseStolen.Add(stolenBaseValue)

		if stolenBaseValue.IsZero() {
			continue
		}

		penaltyPts := fetchConfigInt64(ctx, mgr, constants.SlashPenalty)
		// total slash amount is penaltyPts the RUNE value of the missing funds
		totalBaseToSlash = totalBaseToSlash.Add(common.GetUncappedShare(cosmos.NewUint(uint64(penaltyPts)), cosmos.NewUint(100_00), stolenBaseValue))
		pauseOnSlashThreshold := fetchConfigInt64(ctx, mgr, constants.PauseOnSlashThreshold)
		if pauseOnSlashThreshold > 0 && totalBaseToSlash.GTE(cosmos.NewUint(uint64(pauseOnSlashThreshold))) {
			// set mimirs to pause the chain and ygg funding
			s.keeper.SetMimir(ctx, mimirStopFundYggdrasil, ctx.BlockHeight())
			mimirEvent := NewEventSetMimir(strings.ToUpper(mimirStopFundYggdrasil), strconv.FormatInt(ctx.BlockHeight(), 10))
			if err := mgr.EventMgr().EmitEvent(ctx, mimirEvent); err != nil {
				ctx.Logger().Error("fail to emit set_mimir event", "error", err)
			}

			key := fmt.Sprintf("Halt%sChain", coin.Asset.Chain)
			s.keeper.SetMimir(ctx, key, ctx.BlockHeight())
			mimirEvent = NewEventSetMimir(strings.ToUpper(key), strconv.FormatInt(ctx.BlockHeight(), 10))
			if err := mgr.EventMgr().EmitEvent(ctx, mimirEvent); err != nil {
				ctx.Logger().Error("fail to emit set_mimir event", "error", err)
			}
		}
	}

	totalBaseSlashed := cosmos.ZeroUint()
	for _, member := range membership {
		na, err := s.keeper.GetNodeAccountByPubKey(ctx, member)
		if err != nil {
			ctx.Logger().Error("fail to get node account for slash", "pk", member, "error", err)
			continue
		}
		naBond, err := mgr.Keeper().CalcNodeLiquidityBond(ctx, na)
		if err != nil {
			ctx.Logger().Error("fail to get node account bond", "error", err)
			continue
		}
		if naBond.IsZero() {
			ctx.Logger().Info("validator's bond is zero, can't be slashed", "node address", na.NodeAddress.String())
			continue
		}
		slashAmountRune := common.GetSafeShare(naBond, totalBond, totalBaseToSlash)
		if slashAmountRune.GT(naBond) {
			ctx.Logger().Info("slash amount is larger than bond", "slash amount", slashAmountRune, "bond", naBond)
			slashAmountRune = naBond
		}
		// need to count total slashed bond again , because the node might not have enough bond left
		ctx.Logger().Info("slash node account", "node address", na.NodeAddress.String(), "amount", slashAmountRune.String(), "total slash amount", totalBaseToSlash)
		naBond = common.SafeSub(naBond, slashAmountRune)

		// slash the node account
		slashedAmount, _, err := s.SlashNodeAccountLP(ctx, na, slashAmountRune)
		if err != nil {
			ctx.Logger().Error("fail to slash node account", "error", err)
			continue
		}

		totalBaseSlashed = totalBaseSlashed.Add(slashedAmount)
		for _, coin := range coins {
			metricLabels, _ := ctx.Context().Value(constants.CtxMetricLabels).([]metrics.Label)
			slashAmountRuneFloat, _ := new(big.Float).SetInt(slashAmountRune.BigInt()).Float32()
			telemetry.IncrCounterWithLabels(
				[]string{"mayanode", "bond_slash"},
				slashAmountRuneFloat,
				append(
					metricLabels,
					telemetry.NewLabel("address", na.NodeAddress.String()),
					telemetry.NewLabel("coin_symbol", coin.Asset.Symbol.String()),
					telemetry.NewLabel("coin_chain", string(coin.Asset.Chain)),
					telemetry.NewLabel("vault_type", vault.Type.String()),
					telemetry.NewLabel("vault_status", vault.Status.String()),
				),
			)
		}

		// Ban the node account. Ensure we don't ban more than 1/3rd of any
		// given active or retiring vault
		if vault.IsYggdrasil() {
			// TODO: temporally disabling banning for the theft of funds. This
			// is to give the code time to prove itself reliable before the it
			// starts booting nodes out of the system
			toBan := false // TODO flip this to true
			if naBond.IsZero() {
				toBan = true
			}
			for _, vaultPk := range na.GetSignerMembership() {
				vault, err = s.keeper.GetVault(ctx, vaultPk)
				if err != nil {
					ctx.Logger().Error("fail to get vault", "error", err)
					continue
				}
				if !(vault.Status == ActiveVault || vault.Status == RetiringVault) {
					continue
				}
				activeMembers := 0
				for _, pk := range vault.GetMembership() {
					member, _ := s.keeper.GetNodeAccountByPubKey(ctx, pk)
					if member.Status == NodeActive {
						activeMembers++
					}
				}
				if !HasSuperMajority(activeMembers, len(vault.GetMembership())) {
					toBan = false
					break
				}
			}
			if toBan {
				na.ForcedToLeave = true
				na.LeaveScore = 1 // Set Leave Score to 1, which means the nodes is bad
			}
		}
	}

	if subsidize {
		return subsidizePoolsWithSlashBond(ctx, coins, vault, totalBaseStolen, totalBaseSlashed, mgr)
	}
	return nil
}

// SlashNodeAccountLP slashes a node account based its LP units.
// We take the percentage that the slash represents from the total bond
// of the node and slash that percentage of the LP units to each of the
func (s *SlasherVCUR) SlashNodeAccountLP(ctx cosmos.Context, na NodeAccount, slash cosmos.Uint) (cosmos.Uint, []types.PoolAmt, error) {
	if slash.IsZero() {
		return cosmos.ZeroUint(), nil, nil
	}

	for _, genesis := range GenesisNodes {
		add, err := common.NewAddress(genesis)
		if err != nil {
			ctx.Logger().Error("fail to process genesis node address", "error", err)
			continue
		}
		if na.BondAddress.Equals(add) {
			return cosmos.ZeroUint(), nil, nil
		}
	}

	var slashedAmountsPerPool []types.PoolAmt
	totalSlashed := cosmos.ZeroUint()
	naBond, err := s.keeper.CalcNodeLiquidityBond(ctx, na)
	if err != nil {
		return totalSlashed, slashedAmountsPerPool, ErrInternal(err, "fail to get node account bond")
	}

	polAddress, err := s.keeper.GetModuleAddress(ReserveName)
	if err != nil {
		return totalSlashed, slashedAmountsPerPool, err
	}

	if naBond.IsZero() {
		ctx.Logger().Info("validator's bond is zero, can't be slashed", "node address", na.NodeAddress.String())
		return totalSlashed, slashedAmountsPerPool, errors.New("validator's bond is zero, can't be slashed")
	}

	if slash.GT(naBond) {
		ctx.Logger().Info("slash amount is larger than bond", "slash amount", slash, "bond", naBond)
		slash = naBond
	}

	bp, err := s.keeper.GetBondProviders(ctx, na.NodeAddress)
	if err != nil {
		return totalSlashed, slashedAmountsPerPool, ErrInternal(err, "fail to get node bond providers")
	}

	// It will be at least the length of providers
	liquidityPools := GetLiquidityPools(s.keeper.GetVersion())
	for _, b := range bp.Providers {
		lps, err := s.keeper.GetLiquidityProviderByAssets(ctx, liquidityPools, common.Address(b.BondAddress.String()))
		if err != nil {
			ctx.Logger().Error("fail to get lps for bond provider", "error", err)
			continue
		}
		// Slash corresponding lpunits proportionally to LP.
		for _, lp := range lps {
			pool, err := s.keeper.GetPool(ctx, lp.Asset)
			if err != nil {
				ctx.Logger().Error("fail to get pool", "error", err)
				continue
			}
			if pool.IsAvailable() {
				polLP, err := s.keeper.GetLiquidityProvider(ctx, pool.Asset, polAddress)
				if err != nil {
					ctx.Logger().Error("fail to get pool liquidity provider", "error", err)
					continue
				}

				bondedUnits := lp.GetUnitsBondedToNode(na.NodeAddress)
				// Sanity check
				if bondedUnits.IsZero() {
					continue
				}

				// Remove the slash percentage from the LP units
				slashLPUnits := common.GetSafeShare(slash, naBond, bondedUnits)
				slashedAmount := common.GetSafeShare(slash, naBond, naBond)
				slashedAmountsPerPool = append(slashedAmountsPerPool, types.PoolAmt{
					Asset:  pool.Asset,
					Amount: slashedAmount.BigInt().Int64(),
				})
				totalSlashed = totalSlashed.Add(slashedAmount)

				// Take away corresponding lp units due to slash
				// and add them to the POL
				lp.Unbond(na.NodeAddress, slashLPUnits)
				lp.Units = lp.Units.Sub(slashLPUnits)
				polLP.Units = polLP.Units.Add(slashLPUnits)

				slashLiquidityEvt := types.NewEventSlashLiquidity(na.NodeAddress, pool.Asset, lp.CacaoAddress, slashLPUnits)

				if err := s.eventMgr.EmitEvent(ctx, slashLiquidityEvt); err != nil {
					ctx.Logger().Error("fail to emit slash liquidity event", "error", err)
				}

				s.keeper.SetLiquidityProviders(ctx, LiquidityProviders{lp, polLP})
			}
		}
	}

	return totalSlashed, slashedAmountsPerPool, nil
}

func (s *SlasherVCUR) needsNewVault(ctx cosmos.Context, mgr Manager, nas int, signingTransPeriod, startHeight int64, inhash common.TxID, pk common.PubKey) bool {
	outhashes := mgr.Keeper().GetObservedLink(ctx, inhash)
	if len(outhashes) == 0 {
		return true
	}

	for _, hash := range outhashes {
		voter, err := mgr.Keeper().GetObservedTxOutVoter(ctx, hash)
		if err != nil {
			ctx.Logger().Error("fail to get txout voter", "hash", hash, "error", err)
			continue
		}
		// in the event there are multiple outbounds for a given inhash, we
		// focus on the matching pubkey
		signers := make([]string, 0)
		for _, tx1 := range voter.Txs {
			if tx1.ObservedPubKey.Equals(pk) {
				for _, tx := range voter.Txs {
					if !tx.Tx.ID.Equals(hash) {
						continue
					}
					if len(signers) < len(tx.Signers) {
						signers = tx.Signers
					}
				}
			}
		}
		if len(signers) > 0 {
			if nas > 0 && HasMinority(len(signers), nas) {
				return false
			}
			maxHeight := startHeight + ((int64(len(signers)) + 1) * signingTransPeriod)
			return maxHeight < ctx.BlockHeight()
		}

	}

	return true
}
