package gas

import "context"

type Service interface {

	// GetMeter returns the current transaction-level gas meter. A non-nil meter
	// is always returned. When one is unavailable in the context a dummy instance
	// will be returned.
	GetMeter(ctx context.Context)

	// GetBlockMeter returns the current block-level gas meter. A non-nil meter
	// is always returned. When one is unavailable in the context a dummy instance
	// will be returned.
	GetBlockMeter(ctx context.Context)

	// WithMeter returns a new context with the provided transaction-level gas meter.
	WithMeter(ctx context.Context, meter Meter) context.Context

	// WithMeter returns a new context with the provided block-level gas meter.
	WithBlockMeter(ctx context.Context, meter Meter) context.Context
}
