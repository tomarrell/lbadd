package parser

import (
	"fmt"
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

type errorReporter struct {
	p      *simpleParser
	errs   []error
	sealed bool
}

func (r *errorReporter) incompleteStatement() {
	next, ok := r.p.lookahead()
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
	next, ok := r.p.lookahead()
	if !ok || next.Type() == token.EOF {
		// use this instead of r.prematureEOF() because we can add the
		// information about what tokens were expected
		r.errorf("%w: expected %s", ErrPrematureEOF, expected)
		r.sealed = true
		return
	}

	r.errorf("%w: got %s but expected one of %s at (%d:%d) offset %d length %d", ErrUnexpectedToken, next, expected, next.Line(), next.Col(), next.Offset(), next.Length())
}

func (r *errorReporter) unhandledToken(t token.Token) {
	r.errorf("%w: %s(%s) at (%d:%d) offset %d lenght %d", ErrUnknownToken, t.Type().String(), t.Value(), t.Line(), t.Col(), t.Offset(), t.Length())
}

func (r *errorReporter) unsupportedConstruct(t token.Token) {
	r.errorf("%w: %s(%s) at (%d:%d) offset %d lenght %d", ErrUnsupportedConstruct, t.Type().String(), t.Value(), t.Line(), t.Col(), t.Offset(), t.Length())
}

func (r *errorReporter) errorf(format string, args ...interface{}) {
	r.errs = append(r.errs, fmt.Errorf(format, args...))
}

type reporter interface {
	incompleteStatement()
	prematureEOF()
	unexpectedToken(expected ...token.Type)
	unhandledToken(t token.Token)
	unsupportedConstruct(t token.Token)
}

var _ Parser = (*simpleParser)(nil) // ensure that simpleParser implements Parser

type simpleParser struct {
	scanner scanner.Scanner
}

func New(input string) Parser {
	return &simpleParser{
		scanner: scanner.New([]rune(input)),
	}
}

func (p *simpleParser) Next() (*ast.SQLStmt, []error, bool) {
	if !p.scanner.HasNext() {
		return nil, []error{}, false
	}
	errs := &errorReporter{
		p:    p,
		errs: []error{},
	}
	stmt := p.parseSQLStatement(errs)
	return stmt, errs.errs, true
}

// skipUntil skips tokens until a token is of one of the given types. That token
// will not be consumed, every other token will be consumed and an unexpected
// token error will be reported.
func (p *simpleParser) skipUntil(r reporter, types ...token.Type) {
	for {
		next, ok := p.lookahead()
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		for _, typ := range types {
			if next.Type() == typ {
				return
			}
		}
		r.unexpectedToken(types...)
		p.consumeToken()
	}
}

func (p *simpleParser) skipUntilNoError(types ...token.Type) {
	for {
		next, ok := p.lookahead()
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		for _, typ := range types {
			if next.Type() == typ {
				return
			}
		}
		p.consumeToken()
	}
}

func (p *simpleParser) lookahead() (next token.Token, hasNext bool) {
	if !p.scanner.HasNext() {
		return nil, false
	}

	return p.scanner.Peek(), true
}

func (p *simpleParser) lookaheadWithType(typ token.Type) (token.Token, bool) {
	next, hasNext := p.lookahead()
	return next, hasNext && next.Type() == typ
}

func (p *simpleParser) consumeToken() {
	_ = p.scanner.Next()
}

func (p *simpleParser) parseSQLStatement(r reporter) (stmt *ast.SQLStmt) {
	stmt = &ast.SQLStmt{}

	if next, ok := p.lookaheadWithType(token.KeywordExplain); ok {
		stmt.Explain = next
		p.consumeToken()

		if next, ok := p.lookaheadWithType(token.KeywordQuery); ok {
			stmt.Query = next
			p.consumeToken()

			if next, ok := p.lookaheadWithType(token.KeywordPlan); ok {
				stmt.Plan = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordPlan)
				// At this point, just assume that 'QUERY' was a mistake. Don't
				// abort. It's very unlikely that 'PLAN' occurs somewhere, so
				// assume that the user meant to input 'EXPLAIN <statement>'
				// instead of 'EXPLAIN QUERY PLAN <statement>'.
			}
		}
	}

	// according to the grammar, these are the tokens that initiate a statement
	p.skipUntil(r, token.KeywordAlter, token.KeywordAnalyze, token.KeywordAttach, token.KeywordBegin, token.KeywordCommit, token.KeywordCreate, token.KeywordDelete, token.KeywordDetach, token.KeywordDrop, token.KeywordInsert, token.KeywordPragma, token.KeywordReindex, token.KeywordRelease, token.KeywordRollback, token.KeywordSavepoint, token.KeywordSelect, token.KeywordUpdate, token.KeywordVacuum)

	next, ok := p.lookahead()
	if !ok || next.Type() == token.EOF {
		r.prematureEOF()
		return stmt
	}

	switch next.Type() {
	case token.KeywordAlter:
		stmt.AlterTableStmt = p.parseAlterTableStmt(r)
	case token.StatementSeparator:
		r.incompleteStatement()
		p.consumeToken()
	default:
		r.unsupportedConstruct(next)
		p.skipUntilNoError(token.StatementSeparator, token.EOF)
	}

	next, ok = p.lookahead()
	if ok && (next.Type() != token.StatementSeparator || next.Type() != token.EOF) {
		r.unexpectedToken(token.StatementSeparator, token.EOF)
	}

	return stmt
}

func (p *simpleParser) parseAlterTableStmt(r reporter) (stmt *ast.AlterTableStmt) {
	panic("implement me")
}
