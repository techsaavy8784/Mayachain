package types

import (
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
)

type MsgAddLiquiditySuite struct{}

var _ = Suite(&MsgAddLiquiditySuite{})

func (MsgAddLiquiditySuite) TestMsgAddLiquidity(c *C) {
	addr := GetRandomBech32Addr()
	c.Check(addr.Empty(), Equals, false)
	runeAddress := GetRandomBaseAddress()
	assetAddress := GetRandomBNBAddress()
	txID := GetRandomTxHash()
	c.Check(txID.IsEmpty(), Equals, false)
	tx := common.NewTx(
		txID,
		runeAddress,
		GetRandomBaseAddress(),
		common.Coins{
			common.NewCoin(common.BTCAsset, cosmos.NewUint(100000000)),
		},
		BNBGasFeeSingleton,
		"",
	)
	m := NewMsgAddLiquidity(tx, common.BNBAsset, cosmos.NewUint(100000000), cosmos.NewUint(100000000), runeAddress, assetAddress, common.NoAddress, cosmos.ZeroUint(), addr, 1)
	EnsureMsgBasicCorrect(m, c)
	c.Check(m.Type(), Equals, "add_liquidity")

	inputs := []struct {
		asset     common.Asset
		r         cosmos.Uint
		amt       cosmos.Uint
		runeAddr  common.Address
		assetAddr common.Address
		txHash    common.TxID
		signer    cosmos.AccAddress
	}{
		{
			asset:     common.Asset{},
			r:         cosmos.NewUint(100000000),
			amt:       cosmos.NewUint(100000000),
			runeAddr:  runeAddress,
			assetAddr: assetAddress,
			txHash:    txID,
			signer:    addr,
		},
		{
			asset:     common.BNBAsset,
			r:         cosmos.NewUint(100000000),
			amt:       cosmos.NewUint(100000000),
			runeAddr:  common.NoAddress,
			assetAddr: common.NoAddress,
			txHash:    txID,
			signer:    addr,
		},
		{
			asset:     common.BNBAsset,
			r:         cosmos.NewUint(100000000),
			amt:       cosmos.NewUint(100000000),
			runeAddr:  runeAddress,
			assetAddr: assetAddress,
			txHash:    common.TxID(""),
			signer:    addr,
		},
		{
			asset:     common.BNBAsset,
			r:         cosmos.NewUint(100000000),
			amt:       cosmos.NewUint(100000000),
			runeAddr:  runeAddress,
			assetAddr: assetAddress,
			txHash:    txID,
			signer:    cosmos.AccAddress{},
		},
	}
	for i, item := range inputs {
		tx = common.NewTx(
			item.txHash,
			GetRandomBaseAddress(),
			GetRandomBNBAddress(),
			common.Coins{
				common.NewCoin(item.asset, item.r),
			},
			BNBGasFeeSingleton,
			"",
		)
		m = NewMsgAddLiquidity(tx, item.asset, item.r, item.amt, item.runeAddr, item.assetAddr, common.NoAddress, cosmos.ZeroUint(), item.signer, 1)
		c.Assert(m.ValidateBasicV108(), NotNil, Commentf("%d) %s\n", i, m))
	}
	// If affiliate fee basis point is more than 1000 , the message should be rejected
	m1 := NewMsgAddLiquidity(tx, common.BNBAsset, cosmos.NewUint(100*common.One), cosmos.NewUint(100*common.One), GetRandomBaseAddress(), GetRandomBNBAddress(), GetRandomBaseAddress(), cosmos.NewUint(1024), GetRandomBech32Addr(), 1)
	c.Assert(m1.ValidateBasicV108(), NotNil)

	// check that we can't have zero asset and zero rune amounts in v63
	m1 = NewMsgAddLiquidity(tx, common.BNBAsset, cosmos.ZeroUint(), cosmos.ZeroUint(), GetRandomBaseAddress(), GetRandomBNBAddress(), GetRandomBaseAddress(), cosmos.ZeroUint(), GetRandomBech32Addr(), 1)
	c.Assert(m1.ValidateBasicV63(), NotNil)

	// check that we can have zero asset and zero rune amounts in v108
	m1 = NewMsgAddLiquidity(tx, common.BNBAsset, cosmos.ZeroUint(), cosmos.ZeroUint(), GetRandomBaseAddress(), GetRandomBNBAddress(), GetRandomBaseAddress(), cosmos.ZeroUint(), GetRandomBech32Addr(), 1)
	c.Assert(m1.ValidateBasicV108(), IsNil)
}
