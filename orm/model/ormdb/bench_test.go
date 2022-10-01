package ormdb_test

import (
	"context"
	"fmt"
	"testing"

	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/internal/testkv"
	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/kv"
)

func BenchmarkORMMemory(b *testing.B) {
	bench(b, func(tb *testing.B) fixture {
		return newOrmFixture(tb, ormtest.NewMemoryBackend())
	})
}

func BenchmarkORMLevelDB(b *testing.B) {
	bench(b, func(tb *testing.B) fixture {
		return newOrmFixture(tb, testkv.NewGoLevelDBBackend(b))
	})
}

func BenchmarkManualMemory(b *testing.B) {
	bench(b, func(*testing.B) fixture {
		return &manualFixture{store: dbm.NewMemDB()}
	})
}

func BenchmarkManualLevelDB(b *testing.B) {
	bench(b, func(b *testing.B) fixture {
		db, err := dbm.NewGoLevelDB("test", b.TempDir())
		assert.NilError(b, err)
		return &manualFixture{store: db}
	})
}

func bench(b *testing.B, newFixture func(*testing.B) fixture) {
	f := initFixture(b, newFixture)
	testFixture(b, f)

	b.Run("insert", func(b *testing.B) {
		f := initFixture(b, newFixture)
		benchInsert(b, f)
	})
	b.Run("add update", func(b *testing.B) {
		f := initFixture(b, newFixture)
		benchAddUpdate(b, f)
	})
	b.Run("sub update", func(b *testing.B) {
		f := initFixture(b, newFixture)
		benchSubUpdate(b, f)
	})
	b.Run("delete", func(b *testing.B) {
		f := initFixture(b, newFixture)
		benchDelete(b, f)
	})
	b.Run("get", func(b *testing.B) {
		f := initFixture(b, newFixture)
		b.StopTimer()
		benchInsert(b, f)

		b.StartTimer()
		benchGet(b, f, 10)
	})
}

func initFixture(b *testing.B, newFixture func(*testing.B) fixture) fixture {
	b.StopTimer()
	store := newFixture(b)
	b.StartTimer()
	return store
}

func testFixture(t testing.TB, f fixture) {
	assertBalance := func(expected uint64) {
		bal, err := f.getBalance("acct1", "foo")
		assert.NilError(t, err)
		assert.Equal(t, expected, bal)
	}

	assertBalance(0)

	assert.NilError(t, f.addBalance("acct1", "foo", 10))
	assertBalance(10)

	assert.NilError(t, f.addBalance("acct1", "foo", 5))
	assertBalance(15)

	assert.NilError(t, f.safeSubBalance("acct1", "foo", 10))
	assertBalance(5)

	assert.NilError(t, f.safeSubBalance("acct1", "foo", 5))
	assertBalance(0)
}

func benchInsert(b *testing.B, f fixture) {
	for i := 0; i < b.N; i++ {
		assert.NilError(b, f.addBalance(
			fmt.Sprintf("acct%d", i),
			"foo",
			10,
		))
	}
}

func benchAddUpdate(b *testing.B, f fixture) {
	b.StopTimer()
	benchInsert(b, f)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		assert.NilError(b, f.addBalance(
			fmt.Sprintf("acct%d", i),
			"foo",
			5,
		))
	}
}

func benchSubUpdate(b *testing.B, f fixture) {
	b.StopTimer()
	benchInsert(b, f)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		assert.NilError(b, f.safeSubBalance(
			fmt.Sprintf("acct%d", i),
			"foo",
			5,
		))
	}
}

func benchDelete(b *testing.B, f fixture) {
	b.StopTimer()
	benchInsert(b, f)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		assert.NilError(b, f.safeSubBalance(
			fmt.Sprintf("acct%d", i),
			"foo",
			10,
		))
	}
}

func benchGet(b *testing.B, f fixture, expectedBalance uint64) {
	for i := 0; i < b.N; i++ {
		balance, err := f.getBalance(
			fmt.Sprintf("acct%d", i),
			"foo",
		)
		assert.NilError(b, err)
		assert.Equal(b, expectedBalance, balance)
	}
}

type fixture interface {
	addBalance(acct, denom string, amount uint64) error
	safeSubBalance(acct, denom string, amount uint64) error
	getBalance(acct, denom string) (uint64, error)
}

func newOrmFixture(t testing.TB, backend ormtable.Backend) fixture {
	db, err := ormdb.NewModuleDB(TestBankSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	store, err := testpb.NewBankStore(db)
	assert.NilError(t, err)
	k := &keeper{store}
	return &ormFixture{
		ctx: ormtable.WrapContextDefault(backend),
		k:   k,
	}
}

type ormFixture struct {
	ctx context.Context
	k   *keeper
}

func (o ormFixture) addBalance(acct, denom string, amount uint64) error {
	return o.k.addBalance(o.ctx, acct, denom, amount)
}

func (o ormFixture) safeSubBalance(acct, denom string, amount uint64) error {
	return o.k.safeSubBalance(o.ctx, acct, denom, amount)
}

func (o ormFixture) getBalance(acct, denom string) (uint64, error) {
	return o.k.Balance(o.ctx, acct, denom)
}

var _ fixture = &ormFixture{}

type manualFixture struct {
	store kv.Store
}

func (m manualFixture) addBalance(acct, denom string, amount uint64) error {
	balance, err := manualGetBalance(m.store, acct, denom)
	if err != nil {
		return err
	}

	if balance == nil {
		return manualInsertBalance(m.store, &testpb.Balance{
			Address: acct,
			Denom:   denom,
			Amount:  amount,
		})
	} else {
		balance.Amount += amount
		return manualUpdateBalance(m.store, balance)
	}
}

func (m manualFixture) safeSubBalance(acct, denom string, amount uint64) error {
	balance, err := manualGetBalance(m.store, acct, denom)
	if err != nil {
		return err
	}

	if balance == nil || amount > balance.Amount {
		return fmt.Errorf("insufficient funds")
	}

	balance.Amount -= amount

	if balance.Amount == 0 {
		return manualDeleteBalance(m.store, balance)
	} else {
		return manualUpdateBalance(m.store, balance)
	}
}

func (m manualFixture) getBalance(acct, denom string) (uint64, error) {
	bal, err := manualGetBalance(m.store, acct, denom)
	if err != nil {
		return 0, err
	}

	return bal.Amount, nil
}

var _ fixture = &manualFixture{}

const (
	addressDenomPrefix byte = iota
	denomAddressPrefix
)

func manualInsertBalance(store kv.Store, balance *testpb.Balance) error {
	denom := balance.Denom
	balance.Denom = ""
	addr := balance.Address
	balance.Address = ""

	addressDenomKey := []byte{addressDenomPrefix}
	addressDenomKey = append(addressDenomKey, []byte(addr)...)
	addressDenomKey = append(addressDenomKey, 0x0)
	addressDenomKey = append(addressDenomKey, []byte(denom)...)
	has, err := store.Has(addressDenomKey)
	if err != nil {
		return err
	}

	if has {
		return fmt.Errorf("already exists")
	}

	bz, err := proto.Marshal(balance)
	if err != nil {
		return err
	}
	balance.Denom = denom
	balance.Address = addr

	err = store.Set(addressDenomKey, bz)
	if err != nil {
		return err
	}

	// set denom address index
	denomAddressKey := []byte{denomAddressPrefix}
	denomAddressKey = append(denomAddressKey, []byte(balance.Denom)...)
	denomAddressKey = append(denomAddressKey, 0x0)
	denomAddressKey = append(denomAddressKey, []byte(balance.Address)...)
	err = store.Set(denomAddressKey, []byte{})
	if err != nil {
		return err
	}

	return nil
}

func manualUpdateBalance(store kv.Store, balance *testpb.Balance) error {
	denom := balance.Denom
	balance.Denom = ""
	addr := balance.Address
	balance.Address = ""
	bz, err := proto.Marshal(balance)
	if err != nil {
		return err
	}
	balance.Denom = denom
	balance.Address = addr

	addressDenomKey := []byte{addressDenomPrefix}
	addressDenomKey = append(addressDenomKey, []byte(addr)...)
	addressDenomKey = append(addressDenomKey, 0x0)
	addressDenomKey = append(addressDenomKey, []byte(denom)...)

	return store.Set(addressDenomKey, bz)
}

func manualDeleteBalance(store kv.Store, balance *testpb.Balance) error {
	denom := balance.Denom
	addr := balance.Address

	addressDenomKey := []byte{addressDenomPrefix}
	addressDenomKey = append(addressDenomKey, []byte(addr)...)
	addressDenomKey = append(addressDenomKey, 0x0)
	addressDenomKey = append(addressDenomKey, []byte(denom)...)
	err := store.Delete(addressDenomKey)
	if err != nil {
		return err
	}

	denomAddressKey := []byte{denomAddressPrefix}
	denomAddressKey = append(denomAddressKey, []byte(balance.Denom)...)
	denomAddressKey = append(denomAddressKey, 0x0)
	denomAddressKey = append(denomAddressKey, []byte(balance.Address)...)
	return store.Delete(denomAddressKey)
}

func manualGetBalance(store kv.Store, address, denom string) (*testpb.Balance, error) {
	addressDenomKey := []byte{addressDenomPrefix}
	addressDenomKey = append(addressDenomKey, []byte(address)...)
	addressDenomKey = append(addressDenomKey, 0x0)
	addressDenomKey = append(addressDenomKey, []byte(denom)...)

	bz, err := store.Get(addressDenomKey)
	if err != nil {
		return nil, err
	}

	if bz == nil {
		return &testpb.Balance{Address: address, Denom: denom, Amount: 0}, nil
	}

	balance := testpb.Balance{}
	err = proto.Unmarshal(bz, &balance)
	if err != nil {
		return nil, err
	}

	balance.Address = address
	balance.Denom = denom

	return &balance, nil
}
