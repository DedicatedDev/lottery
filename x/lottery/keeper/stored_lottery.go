package keeper

import (
	"github.com/DedicatedDev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredLottery set a specific storedLottery in the store from its index
func (k Keeper) SetStoredLottery(ctx sdk.Context, storedLottery types.StoredLottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredLotteryKeyPrefix))
	b := k.cdc.MustMarshal(&storedLottery)
	store.Set(types.StoredLotteryKey(
		storedLottery.Index,
	), b)
}

// GetStoredLottery returns a storedLottery from its index
func (k Keeper) GetStoredLottery(
	ctx sdk.Context,
	index string,

) (val types.StoredLottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredLotteryKeyPrefix))

	b := store.Get(types.StoredLotteryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredLottery removes a storedLottery from the store
func (k Keeper) RemoveStoredLottery(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredLotteryKeyPrefix))
	store.Delete(types.StoredLotteryKey(
		index,
	))
}

// GetAllStoredLottery returns all storedLottery
func (k Keeper) GetAllStoredLottery(ctx sdk.Context) (list []types.StoredLottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredLotteryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredLottery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
