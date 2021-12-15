package orm

import "context"

type DB interface {
	OpenRead(context.Context) (*ReadClient, error)
	Open(context.Context) (*Client, error)
}
