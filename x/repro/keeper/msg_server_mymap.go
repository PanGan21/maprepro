package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"maprepro/x/repro/types"
)

func (k msgServer) CreateMymap(goCtx context.Context, msg *types.MsgCreateMymap) (*types.MsgCreateMymapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMymap(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var mymap = types.Mymap{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetMymap(
		ctx,
		mymap,
	)
	return &types.MsgCreateMymapResponse{}, nil
}

func (k msgServer) UpdateMymap(goCtx context.Context, msg *types.MsgUpdateMymap) (*types.MsgUpdateMymapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMymap(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var mymap = types.Mymap{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetMymap(ctx, mymap)

	return &types.MsgUpdateMymapResponse{}, nil
}

func (k msgServer) DeleteMymap(goCtx context.Context, msg *types.MsgDeleteMymap) (*types.MsgDeleteMymapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMymap(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMymap(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteMymapResponse{}, nil
}
