package lottery

import (
	"github.com/DedicatedDev/lottery/x/lottery/keeper"
	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	k.SetSystemInfo(ctx, genState.SystemInfo)
	// Set all the storedLottery
	for _, elem := range genState.StoredLotteryList {
		k.SetStoredLottery(ctx, elem)
	}
	// Set all the storedBid
	for _, elem := range genState.StoredBidList {
		k.SetStoredBid(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StoredLotteryList = k.GetAllStoredLottery(ctx)
	genesis.StoredBidList = k.GetAllStoredBid(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
