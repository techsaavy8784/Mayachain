package kuji

import "strings"

type KujiAssetMapping struct {
	KujiDenom       string
	KujiDecimals    int
	BASEChainSymbol string
}

// KujiAssetMappings maps a Kuji denom to a BASEChain symbol and provides the asset decimals
// CHANGEME: define assets that should be observed by BASEChain here. This also acts a whitelist.
var KujiAssetMappings = []KujiAssetMapping{
	{
		KujiDenom:       "ukuji",
		KujiDecimals:    6,
		BASEChainSymbol: "KUJI",
	},
	{
		KujiDenom:       "factory/kujira1qk00h5atutpsv900x202pxx42npjr9thg58dnqpa72f2p7m2luase444a7/uusk",
		KujiDecimals:    6,
		BASEChainSymbol: "USK",
	},
	{ // deprecated
		KujiDenom:       "factory/kujira1ygfxn0er40klcnck8thltuprdxlck6wvnpkf2k/uyum",
		KujiDecimals:    6,
		BASEChainSymbol: "YUM",
	},
	{
		KujiDenom:       "ibc/507BE7E33F06026652F519AD4D36716251F2D34DF04514A905D3B19A7D8130F7",
		KujiDecimals:    6,
		BASEChainSymbol: "AXLYUM",
	},
}

func GetAssetByKujiDenom(denom string) (KujiAssetMapping, bool) {
	for _, asset := range KujiAssetMappings {
		if strings.EqualFold(asset.KujiDenom, denom) {
			return asset, true
		}
	}
	return KujiAssetMapping{}, false
}

func GetAssetByMayachainSymbol(symbol string) (KujiAssetMapping, bool) {
	for _, asset := range KujiAssetMappings {
		if strings.EqualFold(asset.BASEChainSymbol, symbol) {
			return asset, true
		}
	}
	return KujiAssetMapping{}, false
}
