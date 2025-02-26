package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/reports/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetNextReportID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing report id is set properly",
			subspaceID: 1,
			reportID:   1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReportIDFromBytes(store.Get(types.NextReportIDStoreKey(1)))
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "existing report id is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 1)
			},
			subspaceID: 1,
			reportID:   2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReportIDFromBytes(store.Get(types.NextReportIDStoreKey(1)))
				suite.Require().Equal(uint64(2), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.SetNextReportID(ctx, tc.subspaceID, tc.reportID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetNextReportID() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		shouldErr   bool
		expReportID uint64
	}{
		{
			name:       "non existing report id returns error",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "existing report id is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 1)
			},
			subspaceID:  1,
			shouldErr:   false,
			expReportID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			reportID, err := suite.k.GetNextReportID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReportID, reportID)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteNextReportID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "existing report id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 1)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextReportIDStoreKey(1)))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteNextReportID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_ValidateReport() {
	testCases := []struct {
		name      string
		setup     func()
		report    types.Report
		shouldErr bool
	}{
		{
			name: "invalid report returns error",
			report: types.NewReport(
				0,
				1,
				[]uint32{1},
				"This content is spam",
				types.NewUserTarget("cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5"),
				"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "UserTarget - blocked reporter returns error",
			setup: func() {
				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
						"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
						uint64(1),
					).
					Return(true)
			},
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This content is spam",
				types.NewUserTarget("cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5"),
				"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "UserTarget - valid data returns no error",
			setup: func() {
				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
						"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
						uint64(1),
					).
					Return(false)
			},
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This content is spam",
				types.NewUserTarget("cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5"),
				"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
		},
		{
			name: "PostTarget - not found post returns error",
			setup: func() {
				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.Post{}, false)
			},
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This user is spamming",
				types.NewPostTarget(1),
				"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "PostTarget - blocked user returns error",
			setup: func() {
				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"",
						"This is a new post",
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
						0,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
						"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
						uint64(1),
					).
					Return(true)
			},
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This user is spamming",
				types.NewPostTarget(1),
				"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "PostsData - valid data returns no error",
			setup: func() {
				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"",
						"This is a new post",
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
						0,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos10s22qjua2n3law0ymstm3txm7764mfk2cjawq5",
						"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
						uint64(1),
					).
					Return(false)
			},
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This user is spamming",
				types.NewPostTarget(1),
				"cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}

			err := suite.k.ValidateReport(ctx, tc.report)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveReport() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		report types.Report
		check  func(ctx sdk.Context)
	}{
		{
			name: "post report is stored properly",
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This content is spam",
				types.NewPostTarget(1),
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReport(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewPostTarget(1),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), stored)

				// Check the content key
				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.PostReportStoreKey(1, 1, "cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z")))
			},
		},
		{
			name: "user report is stored properly",
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This content is spam",
				types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReport(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), stored)

				// Check the content key
				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.UserReportStoreKey(1, "cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm", "cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z")))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.SaveReport(ctx, tc.report)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasReport() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		expResult  bool
	}{
		{
			name:       "non existing report returns false",
			subspaceID: 1,
			reportID:   1,
			expResult:  false,
		},
		{
			name: "existing report returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewPostTarget(1),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reportID:   1,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasReport(ctx, tc.subspaceID, tc.reportID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasReported() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reporter   string
		target     types.ReportTarget
		expResult  bool
	}{
		{
			name: "not reported target returns false - different post id",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(2),
					"cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reporter:   "cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
			target:     types.NewPostTarget(1),
			expResult:  false,
		},
		{
			name: "not reported target returns false - different user address",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewUserTarget("cosmos1dzwwn72sevnakh4qhhpzmsqlsj3ehzf9n803yh"),
					"cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reporter:   "cosmos1dzwwn72sevnakh4qhhpzmsqlsj3ehzf9n803yh",
			target:     types.NewUserTarget("cosmos14uhwtt50cge4mywlr8897tef78gkjg75ugc9rq"),
			expResult:  false,
		},
		{
			name: "not reported target returns false - different subspace id",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					2,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reporter:   "cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
			target:     types.NewPostTarget(1),
			expResult:  false,
		},
		{
			name: "not reported target returns false - different reporter",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1hjvrc2rvy0jenpfquk536755x4cgvjqhqj6t3d",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reporter:   "cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
			target:     types.NewPostTarget(1),
			expResult:  false,
		},
		{
			name: "reported post returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reporter:   "cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
			target:     types.NewPostTarget(1),
			expResult:  true,
		},
		{
			name: "reported user returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewUserTarget("cosmos14uhwtt50cge4mywlr8897tef78gkjg75ugc9rq"),
					"cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reporter:   "cosmos1qqjdwjjxxgfpk9kvn0a6gpxmzgvd2z0jtd72e4",
			target:     types.NewUserTarget("cosmos14uhwtt50cge4mywlr8897tef78gkjg75ugc9rq"),
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasReported(ctx, tc.subspaceID, tc.reporter, tc.target)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReport() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		expFound   bool
		expReport  types.Report
	}{
		{
			name:       "non existing report returns false and empty report",
			subspaceID: 1,
			reportID:   1,
			expFound:   false,
			expReport:  types.Report{},
		},
		{
			name: "existing report returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewPostTarget(1),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reportID:   1,
			expFound:   true,
			expReport: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This content is spam",
				types.NewPostTarget(1),
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			report, found := suite.k.GetReport(ctx, tc.subspaceID, tc.reportID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expReport, report)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteReport() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing report is deleted properly",
			subspaceID: 1,
			reportID:   1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReport(ctx, 1, 1))
			},
		},
		{
			name: "existing report is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewPostTarget(1),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reportID:   1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReport(ctx, 1, 1))

				// Check the content key
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.PostReportStoreKey(1, 1, "cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z")))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteReport(ctx, tc.subspaceID, tc.reportID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
