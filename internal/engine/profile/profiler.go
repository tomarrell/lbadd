package profile

import (
	"fmt"
	"time"
)

// Clearer wraps a basic clear method, which will clear the components contents.
// What exactly is cleared, must be documented by the component.
type Clearer interface {
	Clear()
}

// NewProfiler returns a new, ready to use profiler.
func NewProfiler() *Profiler {
	return &Profiler{}
}

// Profiler is a profiler that can collect events.
type Profiler struct {
	events []Event
}

// Event is a simple profiling event. It keeps a back reference to the profiler
// it originated from.
type Event struct {
	origin   *Profiler
	Object   fmt.Stringer
	Start    time.Time
	Duration time.Duration
}

// Enter creates a profiling event. Use like this:
//
//	defer profiler.Enter(MyEvent).Exit()
func (p *Profiler) Enter(object fmt.Stringer) Event {
	if p == nil {
		return Event{}
	}

	return Event{
		origin: p,
		Object: object,
		Start:  time.Now(),
	}
}

// Exit collects the given event. You can call Exit multiple times with the same
// event, it will them appear multiple times in the profiler's profile.
func (p *Profiler) Exit(evt Event) {
	if p == nil {
		return
	}

	p.events = append(p.events, evt)
}

// Clear removes all collected events.
func (p *Profiler) Clear() {
	if p == nil {
		return
	}

	p.events = nil
}

// Profile returns a profile with all collected events from the profiler. The
// collected events are NOT cleared after this. To clear all events, use
// (*Profiler).Clear().
func (p *Profiler) Profile() Profile {
	if p == nil {
		return Profile{}
	}

	return Profile{
		Events: p.events,
	}
}

// Exit passes the event back to the origin profiler. When using this, unlike
// using (*Profiler).Exit(Event), an event duration will be set.
func (e Event) Exit() {
	if e.origin == nil {
		return
	}

	e.origin.Exit(Event{
		origin:   e.origin,
		Object:   e.Object,
		Start:    e.Start,
		Duration: time.Since(e.Start),
	})
}
