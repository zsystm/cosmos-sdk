package orm

import "context"

type DBConnection interface {
	OpenRead(context.Context) (*ReadDB, error)
	Open(context.Context) (*DB, error)
}
