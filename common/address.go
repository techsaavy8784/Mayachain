package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/blang/semver"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/bech32"
	dogchaincfg "github.com/eager7/dogd/chaincfg"
	"github.com/eager7/dogutil"
	eth "github.com/ethereum/go-ethereum/common"
	bchchaincfg "github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"
	dashutil "gitlab.com/mayachain/dashd-go/btcutil"
	dashchaincfg "gitlab.com/mayachain/dashd-go/chaincfg"

	"gitlab.com/mayachain/mayanode/common/cosmos"
)

type Address string

const (
	NoAddress      = Address("")
	NoopAddress    = Address("noop")
	EVMNullAddress = Address("0x0000000000000000000000000000000000000000")
)

var alphaNumRegex = regexp.MustCompile("^[:A-Za-z0-9]*$")

// NewAddress create a new Address. Supports Binance, Bitcoin, and Ethereum
func NewAddress(address string) (Address, error) {
	if len(address) == 0 {
		return NoAddress, nil
	}

	if !alphaNumRegex.MatchString(address) {
		return NoAddress, fmt.Errorf("address format not supported: %s", address)
	}

	// Check is eth address
	if eth.IsHexAddress(address) {
		return Address(address), nil
	}

	// Check bech32 addresses, would succeed any string bech32 encoded (e.g. MAYA, THOR, BNB, ATOM)
	_, _, err := bech32.Decode(address)
	if err == nil {
		return Address(address), nil
	}

	// Check other BTC address formats with mainnet
	_, err = btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check BTC address formats with testnet
	_, err = btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check other LTC address formats with mainnet
	_, err = ltcutil.DecodeAddress(address, &ltcchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check LTC address formats with testnet
	_, err = ltcutil.DecodeAddress(address, &ltcchaincfg.TestNet4Params)
	if err == nil {
		return Address(address), nil
	}

	// Check BCH address formats with mainnet
	_, err = bchutil.DecodeAddress(address, &bchchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check BCH address formats with testnet
	_, err = bchutil.DecodeAddress(address, &bchchaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check BCH address formats with mocknet
	_, err = bchutil.DecodeAddress(address, &bchchaincfg.RegressionNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check DASH address formats with mainnet
	_, err = dashutil.DecodeAddress(address, &dashchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check DASH address formats with testnet
	_, err = dashutil.DecodeAddress(address, &dashchaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check DASH address formats with mocknet
	_, err = dashutil.DecodeAddress(address, &dashchaincfg.RegressionNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check DOGE address formats with mainnet
	_, err = dogutil.DecodeAddress(address, &dogchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check DOGE address formats with testnet
	_, err = dogutil.DecodeAddress(address, &dogchaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check DOGE address formats with mocknet
	_, err = dogutil.DecodeAddress(address, &dogchaincfg.RegressionNetParams)
	if err == nil {
		return Address(address), nil
	}

	return NoAddress, fmt.Errorf("address format not supported: %s", address)
}

// IsValidBCHAddress determinate whether the address is a valid new BCH address format
func (addr Address) IsValidBCHAddress() bool {
	// Check mainnet other formats
	bchAddr, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
	if err == nil {
		switch bchAddr.(type) {
		case *bchutil.LegacyAddressPubKeyHash, *bchutil.LegacyAddressScriptHash:
			return false
		}
		return true
	}
	bchAddr, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
	if err == nil {
		switch bchAddr.(type) {
		case *bchutil.LegacyAddressPubKeyHash, *bchutil.LegacyAddressScriptHash:
			return false
		}
		return true
	}
	bchAddr, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
	if err == nil {
		switch bchAddr.(type) {
		case *bchutil.LegacyAddressPubKeyHash, *bchutil.LegacyAddressScriptHash:
			return false
		}
		return true
	}
	return false
}

// ConvertToNewBCHAddressFormat convert the given BCH to new address format
func ConvertToNewBCHAddressFormat(addr Address, version semver.Version) (Address, error) {
	if !addr.IsChain(BCHChain, version) {
		return NoAddress, fmt.Errorf("address(%s) is not BCH chain", addr)
	}
	network := CurrentChainNetwork
	var param *bchchaincfg.Params
	switch network {
	case MockNet:
		param = &bchchaincfg.RegressionNetParams
	case TestNet:
		param = &bchchaincfg.TestNet3Params
	case MainNet:
		param = &bchchaincfg.MainNetParams
	case StageNet:
		param = &bchchaincfg.MainNetParams
	}
	bchAddr, err := bchutil.DecodeAddress(addr.String(), param)
	if err != nil {
		return NoAddress, fmt.Errorf("fail to decode address(%s), %w", addr, err)
	}
	return getBCHAddress(bchAddr, param)
}

func getBCHAddress(address bchutil.Address, cfg *bchchaincfg.Params) (Address, error) {
	switch address.(type) {
	case *bchutil.LegacyAddressPubKeyHash, *bchutil.AddressPubKeyHash:
		h, err := bchutil.NewAddressPubKeyHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new pubkey hash address: %w", err)
		}
		return NewAddress(h.String())
	case *bchutil.LegacyAddressScriptHash, *bchutil.AddressScriptHash:
		h, err := bchutil.NewAddressScriptHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new address script hash address: %w", err)
		}
		return NewAddress(h.String())
	}
	return NoAddress, fmt.Errorf("invalid address type")
}

// ConvertToNewBCHAddressFormatV83 convert the given BCH to new address format
func ConvertToNewBCHAddressFormatV83(addr Address, version semver.Version) (Address, error) {
	if !addr.IsChain(BCHChain, version) {
		return NoAddress, fmt.Errorf("address(%s) is not BCH chain", addr)
	}
	network := CurrentChainNetwork
	var param *bchchaincfg.Params
	switch network {
	case MockNet:
		param = &bchchaincfg.RegressionNetParams
	case TestNet:
		param = &bchchaincfg.TestNet3Params
	case MainNet:
		param = &bchchaincfg.MainNetParams
	case StageNet:
		param = &bchchaincfg.MainNetParams
	}
	bchAddr, err := bchutil.DecodeAddress(addr.String(), param)
	if err != nil {
		return NoAddress, fmt.Errorf("fail to decode address(%s), %w", addr, err)
	}
	return getBCHAddressV83(bchAddr, param)
}

func getBCHAddressV83(address bchutil.Address, cfg *bchchaincfg.Params) (Address, error) {
	switch address.(type) {
	case *bchutil.LegacyAddressPubKeyHash, *bchutil.AddressPubKeyHash:
		h, err := bchutil.NewAddressPubKeyHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new pubkey hash address: %w", err)
		}
		return NewAddress(h.String())
	case *bchutil.LegacyAddressScriptHash, *bchutil.AddressScriptHash:
		h, err := bchutil.NewAddressScriptHashFromHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new address script hash address: %w", err)
		}
		return NewAddress(h.String())
	}
	return NoAddress, fmt.Errorf("invalid address type")
}

func (addr Address) IsChain(chain Chain, version semver.Version) bool {
	switch {
	case version.GTE(semver.MustParse("1.108.0")):
		return addr.IsChainV108(chain)
	default:
		return addr.IsChainV107(chain)
	}
}

func (addr Address) IsChainV108(chain Chain) bool {
	if chain.IsEVM() {
		return strings.HasPrefix(addr.String(), "0x")
	}
	switch chain {
	case BNBChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "bnb" || prefix == "tbnb"
	case AZTECChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "aztec" || prefix == "taztec" || prefix == "saztec"
	case BASEChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "maya" || prefix == "tmaya" || prefix == "smaya"
	case GAIAChain:
		// Note: Gaia does not use a special prefix for testnet
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "cosmos"
	case THORChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "thor" || prefix == "tthor"
	case KUJIChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "kujira"
	case BTCChain:
		prefix, _, err := bech32.Decode(addr.String())
		if err == nil && (prefix == "bc" || prefix == "tb") {
			return true
		}
		// Check mainnet other formats
		_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		return false
	case LTCChain:
		prefix, _, err := bech32.Decode(addr.String())
		if err == nil && (prefix == "ltc" || prefix == "tltc" || prefix == "rltc") {
			return true
		}
		// Check mainnet other formats
		_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.TestNet4Params)
		if err == nil {
			return true
		}
		return false
	case BCHChain:
		// Check mainnet other formats
		_, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	case DASHChain:
		// Check mainnet other formats
		_, err := dashutil.DecodeAddress(addr.String(), &dashchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = dashutil.DecodeAddress(addr.String(), &dashchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = dashutil.DecodeAddress(addr.String(), &dashchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	case DOGEChain:
		// Check mainnet other formats
		_, err := dogutil.DecodeAddress(addr.String(), &dogchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	default:
		return true // if THORNode don't specifically check a chain yet, assume its ok.
	}
}

func (addr Address) IsChainV107(chain Chain) bool {
	if chain.IsEVM() {
		return strings.HasPrefix(addr.String(), "0x")
	}
	switch chain {
	case BNBChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "bnb" || prefix == "tbnb"
	case AZTECChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "aztec" || prefix == "taztec" || prefix == "saztec"
	case BASEChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "maya" || prefix == "tmaya" || prefix == "smaya"
	case GAIAChain:
		// Note: Gaia does not use a special prefix for testnet
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "cosmos"
	case THORChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "thor" || prefix == "tthor"
	case BTCChain:
		prefix, _, err := bech32.Decode(addr.String())
		if err == nil && (prefix == "bc" || prefix == "tb") {
			return true
		}
		// Check mainnet other formats
		_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		return false
	case LTCChain:
		prefix, _, err := bech32.Decode(addr.String())
		if err == nil && (prefix == "ltc" || prefix == "tltc" || prefix == "rltc") {
			return true
		}
		// Check mainnet other formats
		_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.TestNet4Params)
		if err == nil {
			return true
		}
		return false
	case BCHChain:
		// Check mainnet other formats
		_, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	case DASHChain:
		// Check mainnet other formats
		_, err := dashutil.DecodeAddress(addr.String(), &dashchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = dashutil.DecodeAddress(addr.String(), &dashchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = dashutil.DecodeAddress(addr.String(), &dashchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	case DOGEChain:
		// Check mainnet other formats
		_, err := dogutil.DecodeAddress(addr.String(), &dogchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	default:
		return true // if THORNode don't specifically check a chain yet, assume its ok.
	}
}

func (addr Address) GetChain(version semver.Version) Chain {
	switch {
	case version.GTE(semver.MustParse("1.109.0")):
		return addr.getChainV109(version)
	case version.GTE(semver.MustParse("1.107.0")):
		return addr.getChainV107(version)
	default:
		return addr.getChainV105(version)
	}
}

func (addr Address) getChainV109(version semver.Version) Chain {
	for _, chain := range []Chain{ETHChain, BNBChain, BASEChain, BTCChain, LTCChain, BCHChain, DASHChain, DOGEChain, THORChain, GAIAChain, KUJIChain, AVAXChain, ARBChain} {
		if addr.IsChain(chain, version) {
			return chain
		}
	}
	return EmptyChain
}

func (addr Address) getChainV107(version semver.Version) Chain {
	for _, chain := range []Chain{ETHChain, BNBChain, BASEChain, BTCChain, LTCChain, BCHChain, DASHChain, DOGEChain, THORChain, GAIAChain, KUJIChain, AVAXChain} {
		if addr.IsChain(chain, version) {
			return chain
		}
	}
	return EmptyChain
}

func (addr Address) getChainV105(version semver.Version) Chain {
	for _, chain := range []Chain{ETHChain, BNBChain, BASEChain, BTCChain, LTCChain, BCHChain, DASHChain, DOGEChain, BASEChain, GAIAChain, AVAXChain} {
		if addr.IsChain(chain, version) {
			return chain
		}
	}
	return EmptyChain
}

func (addr Address) GetNetwork(ver semver.Version, chain Chain) ChainNetwork {
	mainNetPredicate := func() ChainNetwork {
		if CurrentChainNetwork == StageNet {
			return StageNet
		}
		return MainNet
	}
	// EVM addresses don't have different prefixes per network
	if chain.IsEVM() {
		return CurrentChainNetwork
	}
	switch chain {
	case BNBChain:
		prefix, _, _ := bech32.Decode(addr.String())
		if strings.EqualFold(prefix, "bnb") {
			return mainNetPredicate()
		}
		if strings.EqualFold(prefix, "tbnb") {
			return TestNet
		}
	case AZTECChain:
		return CurrentChainNetwork
	case BASEChain:
		prefix, _, _ := bech32.Decode(addr.String())
		if strings.EqualFold(prefix, "maya") {
			return mainNetPredicate()
		}
		if strings.EqualFold(prefix, "tmaya") {
			return TestNet
		}
		if strings.EqualFold(prefix, "smaya") {
			return StageNet
		}
	case KUJIChain:
		return CurrentChainNetwork
	case THORChain:
		prefix, _, _ := bech32.Decode(addr.String())
		if strings.EqualFold(prefix, "thor") {
			return mainNetPredicate()
		}
		if strings.EqualFold(prefix, "tthor") {
			return TestNet
		}
	case BTCChain:
		prefix, _, _ := bech32.Decode(addr.String())
		switch prefix {
		case "bc":
			return mainNetPredicate()
		case "tb":
			return TestNet
		case "bcrt":
			return MockNet
		default:
			_, err := btcutil.DecodeAddress(addr.String(), &chaincfg.MainNetParams)
			if err == nil {
				return mainNetPredicate()
			}
			_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.TestNet3Params)
			if err == nil {
				return TestNet
			}
			_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.RegressionNetParams)
			if err == nil {
				return MockNet
			}
		}
	case LTCChain:
		prefix, _, _ := bech32.Decode(addr.String())
		switch prefix {
		case "ltc":
			return mainNetPredicate()
		case "tltc":
			return TestNet
		case "rltc":
			return MockNet
		default:
			_, err := ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.MainNetParams)
			if err == nil {
				return mainNetPredicate()
			}
			_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.TestNet4Params)
			if err == nil {
				return TestNet
			}
			_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.RegressionNetParams)
			if err == nil {
				return MockNet
			}
		}
	case BCHChain:
		// Check mainnet other formats
		_, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
		if err == nil {
			return mainNetPredicate()
		}
		// Check testnet other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
		if err == nil {
			return TestNet
		}
		// Check mocknet / regression other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
		if err == nil {
			return MockNet
		}
	case DASHChain:
		// Check mainnet other formats
		_, err := dashutil.DecodeAddress(addr.String(), &dashchaincfg.MainNetParams)
		if err == nil {
			return mainNetPredicate()
		}
		// Check testnet other formats
		_, err = dashutil.DecodeAddress(addr.String(), &dashchaincfg.TestNet3Params)
		if err == nil {
			return TestNet
		}
		// Check mocknet / regression other formats
		_, err = dashutil.DecodeAddress(addr.String(), &dashchaincfg.RegressionNetParams)
		if err == nil {
			return MockNet
		}
	case DOGEChain:
		// Check mainnet other formats
		_, err := dogutil.DecodeAddress(addr.String(), &dogchaincfg.MainNetParams)
		if err == nil {
			return mainNetPredicate()
		}
		// Check testnet other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.TestNet3Params)
		if err == nil {
			return TestNet
		}
		// Check mocknet / regression other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.RegressionNetParams)
		if err == nil {
			return MockNet
		}
	}
	switch {
	case ver.GTE(semver.MustParse("1.93.0")):
		return CurrentChainNetwork
	default:
		return MockNet
	}
}

func (addr Address) AccAddress() (cosmos.AccAddress, error) {
	return cosmos.AccAddressFromBech32(addr.String())
}

func (addr Address) Equals(addr2 Address) bool {
	return strings.EqualFold(addr.String(), addr2.String())
}

func (addr Address) IsEmpty() bool {
	return strings.TrimSpace(addr.String()) == ""
}

func (addr Address) IsNoop() bool {
	return addr.Equals(NoopAddress)
}

func (addr Address) String() string {
	return string(addr)
}
