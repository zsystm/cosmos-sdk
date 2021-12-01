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
		U64:      3,
		Str:      "abc",
		Bz:       []byte{4, 7, 9},
		Ts:       nil,
		Dur:      nil,
		I32:      10,
		S32:      -3,
		Sf32:     -5,
		I64:      11,
		S64:      7,
		Sf64:     -7,
		F32:      1,
		F64:      2,
		B:        true,
		E:        testpb.Enum_ENUM_FIVE,
		Repeated: []uint32{0, 1, 2, 4},
		Map:      nil,
		Msg: &testpb.B{
			X: "xyz",
		},
		Sum: &testpb.A_Oneof{Oneof: 3},
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
repeated
str
msg {
  x
}
oneof
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
