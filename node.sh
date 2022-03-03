#!/usr/bin/env bash
set -e

# This script is for running a local node, mostly for testing and demo purposes.
# It creates a chain with one validator, called `ALICE`. The native token of
# this chain is "stake", and `ALICE`, being the only existing account in the
# chain at genesis, has lots of them.


make build

# Set variables
CFG_DIR=~/.simapp
BUILD_CMD=./build/simd
VALIDATOR=alice
CHAIN_ID=my-chain

# Cleanup previous installations, if any.
rm -rf $CFG_DIR

# Add ALICE's key to the keyring.
$BUILD_CMD keys add $VALIDATOR --keyring-backend test
VALIDATOR_ADDRESS=$($BUILD_CMD keys show $VALIDATOR -a --keyring-backend test)

echo "-----------"
echo $VALIDATOR_ADDRESS
echo "------------"

# Initialize the genesis file. It is available under $CFG_DIR/config/genesis.json.
$BUILD_CMD init $VALIDATOR --chain-id $CHAIN_ID
$BUILD_CMD add-genesis-account $VALIDATOR_ADDRESS 1000000000000stake
$BUILD_CMD gentx $VALIDATOR 1000000000stake --keyring-backend test --chain-id $CHAIN_ID
$BUILD_CMD collect-gentxs

# Run the node.
$BUILD_CMD start --minimum-gas-prices 0.0000001stake --mode validator
