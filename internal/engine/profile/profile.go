package profile

// Profile is a collection of profiling events that were collected by a
// profiler.
type Profile struct {
	Events []Event
}
