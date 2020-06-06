package parser

import (
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ Parser = (*simpleParser)(nil) // ensure that simpleParser implements Parser

type simpleParser struct {
	scanner scanner.Scanner
}

// NewSimpleParser creates new ready to use parser.
func NewSimpleParser(input string) Parser {
	return &simpleParser{
		scanner: scanner.NewRuleBased([]rune(input), ruleset.Default),
	}
}

func (p *simpleParser) Next() (*ast.SQLStmt, []error, bool) {
	if p.scanner.Peek().Type() == token.EOF {
		return nil, nil, false
	}
	errs := &errorReporter{
		p: p,
	}
	stmt := p.parseSQLStatement(errs)
	return stmt, errs.errs, true
}

// searchNext skips tokens until a token is of one of the given types. That
// token will not be consumed, every other token will be consumed and an
// unexpected token error will be reported.
func (p *simpleParser) searchNext(r reporter, types ...token.Type) {
	for {
		next, ok := p.unsafeLowLevelLookahead()
		if !ok {
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

func (p *simpleParser) skipUntil(types ...token.Type) {
	for {
		next, ok := p.unsafeLowLevelLookahead()
		if !ok {
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

// unsafeLowLevelLookahead is a low level lookahead, only use if needed.
// Remember to check for token.Error, token.EOF and token.StatementSeparator, as
// this will only return hasNext=false if there are no more tokens (which only
// should occur after an EOF token). Any other token will be returned with
// next=<token>,hasNext=true.
func (p *simpleParser) unsafeLowLevelLookahead() (next token.Token, hasNext bool) {
	return p.scanner.Peek(), true
}

// lookahead performs a lookahead while consuming any error or statement
// separator token, and reports an EOF, Error or IncompleteStatement if
// appropriate. If this returns ok=false, return from your parse function
// without reporting any more errors. If ok=false, this means that the next
// token was either a StatementSeparator or EOF, and an error has been reported.
//
// To get any token, even EOF or the a StatementSeparator, use
// (*simpleParser).optionalLookahead.
func (p *simpleParser) lookahead(r reporter) (next token.Token, ok bool) {
	next, ok = p.optionalLookahead(r)

	if !ok || next.Type() == token.EOF {
		r.prematureEOF()
		ok = false
	} else if next.Type() == token.StatementSeparator {
		r.incompleteStatement()
		ok = false
	}
	return
}

// optionalLookahead performs a lookahead while consuming any error token. If
// this returns ok=false, no more tokens are available.
func (p *simpleParser) optionalLookahead(r reporter) (next token.Token, ok bool) {
	next, ok = p.unsafeLowLevelLookahead()

	// drain all error tokens
	for ok && next.Type() == token.Error {
		r.errorToken(next)
		p.consumeToken()
		next, ok = p.unsafeLowLevelLookahead()
	}

	return next, ok
}

func (p *simpleParser) consumeToken() {
	_ = p.scanner.Next()
}
