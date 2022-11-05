package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemInfo: SystemInfo{
			NextId: uint64(DefaultIndex),
		},
		StoredLotteryList: []StoredLottery{},
		StoredBidList:     []StoredBid{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedLottery
	storedLotteryIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredLotteryList {
		index := string(StoredLotteryKey(elem.Index))
		if _, ok := storedLotteryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedLottery")
		}
		storedLotteryIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in storedBid
	storedBidIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredBidList {
		index := string(StoredBidKey(elem.Index))
		if _, ok := storedBidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedBid")
		}
		storedBidIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
