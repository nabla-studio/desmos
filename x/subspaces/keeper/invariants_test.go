package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
)

func (suite *KeeperTestSuite) TestValidSubspacesInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing next subspace id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expBroken: true,
		},
		{
			name: "invalid subspace id compared to next subspace id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 1)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expBroken: true,
		},
		{
			name: "missing next section id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 2)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.DeleteNextSectionID(ctx, 1)
			},
			expBroken: true,
		},
		{
			name: "missing next group id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 2)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.DeleteNextGroupID(ctx, 1)
			},
			expBroken: true,
		},
		{
			name: "invalid subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 2)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Time{},
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 1)
				suite.k.SetNextGroupID(ctx, 1, 1)
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 2)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 1)
				suite.k.SetNextGroupID(ctx, 1, 1)
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidSubspacesInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidSectionsInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "missing subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))
			},
			expBroken: true,
		},
		{
			name: "missing next section id returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid section id compared to next section id returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 1)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))
			},
			expBroken: true,
		},
		{
			name: "missing parent section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					3,
					2,
					"Test section",
					"Test section",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid section breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"",
					"Test section",
				))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidSectionsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidUserGroupsInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			expBroken: true,
		},
		{
			name: "non existing section breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			expBroken: true,
		},
		{
			name: "non existing next group id returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid group id compared to next group id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextGroupID(ctx, 1, 1)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid group breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextGroupID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextGroupID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidUserGroupsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidUserGroupMembersInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4")
			},
			expBroken: true,
		},
		{
			name: "non existing user group breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4")
			},
			expBroken: true,
		},
		{
			name: "invalid data breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.AddUserToGroup(ctx, 1, 0, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4")
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4")
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidUserGroupMembersInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidUserPermissionsInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionEditSubspace),
				)
			},
			expBroken: true,
		},
		{
			name: "non existing section breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionEditSubspace),
				)
			},
			expBroken: true,
		},
		{
			name: "invalid entry breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SetUserPermissions(ctx, 1, 1, "", types.NewPermissions())
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionEditSubspace),
				)
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidUserPermissionsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidUserGrantsInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewUserGrantee("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					&feegrant.BasicAllowance{},
				))
			},
			expBroken: true,
		},
		{
			name: "invalid data breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewUserGrantee("cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4"),
					&feegrant.BasicAllowance{},
				))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewUserGrantee("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					&feegrant.BasicAllowance{},
				))
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidUserGrantsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidGroupGrantsInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{}))
			},
			expBroken: true,
		},
		{
			name: "non existing group breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{},
				))
			},
			expBroken: true,
		},
		{
			name: "invalid data breaks invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveGrant(ctx, types.NewGrant(
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{},
				))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{},
				))
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidGroupGrantsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}
