package ormsql

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSqlite(t *testing.T) {
	_, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}
