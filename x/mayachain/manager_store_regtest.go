//go:build regtest
// +build regtest

package mayachain

import (
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
)

func migrateStoreV86(ctx cosmos.Context, mgr *Mgrs)    {}
func migrateStoreV88(ctx cosmos.Context, mgr Manager)  {}
func migrateStoreV90(ctx cosmos.Context, mgr Manager)  {}
func migrateStoreV96(ctx cosmos.Context, mgr Manager)  {}
func migrateStoreV102(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV104(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV105(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV106(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV107(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV108(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV109(ctx cosmos.Context, mgr Manager) {}
func migrateStoreV110(ctx cosmos.Context, mgr Manager) {}

func migrateStoreV111(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v111", "error", err)
		}
	}()

	// For any in-progress streaming swaps to non-RUNE Native coins,
	// mint the current Out amount to the Pool Module.
	var coinsToMint common.Coins

	iterator := mgr.Keeper().GetSwapQueueIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var msg MsgSwap
		if err := mgr.Keeper().Cdc().Unmarshal(iterator.Value(), &msg); err != nil {
			ctx.Logger().Error("fail to fetch swap msg from queue", "error", err)
			continue
		}

		if !msg.IsStreaming() || !msg.TargetAsset.IsNative() || msg.TargetAsset.IsBase() {
			continue
		}

		swp, err := mgr.Keeper().GetStreamingSwap(ctx, msg.Tx.ID)
		if err != nil {
			ctx.Logger().Error("fail to fetch streaming swap", "error", err)
			continue
		}

		if !swp.Out.IsZero() {
			mintCoin := common.NewCoin(msg.TargetAsset, swp.Out)
			coinsToMint = coinsToMint.Add(mintCoin)
		}
	}

	// The minted coins are for in-progress swaps, so keeping the "swap" in the event field and logs.
	var coinsToTransfer common.Coins
	for _, mintCoin := range coinsToMint {
		if err := mgr.Keeper().MintToModule(ctx, ModuleName, mintCoin); err != nil {
			ctx.Logger().Error("fail to mint coins during swap", "error", err)
		} else {
			// MintBurn event is not currently implemented, will ignore

			// mintEvt := NewEventMintBurn(MintSupplyType, mintCoin.Asset.Native(), mintCoin.Amount, "swap")
			// if err := mgr.EventMgr().EmitEvent(ctx, mintEvt); err != nil {
			// 	ctx.Logger().Error("fail to emit mint event", "error", err)
			// }
			coinsToTransfer = coinsToTransfer.Add(mintCoin)
		}
	}

	if err := mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, coinsToTransfer); err != nil {
		ctx.Logger().Error("fail to move coins during swap", "error", err)
	}
}
