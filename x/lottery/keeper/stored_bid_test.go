package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/DedicatedDev/lottery/testutil/keeper"
	"github.com/DedicatedDev/lottery/testutil/nullify"
	"github.com/DedicatedDev/lottery/x/lottery/keeper"
	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredBid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredBid {
	items := make([]types.StoredBid, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredBid(ctx, items[i])
	}
	return items
}

func TestStoredBidGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredBid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredBid(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredBidRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredBid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredBid(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredBid(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredBidGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredBid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredBid(ctx)),
	)
}
