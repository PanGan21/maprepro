package repro_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "maprepro/testutil/keeper"
	"maprepro/testutil/nullify"
	"maprepro/x/repro"
	"maprepro/x/repro/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MymapList: []types.Mymap{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ReproKeeper(t)
	repro.InitGenesis(ctx, *k, genesisState)
	got := repro.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MymapList, got.MymapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
