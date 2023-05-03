package v2_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/storetesting"

	v2 "github.com/desmos-labs/desmos/v4/x/fees/legacy/v2"
	"github.com/desmos-labs/desmos/v4/x/fees/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps paramstypes.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	keys := sdk.NewKVStoreKeys(types.StoreKey)

	testCases := []struct {
		name          string
		setupSubspace func() mockSubspace
		shouldErr     bool
		check         func(ctx sdk.Context)
	}{
		{
			name: "invalid params returns error",
			setupSubspace: func() mockSubspace {
				return newMockSubspace(types.NewParams([]types.MinFee{
					types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
				}))
			},
			shouldErr: true,
		},
		{
			name: "params migrates properly",
			setupSubspace: func() mockSubspace {
				return newMockSubspace(types.DefaultParams())
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var params types.Params
				bz := store.Get(types.ParamsKey)
				require.NoError(t, cdc.Unmarshal(bz, &params))
				require.Equal(t, types.DefaultParams(), params)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, nil, nil)

			mockSubspace := tc.setupSubspace()

			err := v2.MigrateStore(ctx, keys[types.StoreKey], mockSubspace, cdc)
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