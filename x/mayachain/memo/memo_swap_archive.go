package mayachain

import (
	"fmt"
	"strconv"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

func ParseSwapMemoV1(ctx cosmos.Context, keeper keeper.Keeper, asset common.Asset, parts []string) (SwapMemo, error) {
	var err error
	var order types.OrderType
	if len(parts) < 2 {
		return SwapMemo{}, fmt.Errorf("not enough parameters")
	}
	// DESTADDR can be empty , if it is empty , it will swap to the sender address
	destination := common.NoAddress
	affAddr := common.NoAddress
	affPts := cosmos.ZeroUint()
	if len(parts) > 2 {
		if len(parts[2]) > 0 {
			if keeper == nil {
				destination, err = common.NewAddress(parts[2])
			} else {
				destination, err = FetchAddress(ctx, keeper, parts[2], asset.Chain)
			}
			if err != nil {
				return SwapMemo{}, err
			}
		}
	}
	// price limit can be empty , when it is empty , there is no price protection
	var amount cosmos.Uint
	slip := cosmos.ZeroUint()
	if len(parts) > 3 && len(parts[3]) > 0 {
		amount, err = cosmos.ParseUint(parts[3])
		if err != nil {
			return SwapMemo{}, fmt.Errorf("swap price limit:%s is invalid", parts[3])
		}
		slip = amount
	}

	if len(parts) > 5 && len(parts[4]) > 0 && len(parts[5]) > 0 {
		if keeper == nil {
			affAddr, err = common.NewAddress(parts[4])
		} else {
			affAddr, err = FetchAddress(ctx, keeper, parts[4], common.THORChain)
		}
		if err != nil {
			return SwapMemo{}, err
		}
		var pts uint64
		pts, err = strconv.ParseUint(parts[5], 10, 64)
		if err != nil {
			return SwapMemo{}, err
		}
		affPts = cosmos.NewUint(pts)
	}

	return NewSwapMemo(asset, destination, slip, affAddr, affPts, "", "", cosmos.ZeroUint(), order, 0, 0), nil
}

func ParseSwapMemoV92(ctx cosmos.Context, keeper keeper.Keeper, asset common.Asset, parts []string) (SwapMemo, error) {
	var err error
	var order types.OrderType
	dexAgg := ""
	dexTargetAddress := ""
	dexTargetLimit := cosmos.ZeroUint()
	if len(parts) < 2 {
		return SwapMemo{}, fmt.Errorf("not enough parameters")
	}
	// DESTADDR can be empty , if it is empty , it will swap to the sender address
	destination := common.NoAddress
	affAddr := common.NoAddress
	affPts := cosmos.ZeroUint()
	if len(parts) > 2 {
		if len(parts[2]) > 0 {
			if keeper == nil {
				destination, err = common.NewAddress(parts[2])
			} else {
				destination, err = FetchAddress(ctx, keeper, parts[2], asset.Chain)
			}
			if err != nil {
				return SwapMemo{}, err
			}
		}
	}
	// price limit can be empty , when it is empty , there is no price protection
	var amount cosmos.Uint
	slip := cosmos.ZeroUint()
	if len(parts) > 3 && len(parts[3]) > 0 {
		amount, err = cosmos.ParseUint(parts[3])
		if err != nil {
			return SwapMemo{}, fmt.Errorf("swap price limit:%s is invalid", parts[3])
		}
		slip = amount
	}

	if len(parts) > 5 && len(parts[4]) > 0 && len(parts[5]) > 0 {
		if keeper == nil {
			affAddr, err = common.NewAddress(parts[4])
		} else {
			affAddr, err = FetchAddress(ctx, keeper, parts[4], common.BASEChain)
		}
		if err != nil {
			return SwapMemo{}, err
		}
		var pts uint64
		pts, err = strconv.ParseUint(parts[5], 10, 64)
		if err != nil {
			return SwapMemo{}, err
		}
		affPts = cosmos.NewUint(pts)
	}

	if len(parts) > 6 && len(parts[6]) > 0 {
		dexAgg = parts[6]
	}

	if len(parts) > 7 && len(parts[7]) > 0 {
		dexTargetAddress = parts[7]
	}

	if len(parts) > 8 && len(parts[8]) > 0 {
		dexTargetLimit, err = cosmos.ParseUint(parts[8])
		if err != nil {
			ctx.Logger().Error("invalid dex target limit, ignore it", "limit", parts[8])
			dexTargetLimit = cosmos.ZeroUint()
		}
	}

	return NewSwapMemo(asset, destination, slip, affAddr, affPts, dexAgg, dexTargetAddress, dexTargetLimit, order, 0, 0), nil
}
