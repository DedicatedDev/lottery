package keeper

import (
	"context"
	"fmt"
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
		MinBetAmount: msg.MinBetAmount,
		Fee:          msg.Fee,
	}
	err := storedLottery.ValidateLottery()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetStoredLottery(ctx, storedLottery)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.LotteryCreatedEventType,
		sdk.NewAttribute(types.LotteryCreatedEventTypeCreator, msg.Creator),
		sdk.NewAttribute(types.LotteryCreatedEventTypeLotteryId, newIndex),
		sdk.NewAttribute(types.LotteryCreatedEventTypeLotteryMinBetAmount, fmt.Sprintf("%d", msg.MinBetAmount)),
		sdk.NewAttribute(types.LotteryCreatedEventTypeLotteryFee, fmt.Sprintf("%d", msg.Fee)),
	))
	return &types.MsgCreateLotteryResponse{
		LotteryIndex: newIndex,
	}, nil
}
