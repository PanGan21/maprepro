package repro

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"maprepro/x/repro/keeper"
	"maprepro/x/repro/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the mymap
	for _, elem := range genState.MymapList {
		k.SetMymap(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MymapList = k.GetAllMymap(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
