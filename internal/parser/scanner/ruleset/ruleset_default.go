package ruleset

import (
	"unicode"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

//go:generate go run ../../../tool/generate/keywordtrie ./ruleset_default_keyword_trie.go

var (
	// Default is the ruleset that this application uses by default. The rules
	// are inspired by Sqlite, however they do not care as much about
	// compatibility with other database systems, and are therefore simpler to
	// read and write.
	Default = Ruleset{
		WhitespaceDetector: defaultWhitespaceDetector.Matches,
		LinefeedDetector:   defaultLinefeedDetector.Matches,
		Rules:              defaultRules,
	}
	// defaultWhitespaceDetector matches all the the whitespaces that this ruleset allows
	defaultWhitespaceDetector = matcher.New("whitespace", unicode.Space)
	// defaultLinefeedDetector is the linefeed detector that this ruleset allows
	defaultLinefeedDetector   = matcher.RuneWithDesc("linefeed", '\n')
	defaultStatementSeparator = matcher.RuneWithDesc("statement separator", ';')
	defaultDecimalPoint       = matcher.RuneWithDesc("decimal point", '.')
	defaultExponent           = matcher.RuneWithDesc("exponent indicator", 'E')
	defaultExponentOperator   = matcher.String("+-")
	defaultNumber             = matcher.New("number", unicode.Number)
	// defaultLiteral matches the allowed letters of a literal
	defaultLiteral = matcher.Merge(
		matcher.New("upper", unicode.Upper),
		matcher.New("lower", unicode.Lower),
		matcher.New("title", unicode.Title),
		matcher.String("-_"),
		defaultNumber,
	)
	defaultNumericLiteral = matcher.Merge(
		defaultNumber,
		defaultExponent,
		defaultExponentOperator,
		matcher.RuneWithDesc("X", 'x'),
	)
	defaultQuote          = matcher.String("'\"")
	defaultUnaryOperator  = matcher.String("-+~")
	defaultBinaryOperator = matcher.String("|*/%<>=&!")
	defaultDelimiter      = matcher.String("(),")
	defaultPlaceholder    = matcher.RuneWithDesc("placeholder", '?')
	// the order of the rules are important for some cases. Beware
	defaultRules = []Rule{
		FuncRule(defaultStatementSeparatorRule),
		FuncRule(defaultPlaceholderRule),
		FuncRule(defaultKeywordsRule),
		FuncRule(defaultUnaryOperatorRule),
		FuncRule(defaultBinaryOperatorRule),
		FuncRule(defaultDelimiterRule),
		FuncRule(defaultQuotedLiteralRule),
		FuncRule(defaultNumericLiteralRule),
		FuncRule(defaultUnquotedLiteralRule),
		FuncRule(defaultDecimalPointRule),
	}
)

func defaultStatementSeparatorRule(s RuneScanner) (token.Type, bool) {
	if next, ok := s.Lookahead(); ok && defaultStatementSeparator.Matches(next) {
		s.ConsumeRune()
		return token.StatementSeparator, true
	}
	return token.Unknown, false
}

func defaultPlaceholderRule(s RuneScanner) (token.Type, bool) {
	if next, ok := s.Lookahead(); ok && defaultPlaceholder.Matches(next) {
		s.ConsumeRune()
		return token.Literal, true
	}
	return token.Unknown, false
}

func defaultUnaryOperatorRule(s RuneScanner) (token.Type, bool) {
	if next, ok := s.Lookahead(); ok && defaultUnaryOperator.Matches(next) {
		s.ConsumeRune()
		return token.UnaryOperator, true
	}
	return token.Unknown, false
}

// defaultBinaryOperatorRule scans binary operators. Since binary opeators
// can be of length one and two, it first checks for former, then latter
// and allows the best of the two conditions.
func defaultBinaryOperatorRule(s RuneScanner) (token.Type, bool) {
	if first, ok := s.Lookahead(); ok && defaultBinaryOperator.Matches(first) {
		s.ConsumeRune()
		if first == '*' || first == '/' || first == '%' || first == '&' {
			return token.BinaryOperator, true
		} else if next, ok := s.Lookahead(); ok && defaultBinaryOperator.Matches(next) {
			switch first {
			case '|':
				if next == '|' {
					s.ConsumeRune()
					return token.BinaryOperator, true
				}
			case '<':
				if next == '<' || next == '=' || next == '>' {
					s.ConsumeRune()
					return token.BinaryOperator, true
				}
			case '>':
				if next == '>' || next == '=' {
					s.ConsumeRune()
					return token.BinaryOperator, true
				}
			case '=':
				if next == '=' {
					s.ConsumeRune()
					return token.BinaryOperator, true
				}
			case '!':
				if next == '=' {
					s.ConsumeRune()
					return token.BinaryOperator, true
				}
			}
		}
		// special cases where these operators can be single or can have a suffix
		// The switch case blocks have been designed such that only the cases where
		// there's a meaningful operator, the runes are consumed and returned, in cases
		// where the second operator is not meaningful, the first operator is consumed here.
		// In cases where the second operator is not meaningful, it'll be processed again.
		if first == '<' || first == '>' || first == '=' || first == '|' {
			return token.BinaryOperator, true
		}
	}
	return token.Unknown, false
}

func defaultDelimiterRule(s RuneScanner) (token.Type, bool) {
	if next, ok := s.Lookahead(); ok && defaultDelimiter.Matches(next) {
		s.ConsumeRune()
		return token.Delimiter, true
	}
	return token.Unknown, false
}

func defaultQuotedLiteralRule(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !(ok && defaultQuote.Matches(next)) {
		return token.Unknown, false
	}
	quoteType := next
	s.ConsumeRune()

	for {
		next, ok := s.Lookahead()
		if !ok {
			return token.Unknown, false
		}
		// allowing character escape
		if next == '\\' {
			s.ConsumeRune()
			_, ok = s.Lookahead()
			if !ok {
				return token.Unknown, false
			}
			s.ConsumeRune()
		} else if next == quoteType {
			break
		} else {
			s.ConsumeRune()
		}
	}
	s.ConsumeRune()
	return token.Literal, true
}

func defaultUnquotedLiteralRule(s RuneScanner) (token.Type, bool) {
	if next, ok := s.Lookahead(); !(ok && defaultLiteral.Matches(next)) {
		return token.Unknown, false
	}
	s.ConsumeRune()

	for {
		next, ok := s.Lookahead()
		if !(ok && defaultLiteral.Matches(next)) {
			break
		}
		s.ConsumeRune()
	}

	return token.Literal, true
}

func defaultNumericLiteralRule(s RuneScanner) (token.Type, bool) {
	decimalPointFlag := false
	exponentFlag := false
	exponentOperatorFlag := false
	numericLiteralFlag := false
	// Checking whether the first element is a number or a decimal point.
	// If neither, an unknown token error is raised.
	next, ok := s.Lookahead()
	if !(ok && (defaultNumericLiteral.Matches(next) || defaultDecimalPoint.Matches(next))) {
		return token.Unknown, false
	}
	// If the literal starts with a decimal point, it is recorded in the flag.
	if defaultDecimalPoint.Matches(next) {
		decimalPointFlag = true
	}
	if defaultNumericLiteral.Matches(next) {
		numericLiteralFlag = true
	}
	s.ConsumeRune()
	// case of hexadecimal numbers
	if next == '0' {
		for {
			next, ok = s.Lookahead()
			if !ok || next != 'x' {
				break
			}
			if next == 'x' {
				s.ConsumeRune()
				if next, ok := s.Lookahead(); !(ok && defaultLiteral.Matches(next)) {
					return token.Unknown, false
				}
				s.ConsumeRune()

				for {
					next, ok := s.Lookahead()
					if !(ok && defaultLiteral.Matches(next)) {
						break
					}
					s.ConsumeRune()
				}
				return token.LiteralNumeric, true
			}
		}
	}

	for {
		// in the above case, if there was a `0.34` as a part of the string to read,
		// the loop would have broken without consuming the rune; because the above
		// loop consumes only hexadecimals of form `0xNUMBER`. Since all other cases
		// are not consumed, the `LookAhead` below, gets the previous rune conveniently.
		next, ok := s.Lookahead()
		// continue in case the decimal point/exponent/exponent operator is already not found or not this particular rune.
		if !(ok && (defaultNumericLiteral.Matches(next) || (!decimalPointFlag && defaultDecimalPoint.Matches(next)) || (!exponentFlag && defaultExponent.Matches(next)) || (!exponentOperatorFlag && defaultExponentOperator.Matches(next)))) {
			break
		}
		switch next {
		case '.':
			if decimalPointFlag {
				return token.Unknown, false
			}
			decimalPointFlag = true
			s.ConsumeRune()
		case 'E':
			if exponentFlag {
				return token.Unknown, false
			}
			exponentFlag = true
			s.ConsumeRune()
		case '+', '-':
			// only allow `+` or `-` in case of `E+x` or `E-x`.
			if exponentFlag {
				if exponentOperatorFlag {
					return token.Unknown, false
				}
				exponentOperatorFlag = true
				s.ConsumeRune()
			} else {
				return token.LiteralNumeric, true
			}
		default:
			if defaultNumericLiteral.Matches(next) {
				numericLiteralFlag = true
			}
			s.ConsumeRune()
		}
	}
	// This case checks for "." passing as numericLiterals
	if decimalPointFlag && !numericLiteralFlag {
		return token.Unknown, false
	}
	return token.LiteralNumeric, true
}

func defaultDecimalPointRule(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !(ok && defaultDecimalPoint.Matches(next)) {
		return token.Unknown, false
	}
	s.ConsumeRune()
	return token.Literal, true
}
