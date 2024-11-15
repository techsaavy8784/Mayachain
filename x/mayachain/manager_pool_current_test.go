package mayachain

import (
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
)

type PoolMgrV108Suite struct{}

var _ = Suite(&PoolMgrV108Suite{})

func (s *PoolMgrV108Suite) TestEnableNextPool(c *C) {
	var err error
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	c.Assert(err, IsNil)
	pool := NewPool()
	pool.Asset = common.BNBAsset
	pool.Status = PoolAvailable
	pool.BalanceCacao = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	pool = NewPool()
	pool.Asset = common.BTCAsset // gas pool should be enabled by default
	pool.Status = PoolAvailable
	pool.BalanceCacao = cosmos.NewUint(50 * common.One)
	pool.BalanceAsset = cosmos.NewUint(50 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	ethAsset, err := common.NewAsset("BNB.ETH")
	c.Assert(err, IsNil)
	pool = NewPool()
	pool.Asset = ethAsset
	pool.Status = PoolStaged
	pool.BalanceCacao = cosmos.NewUint(40 * common.One)
	pool.BalanceAsset = cosmos.NewUint(40 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	xmrAsset, err := common.NewAsset("XMR.XMR")
	c.Assert(err, IsNil)
	pool = NewPool()
	pool.Asset = xmrAsset
	pool.Status = PoolStaged
	pool.BalanceCacao = cosmos.NewUint(40 * common.One)
	pool.BalanceAsset = cosmos.NewUint(0 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	// usdAsset
	usdAsset, err := common.NewAsset("BNB.TUSDB")
	c.Assert(err, IsNil)
	pool = NewPool()
	pool.Asset = usdAsset
	pool.Status = PoolStaged
	pool.BalanceCacao = cosmos.NewUint(140 * common.One)
	pool.BalanceAsset = cosmos.NewUint(0 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	poolMgr := newPoolMgrVCUR(k)

	// should enable BTC
	c.Assert(poolMgr.cyclePools(ctx, 100, 1, 0, mgr), IsNil)
	pool, err = k.GetPool(ctx, common.BTCAsset)
	c.Assert(err, IsNil)
	c.Check(pool.Status, Equals, PoolAvailable)

	// should enable ETH
	c.Assert(poolMgr.cyclePools(ctx, 100, 1, 0, mgr), IsNil)
	pool, err = k.GetPool(ctx, ethAsset)
	c.Assert(err, IsNil)
	c.Check(pool.Status, Equals, PoolAvailable)

	// should NOT enable XMR, since it has no assets
	c.Assert(poolMgr.cyclePools(ctx, 100, 1, 10*common.One, mgr), IsNil)
	pool, err = k.GetPool(ctx, xmrAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.IsEmpty(), Equals, false)
	c.Check(pool.Status, Equals, PoolStaged)
	c.Check(pool.BalanceCacao.Uint64(), Equals, uint64(30*common.One))
}

func (s *PoolMgrV108Suite) TestAbandonPool(c *C) {
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	usdAsset, err := common.NewAsset("BNB.TUSDB")
	c.Assert(err, IsNil)
	pool := NewPool()
	pool.Asset = usdAsset
	pool.Status = PoolStaged
	pool.BalanceCacao = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	vault := GetRandomVault()
	vault.Coins = common.Coins{
		common.NewCoin(usdAsset, cosmos.NewUint(100*common.One)),
		common.NewCoin(common.BNBAsset, cosmos.NewUint(100*common.One)),
	}
	c.Assert(k.SetVault(ctx, vault), IsNil)

	runeAddr := GetRandomBaseAddress()
	bnbAddr := GetRandomBNBAddress()
	lp := LiquidityProvider{
		Asset:        usdAsset,
		CacaoAddress: runeAddr,
		AssetAddress: bnbAddr,
		Units:        cosmos.ZeroUint(),
		PendingCacao: cosmos.ZeroUint(),
		PendingAsset: cosmos.ZeroUint(),
	}
	k.SetLiquidityProvider(ctx, lp)

	poolMgr := newPoolMgrVCUR(k)

	// cycle pools
	c.Assert(poolMgr.cyclePools(ctx, 100, 1, 100*common.One, mgr), IsNil)

	// check pool was deleted
	pool, err = k.GetPool(ctx, usdAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.BalanceCacao.IsZero(), Equals, true)
	c.Assert(pool.BalanceAsset.IsZero(), Equals, true)

	// check vault remove pool asset
	vault, err = k.GetVault(ctx, vault.PubKey)
	c.Assert(err, IsNil)
	c.Assert(vault.HasAsset(usdAsset), Equals, false)
	c.Assert(vault.CoinLength(), Equals, 1)

	// check that liquidity provider got removed
	count := 0
	iterator := k.GetLiquidityProviderIterator(ctx, usdAsset)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		count++
	}
	c.Assert(count, Equals, 0)
}

func (s *PoolMgrV108Suite) TestDemotePoolWithLowLiquidityFees(c *C) {
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	usdAsset, err := common.NewAsset("BNB.TUSDB")
	c.Assert(err, IsNil)
	pool := NewPool()
	pool.Asset = usdAsset
	pool.Status = PoolStaged
	pool.BalanceCacao = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	poolBNB := NewPool()
	poolBNB.Asset = common.BNBAsset
	poolBNB.Status = PoolAvailable
	poolBNB.BalanceCacao = cosmos.NewUint(100000 * common.One)
	poolBNB.BalanceAsset = cosmos.NewUint(100000 * common.One)
	c.Assert(k.SetPool(ctx, poolBNB), IsNil)

	bnbETH, err := common.NewAsset("BNB.ETH-1C9")
	c.Assert(err, IsNil)
	poolLoki := NewPool()
	poolLoki.Asset = bnbETH
	poolLoki.Status = PoolAvailable
	poolLoki.BalanceCacao = cosmos.NewUint(100000 * common.One)
	poolLoki.BalanceAsset = cosmos.NewUint(100000 * common.One)
	c.Assert(k.SetPool(ctx, poolLoki), IsNil)

	vault := GetRandomVault()
	vault.Coins = common.Coins{
		common.NewCoin(usdAsset, cosmos.NewUint(100*common.One)),
		common.NewCoin(common.BNBAsset, cosmos.NewUint(100*common.One)),
	}
	c.Assert(k.SetVault(ctx, vault), IsNil)

	runeAddr := GetRandomBaseAddress()
	bnbAddr := GetRandomBNBAddress()
	lp := LiquidityProvider{
		Asset:        usdAsset,
		CacaoAddress: runeAddr,
		AssetAddress: bnbAddr,
		Units:        cosmos.ZeroUint(),
		PendingCacao: cosmos.ZeroUint(),
		PendingAsset: cosmos.ZeroUint(),
	}
	k.SetLiquidityProvider(ctx, lp)
	k.SetMimir(ctx, constants.MinimumPoolLiquidityFee.String(), 100000000)

	poolMgr := newPoolMgrVCUR(k)

	// cycle pools
	c.Assert(poolMgr.cyclePools(ctx, 100, 1, 100*common.One, mgr), IsNil)

	// check pool was deleted
	pool, err = k.GetPool(ctx, usdAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.BalanceCacao.IsZero(), Equals, true)
	c.Assert(pool.BalanceAsset.IsZero(), Equals, true)

	// check vault remove pool asset
	vault, err = k.GetVault(ctx, vault.PubKey)
	c.Assert(err, IsNil)
	c.Assert(vault.HasAsset(usdAsset), Equals, false)
	c.Assert(vault.CoinLength(), Equals, 1)

	// check that liquidity provider got removed
	count := 0
	iterator := k.GetLiquidityProviderIterator(ctx, usdAsset)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		count++
	}
	c.Assert(count, Equals, 0)
	afterBNBPool, err := k.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(afterBNBPool.Status == PoolAvailable, Equals, true)
	afterBNBEth, err := k.GetPool(ctx, bnbETH)
	c.Assert(err, IsNil)
	c.Assert(afterBNBEth.Status == PoolStaged, Equals, true)
}

func (s *PoolMgrV108Suite) TestPoolMeetTradingVolumeCriteria(c *C) {
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	pm := newPoolMgrVCUR(k)

	asset := common.BTCAsset

	pool := Pool{
		Asset:        asset,
		BalanceAsset: cosmos.NewUint(1000 * common.One),
		BalanceCacao: cosmos.NewUint(1000 * common.One),
	}

	minFee := cosmos.ZeroUint()
	meets := pm.poolMeetTradingVolumeCriteria(ctx, mgr, pool, minFee)
	c.Assert(meets, Equals, true,
		Commentf("all pools should meet criteria if min fees is zero"))

	minFee = cosmos.NewUint(10 * common.One)
	meets = pm.poolMeetTradingVolumeCriteria(ctx, mgr, pool, minFee)
	c.Assert(meets, Equals, false,
		Commentf("pool with no fees collected should not meet criteria"))

	err := k.AddToLiquidityFees(ctx, pool.Asset, minFee.Add(cosmos.NewUint(1)))
	c.Assert(err, IsNil)
	cur, err := k.GetRollingPoolLiquidityFee(ctx, pool.Asset)
	c.Assert(err, IsNil)
	c.Assert(cur, Equals, minFee.Add(cosmos.NewUint(1)).Uint64())

	meets = pm.poolMeetTradingVolumeCriteria(ctx, mgr, pool, minFee)
	c.Assert(meets, Equals, true,
		Commentf("pool should meet min fee criteria"))
}

func (s *PoolMgrV108Suite) TestRemoveAssetFromVault(c *C) {
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	pm := newPoolMgrVCUR(k)

	asset := common.BTCAsset

	v0 := GetRandomVault()
	v0.Coins = common.Coins{
		common.NewCoin(common.ETHAsset, cosmos.NewUint(1*common.One)),
		common.NewCoin(common.BNBAsset, cosmos.NewUint(10*common.One)),
	}
	c.Assert(k.SetVault(ctx, v0), IsNil)
	c.Assert(v0.HasAsset(asset), Equals, false,
		Commentf("vault0 should not have asset balance"))

	v1 := GetRandomVault()
	v1.Coins = common.Coins{
		common.NewCoin(asset, cosmos.NewUint(100*common.One)),
		common.NewCoin(common.ETHAsset, cosmos.NewUint(1000*common.One)),
	}
	c.Assert(k.SetVault(ctx, v1), IsNil)
	c.Assert(v1.HasAsset(asset), Equals, true,
		Commentf("vault1 should have asset balance"))

	pm.removeAssetFromVault(ctx, asset, mgr)

	v0, err := k.GetVault(ctx, v0.PubKey)
	c.Assert(err, IsNil)
	c.Assert(v0.HasAsset(asset), Equals, false,
		Commentf("vault0 should still not have asset balance"))

	v1, err = k.GetVault(ctx, v1.PubKey)
	c.Assert(err, IsNil)
	c.Assert(v1.HasAsset(asset), Equals, false,
		Commentf("vault1 should no longer have asset"))
}

func (s *PoolMgrV108Suite) TestRemoveLiquidityProviders(c *C) {
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	pm := newPoolMgrVCUR(k)

	countLiquidityProviders := func(ctx cosmos.Context, k keeper.Keeper, asset common.Asset) int {
		count := 0
		iterator := k.GetLiquidityProviderIterator(ctx, asset)
		defer iterator.Close()
		for ; iterator.Valid(); iterator.Next() {
			count++
		}
		return count
	}

	asset := common.BTCAsset

	c.Assert(countLiquidityProviders(ctx, k, asset), Equals, 0,
		Commentf("should not have lps before adding"))

	k.SetLiquidityProvider(ctx, LiquidityProvider{
		Asset:        asset,
		CacaoAddress: GetRandomBaseAddress(),
	})
	k.SetLiquidityProvider(ctx, LiquidityProvider{
		Asset:        asset,
		CacaoAddress: GetRandomBaseAddress(),
	})
	c.Assert(countLiquidityProviders(ctx, k, asset), Equals, 2,
		Commentf("should have 2 lps after adding"))

	pm.removeLiquidityProviders(ctx, asset, mgr)

	c.Assert(countLiquidityProviders(ctx, k, asset), Equals, 0,
		Commentf("should have 0 lps after removing"))
}
