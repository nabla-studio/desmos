package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/golang/mock/gomock"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v5/app"

	"github.com/desmos-labs/desmos/v5/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v5/x/reactions/testutil"
	"github.com/desmos-labs/desmos/v5/x/reactions/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       storetypes.StoreKey

	ak *testutil.MockProfilesKeeper
	rk *testutil.MockRelationshipsKeeper
	pk *testutil.MockPostsKeeper
	sk *testutil.MockSubspacesKeeper
	k  keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
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

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, _ = app.MakeCodecs()

	// Mocks initializations
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.ak = testutil.NewMockProfilesKeeper(ctrl)
	suite.rk = testutil.NewMockRelationshipsKeeper(ctrl)
	suite.pk = testutil.NewMockPostsKeeper(ctrl)
	suite.sk = testutil.NewMockSubspacesKeeper(ctrl)

	suite.k = keeper.NewKeeper(suite.cdc, keys[types.StoreKey], suite.ak, suite.sk, suite.rk, suite.pk)
}
