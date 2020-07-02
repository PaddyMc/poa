package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Poa params default values
const (
	// Epoch reflects number of blocks until voting resets
	DefaultEpoch uint16 = 30000

	// Default maximum number of bonded validators
	DefaultMaxValidators uint16 = 100
)

// nolint - Keys for parameter access
var (
	KeyEpoch		= []byte("Epoch")
	KeyMaxValidators	= []byte("MaxValidators")
)

var _ params.ParamSet = (*Params)(nil)

// Params defines the high level settings for poa module
type Params struct {
	Epoch		    uint16	  `json:"epoch" yaml:"epoch"`
	MaxValidators       uint16        `json:"max_validators" yaml:"max_validators"`               // maximum number of validators (max uint16 = 65535)
}

// NewParams creates a new Params instance
func NewParams(epoch uint16, maxValidators uint16) Params {
	return Params{
		Epoch:		     epoch,
		MaxValidators:       maxValidators,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyEpoch, &p.Epoch, validateInt),
		params.NewParamSetPair(KeyMaxValidators, &p.MaxValidators, validateInt),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultEpoch, DefaultMaxValidators)
}

func validateInt(i interface{}) error {
	v, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}
