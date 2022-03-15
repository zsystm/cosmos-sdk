package ormtable_test

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"

	"github.com/regen-network/gocuke"
	"google.golang.org/grpc/status"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
)

func TestFeatures(t *testing.T) {
	gocuke.NewRunner(t, &suite{}).Run()
}

type suite struct {
	gocuke.TestingT
	table ormtable.Table
	ctx   context.Context
	err   error
}

func (s *suite) Before() {
	var err error
	s.table, err = ormtable.Build(ormtable.Options{
		MessageType: (&testpb.ExampleTable{}).ProtoReflect().Type(),
	})
	assert.NilError(s, err)
	s.ctx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
}

func (s *suite) AnExistingEntity(docString gocuke.DocString) {
	existing := s.exampleTableFromDocString(docString)
	assert.NilError(s, s.table.Insert(s.ctx, existing))
}

func (s suite) exampleTableFromDocString(docString gocuke.DocString) *testpb.ExampleTable {
	ex := &testpb.ExampleTable{}
	assert.NilError(s, protojson.Unmarshal([]byte(docString.Content), ex))
	return ex
}

func (s *suite) IInsert(docString gocuke.DocString) {
	ex := s.exampleTableFromDocString(docString)
	s.err = s.table.Insert(s.ctx, ex)
}

func (s *suite) ExpectAError(a string) {
	assert.ErrorIs(s, s.err, s.toError(a))
}

func (s *suite) toError(str string) error {
	switch str {
	case "already exists":
		return ormerrors.AlreadyExists
	default:
		s.Fatalf("missing case for error %s", str)
		return nil
	}
}

func (s *suite) ExpectGrpcErrorCode(a string) {
	assert.Equal(s, a, status.Code(s.err).String())
}
