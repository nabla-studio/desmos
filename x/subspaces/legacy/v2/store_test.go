package v2_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/testutil/storetesting"
	v2 "github.com/desmos-labs/desmos/v5/x/subspaces/legacy/v2"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "groups permissions are sanitized properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				group := v2.NewUserGroup(11, 11, "Test group", "", 0b11111111111111111111111111000001)
				kvStore.Set(v2.GroupStoreKey(group.SubspaceID, group.ID), cdc.MustMarshal(&group))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Check the permissions
				var group v2.UserGroup
				cdc.MustUnmarshal(kvStore.Get(v2.GroupStoreKey(11, 11)), &group)
				require.Equal(t, v2.PermissionWrite, group.Permissions)
			},
		},
		{
			name: "user permissions are sanitized properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				addr, err := sdk.AccAddressFromBech32("cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv")
				require.NoError(t, err)

				kvStore.Set(v2.UserPermissionStoreKey(11, addr), v2.MarshalPermission(0b11111111111111111111111111000001))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				addr, err := sdk.AccAddressFromBech32("cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv")
				require.NoError(t, err)

				// Check the permissions
				stored := v2.UnmarshalPermission(kvStore.Get(v2.UserPermissionStoreKey(11, addr)))
				require.Equal(t, v2.PermissionWrite, stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v2.MigrateStore(ctx, keys[types.StoreKey], cdc)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
