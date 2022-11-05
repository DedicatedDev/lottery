package keeper

import (
	"github.com/DedicatedDev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredBid set a specific storedBid in the store from its index
func (k Keeper) SetStoredBid(ctx sdk.Context, storedBid types.StoredBid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBidKeyPrefix))
	b := k.cdc.MustMarshal(&storedBid)
	store.Set(types.StoredBidKey(
		storedBid.Index,
	), b)
}

// GetStoredBid returns a storedBid from its index
func (k Keeper) GetStoredBid(
	ctx sdk.Context,
	index string,

) (val types.StoredBid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBidKeyPrefix))

	b := store.Get(types.StoredBidKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredBid removes a storedBid from the store
func (k Keeper) RemoveStoredBid(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBidKeyPrefix))
	store.Delete(types.StoredBidKey(
		index,
	))
}

// GetAllStoredBid returns all storedBid
func (k Keeper) GetAllStoredBid(ctx sdk.Context) (list []types.StoredBid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredBid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
