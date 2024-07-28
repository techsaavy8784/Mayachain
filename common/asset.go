package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/blang/semver"
	"github.com/gogo/protobuf/jsonpb"
)

type Assets []Asset

var (
	// EmptyAsset empty asset, not valid
	EmptyAsset = Asset{Chain: EmptyChain, Symbol: "", Ticker: "", Synth: false}
	// RUNEAsset RUNE
	RUNEAsset = Asset{Chain: THORChain, Symbol: "RUNE", Ticker: "RUNE", Synth: false}
	// ATOMAsset ATOM
	ATOMAsset = Asset{Chain: GAIAChain, Symbol: "ATOM", Ticker: "ATOM", Synth: false}
	// BNBAsset BNB
	BNBAsset = Asset{Chain: BNBChain, Symbol: "BNB", Ticker: "BNB", Synth: false}
	// BTCAsset BTC
	BTCAsset = Asset{Chain: BTCChain, Symbol: "BTC", Ticker: "BTC", Synth: false}
	// LTCAsset BTC
	LTCAsset = Asset{Chain: LTCChain, Symbol: "LTC", Ticker: "LTC", Synth: false}
	// BCHAsset BCH
	BCHAsset = Asset{Chain: BCHChain, Symbol: "BCH", Ticker: "BCH", Synth: false}
	// DASHAsset DASH
	DASHAsset = Asset{Chain: DASHChain, Symbol: "DASH", Ticker: "DASH", Synth: false}
	// DOGEAsset DOGE
	DOGEAsset = Asset{Chain: DOGEChain, Symbol: "DOGE", Ticker: "DOGE", Synth: false}
	// ETHAsset ETH
	ETHAsset = Asset{Chain: ETHChain, Symbol: "ETH", Ticker: "ETH", Synth: false}
	// USDTAsset ETH
	USDTAsset   = Asset{Chain: ETHChain, Symbol: "USDT-0xdAC17F958D2ee523a2206206994597C13D831ec7", Ticker: "USDT", Synth: false}
	USDTAssetV1 = Asset{Chain: ETHChain, Symbol: "USDT-0xdAC17F958D2ee523a2206206994597C13D831ec7", Ticker: "ETH", Synth: false}
	// USDCAsset ETH
	USDCAsset   = Asset{Chain: ETHChain, Symbol: "USDC-0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Ticker: "USDC", Synth: false}
	USDCAssetV1 = Asset{Chain: ETHChain, Symbol: "USDC-0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Ticker: "ETH", Synth: false}
	// WSTETHAsset ETH
	WSTETHAsset   = Asset{Chain: ETHChain, Symbol: "WSTETH-0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0", Ticker: "WSTETH", Synth: false}
	WSTETHAssetV1 = Asset{Chain: ETHChain, Symbol: "WSTETH-0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0", Ticker: "ETH", Synth: false}
	// PEPEAsset ETH
	PEPEAsset = Asset{Chain: ETHChain, Symbol: "PEPE-0x25D887CE7A35172C62FEBFD67A1856F20FAEBB00", Ticker: "PEPE", Synth: false}

	// ETHAsset ARB
	AETHAsset = Asset{Chain: ARBChain, Symbol: "ETH", Ticker: "ETH", Synth: false}
	// USDTAsset ARB
	AUSDTAsset = Asset{Chain: ARBChain, Symbol: "USDT-0XFD086BC7CD5C481DCC9C85EBE478A1C0B69FCBB9", Ticker: "USDT", Synth: false}
	// USDCAsset ARB
	AUSDCAsset = Asset{Chain: ARBChain, Symbol: "USDC-0XAF88D065E77C8CC2239327C5EDB3A432268E5831", Ticker: "USDC", Synth: false}
	// DAIAsset ARB
	ADAIAsset = Asset{Chain: ARBChain, Symbol: "DAI-0XDA10009CBD5D07DD0CECC66161FC93D7C9000DA1", Ticker: "DAI", Synth: false}
	// PEPEAsset ARB
	APEPEAsset = Asset{Chain: ARBChain, Symbol: "PEPE-0X25D887CE7A35172C62FEBFD67A1856F20FAEBB00", Ticker: "PEPE", Synth: false}
	// WSTETHAsset ARB
	AWSTETHAsset = Asset{Chain: ARBChain, Symbol: "WSTETH-0X5979D7B546E38E414F7E9822514BE443A4800529", Ticker: "WSTETH", Synth: false}
	// WBTCAsset ARB
	AWBTCAsset = Asset{Chain: ARBChain, Symbol: "WBTC-0X2F2A2543B76A4166549F7AAB2E75BEF0AEFC5B0F", Ticker: "WBTC", Synth: false}
	// ATGTAsset ARB
	ATGTAsset = Asset{Chain: ARBChain, Symbol: "TGT-0x429FED88F10285E61B12BDF00848315FBDFCC341", Ticker: "TGT", Synth: false}

	// KUJIAsset KUJI
	KUJIAsset = Asset{Chain: KUJIChain, Symbol: "KUJI", Ticker: "KUJI", Synth: false}
	// USKAsset KUJI
	USKAsset = Asset{Chain: KUJIChain, Symbol: "USK", Ticker: "KUJI", Synth: false}

	// AVAXAsset AVAX
	AVAXAsset = Asset{Chain: AVAXChain, Symbol: "AVAX", Ticker: "AVAX", Synth: false}

	// BaseNative CACAO on mayachain
	BaseNative = Asset{Chain: BASEChain, Symbol: "CACAO", Ticker: "CACAO", Synth: false}
	MayaNative = Asset{Chain: BASEChain, Symbol: "MAYA", Ticker: "MAYA", Synth: false}
)

// NewAsset parse the given input into Asset object
func NewAsset(input string) (Asset, error) {
	var err error
	var asset Asset
	var sym string
	var parts []string
	if strings.Count(input, "/") > 0 {
		parts = strings.SplitN(input, "/", 2)
		asset.Synth = true
	} else {
		parts = strings.SplitN(input, ".", 2)
		asset.Synth = false
	}
	if len(parts) == 1 {
		asset.Chain = BASEChain
		sym = parts[0]
	} else {
		asset.Chain, err = NewChain(parts[0])
		if err != nil {
			return EmptyAsset, err
		}
		sym = parts[1]
	}

	asset.Symbol, err = NewSymbol(sym)
	if err != nil {
		return EmptyAsset, err
	}

	parts = strings.SplitN(sym, "-", 2)
	asset.Ticker, err = NewTicker(parts[0])
	if err != nil {
		return EmptyAsset, err
	}

	return asset, nil
}

func NewAssetWithShortCodes(version semver.Version, input string) (Asset, error) {
	switch {
	case version.GTE(semver.MustParse("1.110.0")):
		return NewAssetWithShortCodesV110(input)
	default:
		return NewAsset(input)
	}
}

func NewAssetWithShortCodesV110(input string) (Asset, error) {
	shorts := make(map[string]string)

	// One letter
	shorts[AETHAsset.ShortCode()] = AETHAsset.String()
	shorts[BaseAsset().ShortCode()] = BaseAsset().String()
	shorts[BTCAsset.ShortCode()] = BTCAsset.String()
	shorts[DASHAsset.ShortCode()] = DASHAsset.String()
	shorts[ETHAsset.ShortCode()] = ETHAsset.String()
	shorts[KUJIAsset.ShortCode()] = KUJIAsset.String()
	shorts[RUNEAsset.ShortCode()] = RUNEAsset.String()

	// Two letter
	shorts[AETHAsset.TwoLetterShortCode()] = AETHAsset.String()
	shorts[BaseAsset().TwoLetterShortCode()] = BaseAsset().String()
	shorts[BTCAsset.TwoLetterShortCode()] = BTCAsset.String()
	shorts[DASHAsset.TwoLetterShortCode()] = DASHAsset.String()
	shorts[ETHAsset.TwoLetterShortCode()] = ETHAsset.String()
	shorts[KUJIAsset.TwoLetterShortCode()] = KUJIAsset.String()
	shorts[RUNEAsset.TwoLetterShortCode()] = RUNEAsset.String()

	// Two letter non-gas assets
	// ARB
	shorts[AUSDTAsset.TwoLetterShortCode()] = AUSDTAsset.String()
	shorts[AUSDCAsset.TwoLetterShortCode()] = AUSDCAsset.String()
	shorts[ADAIAsset.TwoLetterShortCode()] = ADAIAsset.String()
	shorts[APEPEAsset.TwoLetterShortCode()] = APEPEAsset.String()
	shorts[AWSTETHAsset.TwoLetterShortCode()] = AWSTETHAsset.String()
	shorts[AWBTCAsset.TwoLetterShortCode()] = AWBTCAsset.String()
	// ETH
	shorts[USDTAsset.TwoLetterShortCode()] = USDTAsset.String()
	shorts[USDCAsset.TwoLetterShortCode()] = USDCAsset.String()
	shorts[PEPEAsset.TwoLetterShortCode()] = PEPEAsset.String()
	shorts[WSTETHAsset.TwoLetterShortCode()] = WSTETHAsset.String()
	// KUJI
	shorts[USKAsset.TwoLetterShortCode()] = USKAsset.String()

	long, ok := shorts[input]
	if ok {
		input = long
	}

	return NewAsset(input)
}

// Equals determinate whether two assets are equivalent
func (a Asset) Equals(a2 Asset) bool {
	return a.Chain.Equals(a2.Chain) && a.Symbol.Equals(a2.Symbol) && a.Ticker.Equals(a2.Ticker) && a.Synth == a2.Synth
}

func (a Asset) GetChain() Chain {
	if a.Synth {
		return BASEChain
	}
	return a.Chain
}

// Get layer1 asset version
func (a Asset) GetLayer1Asset() Asset {
	if !a.IsSyntheticAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  false,
	}
}

// Get synthetic asset of asset
func (a Asset) GetSyntheticAsset() Asset {
	if a.IsSyntheticAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  true,
	}
}

// Check if asset is a pegged asset
func (a Asset) IsSyntheticAsset() bool {
	return a.Synth
}

func (a Asset) IsVaultAsset() bool {
	return a.IsSyntheticAsset()
}

// Native return native asset, only relevant on THORChain
func (a Asset) Native() string {
	if a.IsBase() {
		return "cacao"
	}
	if a.Equals(MayaNative) {
		return "maya"
	}
	return strings.ToLower(a.String())
}

// IsEmpty will be true when any of the field is empty, chain,symbol or ticker
func (a Asset) IsEmpty() bool {
	return a.Chain.IsEmpty() || a.Symbol.IsEmpty() || a.Ticker.IsEmpty()
}

// String implement fmt.Stringer , return the string representation of Asset
func (a Asset) String() string {
	div := "."
	if a.Synth {
		div = "/"
	}
	return fmt.Sprintf("%s%s%s", a.Chain.String(), div, a.Symbol.String())
}

// ShortCode returns the short code for the asset.
func (a Asset) ShortCode() string {
	switch a.String() {
	case "ARB.ETH":
		return "a"
	case "BTC.BTC":
		return "b"
	case "DASH.DASH":
		return "d"
	case "ETH.ETH":
		return "e"
	case "KUJI.KUJI":
		return "k"
	case "MAYA.CACAO":
		return "c"
	case "THOR.RUNE":
		return "r"
	default:
		return ""
	}
}

// ShortCode returns the short code for the asset.
func (a Asset) TwoLetterShortCode() string {
	switch a.String() {
	case "ARB.ETH":
		return "ae"
	case "ARB.USDT-0XFD086BC7CD5C481DCC9C85EBE478A1C0B69FCBB9":
		return "at"
	case "ARB.USDC-0XAF88D065E77C8CC2239327C5EDB3A432268E5831":
		return "ac"
	case "ARB.DAI-0XDA10009CBD5D07DD0CECC66161FC93D7C9000DA1":
		return "ad"
	case "ARB.PEPE-0X25D887CE7A35172C62FEBFD67A1856F20FAEBB00":
		return "ap"
	case "ARB.WSTETH-0X5979D7B546E38E414F7E9822514BE443A4800529":
		return "aw"
	case "ARB.WBTC-0X2F2A2543B76A4166549F7AAB2E75BEF0AEFC5B0F":
		return "ab"
	case "BTC.BTC":
		return "bb"
	case "DASH.DASH":
		return "dd"
	case "ETH.USDT-0XDAC17F958D2EE523A2206206994597C13D831EC7":
		return "et"
	case "ETH.USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48":
		return "ec"
	case "ETH.PEPE-0X6982508145454CE325DDBE47A25D4EC3D2311933":
		return "ep"
	case "ETH.WSTETH-0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0":
		return "ew"
	case "KUJI.USK":
		return "ku"
	case "MAYA.CACAO":
		return "mc"
	case "THOR.RUNE":
		return "tr"
	default:
		return ""
	}
}

// IsGasAsset check whether asset is base asset used to pay for gas
func (a Asset) IsGasAsset() bool {
	gasAsset := a.GetChain().GetGasAsset()
	if gasAsset.IsEmpty() {
		return false
	}
	return a.Equals(gasAsset)
}

// IsCacao is a helper function ,return true only when the asset represent RUNE
func (a Asset) IsBase() bool {
	return a.Equals(BaseNative)
}

// IsNativeRune is a helper function, return true only when the asset represent NATIVE RUNE
func (a Asset) IsNativeBase() bool {
	return a.IsBase() && a.Chain.IsBASEChain()
}

// IsNative is a helper function, returns true when the asset is a native
// asset to THORChain (ie rune, a synth, etc)
func (a Asset) IsNative() bool {
	return a.GetChain().IsBASEChain()
}

// IsBNB is a helper function, return true only when the asset represent BNB
func (a Asset) IsBNB() bool {
	return a.Equals(BNBAsset)
}

// MarshalJSON implement Marshaler interface
func (a Asset) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// UnmarshalJSON implement Unmarshaler interface
func (a *Asset) UnmarshalJSON(data []byte) error {
	var err error
	var assetStr string
	if err = json.Unmarshal(data, &assetStr); err != nil {
		return err
	}
	if assetStr == "." {
		*a = EmptyAsset
		return nil
	}
	*a, err = NewAsset(assetStr)
	return err
}

// MarshalJSONPB implement jsonpb.Marshaler
func (a Asset) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) {
	return a.MarshalJSON()
}

// UnmarshalJSONPB implement jsonpb.Unmarshaler
func (a *Asset) UnmarshalJSONPB(unmarshal *jsonpb.Unmarshaler, content []byte) error {
	return a.UnmarshalJSON(content)
}

// Contains checks if the array contains the specified element
func (as *Assets) Contains(a Asset) bool {
	for _, asset := range *as {
		if asset.Equals(a) {
			return true
		}
	}
	return false
}

// BaseAsset return RUNE Asset depends on different environment
func BaseAsset() Asset {
	return BaseNative
}

// Replace pool name "." with a "-" for Mimir key checking.
func (a Asset) MimirString() string {
	return a.Chain.String() + "-" + a.Symbol.String()
}

// GetAsset returns true if the asset exists in the list of assets
func ContainsAsset(asset Asset, assets []Asset) bool {
	for _, a := range assets {
		if a.Equals(asset) {
			return true
		}
	}
	return false
}
