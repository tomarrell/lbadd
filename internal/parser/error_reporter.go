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

func (r *errorReporter) expectedExpression() {
	if r.sealed {
		return
	}
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok || next.Type() == token.EOF {
		// use this instead of r.prematureEOF() because we can add the
		// information about what tokens were expected
		r.errorf("%w: expected expression", ErrPrematureEOF)
		r.sealed = true
		return
	}
	r.errorf("%w: got %s but expected expression at (%d:%d) offset %d length %d", ErrUnexpectedToken, next, next.Line(), next.Col(), next.Offset(), next.Length())
}

type reporter interface {
	// errorToken reports the given token as error. This makes sense, if the
	// scanner emitted an error token. If so, use this method to report it as
	// error.
	errorToken(t token.Token)
	// incompleteStatement indicates that the statement was incomplete, i. e.
	// that required parts are missing.
	incompleteStatement()
	// prematureEOF is a more specific version of an incompleteStatement, but
	// indicating that EOF was encountered instead of the expected token.
	prematureEOF()
	// unexpectedToken reports that the current token is unexpected. The given
	// token types are the token types that were expected at this point, and
	// will also be reported in order to make the error message more helpful.
	unexpectedToken(expected ...token.Type)
	// unexpectedSingleRuneToken is similar to unexpectedToken, but instead of a
	// token type, it prints the expected runes.
	unexpectedSingleRuneToken(typ token.Type, expected ...rune)
	// unhandledToken indicates that the handle for a specific token at this
	// point was not implemented yet.
	unhandledToken(t token.Token)
	// unsupportedConstruct indicates, that the found token can be recognized,
	// but is not supported for some reason (not implemented, no database
	// support, etc.).
	unsupportedConstruct(t token.Token)
	// expectedExpression is similar to unexpectedToken() but instead of
	// expecting tokens, it expects an expression.
	expectedExpression()
}
