package keeper_test

import (
	"testing"

	testkeeper "github.com/mkXultra/oracle/testutil/keeper"
	"github.com/mkXultra/oracle/x/consuming/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ConsumingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
