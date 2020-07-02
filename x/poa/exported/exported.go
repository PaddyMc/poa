package exported

import (
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ProposalI propsal for adding new validators to the consensus 
type ProposalI interface {
	GetValidatorAddr()	sdk.ValAddress
	GetStartBlock()		int32
	GetVotes()		int32
}

// ValidatorI expected validator functions
type ValidatorI interface {
	GetValidatorAddr()	sdk.ValAddress
	IsJailed()		bool
	IsVerified()		bool
	GetMoniker()		string
	GetWaitBlocks()		int32
	GetStatus()		sdk.BondStatus
	GetConsPubKey()		crypto.PubKey
	GetConsAddr()		sdk.ConsAddress
	GetConsensusPower()	int64
}

