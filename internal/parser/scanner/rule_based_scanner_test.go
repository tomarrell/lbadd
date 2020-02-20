package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func TestRuleBasedScanner(t *testing.T) {
	inputs := []struct {
		query   string
		ruleset ruleset.Ruleset
		want    []token.Token
	}{
		{
			"SELECT FROM \"WHERE\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 7, token.Literal, "\"WHERE\""),
				token.New(1, 20, 19, 0, token.EOF, ""),
			},
		},
		{
			"SELECT FROM \"WHERE",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 6, token.Error, "unexpected token: '\"WHERE' at offset 12"),
				token.New(1, 19, 18, 0, token.EOF, ""),
			},
		},
		{
			"SELECT      FROM || & +7 59 \"foobar\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 13, 12, 4, token.KeywordFrom, "FROM"),
				token.New(1, 18, 17, 2, token.BinaryOperator, "||"),
				token.New(1, 21, 20, 1, token.BinaryOperator, "&"),
				token.New(1, 23, 22, 1, token.UnaryOperator, "+"),
				token.New(1, 24, 23, 1, token.Literal, "7"),
				token.New(1, 26, 25, 2, token.Literal, "59"),
				token.New(1, 29, 28, 8, token.Literal, "\"foobar\""),
				token.New(1, 37, 36, 0, token.EOF, ""),
			},
		},
		{
			"SELECT FROM 'WHERE'",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 7, token.Literal, "'WHERE'"),
				token.New(1, 20, 19, 0, token.EOF, ""),
			},
		},
		{
			"SELECT \"myCol FROM \"myTable\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 13, token.Literal, `"myCol FROM "`),
				token.New(1, 21, 20, 7, token.Literal, "myTable"),
				token.New(1, 28, 27, 1, token.Error, "unexpected token: '\"' at offset 27"),
				token.New(1, 29, 28, 0, token.EOF, ""),
			},
		},
		{
			"SELECT \" FROM",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 1, token.Error, `unexpected token: '"' at offset 7`),
				token.New(1, 10, 9, 4, token.KeywordFrom, "FROM"),
				token.New(1, 14, 13, 0, token.EOF, ""),
			},
		},
		{
			`SELECT FROM "this \" can be anything"`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 25, token.Literal, `"this \" can be anything"`),
				token.New(1, 38, 37, 0, token.EOF, ""),
			},
		},
		{
			`SELECT FROM 'this \" can be anything'`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 25, token.Literal, `'this \" can be anything'`),
				token.New(1, 38, 37, 0, token.EOF, ""),
			},
		},
		{
			`SELECT FROM 'this \' can be anything'`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 25, token.Literal, `'this \' can be anything'`),
				token.New(1, 38, 37, 0, token.EOF, ""),
			},
		},
		{
			`SELECT FROM "this \' can be anything"`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 25, token.Literal, `"this \' can be anything"`),
				token.New(1, 38, 37, 0, token.EOF, ""),
			},
		},
	}
	for _, input := range inputs {
		t.Run("ruleset=default/"+input.query, _TestRuleBasedScannerWithRuleset(input.query, input.ruleset, input.want))
	}
}

func _TestRuleBasedScannerWithRuleset(input string, ruleset ruleset.Ruleset, want []token.Token) func(*testing.T) {
	return func(t *testing.T) {
		assert := assert.New(t)

		var got []token.Token

		// create the scanner to be tested
		sc := NewRuleBased([]rune(input), ruleset)

		// collect all whitespaces
		for {
			next := sc.Next()
			got = append(got, next)
			if next.Type() == token.EOF {
				break
			}
		}
		assert.Equalf(len(want), len(got), "did not receive as much tokens as expected (expected %d, but got %d)", len(want), len(got))

		limit := len(want)
		if len(got) < limit {
			limit = len(got)
		}

		for i := 0; i < limit; i++ {
			assert.Equal(want[i].Line(), got[i].Line(), "Line doesn't match")
			assert.Equal(want[i].Col(), got[i].Col(), "Col doesn't match")
			assert.Equal(want[i].Offset(), got[i].Offset(), "Offset doesn't match")
			assert.Equal(want[i].Length(), got[i].Length(), "Length doesn't match")
			assert.Equal(want[i].Type(), got[i].Type(), "Type doesn't match")
			assert.Equal(want[i].Value(), got[i].Value(), "Value doesn't match")
		}
	}
}
