package btree

import "fmt"

type levels map[int]*level

type level struct {
	entries [][]*Entry
}

func (l *level) String() string {
	out := "["

	for _, e := range l.entries {
		out += fmt.Sprintf("%v", e)
	}

	out += "]"

	return out
}
