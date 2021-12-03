package ormsql

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
)

func TestMessageCodec(t *testing.T) {
	b := newBuilder(protojson.MarshalOptions{})
	msgType := (&testpb.A{}).ProtoReflect().Type()
	tableDesc := proto.GetExtension(msgType.Descriptor().Options(), ormv1alpha1.E_Table).(*ormv1alpha1.TableDescriptor)
	msgCdc, err := b.makeMessageCodec(msgType, tableDesc)
	assert.NilError(t, err)
	t.Logf("%+v", msgCdc.structType)
	x := &testpb.A{U32: 7, I32: 4,
		Map:      map[string]uint32{"abc": 4},
		Msg:      &testpb.B{X: "foo"},
		Repeated: []uint32{1, 2, 4, 7, 9},
	}
	val := msgCdc.encode(x.ProtoReflect())
	t.Logf("%+v", val)

	//db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("file:test.sqlite"), &gorm.Config{})
	assert.NilError(t, err)
	assert.NilError(t, msgCdc.autoMigrate(db))
	msgCdc.save(db, x.ProtoReflect())
}
