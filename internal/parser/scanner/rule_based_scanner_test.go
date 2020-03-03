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
			"SELECT FROM WHERE",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
				token.New(1, 13, 12, 5, token.KeywordWhere, "WHERE"),
				token.New(1, 18, 17, 0, token.EOF, ""),
			},
		},
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
				token.New(1, 13, 12, 1, token.Error, `unexpected token: '"' at offset 12`),
				token.New(1, 14, 13, 5, token.KeywordWhere, "WHERE"),
				token.New(1, 19, 18, 0, token.EOF, ""),
			},
		},
		{
			"SELECT      FROM || & +7 5 \"foobar\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 13, 12, 4, token.KeywordFrom, "FROM"),
				token.New(1, 18, 17, 2, token.BinaryOperator, "||"),
				token.New(1, 21, 20, 1, token.BinaryOperator, "&"),
				token.New(1, 23, 22, 1, token.UnaryOperator, "+"),
				token.New(1, 24, 23, 1, token.Literal, "7"),
				token.New(1, 26, 25, 1, token.Literal, "5"),
				token.New(1, 28, 27, 8, token.Literal, "\"foobar\""),
				token.New(1, 36, 35, 0, token.EOF, ""),
			},
		},
		{
			"SELECT      FROM || & +7 5.9 \"foobar\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 13, 12, 4, token.KeywordFrom, "FROM"),
				token.New(1, 18, 17, 2, token.BinaryOperator, "||"),
				token.New(1, 21, 20, 1, token.BinaryOperator, "&"),
				token.New(1, 23, 22, 1, token.UnaryOperator, "+"),
				token.New(1, 24, 23, 1, token.Literal, "7"),
				token.New(1, 26, 25, 3, token.Literal, "5.9"),
				token.New(1, 30, 29, 8, token.Literal, "\"foobar\""),
				token.New(1, 38, 37, 0, token.EOF, ""),
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
			`|| * / % + - ~ << >> & | < <= > >= = == != <> !>> >>`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 2, token.BinaryOperator, "||"),
				token.New(1, 4, 3, 1, token.BinaryOperator, "*"),
				token.New(1, 6, 5, 1, token.BinaryOperator, "/"),
				token.New(1, 8, 7, 1, token.BinaryOperator, "%"),
				token.New(1, 10, 9, 1, token.UnaryOperator, "+"),
				token.New(1, 12, 11, 1, token.UnaryOperator, "-"),
				token.New(1, 14, 13, 1, token.UnaryOperator, "~"),
				token.New(1, 16, 15, 2, token.BinaryOperator, "<<"),
				token.New(1, 19, 18, 2, token.BinaryOperator, ">>"),
				token.New(1, 22, 21, 1, token.BinaryOperator, "&"),
				token.New(1, 24, 23, 1, token.BinaryOperator, "|"),
				token.New(1, 26, 25, 1, token.BinaryOperator, "<"),
				token.New(1, 28, 27, 2, token.BinaryOperator, "<="),
				token.New(1, 31, 30, 1, token.BinaryOperator, ">"),
				token.New(1, 33, 32, 2, token.BinaryOperator, ">="),
				token.New(1, 36, 35, 1, token.BinaryOperator, "="),
				token.New(1, 38, 37, 2, token.BinaryOperator, "=="),
				token.New(1, 41, 40, 2, token.BinaryOperator, "!="),
				token.New(1, 44, 43, 2, token.BinaryOperator, "<>"),
				token.New(1, 47, 46, 1, token.Error, "unexpected token: '!' at offset 46"),
				token.New(1, 48, 47, 2, token.BinaryOperator, ">>"),
				token.New(1, 51, 50, 2, token.BinaryOperator, ">>"),
				token.New(1, 53, 52, 0, token.EOF, ""),
			},
		},
		{
			`7 7.5 8.9.8 8.0 0.4 10 10000 18907.890 1890976.09.977`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 1, token.Literal, "7"),
				token.New(1, 3, 2, 3, token.Literal, "7.5"),
				token.New(1, 7, 6, 3, token.Literal, "8.9"),
				token.New(1, 10, 9, 2, token.Literal, ".8"),
				token.New(1, 13, 12, 3, token.Literal, "8.0"),
				token.New(1, 17, 16, 3, token.Literal, "0.4"),
				token.New(1, 21, 20, 2, token.Literal, "10"),
				token.New(1, 24, 23, 5, token.Literal, "10000"),
				token.New(1, 30, 29, 9, token.Literal, "18907.890"),
				token.New(1, 40, 39, 10, token.Literal, "1890976.09"),
				token.New(1, 50, 49, 4, token.Literal, ".977"),
				token.New(1, 54, 53, 0, token.EOF, ""),
			},
		},
		{
			`11.672E19 11.672E+19 11.657EE19 0xCAFEBABE 2.5E-1 1.2.3.4.5.6.7 5.hello something.4`,
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 9, token.Literal, "11.672E19"),
				token.New(1, 11, 10, 10, token.Literal, "11.672E+19"),
				token.New(1, 22, 21, 2, token.Literal, "11"),
				token.New(1, 24, 23, 1, token.Error, "unexpected token: '.' at offset 23"),
				token.New(1, 25, 24, 7, token.Literal, "657EE19"),
				token.New(1, 33, 32, 10, token.Literal, "0xCAFEBABE"),
				token.New(1, 44, 43, 6, token.Literal, "2.5E-1"),
				token.New(1, 51, 50, 3, token.Literal, "1.2"),
				token.New(1, 54, 53, 2, token.Literal, ".3"),
				token.New(1, 56, 55, 2, token.Literal, ".4"),
				token.New(1, 58, 57, 2, token.Literal, ".5"),
				token.New(1, 60, 59, 2, token.Literal, ".6"),
				token.New(1, 62, 61, 2, token.Literal, ".7"),
				token.New(1, 65, 64, 2, token.Literal, "5."),
				token.New(1, 67, 66, 5, token.Literal, "hello"),
				token.New(1, 73, 72, 9, token.Literal, "something"),
				token.New(1, 82, 81, 2, token.Literal, ".4"),
				token.New(1, 84, 83, 0, token.EOF, ""),
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
