package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _, _, _, _, _, _, _ sdk.Msg = &MsgSignMint{}, &MsgTriggerSignMint{}, &MsgSignWithdraw{},
	&MsgTriggerSignWithdraw{}, &MsgClaimFee{}, &MsgMigrateVault{}, &MsgResetTotalSupply{}

func NewMsgSignMint(sender string) *MsgSignMint {
	return &MsgSignMint{
		Sender: sender,
	}
}

func (msg *MsgSignMint) Route() string {
	return RouterKey
}

func (msg *MsgSignMint) Type() string {
	return "SignMint"
}

func (msg *MsgSignMint) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSignMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgTriggerSignMint(sender string) *MsgTriggerSignMint {
	return &MsgTriggerSignMint{
		Sender: sender,
	}
}

func (msg *MsgTriggerSignMint) Route() string {
	return RouterKey
}

func (msg *MsgTriggerSignMint) Type() string {
	return "TriggerSignMint"
}

func (msg *MsgTriggerSignMint) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgTriggerSignMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerSignMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgSignWithdraw(sender string) *MsgSignWithdraw {
	return &MsgSignWithdraw{
		Sender: sender,
	}
}

func (msg *MsgSignWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgSignWithdraw) Type() string {
	return "SignWithdraw"
}

func (msg *MsgSignWithdraw) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSignWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgTriggerSignWithdraw(sender string) *MsgTriggerSignWithdraw {
	return &MsgTriggerSignWithdraw{
		Sender: sender,
	}
}

func (msg *MsgTriggerSignWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgTriggerSignWithdraw) Type() string {
	return "TriggerSignWithdraw"
}

func (msg *MsgTriggerSignWithdraw) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgTriggerSignWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerSignWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgClaimFee(
	delAddress string, chainId uint64, tokenAddress string, nonce uint64, sender string) *MsgClaimFee {
	return &MsgClaimFee{
		DelegatorAddress: delAddress,
		ChainId:          chainId,
		TokenAddress:     tokenAddress,
		Nonce:            nonce,
		Sender:           sender,
	}
}

func (msg *MsgClaimFee) Route() string {
	return RouterKey
}

func (msg *MsgClaimFee) Type() string {
	return "ClaimFee"
}

func (msg *MsgClaimFee) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClaimFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimFee) ValidateBasic() error {
	if msg.DelegatorAddress == "" && !msg.IsValidator {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgClaimRefund(sender string) *MsgClaimRefund {
	return &MsgClaimRefund{
		Sender: sender,
	}
}

func (msg *MsgClaimRefund) Route() string {
	return RouterKey
}

func (msg *MsgClaimRefund) Type() string {
	return "ClaimRefund"
}

func (msg *MsgClaimRefund) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClaimRefund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgMigrateVault(chainId uint64, token, amount, toAddress, sender string) *MsgMigrateVault {
	return &MsgMigrateVault{
		ChainId:   chainId,
		Token:     token,
		Amount:    amount,
		ToAddress: toAddress,
		Sender:    sender,
	}
}

func (msg *MsgMigrateVault) Route() string {
	return RouterKey
}

func (msg *MsgMigrateVault) Type() string {
	return "MigrateVault"
}

func (msg *MsgMigrateVault) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgMigrateVault) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMigrateVault) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetSender())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgResetTotalSupply(sender string) *MsgResetTotalSupply {
	return &MsgResetTotalSupply{
		Sender: sender,
	}
}

func (msg *MsgResetTotalSupply) Route() string {
	return RouterKey
}

func (msg *MsgResetTotalSupply) Type() string {
	return "ResetTotalSupply"
}

func (msg *MsgResetTotalSupply) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgResetTotalSupply) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResetTotalSupply) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
