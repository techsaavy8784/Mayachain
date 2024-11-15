package mayachain

import (
	"strconv"
	"strings"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
)

// SwapQueueVCUR is going to manage the swaps queue
type SwapQueueVCUR struct {
	k keeper.Keeper
}

// newSwapQueueVCUR create a new vault manager
func newSwapQueueVCUR(k keeper.Keeper) *SwapQueueVCUR {
	return &SwapQueueVCUR{k: k}
}

// FetchQueue - grabs all swap queue items from the kvstore and returns them
func (vm *SwapQueueVCUR) FetchQueue(ctx cosmos.Context, mgr Manager) (swapItems, error) { // nolint
	items := make(swapItems, 0)
	iterator := vm.k.GetSwapQueueIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var msg MsgSwap
		if err := vm.k.Cdc().Unmarshal(iterator.Value(), &msg); err != nil {
			ctx.Logger().Error("fail to fetch swap msg from queue", "error", err)
			continue
		}

		ss := strings.Split(string(iterator.Key()), "-")
		i, err := strconv.Atoi(ss[len(ss)-1])
		if err != nil {
			ctx.Logger().Error("fail to parse swap queue msg index", "key", iterator.Key(), "error", err)
			continue
		}

		// exclude streaming swaps when its not "their time". Always want to
		// allow the first sub-swap immediately (ie no LastHeight yet)
		if msg.IsStreaming() {
			pausedStreaming := fetchConfigInt64(ctx, mgr, constants.StreamingSwapPause)
			if pausedStreaming > 0 {
				continue
			}
			swp := msg.GetStreamingSwap()
			if vm.k.StreamingSwapExists(ctx, msg.Tx.ID) {
				var err error
				swp, err = vm.k.GetStreamingSwap(ctx, msg.Tx.ID)
				if err != nil {
					ctx.Logger().Error("fail to fetch streaming swap", "error", err)
					continue
				}
			}
			if swp.LastHeight > 0 { // if we don't have a height, do first swap attempt now
				if swp.LastHeight >= ctx.BlockHeight() {
					// last swap must be in the past
					continue // skip
				}
				if (ctx.BlockHeight()-swp.LastHeight)%int64(swp.Interval) != 0 {
					continue // skip
				}
				if isTradingHalt(ctx, &msg, mgr) {
					// if trading/chain is halted, skip
					continue // skip
				}
			}
		}

		items = append(items, swapItem{
			msg:   msg,
			index: i,
			fee:   cosmos.ZeroUint(),
			slip:  cosmos.ZeroUint(),
		})
	}

	return items, nil
}

// EndBlock trigger the real swap to be processed
func (vm *SwapQueueVCUR) EndBlock(ctx cosmos.Context, mgr Manager) error {
	handler := NewInternalHandler(mgr)

	minSwapsPerBlock, err := vm.k.GetMimir(ctx, constants.MinSwapsPerBlock.String())
	if minSwapsPerBlock < 0 || err != nil {
		minSwapsPerBlock = mgr.GetConstants().GetInt64Value(constants.MinSwapsPerBlock)
	}
	maxSwapsPerBlock, err := vm.k.GetMimir(ctx, constants.MaxSwapsPerBlock.String())
	if maxSwapsPerBlock < 0 || err != nil {
		maxSwapsPerBlock = mgr.GetConstants().GetInt64Value(constants.MaxSwapsPerBlock)
	}
	synthVirtualDepthMult, err := vm.k.GetMimir(ctx, constants.VirtualMultSynthsBasisPoints.String())
	if synthVirtualDepthMult < 1 || err != nil {
		synthVirtualDepthMult = mgr.GetConstants().GetInt64Value(constants.VirtualMultSynthsBasisPoints)
	}

	swaps, err := vm.FetchQueue(ctx, mgr)
	if err != nil {
		ctx.Logger().Error("fail to fetch swap queue from store", "error", err)
		return err
	}
	swaps, err = vm.scoreMsgs(ctx, swaps, synthVirtualDepthMult, mgr)
	if err != nil {
		ctx.Logger().Error("fail to fetch swap items", "error", err)
		// continue, don't exit, just do them out of order (instead of not at all)
	}
	swaps = swaps.Sort()

	for i := int64(0); i < vm.getTodoNum(int64(len(swaps)), minSwapsPerBlock, maxSwapsPerBlock); i++ {
		pick := swaps[i]
		// grab swp BEFORE a streaming swap modified the msg.Tx.Coins[0].Amount
		// value. This is used later to refund the correct amount
		swp := pick.msg.GetStreamingSwap()

		triggerRefund := false
		_, handleErr := handler(ctx, &pick.msg)
		if handleErr != nil {
			ctx.Logger().Error("fail to swap", "msg", pick.msg.Tx.String(), "error", handleErr)

			var refundErr error
			triggerRefund = !pick.msg.IsStreaming()

			if pick.msg.IsStreaming() {
				if vm.k.StreamingSwapExists(ctx, pick.msg.Tx.ID) {
					var getErr error
					swp, getErr = vm.k.GetStreamingSwap(ctx, pick.msg.Tx.ID)
					if getErr != nil {
						ctx.Logger().Error("fail to fetch streaming swap", "error", getErr)
						return getErr
					}
				}

				// if we haven't made any swaps yet, its safe to do a regular
				// refund. Otherwise allow later code to do partial refunds
				triggerRefund = swp.In.IsZero() && swp.Out.IsZero()
				if triggerRefund {
					// revert the tx amount to the be original deposit amount
					pick.msg.Tx.Coins[0].Amount = swp.Deposit
					vm.k.RemoveStreamingSwap(ctx, pick.msg.Tx.ID)
					vm.k.RemoveSwapQueueItem(ctx, pick.msg.Tx.ID, pick.index)
				}
			}

			if triggerRefund {
				// Get the full ObservedTx from the TxID, for the vault ObservedPubKey to first try to refund from.
				voter, voterErr := mgr.Keeper().GetObservedTxInVoter(ctx, pick.msg.Tx.ID)
				if voterErr == nil && !voter.Tx.IsEmpty() {
					refundErr = refundTx(ctx, ObservedTx{Tx: pick.msg.Tx, ObservedPubKey: voter.Tx.ObservedPubKey}, mgr, CodeSwapFail, handleErr.Error(), "")
				} else {
					// If the full ObservedTx could not be retrieved, proceed with just the MsgSwap's Tx (no ObservedPubKey).
					ctx.Logger().Error("fail to get non-empty observed tx", "error", voterErr)
					refundErr = refundTx(ctx, ObservedTx{Tx: pick.msg.Tx}, mgr, CodeSwapFail, handleErr.Error(), "")
				}

				if refundErr != nil {
					ctx.Logger().Error("fail to refund swap", "error", refundErr)
				}
			}
		}

		if pick.msg.IsStreaming() {
			swp, err := vm.k.GetStreamingSwap(ctx, pick.msg.Tx.ID)
			if err != nil {
				ctx.Logger().Error("fail to fetch streaming swap", "error", err)
				return err
			}
			swp.Count += 1
			if handleErr != nil {
				swp.FailedSwaps = append(swp.FailedSwaps, swp.Count)
				swp.FailedSwapReasons = append(swp.FailedSwapReasons, handleErr.Error())
			}
			swp.LastHeight = ctx.BlockHeight()
			if !triggerRefund {
				mgr.Keeper().SetStreamingSwap(ctx, swp)
			}
			if swp.Valid() == nil && swp.IsDone() {
				vm.k.RemoveSwapQueueItem(ctx, pick.msg.Tx.ID, pick.index)
				vm.k.RemoveStreamingSwap(ctx, pick.msg.Tx.ID)

				memo, err := ParseMemoWithMAYANames(ctx, vm.k, pick.msg.Tx.Memo)
				if err != nil {
					return err
				}
				isSaversAdd := memo.IsType(TxAdd)

				// If this is a savers add skip scheduling outbound
				if !swp.Out.IsZero() && !isSaversAdd {
					dexAgg := ""
					if len(pick.msg.Aggregator) > 0 {
						dexAgg, err = FetchDexAggregator(
							mgr.GetVersion(),
							pick.msg.TargetAsset.GetChain(),
							pick.msg.Aggregator,
						)
						if err != nil {
							return err
						}
					}
					dexAggTargetAsset := pick.msg.AggregatorTargetAddress

					toi := TxOutItem{
						Chain:                 pick.msg.TargetAsset.GetChain(),
						InHash:                pick.msg.Tx.ID,
						ToAddress:             pick.msg.Destination,
						Coin:                  common.NewCoin(pick.msg.TargetAsset, swp.Out),
						Aggregator:            dexAgg,
						AggregatorTargetAsset: dexAggTargetAsset,
						AggregatorTargetLimit: pick.msg.AggregatorTargetLimit,
					}

					if _, err := mgr.TxOutStore().TryAddTxOutItem(ctx, mgr, toi, cosmos.ZeroUint()); err != nil {
						ctx.Logger().Error("fail streaming swap outbound", "error", err)
						unrefundableCoinCleanup(ctx, mgr, toi, "failed_outbound")
					}
				}

				if swp.Deposit.GT(swp.In) {
					remainder := common.SafeSub(swp.Deposit, swp.In)
					source := pick.msg.Tx.Coins[0].Asset
					refundCoin := common.NewCoin(source, remainder)
					refundCoinTx := pick.msg.Tx
					refundCoinTx.Coins = common.NewCoins(refundCoin)
					// As this is a streaming swap's partial refund, the vault context may have changed, so do vault selection.
					if refundErr := refundTx(ctx, ObservedTx{Tx: refundCoinTx}, mgr, CodeSwapFail, handleErr.Error(), ""); refundErr != nil {
						ctx.Logger().Error("fail to partial-refund swap", "error", refundErr)
					}
				}

				evt := NewEventStreamingSwap(pick.msg.Tx.Coins[0].Asset, pick.msg.TargetAsset, swp)
				if err := mgr.EventMgr().EmitEvent(ctx, evt); err != nil {
					ctx.Logger().Error("fail to emit streaming swap event", "error", err)
				}
			}
		} else {
			vm.k.RemoveSwapQueueItem(ctx, pick.msg.Tx.ID, pick.index)
		}
	}
	return nil
}

// getTodoNum - determine how many swaps to do.
func (vm *SwapQueueVCUR) getTodoNum(queueLen, minSwapsPerBlock, maxSwapsPerBlock int64) int64 {
	// Do half the length of the queue. Unless...
	//	1. The queue length is greater than maxSwapsPerBlock
	//  2. The queue length is less than minSwapsPerBlock
	todo := queueLen / 2
	if minSwapsPerBlock >= queueLen {
		todo = queueLen
	}
	if maxSwapsPerBlock < todo {
		todo = maxSwapsPerBlock
	}
	return todo
}

// scoreMsgs - this takes a list of MsgSwap, and converts them to a scored
// swapItem list
func (vm *SwapQueueVCUR) scoreMsgs(ctx cosmos.Context, items swapItems, synthVirtualDepthMult int64, mgr Manager) (swapItems, error) {
	pools := make(map[common.Asset]Pool)

	for i, item := range items {
		// the asset customer send
		sourceAsset := item.msg.Tx.Coins[0].Asset
		// the asset customer want
		targetAsset := item.msg.TargetAsset

		assets := common.Assets{sourceAsset, targetAsset}
		for _, a := range assets {
			if a.IsBase() {
				continue
			}

			if _, ok := pools[a]; !ok {
				var err error
				pools[a], err = vm.k.GetPool(ctx, a.GetLayer1Asset())
				if err != nil {
					ctx.Logger().Error("fail to get pool", "pool", a, "error", err)
					continue
				}
			}
		}

		nonBaseAsset := sourceAsset
		if nonBaseAsset.IsBase() {
			nonBaseAsset = targetAsset
		}
		pool := pools[nonBaseAsset]
		if pool.IsEmpty() || pool.BalanceCacao.IsZero() || pool.BalanceAsset.IsZero() {
			continue
		}
		// synths may be redeemed on unavailable pools, score them
		if !pool.IsAvailable() && !sourceAsset.IsSyntheticAsset() {
			continue
		}
		virtualDepthMult := int64(10_000)
		if nonBaseAsset.IsSyntheticAsset() {
			virtualDepthMult = synthVirtualDepthMult
		}
		vm.getLiquidityFeeAndSlip(ctx, pool, item.msg.Tx.Coins[0], &items[i], virtualDepthMult, mgr)

		if sourceAsset.IsBase() || targetAsset.IsBase() {
			// single swap , stop here
			continue
		}
		// double swap , thus need to convert source coin to RUNE and calculate fee and slip again
		runeCoin := common.NewCoin(common.BaseAsset(), pool.AssetValueInRune(item.msg.Tx.Coins[0].Amount))
		nonBaseAsset = targetAsset
		pool = pools[nonBaseAsset]
		if pool.IsEmpty() || !pool.IsAvailable() || pool.BalanceCacao.IsZero() || pool.BalanceAsset.IsZero() {
			continue
		}
		virtualDepthMult = int64(10_000)
		if targetAsset.IsSyntheticAsset() {
			virtualDepthMult = synthVirtualDepthMult
		}
		vm.getLiquidityFeeAndSlip(ctx, pool, runeCoin, &items[i], virtualDepthMult, mgr)
	}

	return items, nil
}

// getLiquidityFeeAndSlip calculate liquidity fee and slip, fee is in RUNE
func (vm *SwapQueueVCUR) getLiquidityFeeAndSlip(ctx cosmos.Context, pool Pool, sourceCoin common.Coin, item *swapItem, virtualDepthMult int64, mgr Manager) {
	// Get our X, x, Y values
	var X, x, Y cosmos.Uint
	x = sourceCoin.Amount
	if sourceCoin.Asset.IsBase() {
		X = pool.BalanceCacao
		Y = pool.BalanceAsset
	} else {
		Y = pool.BalanceCacao
		X = pool.BalanceAsset
	}

	X = common.GetUncappedShare(cosmos.NewUint(uint64(virtualDepthMult)), cosmos.NewUint(10_000), X)
	Y = common.GetUncappedShare(cosmos.NewUint(uint64(virtualDepthMult)), cosmos.NewUint(10_000), Y)

	swapper, err := GetSwapper(vm.k.GetVersion())
	if err != nil {
		ctx.Logger().Error("fail to fetch swapper", "error", err)
		swapper = newSwapperV95()
	}
	fee := swapper.CalcLiquidityFee(X, x, Y)
	if sourceCoin.Asset.IsBase() {
		fee = pool.AssetValueInRune(fee)
	}

	slipFeeAddedBasisPoints := getSlipFeeAddedBasisPoints(ctx, mgr)

	slip := swapper.CalcSwapSlip(X, x, cosmos.NewUint(slipFeeAddedBasisPoints))
	item.fee = item.fee.Add(fee)
	item.slip = item.slip.Add(slip)
}
