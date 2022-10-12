package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "wonderfulbet/testutil/keeper"
	"wonderfulbet/x/wonderfulbet/keeper"
	"wonderfulbet/x/wonderfulbet/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WonderfulbetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
