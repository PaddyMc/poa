## Hooks

The following hooks can be registered with the `poa` module:

- `AfterValidatorCreated(Context, ValAddress)`
    - called when a validator is created
- `AfterValidatorEdited(Context, ValAddress)`
    - called when a validator's state is changed
- `AfterValidatorRemoved(Context, ValAddress)`
    - called when a validator is deleted
- `AfterVoteCast(Context, ValAddress)`
    - called when a validator casts a vote
