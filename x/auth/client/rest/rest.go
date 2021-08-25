package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

// REST query and parameter values
const (
	MethodGet = "GET"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router, storeName string) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc(
		"/auth/accounts/{address}", QueryAccountRequestHandlerFn(storeName, clientCtx),
	).Methods(MethodGet)

	r.HandleFunc(
		"/auth/params",
		queryParamsHandler(clientCtx),
	).Methods(MethodGet)
}

// RegisterTxRoutes registers all transaction routes on the provided router.
func RegisterTxRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc("/txs/{hash}", QueryTxRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/txs", QueryTxsRequestHandlerFn(clientCtx)).Methods("GET")
}

// BroadcastReq defines a tx broadcasting request. The broadcast endpoint has
// been removed, but this type is still used in various places (e.g. multisign).
type BroadcastReq struct {
	Tx   legacytx.StdTx `json:"tx" yaml:"tx"`
	Mode string         `json:"mode" yaml:"mode"`
}

var _ codectypes.UnpackInterfacesMessage = BroadcastReq{}

// UnpackInterfaces implements the UnpackInterfacesMessage interface.
func (m BroadcastReq) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return m.Tx.UnpackInterfaces(unpacker)
}
