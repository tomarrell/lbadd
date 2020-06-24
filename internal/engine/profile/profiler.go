package profile

import (
	"fmt"
	"time"
)

type Clearer interface {
	Clear()
}

func NewProfiler() *Profiler {
	return &Profiler{}
}

type Profiler struct {
	events []Event
}

type Event struct {
	origin   *Profiler
	Object   fmt.Stringer
	Start    time.Time
	Duration time.Duration
}

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

func (p *Profiler) Exit(evt Event) {
	if p == nil {
		return
	}

	p.events = append(p.events, evt)
}

func (p *Profiler) Clear() {
	if p == nil {
		return
	}

	p.events = nil
}

func (p *Profiler) Profile() Profile {
	if p == nil {
		return Profile{}
	}

	return Profile{
		Events: p.events,
	}
}

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
