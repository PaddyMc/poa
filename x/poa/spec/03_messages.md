## Messages

### **MsgCreateValidator**

---

A validator is created using the `MsgCreateValidator` message

```go
type MsgCreateValidator struct {
    Description            Description
    ValidatorAddr          sdk.ValAddress
    PubKey                 crypto.PubKey
}
```

This message is expected to fail if:

- Another validator with this operator address is already registered
- Another validator with this `PubKey` is already registered
- The description fields are too large

This message creates and stores the `Validator` object at appropriate indexes. The validator then sends the MsgCastVote message for the validator.

### **MsgEditValidator**

---

The `Description` of a validator can be updated using the `MsgEditValidator`.

```go
type MsgEditValidator struct {
    Description            Description
    ValidatorAddr          sdk.ValAddress
}
```

This message is expected to fail if:

- The description fields are too large

This message stores the updated `Validator` object.

### **MsgRemoveValidator**

---

The `Description` of a validator can be updated using the `MsgEditCandidacy`.

```go
type MsgRemoveValidator struct {
    ValidatorAddr          sdk.ValAddress
}
```

This message is expected to fail if:

- The `Validator` object cannot be found

This message removes the `Validator` object.

### **MsgCastVote**

---

A `Vote` can be cast as a validator using the `MsgCastVote`.

```go
type MsgCastVote struct {
    ValidatorAddr          sdk.ValAddress
    inFavor                bool
}
```

This updates the `Proposal` object and creates one if it does not exist. If `inFavor` is true, increment the `Proposal.Votes` value, if false decrement the `Proposal.Votes` value.
