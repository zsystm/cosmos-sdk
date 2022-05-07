package event

import (
	"context"
)

type Service interface {
	GetManager(context.Context) Manager
	WithManager(ctx context.Context, manager Manager) context.Context
}
