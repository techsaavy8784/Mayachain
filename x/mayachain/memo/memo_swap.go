package mayachain

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/blang/semver"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

type SwapMemo struct {
	MemoBase
	Destination          common.Address
	SlipLimit            cosmos.Uint
	AffiliateAddress     common.Address
	AffiliateBasisPoints cosmos.Uint
	DexAggregator        string
	DexTargetAddress     string
	DexTargetLimit       *cosmos.Uint
	OrderType            types.OrderType
	StreamInterval       uint64
	StreamQuantity       uint64
}

func (m SwapMemo) GetDestination() common.Address       { return m.Destination }
func (m SwapMemo) GetSlipLimit() cosmos.Uint            { return m.SlipLimit }
func (m SwapMemo) GetAffiliateAddress() common.Address  { return m.AffiliateAddress }
func (m SwapMemo) GetAffiliateBasisPoints() cosmos.Uint { return m.AffiliateBasisPoints }
func (m SwapMemo) GetDexAggregator() string             { return m.DexAggregator }
func (m SwapMemo) GetDexTargetAddress() string          { return m.DexTargetAddress }
func (m SwapMemo) GetDexTargetLimit() *cosmos.Uint      { return m.DexTargetLimit }
func (m SwapMemo) GetOrderType() types.OrderType        { return m.OrderType }
func (m SwapMemo) GetStreamInterval() uint64            { return m.StreamInterval }
func (m SwapMemo) GetStreamQuantity() uint64            { return m.StreamQuantity }

func (m SwapMemo) String() string {
	return m.string(false)
}

func (m SwapMemo) ShortString() string {
	return m.string(true)
}

func (m SwapMemo) string(short bool) string {
	slipLimit := m.SlipLimit.String()
	if m.SlipLimit.IsZero() {
		slipLimit = ""
	}

	// prefer short notation for generate swap memo
	txType := m.TxType.String()
	if m.TxType == TxSwap {
		txType = "="
	}

	if m.StreamInterval > 0 || m.StreamQuantity > 1 {
		slipLimit = fmt.Sprintf("%s/%d/%d", m.SlipLimit.String(), m.StreamInterval, m.StreamQuantity)
	}

	var assetString string
	if short && len(m.Asset.ShortCode()) > 0 {
		assetString = m.Asset.ShortCode()
	} else {
		assetString = m.Asset.String()
	}

	args := []string{
		txType,
		assetString,
		m.Destination.String(),
		slipLimit,
		m.AffiliateAddress.String(),
		m.AffiliateBasisPoints.String(),
		m.DexAggregator,
		m.DexTargetAddress,
	}

	last := 3
	if !m.SlipLimit.IsZero() || m.StreamInterval > 0 || m.StreamQuantity > 1 {
		last = 4
	}

	if !m.AffiliateAddress.IsEmpty() {
		last = 6
	}

	if m.DexAggregator != "" {
		last = 8
	}

	if m.DexTargetLimit != nil && !m.DexTargetLimit.IsZero() {
		args = append(args, m.DexTargetLimit.String())
		last = 9
	}

	return strings.Join(args[:last], ":")
}

func NewSwapMemo(asset common.Asset, dest common.Address, slip cosmos.Uint, affAddr common.Address, affPts cosmos.Uint, dexAgg, dexTargetAddress string, dexTargetLimit cosmos.Uint, orderType types.OrderType, interval uint64, quan uint64) SwapMemo {
	swapMemo := SwapMemo{
		MemoBase:             MemoBase{TxType: TxSwap, Asset: asset},
		Destination:          dest,
		SlipLimit:            slip,
		AffiliateAddress:     affAddr,
		AffiliateBasisPoints: affPts,
		DexAggregator:        dexAgg,
		DexTargetAddress:     dexTargetAddress,
		OrderType:            orderType,
		StreamInterval:       interval,
		StreamQuantity:       quan,
	}
	if !dexTargetLimit.IsZero() {
		swapMemo.DexTargetLimit = &dexTargetLimit
	}
	return swapMemo
}

func ParseSwapMemo(ctx cosmos.Context, keeper keeper.Keeper, asset common.Asset, parts []string) (SwapMemo, error) {
	if keeper == nil {
		return ParseSwapMemoV1(ctx, keeper, asset, parts)
	}
	switch {

	case keeper.GetVersion().GTE(semver.MustParse("1.110.0")):
		return ParseSwapMemoV110(ctx, keeper, asset, parts)
	case keeper.GetVersion().GTE(semver.MustParse("1.92.0")):
		return ParseSwapMemoV92(ctx, keeper, asset, parts)
	default:
		return ParseSwapMemoV1(ctx, keeper, asset, parts)
	}
}

func ParseSwapMemoV110(ctx cosmos.Context, keeper keeper.Keeper, asset common.Asset, parts []string) (SwapMemo, error) {
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
	var limitStr string
	slip := cosmos.ZeroUint()
	streamInterval := uint64(0)
	streamQuantity := uint64(0)
	if len(parts) > 3 && len(parts[3]) > 0 {
		limitStr = parts[3]
		if strings.Contains(parts[3], "/") {
			split := strings.SplitN(limitStr, "/", 3)
			for i := range split {
				if split[i] == "" {
					split[i] = "0"
				}
			}
			if len(split) < 1 {
				return SwapMemo{}, fmt.Errorf("invalid streaming swap format: %s", parts[3])
			}
			slip, err = cosmos.ParseUint(split[0])
			if err != nil {
				return SwapMemo{}, fmt.Errorf("swap price limit:%s is invalid", parts[3])
			}
			if len(split) > 1 {
				streamInterval, err = strconv.ParseUint(split[1], 10, 64)
				if err != nil {
					return SwapMemo{}, fmt.Errorf("swap stream interval:%s is invalid", parts[3])
				}
			}

			if len(split) > 2 {
				streamQuantity, err = strconv.ParseUint(split[2], 10, 64)
				if err != nil {
					return SwapMemo{}, fmt.Errorf("swap stream quantity:%s is invalid", parts[3])
				}
			}
		} else {
			var amount cosmos.Uint
			amount, err = cosmos.ParseUint(parts[3])
			if err != nil {
				return SwapMemo{}, fmt.Errorf("swap price limit:%s is invalid", parts[3])
			}
			slip = amount
		}
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

	return NewSwapMemo(asset, destination, slip, affAddr, affPts, dexAgg, dexTargetAddress, dexTargetLimit, order, streamInterval, streamQuantity), nil
}
