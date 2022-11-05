package keeper

import (
	"context"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredLotteryAll(c context.Context, req *types.QueryAllStoredLotteryRequest) (*types.QueryAllStoredLotteryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedLotterys []types.StoredLottery
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedLotteryStore := prefix.NewStore(store, types.KeyPrefix(types.StoredLotteryKeyPrefix))

	pageRes, err := query.Paginate(storedLotteryStore, req.Pagination, func(key []byte, value []byte) error {
		var storedLottery types.StoredLottery
		if err := k.cdc.Unmarshal(value, &storedLottery); err != nil {
			return err
		}

		storedLotterys = append(storedLotterys, storedLottery)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredLotteryResponse{StoredLottery: storedLotterys, Pagination: pageRes}, nil
}

func (k Keeper) StoredLottery(c context.Context, req *types.QueryGetStoredLotteryRequest) (*types.QueryGetStoredLotteryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredLottery(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredLotteryResponse{StoredLottery: val}, nil
}
