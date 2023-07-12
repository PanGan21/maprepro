package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"maprepro/x/repro/types"
)

// SetMymap set a specific mymap in the store from its index
func (k Keeper) SetMymap(ctx sdk.Context, mymap types.Mymap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MymapKeyPrefix))
	b := k.cdc.MustMarshal(&mymap)
	store.Set(types.MymapKey(
		mymap.Index,
	), b)
}

// GetMymap returns a mymap from its index
func (k Keeper) GetMymap(
	ctx sdk.Context,
	index string,

) (val types.Mymap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MymapKeyPrefix))

	b := store.Get(types.MymapKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMymap removes a mymap from the store
func (k Keeper) RemoveMymap(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MymapKeyPrefix))
	store.Delete(types.MymapKey(
		index,
	))
}

// GetAllMymap returns all mymap
func (k Keeper) GetAllMymap(ctx sdk.Context) (list []types.Mymap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MymapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mymap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
