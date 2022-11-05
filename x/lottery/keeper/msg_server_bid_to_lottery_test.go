package keeper_test

import (
	"context"
	"fmt"
	"testing"

	keepertest "github.com/DedicatedDev/lottery/testutil/keeper"
	"github.com/DedicatedDev/lottery/x/lottery"
	"github.com/DedicatedDev/lottery/x/lottery/keeper"
	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneLotteryForBid(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateLottery(context, &types.MsgCreateLottery{
		Creator:      alice,
		MinBetAmount: DefaultMinBetAmount,
		Fee:          DefaultFee,
	})
	return server, *k, context
}
func TestBidToLottery(t *testing.T) {
	msgServer, _, ctx := setupMsgServerWithOneLotteryForBid(t)
	bidToLotteryRes, err := msgServer.BidToLottery(ctx, &types.MsgBidToLottery{
		Creator:   bob,
		LotteryId: 1,
		BidAmount: 10,
	})
	require.Nil(t, err)
	newBidIndex, _ := types.GenValidBidIndex(bob, 1)
	require.EqualValues(t, types.MsgBidToLotteryResponse{
		BidId:    newBidIndex,
		BidCount: 1,
	}, *bidToLotteryRes)
}

func TestBidToLotteryCannotBidWrong(t *testing.T) {
	msgServer, k, ctx := setupMsgServerWithOneLotteryForBid(t)
	context := sdk.UnwrapSDKContext(ctx)
	storedLottery, _ := k.GetStoredLottery(sdk.UnwrapSDKContext(context), "1")
	storedLottery.MinBetAmount = 4
	k.SetStoredLottery(context, storedLottery)
	_, err := msgServer.BidToLottery(ctx, &types.MsgBidToLottery{
		Creator:   bob,
		LotteryId: 1,
		BidAmount: 10,
	})
	require.NotNil(t, err)
}

func TestBidToLotteryWithWrongFund(t *testing.T) {
	msgServer, _, ctx := setupMsgServerWithOneLotteryForBid(t)
	_, err := msgServer.BidToLottery(ctx, &types.MsgBidToLottery{
		Creator:   bob,
		LotteryId: 1,
		BidAmount: 3,
	})
	require.NotNil(t, err)
	require.EqualValues(t, err, types.ErrInvalidBetAmount)
}

func TestBidToLotteryWithWrongCreator(t *testing.T) {
	msgServer, _, ctx := setupMsgServerWithOneLotteryForBid(t)
	_, err := msgServer.BidToLottery(ctx, &types.MsgBidToLottery{
		Creator:   "",
		LotteryId: 1,
		BidAmount: 3,
	})
	require.NotNil(t, err)
	require.EqualValues(t, err.Error(), "empty address string is not allowed")
}

func TestBidToLotteryWithWrongLotteryId(t *testing.T) {
	msgServer, _, ctx := setupMsgServerWithOneLotteryForBid(t)
	_, err := msgServer.BidToLottery(ctx, &types.MsgBidToLottery{
		Creator:   bob,
		LotteryId: 39,
		BidAmount: 3,
	})
	require.NotNil(t, err)
	expectedErr := sdkerrors.Wrapf(types.ErrLotteryNotFound, "%d", 39)
	require.EqualValues(t, err.Error(), expectedErr.Error())
}

func TestBidToLotteryEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneLotteryForBid(t)
	msgServer.BidToLottery(context, &types.MsgBidToLottery{
		Creator:   bob,
		LotteryId: 1,
		BidAmount: 10,
	})
	newIndex, _ := types.GenValidBidIndex(bob, 1)
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-bid-created",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: bob},
			{Key: "bid-id", Value: newIndex},
			{Key: "lottery-id", Value: "1"},
			{Key: "total-bidder", Value: fmt.Sprintf("%d", 1)},
		},
	}, event)
}

func TestBidToLottery2Emitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneLotteryForBid(t)
	msgServer.BidToLottery(context, &types.MsgBidToLottery{
		Creator:   bob,
		LotteryId: 1,
		BidAmount: 10,
	})
	msgServer.BidToLottery(context, &types.MsgBidToLottery{
		Creator:   carol,
		LotteryId: 1,
		BidAmount: 20,
	})
	newIndex, _ := types.GenValidBidIndex(carol, 1)
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[0]
	fmt.Println(event)
	require.Equal(t, "new-bid-created", event.Type)
	require.EqualValues(t, []sdk.Attribute{
		{Key: "creator", Value: carol},
		{Key: "bid-id", Value: newIndex},
		{Key: "lottery-id", Value: "1"},
		{Key: "total-bidder", Value: fmt.Sprintf("%d", 2)},
	}, event.Attributes[4:])
}
