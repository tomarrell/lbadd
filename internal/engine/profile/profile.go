package profile

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Profile is a collection of profiling events that were collected by a
// profiler.
type Profile struct {
	Events []Event
}

func (p Profile) String() string {
	var buf bytes.Buffer

	evts := p.Events
	sort.Slice(evts, func(i, j int) bool { return strings.Compare(evts[i].Object.String(), evts[j].Object.String()) < 0 })

	firstEvt, lastEvt := evts[0], evts[0]
	for _, evt := range evts {
		if evt.Start.Before(firstEvt.Start) {
			firstEvt = evt
		}
		if evt.Start.After(lastEvt.Start) {
			lastEvt = evt
		}
	}

	startTime := firstEvt.Start
	endTime := lastEvt.Start.Add(lastEvt.Duration)

	_, _ = fmt.Fprintf(&buf, "Profile\n\tfrom %v\n\tto   %v\n\ttook %v\n", fmtTime(startTime), fmtTime(endTime), endTime.Sub(startTime))
	_, _ = fmt.Fprintf(&buf, "Events (%v):\n", len(evts))

	buckets := make(map[string][]Event)
	for _, evt := range evts {
		str := evt.Object.String()
		buckets[str] = append(buckets[str], evt)
	}

	for bucket, bucketEvts := range buckets {
		_, _ = fmt.Fprintf(&buf, "\t%v (%v events)\n", bucket, len(bucketEvts))
		totalDuration := 0 * time.Second
		for _, bucketEvt := range bucketEvts {
			totalDuration += bucketEvt.Duration
			_, _ = fmt.Fprintf(&buf, "\t\t- %v took %v\n", fmtTime(bucketEvt.Start), bucketEvt.Duration)
		}
		_, _ = fmt.Fprintf(&buf, "\t\taverage %v, total %v\n", totalDuration/time.Duration(len(bucketEvts)), totalDuration)
	}

	return buf.String()
}

func fmtTime(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}
