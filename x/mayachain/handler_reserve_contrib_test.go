package mayachain

import (
	"errors"

	"github.com/blang/semver"
	se "github.com/cosmos/cosmos-sdk/types/errors"
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
	cosmos "gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	keeper "gitlab.com/mayachain/mayanode/x/mayachain/keeper"
)

type HandlerReserveContributorSuite struct{}

var _ = Suite(&HandlerReserveContributorSuite{})

type reserveContributorKeeper struct {
	keeper.Keeper
	errGetNetwork bool
	errSetNetwork bool
}

func newReserveContributorKeeper(k keeper.Keeper) *reserveContributorKeeper {
	return &reserveContributorKeeper{
		Keeper: k,
	}
}

func (k *reserveContributorKeeper) GetNetwork(ctx cosmos.Context) (Network, error) {
	if k.errGetNetwork {
		return Network{}, errKaboom
	}
	return k.Keeper.GetNetwork(ctx)
}

func (k *reserveContributorKeeper) SetNetwork(ctx cosmos.Context, data Network) error {
	if k.errSetNetwork {
		return errKaboom
	}
	return k.Keeper.SetNetwork(ctx, data)
}

type reserveContributorHandlerHelper struct {
	ctx                cosmos.Context
	version            semver.Version
	keeper             *reserveContributorKeeper
	nodeAccount        NodeAccount
	constAccessor      constants.ConstantValues
	reserveContributor ReserveContributor
}

func newReserveContributorHandlerHelper(c *C) reserveContributorHandlerHelper {
	ctx, k := setupKeeperForTest(c)
	ctx = ctx.WithBlockHeight(1023)

	version := GetCurrentVersion()
	keeper := newReserveContributorKeeper(k)

	// active account
	nodeAccount := GetRandomValidatorNode(NodeActive)
	bp := NewBondProviders(nodeAccount.NodeAddress)
	acc, err := nodeAccount.BondAddress.AccAddress()
	c.Assert(err, IsNil)
	bp.Providers = append(bp.Providers, NewBondProvider(acc))
	bp.Providers[0].Bonded = true
	SetupLiquidityBondForTest(c, ctx, k, common.BNBAsset, nodeAccount.BondAddress, nodeAccount, cosmos.NewUint(100*common.One))
	c.Assert(k.SetBondProviders(ctx, bp), IsNil)
	c.Assert(keeper.SetNodeAccount(ctx, nodeAccount), IsNil)
	constAccessor := constants.GetConstantValues(version)

	reserveContributor := ReserveContributor{
		Address: GetRandomBNBAddress(),
		Amount:  cosmos.NewUint(100 * common.One),
	}
	return reserveContributorHandlerHelper{
		ctx:                ctx,
		version:            version,
		keeper:             keeper,
		nodeAccount:        nodeAccount,
		constAccessor:      constAccessor,
		reserveContributor: reserveContributor,
	}
}

func (h HandlerReserveContributorSuite) TestReserveContributorHandler(c *C) {
	testCases := []struct {
		name           string
		messageCreator func(helper reserveContributorHandlerHelper) cosmos.Msg
		runner         func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error)
		expectedResult error
		validator      func(helper reserveContributorHandlerHelper, msg cosmos.Msg, result *cosmos.Result, c *C)
	}{
		{
			name: "invalid message should return error",
			messageCreator: func(helper reserveContributorHandlerHelper) cosmos.Msg {
				return NewMsgNoOp(GetRandomObservedTx(), helper.nodeAccount.NodeAddress, "")
			},
			runner: func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error) {
				return handler.Run(helper.ctx, msg)
			},
			expectedResult: errInvalidMessage,
		},
		{
			name: "empty signer should return an error",
			messageCreator: func(helper reserveContributorHandlerHelper) cosmos.Msg {
				return NewMsgReserveContributor(GetRandomTx(), helper.reserveContributor, cosmos.AccAddress{})
			},
			runner: func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error) {
				return handler.Run(helper.ctx, msg)
			},
			expectedResult: se.ErrInvalidAddress,
		},
		{
			name: "empty contributor address should return an error",
			messageCreator: func(helper reserveContributorHandlerHelper) cosmos.Msg {
				return NewMsgReserveContributor(GetRandomTx(), ReserveContributor{
					Address: common.NoAddress,
					Amount:  cosmos.NewUint(100),
				}, helper.nodeAccount.NodeAddress)
			},
			runner: func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error) {
				return handler.Run(helper.ctx, msg)
			},
			expectedResult: se.ErrUnknownRequest,
		},
		{
			name: "empty contributor amount should return an error",
			messageCreator: func(helper reserveContributorHandlerHelper) cosmos.Msg {
				return NewMsgReserveContributor(GetRandomTx(), ReserveContributor{
					Address: GetRandomBNBAddress(),
					Amount:  cosmos.ZeroUint(),
				}, helper.nodeAccount.NodeAddress)
			},
			runner: func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error) {
				return handler.Run(helper.ctx, msg)
			},
			expectedResult: se.ErrUnknownRequest,
		},
		{
			name: "invalid tx should return an error",
			messageCreator: func(helper reserveContributorHandlerHelper) cosmos.Msg {
				tx := GetRandomTx()
				tx.ID = ""
				return NewMsgReserveContributor(tx, helper.reserveContributor, helper.nodeAccount.NodeAddress)
			},
			runner: func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error) {
				return handler.Run(helper.ctx, msg)
			},
			expectedResult: se.ErrUnknownRequest,
		},
		{
			name: "normal reserve contribute message should return success",
			messageCreator: func(helper reserveContributorHandlerHelper) cosmos.Msg {
				tx := GetRandomTx()
				tx.Coins = common.NewCoins(common.NewCoin(common.BaseAsset(), cosmos.NewUint(10*common.One)))
				return NewMsgReserveContributor(tx, helper.reserveContributor, helper.nodeAccount.NodeAddress)
			},
			runner: func(handler ReserveContributorHandler, helper reserveContributorHandlerHelper, msg cosmos.Msg) (*cosmos.Result, error) {
				return handler.Run(helper.ctx, msg)
			},
			expectedResult: nil,
		},
	}
	for _, tc := range testCases {
		_, mgr := setupManagerForTest(c)
		helper := newReserveContributorHandlerHelper(c)
		mgr.K = helper.keeper
		handler := NewReserveContributorHandler(mgr)
		msg := tc.messageCreator(helper)
		result, err := tc.runner(handler, helper, msg)
		if tc.expectedResult == nil {
			c.Check(err, IsNil)
		} else {
			c.Check(errors.Is(err, tc.expectedResult), Equals, true, Commentf("name:%s", tc.name))
		}
		if tc.validator != nil {
			tc.validator(helper, msg, result, c)
		}
	}
}
