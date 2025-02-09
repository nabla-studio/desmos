package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/feegrant"

	"github.com/desmos-labs/desmos/v5/testutil/simtesting"
	"github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// DONTCOVER

// SimulateMsgGrantAllowance tests and runs a single MsgGrantAllowance
func SimulateMsgGrantAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, grantee, signer, skip := randomGrantAllowanceFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgGrantAllowance", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgGrantAllowance(subspaceID, granter, grantee, &feegrant.BasicAllowance{})

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomGrantAllowanceFields returns the data used to build a random MsgGrantAllowance
func randomGrantAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, granter string, grantee types.Grantee, signer simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}

	subspaceID = RandomSubspace(r, subspaces).ID

	// Get a granter
	granters := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionManageAllowances))
	granter = RandomAddress(r, granters)
	if len(granters) == 0 {
		skip = true
		return
	}

	// 50% of having a user grantee, otherwise a group grantee
	if r.Intn(100) < 50 {
		accounts := ak.GetAllAccounts(ctx)
		granteeAddr := RandomAuthAccount(r, accounts).GetAddress().String()

		if k.HasUserGrant(ctx, subspaceID, granteeAddr) {
			// Skip because grant does exist
			skip = true
			return
		}
		if granteeAddr == granter {
			// Skip because granting to itself is not allowed
			skip = true
			return
		}

		grantee = types.NewUserGrantee(granteeAddr)
	} else {
		groups := k.GetSubspaceUserGroups(ctx, subspaceID)
		if len(groups) == 0 {
			// Skip because there are no groups
			skip = true
			return
		}

		group := RandomGroup(r, groups)
		if group.ID == 0 {
			// Skip because we cannot grant the default group
			skip = true
			return
		}

		if k.HasGroupGrant(ctx, subspaceID, group.ID) {
			// Skip because grant does exist
			skip = true
			return
		}

		grantee = types.NewGroupGrantee(group.ID)
	}

	// Get a signer account
	acc := GetAccount(granter, accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, granter, grantee, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRevokeAllowance tests and runs a single MsgRevokeAllowance
func SimulateMsgRevokeAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, grantee, signer, skip := randomRevokeAllowanceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRevokeAllowance", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRevokeAllowance(subspaceID, granter, grantee)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomRevokeAllowanceFields returns the data used to build a random MsgRevokeAllowance
func randomRevokeAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, granter string, grantee types.Grantee, signer simtypes.Account, skip bool) {
	// 50% of having user grants, otherwise group grants
	var grants []types.Grant
	if r.Intn(100) < 50 {
		grants = k.GetAllUserGrants(ctx)
	} else {
		grants = k.GetAllUserGroupsGrants(ctx)
	}

	if len(grants) == 0 {
		// Skip because there are no grants
		skip = true
		return
	}

	grant := RandomGrant(r, grants)
	subspaceID = grant.SubspaceID
	grantee = grant.Grantee.GetCachedValue().(types.Grantee)

	// Get a granter
	granters := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionManageAllowances))
	if len(granters) == 0 {
		skip = true
		return
	}

	granter = RandomAddress(r, granters)

	// Get a signer account
	acc := GetAccount(granter, accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, granter, grantee, signer, false
}
