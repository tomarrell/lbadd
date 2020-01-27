package scanner

import (
	"strings"
	"unicode"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var keywordSlice = []string{
	"ABORT",
	"ACTION",
	"ADD",
	"AFTER",
	"ALL",
	"ALTER",
	"ANALYZE",
	"AND",
	"AS",
	"ASC",
	"ATTACH",
	"AUTO",
	"BEFORE",
	"BEGIN",
	"BETWEEN",
	"BY",
	"CASCADE",
	"CASE",
	"CAST",
	"CHECK",
	"COLLATE",
	"COLUMN",
	"COMMIT",
	"CONFLICT",
	"CONSTRAINT",
	"CREATE",
	"CROSS",
	"CURRENT",
	"CURRENT_DATE",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"DATABASE",
	"DEFAULT",
	"DEFERRABLE",
	"DEFERRED",
	"DELETE",
	"DESC",
	"DETACH",
	"DISTINCT",
	"DO",
	"DROP",
	"EACH",
	"ELSE",
	"END",
	"ESCAPE",
	"EXCEPT",
	"EXCLUDE",
	"EXCLUSIVE",
	"EXISTS",
	"EXPLAIN",
	"FAIL",
	"FILTER",
	"FIRST",
	"FOLLOWING",
	"FOR",
	"FOREIGN",
	"FROM",
	"FULL",
	"GLOB",
	"GROUP",
	"GROUPS",
	"HAVING",
	"IF",
	"IGNORE",
	"IMMEDIATE",
	"IN",
	"INDEX",
	"INDEXED",
	"INITIALLY",
	"INNER",
	"INSERT",
	"INSTEAD",
	"INTERSECT",
	"INTO",
	"IS",
	"ISNULL",
	"JOIN",
	"KEY",
	"LAST",
	"LEFT",
	"LIKE",
	"LIMIT",
	"MATCH",
	"NATURAL",
	"NO",
	"NOT",
	"NOTHING",
	"NOTNULL",
	"OF",
	"OFFSET",
	"ON",
	"OR",
	"ORDER",
	"OTHERS",
	"OUTER",
	"OVER",
	"PARTITION",
	"PLAN",
	"PRAGMA",
	"PRECEDING",
	"PRIMARY",
	"QUERY",
	"RAISE",
	"RANGE",
	"RECURSIVE",
	"REFERENCES",
	"REGEXP",
	"REINDEX",
	"RELEASE",
	"RENAME",
	"REPLACE",
	"RESTRICT",
	"RIGHT",
	"ROLLBACK",
	"ROW",
	"ROWS",
	"SAVEPOINT",
	"SELECT",
	"SET",
	"TABLE",
	"TEMP",
	"TEMPORARY",
	"THEN",
	"TIES",
	"TO",
	"TRANSACTION",
	"TRIGGER",
	"UNBOUNDED",
	"UNION",
	"UNIQUE",
	"UPDATE",
	"USING",
	"VACUUM",
	"VALUES",
	"VIEW",
	"VIRTUAL",
	"WHEN",
	"WHERE",
	"WINDOW",
	"WITH",
	"WITHOUT",
}

var keywordMap map[string]token.Type = map[string]token.Type{
	"ABORT":             token.KeywordAbort,
	"ACTION":            token.KeywordAction,
	"ADD":               token.KeywordAdd,
	"AFTER":             token.KeywordAdd,
	"ALL":               token.KeywordAll,
	"ALTER":             token.KeywordAlter,
	"ANALYZE":           token.KeywordAnalyze,
	"AND":               token.KeywordAnd,
	"AS":                token.KeywordAnd,
	"ASC":               token.KeywordAsc,
	"ATTACH":            token.KeywordAttach,
	"AUTO":              token.KeywordAuto,
	"BEFORE":            token.KeywordBefore,
	"BEGIN":             token.KeywordBegin,
	"BETWEEN":           token.KeywordBetween,
	"BY":                token.KeywordBy,
	"CASCADE":           token.KeywordCascade,
	"CASE":              token.KeywordCase,
	"CAST":              token.KeywordCast,
	"CHECK":             token.KeywordCheck,
	"COLLATE":           token.KeywordCollate,
	"COLUMN":            token.KeywordColumn,
	"COMMIT":            token.KeywordCommit,
	"CONFLICT":          token.KeywordConflict,
	"CONSTRAINT":        token.KeywordConstraint,
	"CREATE":            token.KeywordCreate,
	"CROSS":             token.KeywordCross,
	"CURRENT":           token.KeywordCurrent,
	"CURRENT_DATE":      token.KeywordCurrentDate,
	"CURRENT_TIME":      token.KeywordCurrentTime,
	"CURRENT_TIMESTAMP": token.KeywordCurrentTimestamp,
	"DATABASE":          token.KeywordDatabase,
	"DEFAULT":           token.KeywordDefault,
	"DEFERRABLE":        token.KeywordDeferrable,
	"DEFERRED":          token.KeywordDeferred,
	"DELETE":            token.KeywordDelete,
	"DESC":              token.KeywordDesc,
	"DETACH":            token.KeywordDetach,
	"DISTINCT":          token.KeywordDistinct,
	"DO":                token.KeywordDo,
	"DROP":              token.KeywordDrop,
	"EACH":              token.KeywordEach,
	"ELSE":              token.KeywordElse,
	"END":               token.KeywordEnd,
	"ESCAPE":            token.KeywordEscape,
	"EXCEPT":            token.KeywordExcept,
	"EXCLUDE":           token.KeywordExclude,
	"EXCLUSIVE":         token.KeywordExclusive,
	"EXISTS":            token.KeywordExists,
	"EXPLAIN":           token.KeywordExplain,
	"FAIL":              token.KeywordFail,
	"FILTER":            token.KeywordFilter,
	"FIRST":             token.KeywordFirst,
	"FOLLOWING":         token.KeywordFollowing,
	"FOR":               token.KeywordFor,
	"FOREIGN":           token.KeywordForeign,
	"FROM":              token.KeywordFrom,
	"FULL":              token.KeywordFull,
	"GLOB":              token.KeywordGlob,
	"GROUP":             token.KeywordGroup,
	"GROUPS":            token.KeywordGroups,
	"HAVING":            token.KeywordHaving,
	"IF":                token.KeywordIf,
	"IGNORE":            token.KeywordIgnore,
	"IMMEDIATE":         token.KeywordImmediate,
	"IN":                token.KeywordIn,
	"INDEX":             token.KeywordIndex,
	"INDEXED":           token.KeywordIndexed,
	"INITIALLY":         token.KeywordInitially,
	"INNER":             token.KeywordInner,
	"INSERT":            token.KeywordInsert,
	"INSTEAD":           token.KeywordInstead,
	"INTERSECT":         token.KeywordIntersect,
	"INTO":              token.KeywordInto,
	"IS":                token.KeywordIs,
	"ISNULL":            token.KeywordIsnull,
	"JOIN":              token.KeywordJoin,
	"KEY":               token.KeywordKey,
	"LAST":              token.KeywordLast,
	"LEFT":              token.KeywordLeft,
	"LIKE":              token.KeywordLike,
	"LIMIT":             token.KeywordLimit,
	"MATCH":             token.KeywordMatch,
	"NATURAL":           token.KeywordNatural,
	"NO":                token.KeywordNo,
	"NOT":               token.KeywordNot,
	"NOTHING":           token.KeywordNothing,
	"NOTNULL":           token.KeywordNotnull,
	"OF":                token.KeywordOf,
	"OFFSET":            token.KeywordOffset,
	"ON":                token.KeywordOn,
	"OR":                token.KeywordOr,
	"ORDER":             token.KeywordOrder,
	"OTHERS":            token.KeywordOthers,
	"OUTER":             token.KeywordOuter,
	"OVER":              token.KeywordOver,
	"PARTITION":         token.KeywordPartition,
	"PLAN":              token.KeywordPlan,
	"PRAGMA":            token.KeywordPragma,
	"PRECEDING":         token.KeywordPreceding,
	"PRIMARY":           token.KeywordPrimary,
	"QUERY":             token.KeywordQuery,
	"RAISE":             token.KeywordRaise,
	"RANGE":             token.KeywordRange,
	"RECURSIVE":         token.KeywordRecursive,
	"REFERENCES":        token.KeywordReferences,
	"REGEXP":            token.KeywordRegexp,
	"REINDEX":           token.KeywordReindex,
	"RELEASE":           token.KeywordRelease,
	"RENAME":            token.KeywordRename,
	"REPLACE":           token.KeywordReplace,
	"RESTRICT":          token.KeywordRestrict,
	"RIGHT":             token.KeywordRight,
	"ROLLBACK":          token.KeywordRollback,
	"ROW":               token.KeywordRow,
	"ROWS":              token.KeywordRows,
	"SAVEPOINT":         token.KeywordSavepoint,
	"SELECT":            token.KeywordSelect,
	"SET":               token.KeywordSet,
	"TABLE":             token.KeywordTable,
	"TEMP":              token.KeywordTemp,
	"TEMPORARY":         token.KeywordTemporary,
	"THEN":              token.KeywordThen,
	"TIES":              token.KeywordTies,
	"TO":                token.KeywordTo,
	"TRANSACTION":       token.KeywordTransaction,
	"TRIGGER":           token.KeywordTrigger,
	"UNBOUNDED":         token.KeywordUnbounded,
	"UNION":             token.KeywordUnion,
	"UNIQUE":            token.KeywordUnique,
	"UPDATE":            token.KeywordUpdate,
	"USING":             token.KeywordUsing,
	"VACUUM":            token.KeywordVacuum,
	"VALUES":            token.KeywordValues,
	"VIEW":              token.KeywordView,
	"VIRTUAL":           token.KeywordVirtual,
	"WHEN":              token.KeywordWhen,
	"WHERE":             token.KeywordWhere,
	"WINDOW":            token.KeywordWindow,
	"WITH":              token.KeywordWith,
	"WITHOUT":           token.KeywordWithout,
}

var (
	whiteSpace            = matcher.Merge(formFeed, noBreakSpace, space, horizontalTab, unicodeSpace, verticalTab, zeroWidthJoiner, zeroWidthNoBreakSpace, zeroWidthNonJoiner)
	formFeed              = matcher.RuneWithDesc("<FF>", '\u000C')
	noBreakSpace          = matcher.RuneWithDesc("<NBSP>", '\u00A0')
	space                 = matcher.RuneWithDesc("<SP>", '\u0020')
	horizontalTab         = matcher.RuneWithDesc("<TAB>", '\u0009')
	unicodeSpace          = matcher.New("<USP>", unicode.Z)
	verticalTab           = matcher.RuneWithDesc("<VT>", '\u000B')
	zeroWidthJoiner       = matcher.RuneWithDesc("<ZWJ>", '\u200D')
	zeroWidthNoBreakSpace = matcher.RuneWithDesc("<ZWNBSP>", '\uFEFF')
	zeroWidthNonJoiner    = matcher.RuneWithDesc("<ZWNJ>", '\u200C')
)

func (s *scanner) consumeRune() {
	s.col++
	s.accept(matcher.String(" "))
}

// func scanDoubleQuote(s *scanner) token.Token {
// 	s.accept(matcher.String("\""))
// 	return createToken(token.SQLSpecialCharacter)
// }

// func scanQuote(s *scanner) token.Token {
// 	s.accept(matcher.String("'"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanLeftParanthesis(s *scanner) token.Token {
// 	s.accept(matcher.String("("))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanRightParanthesis(s *scanner) token.Token {
// 	s.accept(matcher.String(")"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanComma(s *scanner) token.Token {
// 	s.accept(matcher.String(":true,"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanPeriod(s *scanner) token.Token {
// 	s.accept(matcher.String("."))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanSolidus(s *scanner) token.Token {
// 	s.accept(matcher.String("/"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanReverseSolidus(s *scanner) token.Token {
// 	s.accept(matcher.String("\\"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanColon(s *scanner) token.Token {
// 	s.accept(matcher.String(":"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanSemiColon(s *scanner) token.Token {
// 	s.accept(matcher.String(";"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// // needs more work
// func scanQuestioMarkOrTrigraphs(s *scanner) token.Token {
// 	s.accept(matcher.String("="))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanLeftBracket(s *scanner) token.Token {
// 	s.accept(matcher.String("["))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanRightBracket(s *scanner) token.Token {
// 	s.accept(matcher.String("]"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanCircumflex(s *scanner) token.Token {
// 	s.accept(matcher.String("^"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanUnderscore(s *scanner) token.Token {
// 	s.accept(matcher.String("_"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanVerticalBar(s *scanner) token.Token {
// 	s.accept(matcher.String("|"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanLeftBrace(s *scanner) token.Token {
// 	s.accept(matcher.String("{"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanRightBrace(s *scanner) token.Token {
// 	s.accept(matcher.String("}"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanDollarSign(s *scanner) token.Token {
// 	s.accept(matcher.String("$"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

func (s *scanner) scanKeyword() token.Token {
	nextPos := s.seekNext(s.start)
	input := []rune(strings.ToUpper(string(s.input[s.start:nextPos])))
	for _, k := range keywordSlice {
		keyword := []rune(k)
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == input[i] {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(s.input[s.start:nextPos]))
			return s.createToken(keywordMap[k])
		}
	}
	s.start = nextPos
	s.pos = nextPos
	return s.unexpectedRune(string(input))
}

func (s *scanner) scanOperator() token.Token {
	return s.unexpectedRune("nil")
}

func (s *scanner) scanLiteral() token.Token {
	return s.unexpectedRune("nil")
}

func (s *scanner) scanSpace() token.Token {
	if s.input[s.start] == '\n' {
		s.line++
		s.col = 1
	}
	s.pos++
	s.start = s.pos
	return nil
}
