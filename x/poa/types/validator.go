package types

import (
//	"bytes"
//	"fmt"
//	"strings"
//	"time"
//
//	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
//	tmtypes "github.com/tendermint/tendermint/types"
//	yaml "gopkg.in/yaml.v2"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Add extra data to struct and propagate

type Validator struct {
	OperatorAddress         sdk.ValAddress           `json:"operator_address" yaml:"operator_address"`
	ConsPubKey              crypto.PubKey            `json:"consensus_pubkey" yaml:"consensus_pubkey"`
	Jailed                  bool                     `json:"jailed" yaml:"jailed"`
	Status                  sdk.BondStatus           `json:"status" yaml:"status"`
	Description             stakingtypes.Description `json:"description" yaml:"description"`
}

// NewValidator - create a new Validator
func NewValidator(operator sdk.ValAddress, pubKey crypto.PubKey, description stakingtypes.Description) Validator {
	return Validator{
		OperatorAddress:         operator,
		ConsPubKey:              pubKey,
		Jailed:                  false,
		Status:                  sdk.Bonded,
		Description:             description,
	}
}

// unmarshal a redelegation from a store value
func UnmarshalValidator(cdc *codec.Codec, value []byte) (validator Validator, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &validator)
	return validator, err
}

// this is a helper struct used for JSON de- and encoding only
type bechValidator struct {
	OperatorAddress         sdk.ValAddress           `json:"operator_address" yaml:"operator_address"`
	ConsPubKey              string                   `json:"consensus_pubkey" yaml:"consensus_pubkey"`
	Jailed                  bool                     `json:"jailed" yaml:"jailed"`
	Status                  sdk.BondStatus           `json:"status" yaml:"status"`
	Description             stakingtypes.Description `json:"description" yaml:"description"`
}

// MarshalJSON marshals the validator to JSON using Bech32
func (v Validator) MarshalJSON() ([]byte, error) {
	bechConsPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, v.ConsPubKey)
	if err != nil {
		return nil, err
	}

	return codec.Cdc.MarshalJSON(bechValidator{
		OperatorAddress:         v.OperatorAddress,
		ConsPubKey:              bechConsPubKey,
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		Description:             v.Description,
	})
}

// UnmarshalJSON unmarshals the validator from JSON using Bech32
func (v *Validator) UnmarshalJSON(data []byte) error {
	bv := &bechValidator{}
	if err := codec.Cdc.UnmarshalJSON(data, bv); err != nil {
		return err
	}
	consPubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, bv.ConsPubKey)
	if err != nil {
		return err
	}
	*v = Validator{
		OperatorAddress:         bv.OperatorAddress,
		ConsPubKey:              consPubKey,
		Jailed:                  bv.Jailed,
		Status:                  bv.Status,
		Description:             bv.Description,
	}
	return nil
}

// return the TM validator address
func (v Validator) ConsAddress() sdk.ConsAddress {
	return sdk.ConsAddress(v.ConsPubKey.Address())
}


// get the consensus-engine power
func (v Validator) ConsensusPower() int64 {
	return 1//v.PotentialConsensusPower()
}

// UpdateStatus updates the location of the shares within a validator
func (v Validator) UpdateStatus(newStatus sdk.BondStatus) Validator {
	v.Status = newStatus
	return v
}

// Validators is a collection of Validator
type Validators []Validator

func (v Validator) IsJailed() bool               { return v.Jailed }
func (v Validator) GetMoniker() string           { return v.Description.Moniker }
func (v Validator) GetStatus() sdk.BondStatus    { return v.Status }
func (v Validator) GetOperator() sdk.ValAddress  { return v.OperatorAddress }
func (v Validator) GetConsPubKey() crypto.PubKey { return v.ConsPubKey }
func (v Validator) GetConsAddr() sdk.ConsAddress { return sdk.ConsAddress(v.ConsPubKey.Address()) }
func (v Validator) GetConsensusPower() int64     { return v.ConsensusPower() }

// String returns a human readable string representation of a validator.
//func (v Validator) String() string {
//	bechConsPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, v.ConsPubKey)
//	if err != nil {
//		panic(err)
//	}
//	return fmt.Sprintf(`Validator
// 			    Operator Address:           %s
// 			    Validator Consensus Pubkey: %s
// 			    Jailed:                     %v
// 			    Status:                     %s
// 			    Description:                %s`,
//			    v.OperatorAddress, bechConsPubKey,
//			    v.Jailed, v.Status,
//			    v.Description
//			   )
//}
