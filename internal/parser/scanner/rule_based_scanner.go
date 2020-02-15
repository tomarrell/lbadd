package scanner

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ Scanner = (*ruleBasedScanner)(nil)

type ruleBasedScanner struct {
	input []rune

	cache token.Token

	whitespaceDetector matcher.M
	rules              []ruleset.Rule

	start int
	pos   int
}

func NewRuleBased(input []rune, ruleset ruleset.Ruleset) *ruleBasedScanner {
	return &ruleBasedScanner{
		input:              input,
		cache:              nil,
		whitespaceDetector: ruleset.WhitespaceDetector,
		rules:              ruleset.Rules,
		start:              0,
		pos:                0,
	}
}

func (s *ruleBasedScanner) Next() token.Token {
	tok := s.Peek()
	s.cache = nil
	return tok
}

func (s *ruleBasedScanner) Peek() token.Token {
	if s.cache == nil {
		s.cache = s.computeNext()
	}
	return s.cache
}

func (s *ruleBasedScanner) done() bool {
	return s.pos >= len(s.input)
}

func (s *ruleBasedScanner) computeNext() token.Token {
	if s.done() {
		return s.eof()
	}

	s.drainWhitespace()
	return s.applyRule()
}

func (s *ruleBasedScanner) applyRule() token.Token {
	// try to apply all rules in the given order
	for _, rule := range s.rules {
		typ, ok := rule.Apply(s)
		if ok {
			return s.createToken(typ)
		}
	}

	// no rules matched, create an error token
	s.seekNextWhitespace()
	return s.unexpectedToken()
}

func (s *ruleBasedScanner) seekNextWhitespace() {
	for {
		next, ok := s.Lookahead()
		if !ok || s.whitespaceDetector.Matches(next) {
			break
		}
		s.ConsumeRune()
	}
}

func (s *ruleBasedScanner) drainWhitespace() {
	for {
		next, ok := s.Lookahead()
		if !(ok && s.whitespaceDetector.Matches(next)) {
			break
		}
		s.ConsumeRune()
	}
	_ = s.createToken(token.Unknown) // discard consumed tokens
}

func (s *ruleBasedScanner) candidate() string {
	return string(s.input[s.start:s.pos])
}

func (s *ruleBasedScanner) eof() token.Token {
	return s.createTokenWithValue(token.EOF, "EOF")
}

func (s *ruleBasedScanner) unexpectedToken() token.Token {
	return s.createTokenWithValue(token.Error, fmt.Sprintf("unexpected token '%v' at offset %v", s.candidate(), s.start))
}

func (s *ruleBasedScanner) createToken(t token.Type) token.Token {
	return s.createTokenWithValue(t, s.candidate())
}

func (s *ruleBasedScanner) createTokenWithValue(t token.Type, val string) token.Token {
	tok := token.New(-1, -1, s.start, s.start-s.pos, t, val)
	s.start = s.pos
	return tok
}

// runeScanner

func (s *ruleBasedScanner) Lookahead() (rune, bool) {
	if !s.done() {
		return s.input[s.pos], true
	}
	return 0, false
}

func (s *ruleBasedScanner) ConsumeRune() {
	s.pos++
}
