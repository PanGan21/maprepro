package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MymapList: []Mymap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in mymap
	mymapIndexMap := make(map[string]struct{})

	for _, elem := range gs.MymapList {
		index := string(MymapKey(elem.Index))
		if _, ok := mymapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mymap")
		}
		mymapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
