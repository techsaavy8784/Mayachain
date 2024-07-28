package mayachain

import (
	"strconv"

	"gitlab.com/mayachain/mayanode/common"
	cosmos "gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
)

func ParseAddLiquidityMemoV1(ctx cosmos.Context, keeper keeper.Keeper, asset common.Asset, parts []string) (AddLiquidityMemo, error) {
	var err error
	addr := common.NoAddress
	affAddr := common.NoAddress
	affPts := cosmos.ZeroUint()
	if len(parts) >= 3 && len(parts[2]) > 0 {
		if keeper == nil {
			addr, err = common.NewAddress(parts[2])
		} else {
			addr, err = FetchAddress(ctx, keeper, parts[2], asset.Chain)
		}
		if err != nil {
			return AddLiquidityMemo{}, err
		}
	}

	if len(parts) > 4 && len(parts[3]) > 0 && len(parts[4]) > 0 {
		if keeper == nil {
			affAddr, err = common.NewAddress(parts[3])
		} else {
			affAddr, err = FetchAddress(ctx, keeper, parts[3], common.BASEChain)
		}
		if err != nil {
			return AddLiquidityMemo{}, err
		}
		var pts uint64
		pts, err = strconv.ParseUint(parts[4], 10, 64)
		if err != nil {
			return AddLiquidityMemo{}, err
		}
		affPts = cosmos.NewUint(pts)
	}
	return NewAddLiquidityMemo(asset, addr, affAddr, affPts, 0), nil
}
