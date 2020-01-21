package scanner

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var keywordsWithA map[string]bool = map[string]bool{
	"ABORT":          true,
	"ACTION":         true,
	"ADD":            true,
	"AFTER":          true,
	"ALL":            true,
	"ALTER":          true,
	"ANALYZE":        true,
	"AND":            true,
	"AS":             true,
	"ASC":            true,
	"ATTACH":         true,
	"AUTO INCREMENT": true,
}

var keywordsWithB map[string]bool = map[string]bool{
	"BEFORE":  true,
	"BEGIN":   true,
	"BETWEEN": true,
	"BY":      true,
}

var keywordsWithC map[string]bool = map[string]bool{
	"CASCADE":           true,
	"CASE":              true,
	"CAST":              true,
	"CHECK":             true,
	"COLLATE":           true,
	"COLUMN":            true,
	"COMMIT":            true,
	"CONFLICT":          true,
	"CONSTRAINT":        true,
	"CREATE":            true,
	"CROSS":             true,
	"CURRENT":           true,
	"CURRENT_DATE":      true,
	"CURRENT_TIME":      true,
	"CURRENT_TIMESTAMP": true,
}

var keywordsWithD map[string]bool = map[string]bool{
	"DATABASE":   true,
	"DEFAULT":    true,
	"DEFERRABLE": true,
	"DEFERRED":   true,
	"DELETE":     true,
	"DESC":       true,
	"DETACH":     true,
	"DISTINCT":   true,
	"DO":         true,
	"DROP":       true,
}

var keywordsWithE map[string]bool = map[string]bool{
	"EACH":      true,
	"ELSE":      true,
	"END":       true,
	"ESCAPE":    true,
	"EXCEPT":    true,
	"EXCLUDE":   true,
	"EXCLUSIVE": true,
	"EXISTS":    true,
	"EXPLAIN":   true,
}

var keywordsWithF map[string]bool = map[string]bool{
	"FAIL":      true,
	"FILTER":    true,
	"FIRST":     true,
	"FOLLOWING": true,
	"FOR":       true,
	"FOREIGN":   true,
	"FROM":      true,
	"FULL":      true,
}

var keywordsWithG map[string]bool = map[string]bool{
	"GLOB":   true,
	"GROUP":  true,
	"GROUPS": true,
}

var keywordsWithH map[string]bool = map[string]bool{
	"HAVING": true,
}

var keywordsWithI map[string]bool = map[string]bool{
	"IF":        true,
	"IGNORE":    true,
	"IMMEDIATE": true,
	"IN":        true,
	"INDEX":     true,
	"INDEXED":   true,
	"INITIALLY": true,
	"INNER":     true,
	"INSERT":    true,
	"INSTEAD":   true,
	"INTERSECT": true,
	"INTO":      true,
	"IS":        true,
	"ISNULL":    true,
}

var keywordsWithJ map[string]bool = map[string]bool{
	"JOIN": true,
}

var keywordsWithK map[string]bool = map[string]bool{
	"KEY": true,
}

var keywordsWithL map[string]bool = map[string]bool{
	"LAST":  true,
	"LEFT":  true,
	"LIKE":  true,
	"LIMIT": true,
}

var keywordsWithM map[string]bool = map[string]bool{
	"MATCH": true,
}

var keywordsWithN map[string]bool = map[string]bool{
	"NATURAL": true,
	"NO":      true,
	"NOT":     true,
	"NOTHING": true,
	"NOTNULL": true,
	"NULL":    true,
}

var keywordsWithO map[string]bool = map[string]bool{
	"OF":     true,
	"OFFSET": true,
	"ON":     true,
	"OR":     true,
	"ORDER":  true,
	"OTHERS": true,
	"OUTER":  true,
	"OVER":   true,
}

var keywordsWithP map[string]bool = map[string]bool{
	"PARTITION": true,
	"PLAN":      true,
	"PRAGMA":    true,
	"PRECEDING": true,
	"PRIMARY":   true,
}

var keywordsWithQ map[string]bool = map[string]bool{
	"QUERY": true,
}

var keywordsWithR map[string]bool = map[string]bool{
	"RAISE":      true,
	"RANGE":      true,
	"RECURSIVE":  true,
	"REFERENCES": true,
	"REGEXP":     true,
	"REINDEX":    true,
	"RELEASE":    true,
	"RENAME":     true,
	"REPLACE":    true,
	"RESTRICT":   true,
	"RIGHT":      true,
	"ROLLBACK":   true,
	"ROW":        true,
	"ROWS":       true,
}

var keywordsWithS map[string]bool = map[string]bool{
	"SAVEPOINT": true,
	"SELECT":    true,
	"SET":       true,
}

var keywordsWithT map[string]bool = map[string]bool{
	"TABLE":       true,
	"TEMP":        true,
	"TEMPORARY":   true,
	"THEN":        true,
	"TIES":        true,
	"TO":          true,
	"TRANSACTION": true,
	"TRIGGER":     true,
}

var keywordsWithU map[string]bool = map[string]bool{
	"UNBOUNDED": true,
	"UNION":     true,
	"UNIQUE":    true,
	"UPDATE":    true,
	"USING":     true,
}

var keywordsWithV map[string]bool = map[string]bool{
	"VACUUM":  true,
	"VALUES":  true,
	"VIEW":    true,
	"VIRTUAL": true,
}

var keywordsWithw map[string]bool = map[string]bool{
	"WHEN":    true,
	"WHERE":   true,
	"WINDOW":  true,
	"WITH":    true,
	"WITHOUT": true,
}

// func scanSpace(s *scanner) token.Token {
// 	s.accept(matcher.String(" "))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
// }

// func scanDoubleQuote(s *scanner) token.Token {
// 	s.accept(matcher.String("\""))
// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
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

func scanSelectOperator(s *scanner) token.Token {
	fmt.Println(string(s.input[s.start:s.seekNext(s.start)]))
	// if s.acceptString("SELECT") {
	// 	return createToken(s.line, s.col, s.start, s.pos, token.SQLSpecialCharacter, string(s.input[s.start:s.pos]), s)
	// }

	return createToken(s.line, s.col, s.start, s.pos, token.KeywordSelect, string(s.input[s.start:s.pos]), s)
}
