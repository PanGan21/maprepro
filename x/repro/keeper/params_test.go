package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "maprepro/testutil/keeper"
	"maprepro/x/repro/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ReproKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
