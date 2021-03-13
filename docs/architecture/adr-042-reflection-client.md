# ADR 042: reflection client

## Changelog

- 13-03-2021: DRAFT

## Status

DRAFT

## Authors

- Frojdi Dymylja (fdymylja)

## Abstract

This ADR introduces a new cosmos-sdk client which can interact with any chain (queries and txs) in a codec and configuration agnostic way.

## Context

Currently, in any app built using the sdk, clients are created using the codec. The codec is an object which is specific to the application and its version. 

On top of that we have specific configurations (such as address prefix), which, with the current tooling, force the creation of clients to be reliant on application specific code.

The creation of different application clients, is something that can be addressed at compile time only.

For this reason building generalized (golang) tooling for the sdk, becomes troublesome especially if you want to support multiple chains and multiple versions.

To highlight possible use cases:

- One single rosetta instance could support multiple cosmos-sdk based applications (with different cosmos-sdk stargate versions)
- A single CLI capable of querying and sending txs to multiple applications
- Generalized tooling for chain-analysis (block explorers and so on)

To sum it up, in the sdk we already provide tx encoding and decoding endpoints, sending JSON via grpc-gateway, but we don't provide a library which can be used a single point of entry for any possible interaction with any possible chain.

This ADR aims to provide sdk-specific tooling that would allow for grpc-curl -like capabilities.


## Decision

To meet the needs highlighted above, we'd like to introduce the following (backwards-compatible) change:

- Extend the current [reflection](https://github.com/cosmos/cosmos-sdk/blob/master/proto/cosmos/base/reflection/v1beta1/reflection.proto) service in the sdk to allow better reflection capabilities for sdk specific codec interactions, a draft of the changes can be found [here](https://github.com/cosmos/cosmos-sdk/blob/fdymylja/cosmos-reflection/proto/cosmos/base/reflection/v1beta1/reflection.proto) (also check further discussion 1).

Note: the following API is based on protov2 API (which is already a dependency in the SDK, and what the SDK is planning to migrate to).

On top of this change, we'd like to add a new package called in client called `reflection` (position and name TBD). 

The `reflection` package would expose the following types:

- `reflection.Codec`: which is a codec capable of building itself at runtime. Its duty is to resolve and create JSON/Proto specific marshalers for the target chain. It can do that by asking the chain to provide the proto byte descriptors required to build a dynamic codec.
- `reflection.UnsignedTxBuilder`: its duty is to be able to facilitate the creation of application specific transactions, which are unsigned. 
  ```golang
    package reflection
    type UnsignedTxBuilder interface {
        SetMemo(memo string)
        SetGasLimit(limit uint64)
        ...
        AddSigner(signer SignerInfo)
        SignedBuilder() (SignedBuilder, error)
    }
  ```
- `reflection.SignedTxBuilder`: this object is returned by `reflection.UnsignedTxBuilder` and it serves the purpose of turning an unsigned transaction into a signed one, by providing information on the required signing bytes for a given pub key and setting the signatures.
    ```golang
        package reflection
        type SignedBuilder interface {
       	 BytesToSign(signer cryptotypes.PubKey) ([]byte, error)
         SetSignature(signer cryptotypes.PubKey, sig []byte) error
        }
    ```

- `reflection.Client`: this is the client that consumers will use to interact with any chain, by providing the gRPC and Tendermint endpoint. During its initialization it will build the codec and will provide TX and Query capabilities towards a chain. Multiple clients can be instantiated to interact with multiple chains at the same time.
    ```golang
        type Client interface {
            Query(ctx context.Context, method string, request proto.Message) (resp proto.Message, err error)
            QueryUnstructured(ctx context.Context, method string, request unstructured.Map) (resp proto.Message, err error)
            Tx(ctx context.Context, method string, msgs []proto.Message) (resp proto.Message, err error)
            TxUnstructured(ctx context.Context, method string, msgs []unstructured.Map) (resp proto.Message, error)
            
            ListQueries() []QueryDescriptor
            ListDeliverables() []DeliverableDescriptor
            
            Codec() Codec
        }
    ```
- `reflection.AccountInfoProvider`: this is an interface which provides account information, such as sequence, address (which are app configuration dependent) etc.   
- `unstructured.Map`: this object is a `map[string]interface` type which can marshal itself (recursively too if needed) into dynamic proto messages given a `protoreflect.MessageDescriptor`. Example interaction:
```golang
package example
func example() {
  resp, err := client.Tx(
    context.TODO(), "cosmos.bank.v1beta1.MsgSend",
    unstructured.Map{
      "from_address": "cosmos1ujtnemf6jmfm995j000qdry064n5lq854gfe3j",
      "to_address":   "cosmos1caa3es6q3mv8t4gksn9wjcwyzw7cnf5gn5cx7j",
      "amount": []unstructured.Map{
        {
          "denom":  "stake",
          "amount": "10",
        },
      },
    },
  )
}
```

## Consequences

### Backwards Compatibility

Backwards compatibility is fully respected. This ADR only adds a new cosmos-sdk client.

### Positive

- Introduction of a client which is non-compile-time reliant.
- Enables the construction of tooling which can be application independent.
- Allows for interactions between multiple chains and multiple versions of the chains as long as they are a post-stargate release.
- Abstracts away a lot of complexity that comes from building transactions (app specific address, app specific messages etc)

### Negative

- TxBuilder approach differs a little from the sdk default one. (Although it might be better UX wise)



## Further Discussions

1. Currently this ADR expects a change in the sdk reflection service, this shouldn't be needed as gRPC already offers a reflection service (needed to compute proto descriptor bytes, which in the end are the building pieces for our codec). Despite this, the reflection service appears not to be resolving types correctly. If anyone could provide a correct flow for this, it'd be better to use grpc reflection instead of sdk one.

## References

- {reference link}