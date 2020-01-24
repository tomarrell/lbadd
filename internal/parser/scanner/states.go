package scanner

import (
	"strings"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var keywordsWithA = []string{
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
}

var keywordMapWithA map[string]token.Type = map[string]token.Type{
	"ABORT":   token.KeywordAbort,
	"ACTION":  token.KeywordAction,
	"ADD":     token.KeywordAdd,
	"AFTER":   token.KeywordAdd,
	"ALL":     token.KeywordAll,
	"ALTER":   token.KeywordAlter,
	"ANALYZE": token.KeywordAnalyze,
	"AND":     token.KeywordAnd,
	"AS":      token.KeywordAnd,
	"ASC":     token.KeywordAsc,
	"ATTACH":  token.KeywordAttach,
	"AUTO":    token.KeywordAuto,
}

var keywordsWithB = []string{
	"BEFORE",
	"BEGIN",
	"BETWEEN",
	"BY",
}

var keywordMapWithB map[string]token.Type = map[string]token.Type{
	"BEFORE":  token.KeywordBefore,
	"BEGIN":   token.KeywordBegin,
	"BETWEEN": token.KeywordBetween,
	"BY":      token.KeywordBy,
}

var keywordsWithC = []string{
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
}

var keywordMapWithC map[string]token.Type = map[string]token.Type{
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
}

var keywordsWithD = []string{
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
}

var keywordMapWithD map[string]token.Type = map[string]token.Type{
	"DATABASE":   token.KeywordDatabase,
	"DEFAULT":    token.KeywordDefault,
	"DEFERRABLE": token.KeywordDeferrable,
	"DEFERRED":   token.KeywordDeferred,
	"DELETE":     token.KeywordDelete,
	"DESC":       token.KeywordDesc,
	"DETACH":     token.KeywordDetach,
	"DISTINCT":   token.KeywordDistinct,
	"DO":         token.KeywordDo,
	"DROP":       token.KeywordDrop,
}

var keywordsWithE = []string{
	"EACH",
	"ELSE",
	"END",
	"ESCAPE",
	"EXCEPT",
	"EXCLUDE",
	"EXCLUSIVE",
	"EXISTS",
	"EXPLAIN",
}

var keywordMapWithE map[string]token.Type = map[string]token.Type{
	"EACH":      token.KeywordEach,
	"ELSE":      token.KeywordElse,
	"END":       token.KeywordEnd,
	"ESCAPE":    token.KeywordEscape,
	"EXCEPT":    token.KeywordExcept,
	"EXCLUDE":   token.KeywordExclude,
	"EXCLUSIVE": token.KeywordExclusive,
	"EXISTS":    token.KeywordExists,
	"EXPLAIN":   token.KeywordExplain,
}

var keywordsWithF = []string{
	"FAIL",
	"FILTER",
	"FIRST",
	"FOLLOWING",
	"FOR",
	"FOREIGN",
	"FROM",
	"FULL",
}

var keywordMapWithF map[string]token.Type = map[string]token.Type{
	"FAIL":      token.KeywordFail,
	"FILTER":    token.KeywordFilter,
	"FIRST":     token.KeywordFirst,
	"FOLLOWING": token.KeywordFollowing,
	"FOR":       token.KeywordFor,
	"FOREIGN":   token.KeywordForeign,
	"FROM":      token.KeywordFrom,
	"FULL":      token.KeywordFull,
}

var keywordsWithG = []string{
	"GLOB",
	"GROUP",
	"GROUPS",
}

var keywordMapWithG map[string]token.Type = map[string]token.Type{
	"GLOB":   token.KeywordGlob,
	"GROUP":  token.KeywordGroup,
	"GROUPS": token.KeywordGroups,
}

var keywordsWithH = []string{
	"HAVING",
}

var keywordMapWithH map[string]token.Type = map[string]token.Type{
	"HAVING": token.KeywordHaving,
}

var keywordsWithI = []string{
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
}

var keywordMapWithI map[string]token.Type = map[string]token.Type{
	"IF":        token.KeywordIf,
	"IGNORE":    token.KeywordIgnore,
	"IMMEDIATE": token.KeywordImmediate,
	"IN":        token.KeywordIn,
	"INDEX":     token.KeywordIndex,
	"INDEXED":   token.KeywordIndexed,
	"INITIALLY": token.KeywordInitially,
	"INNER":     token.KeywordInner,
	"INSERT":    token.KeywordInsert,
	"INSTEAD":   token.KeywordInstead,
	"INTERSECT": token.KeywordIntersect,
	"INTO":      token.KeywordInto,
	"IS":        token.KeywordIs,
	"ISNULL":    token.KeywordIsnull,
}

var keywordsWithJ = []string{
	"JOIN",
}

var keywordMapWithJ map[string]token.Type = map[string]token.Type{
	"JOIN": token.KeywordJoin,
}

var keywordsWithK = []string{
	"KEY",
}

var keywordMapWithK map[string]token.Type = map[string]token.Type{
	"KEY": token.KeywordKey,
}

var keywordsWithL = []string{
	"LAST",
	"LEFT",
	"LIKE",
	"LIMIT",
}

var keywordMapWithL map[string]token.Type = map[string]token.Type{
	"LAST":  token.KeywordLast,
	"LEFT":  token.KeywordLeft,
	"LIKE":  token.KeywordLike,
	"LIMIT": token.KeywordLimit,
}

var keywordsWithM = []string{
	"MATCH",
}

var keywordMapWithM map[string]token.Type = map[string]token.Type{
	"MATCH": token.KeywordMatch,
}

var keywordsWithN = []string{
	"NATURAL",
	"NO",
	"NOT",
	"NOTHING",
	"NOTNULL",
	"NULL",
}

var keywordMapWithN map[string]token.Type = map[string]token.Type{
	"NATURAL": token.KeywordNatural,
	"NO":      token.KeywordNo,
	"NOT":     token.KeywordNot,
	"NOTHING": token.KeywordNothing,
	"NOTNULL": token.KeywordNotnull,
	"NULL":    token.KeywordNull,
}

var keywordsWithO = []string{
	"OF",
	"OFFSET",
	"ON",
	"OR",
	"ORDER",
	"OTHERS",
	"OUTER",
	"OVER",
}

var keywordMapWithO map[string]token.Type = map[string]token.Type{
	"OF":     token.KeywordOf,
	"OFFSET": token.KeywordOffset,
	"ON":     token.KeywordOn,
	"OR":     token.KeywordOr,
	"ORDER":  token.KeywordOrder,
	"OTHERS": token.KeywordOthers,
	"OUTER":  token.KeywordOuter,
	"OVER":   token.KeywordOver,
}

var keywordsWithP = []string{
	"PARTITION",
	"PLAN",
	"PRAGMA",
	"PRECEDING",
	"PRIMARY",
}

var keywordMapWithP map[string]token.Type = map[string]token.Type{
	"PARTITION": token.KeywordPartition,
	"PLAN":      token.KeywordPlan,
	"PRAGMA":    token.KeywordPragma,
	"PRECEDING": token.KeywordPreceding,
	"PRIMARY":   token.KeywordPrimary,
}

var keywordsWithQ = []string{
	"QUERY",
}

var keywordMapWithQ map[string]token.Type = map[string]token.Type{
	"QUERY": token.KeywordQuery,
}

var keywordsWithR = []string{
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
}

var keywordMapWithR map[string]token.Type = map[string]token.Type{
	"RAISE":      token.KeywordRaise,
	"RANGE":      token.KeywordRange,
	"RECURSIVE":  token.KeywordRecursive,
	"REFERENCES": token.KeywordReferences,
	"REGEXP":     token.KeywordRegexp,
	"REINDEX":    token.KeywordReindex,
	"RELEASE":    token.KeywordRelease,
	"RENAME":     token.KeywordRename,
	"REPLACE":    token.KeywordReplace,
	"RESTRICT":   token.KeywordRestrict,
	"RIGHT":      token.KeywordRight,
	"ROLLBACK":   token.KeywordRollback,
	"ROW":        token.KeywordRow,
	"ROWS":       token.KeywordRows,
}

var keywordsWithS = []string{
	"SAVEPOINT",
	"SELECT",
	"SET",
}

var keywordMapWithS map[string]token.Type = map[string]token.Type{
	"SAVEPOINT": token.KeywordSavepoint,
	"SELECT":    token.KeywordSelect,
	"SET":       token.KeywordSet,
}

var keywordsWithT = []string{
	"TABLE",
	"TEMP",
	"TEMPORARY",
	"THEN",
	"TIES",
	"TO",
	"TRANSACTION",
	"TRIGGER",
}

var keywordMapWithT map[string]token.Type = map[string]token.Type{
	"TABLE":       token.KeywordTable,
	"TEMP":        token.KeywordTemp,
	"TEMPORARY":   token.KeywordTemporary,
	"THEN":        token.KeywordThen,
	"TIES":        token.KeywordTies,
	"TO":          token.KeywordTo,
	"TRANSACTION": token.KeywordTransaction,
	"TRIGGER":     token.KeywordTrigger,
}

var keywordsWithU = []string{
	"UNBOUNDED",
	"UNION",
	"UNIQUE",
	"UPDATE",
	"USING",
}

var keywordMapWithU map[string]token.Type = map[string]token.Type{
	"UNBOUNDED": token.KeywordUnbounded,
	"UNION":     token.KeywordUnion,
	"UNIQUE":    token.KeywordUnique,
	"UPDATE":    token.KeywordUpdate,
	"USING":     token.KeywordUsing,
}

var keywordsWithV = []string{
	"VACUUM",
	"VALUES",
	"VIEW",
	"VIRTUAL",
}

var keywordMapWithV map[string]token.Type = map[string]token.Type{
	"VACUUM":  token.KeywordVacuum,
	"VALUES":  token.KeywordValues,
	"VIEW":    token.KeywordView,
	"VIRTUAL": token.KeywordVirtual,
}

var keywordsWithW = []string{
	"WHEN",
	"WHERE",
	"WINDOW",
	"WITH",
	"WITHOUT",
}

var keywordMapWithW map[string]token.Type = map[string]token.Type{
	"WHEN":    token.KeywordWhen,
	"WHERE":   token.KeywordWhere,
	"WINDOW":  token.KeywordWindow,
	"WITH":    token.KeywordWith,
	"WITHOUT": token.KeywordWithout,
}

func scanSpace(s *scanner) {
	s.accept(matcher.String(" "))
}

// func scanDoubleQuote(s *scanner) token.Token {
// 	s.accept(matcher.String("\""))
// 	return createToken(token.SQLSpecialCharacter)
// }

// func scanPercent(s *scanner) token.Token {
// 	s.accept(matcher.String("%"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanAmpersand(s *scanner) token.Token {
// 	s.accept(matcher.String("&"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
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

// func scanAsterisk(s *scanner) token.Token {
// 	s.accept(matcher.String("*"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanPlusSign(s *scanner) token.Token {
// 	s.accept(matcher.String("+"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanComma(s *scanner) token.Token {
// 	s.accept(matcher.String(":true,"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanMinusSign(s *scanner) token.Token {
// 	s.accept(matcher.String("-"))
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

// func scanLessThanOperator(s *scanner) token.Token {
// 	s.accept(matcher.String("<"))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanEqualsOperator(s *scanner) token.Token {
// 	s.accept(matcher.String("="))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanGreaterThanOperator(s *scanner) token.Token {
// 	s.accept(matcher.String("<"))
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

func scanAKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithA {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithA[k])
		}
	}
	return s.unexpectedRune()
}

func scanBKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithB {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithB[k])
		}
	}
	return s.unexpectedRune()
}

func scanCKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithC {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithC[k])
		}
	}
	return s.unexpectedRune()
}

func scanDKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithD {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithD[k])
		}
	}
	return s.unexpectedRune()
}

func scanEKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithE {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithE[k])
		}
	}
	return s.unexpectedRune()
}

func scanFKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithF {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithF[k])
		}
	}
	return s.unexpectedRune()
}

func scanGKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithG {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithG[k])
		}
	}
	return s.unexpectedRune()
}

func scanHKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithH {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithH[k])
		}
	}
	return s.unexpectedRune()
}

func scanIKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithI {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithI[k])
		}
	}
	return s.unexpectedRune()
}

func scanJKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithJ {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithJ[k])
		}
	}
	return s.unexpectedRune()
}

func scanKKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithK {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithK[k])
		}
	}
	return s.unexpectedRune()
}

func scanLKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithL {
		keyword := strings.Split(k, "")
		j := 0
		for i := range keyword {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithL[k])
		}
	}
	s.acceptString(string(input))
	return s.createToken(token.KeywordAbort)
}

func scanMKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithM {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithM[k])
		}
	}
	return s.unexpectedRune()
}

func scanNKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithN {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithN[k])
		}
	}
	return s.unexpectedRune()
}

func scanOKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithO {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithO[k])
		}
	}
	return s.unexpectedRune()
}

func scanPKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithP {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithP[k])
		}
	}
	return s.unexpectedRune()
}

func scanQKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithQ {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithQ[k])
		}
	}
	return s.unexpectedRune()
}

func scanRKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithR {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithR[k])
		}
	}
	return s.unexpectedRune()
}

func scanSKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithS {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithS[k])
		}
	}
	return s.unexpectedRune()
}

func scanTKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithT {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithT[k])
		}
	}
	return s.unexpectedRune()
}

func scanUKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithU {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithU[k])
		}
	}
	return s.unexpectedRune()
}
func scanVKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithV {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithV[k])
		}
	}
	return s.unexpectedRune()
}

func scanWKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := s.input[s.start:nextRune]
	for _, k := range keywordsWithW {
		keyword := strings.Split(k, "")
		j := 0
		length := len(input)
		if length > len(keyword) {
			length = len(keyword)
		}
		for i := 0; i < length; i++ {
			if keyword[i] == string(input[i]) {
				j++
			}
		}
		if j == len(keyword) {
			s.acceptString(string(input))
			return s.createToken(keywordMapWithW[k])
		}
	}
	return s.unexpectedRune()
}
