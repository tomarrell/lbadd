package scanner

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// Scanner is the interface that describes a scanner.
type Scanner interface {
	HasNext() bool
	Next() token.Token
	Peek() token.Token
}

type scanner struct {
	input          []rune
	start          int
	pos, line, col int
	closed         bool
}

// checkpoint represents the state of a scanner at any given point in time. A
// scanner can be restored to a checkpoint.
//
//	var s *Scanner
//	...
//	chck := s.checkpoint() // create a checkpoint
//	...                     // accept(), next(), goback(), whatever
//	s.restore(chck)        // scanner is in the same state as when the checkpoint was created
//
// This is useful when a state should not return an error if something
// unexpected was read, but for example another grammar production should be
// tried. To guarantee that both tries start with the same scanner state, a
// checkpoint can be used.
type checkpoint struct {
	start int
	pos   int

	startLine, startCol int
	line, lastCol, col  int
}

// New returns a new scanner
func New(input []rune) Scanner {
	// LoadTrie()
	return &scanner{
		input: input,
		start: 0,
		pos:   0,

		closed: false,
	}
}

// HasNext checks for existance of next token, returns true if exists, false otherwise.
func (s *scanner) HasNext() bool {
	return !s.done()
}

// Next reads the next token. This is basically starting from the initial state until a
// token gets emitted. If an error occurs, simply return an error token.
func (s *scanner) Next() token.Token {
	next := s.peekRune()
	switch next {
	case ';':
		s.consumeRune()
		return s.createToken(token.StatementSeparator)
	case '|', '*', '/', '%', '<', '>', '&', '=', '!', '~':
		return s.scanOperator()
	default:
		if ('a' <= next && next <= 'w') ||
			('A' <= next && next <= 'W') {
			return s.scanKeyword()
		}
		if whiteSpace.Matches(next) {
			return s.scanSpace()
		}
		return s.scanLiteral()
		// case '"':
		// 	return scanDoubleQuote(s)
		// case '\'':
		// 	return scanQuote(s)
		// case '(':
		// 	return scanLeftParanthesis(s)
		// case ')':
		// 	return scanRightParanthesis(s)
		// case ',':
		// 	return scanComma(s)
		// case '.':
		// 	return scanPeriod(s)
		// case '/':
		// 	return scanSolidus(s)
		// case '\\':
		// 	return scanReverseSolidus(s)
		// case ':':
		// 	return scanColon(s)
		// case ';':
		// 	return scanSemiColon(s)
		// case '?':
		// 	return scanQuestioMarkOrTrigraphs(s)
		// case '[':
		// 	return scanLeftBracket(s)
		// case ']':
		// 	return scanRightBracket(s)
		// case '^':
		// 	return scanCircumflex(s)
		// case '_':
		// 	return scanUnderscore(s)
		// case '|':
		// 	return scanVerticalBar(s)
		// case '{':
		// 	return scanLeftBrace(s)
		// case '}':
		// 	return scanRightBrace(s)
		// case '$':
		// 	return scanDollarSign(s)

	}
	return nil
}

func (s *scanner) Peek() token.Token {
	panic("implement me")
}

// Close will cause this scanner to not execute any more states. The execution
// of the current state cannot be aborted, but the scanner will stop executing
// states after the current state has finished.
func (s *scanner) Close() error {
	s.closed = true
	return nil
}

// done determines whether the scanner is done with its work. This is the case,
// if either the scanner was closed, or the scanner has reached the end of its
// input.
func (s *scanner) done() bool {
	return s.closed ||
		s.pos >= len(s.input)
}

// next returns the next rune of the scanners input and advances its pointer by
// one position. This method also updates the line and col information of the
// scanner. If the scanner.done()==true and this method is called, it will panic
// with a syntax error.
//
// The process of advancing the pointer and returning the read rune is called
// "consuming a rune" or "accepting a rune".
func (s *scanner) next() rune {
	// get the actual next rune
	next := s.input[s.pos]
	if next == '\n' {
		s.line++
		s.col = 1
	}
	// update current scanner position
	s.pos++

	return next
}

// peek returns the next rune of the scanners input without consuming it.
func (s *scanner) peekRune() rune {
	return s.input[s.pos]
}

// goback decrements the scanner's position by one and updates its line and col
// information.
func (s *scanner) goback() {
	s.pos--
}

// ignore discards all accepted runes. This is done by simply setting the start
// position of the last read token to the current scanner position.
func (s *scanner) ignore() {
	s.start = s.pos
}

// accept accepts exactly one rune matched by the given matcher. This means,
// that: If the next rune is matched by the scanner, it is consumed and ok=true
// is returned. If the next rune is NOT matched, it is unread and ok=false is
// returned. This implies, that accept(Alphanumeric) will actually do nothing if
// the next rune is not Alphanumeric. However, if the next rune is Alphanumeric,
// it will be accepted.
func (s *scanner) accept(m matcher.M) bool {
	if m.Matches(s.next()) {
		return true
	}
	s.goback()
	return false
}

// acceptMultiple accepts multiple runes that are matched by the given matcher.
// See the godoc on (*scanner).accept for more information. The amount of
// matched runes is returned.
func (s *scanner) acceptMultiple(m matcher.M) (matched uint) {
	for s.accept(m) {
		matched++
	}
	return
}

// acceptString accepts the exact sequence of runes that the given string
// represents, or does nothing, if the string is not matched.
//
//	input := []rune(".hello")
//	...
//	s.acceptString("hello") // will do nothing, as the next rune is '.'
//	s.next()                // advance the position by one (next rune is now 'h')
//	s.acceptString("hello") // will accept 5 runes, the scanner has reached its EOF now
func (s *scanner) acceptString(str string) bool {
	if s.peekString(str) {
		s.pos += len(str)
		return true
	}
	return false
}

// peekString works like (*scanner).acceptString, except it doesn't consume any
// runes. It just peeks, if the given string lays ahead.
func (s *scanner) peekString(str string) bool {
	for i, r := range str {
		if r != s.input[s.pos+i] {
			return false
		}
	}
	return true
}

// createToken creates a token with given parameters
func (s *scanner) createToken(t token.Type) token.Token {
	tk := token.New(s.line, s.col, s.start, s.pos-s.start, t, string(s.input[s.start:s.pos]))
	s.start = s.pos
	return tk
}

// seekNext returns the position of the end of a keyword.
// It takes the start position of the keyword.
func (s *scanner) seekNext(start int) int {
	for start < len(s.input) && s.input[start] != ' ' { //} Whitespace.Matches(s.input[start]) {
		start++
	}
	return start
}

func (s *scanner) unexpectedRune(value string) token.Token {
	tk := token.New(s.line, s.col, s.start, s.pos, token.Error, value)
	s.start = s.pos
	return tk
}
