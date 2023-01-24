package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _, _, _, _, _ sdk.Msg = &MsgInitWithdraw{}, &MsgSendMySig{}, &MsgSignAgain{}, &MsgSyncFarming{}, &MsgSetLiquidityCap{}

func NewMsgInitWithdraw(creator string) *MsgInitWithdraw {
	return &MsgInitWithdraw{
		Creator: creator,
	}
}

func (msg *MsgInitWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgInitWithdraw) Type() string {
	return "InitWithdraw"
}

func (msg *MsgInitWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgSendMySig) Route() string {
	return RouterKey
}

func (msg *MsgSendMySig) Type() string {
	return "SendMySig"
}

func (msg *MsgSendMySig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendMySig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendMySig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgSignAgain) Route() string {
	return RouterKey
}

func (msg *MsgSignAgain) Type() string {
	return "SignAgain"
}

func (msg *MsgSignAgain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignAgain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignAgain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgUpdateLatestSigners(creator string) *MsgUpdateLatestSigners {
	return &MsgUpdateLatestSigners{
		Creator: creator,
	}
}

func (msg *MsgUpdateLatestSigners) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLatestSigners) Type() string {
	return "UpdateLatestSigners"
}

func (msg *MsgUpdateLatestSigners) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLatestSigners) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLatestSigners) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSyncFarming(
	lpAddress string,
	chainId uint64,
	tokenAddress string,
	creator string) *MsgSyncFarming {
	return &MsgSyncFarming{
		LpAddress:    lpAddress,
		ChainId:      chainId,
		TokenAddress: tokenAddress,
		Creator:      creator,
	}
}

func (msg *MsgSyncFarming) Route() string {
	return RouterKey
}

func (msg *MsgSyncFarming) Type() string {
	return "SyncFarming"
}

func (msg *MsgSyncFarming) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSyncFarming) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSyncFarming) ValidateBasic() error {
	if msg.LpAddress == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid LP address")
	}
	if msg.TokenAddress == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid token address")
	}
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgTriggerSetRefund) Route() string {
	return RouterKey
}

func (msg *MsgTriggerSetRefund) Type() string {
	return "TriggerSetRefund"
}

func (msg *MsgTriggerSetRefund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTriggerSetRefund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerSetRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgSetLiquidityCap) Route() string {
	return RouterKey
}

func (msg *MsgSetLiquidityCap) Type() string {
	return "SetLiquidityCap"
}

func (msg *MsgSetLiquidityCap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetLiquidityCap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetLiquidityCap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
