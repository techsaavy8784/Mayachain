package mayachain

import (
	"fmt"
	"strings"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"

	"github.com/blang/semver"
)

// TXTYPE:STATE1:STATE2:STATE3:FINALMEMO

type TxType uint8

const (
	TxUnknown TxType = iota
	TxAdd
	TxWithdraw
	TxSwap
	TxLimitOrder
	TxOutbound
	TxDonate
	TxBond
	TxUnbond
	TxLeave
	TxYggdrasilFund
	TxYggdrasilReturn
	TxReserve
	TxRefund
	TxMigrate
	TxRagnarok
	TxNoOp
	TxConsolidate
	TxMAYAName
	TxForgiveSlash
)

var stringToTxTypeMap = map[string]TxType{
	"add":         TxAdd,
	"+":           TxAdd,
	"withdraw":    TxWithdraw,
	"wd":          TxWithdraw,
	"-":           TxWithdraw,
	"swap":        TxSwap,
	"s":           TxSwap,
	"=":           TxSwap,
	"out":         TxOutbound,
	"donate":      TxDonate,
	"d":           TxDonate,
	"bond":        TxBond,
	"unbond":      TxUnbond,
	"leave":       TxLeave,
	"yggdrasil+":  TxYggdrasilFund,
	"yggdrasil-":  TxYggdrasilReturn,
	"reserve":     TxReserve,
	"refund":      TxRefund,
	"migrate":     TxMigrate,
	"ragnarok":    TxRagnarok,
	"noop":        TxNoOp,
	"consolidate": TxConsolidate,
	"name":        TxMAYAName,
	"n":           TxMAYAName,
	"~":           TxMAYAName,
}

var txToStringMap = map[TxType]string{
	TxAdd:             "add",
	TxWithdraw:        "withdraw",
	TxSwap:            "swap",
	TxOutbound:        "out",
	TxRefund:          "refund",
	TxDonate:          "donate",
	TxBond:            "bond",
	TxUnbond:          "unbond",
	TxLeave:           "leave",
	TxYggdrasilFund:   "yggdrasil+",
	TxYggdrasilReturn: "yggdrasil-",
	TxReserve:         "reserve",
	TxMigrate:         "migrate",
	TxRagnarok:        "ragnarok",
	TxNoOp:            "noop",
	TxConsolidate:     "consolidate",
	TxMAYAName:        "mayaname",
	TxForgiveSlash:    "forgive_slash",
}

// converts a string into a txType
func StringToTxType(s string) (TxType, error) {
	// THORNode can support Abbreviated MEMOs , usually it is only one character
	sl := strings.ToLower(s)
	if t, ok := stringToTxTypeMap[sl]; ok {
		return t, nil
	}

	return TxUnknown, fmt.Errorf("invalid tx type: %s", s)
}

func (tx TxType) IsInbound() bool {
	switch tx {
	case TxAdd, TxWithdraw, TxSwap, TxDonate, TxBond, TxUnbond, TxLeave, TxReserve, TxNoOp, TxMAYAName, TxForgiveSlash:
		return true
	default:
		return false
	}
}

func (tx TxType) IsOutbound() bool {
	switch tx {
	case TxOutbound, TxRefund, TxRagnarok:
		return true
	default:
		return false
	}
}

func (tx TxType) IsInternal() bool {
	switch tx {
	case TxYggdrasilFund, TxYggdrasilReturn, TxMigrate, TxConsolidate:
		return true
	default:
		return false
	}
}

// HasOutbound whether the txtype might trigger outbound tx
func (tx TxType) HasOutbound() bool {
	switch tx {
	case TxAdd, TxBond, TxDonate, TxYggdrasilReturn, TxReserve, TxMigrate, TxRagnarok:
		return false
	default:
		return true
	}
}

func (tx TxType) IsEmpty() bool {
	return tx == TxUnknown
}

// Check if two txTypes are the same
func (tx TxType) Equals(tx2 TxType) bool {
	return tx == tx2
}

// Converts a txType into a string
func (tx TxType) String() string {
	return txToStringMap[tx]
}

type Memo interface {
	IsType(tx TxType) bool
	GetType() TxType
	IsEmpty() bool
	IsInbound() bool
	IsOutbound() bool
	IsInternal() bool
	String() string
	GetAsset() common.Asset
	GetAmount() cosmos.Uint
	GetDestination() common.Address
	GetSlipLimit() cosmos.Uint
	GetTxID() common.TxID
	GetAccAddress() cosmos.AccAddress
	GetBlockHeight() int64
	GetDexAggregator() string
	GetDexTargetAddress() string
	GetDexTargetLimit() *cosmos.Uint
}

type MemoBase struct {
	TxType TxType
	Asset  common.Asset
}

func (m MemoBase) String() string                   { return "" }
func (m MemoBase) GetType() TxType                  { return m.TxType }
func (m MemoBase) IsType(tx TxType) bool            { return m.TxType.Equals(tx) }
func (m MemoBase) GetAsset() common.Asset           { return m.Asset }
func (m MemoBase) GetAmount() cosmos.Uint           { return cosmos.ZeroUint() }
func (m MemoBase) GetDestination() common.Address   { return "" }
func (m MemoBase) GetSlipLimit() cosmos.Uint        { return cosmos.ZeroUint() }
func (m MemoBase) GetTxID() common.TxID             { return "" }
func (m MemoBase) GetAccAddress() cosmos.AccAddress { return cosmos.AccAddress{} }
func (m MemoBase) GetBlockHeight() int64            { return 0 }
func (m MemoBase) IsOutbound() bool                 { return m.TxType.IsOutbound() }
func (m MemoBase) IsInbound() bool                  { return m.TxType.IsInbound() }
func (m MemoBase) IsInternal() bool                 { return m.TxType.IsInternal() }
func (m MemoBase) IsEmpty() bool                    { return m.TxType.IsEmpty() }
func (m MemoBase) GetDexAggregator() string         { return "" }
func (m MemoBase) GetDexTargetAddress() string      { return "" }
func (m MemoBase) GetDexTargetLimit() *cosmos.Uint  { return nil }

func parseBase(version semver.Version, memo string) (MemoBase, []string, error) {
	if version.GTE(semver.MustParse("1.108.0")) {
		memo = strings.Split(memo, "|")[0]
	}
	parts := strings.Split(memo, ":")
	mem := MemoBase{TxType: TxUnknown}
	if len(memo) == 0 {
		return mem, parts, fmt.Errorf("memo can't be empty")
	}
	var err error
	mem.TxType, err = StringToTxType(parts[0])
	if err != nil {
		return mem, parts, err
	}

	switch mem.TxType {
	case TxDonate, TxAdd, TxSwap, TxWithdraw:
		if len(parts) < 2 {
			return mem, parts, fmt.Errorf("cannot parse given memo: length %d", len(parts))
		}
		mem.Asset, err = common.NewAssetWithShortCodes(version, parts[1])
		if err != nil {
			return mem, parts, err
		}
	}

	return mem, parts, nil
}

func ParseMemo(version semver.Version, memo string) (mem Memo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panicked parsing memo(%s), err: %s", memo, r)
		}
	}()

	mem, parts, err := parseBase(version, memo)
	if err != nil {
		return mem, err
	}

	asset := mem.GetAsset()

	switch mem.GetType() {
	case TxLeave:
		return ParseLeaveMemo(parts)
	case TxDonate:
		return NewDonateMemo(asset), nil
	case TxAdd:
		return ParseAddLiquidityMemo(cosmos.Context{}, nil, asset, parts)
	case TxWithdraw:
		return ParseWithdrawLiquidityMemo(cosmos.Context{}, nil, asset, parts)
	case TxSwap:
		return ParseSwapMemo(cosmos.Context{}, nil, asset, parts)
	case TxOutbound:
		return ParseOutboundMemo(parts)
	case TxRefund:
		return ParseRefundMemo(parts)
	case TxBond:
		return ParseBondMemo(version, parts)
	case TxUnbond:
		return ParseUnbondMemo(version, parts)
	case TxYggdrasilFund:
		return ParseYggdrasilFundMemo(parts)
	case TxYggdrasilReturn:
		return ParseYggdrasilReturnMemo(parts)
	case TxReserve:
		return NewReserveMemo(), nil
	case TxMigrate:
		return ParseMigrateMemo(parts)
	case TxRagnarok:
		return ParseRagnarokMemo(parts)
	case TxNoOp:
		return ParseNoOpMemo(parts)
	case TxConsolidate:
		return ParseConsolidateMemo(parts)
	case TxForgiveSlash:
		return ParseForgiveSlashMemo(parts)
	default:
		return mem, fmt.Errorf("TxType not supported: %s", mem.GetType().String())
	}
}

func ParseMemoWithMAYANames(ctx cosmos.Context, keeper keeper.Keeper, memo string) (mem Memo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panicked parsing memo(%s), err: %s", memo, r)
		}
	}()

	mem, parts, err := parseBase(keeper.GetVersion(), memo)
	if err != nil {
		return mem, err
	}

	asset := mem.GetAsset()

	switch mem.GetType() {
	case TxLeave:
		return ParseLeaveMemo(parts)
	case TxDonate:
		return NewDonateMemo(asset), nil
	case TxAdd:
		return ParseAddLiquidityMemo(ctx, keeper, asset, parts)
	case TxWithdraw:
		return ParseWithdrawLiquidityMemo(ctx, keeper, asset, parts)
	case TxSwap:
		return ParseSwapMemo(ctx, keeper, asset, parts)
	case TxOutbound:
		return ParseOutboundMemo(parts)
	case TxRefund:
		return ParseRefundMemo(parts)
	case TxBond:
		return ParseBondMemo(keeper.GetVersion(), parts)
	case TxUnbond:
		return ParseUnbondMemo(keeper.GetVersion(), parts)
	case TxYggdrasilFund:
		return ParseYggdrasilFundMemo(parts)
	case TxYggdrasilReturn:
		return ParseYggdrasilReturnMemo(parts)
	case TxReserve:
		return NewReserveMemo(), nil
	case TxMigrate:
		return ParseMigrateMemo(parts)
	case TxRagnarok:
		return ParseRagnarokMemo(parts)
	case TxNoOp:
		return ParseNoOpMemo(parts)
	case TxConsolidate:
		return ParseConsolidateMemo(parts)
	case TxMAYAName:
		return ParseManageMAYANameMemo(keeper.GetVersion(), parts)
	case TxForgiveSlash:
		return ParseForgiveSlashMemo(parts)
	default:
		return mem, fmt.Errorf("TxType not supported: %s", mem.GetType().String())
	}
}

func FetchAddress(ctx cosmos.Context, keeper keeper.Keeper, name string, chain common.Chain) (common.Address, error) {
	// if name is an address, return as is
	addr, err := common.NewAddress(name)
	if err == nil {
		return addr, nil
	}

	parts := strings.SplitN(name, ".", 2)
	if len(parts) > 1 {
		chain, err = common.NewChain(parts[1])
		if err != nil {
			return common.NoAddress, err
		}
	}

	if keeper.MAYANameExists(ctx, parts[0]) {
		var mayaname types.MAYAName
		mayaname, err = keeper.GetMAYAName(ctx, parts[0])
		if err != nil {
			return common.NoAddress, err
		}
		return mayaname.GetAlias(chain), nil
	}

	return common.NoAddress, fmt.Errorf("%s is not recognizable", name)
}
