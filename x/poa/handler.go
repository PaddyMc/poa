package poa

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the poa type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName,  msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handle<Action> does x
func handleMsgCreateValidator(ctx sdk.Context, msg MsgCreateValidator, k Keeper) (*sdk.Result, error) {
	//if _, found := k.GetValidator(ctx, msg.Name); found {
	//	return nil, stakingtypes.ErrValidatorOwnerExists
	//}

	//if _, found := k.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(msg.PubKey)); found {
	//	return nil, stakingtypes.ErrValidatorPubKeyExists
	//}

	//if _, err := msg.Description.EnsureLength(); err != nil {
	//	return nil, err
	//}

	//if ctx.ConsensusParams() != nil {
	//	tmPubKey := tmtypes.TM2PB.PubKey(msg.PubKey)
	//	if !tmstrings.StringInSlice(tmPubKey.Type, ctx.ConsensusParams().Validator.PubKeyTypes) {
	//		return nil, sdkerrors.Wrapf(stakingtypes.ErrValidatorPubKeyTypeNotSupported,
	//			"got: %s, valid: %s", tmPubKey.Type, ctx.ConsensusParams().Validator.PubKeyTypes,
	//		)
	//	}
	//}

	//validator := NewValidator(msg.ValidatorAddress, msg.PubKey, msg.Description)

	//k.SetValidator(ctx, validator)
	//k.SetValidatorByConsAddr(ctx, validator)
	//k.SetNewValidatorByPowerIndex(ctx, validator)

	//// call the after-creation hook
	//k.AfterValidatorCreated(ctx, validator.OperatorAddress)

	//ctx.EventManager().EmitEvents(sdk.Events{
	//	sdk.NewEvent(
	//		EventTypeCreateValidator,
	//		sdk.NewAttribute(AttributeKeyValidator, msg.ValidatorAddress.String()),
	//	),
	//	sdk.NewEvent(
	//		sdk.EventTypeMessage,
	//		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
	//	),
	//})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

