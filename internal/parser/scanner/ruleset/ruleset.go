package ruleset

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

type Ruleset struct {
	WhitespaceDetector matcher.M
	Rules              []Rule
}

type Rule interface {
	Apply(RuneScanner) (token.Type, bool)
}

type FuncRule func(RuneScanner) (token.Type, bool)

func (r FuncRule) Apply(s RuneScanner) (token.Type, bool) { return r(s) }

type RuneScanner interface {
	Lookahead() (rune, bool)
	ConsumeRune()
}
