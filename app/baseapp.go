package app

import (
	"fmt"
	"reflect"

	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/container"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AppName string

type baseappInputs struct {
	Name      AppName       `optional:"true"`
	Logger    log.Logger    `optional:"true"`
	DB        dbm.DB        `optional:"true"`
	TxDecoder sdk.TxDecoder `optional:"true"`
	Options   []func(*baseapp.BaseApp)
}

var nilOpt func(*baseapp.BaseApp)

var BaseAppProvider = container.Options(
	container.AutoGroupTypes(reflect.TypeOf(nilOpt)),
	container.Provide(
		func(inputs baseappInputs) *baseapp.BaseApp {
			name := inputs.Name
			if name == "" {
				name = "simapp"
			}

			logger := inputs.Logger
			if logger == nil {
				logger = log.NewNopLogger()
			}

			db := inputs.DB
			if db == nil {
				db = dbm.NewMemDB()
			}

			txDecoder := inputs.TxDecoder
			if txDecoder == nil {
				txDecoder = func(txBytes []byte) (sdk.Tx, error) {
					return nil, fmt.Errorf("no TxDecoder, can't decode transactions")
				}
			}

			return baseapp.NewBaseApp(string(name), logger, db, txDecoder, inputs.Options...)
		},
	))
