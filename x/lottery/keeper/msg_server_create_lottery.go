package keeper

import (
	"context"
	"strconv"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateLottery(goCtx context.Context, msg *types.MsgCreateLottery) (*types.MsgCreateLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatInt(int64(systemInfo.NextId), 10)
	storedLottery := types.StoredLottery{
		Index:        newIndex,
		BidCount:     0,
		MinBetAmount: msg.MintBetAmount,
		Fee:          msg.Fee,
	}
	err := storedLottery.ValidateLottery()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetStoredLottery(ctx, storedLottery)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	return &types.MsgCreateLotteryResponse{
		LotteryIndex: newIndex,
	}, nil
}
