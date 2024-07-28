//go:build testnet || mocknet
// +build testnet mocknet

package avaxtokens

import (
	_ "embed"
)

//go:embed avax_mocknet_V95.json
var AVAXTokenListRawV95 []byte
