package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
        exported "github.com/PaddyMc/poa/x/poa/exported"
)

// ParamSubspace defines the expected Subspace interfacace
type ParamSubspace interface {
	WithKeyTable(table params.KeyTable) params.Subspace
	Get(ctx sdk.Context, key []byte, ptr interface{})
	GetParamSet(ctx sdk.Context, ps params.ParamSet)
	SetParamSet(ctx sdk.Context, ps params.ParamSet)
}

// ValidatorSet expected properties for the set of all validators
type ValidatorSet interface {
        //IterateValidators(sdk.Context,
        //        func(index int64, validator exported.ValidatorI) (stop bool))
        //IterateLastBlockValidators(sdk.Context,
        //        func(index int64, validator exported.ValidatorI) (stop bool))
        GetValidator(sdk.Context, sdk.ValAddress) (found bool)
	//Jail(sdk.Context, sdk.ConsAddress)
        //Unjail(sdk.Context, sdk.ConsAddress)
        //MaxValidators(sdk.Context) uint16
}

// Proposal expected properties for the proposals
type Proposal interface {
        IterateProposal(sdk.Context,
                func(index int64, validator exported.ValidatorI) (stop bool))
}

// PoaHooks event hooks for poa validator object (noalias)
type PoaHooks interface {
        AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress)
        AfterValidatorEdited(ctx sdk.Context, valAddr sdk.ValAddress)
        AfterValidatorRemoved(ctx sdk.Context, valAddr sdk.ValAddress)
        AfterVoteCast(ctx sdk.Context, valAddr sdk.ValAddress)
}
