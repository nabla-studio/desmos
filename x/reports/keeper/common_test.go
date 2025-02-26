package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/mock/gomock"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/x/reports/keeper"
	"github.com/desmos-labs/desmos/v5/x/reports/testutil"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc codec.Codec
	ctx sdk.Context

	storeKey storetypes.StoreKey
	k        keeper.Keeper

	ak *testutil.MockProfilesKeeper
	sk *testutil.MockSubspacesKeeper
	rk *testutil.MockRelationshipsKeeper
	pk *testutil.MockPostsKeeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, _ = app.MakeCodecs()

	// Mocks initializations
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	// Setup keepers
	suite.sk = testutil.NewMockSubspacesKeeper(ctrl)
	suite.rk = testutil.NewMockRelationshipsKeeper(ctrl)
	suite.ak = testutil.NewMockProfilesKeeper(ctrl)
	suite.pk = testutil.NewMockPostsKeeper(ctrl)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.ak,
		suite.sk,
		suite.rk,
		suite.pk,
		authtypes.NewModuleAddress("gov").String(),
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
