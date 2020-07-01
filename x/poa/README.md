# Proof Of Authority module for the cosmos-sdk

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

## For more info navigate to the /spec folder
