//go:build testnet || mocknet
// +build testnet mocknet

package ethtokens

import (
	_ "embed"
)

//go:embed eth_testnet_V93.json
var ETHTokenListRawV93 []byte

//go:embed eth_testnet_V95.json
var ETHTokenListRawV95 []byte

//go:embed eth_testnet_V106.json
var ETHTokenListRawV106 []byte

//go:embed eth_testnet_V109.json
var ETHTokenListRawV109 []byte

//go:embed eth_testnet_latest.json
var ETHTokenListRawV110 []byte
