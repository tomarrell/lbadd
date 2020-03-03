package scanner

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ Scanner = (*ruleBasedScanner)(nil)

type ruleBasedScanner struct {
	input []rune

	cache token.Token

	whitespaceDetector ruleset.DetectorFunc
	linefeedDetector   ruleset.DetectorFunc
	rules              []ruleset.Rule

	state
}

type state struct {
	start     int
	startLine int
	startCol  int
	pos       int
	line      int
	col       int
}

// NewRuleBased creates a new, ready to use rule based scanner with the given
// ruleset, that will process the given input rune slice.
func NewRuleBased(input []rune, ruleset ruleset.Ruleset) Scanner {
	return &ruleBasedScanner{
		input:              input,
		cache:              nil,
		whitespaceDetector: ruleset.WhitespaceDetector,
		linefeedDetector:   ruleset.LinefeedDetector,
		rules:              ruleset.Rules,
		state: state{
			start:     0,
			startLine: 1,
			startCol:  1,
			pos:       0,
			line:      1,
			col:       1,
		},
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

func (s *ruleBasedScanner) checkpoint() state {
	return s.state
}

func (s *ruleBasedScanner) restore(chck state) {
	s.state = chck
}

func (s *ruleBasedScanner) done() bool {
	return s.pos >= len(s.input)
}

func (s *ruleBasedScanner) computeNext() token.Token {
	s.drainWhitespace()

	if s.done() {
		return s.eof()
	}
	return s.applyRule()
}

func (s *ruleBasedScanner) applyRule() token.Token {
	// try to apply all rules in the given order
	for _, rule := range s.rules {
		chck := s.checkpoint()
		typ, ok := rule.Apply(s)
		if ok {
			return s.token(typ)
		}
		s.restore(chck)
	}

	// no rules matched, create an error token
	s.ConsumeRune() // skip the one offending rune
	return s.unexpectedToken()
}

func (s *ruleBasedScanner) drainWhitespace() {
	for {
		next, ok := s.Lookahead()
		if !(ok && s.whitespaceDetector(next)) {
			break
		}
		s.ConsumeRune()
	}
	_ = s.token(token.Unknown) // discard consumed tokens
}

func (s *ruleBasedScanner) candidate() string {
	return string(s.input[s.start:s.pos])
}

func (s *ruleBasedScanner) eof() token.Token {
	return s.token(token.EOF)
}

func (s *ruleBasedScanner) unexpectedToken() token.Token {
	return s.errorToken(fmt.Errorf("%w: '%v' at offset %v", ErrUnexpectedToken, s.candidate(), s.start))
}

func (s *ruleBasedScanner) token(t token.Type) token.Token {
	tok := token.New(s.startLine, s.startCol, s.start, s.pos-s.start, t, s.candidate())
	s.updateStartPositions()
	return tok
}

func (s *ruleBasedScanner) errorToken(err error) token.Token {
	tok := token.NewErrorToken(s.startLine, s.startCol, s.start, s.pos-s.start, token.Error, err)
	s.updateStartPositions()
	return tok
}

func (s *ruleBasedScanner) updateStartPositions() {
	s.start = s.pos
	s.startLine = s.line
	s.startCol = s.col
}

// runeScanner

func (s *ruleBasedScanner) Lookahead() (rune, bool) {
	if !s.done() {
		return s.input[s.pos], true
	}
	return 0, false
}

func (s *ruleBasedScanner) ConsumeRune() {
	if s.linefeedDetector(s.input[s.pos]) {
		s.line++
		s.col = 1
	} else {
		s.col++
	}
	s.pos++
}
