# Sample POA Cosmos Blockchain

## To view the spec

**[Link to Spec](./x/poa/spec/README.md)**

## To build

```sh
go mod tidy
make install
```

## To setup

```sh
appd init appd
appcli keys add validator
appd add-genesis-account validator 100000000stake
appd gentx --name validator
appd collect-gentxs
```

## To run

### CLI

```sh
appcli  -h
```

### Daemon

```sh
appd start
```

