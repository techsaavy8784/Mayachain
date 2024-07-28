//go:build !testnet && !mocknet && !stagenet
// +build !testnet,!mocknet,!stagenet

package aggregators

import (
	"gitlab.com/mayachain/mayanode/common"
)

func DexAggregatorsV110() []Aggregator {
	return []Aggregator{
		// RangoDiamond Ethereum
		{common.ETHChain, `0x69460570c93f9DE5E2edbC3052bf10125f0Ca22d`, 400_000},
		// RangoThorchainOutputAggUniV3_COMPACT_Fee500
		{common.ETHChain, `0x70F75937546fB26c6FD3956eBBfb285f41526186`, 400_000},
		// RangoThorchainOutputAggUniV3_COMPACT_Fee3000
		{common.ETHChain, `0xd1687354CBA0e56facd0c44eD0F69D97F5734Dc1`, 400_000},
		// RangoThorchainOutputAggUniV3_COMPACT_Fee10000
		{common.ETHChain, `0xaFa4cBA6db85515f66E3ed7d6784e8cf5b689E2D`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_SUSHI
		{common.ETHChain, `0x0964347B0019eb227c901220ce7d66BB01479220`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_UNI
		{common.ETHChain, `0x6f281993AB68216F8898c593C4578C8a4a76F063`, 400_000},
		// TSAggregatorPancakeSwap Ethereum V2
		{common.ETHChain, `0x35CF22003c90126528fbe95b21bB3ADB2ca8c53D`, 400_000},
		// TSAggregatorGeneric
		{common.ETHChain, `0xd31f7e39afECEc4855fecc51b693F9A0Cec49fd2`, 400_000},
		// RangoThorchainOutputAggUniV2
		{common.ETHChain, `0x2a7813412b8da8d18Ce56FE763B9eb264D8e28a8`, 400_000},
		// RangoThorchainOutputAggUniV3
		{common.ETHChain, `0xbB8De86F3b041B3C084431dcf3159fE4827c5F0D`, 400_000},
		// TSAggregatorUniswapV2 - short notation
		{common.ETHChain, `0x86904eb2b3c743400d03f929f2246efa80b91215`, 400_000},
		// TSAggregatorSushiswap - short notation
		{common.ETHChain, `0xbf365e79aa44a2164da135100c57fdb6635ae870`, 400_000},
		// TSAggregatorUniswapV3 100 - short notation
		{common.ETHChain, `0xbd68cbe6c247e2c3a0e36b8f0e24964914f26ee8`, 400_000},
		// TSAggregatorUniswapV3 500 - short notation
		{common.ETHChain, `0xe4ddca21881bac219af7f217703db0475d2a9f02`, 400_000},
		// TSAggregatorUniswapV3 3000 - short notation
		{common.ETHChain, `0x11733abf0cdb43298f7e949c930188451a9a9ef2`, 400_000},
		// TSAggregatorUniswapV3 10000 - short notation
		{common.ETHChain, `0xb33874810e5395eb49d8bd7e912631db115d5a03`, 400_000},
		// TSLedgerAdapter
		{common.ETHChain, `0xB81C7C2D2d078205D7FA515DDB2dEA3d896F4016`, 500_000},
		// TSAggregatorUniswapV2 Ethereum gen2 V2.5 - tax tokens
		{common.ETHChain, `0x0fA226e8BCf45ec2f3c3163D2d7ba0d2aAD2eBcF`, 400_000},
		// TSWrapperLedger_V1
		{common.ETHChain, `0xE4e8313AbbADc8E18543EC9528f67Fde2e44D3D6`, 600_000},
		// TSWrapperTCRouterV4_V1
		{common.ETHChain, `0x94B7F2145C328DaB2EC56aB982CaB95F00941aE7`, 400_000},
		// LiFi - ETH
		{common.ETHChain, `0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE`, 800_000},
		// LiFi Staging - ETH
		{common.ETHChain, `0xbEbCDb5093B47Cd7add8211E4c77B6826aF7bc5F`, 800_000},
		// LiFi - ARB
		{common.ARBChain, `0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE`, 800_000},
		// LiFi Staging - ARB
		{common.ARBChain, `0xbEbCDb5093B47Cd7add8211E4c77B6826aF7bc5F`, 800_000},
		// OKXRouter - ETH
		{common.ETHChain, `0xFc99f58A8974A4bc36e60E2d490Bb8D72899ee9f`, 800_000},
		// OKXRouter - ARB
		{common.ARBChain, `0xFc99f58A8974A4bc36e60E2d490Bb8D72899ee9f`, 800_000},
		// PlexusDiamond Ethereum
		{common.ETHChain, `0x7E1B771bBA8f017E0988F068C320b8852F4A3265`, 600_000},
		// PlexusDiamond Arbitrum
		{common.ARBChain, `0x52cdB00b69f11C4cA932fA9108C6BFdD65F20d62`, 600_000},
		// TSAggregatorStargate Ethereum gen2 V1
		{common.ETHChain, `0x1204b5Bf0D6d48E718B1d9753A4166A7258B8432`, 800_000},
		// TSAggregatorStargate Ethereum gen2 V2
		{common.ETHChain, `0x48f68ff093b3b3A80D2FC97488EaD97E16b86283`, 800_000},
		// LayerZero Executor Ethereum
		{common.ETHChain, `0xe93685f3bBA03016F02bD1828BaDD6195988D950`, 800_000},
		// LayerZero Executor Arbitrum
		{common.ARBChain, `0xe93685f3bBA03016F02bD1828BaDD6195988D950`, 800_000},
		// SquidRouter MultiCall Ethereum
		{common.ETHChain, `0x4fd39C9E151e50580779bd04B1f7eCc310079fd3`, 800_000},
		// TSAggregatorChainflip V1
		{common.ETHChain, `0xd917C32Cfa50b0FcA692A208118AfE4d1d06f1f1`, 800_000},
		// TSAggregatorWoofi - ARB
		{common.ARBChain, `0xcD6f8B17d19077f7f5a8Fa30C9b645B689626ceb`, 400_000},
		// Corn - Uniswap 500
		{common.ETHChain, `0x9E9b95b22bA7A1A875D3f93C2Ec54C015C14Aa9c`, 800_000},
		// Unizen (Prod) - ETH
		{common.ETHChain, `0xd3f64BAa732061F8B3626ee44bab354f854877AC`, 600_000},
		// Unizen (Prod) - ARB
		{common.ARBChain, `0x1C7F7e0258c81CF41bcEa31ea4bB5191914Bf7D7`, 600_000},
		// Unizen (Dev) - ETH
		{common.ETHChain, `0x4432d5179E29FBBd1a91F5c6949a220999790AAC`, 600_000},
	}
}