package keeperv1

import (
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
)

type InvariantsSuite struct{}

var _ = Suite(&InvariantsSuite{})

func (s *InvariantsSuite) TestAsgardInvariant(c *C) {
	ctx, k := setupKeeperForTest(c)

	// empty the starting balance of asgard
	cacaoBal := k.GetRuneBalanceOfModule(ctx, AsgardName)
	coins := common.NewCoins(common.NewCoin(common.BaseAsset(), cacaoBal))
	c.Assert(k.SendFromModuleToModule(ctx, AsgardName, ReserveName, coins), IsNil)

	pool := NewPool()
	pool.Asset = common.BTCAsset
	pool.BalanceCacao = cosmos.NewUint(1000)
	pool.PendingInboundCacao = cosmos.NewUint(100)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	// savers pools are not included in expectations
	pool = NewPool()
	pool.Asset = common.BTCAsset.GetSyntheticAsset()
	pool.BalanceCacao = cosmos.NewUint(666)
	pool.PendingInboundCacao = cosmos.NewUint(777)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	swapMsg := MsgSwap{
		Tx: GetRandomTx(),
	}
	swapMsg.Tx.Coins = common.NewCoins(common.NewCoin(common.BaseAsset(), cosmos.NewUint(2000)))
	c.Assert(k.SetSwapQueueItem(ctx, swapMsg, 0), IsNil)

	// synth swaps are ignored
	swapMsg.Tx.Coins = common.NewCoins(common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(666)))
	c.Assert(k.SetSwapQueueItem(ctx, swapMsg, 1), IsNil)

	// layer1 swaps are ignored
	swapMsg.Tx.Coins = common.NewCoins(common.NewCoin(common.BTCAsset, cosmos.NewUint(777)))
	c.Assert(k.SetSwapQueueItem(ctx, swapMsg, 2), IsNil)

	invariant := AsgardInvariant(k)

	msg, broken := invariant(ctx)
	c.Assert(broken, Equals, true)
	c.Assert(len(msg), Equals, 2)
	c.Assert(msg[0], Equals, "insolvent: 666btc/btc")
	c.Assert(msg[1], Equals, "insolvent: 3100cacao")

	// send the expected amount to asgard
	expCoins := common.NewCoins(
		common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(666)),
		common.NewCoin(common.BaseAsset(), cosmos.NewUint(3100)),
	)
	for _, coin := range expCoins {
		c.Assert(k.MintToModule(ctx, ModuleName, coin), IsNil)
	}
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, AsgardName, expCoins), IsNil)

	msg, broken = invariant(ctx)
	c.Assert(broken, Equals, false)
	c.Assert(msg, IsNil)

	// send a little more to make asgard oversolvent
	extraCoins := common.NewCoins(common.NewCoin(common.BaseAsset(), cosmos.NewUint(1)))
	c.Assert(k.SendFromModuleToModule(ctx, ReserveName, AsgardName, extraCoins), IsNil)

	msg, broken = invariant(ctx)
	c.Assert(broken, Equals, true)
	c.Assert(len(msg), Equals, 1)
	c.Assert(msg[0], Equals, "oversolvent: 1cacao")
}

func (s *InvariantsSuite) TestNodeRewardsInvariant(c *C) {
	ctx, k := setupKeeperForTest(c)

	network := NewNetwork()
	network.BondRewardRune = cosmos.NewUint(2000)
	c.Assert(k.SetNetwork(ctx, network), IsNil)

	invariant := NodeRewardsInvariant(k)

	msg, broken := invariant(ctx)
	c.Assert(broken, Equals, true)
	c.Assert(len(msg), Equals, 1)
	c.Assert(msg[0], Equals, "insolvent: 2000cacao")

	expCacao := common.NewCoin(common.BaseAsset(), cosmos.NewUint(2000))
	c.Assert(k.MintToModule(ctx, ModuleName, expCacao), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, BondName, common.NewCoins(expCacao)), IsNil)

	msg, broken = invariant(ctx)
	c.Assert(broken, Equals, false)
	c.Assert(msg, IsNil)

	// send more to make bond oversolvent
	c.Assert(k.MintToModule(ctx, ModuleName, expCacao), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, BondName, common.NewCoins(expCacao)), IsNil)

	msg, broken = invariant(ctx)
	c.Assert(broken, Equals, true)
	c.Assert(len(msg), Equals, 1)
	c.Assert(msg[0], Equals, "oversolvent: 2000cacao")
}

func (s *InvariantsSuite) TestTHORChainInvariant(c *C) {
	ctx, k := setupKeeperForTest(c)

	invariant := MAYAChainInvariant(k)

	// should pass since it has no coins
	msg, broken := invariant(ctx)
	c.Assert(broken, Equals, false)
	c.Assert(msg, IsNil)

	// send some coins to make it oversolvent
	coins := common.NewCoins(common.NewCoin(common.BaseAsset(), cosmos.NewUint(1)))
	c.Assert(k.SendFromModuleToModule(ctx, AsgardName, ModuleName, coins), IsNil)

	msg, broken = invariant(ctx)
	c.Assert(broken, Equals, true)
	c.Assert(len(msg), Equals, 1)
	c.Assert(msg[0], Equals, "oversolvent: 1cacao")
}
