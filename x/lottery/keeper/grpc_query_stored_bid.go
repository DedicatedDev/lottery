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

func (k Keeper) StoredBidAll(c context.Context, req *types.QueryAllStoredBidRequest) (*types.QueryAllStoredBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedBids []types.StoredBid
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedBidStore := prefix.NewStore(store, types.KeyPrefix(types.StoredBidKeyPrefix))

	pageRes, err := query.Paginate(storedBidStore, req.Pagination, func(key []byte, value []byte) error {
		var storedBid types.StoredBid
		if err := k.cdc.Unmarshal(value, &storedBid); err != nil {
			return err
		}

		storedBids = append(storedBids, storedBid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredBidResponse{StoredBid: storedBids, Pagination: pageRes}, nil
}

func (k Keeper) StoredBid(c context.Context, req *types.QueryGetStoredBidRequest) (*types.QueryGetStoredBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredBid(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredBidResponse{StoredBid: val}, nil
}
