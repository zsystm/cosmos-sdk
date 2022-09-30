module cosmossdk.io/api/cosmos/authn

go 1.19

require (
	cosmossdk.io/api v0.2.1 // indirect
	github.com/cosmos/cosmos-sdk/orm v1.0.0-alpha.12 // indirect
)

replace (
	cosmossdk.io/api => ../..
	github.com/cosmos/cosmos-sdk/orm => ../../../orm
)
