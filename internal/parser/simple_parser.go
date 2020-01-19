package parser

import (
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ Parser = (*simpleParser)(nil) // ensure that simpleParser implements Parser

type simpleParser struct {
	scanner scanner.Scanner
}

func New() Parser {
	return &simpleParser{}
}

func (p *simpleParser) HasNext() bool {
	panic("implement me")
}

func (p *simpleParser) Next() *ast.SqlStmt {
	if !p.HasNext() {
		panic("no more statements, check with HasNext() before calling Next()")
	}
	panic("implement me")
}

func (p *simpleParser) expect(t token.Type) (token.Token, bool) {
	if !p.scanner.HasNext() {
		return nil, false
	}

	// check if the next token's type matches
	tk := p.scanner.Peek()
	if tk.Type() == t {
		return p.scanner.Next(), true // if it matches, consume the token
	}
	return tk, false
}

func (p *simpleParser) parseSqlStatement() *ast.SqlStmt {
	_, _ = p.expect(token.KeywordExplain)
	// TODO(TimSatke): QUERY PLAN

	if !p.scanner.HasNext() {
		panic("implement error handling")
	}

	next := p.scanner.Next()
	switch next.Type() {
	default:
		panic("implement next rules")
	}

	panic("implement me")
}
