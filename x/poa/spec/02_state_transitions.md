# State Transitions

This document describes the state transition operations pertaining to:

1. [Validators](./02_state_transitions.md#validators)
2. [Proposals](./02_state_transitions.md#proposals)


## Validators

---

State transitions in validators are performed on every `EndBlock` and are used update the active `ValidatorSet`.

### `INTURN` **to** `NOTURN`

After a validator has proposed and signed a block the following operations occur:

- set `validator.Status` to `NOTURN`
- set `validator.WaitBlocks` to `Floor(VALIDATOR_COUNT / 2) + 1`
- update the `Validator` object for this validator in the store

### `NOTURN` **to** `INTURN`

When a validator can propose and sign blocks:

- set `validator.Status` to `INTURN`
- update the `Validator` object for this validator

### `Jail` **||** `Unjail`

When a validator is jailed it is effectively removed from the Tendermint set. This process may be also be reversed. The following operations occur:

- set `Validator.Jailed` to true | false
- update the `Validator` object for this validator

### `Verified` **||** `Unverfied`

When a validator is verified it is added the Tendermint set. The following operations occur:

- set `Validator.Verified` to true
- update the `Validator` object for this validator

## Proposals

---

State transitions for proposals are performed on every `BeginBlock` and are used update the active `ValidatorSet`.

### `Votes ++`

When a vote has been cast in favour of a validator:

- Increment `Proposal.Votes` by 1

### `Votes --`

When a vote has been cast against a new validator:

- Decrement `Proposal.Votes` by 1
