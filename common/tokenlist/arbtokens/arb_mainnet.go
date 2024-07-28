//go:build !testnet && !mocknet
// +build !testnet,!mocknet

package arbtokens

import (
	_ "embed"
)

//go:embed arb_mainnet_V109.json
var ARBTokenListRawV109 []byte

//go:embed arb_mainnet_latest.json
var ARBTokenListRawV110 []byte
