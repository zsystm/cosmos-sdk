package ormgraphql

import (
	"encoding/json"
	"testing"

	"github.com/graphql-go/graphql"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
)

func TestMessage(t *testing.T) {
	b := NewBuilder()
	a := &testpb.A{
		U32:      14,
		U64:      0,
		Str:      "",
		Bz:       nil,
		Ts:       nil,
		Dur:      nil,
		I32:      10,
		S32:      0,
		Sf32:     0,
		I64:      0,
		S64:      0,
		Sf64:     0,
		F32:      0,
		F64:      0,
		B:        true,
		E:        0,
		Repeated: nil,
		Map:      nil,
		Msg:      nil,
		Sum:      nil,
	}
	obj, err := b.protoMessageToGraphqlObject(a.ProtoReflect().Descriptor())
	assert.NilError(t, err)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"a": &graphql.Field{
				Type: obj,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(map[string]interface{})["a"], nil
				},
			},
		},
	}),
	})
	assert.NilError(t, err)

	query := `{ a { 
i32
s32
sf32
u32
f32
u64
f64
b
} }`
	res := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		RootObject:    map[string]interface{}{"a": a.ProtoReflect()},
	})
	assert.Equal(t, 0, len(res.Errors), res.Errors)
	bz, err := json.Marshal(res.Data)
	assert.NilError(t, err)
	assert.Equal(t, ``, string(bz))
}
