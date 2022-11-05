package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DedicatedDev/lottery/testutil/keeper"
	"github.com/DedicatedDev/lottery/x/lottery"
	"github.com/DedicatedDev/lottery/x/lottery/keeper"
	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const DefaultMinBetAmount = 5
const DefaultFee = 1
const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateLottery(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
func TestCreateLottery(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateLottery(t)
	createResponse, err := msgServer.CreateLottery(context, &types.MsgCreateLottery{
		Creator:       alice,
		MintBetAmount: DefaultMinBetAmount,
		Fee:           DefaultFee,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateLotteryResponse{
		LotteryIndex: "1",
	}, *createResponse)
}

func TestCreate1LotteryHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateLottery(t)
	msgSrvr.CreateLottery(context, &types.MsgCreateLottery{
		Creator:       alice,
		MintBetAmount: DefaultMinBetAmount,
		Fee:           DefaultFee,
	})
	systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId: 2,
	}, systemInfo)
	game1, found1 := keeper.GetStoredLottery(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredLottery{
		Index:        "1",
		BidCount:     0,
		MinBetAmount: DefaultMinBetAmount,
		Fee:          DefaultFee,
	}, game1)
}
