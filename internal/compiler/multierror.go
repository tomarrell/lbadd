package compiler

import "strings"

// MultiError is a wrapper for an error slice, which provides convenient
// wrapping of multiple errors into a single error. MultiError is not safe for
// concurrent use.
type MultiError struct {
	errs []error
}

// Append appends an error to the multi error. If the error is nil, it
// will still be added.
func (e *MultiError) Append(err error) {
	e.errs = append(e.errs, err)
}

func (e *MultiError) Error() string {
	if len(e.errs) == 0 {
		return ""
	} else if len(e.errs) == 1 {
		return e.errs[0].Error()
	}

	var buf strings.Builder
	buf.WriteString("multiple errors:")
	for _, err := range e.errs {
		buf.WriteString("\n\t" + err.Error())
	}
	return buf.String()
}
