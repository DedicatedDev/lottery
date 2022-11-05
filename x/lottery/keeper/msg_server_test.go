package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DedicatedDev/lottery/testutil/keeper"
	"github.com/DedicatedDev/lottery/x/lottery"
	"github.com/DedicatedDev/lottery/x/lottery/keeper"
	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
