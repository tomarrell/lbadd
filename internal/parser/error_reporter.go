package parser

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

type errorReporter struct {
	p      *simpleParser
	errs   []error
	sealed bool
}

func (r *errorReporter) errorToken(t token.Token) {
	r.errorf("%w: %s", ErrScanner, t.Value())
}

func (r *errorReporter) incompleteStatement() {
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok {
		r.errorf("%w: EOF", ErrIncompleteStatement)
	} else {
		r.errorf("%w: %s at (%d:%d) offset %d length %d", ErrIncompleteStatement, next.Type().String(), next.Line(), next.Col(), next.Offset(), next.Length())
	}
}

func (r *errorReporter) prematureEOF() {
	r.errorf("%w", ErrPrematureEOF)
	r.sealed = true
}

func (r *errorReporter) unexpectedToken(expected ...token.Type) {
	if r.sealed {
		return
	}
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok || next.Type() == token.EOF {
		// use this instead of r.prematureEOF() because we can add the
		// information about what tokens were expected
		r.errorf("%w: expected %s", ErrPrematureEOF, expected)
		r.sealed = true
		return
	}

	r.errorf("%w: got %s but expected one of %s at (%d:%d) offset %d length %d", ErrUnexpectedToken, next, expected, next.Line(), next.Col(), next.Offset(), next.Length())
}

func (r *errorReporter) unexpectedSingleRuneToken(typ token.Type, expected ...rune) {
	if r.sealed {
		return
	}
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok || next.Type() == token.EOF {
		// use this instead of r.prematureEOF() because we can add the
		// information about what tokens were expected
		r.errorf("%w: expected %s (more precisely one of %v)", ErrPrematureEOF, typ, expected)
		r.sealed = true
		return
	}

	r.errorf("%w: got %s but expected one of %s at (%d:%d) offset %d length %d", ErrUnexpectedToken, next, typ, next.Line(), next.Col(), next.Offset(), next.Length())
}

func (r *errorReporter) unhandledToken(t token.Token) {
	r.errorf("%w: %s(%s) at (%d:%d) offset %d length %d", ErrUnknownToken, t.Type().String(), t.Value(), t.Line(), t.Col(), t.Offset(), t.Length())
}

func (r *errorReporter) unsupportedConstruct(t token.Token) {
	r.errorf("%w: %s(%s) at (%d:%d) offset %d length %d", ErrUnsupportedConstruct, t.Type().String(), t.Value(), t.Line(), t.Col(), t.Offset(), t.Length())
}

func (r *errorReporter) errorf(format string, args ...interface{}) {
	r.errs = append(r.errs, fmt.Errorf(format, args...))
}

type reporter interface {
	errorToken(t token.Token)
	incompleteStatement()
	prematureEOF()
	unexpectedToken(expected ...token.Type)
	unexpectedSingleRuneToken(typ token.Type, expected ...rune)
	unhandledToken(t token.Token)
	unsupportedConstruct(t token.Token)
}
