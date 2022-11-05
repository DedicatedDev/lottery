package keeper_test

import (
	"testing"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func TestCreateLottery(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateLottery(context, &types.MsgCreateLottery{
		Creator:       alice,
		MintBetAmount: 5,
		Fee:           1,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateLotteryResponse{
		LotteryIndex: "",
	}, *createResponse)
}
