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
1. **[State Transitions](02_state_transitions.md)**
1. **[Messages](03_messages.md)**
1. **[Begin-Block](04_begin_block.md)**
1. **[End-Block](05_end_block.md)**
1. **[05_hooks](06_hooks.md)**
1. **[Events](07_events.md)**
1. **[Parameters](08_parameters.md)**
