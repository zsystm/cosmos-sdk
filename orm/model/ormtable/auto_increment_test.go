package ormtable_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/orm/internal/testkv"
	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestAutoIncrementScenario(t *testing.T) {
	table, err := ormtable.Build(ormtable.Options{
		MessageType: (&testpb.ExampleAutoIncrementTable{}).ProtoReflect().Type(),
	})
	assert.NilError(t, err)

	// first run tests with a split index-commitment store
	runAutoIncrementScenario(t, table, testkv.NewSplitMemBackend())

	// now run with shared store and debugging
	debugBuf := &strings.Builder{}
	store := testkv.NewDebugBackend(
		ormtable.ContextOptions{CommitmentStore: dbm.NewMemDB()},
		&testkv.EntryCodecDebugger{
			EntryCodec: table,
			Print:      func(s string) { debugBuf.WriteString(s + "\n") },
		},
	)
	runAutoIncrementScenario(t, table, store)

	golden.Assert(t, debugBuf.String(), "test_auto_inc.golden")
	checkEncodeDecodeEntries(t, table, store.IndexStoreReader())
}

func runAutoIncrementScenario(t *testing.T, table ormtable.Table, context ormtable.Context) {
	err := table.Save(context, &testpb.ExampleAutoIncrementTable{Id: 5}, ormtable.SAVE_MODE_DEFAULT)
	assert.ErrorContains(t, err, "update")

	ex1 := &testpb.ExampleAutoIncrementTable{X: "foo", Y: 5}
	assert.NilError(t, table.Save(context, ex1, ormtable.SAVE_MODE_DEFAULT))
	assert.Equal(t, uint64(1), ex1.Id)

	buf := &bytes.Buffer{}
	assert.NilError(t, table.ExportJSON(context, buf))
	golden.Assert(t, string(buf.Bytes()), "auto_inc_json.golden")

	assert.NilError(t, table.ValidateJSON(bytes.NewReader(buf.Bytes())))
	store2 := testkv.NewSplitMemBackend()
	assert.NilError(t, table.ImportJSON(store2, bytes.NewReader(buf.Bytes())))
	assertTablesEqual(t, table, context, store2)
}

func TestBadJSON(t *testing.T) {
	table, err := ormtable.Build(ormtable.Options{
		MessageType: (&testpb.ExampleAutoIncrementTable{}).ProtoReflect().Type(),
	})
	assert.NilError(t, err)

	store := testkv.NewSplitMemBackend()
	f, err := os.Open("testdata/bad_auto_inc.json")
	assert.NilError(t, err)
	assert.ErrorContains(t, table.ImportJSON(store, f), "invalid ID")

	f, err = os.Open("testdata/bad_auto_inc2.json")
	assert.NilError(t, err)
	assert.ErrorContains(t, table.ImportJSON(store, f), "invalid ID")
}
