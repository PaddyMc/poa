# State

## Params

---

Params is a module-wide configuration structure that stores system parameters and defines the overall functioning of the `poa` module.

- Params: `Paramsspace("poa") -> amino(Params)`

```go
type Params struct {
    Epoch               uint16              // Number of block to wait until voting is reset
    MaxValidators       uint16              // Max number of validators in the consensus
}
```

## Validator

---

Validators objects should be primarily stored and accessed by the `OperatorAddr`, an SDK validator address for the operator of the validator.

- Validators: `0x21 | OperatorAddr -> amino(Validator)`

Validators can have one of two statuses:

1. `INTURN`: The validator can propose blocks
2. `NOTURN`: The validator cannot propose blocks

Validators cannot sign blocks for `Floor(VALIDATOR_COUNT / 2) + 1` blocks after signing a block.

Each validator's state is stored in a Validator struct:

```go
type Validator struct {
    OperatorAddress         sdk.ValAddress 
    ConsPubKey              crypto.PubKey   // the consensus public key
    Status                  sdk.Status      // validator status (INTURN/NOTURN)
    WaitBlocks              int32           // wait for number of blocks before being allowed to create a new block
    Verified                bool            // has been voted to join the validator set
    Jailed                  bool
    Description             Description
}

type Description struct {
    Moniker          string                 // name
    Identity         string                 // optional identity signature
    Website          string
    SecurityContact  string
    Details          string
}
```

## Proposal

---

Proposal objects should be primarily stored and accessed by the `OperatorAddr`

- Proposal: `0x22 | OperatorAddr -> amino(Proposal)`

Each proposal state is stored in a Proposal struct:

```go
type Proposal struct {
    OperatorAddress         sdk.ValAddress   
    StartBlock              uint32          // Block number the Proposal started on
    Votes                   uint32          // Number of votes on a Proposal
}
```
