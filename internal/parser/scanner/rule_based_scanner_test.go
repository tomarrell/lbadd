package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func TestRuleBasedScanner(t *testing.T) {
	inputs := []struct {
		name    string
		query   string
		ruleset ruleset.Ruleset
		want    []token.Token
	}{
		{
			"SELECT FROM WHERE",
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
			"SELECT FROM Literal",
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
			"unclosed literal",
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
			"many whitespaces with delimiters and literals",
			"SELECT      FROM || & +7 5 \"foobar\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 13, 12, 4, token.KeywordFrom, "FROM"),
				token.New(1, 18, 17, 2, token.BinaryOperator, "||"),
				token.New(1, 21, 20, 1, token.BinaryOperator, "&"),
				token.New(1, 23, 22, 1, token.UnaryOperator, "+"),
				token.New(1, 24, 23, 1, token.LiteralNumeric, "7"),
				token.New(1, 26, 25, 1, token.LiteralNumeric, "5"),
				token.New(1, 28, 27, 8, token.Literal, "\"foobar\""),
				token.New(1, 36, 35, 0, token.EOF, ""),
			},
		},
		{
			"fractional numeric literal",
			"SELECT      FROM || & +7 5.9 \"foobar\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 13, 12, 4, token.KeywordFrom, "FROM"),
				token.New(1, 18, 17, 2, token.BinaryOperator, "||"),
				token.New(1, 21, 20, 1, token.BinaryOperator, "&"),
				token.New(1, 23, 22, 1, token.UnaryOperator, "+"),
				token.New(1, 24, 23, 1, token.LiteralNumeric, "7"),
				token.New(1, 26, 25, 3, token.LiteralNumeric, "5.9"),
				token.New(1, 30, 29, 8, token.Literal, "\"foobar\""),
				token.New(1, 38, 37, 0, token.EOF, ""),
			},
		},
		{
			"single quote literal",
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
			"unclosed literal",
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
			"misplaced quote",
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
			"literal closing escape double quote",
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
			"literal closing escape single quote",
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
			"unary and binary operators",
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
			"numeric literals",
			"7 7.5 8.9.8 8.0 0.4 10 10000 18907.890 1890976.09.977",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 1, token.LiteralNumeric, "7"),
				token.New(1, 3, 2, 3, token.LiteralNumeric, "7.5"),
				token.New(1, 7, 6, 3, token.LiteralNumeric, "8.9"),
				token.New(1, 10, 9, 2, token.LiteralNumeric, ".8"),
				token.New(1, 13, 12, 3, token.LiteralNumeric, "8.0"),
				token.New(1, 17, 16, 3, token.LiteralNumeric, "0.4"),
				token.New(1, 21, 20, 2, token.LiteralNumeric, "10"),
				token.New(1, 24, 23, 5, token.LiteralNumeric, "10000"),
				token.New(1, 30, 29, 9, token.LiteralNumeric, "18907.890"),
				token.New(1, 40, 39, 10, token.LiteralNumeric, "1890976.09"),
				token.New(1, 50, 49, 4, token.LiteralNumeric, ".977"),
				token.New(1, 54, 53, 0, token.EOF, ""),
			},
		},
		{
			"numeric literals with exponents, no comma leading and or trailing digits, hex literals",
			"11.672E19 11.672E+19 11.657EE19 0xCAFEBABE 2.5E-1 1.2.3.4.5.6.7 5.hello something.4 ",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 9, token.LiteralNumeric, "11.672E19"),
				token.New(1, 11, 10, 10, token.LiteralNumeric, "11.672E+19"),
				token.New(1, 22, 21, 2, token.Literal, "11"),
				token.New(1, 24, 23, 1, token.Literal, "."),
				token.New(1, 25, 24, 7, token.Literal, "657EE19"),
				token.New(1, 33, 32, 10, token.LiteralNumeric, "0xCAFEBABE"),
				token.New(1, 44, 43, 6, token.LiteralNumeric, "2.5E-1"),
				token.New(1, 51, 50, 3, token.LiteralNumeric, "1.2"),
				token.New(1, 54, 53, 2, token.LiteralNumeric, ".3"),
				token.New(1, 56, 55, 2, token.LiteralNumeric, ".4"),
				token.New(1, 58, 57, 2, token.LiteralNumeric, ".5"),
				token.New(1, 60, 59, 2, token.LiteralNumeric, ".6"),
				token.New(1, 62, 61, 2, token.LiteralNumeric, ".7"),
				token.New(1, 65, 64, 2, token.LiteralNumeric, "5."),
				token.New(1, 67, 66, 5, token.Literal, "hello"),
				token.New(1, 73, 72, 9, token.Literal, "something"),
				token.New(1, 82, 81, 2, token.LiteralNumeric, ".4"),
				token.New(1, 85, 84, 0, token.EOF, ""),
			},
		},
		{
			"asc desc regression",
			"ASC DESC",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 3, token.KeywordAsc, "ASC"),
				token.New(1, 5, 4, 4, token.KeywordDesc, "DESC"),
				token.New(1, 9, 8, 0, token.EOF, ""),
			},
		},
		{
			"placeholder as literal",
			"SELECT * FROM users WHERE name = ?;",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
				token.New(1, 8, 7, 1, token.BinaryOperator, "*"),
				token.New(1, 10, 9, 4, token.KeywordFrom, "FROM"),
				token.New(1, 15, 14, 5, token.Literal, "users"),
				token.New(1, 21, 20, 5, token.KeywordWhere, "WHERE"),
				token.New(1, 27, 26, 4, token.Literal, "name"),
				token.New(1, 32, 31, 1, token.BinaryOperator, "="),
				token.New(1, 34, 33, 1, token.Literal, "?"),
				token.New(1, 35, 34, 1, token.StatementSeparator, ";"),
				token.New(1, 36, 35, 0, token.EOF, ""),
			},
		},
		{
			"placeholder within unquoted literals",
			"foobar?snafu",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 6, token.Literal, "foobar"),
				token.New(1, 7, 6, 1, token.Literal, "?"),
				token.New(1, 8, 7, 5, token.Literal, "snafu"),
				token.New(1, 13, 12, 0, token.EOF, ""),
			},
		},
		{
			"underscore in single unquoted token",
			"alpha_beta",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 10, token.Literal, "alpha_beta"),
				token.New(1, 11, 10, 0, token.EOF, ""),
			},
		},
		{
			"underscore in single quoted token",
			"\"alpha_beta\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 12, token.Literal, "\"alpha_beta\""),
				token.New(1, 13, 12, 0, token.EOF, ""),
			},
		},
		{
			"dash in single unquoted token",
			"alpha-beta",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 10, token.Literal, "alpha-beta"),
				token.New(1, 11, 10, 0, token.EOF, ""),
			},
		},
		{
			"dash in single quoted token",
			"\"alpha-beta\"",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 12, token.Literal, "\"alpha-beta\""),
				token.New(1, 13, 12, 0, token.EOF, ""),
			},
		},
		{
			"binary expression",
			"2+3",
			ruleset.Default,
			[]token.Token{
				token.New(1, 1, 0, 1, token.LiteralNumeric, "2"),
				token.New(1, 2, 1, 1, token.UnaryOperator, "+"),
				token.New(1, 3, 2, 1, token.LiteralNumeric, "3"),
				token.New(1, 4, 3, 0, token.EOF, ""),
			},
		},
	}
	for _, input := range inputs {
		t.Run("ruleset=default/"+input.name, _TestRuleBasedScannerWithRuleset(input.query, input.ruleset, input.want))
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
			assert.Equal(want[i].Type(), got[i].Type(), "Type doesn't match, want %s, but got %s", want[i].Type().String(), got[i].Type().String())
			assert.Equal(want[i].Value(), got[i].Value(), "Value doesn't match")
		}
	}
}

func TestRuleBasedSannerStatementEndingInWhitespace(t *testing.T) {
	assert := assert.New(t)

	stmt := "SELECT "
	sc := NewRuleBased([]rune(stmt), ruleset.Default)
	next := sc.Next()
	assert.Equal(token.KeywordSelect, next.Type())
	eof := sc.Next()
	assert.Equal(token.EOF, eof.Type())
}
