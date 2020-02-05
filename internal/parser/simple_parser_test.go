package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ scanner.Scanner = (*_testScanner)(nil) // ensure that testScanner implements scanner.Scanner

type _testScanner struct {
	pos    int
	tokens []token.Token
}

func scannerOf(tokens ...token.Token) *_testScanner {
	return &_testScanner{tokens: tokens}
}

func (s *_testScanner) HasNext() bool {
	return s.pos < len(s.tokens)
}

func (s *_testScanner) Next() token.Token {
	tk := s.Peek()
	s.pos++
	return tk
}

func (s *_testScanner) Peek() token.Token {
	return s.tokens[s.pos]
}

func Test_simpleParser_Next(t *testing.T) {
	tests := []struct {
		name   string
		tokens []token.Token
		stmt   *ast.SQLStmt
		errs   []error
		ok     bool
	}{
		{
			"no tokens",
			[]token.Token{},
			nil,
			[]error{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			p := &simpleParser{
				scanner: scannerOf(tt.tokens...),
			}
			stmt, errs, ok := p.Next()

			assert.Equal(stmt, tt.stmt)
			assert.Equal(errs, tt.errs)
			assert.Equal(ok, tt.ok)
		})
	}
}
