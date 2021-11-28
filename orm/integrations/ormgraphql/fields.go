package ormgraphql

import (
	"github.com/iancoleman/strcase"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

func fieldsCamelCase(fields ormkv.Fields) {
	strcase.ToCamel(fields.String())
}
