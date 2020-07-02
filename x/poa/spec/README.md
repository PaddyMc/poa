# poa module specification

## Abstract

This paper specifies a `poa` module of the Cosmos-SDK. The `poa` module will define a Proof-Of-Authority (POA) system that will allow the Cosmos-SDK to leverage a POA algorithm for building blockchains.

The `poa` module will have the following characteristics:

- Trusted set of validators on a network
- Produces blocks at a reliable rate
- Not as many adversarial conditions on the network
- Potentially higher performance than other consensus algorithms

Two main data structures are used to allow the `poa` module to function correctly:

- `Validator`: Represents consensus validators that are used to create and confirm blocks
- `Proposal`: Allows verified validators to vote if new validators can join the consensus

To allow for easy development and avoid common pitfalls of defining a new algorithm a modified version of the Clique POA algorithm will be used.

## Contents

1. **[State](01_state.md)**
    - [Params](01_state.md#params)
    - [Validator](01_state.md#validator)
    - [Proposal](01_state.md#proposal)
1. **[State Transitions](02_state_transitions.md)**
    - [Validator](02_state_transitions.md#validator)
    - [Proposal](02_state_transitions.md#proposal)
1. **[Messages](03_messages.md)**
    - [MsgCreateValidator](03_messages.md#MsgCreateValidator)
    - [MsgEditValidator](03_messages.md#MsgEditValidator)
    - [MsgRemoveValidator](03_messages.md#MsgRemoveValidator)
    - [MsgCastVote](03_messages.md#MsgCastVote)
1. **[Begin-Block](04_begin_block.md)**
    - [Proposal changes](04_begin_block#Proposal-changes)
1. **[End-Block](05_end_block.md)**
    - [Validator set changes](04_begin_block#Validator-set-changes)
1. **[Hooks](06_hooks.md)**
1. **[Events](07_events.md)**
   - [EndBlocker](07_events.md#endblocker)
   - [Handlers](07_events.md#handlers)
1. **[Parameters](08_parameters.md)**
