package repro

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"maprepro/testutil/sample"
	reprosimulation "maprepro/x/repro/simulation"
	"maprepro/x/repro/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = reprosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateMymap = "op_weight_msg_mymap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMymap int = 100

	opWeightMsgUpdateMymap = "op_weight_msg_mymap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMymap int = 100

	opWeightMsgDeleteMymap = "op_weight_msg_mymap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMymap int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	reproGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MymapList: []types.Mymap{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&reproGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMymap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMymap, &weightMsgCreateMymap, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMymap = defaultWeightMsgCreateMymap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMymap,
		reprosimulation.SimulateMsgCreateMymap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMymap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMymap, &weightMsgUpdateMymap, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMymap = defaultWeightMsgUpdateMymap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMymap,
		reprosimulation.SimulateMsgUpdateMymap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMymap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMymap, &weightMsgDeleteMymap, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMymap = defaultWeightMsgDeleteMymap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMymap,
		reprosimulation.SimulateMsgDeleteMymap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
