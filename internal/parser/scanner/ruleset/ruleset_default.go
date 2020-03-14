package ruleset

import (
	"bytes"
	"unicode"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

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
	// the order of the rules are important for some cases. Beware
	defaultRules = []Rule{
		FuncRule(defaultStatementSeparatorRule),
		FuncRule(defaultKeywordsRule),
		FuncRule(defaultUnaryOperatorRule),
		FuncRule(defaultBinaryOperatorRule),
		FuncRule(defaultDelimiterRule),
		FuncRule(defaultQuotedLiteralRule),
		FuncRule(defaultNumericLiteralRule),
		FuncRule(defaultUnquotedLiteralRule),
	}
	defaultKeywords = map[string]token.Type{"ABORT": token.KeywordAbort, "ACTION": token.KeywordAction, "ADD": token.KeywordAdd, "AFTER": token.KeywordAfter, "ALL": token.KeywordAll, "ALTER": token.KeywordAlter, "ANALYZE": token.KeywordAnalyze, "AND": token.KeywordAnd, "AS": token.KeywordAs, "ASC": token.KeywordAsc, "ATTACH": token.KeywordAttach, "AUTOINCREMENT": token.KeywordAutoincrement, "BEFORE": token.KeywordBefore, "BEGIN": token.KeywordBegin, "BETWEEN": token.KeywordBetween, "BY": token.KeywordBy, "CASCADE": token.KeywordCascade, "CASE": token.KeywordCase, "CAST": token.KeywordCast, "CHECK": token.KeywordCheck, "COLLATE": token.KeywordCollate, "COLUMN": token.KeywordColumn, "COMMIT": token.KeywordCommit, "CONFLICT": token.KeywordConflict, "CONSTRAINT": token.KeywordConstraint, "CREATE": token.KeywordCreate, "CROSS": token.KeywordCross, "CURRENT": token.KeywordCurrent, "CURRENT_DATE": token.KeywordCurrentDate, "CURRENT_TIME": token.KeywordCurrentTime, "CURRENT_TIMESTAMP": token.KeywordCurrentTimestamp, "DATABASE": token.KeywordDatabase, "DEFAULT": token.KeywordDefault, "DEFERRABLE": token.KeywordDeferrable, "DEFERRED": token.KeywordDeferred, "DELETE": token.KeywordDelete, "DESC": token.KeywordDesc, "DETACH": token.KeywordDetach, "DISTINCT": token.KeywordDistinct, "DO": token.KeywordDo, "DROP": token.KeywordDrop, "EACH": token.KeywordEach, "ELSE": token.KeywordElse, "END": token.KeywordEnd, "ESCAPE": token.KeywordEscape, "EXCEPT": token.KeywordExcept, "EXCLUDE": token.KeywordExclude, "EXCLUSIVE": token.KeywordExclusive, "EXISTS": token.KeywordExists, "EXPLAIN": token.KeywordExplain, "FAIL": token.KeywordFail, "FILTER": token.KeywordFilter, "FIRST": token.KeywordFirst, "FOLLOWING": token.KeywordFollowing, "FOR": token.KeywordFor, "FOREIGN": token.KeywordForeign, "FROM": token.KeywordFrom, "FULL": token.KeywordFull, "GLOB": token.KeywordGlob, "GROUP": token.KeywordGroup, "GROUPS": token.KeywordGroups, "HAVING": token.KeywordHaving, "IF": token.KeywordIf, "IGNORE": token.KeywordIgnore, "IMMEDIATE": token.KeywordImmediate, "IN": token.KeywordIn, "INDEX": token.KeywordIndex, "INDEXED": token.KeywordIndexed, "INITIALLY": token.KeywordInitially, "INNER": token.KeywordInner, "INSERT": token.KeywordInsert, "INSTEAD": token.KeywordInstead, "INTERSECT": token.KeywordIntersect, "INTO": token.KeywordInto, "IS": token.KeywordIs, "ISNULL": token.KeywordIsnull, "JOIN": token.KeywordJoin, "KEY": token.KeywordKey, "LAST": token.KeywordLast, "LEFT": token.KeywordLeft, "LIKE": token.KeywordLike, "LIMIT": token.KeywordLimit, "MATCH": token.KeywordMatch, "NATURAL": token.KeywordNatural, "NO": token.KeywordNo, "NOT": token.KeywordNot, "NOTHING": token.KeywordNothing, "NOTNULL": token.KeywordNotnull, "NULL": token.KeywordNull, "OF": token.KeywordOf, "OFFSET": token.KeywordOffset, "ON": token.KeywordOn, "OR": token.KeywordOr, "ORDER": token.KeywordOrder, "OTHERS": token.KeywordOthers, "OUTER": token.KeywordOuter, "OVER": token.KeywordOver, "PARTITION": token.KeywordPartition, "PLAN": token.KeywordPlan, "PRAGMA": token.KeywordPragma, "PRECEDING": token.KeywordPreceding, "PRIMARY": token.KeywordPrimary, "QUERY": token.KeywordQuery, "RAISE": token.KeywordRaise, "RANGE": token.KeywordRange, "RECURSIVE": token.KeywordRecursive, "REFERENCES": token.KeywordReferences, "REGEXP": token.KeywordRegexp, "REINDEX": token.KeywordReindex, "RELEASE": token.KeywordRelease, "RENAME": token.KeywordRename, "REPLACE": token.KeywordReplace, "RESTRICT": token.KeywordRestrict, "RIGHT": token.KeywordRight, "ROLLBACK": token.KeywordRollback, "ROW": token.KeywordRow, "ROWS": token.KeywordRows, "SAVEPOINT": token.KeywordSavepoint, "SELECT": token.KeywordSelect, "SET": token.KeywordSet, "TABLE": token.KeywordTable, "TEMP": token.KeywordTemp, "TEMPORARY": token.KeywordTemporary, "THEN": token.KeywordThen, "TIES": token.KeywordTies, "TO": token.KeywordTo, "TRANSACTION": token.KeywordTransaction, "TRIGGER": token.KeywordTrigger, "UNBOUNDED": token.KeywordUnbounded, "UNION": token.KeywordUnion, "UNIQUE": token.KeywordUnique, "UPDATE": token.KeywordUpdate, "USING": token.KeywordUsing, "VACUUM": token.KeywordVacuum, "VALUES": token.KeywordValues, "VIEW": token.KeywordView, "VIRTUAL": token.KeywordVirtual, "WHEN": token.KeywordWhen, "WHERE": token.KeywordWhere, "WINDOW": token.KeywordWindow, "WITH": token.KeywordWith, "WITHOUT": token.KeywordWithout}
)

func defaultStatementSeparatorRule(s RuneScanner) (token.Type, bool) {
	if next, ok := s.Lookahead(); ok && defaultStatementSeparator.Matches(next) {
		s.ConsumeRune()
		return token.StatementSeparator, true
	}
	return token.Unknown, false
}

func defaultKeywordsRule(s RuneScanner) (token.Type, bool) {
	// read word
	var buf bytes.Buffer
	for {
		next, ok := s.Lookahead()
		if !ok || defaultWhitespaceDetector.Matches(next) || !defaultLiteral.Matches(next) {
			break
		}
		_, _ = buf.WriteRune(next)
		s.ConsumeRune()
	}
	candidate := buf.String() // candidate is the next word that may be a keyword

	// check if the candidate is a keyword
	if typ, ok := defaultKeywords[candidate]; ok {
		return typ, true
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
				return defaultUnquotedLiteralRule(s)
			}
		}
	}
	for {
		// in the above case, if there was a `0.34` as a part of the string to read,
		// the loop would have broken without consuming the rune; because the above
		// loop consumes only hexadecimals of form `0xNUMBER`. Since all other cases
		// are not consumed, the `LookAhead` below, gets the previous rune conveniently.
		next, ok := s.Lookahead()
		// continue on cases where the decimal point/exponent/exponent opeartor is already not found or not this particular rune.
		if !(ok && (defaultNumericLiteral.Matches(next) || (!decimalPointFlag && defaultDecimalPoint.Matches(next)) || (!exponentFlag && defaultExponent.Matches(next)) || (!exponentOperatorFlag && defaultExponentOperator.Matches(next)))) {
			break
		}
		if next == '.' {
			if decimalPointFlag {
				return token.Unknown, false
			}
			decimalPointFlag = true
		}
		if next == 'E' {
			if exponentFlag {
				return token.Unknown, false
			}
			exponentFlag = true
		}
		if next == '+' || next == '-' {
			if exponentOperatorFlag {
				return token.Unknown, false
			}
			exponentOperatorFlag = true
		}
		if defaultDecimalPoint.Matches(next) {
			decimalPointFlag = true
		}
		s.ConsumeRune()
	}
	return token.Literal, true
}
