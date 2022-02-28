package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mkXultra/oracle/testutil/keeper"
	"github.com/mkXultra/oracle/x/consuming/keeper"
	"github.com/mkXultra/oracle/x/consuming/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ConsumingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
