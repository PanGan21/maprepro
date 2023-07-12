package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "maprepro/testutil/keeper"
	"maprepro/testutil/nullify"
	"maprepro/x/repro/keeper"
	"maprepro/x/repro/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMymap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mymap {
	items := make([]types.Mymap, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMymap(ctx, items[i])
	}
	return items
}

func TestMymapGet(t *testing.T) {
	keeper, ctx := keepertest.ReproKeeper(t)
	items := createNMymap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMymap(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMymapRemove(t *testing.T) {
	keeper, ctx := keepertest.ReproKeeper(t)
	items := createNMymap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMymap(ctx,
			item.Index,
		)
		_, found := keeper.GetMymap(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestMymapGetAll(t *testing.T) {
	keeper, ctx := keepertest.ReproKeeper(t)
	items := createNMymap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMymap(ctx)),
	)
}
