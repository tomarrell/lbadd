package ruleset

import (
	"testing"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var (
	keywords = []string{"ATTACH", "DROP", "SELECT", "VACUUM"}
	Result   interface{}
)

func BenchmarkKeywordScan(b *testing.B) {
	_benchKeywords(b, defaultKeywordsRule)
}

func _benchKeywords(b *testing.B, fn func(s RuneScanner) (token.Type, bool)) {
	for _, keyword := range keywords {
		b.Run(keyword, _bench(b, createRuneScanner(keyword), fn))
	}
}

func createRuneScanner(input string) RuneScanner {
	return &runeScanner{
		input: []rune(input),
	}
}

type runeScanner struct {
	input []rune
	idx   int
}

func (s *runeScanner) ConsumeRune() {
	s.idx++
}

func (s *runeScanner) Lookahead() (rune, bool) {
	if s.idx >= len(s.input) {
		return 0, false
	}
	return s.input[s.idx], true
}

func _bench(b *testing.B, rs RuneScanner, fn func(s RuneScanner) (token.Type, bool)) func(*testing.B) {
	return func(b *testing.B) {
		var ok bool
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, ok = fn(rs)
		}
		Result = ok
	}
}
