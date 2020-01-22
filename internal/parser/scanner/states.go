package scanner

import (
	"fmt"

	// "github.com/dghubble/trie"
	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// LoadTrie inserts all keywords on to the trie
func LoadTrie() {
	trie := NewRuneTrie()
	for k, v := range keywordsWithS {
		trie.Put(k, v)
	}
	fmt.Println(trie.Get("SELECS"))
}

var keywordsWithA map[string]token.Type = map[string]token.Type{
	"ABORT":          token.KeywordAbort,
	"ACTION":         token.KeywordAction,
	"ADD":            token.KeywordAdd,
	"AFTER":          token.KeywordAdd,
	"ALL":            token.KeywordAll,
	"ALTER":          token.KeywordAlter,
	"ANALYZE":        token.KeywordAnalyze,
	"AND":            token.KeywordAnd,
	"AS":             token.KeywordAnd,
	"ASC":            token.KeywordAsc,
	"ATTACH":         token.KeywordAttach,
	"AUTO INCREMENT": token.KeywordAutoincrement,
}

var keywordsWithB map[string]token.Type = map[string]token.Type{
	"BEFORE":  token.KeywordBefore,
	"BEGIN":   token.KeywordBegin,
	"BETWEEN": token.KeywordBetween,
	"BY":      token.KeywordBy,
}

var keywordsWithC map[string]token.Type = map[string]token.Type{
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

var keywordsWithD map[string]token.Type = map[string]token.Type{
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

var keywordsWithE map[string]token.Type = map[string]token.Type{
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

var keywordsWithF map[string]token.Type = map[string]token.Type{
	"FAIL":      token.KeywordFail,
	"FILTER":    token.KeywordFilter,
	"FIRST":     token.KeywordFirst,
	"FOLLOWING": token.KeywordFollowing,
	"FOR":       token.KeywordFor,
	"FOREIGN":   token.KeywordForeign,
	"FROM":      token.KeywordFrom,
	"FULL":      token.KeywordFull,
}

var keywordsWithG map[string]token.Type = map[string]token.Type{
	"GLOB":   token.KeywordGlob,
	"GROUP":  token.KeywordGroup,
	"GROUPS": token.KeywordGroups,
}

var keywordsWithH map[string]token.Type = map[string]token.Type{
	"HAVING": token.KeywordHaving,
}

var keywordsWithI map[string]token.Type = map[string]token.Type{
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

var keywordsWithJ map[string]token.Type = map[string]token.Type{
	"JOIN": token.KeywordJoin,
}

var keywordsWithK map[string]token.Type = map[string]token.Type{
	"KEY": token.KeywordKey,
}

var keywordsWithL map[string]token.Type = map[string]token.Type{
	"LAST":  token.KeywordLast,
	"LEFT":  token.KeywordLeft,
	"LIKE":  token.KeywordLike,
	"LIMIT": token.KeywordLimit,
}

var keywordsWithM map[string]token.Type = map[string]token.Type{
	"MATCH": token.KeywordMatch,
}

var keywordsWithN map[string]token.Type = map[string]token.Type{
	"NATURAL": token.KeywordNatural,
	"NO":      token.KeywordNo,
	"NOT":     token.KeywordNot,
	"NOTHING": token.KeywordNothing,
	"NOTNULL": token.KeywordNotnull,
	"NULL":    token.KeywordNull,
}

var keywordsWithO map[string]token.Type = map[string]token.Type{
	"OF":     token.KeywordOf,
	"OFFSET": token.KeywordOffset,
	"ON":     token.KeywordOn,
	"OR":     token.KeywordOr,
	"ORDER":  token.KeywordOrder,
	"OTHERS": token.KeywordOthers,
	"OUTER":  token.KeywordOuter,
	"OVER":   token.KeywordOver,
}

var keywordsWithP map[string]token.Type = map[string]token.Type{
	"PARTITION": token.KeywordPartition,
	"PLAN":      token.KeywordPlan,
	"PRAGMA":    token.KeywordPragma,
	"PRECEDING": token.KeywordPreceding,
	"PRIMARY":   token.KeywordPrimary,
}

var keywordsWithQ map[string]token.Type = map[string]token.Type{
	"QUERY": token.KeywordQuery,
}

var keywordsWithR map[string]token.Type = map[string]token.Type{
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

var keywordsWithS map[string]token.Type = map[string]token.Type{
	"SAVEPOINT": token.KeywordSavepoint,
	"SELECT":    token.KeywordSelect,
	"SET":       token.KeywordSet,
}

var keywordsWithT map[string]token.Type = map[string]token.Type{
	"TABLE":       token.KeywordTable,
	"TEMP":        token.KeywordTemp,
	"TEMPORARY":   token.KeywordTemporary,
	"THEN":        token.KeywordThen,
	"TIES":        token.KeywordTies,
	"TO":          token.KeywordTo,
	"TRANSACTION": token.KeywordTransaction,
	"TRIGGER":     token.KeywordTrigger,
}

var keywordsWithU map[string]token.Type = map[string]token.Type{
	"UNBOUNDED": token.KeywordUnbounded,
	"UNION":     token.KeywordUnion,
	"UNIQUE":    token.KeywordUnique,
	"UPDATE":    token.KeywordUpdate,
	"USING":     token.KeywordUsing,
}

var keywordsWithV map[string]token.Type = map[string]token.Type{
	"VACUUM":  token.KeywordVacuum,
	"VALUES":  token.KeywordValues,
	"VIEW":    token.KeywordView,
	"VIRTUAL": token.KeywordVirtual,
}

var keywordsWithW map[string]token.Type = map[string]token.Type{
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

// func scanAKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithA[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithA[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanBKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithB[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithB[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanCKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithC[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithC[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanDKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithD[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithD[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanEKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithE[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithE[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanFKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithF[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithF[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanGKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithG[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithG[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanHKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithH[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithH[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanIKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithI[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithI[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanJKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithJ[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithJ[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanKKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithK[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithK[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanLKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithL[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithL[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanMKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithM[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithM[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanNKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithN[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithN[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanOKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithO[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithO[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanPKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithP[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithP[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanQKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithQ[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithQ[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanRKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithR[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithR[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

func scanSKeyword(s *scanner) token.Token {
	nextRune := s.seekNext(s.start)
	input := string(s.input[s.start:nextRune])
	if _, ok := keywordsWithS[input]; ok {
		s.acceptString(input)
		return s.createToken(keywordsWithS[input])
	}
	return s.createToken(keywordsWithS[input])
}

// func scanTKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithT[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithT[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanUKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithU[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithU[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }
// func scanVKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithV[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithV[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }

// func scanWKeyword(s *scanner) token.Token {
// 	nextRune := s.seekNext(s.start)
// 	input := string(s.input[s.start:nextRune])
// 	if _, ok := keywordsWithW[input]; ok {
// 		s.acceptString(input)
// 		return s.createToken(keywordsWithW[input])
// 	}
// 	return s.unexpectedRune(s.input[s.start:nextRune])
// }
