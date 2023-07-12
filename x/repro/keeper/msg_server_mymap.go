package keeper

import (
	"context"

	"maprepro/x/repro/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	mockThirdMapList := types.ThirdMapList{
		ThirdMap: []types.ThirdMap{
			{
				ThirdMap: map[string]int32{
					"key1": 10,
					"key2": 20,
				},
			},
			{
				ThirdMap: map[string]int32{
					"key3": 30,
					"key4": 40,
				},
			},
			{
				ThirdMap: map[string]int32{
					"key5": 30,
					"key6": 40,
				},
			},
			{
				ThirdMap: map[string]int32{
					"key7": 30,
					"key8": 40,
				},
			},
			{
				ThirdMap: map[string]int32{
					"key9":  30,
					"key10": 40,
				},
			},
		},
	}

	mockSecMap := types.SecMap{
		ThirdMaps: map[string]types.ThirdMapList{
			"secMap1": mockThirdMapList,
			"secMap2": mockThirdMapList,
			"secMap3": mockThirdMapList,
			"secMap4": mockThirdMapList,
		},
	}

	mockInnerMap := types.InnerMap{
		InnerMap: map[string]types.SecMap{
			"innerMap1": mockSecMap,
			"innerMap2": mockSecMap,
			"innerMap3": mockSecMap,
			"innerMap4": mockSecMap,
		},
	}

	var mymap = types.Mymap{
		Creator:  msg.Creator,
		Index:    msg.Index,
		InnerMap: &mockInnerMap,
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
