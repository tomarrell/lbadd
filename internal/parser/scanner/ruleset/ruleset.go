package ruleset

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// DetectorFunc is a function that can detect if a given rune has a
// implementation specific attribute.
type DetectorFunc func(rune) bool

// Ruleset is a collection of a whitespace detector, a linefeed detector and a
// slice of rules. A rule based scanner can work with this to create tokens from
// the given rules.
type Ruleset struct {
	WhitespaceDetector DetectorFunc
	LinefeedDetector   DetectorFunc
	Rules              []Rule
}

// Rule describes a single scanner rule that can theoretically be applied. If
// the rule is applicable, is decided by the rule itself. If it returns
// applicable=false, the it is the rule based scanner's responsibility, to reset
// its position information to where it was, before the rule was applied.
//
// A rule takes a RuneScanner as parameter, which can be used to evaluate runes
// and advance the position pointer of the rule based scanner. After consuming n
// matching runes, return a token.Type. The rule based scanner will create a
// token from its position information, the consumed runes and the returned
// token.Type.
//
//  input := "hello"
//  ...
//  func (r myRule) Apply(s RuneScanner) (token.Type, bool) {
//      r.ConsumeRune()
//      r.ConsumeRune()
//      r.ConsumeRune()
//      r.ConsumeRune()
//      r.ConsumeRune()
//      return token.Literal, true
//  }
//
// The above example will cause the rule based scanner to emit a token with
// length=5, value=hello, offset=0.
//
// To use a func(RuneScanner) (token.Type, bool) as Rule, see ruleset.FuncRule.
type Rule interface {
	// Apply tries to apply this rule to the given RuneScanner. If the rule
	// determines itself as not applicable to the current RuneScanner, return
	// applicable=false, and the rule based scanner will handle everything else.
	Apply(RuneScanner) (token token.Type, applicable bool)
}

// FuncRule is an type alias that works as implementation for a ruleset.Rule.
//
//  func myRule(s RuneScanner) (token.Type, bool) { ... }
//  var _ Rule = FuncRule(myRule)
type FuncRule func(RuneScanner) (token.Type, bool)

// Apply implements (ruleset.Rule).Apply.
func (r FuncRule) Apply(s RuneScanner) (token.Type, bool) { return r(s) }

// RuneScanner is a capsule that limits the interaction of the scanner to a
// subset of two core functions. It is not guaranteed that the value in this
// interface will always be of the same type. This implies, that you shouldn't
// place type assertions on this interface.
type RuneScanner interface {
	Lookahead() (rune, bool)
	ConsumeRune()
}
