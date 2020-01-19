package scanner

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func scanSpace(s *scanner) token.Token {
	s.accept(matcher.String(" "))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanDoubleQuote(s *scanner) token.Token {
	s.accept(matcher.String("\""))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanPercent(s *scanner) token.Token {
	s.accept(matcher.String("%"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanAmpersand(s *scanner) token.Token {
	s.accept(matcher.String("&"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanQuote(s *scanner) token.Token {
	s.accept(matcher.String("'"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanLeftParanthesis(s *scanner) token.Token {
	s.accept(matcher.String("("))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanRightParanthesis(s *scanner) token.Token {
	s.accept(matcher.String(")"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanAsterisk(s *scanner) token.Token {
	s.accept(matcher.String("*"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanPlusSign(s *scanner) token.Token {
	s.accept(matcher.String("+"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanComma(s *scanner) token.Token {
	s.accept(matcher.String(","))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanMinusSign(s *scanner) token.Token {
	s.accept(matcher.String("-"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanPeriod(s *scanner) token.Token {
	s.accept(matcher.String("."))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanSolidus(s *scanner) token.Token {
	s.accept(matcher.String("/"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanReverseSolidus(s *scanner) token.Token {
	s.accept(matcher.String("\\"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanColon(s *scanner) token.Token {
	s.accept(matcher.String(":"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanSemiColon(s *scanner) token.Token {
	s.accept(matcher.String(";"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanLessThanOperator(s *scanner) token.Token {
	s.accept(matcher.String("<"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanEqualsOperator(s *scanner) token.Token {
	s.accept(matcher.String("="))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanGreaterThanOperator(s *scanner) token.Token {
	s.accept(matcher.String("<"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

// needs more work
func scanQuestioMarkOrTrigraphs(s *scanner) token.Token {
	s.accept(matcher.String("="))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanLeftBracket(s *scanner) token.Token {
	s.accept(matcher.String("["))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanRightBracket(s *scanner) token.Token {
	s.accept(matcher.String("]"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanCircumflex(s *scanner) token.Token {
	s.accept(matcher.String("^"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanUnderscore(s *scanner) token.Token {
	s.accept(matcher.String("_"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanVerticalBar(s *scanner) token.Token {
	s.accept(matcher.String("|"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanLeftBrace(s *scanner) token.Token {
	s.accept(matcher.String("{"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanRightBrace(s *scanner) token.Token {
	s.accept(matcher.String("}"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanDollarSign(s *scanner) token.Token {
	s.accept(matcher.String("$"))
	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
}

func scanSelectOperator(s *scanner) token.Token {
	if s.acceptString("SELECT") {
		return token.New(s.line, s.col, s.pos-s.start, len("SELECT"), token.SQLSpecialCharacter, "SELECT")
	}
	return token.New(s.line, s.col, s.pos-s.start, s.pos, token.Error, fmt.Sprintf("Expected SELECT operator."))
}
