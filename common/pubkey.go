package common

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	dashutil "gitlab.com/mayachain/dashd-go/btcutil"
	dashchaincfg "gitlab.com/mayachain/dashd-go/chaincfg"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/codec"
	dogchaincfg "github.com/eager7/dogd/chaincfg"
	"github.com/eager7/dogutil"
	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"

	bchchaincfg "github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"

	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/crypto"

	"gitlab.com/mayachain/mayanode/common/cosmos"
)

// PubKey used in thorchain, it should be bech32 encoded string
// thus it will be something like
// tmayapub1addwnpepqt7qug8vk9r3saw8n4r803ydj2g3dqwx0mvq5akhnze86fc536xcy7cau6l
// tmayapub1addwnpepqdqvd4r84lq9m54m5kk9sf4k6kdgavvch723pcgadulxd6ey9u70kujkcf2
type (
	PubKey  string
	PubKeys []PubKey
)

var EmptyPubKey PubKey

var EmptyPubKeySet PubKeySet

// NewPubKey create a new instance of PubKey
// key is bech32 encoded string
func NewPubKey(key string) (PubKey, error) {
	if len(key) == 0 {
		return EmptyPubKey, nil
	}
	_, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, key)
	if err != nil {
		return EmptyPubKey, fmt.Errorf("%s is not bech32 encoded pub key,err : %w", key, err)
	}
	return PubKey(key), nil
}

// NewPubKeyFromCrypto
func NewPubKeyFromCrypto(pk crypto.PubKey) (PubKey, error) {
	tmp, err := codec.FromTmPubKeyInterface(pk)
	if err != nil {
		return EmptyPubKey, fmt.Errorf("fail to create PubKey from crypto.PubKey,err:%w", err)
	}
	s, err := cosmos.Bech32ifyPubKey(cosmos.Bech32PubKeyTypeAccPub, tmp)
	if err != nil {
		return EmptyPubKey, fmt.Errorf("fail to create PubKey from crypto.PubKey,err:%w", err)
	}
	return PubKey(s), nil
}

// Equals check whether two are the same
func (pubKey PubKey) Equals(pubKey1 PubKey) bool {
	return pubKey == pubKey1
}

// IsEmpty to check whether it is empty
func (pubKey PubKey) IsEmpty() bool {
	return len(pubKey) == 0
}

// String stringer implementation
func (pubKey PubKey) String() string {
	return string(pubKey)
}

// EVMPubkeyToAddress converts a pubkey of an EVM chain to the corresponding address
func (pubKey PubKey) EVMPubkeyToAddress() (Address, error) {
	// retrieve compressed pubkey bytes from bechh32 encoded str
	pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
	if err != nil {
		return NoAddress, err
	}
	// parse compressed bytes removing 5 first bytes (amino encoding) to get uncompressed
	pub, err := secp256k1.ParsePubKey(pk.Bytes(), secp256k1.S256())
	if err != nil {
		return NoAddress, err
	}
	str := strings.ToLower(eth.PubkeyToAddress(*pub.ToECDSA()).String())
	return NewAddress(str)
}

// GetAddress will return an address for the given chain
func (pubKey PubKey) GetAddress(chain Chain) (Address, error) {
	if pubKey.IsEmpty() {
		return NoAddress, nil
	}
	chainNetwork := CurrentChainNetwork
	if chain.IsEVM() {
		return pubKey.EVMPubkeyToAddress()
	}
	switch chain {
	case BNBChain, BASEChain, THORChain, GAIAChain, KUJIChain:
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
		if err != nil {
			return NoAddress, err
		}
		str, err := ConvertAndEncode(chain.AddressPrefix(chainNetwork), pk.Address().Bytes())
		if err != nil {
			return NoAddress, fmt.Errorf("fail to bech32 encode the address, err: %w", err)
		}
		return NewAddress(str)
	case BTCChain:
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
		if err != nil {
			return NoAddress, err
		}
		var net *chaincfg.Params
		switch chainNetwork {
		case MockNet:
			net = &chaincfg.RegressionNetParams
		case TestNet:
			net = &chaincfg.TestNet3Params
		case MainNet, StageNet:
			net = &chaincfg.MainNetParams
		}
		addr, err := btcutil.NewAddressWitnessPubKeyHash(pk.Address().Bytes(), net)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to bech32 encode the address, err: %w", err)
		}
		return NewAddress(addr.String())
	case LTCChain:
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
		if err != nil {
			return NoAddress, err
		}
		var net *ltcchaincfg.Params
		switch chainNetwork {
		case MockNet:
			net = &ltcchaincfg.RegressionNetParams
		case TestNet:
			net = &ltcchaincfg.TestNet4Params
		case MainNet, StageNet:
			net = &ltcchaincfg.MainNetParams
		}
		addr, err := ltcutil.NewAddressWitnessPubKeyHash(pk.Address().Bytes(), net)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to bech32 encode the address, err: %w", err)
		}
		return NewAddress(addr.String())
	case DASHChain:
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
		if err != nil {
			return NoAddress, err
		}
		var net *dashchaincfg.Params
		switch chainNetwork {
		case MockNet:
			net = &dashchaincfg.RegressionNetParams
		case TestNet:
			net = &dashchaincfg.TestNet3Params
		case MainNet, StageNet:
			net = &dashchaincfg.MainNetParams
		}
		addr, err := dashutil.NewAddressPubKeyHash(pk.Address().Bytes(), net)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to encode the address, err: %w", err)
		}
		return NewAddress(addr.String())
	case DOGEChain:
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
		if err != nil {
			return NoAddress, err
		}
		var net *dogchaincfg.Params
		switch chainNetwork {
		case MockNet:
			net = &dogchaincfg.RegressionNetParams
		case TestNet:
			net = &dogchaincfg.TestNet3Params
		case MainNet, StageNet:
			net = &dogchaincfg.MainNetParams
		}
		addr, err := dogutil.NewAddressPubKeyHash(pk.Address().Bytes(), net)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to encode the address, err: %w", err)
		}
		return NewAddress(addr.String())
	case BCHChain:
		pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, string(pubKey))
		if err != nil {
			return NoAddress, err
		}
		var net *bchchaincfg.Params
		switch chainNetwork {
		case MockNet:
			net = &bchchaincfg.RegressionNetParams
		case TestNet:
			net = &bchchaincfg.TestNet3Params
		case MainNet, StageNet:
			net = &bchchaincfg.MainNetParams
		}
		addr, err := bchutil.NewAddressPubKeyHash(pk.Address().Bytes(), net)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to encode the address, err: %w", err)
		}
		return NewAddress(addr.String())
	}

	return NoAddress, nil
}

func (pubKey PubKey) GetThorAddress() (cosmos.AccAddress, error) {
	addr, err := pubKey.GetAddress(BASEChain)
	if err != nil {
		return nil, err
	}
	return cosmos.AccAddressFromBech32(addr.String())
}

// MarshalJSON to Marshals to JSON using Bech32
func (pubKey PubKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(pubKey.String())
}

// UnmarshalJSON to Unmarshal from JSON assuming Bech32 encoding
func (pubKey *PubKey) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	pk, err := NewPubKey(s)
	if err != nil {
		return err
	}
	*pubKey = pk
	return nil
}

func (pks PubKeys) Valid() error {
	for _, pk := range pks {
		if _, err := NewPubKey(pk.String()); err != nil {
			return err
		}
	}
	return nil
}

func (pks PubKeys) Contains(pk PubKey) bool {
	for _, p := range pks {
		if p.Equals(pk) {
			return true
		}
	}
	return false
}

// Equals check whether two pub keys are identical
func (pks PubKeys) Equals(newPks PubKeys) bool {
	if len(pks) != len(newPks) {
		return false
	}

	source := make(PubKeys, len(pks))
	dest := make(PubKeys, len(newPks))
	copy(source, pks)
	copy(dest, newPks)

	// sort both lists
	sort.Slice(source[:], func(i, j int) bool {
		return source[i].String() < source[j].String()
	})
	sort.Slice(dest[:], func(i, j int) bool {
		return dest[i].String() < dest[j].String()
	})
	for i := range source {
		if !source[i].Equals(dest[i]) {
			return false
		}
	}
	return true
}

// String implement stringer interface
func (pks PubKeys) String() string {
	strs := make([]string, len(pks))
	for i := range pks {
		strs[i] = pks[i].String()
	}
	return strings.Join(strs, ", ")
}

func (pks PubKeys) Strings() []string {
	allStrings := make([]string, len(pks))
	for i, pk := range pks {
		allStrings[i] = pk.String()
	}
	return allStrings
}

// ConvertAndEncode converts from a base64 encoded byte string to hex or base32 encoded byte string and then to bech32
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("encoding bech32 failed,%w", err)
	}
	return bech32.Encode(hrp, converted)
}

// NewPubKeySet create a new instance of PubKeySet , which contains two keys
func NewPubKeySet(secp256k1, ed25519 PubKey) PubKeySet {
	return PubKeySet{
		Secp256k1: secp256k1,
		Ed25519:   ed25519,
	}
}

// IsEmpty will determinate whether PubKeySet is an empty
func (pks PubKeySet) IsEmpty() bool {
	return pks.Secp256k1.IsEmpty() || pks.Ed25519.IsEmpty()
}

// Equals check whether two PubKeySet are the same
func (pks PubKeySet) Equals(pks1 PubKeySet) bool {
	return pks.Ed25519.Equals(pks1.Ed25519) && pks.Secp256k1.Equals(pks1.Secp256k1)
}

func (pks PubKeySet) Contains(pk PubKey) bool {
	return pks.Ed25519.Equals(pk) || pks.Secp256k1.Equals(pk)
}

// String implement fmt.Stinger
func (pks PubKeySet) String() string {
	return fmt.Sprintf(`
	secp256k1: %s
	ed25519: %s
`, pks.Secp256k1.String(), pks.Ed25519.String())
}

// GetAddress
func (pks PubKeySet) GetAddress(chain Chain) (Address, error) {
	switch chain.GetSigningAlgo() {
	case SigningAlgoSecp256k1:
		return pks.Secp256k1.GetAddress(chain)
	case SigningAlgoEd25519:
		return pks.Ed25519.GetAddress(chain)
	}
	return NoAddress, fmt.Errorf("unknown signing algorithm")
}
