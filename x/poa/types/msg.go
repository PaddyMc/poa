package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//const RouterKey = ModuleName

// Route should return the name of the module
func (msg MsgCreateValidator) Route() string { return ModuleName }

// Type should return the action
func (msg MsgCreateValidator) Type() string { return "create_validator" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateValidator) ValidateBasic() error {
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateValidator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgCreateValidator defines a SetName message
type MsgCreateValidator struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgCreateValidator is a constructor function for MsgCreateValidator
func NewMsgCreateValidator(name string, owner sdk.AccAddress) MsgCreateValidator {
	return MsgCreateValidator{
		Name:  name,
		Owner: owner,
	}
}
