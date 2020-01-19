package matcher

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

// M is the definition of a character class, which can tell whether a rune is
// part of that character definition or not.
type M struct {
	rt   *unicode.RangeTable
	desc string
}

// Matches describes, whether the given rune is contained in this matcher's
// range table.
func (m M) Matches(r rune) bool { return unicode.Is(m.rt, r) }

// String returns a human readable string description of the runes that this
// matcher matches.
func (m M) String() string { return m.desc }

// Matcher constructors

// New creates a new matcher from a given match function.
func New(desc string, rt *unicode.RangeTable) M {
	return M{rt, desc}
}

// String creates a matcher that checks whether a rune is part of the given
// string.
func String(s string) M {
	return M{
		rt:   rangetable.New([]rune(s)...),
		desc: "one of '" + s + "'",
	}
}

// RuneWithDesc creates a matcher that matches only the given rune. The
// description is the string representation of this matcher. This is useful when
// dealing with whitespace characters.
func RuneWithDesc(desc string, exp rune) M {
	return M{
		rt:   rangetable.New(exp),
		desc: desc,
	}
}

// Merge creates a new matcher, that accepts runes that are matched by one or
// more of the given matchers.
func Merge(ms ...M) M {
	var rtms []*unicode.RangeTable
	descs := make([]string, len(ms))

	for i, m := range ms {
		descs[i] = m.String()
		rtms = append(rtms, m.rt)
	}

	return M{
		rt:   rangetable.Merge(rtms...),
		desc: strings.Join(descs, " or "),
	}
}
