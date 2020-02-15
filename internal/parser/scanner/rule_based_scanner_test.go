package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func TestRuleBasedScanner(t *testing.T) {
	t.Run("ruleset=default", _TestRuleBasedScannerWithRuleset("SELECT FROM WHERE", ruleset.Default, []token.Token{}))
}

func _TestRuleBasedScannerWithRuleset(input string, ruleset ruleset.Ruleset, tokens []token.Token) func(*testing.T) {
	return func(t *testing.T) {
		require := require.New(t)

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

		require.Equal(len(tokens), len(got), "did not receive as much tokens as expected")

		for i := range got {
			// TODO: extract ast/tool/cmp to project level and open up so that AST and Tokens can be
			// compared with it
			require.True(token.Equal(got[i], tokens[i]))
		}
	}
}
