package module

import (
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/genesis"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UseCoreAPIModule wraps the core API module as an AppModule that this version
// of the SDK can use.
func UseCoreAPIModule(name string, module appmodule.AppModule) AppModule {
	return coreAppModuleAdapator{
		name:   name,
		module: module,
	}
}

type coreAppModuleAdapator struct {
	name   string
	module appmodule.AppModule
}

func (c coreAppModuleAdapator) Name() string {
	return c.name
}

func (c coreAppModuleAdapator) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	if mod, ok := c.module.(interface {
		RegisterLegacyAminoCodec(amino *codec.LegacyAmino)
	}); ok {
		mod.RegisterLegacyAminoCodec(amino)
	}
}

func (c coreAppModuleAdapator) RegisterInterfaces(registry types.InterfaceRegistry) {
	if mod, ok := c.module.(interface {
		RegisterInterfaces(registry types.InterfaceRegistry)
	}); ok {
		mod.RegisterInterfaces(registry)
	}
}

func (c coreAppModuleAdapator) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	if mod, ok := c.module.(appmodule.HasGenesis); ok {
		target := genesis.RawJSONTarget{}
		err := mod.DefaultGenesis(target.Target())
		if err != nil {
			panic(err)
		}

		rawJson, err := target.JSON()
		if err != nil {
			panic(err)
		}

		return rawJson
	}

	return nil
}

func (c coreAppModuleAdapator) ValidateGenesis(codec codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {
	if mod, ok := c.module.(appmodule.HasGenesis); ok {
		source, err := genesis.SourceFromRawJSON(message)
		if err != nil {
			return err
		}

		return mod.ValidateGenesis(source)
	}

	return nil
}

func (c coreAppModuleAdapator) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	if mod, ok := c.module.(interface {
		RegisterRESTRoutes(context client.Context, router *mux.Router)
	}); ok {
		mod.RegisterRESTRoutes(context, router)
	}
}

func (c coreAppModuleAdapator) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	if mod, ok := c.module.(interface {
		RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux)
	}); ok {
		mod.RegisterGRPCGatewayRoutes(context, mux)
	}
}

func (c coreAppModuleAdapator) GetTxCmd() *cobra.Command {
	if mod, ok := c.module.(interface {
		GetTxCmd() *cobra.Command
	}); ok {
		return mod.GetTxCmd()
	}

	return nil
}

func (c coreAppModuleAdapator) GetQueryCmd() *cobra.Command {
	if mod, ok := c.module.(interface {
		GetQueryCmd() *cobra.Command
	}); ok {
		return mod.GetQueryCmd()
	}

	return nil
}

func (c coreAppModuleAdapator) InitGenesis(context sdk.Context, codec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	if mod, ok := c.module.(appmodule.HasGenesis); ok {
		source, err := genesis.SourceFromRawJSON(message)
		if err != nil {
			panic(err)
		}
		err = mod.InitGenesis(sdk.WrapSDKContext(context), source)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (c coreAppModuleAdapator) ExportGenesis(context sdk.Context, codec codec.JSONCodec) json.RawMessage {
	if mod, ok := c.module.(appmodule.HasGenesis); ok {
		target := genesis.RawJSONTarget{}
		err := mod.ExportGenesis(sdk.WrapSDKContext(context), target.Target())
		if err != nil {
			panic(err)
		}

		rawJson, err := target.JSON()
		if err != nil {
			panic(err)
		}

		return rawJson
	}

	return nil
}

func (c coreAppModuleAdapator) RegisterInvariants(registry sdk.InvariantRegistry) {
	if mod, ok := c.module.(interface {
		RegisterInvariants(registry sdk.InvariantRegistry)
	}); ok {
		mod.RegisterInvariants(registry)
	}
}

func (c coreAppModuleAdapator) Route() sdk.Route {
	if mod, ok := c.module.(interface {
		Route() sdk.Route
	}); ok {
		return mod.Route()
	}

	return sdk.Route{}
}

func (c coreAppModuleAdapator) QuerierRoute() string {
	if mod, ok := c.module.(interface {
		QuerierRoute() string
	}); ok {
		return mod.QuerierRoute()
	}

	return ""
}

func (c coreAppModuleAdapator) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (c coreAppModuleAdapator) RegisterServices(c2 Configurator) {
	if mod, ok := c.module.(interface {
		RegisterServices(c2 Configurator)
	}); ok {
		mod.RegisterServices(c2)
	}
}

func (c coreAppModuleAdapator) ConsensusVersion() uint64 {
	if mod, ok := c.module.(interface {
		ConsensusVersion() uint64
	}); ok {
		return mod.ConsensusVersion()
	}

	return 0
}
