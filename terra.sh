#!/usr/bin/env bash
set -e

# Run a node in another terminal, for populating live accSeq/accNum.

# Set variables
CFG_DIR=~/.simapp
BUILD_CMD=./build/simd
ALICE=alice
BOB=bob
MULTI=multi_alice_bob
CHAIN_ID=my-chain

COMMON_FLAGS="--keyring-backend test --chain-id $CHAIN_ID"
ALICE_ADDRESS=$($BUILD_CMD keys show $ALICE -a --keyring-backend test)
$BUILD_CMD keys add --multisig alice $MULTI --keyring-backend test
MULTI_ADDRESS2=$($BUILD_CMD keys show $MULTI -a --keyring-backend test)
# MULTI_ADDRESS="cosmos1hd6fsrvnz6qkp87s3u86ludegq97agxsdkwzyh" # dummy address, to simulate a multisig that's not in the keyring.
MULTI_ADDRESS="cosmos17m78ka80fqkkw2c4ww0v4xm5nsu2drgrlm8mn2" # dummy address, to simulate a multisig that's not in the keyring.

# Create unsigned tx
echo "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"$MULTI_ADDRESS\",\"to_address\":\"$ALICE_ADDRESS\",\"amount\":[{\"denom\":\"stake\",\"amount\":\"10\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[],\"fee\":{\"amount\":[],\"gas_limit\":\"200000\",\"payer\":\"\",\"granter\":\"\"},\"tip\":null},\"signatures\":[]}" > unsigned.json

# echo "=============="
# echo $ALICE_ADDRESS
# echo $MULTI_ADDRESS
# echo $MULTI_ADDRESS2
# echo "=============="

$BUILD_CMD tx sign unsigned.json --multisig $MULTI_ADDRESS $COMMON_FLAGS --from $ALICE

# echo "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"$MULTI_ADDRESS2\",\"to_address\":\"$ALICE_ADDRESS\",\"amount\":[{\"denom\":\"stake\",\"amount\":\"10\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[],\"fee\":{\"amount\":[],\"gas_limit\":\"200000\",\"payer\":\"\",\"granter\":\"\"},\"tip\":null},\"signatures\":[]}" > unsigned.json

# $BUILD_CMD tx sign unsigned.json --multisig $MULTI $COMMON_FLAGS --from $ALICE
