package ormtable_test

import (
	"context"
	"testing"

	"github.com/aaronc/gocuke"
	"google.golang.org/grpc/status"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
)

func TestFeatures(t *testing.T) {
	gocuke.NewRunner(t, func(t gocuke.TestingT) gocuke.StepDefinitions {
		return &suite{TestingT: t}
	}).Run()
}

type suite struct {
	gocuke.TestingT
	table    ormtable.Table
	ctx      context.Context
	existing *testpb.ExampleTable
	err      error
}

func (s *suite) Before() {
	var err error
	s.table, err = ormtable.Build(ormtable.Options{
		MessageType: (&testpb.ExampleTable{}).ProtoReflect().Type(),
	})
	assert.NilError(s, err)
	s.ctx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
}

func (s *suite) AnExistingEntity() {
	s.existing = &testpb.ExampleTable{}
	assert.NilError(s, s.table.Insert(s.ctx, s.existing))
}

func (s *suite) IInsertAnEntityWithTheSamePrimaryKey() {
	s.err = s.table.Insert(s.ctx, s.existing)
}

func (s *suite) IInsertAnEntityWithTheSameUniqueKey() {
	panic("TODO")
}

func (s *suite) APrimaryKeyForANonexistingEntity() {
	panic("TODO")
}

func (s *suite) IGetTheEntityByPrimaryKey() {
	panic("TODO")
}

func (s *suite) AUniqueKeyForANonexistingEntity() {
	panic("TODO")
}

func (s *suite) IGetTheEntityByUniqueKey() {
	panic("TODO")
}

func (s *suite) IUpdateAnEntityThatDoesntExist() {
	panic("TODO")
}

func (s *suite) IUpdateAnotherEntityToHaveTheSameUniqueKey() {
	panic("TODO")
}

func (s *suite) ExpectAError(a string) {
	panic("TODO")
}

func (s *suite) ExpectGrpcErrorCode(a string) {
	assert.Equal(s, a, status.Code(s.err).String())
}
