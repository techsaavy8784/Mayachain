package mayachain

import (
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/keeper"
)

type NetworkManagerV108TestSuite struct{}

var _ = Suite(&NetworkManagerV108TestSuite{})

func (s *NetworkManagerV108TestSuite) SetUpSuite(c *C) {
	SetupConfigForTest()
}

func (s *NetworkManagerV108TestSuite) TestRagnarokChain(c *C) {
	ctx, _ := setupKeeperForTest(c)
	ctx = ctx.WithBlockHeight(100000)

	activeVault := GetRandomVault()
	activeVault.StatusSince = ctx.BlockHeight() - 10
	activeVault.Coins = common.Coins{
		common.NewCoin(common.BNBAsset, cosmos.NewUint(100*common.One)),
	}
	retireVault := GetRandomVault()
	retireVault.Chains = common.Chains{common.BNBChain, common.BTCChain}.Strings()
	yggVault := GetRandomVault()
	yggVault.Type = YggdrasilVault
	yggVault.Coins = common.Coins{
		common.NewCoin(common.BTCAsset, cosmos.NewUint(3*common.One)),
		common.NewCoin(common.BaseAsset(), cosmos.NewUint(300*common.One)),
	}

	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceCacao = cosmos.NewUint(1000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)

	bnbPool := NewPool()
	bnbPool.Asset = common.BNBAsset
	bnbPool.BalanceCacao = cosmos.NewUint(1000 * common.One)
	bnbPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	bnbPool.LPUnits = cosmos.NewUint(1600)

	addr := GetRandomBaseAddress()
	lps := LiquidityProviders{
		{
			CacaoAddress:              addr,
			AssetAddress:              GetRandomBTCAddress(),
			LastAddHeight:             5,
			Units:                     btcPool.LPUnits.QuoUint64(2),
			PendingCacao:              cosmos.ZeroUint(),
			PendingAsset:              cosmos.ZeroUint(),
			AssetDepositValue:         cosmos.ZeroUint(),
			CacaoDepositValue:         cosmos.ZeroUint(),
			WithdrawCounter:           cosmos.ZeroUint(),
			LastWithdrawCounterHeight: 0,
		},
		{
			CacaoAddress:              GetRandomBaseAddress(),
			AssetAddress:              GetRandomBTCAddress(),
			LastAddHeight:             10,
			Units:                     btcPool.LPUnits.QuoUint64(2),
			PendingCacao:              cosmos.ZeroUint(),
			PendingAsset:              cosmos.ZeroUint(),
			AssetDepositValue:         cosmos.ZeroUint(),
			CacaoDepositValue:         cosmos.ZeroUint(),
			WithdrawCounter:           cosmos.ZeroUint(),
			LastWithdrawCounterHeight: 0,
		},
	}

	keeper := &TestRagnarokChainKeeper{
		na:          GetRandomValidatorNode(NodeActive),
		activeVault: activeVault,
		retireVault: retireVault,
		yggVault:    yggVault,
		pools:       Pools{bnbPool, btcPool},
		lps:         lps,
	}

	mgr := NewDummyMgrWithKeeper(keeper)

	networkMgr := newNetworkMgrV108(keeper, mgr.TxOutStore(), mgr.EventMgr())

	// the first round should just recall yggdrasil fund
	err := networkMgr.manageChains(ctx, mgr)
	c.Assert(err, IsNil)
	c.Check(keeper.pools[1].Asset.Equals(common.BTCAsset), Equals, true)
	c.Check(keeper.pools[1].LPUnits.IsZero(), Equals, false, Commentf("%d\n", keeper.pools[1].LPUnits.Uint64()))
	c.Check(keeper.pools[0].LPUnits.Equal(cosmos.NewUint(1600)), Equals, true)
	for _, skr := range keeper.lps {
		c.Check(skr.Units.IsZero(), Equals, false)
	}

	// the first round should just recall yggdrasil fund
	ctx = ctx.WithBlockHeight(200000)
	err = networkMgr.manageChains(ctx, mgr)
	c.Assert(err, IsNil)
	c.Check(keeper.pools[1].Asset.Equals(common.BTCAsset), Equals, true)
	c.Check(keeper.pools[1].LPUnits.IsZero(), Equals, true, Commentf("%d\n", keeper.pools[1].LPUnits.Uint64()))
	c.Check(keeper.pools[0].LPUnits.Equal(cosmos.NewUint(1600)), Equals, true)
	for _, skr := range keeper.lps {
		c.Check(skr.Units.IsZero(), Equals, true)
	}
	// ensure we have requested for ygg funds to be returned
	txOutStore := mgr.TxOutStore()
	c.Assert(err, IsNil)
	items, err := txOutStore.GetOutboundItems(ctx)
	c.Assert(err, IsNil)

	// 1 ygg return + 4 withdrawals
	c.Check(items, HasLen, 3, Commentf("Len %d", items))
	c.Check(items[0].Memo, Equals, NewYggdrasilReturn(100000).String())
	c.Check(items[0].Chain.Equals(common.BTCChain), Equals, true)

	ctx, mgr1 := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr1.Keeper())
	mgr.K = helper
	networkMgr1 := newNetworkMgrV108(helper, mgr1.TxOutStore(), mgr1.EventMgr())
	// fail to get active nodes should error out
	helper.failToListActiveAccounts = true
	c.Assert(networkMgr1.ragnarokChain(ctx, common.BNBChain, 1, mgr), NotNil)
	helper.failToListActiveAccounts = false

	// no active nodes , should error
	c.Assert(networkMgr1.ragnarokChain(ctx, common.BNBChain, 1, mgr), NotNil)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)

	// fail to get pools should error out
	helper.failGetPools = true
	c.Assert(networkMgr1.ragnarokChain(ctx, common.BNBChain, 1, mgr), NotNil)
	helper.failGetPools = false
}

func (s *NetworkManagerV108TestSuite) TestUpdateNetwork(c *C) {
	ctx, mgr := setupManagerForTest(c)
	ver := GetCurrentVersion()
	constAccessor := constants.GetConstantValues(ver)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())

	// fail to get Network should return error
	helper.failGetNetwork = true
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), NotNil)
	helper.failGetNetwork = false

	// TotalReserve is zero , should not doing anything
	vd := NewNetwork()
	err := mgr.Keeper().SetNetwork(ctx, vd)
	c.Assert(err, IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	p := NewPool()
	p.Asset = common.BNBAsset
	p.BalanceCacao = cosmos.NewUint(common.One * 100)
	p.BalanceAsset = cosmos.NewUint(common.One * 100)
	p.Status = PoolAvailable
	c.Assert(helper.SetPool(ctx, p), IsNil)
	// no active node , thus no bond
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	// with liquidity fee , and bonds
	c.Assert(helper.Keeper.AddToLiquidityFees(ctx, common.BNBAsset, cosmos.NewUint(50*common.One)), IsNil)

	reserveBalanceBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, ReserveName)
	mayaBalanceBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, MayaFund)
	bondBalanceBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, BondName)

	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	reserveBalanceAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, ReserveName)
	mayaBalanceAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, MayaFund)
	bondBalanceAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, BondName)

	c.Check(reserveBalanceAfter.Sub(reserveBalanceBefore).Uint64(), Equals, uint64(5*common.One))
	c.Check(mayaBalanceAfter.Sub(mayaBalanceBefore).Uint64(), Equals, uint64(5*common.One))
	c.Assert(bondBalanceAfter.Sub(bondBalanceBefore).Uint64(), Equals, uint64(40*common.One))

	// add bond
	na := GetRandomValidatorNode(NodeActive)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, na), IsNil)
	SetupLiquidityBondForTest(c, ctx, mgr.Keeper(), common.BTCAsset, na.BondAddress, na, cosmos.NewUint(100*common.One))
	na1 := GetRandomValidatorNode(NodeActive)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, na1), IsNil)
	SetupLiquidityBondForTest(c, ctx, mgr.Keeper(), common.BTCAsset, na1.BondAddress, na1, cosmos.NewUint(100*common.One))
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	// fail to get total liquidity fee should result an error
	helper.failGetTotalLiquidityFee = true
	if common.BaseAsset().Equals(common.BaseNative) {
		FundModule(c, ctx, helper, ReserveName, 100)
	}
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), NotNil)
	helper.failGetTotalLiquidityFee = false

	helper.failToListActiveAccounts = true
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), NotNil)
}

func (s *NetworkManagerV108TestSuite) TestCalcBlockRewards(c *C) {
	ctx, k := setupKeeperForTest(c)
	mgr := NewDummyMgrWithKeeper(k)
	networkMgr := newNetworkMgrV108(k, mgr.TxOutStore(), mgr.EventMgr())

	bondR, bondShare := networkMgr.calcBlockRewards(ctx, cosmos.NewUint(1000*common.One), cosmos.NewUint(751*common.One), cosmos.NewUint(100*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(99_60000000), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(9960), Commentf("%d", bondShare.Uint64()))

	// bonded should always be less or equal to total liquidity since bond is liquidity
	bondR, bondShare = networkMgr.calcBlockRewards(ctx, cosmos.NewUint(1000*common.One), cosmos.NewUint(2000*common.One), cosmos.NewUint(1000*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(1000*common.One), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(10_000), Commentf("%d", bondShare.Uint64()))

	// no fees
	bondR, bondShare = networkMgr.calcBlockRewards(ctx, cosmos.NewUint(1000*common.One), cosmos.NewUint(900*common.One), cosmos.NewUint(0*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(0), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(4000), Commentf("%d", bondShare.Uint64()))

	// really over bonded pays correctly, but practically 0%
	bondR, bondShare = networkMgr.calcBlockRewards(ctx, cosmos.NewUint(1000*common.One), cosmos.NewUint(999_99999999), cosmos.NewUint(1000*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(4), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(0), Commentf("%d", bondShare.Uint64()))

	bondR, bondShare = networkMgr.calcBlockRewards(ctx, cosmos.NewUint(2000*common.One), cosmos.NewUint(1000*common.One), cosmos.NewUint(1000*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(1000*common.One), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(10_000), Commentf("%d", bondShare.Uint64()))

	// IncentiveCurveControl mimir set to 9000 (90%)
	networkMgr.k.SetMimir(ctx, "IncentiveCurveControl", 9000)
	bondR, bondShare = networkMgr.calcBlockRewards(ctx, cosmos.NewUint(1000*common.One), cosmos.NewUint(900*common.One), cosmos.NewUint(100*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(90*common.One), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(9000), Commentf("%d", bondShare.Uint64()))

	// IncentiveCurveControl mimir set to 10001 (out of range)
	// Should get 40% of the fees since 4(1-0.9) = 0.4
	networkMgr.k.SetMimir(ctx, "IncentiveCurveControl", 10001)
	bondR, bondShare = networkMgr.calcBlockRewards(ctx, cosmos.NewUint(1000*common.One), cosmos.NewUint(900*common.One), cosmos.NewUint(100*common.One))
	c.Check(bondR.Uint64(), Equals, uint64(40*common.One), Commentf("%d", bondR.Uint64()))
	c.Check(bondShare.Uint64(), Equals, uint64(4000), Commentf("%d", bondShare.Uint64()))
}

func (s *NetworkManagerV108TestSuite) TestCalcPoolDeficit(c *C) {
	pool1Fees := cosmos.NewUint(1000)
	pool2Fees := cosmos.NewUint(3000)
	totalFees := cosmos.NewUint(4000)

	mgr := NewDummyMgr()
	networkMgr := newNetworkMgrV108(keeper.KVStoreDummy{}, mgr.TxOutStore(), mgr.EventMgr())

	lpDeficit := cosmos.NewUint(1120)
	amt1 := networkMgr.calcPoolDeficit(lpDeficit, totalFees, pool1Fees)
	amt2 := networkMgr.calcPoolDeficit(lpDeficit, totalFees, pool2Fees)

	c.Check(amt1.Equal(cosmos.NewUint(280)), Equals, true, Commentf("%d", amt1.Uint64()))
	c.Check(amt2.Equal(cosmos.NewUint(840)), Equals, true, Commentf("%d", amt2.Uint64()))
}

func (*NetworkManagerV108TestSuite) TestProcessGenesisSetup(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	ctx = ctx.WithBlockHeight(1)
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	// no active account
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)

	nodeAccount := GetRandomValidatorNode(NodeActive)
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, nodeAccount), IsNil)
	c.Assert(networkMgr.EndBlock(ctx, mgr), IsNil)
	// make sure asgard vault get created
	vaults, err := mgr.Keeper().GetAsgardVaults(ctx)
	c.Assert(err, IsNil)
	c.Assert(vaults, HasLen, 1)

	// fail to get asgard vaults should return an error
	helper.failToGetAsgardVaults = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failToGetAsgardVaults = false

	// vault already exist , it should not do anything , and should not error
	c.Assert(networkMgr.EndBlock(ctx, mgr), IsNil)

	ctx, mgr = setupManagerForTest(c)
	helper = NewVaultGenesisSetupTestHelper(mgr.Keeper())
	ctx = ctx.WithBlockHeight(1)
	mgr.K = helper
	networkMgr = newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	helper.failToListActiveAccounts = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failToListActiveAccounts = false

	helper.failToSetVault = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failToSetVault = false

	helper.failGetRetiringAsgardVault = true
	ctx = ctx.WithBlockHeight(1024)
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failGetRetiringAsgardVault = false

	helper.failGetActiveAsgardVault = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failGetActiveAsgardVault = false
}

func (*NetworkManagerV108TestSuite) TestGetTotalActiveBond(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	helper.failToListActiveAccounts = true
	bond, err := networkMgr.getTotalActiveBond(ctx)
	c.Assert(err, NotNil)
	c.Assert(bond.Equal(cosmos.ZeroUint()), Equals, true)
	helper.failToListActiveAccounts = false
	na := GetRandomValidatorNode(NodeActive)
	bp := NewBondProviders(na.NodeAddress)
	acc, err := na.BondAddress.AccAddress()
	c.Assert(err, IsNil)
	bp.Providers = append(bp.Providers, NewBondProvider(acc))
	bp.Providers[0].Bonded = true
	SetupLiquidityBondForTest(c, ctx, helper.Keeper, common.BNBAsset, na.BondAddress, na, cosmos.NewUint(100*common.One))
	c.Assert(helper.Keeper.SetBondProviders(ctx, bp), IsNil)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, na), IsNil)
	bond, err = networkMgr.getTotalActiveBond(ctx)
	c.Assert(err, IsNil)
	c.Assert(bond.Uint64() > 0, Equals, true)
}

func (*NetworkManagerV108TestSuite) TestGetTotalLiquidityRune(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	p := NewPool()
	p.Asset = common.BNBAsset
	p.BalanceCacao = cosmos.NewUint(common.One * 100)
	p.BalanceAsset = cosmos.NewUint(common.One * 100)
	p.Status = PoolAvailable
	c.Assert(helper.SetPool(ctx, p), IsNil)
	pools, totalLiquidity, err := networkMgr.getTotalProvidedLiquidityRune(ctx)
	c.Assert(err, IsNil)
	c.Assert(pools, HasLen, 1)
	c.Assert(totalLiquidity.Equal(p.BalanceCacao), Equals, true)
}

func (*NetworkManagerV108TestSuite) TestPayPoolRewards(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	p := NewPool()
	p.Asset = common.BNBAsset
	p.BalanceCacao = cosmos.NewUint(common.One * 100)
	p.BalanceAsset = cosmos.NewUint(common.One * 100)
	p.Status = PoolAvailable
	c.Assert(helper.SetPool(ctx, p), IsNil)
	c.Assert(networkMgr.payPoolRewards(ctx, []cosmos.Uint{cosmos.NewUint(100 * common.One)}, Pools{p}), IsNil)
	helper.failToSetPool = true
	c.Assert(networkMgr.payPoolRewards(ctx, []cosmos.Uint{cosmos.NewUint(100 * common.One)}, Pools{p}), NotNil)
}

func (*NetworkManagerV108TestSuite) TestFindChainsToRetire(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	// fail to get active asgard vault
	helper.failGetActiveAsgardVault = true
	chains, err := networkMgr.findChainsToRetire(ctx)
	c.Assert(err, NotNil)
	c.Assert(chains, HasLen, 0)
	helper.failGetActiveAsgardVault = false

	// fail to get retire asgard vault
	helper.failGetRetiringAsgardVault = true
	chains, err = networkMgr.findChainsToRetire(ctx)
	c.Assert(err, NotNil)
	c.Assert(chains, HasLen, 0)
	helper.failGetRetiringAsgardVault = false
}

func (*NetworkManagerV108TestSuite) TestRecallChainFunds(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())
	helper.failToListActiveAccounts = true
	c.Assert(networkMgr.RecallChainFunds(ctx, common.BNBChain, mgr, common.PubKeys{}), NotNil)
	helper.failToListActiveAccounts = false

	helper.failGetActiveAsgardVault = true
	c.Assert(networkMgr.RecallChainFunds(ctx, common.BNBChain, mgr, common.PubKeys{}), NotNil)
	helper.failGetActiveAsgardVault = false
}

func (s *NetworkManagerV108TestSuite) TestRecoverPoolDeficit(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV108(helper, mgr.TxOutStore(), mgr.EventMgr())

	pools := Pools{
		Pool{
			Asset:        common.BNBAsset,
			BalanceCacao: cosmos.NewUint(common.One * 2000),
			BalanceAsset: cosmos.NewUint(common.One * 2000),
			Status:       PoolAvailable,
		},
	}
	c.Assert(helper.Keeper.SetPool(ctx, pools[0]), IsNil)

	totalLiquidityFees := cosmos.NewUint(50 * common.One)
	c.Assert(helper.Keeper.AddToLiquidityFees(ctx, common.BNBAsset, totalLiquidityFees), IsNil)

	lpDeficit := cosmos.NewUint(totalLiquidityFees.Uint64())

	bondBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, BondName)
	asgardBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, AsgardName)
	reserveBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, ReserveName)

	poolAmts, err := networkMgr.deductPoolRewardDeficit(ctx, pools, totalLiquidityFees, lpDeficit)
	c.Assert(err, IsNil)
	c.Assert(len(poolAmts), Equals, 1)

	bondAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, BondName)
	asgardAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, AsgardName)
	reserveAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, ReserveName)

	// bond module is not touched
	c.Assert(bondAfter.String(), Equals, bondBefore.String())

	// deficit moves from asgard to reserve
	c.Assert(asgardAfter.String(), Equals, asgardBefore.Sub(lpDeficit).String())
	c.Assert(reserveAfter.String(), Equals, reserveBefore.Add(lpDeficit).String())

	// deficit rune is deducted from the pool record
	pool, err := helper.Keeper.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.BalanceCacao.String(), Equals, pools[0].BalanceCacao.Sub(lpDeficit).String())
}

func (s *NetworkManagerV108TestSuite) TestSaverYield(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	net := newNetworkMgrV108(mgr.Keeper(), mgr.TxOutStore(), mgr.EventMgr())
	mgr.Keeper().SetMimir(ctx, constants.SynthYieldCycle.String(), 5_000)

	// mint synths
	coin := common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(10*common.One))
	c.Assert(mgr.Keeper().MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	spool := NewPool()
	spool.Asset = common.BTCAsset.GetSyntheticAsset()
	spool.BalanceAsset = coin.Amount
	spool.LPUnits = cosmos.NewUint(100)
	c.Assert(mgr.Keeper().SetPool(ctx, spool), IsNil)

	// first pool
	pool := NewPool()
	pool.Asset = common.BTCAsset
	pool.BalanceCacao = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.LPUnits = cosmos.NewUint(100)
	pool.CalcUnits(mgr.GetVersion(), coin.Amount)
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	c.Assert(net.paySaverYield(ctx, common.BTCAsset, cosmos.NewUint(50*common.One)), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset)
	c.Assert(err, IsNil)
	c.Assert(spool.BalanceAsset.Uint64(), Equals, uint64(1102400000), Commentf("%d", spool.BalanceAsset.Uint64()))
}

func (s *NetworkManagerV108TestSuite) TestRagnarokPool(c *C) {
	ctx, k := setupKeeperForTest(c)
	ctx = ctx.WithBlockHeight(100000)
	na := GetRandomValidatorNode(NodeActive)
	bp := NewBondProviders(na.NodeAddress)
	acc, err := na.BondAddress.AccAddress()
	c.Assert(err, IsNil)
	bp.Providers = append(bp.Providers, NewBondProvider(acc))
	bp.Providers[0].Bonded = true
	c.Assert(k.SetNodeAccount(ctx, na), IsNil)
	c.Assert(k.SetBondProviders(ctx, bp), IsNil)
	activeVault := GetRandomVault()
	activeVault.StatusSince = ctx.BlockHeight() - 10
	activeVault.Coins = common.Coins{
		common.NewCoin(common.BNBAsset, cosmos.NewUint(100*common.One)),
	}
	c.Assert(k.SetVault(ctx, activeVault), IsNil)
	retireVault := GetRandomVault()
	retireVault.Chains = common.Chains{common.BNBChain, common.BTCChain}.Strings()
	yggVault := GetRandomVault()
	yggVault.PubKey = na.PubKeySet.Secp256k1
	yggVault.Type = YggdrasilVault
	yggVault.Coins = common.Coins{
		common.NewCoin(common.BTCAsset, cosmos.NewUint(3*common.One)),
	}
	c.Assert(k.SetVault(ctx, yggVault), IsNil)
	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceCacao = cosmos.NewUint(1000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)
	btcPool.Status = PoolAvailable
	c.Assert(k.SetPool(ctx, btcPool), IsNil)

	// Add liquidity for the node
	SetupLiquidityBondForTest(c, ctx, k, common.BNBAsset, common.Address(na.NodeAddress.String()), na, cosmos.NewUint(100*common.One))

	bnbPool := NewPool()
	bnbPool.Asset = common.BNBAsset
	bnbPool.BalanceCacao = cosmos.NewUint(1000 * common.One)
	bnbPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	bnbPool.LPUnits = cosmos.NewUint(1600)
	bnbPool.Status = PoolAvailable
	c.Assert(k.SetPool(ctx, bnbPool), IsNil)
	addr := GetRandomBaseAddress()
	lps := LiquidityProviders{
		{
			Asset:             common.BTCAsset,
			CacaoAddress:      addr,
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     5,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingCacao:      cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			CacaoDepositValue: cosmos.ZeroUint(),
		},
		{
			Asset:             common.BTCAsset,
			CacaoAddress:      GetRandomBaseAddress(),
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     10,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingCacao:      cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			CacaoDepositValue: cosmos.ZeroUint(),
		},
	}
	k.SetLiquidityProvider(ctx, lps[0])
	k.SetLiquidityProvider(ctx, lps[1])
	mgr := NewDummyMgrWithKeeper(k)
	networkMgr := newNetworkMgrV108(k, mgr.TxOutStore(), mgr.EventMgr())

	ctx = ctx.WithBlockHeight(1)
	// block height not correct , doesn't take any actions
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	var tempPool Pool
	for _, a := range []common.Asset{common.BTCAsset, common.BNBAsset} {
		tempPool, err = k.GetPool(ctx, a)
		c.Assert(err, IsNil)
		c.Assert(tempPool.Status, Equals, PoolAvailable)
	}
	interval := mgr.GetConstants().GetInt64Value(constants.FundMigrationInterval)
	// mimir didn't set , it should not take any actions
	ctx = ctx.WithBlockHeight(interval * 5)
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)

	// happy path
	networkMgr.k.SetMimir(ctx, "RAGNAROK-BTC-BTC", 1)
	// first round , it should recall yggdrasil
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	items, _ := mgr.txOutStore.GetOutboundItems(ctx)
	c.Assert(items, HasLen, 1)
	c.Assert(items[0].Memo, Equals, "YGGDRASIL-:200")

	// second round, ragnarok
	ctx = ctx.WithBlockHeight(interval * 6)
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	items, _ = mgr.txOutStore.GetOutboundItems(ctx)
	c.Assert(items, HasLen, 3)

	tempPool, err = k.GetPool(ctx, common.BTCAsset)
	c.Assert(err, IsNil)
	c.Assert(tempPool.Status, Equals, PoolSuspended)

	tempPool, err = k.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(tempPool.Status, Equals, PoolAvailable)

	// when there are none gas token pool , and it is active , gas asset token pool should not be ragnarok
	busdPool := NewPool()
	busdAsset, err := common.NewAsset("BNB.BUSD-BD1")
	c.Assert(err, IsNil)
	busdPool.Asset = busdAsset
	busdPool.BalanceCacao = cosmos.NewUint(1000 * common.One)
	busdPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	busdPool.LPUnits = cosmos.NewUint(1600)
	busdPool.Status = PoolAvailable
	c.Assert(k.SetPool(ctx, busdPool), IsNil)

	networkMgr.k.SetMimir(ctx, "RAGNAROK-BNB-BNB", 1)
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	tempPool, err = k.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(tempPool.Status, Equals, PoolAvailable)
}

func (s *NetworkManagerV108TestSuite) TestCleanupAsgardIndex(c *C) {
	ctx, k := setupKeeperForTest(c)
	vault1 := NewVault(1024, ActiveVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault1), IsNil)
	vault2 := NewVault(1024, RetiringVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault2), IsNil)
	vault3 := NewVault(1024, InitVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault3), IsNil)
	vault4 := NewVault(1024, InactiveVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault4), IsNil)
	mgr := NewDummyMgrWithKeeper(k)
	networkMgr := newNetworkMgrV108(k, mgr.TxOutStore(), mgr.EventMgr())
	c.Assert(networkMgr.cleanupAsgardIndex(ctx), IsNil)
	containsVault := func(vaults Vaults, pubKey common.PubKey) bool {
		for _, item := range vaults {
			if item.PubKey.Equals(pubKey) {
				return true
			}
		}
		return false
	}
	asgards, err := k.GetAsgardVaults(ctx)
	c.Assert(err, IsNil)
	c.Assert(containsVault(asgards, vault1.PubKey), Equals, true)
	c.Assert(containsVault(asgards, vault2.PubKey), Equals, true)
	c.Assert(containsVault(asgards, vault3.PubKey), Equals, true)
	c.Assert(containsVault(asgards, vault4.PubKey), Equals, false)
}

func (*NetworkManagerV108TestSuite) TestPOLLiquidityAdd(c *C) {
	ctx, mgr := setupManagerForTest(c)

	net := newNetworkMgrV108(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())
	max := cosmos.NewUint(100)

	polAddress, err := mgr.Keeper().GetModuleAddress(ReserveName)
	c.Assert(err, IsNil)
	asgardAddress, err := mgr.Keeper().GetModuleAddress(AsgardName)
	c.Assert(err, IsNil)
	na := GetRandomValidatorNode(NodeActive)
	signer := na.NodeAddress
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, na), IsNil)

	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceCacao = cosmos.NewUint(2000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(20 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)
	c.Assert(mgr.Keeper().SetPool(ctx, btcPool), IsNil)

	// hit max
	util := cosmos.NewUint(1500)
	target := cosmos.NewUint(1000)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err := mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(7), Commentf("%d", lp.Units.Uint64()))

	// doesn't hit max
	util = cosmos.NewUint(1050)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(10), Commentf("%d", lp.Units.Uint64()))

	// no change needed
	util = cosmos.NewUint(1000)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(10), Commentf("%d", lp.Units.Uint64()))

	// not enough balance in the reserve module
	max = cosmos.NewUint(10000)
	util = cosmos.NewUint(50_000)
	btcPool.BalanceCacao = cosmos.NewUint(90000000000 * common.One)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(10), Commentf("%d", lp.Units.Uint64()))
}

func (*NetworkManagerV108TestSuite) TestPOLLiquidityWithdraw(c *C) {
	ctx, mgr := setupManagerForTest(c)

	net := newNetworkMgrV108(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())
	max := cosmos.NewUint(100)

	polAddress, err := mgr.Keeper().GetModuleAddress(ReserveName)
	c.Assert(err, IsNil)
	asgardAddress, err := mgr.Keeper().GetModuleAddress(AsgardName)
	c.Assert(err, IsNil)
	na := GetRandomValidatorNode(NodeActive)
	signer := na.NodeAddress
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, na), IsNil)

	vault := GetRandomVault()
	c.Assert(mgr.Keeper().SetVault(ctx, vault), IsNil)

	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceCacao = cosmos.NewUint(2000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(20 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)
	c.Assert(mgr.Keeper().SetPool(ctx, btcPool), IsNil)

	lps := LiquidityProviders{
		{
			Asset:             btcPool.Asset,
			CacaoAddress:      GetRandomBNBAddress(),
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     5,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingCacao:      cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			CacaoDepositValue: cosmos.ZeroUint(),
		},
		{
			Asset:             btcPool.Asset,
			CacaoAddress:      polAddress,
			AssetAddress:      common.NoAddress,
			LastAddHeight:     10,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingCacao:      cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			CacaoDepositValue: cosmos.ZeroUint(),
		},
	}
	for _, lp := range lps {
		mgr.Keeper().SetLiquidityProvider(ctx, lp)
	}

	// hit max
	util := cosmos.NewUint(500)
	target := cosmos.NewUint(1000)
	c.Assert(net.removePOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err := mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(792), Commentf("%d", lp.Units.Uint64()))
	// To withdraw max 1% (100 basis points) of the pool RUNE depth, asymmetrically withdraw as RUNE 0.5% of all pool units.
	// 0.5% of 1600 is 8; 800 minus 8 is 792.

	// doesn't hit max
	util = cosmos.NewUint(950)
	c.Assert(net.removePOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(788), Commentf("%d", lp.Units.Uint64()))
	// To withdraw 0.5% of the pool RUNE depth, asymmetrically withdraw as RUNE 0.25% of all pool units.
	// 0.25% of 1592 is 3.98 which rounds to 4; 792 minus 4 is 788.

	// no change needed
	util = cosmos.NewUint(1000)
	c.Assert(net.removePOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(788), Commentf("%d", lp.Units.Uint64()))
}

func (s *NetworkManagerV108TestSuite) TestSaverYieldFunc(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	net := newNetworkMgrV108(mgr.Keeper(), mgr.TxOutStore(), mgr.EventMgr())

	// mint synths
	coin := common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(10*common.One))
	c.Assert(mgr.Keeper().MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	spool := NewPool()
	spool.Asset = common.BTCAsset.GetSyntheticAsset()
	spool.BalanceAsset = coin.Amount
	spool.LPUnits = cosmos.NewUint(100)
	c.Assert(mgr.Keeper().SetPool(ctx, spool), IsNil)

	// first pool
	pool := NewPool()
	pool.Asset = common.BTCAsset
	pool.BalanceCacao = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.LPUnits = cosmos.NewUint(100)
	pool.CalcUnits(mgr.GetVersion(), coin.Amount)
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	c.Assert(net.paySaverYield(ctx, common.BTCAsset, cosmos.NewUint(50*common.One)), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset)
	c.Assert(err, IsNil)
	c.Assert(spool.BalanceAsset.String(), Equals, "1102400000", Commentf("%d", spool.BalanceAsset.Uint64()))
}

func (s *NetworkManagerV108TestSuite) TestSaverYieldCall(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	ver := GetCurrentVersion()
	constAccessor := constants.GetConstantValues(ver)

	na := GetRandomValidatorNode(NodeActive)
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, na), IsNil)

	coin := common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(10*common.One))
	spool := NewPool()
	spool.Asset = common.BTCAsset.GetSyntheticAsset()
	spool.BalanceAsset = coin.Amount
	spool.LPUnits = cosmos.NewUint(100)
	c.Assert(mgr.Keeper().SetPool(ctx, spool), IsNil)

	// first pool
	pool := NewPool()
	pool.Asset = common.BTCAsset
	pool.BalanceCacao = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.LPUnits = cosmos.NewUint(100)
	pool.CalcUnits(mgr.GetVersion(), coin.Amount)
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	networkMgr := newNetworkMgrV108(mgr.Keeper(), mgr.TxOutStore(), mgr.EventMgr())

	// test no fees collected
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset.GetSyntheticAsset())
	c.Assert(err, IsNil)
	c.Check(spool.BalanceAsset.Uint64(), Equals, uint64(1000000000), Commentf("%d", spool.BalanceAsset.Uint64()))

	c.Assert(mgr.Keeper().AddToLiquidityFees(ctx, pool.Asset, cosmos.NewUint(50*common.One)), IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset.GetSyntheticAsset())
	c.Assert(err, IsNil)
	c.Check(spool.BalanceAsset.String(), Equals, "1000000000", Commentf("%d", spool.BalanceAsset.Uint64()))

	// check we don't give yield when synth utilization is too high
	// add some synths
	coins := cosmos.NewCoins(cosmos.NewCoin("btc/btc", cosmos.NewInt(101*common.One))) // 51% utilization
	c.Assert(mgr.coinKeeper.MintCoins(ctx, ModuleName, coins), IsNil)
	c.Assert(mgr.Keeper().AddToLiquidityFees(ctx, pool.Asset, cosmos.NewUint(50*common.One)), IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset.GetSyntheticAsset())
	c.Assert(err, IsNil)
	c.Check(spool.BalanceAsset.String(), Equals, "1000000000", Commentf("%d", spool.BalanceAsset.Uint64()))
}
