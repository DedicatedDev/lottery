package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BidToLottery(goCtx context.Context, msg *types.MsgBidToLottery) (*types.MsgBidToLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check lottery exist or not
	storedLottery, found := k.Keeper.GetStoredLottery(ctx, strconv.FormatInt(int64(msg.LotteryId), 10))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrLotteryNotFound, "%d", msg.LotteryId)
	}
	err := storedLottery.ValidateLottery()
	if err != nil {
		return nil, err
	}
	newBidId, _ := types.GenValidBidIndex(msg.Creator, msg.LotteryId)
	newBid := types.StoredBid{
		Index:     newBidId,
		LotteryId: msg.LotteryId,
		BetAmount: msg.BidAmount,
		Bidder:    msg.Creator,
	}
	err = newBid.ValidateBid()
	if err != nil {
		return nil, err
	}
	if newBid.BetAmount < storedLottery.Fee+storedLottery.MinBetAmount {
		return nil, sdkerrors.Wrap(types.ErrBidNotEnoughFund, fmt.Sprintf("LotteryId:%d", msg.LotteryId))
	}

	_, found = k.Keeper.GetStoredBid(ctx, newBidId)
	k.SetStoredBid(ctx, newBid)
	if !found {
		storedLottery.BidCount++
	}
	k.SetStoredLottery(ctx, storedLottery)
	return &types.MsgBidToLotteryResponse{
		BidId:    newBidId,
		BidCount: storedLottery.BidCount,
	}, nil
}
