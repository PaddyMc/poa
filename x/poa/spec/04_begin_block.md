# Begin-Block

At each ABCI `BeginBlock` call, the operations to update proposal are executed.

## Proposal changes

---

- Propose new validators to be added to the validator set
- Votes are calculated live as the chain progresses
- If the votes on a proposal to add a new validator are greater than 2/3 of the validator set, set the `Validator.Verified` value to true
- If `Proposal.StartBlock` + `Epoch` > the current block number purge the `Proposal`
