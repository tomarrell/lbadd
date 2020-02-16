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
				token.New(1, 18, 17, 0, token.EOF, "EOF"),
			},
		},
	}
	for _, input := range inputs {
		t.Run("ruleset=default", _TestRuleBasedScannerWithRuleset(input.query, input.ruleset, input.want))
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
