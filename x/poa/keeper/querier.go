package keeper

import (
//	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

//	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/PaddyMc/poa/x/poa/types"
)

const (
	QueryValidator = "validator"
	QueryParams    = "params"
)

// NewQuerier creates a new querier for poa clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryValidator:
			return queryValidator(ctx, req, k)

		case QueryParams:
			return queryParams(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown poa query endpoint")
		}
	}
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	//var params types.QueryValidatorParams

	//err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	//if err != nil {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	//}

	//validator, found := k.GetValidator(ctx, params.ValidatorAddr)
	//if !found {
	//	return nil, stakingtypes.ErrNoValidatorFound
	//}

	//res, err := codec.MarshalJSONIndent(types.ModuleCdc, validator)
	//if err != nil {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	//}

	return nil,nil
}


func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}


