# Hooks

The following hooks can be registered with the `poa` module:

1. `AfterValidatorCreated(Context, ValAddress)`

   - called when a validator is created

1. `AfterValidatorEdited(Context, ValAddress)`

   - called when a validator's state is changed

1. `AfterValidatorRemoved(Context, ValAddress)`

   - called when a validator is deleted

1. `AfterVoteCast(Context, ValAddress)`
   - called when a validator casts a vote
