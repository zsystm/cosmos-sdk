package baseapp

// CheckState is an exported method to be able to access baseapp's
// checkState in tests.
//
// This method is only accessible in baseapp tests.
func (app *BaseApp) CheckState() *state {
	return app.checkState
}

// DeliverState is an exported method to be able to access baseapp's
// deliverState in tests.
//
// This method is only accessible in baseapp tests.
func (app *BaseApp) DeliverState() *state {
	return app.deliverState
}
