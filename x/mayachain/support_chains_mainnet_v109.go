//go:build !testnet && !mocknet && !stagenet
// +build !testnet,!mocknet,!stagenet

// go:build !testnet && !mocknet && !stagenet
package mayachain

import "gitlab.com/mayachain/mayanode/common"

// Supported chains for mainnet
var SUPPORT_CHAINS_V109 = common.Chains{
	common.BASEChain,
	common.BTCChain,
	common.DASHChain,
	common.ETHChain,
	common.KUJIChain,
	common.THORChain,
	common.ARBChain,
}
