package engine

// ExecutionContext is a context that is passed down throughout a complete
// evaluation. It may be populated further.
type ExecutionContext struct {
	*executionContext
}

type executionContext struct {
}
