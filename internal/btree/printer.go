package btree

import "fmt"

type levels map[int][][]levelGroup

type levelGroup struct {
	parIdx  int
	entries []*Entry
}

func (l *levelGroup) String() string {
	if len(l.entries) == 0 {
		return "[]"
	}

	out := fmt.Sprintf("[%v", l.entries[0])

	for _, e := range l.entries[1:] {
		out += fmt.Sprintf(" %v", e)
	}

	out += "]"

	return out
}
