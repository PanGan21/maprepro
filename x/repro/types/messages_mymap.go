package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMymap = "create_mymap"
	TypeMsgUpdateMymap = "update_mymap"
	TypeMsgDeleteMymap = "delete_mymap"
)

var _ sdk.Msg = &MsgCreateMymap{}

func NewMsgCreateMymap(
	creator string,
	index string,
	policy Policy,

) *MsgCreateMymap {
	return &MsgCreateMymap{
		Creator: creator,
		Index:   index,
		Policy:  &policy,
	}
}

func (msg *MsgCreateMymap) Route() string {
	return RouterKey
}

func (msg *MsgCreateMymap) Type() string {
	return TypeMsgCreateMymap
}

func (msg *MsgCreateMymap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMymap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMymap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMymap{}

func NewMsgUpdateMymap(
	creator string,
	index string,

) *MsgUpdateMymap {
	return &MsgUpdateMymap{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateMymap) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMymap) Type() string {
	return TypeMsgUpdateMymap
}

func (msg *MsgUpdateMymap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMymap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMymap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMymap{}

func NewMsgDeleteMymap(
	creator string,
	index string,

) *MsgDeleteMymap {
	return &MsgDeleteMymap{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteMymap) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMymap) Type() string {
	return TypeMsgDeleteMymap
}

func (msg *MsgDeleteMymap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMymap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMymap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
