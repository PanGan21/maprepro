package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"maprepro/x/repro/types"
)

func (k Keeper) MymapAll(goCtx context.Context, req *types.QueryAllMymapRequest) (*types.QueryAllMymapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mymaps []types.Mymap
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	mymapStore := prefix.NewStore(store, types.KeyPrefix(types.MymapKeyPrefix))

	pageRes, err := query.Paginate(mymapStore, req.Pagination, func(key []byte, value []byte) error {
		var mymap types.Mymap
		if err := k.cdc.Unmarshal(value, &mymap); err != nil {
			return err
		}

		mymaps = append(mymaps, mymap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMymapResponse{Mymap: mymaps, Pagination: pageRes}, nil
}

func (k Keeper) Mymap(goCtx context.Context, req *types.QueryGetMymapRequest) (*types.QueryGetMymapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetMymap(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMymapResponse{Mymap: val}, nil
}
