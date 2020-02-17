package parser

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

type errorReporter struct {
	p      *simpleParser
	errs   []error
	sealed bool
}

func (r *errorReporter) errorToken(t token.Token) {
	r.errorf("%w: %s", ErrScanner, t.Value())
}

func (r *errorReporter) incompleteStatement() {
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok {
		r.errorf("%w: EOF", ErrIncompleteStatement)
	} else {
		r.errorf("%w: %s at (%d:%d) offset %d length %d", ErrIncompleteStatement, next.Type().String(), next.Line(), next.Col(), next.Offset(), next.Length())
	}
}

func (r *errorReporter) prematureEOF() {
	r.errorf("%w", ErrPrematureEOF)
	r.sealed = true
}

func (r *errorReporter) unexpectedToken(expected ...token.Type) {
	if r.sealed {
		return
	}
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok || next.Type() == token.EOF {
		// use this instead of r.prematureEOF() because we can add the
		// information about what tokens were expected
		r.errorf("%w: expected %s", ErrPrematureEOF, expected)
		r.sealed = true
		return
	}

	r.errorf("%w: got %s but expected one of %s at (%d:%d) offset %d length %d", ErrUnexpectedToken, next, expected, next.Line(), next.Col(), next.Offset(), next.Length())
}

func (r *errorReporter) unexpectedSingleRuneToken(typ token.Type, expected ...rune) {
	if r.sealed {
		return
	}
	next, ok := r.p.unsafeLowLevelLookahead()
	if !ok || next.Type() == token.EOF {
		// use this instead of r.prematureEOF() because we can add the
		// information about what tokens were expected
		r.errorf("%w: expected %s (more precisely one of %v)", ErrPrematureEOF, typ, expected)
		r.sealed = true
		return
	}

	r.errorf("%w: got %s but expected one of %s at (%d:%d) offset %d length %d", ErrUnexpectedToken, next, typ, next.Line(), next.Col(), next.Offset(), next.Length())
}

func (r *errorReporter) unhandledToken(t token.Token) {
	r.errorf("%w: %s(%s) at (%d:%d) offset %d lenght %d", ErrUnknownToken, t.Type().String(), t.Value(), t.Line(), t.Col(), t.Offset(), t.Length())
}

func (r *errorReporter) unsupportedConstruct(t token.Token) {
	r.errorf("%w: %s(%s) at (%d:%d) offset %d lenght %d", ErrUnsupportedConstruct, t.Type().String(), t.Value(), t.Line(), t.Col(), t.Offset(), t.Length())
}

func (r *errorReporter) errorf(format string, args ...interface{}) {
	r.errs = append(r.errs, fmt.Errorf(format, args...))
}

type reporter interface {
	errorToken(t token.Token)
	incompleteStatement()
	prematureEOF()
	unexpectedToken(expected ...token.Type)
	unexpectedSingleRuneToken(typ token.Type, expected ...rune)
	unhandledToken(t token.Token)
	unsupportedConstruct(t token.Token)
}

var _ Parser = (*simpleParser)(nil) // ensure that simpleParser implements Parser

type simpleParser struct {
	scanner scanner.Scanner
}

// NewSimpleParser creates new ready to use parser.
func NewSimpleParser(input string) Parser {
	return &simpleParser{
		scanner: scanner.NewRuleBased([]rune(input), ruleset.Default),
	}
}

func (p *simpleParser) Next() (*ast.SQLStmt, []error, bool) {
	if p.scanner.Peek().Type() == token.EOF {
		return nil, []error{}, false
	}
	errs := &errorReporter{
		p:    p,
		errs: []error{},
	}
	stmt := p.parseSQLStatement(errs)
	return stmt, errs.errs, true
}

// searchNext skips tokens until a token is of one of the given types. That
// token will not be consumed, every other token will be consumed and an
// unexpected token error will be reported.
func (p *simpleParser) searchNext(r reporter, types ...token.Type) {
	for {
		next, ok := p.unsafeLowLevelLookahead()
		if !ok {
			return
		}
		for _, typ := range types {
			if next.Type() == typ {
				return
			}
		}
		r.unexpectedToken(types...)
		p.consumeToken()
	}
}

func (p *simpleParser) skipUntil(types ...token.Type) {
	for {
		next, ok := p.unsafeLowLevelLookahead()
		if !ok {
			return
		}
		for _, typ := range types {
			if next.Type() == typ {
				return
			}
		}
		p.consumeToken()
	}
}

// unsafeLowLevelLookahead is a low level lookahead, only use if needed.
// Remember to check for token.Error, token.EOF and token.StatementSeparator, as
// this will only return hasNext=false if there are no more tokens (which only
// should occur after an EOF token). Any other token will be returned with
// next=<token>,hasNext=true.
func (p *simpleParser) unsafeLowLevelLookahead() (next token.Token, hasNext bool) {
	return p.scanner.Peek(), true
}

func (p *simpleParser) lookaheadWithType(r reporter, typ token.Type /* ensure at compile time that at least one type is specified */, typs ...token.Type) (token.Token, bool) {
	next, hasNext := p.lookahead(r)
	contains := next.Type() == typ
	for _, t := range typs {
		contains = contains || (next.Type() == t)
	}
	return next, hasNext && contains
}

// lookahead performs a lookahead while consuming any error or statement
// separator token, and reports an EOF, Error or IncompleteStatement if
// appropriate. If this returns ok=false, return from your parse function
// without reporting any more errors. If ok=false, this means that the next
// token was either a StatementSeparator or EOF, and an error has been reported.
func (p *simpleParser) lookahead(r reporter) (next token.Token, ok bool) {
	next, ok = p.optionalLookahead(r)

	if !ok || next.Type() == token.EOF {
		r.prematureEOF()
		ok = false
	} else if next.Type() == token.StatementSeparator {
		r.incompleteStatement()
		ok = false
	}
	return
}

// optionalLookahead performs a lookahead while consuming any error token. If
// this returns ok=false, no more tokens are available.
func (p *simpleParser) optionalLookahead(r reporter) (next token.Token, ok bool) {
	next, ok = p.unsafeLowLevelLookahead()

	// drain all error tokens
	for ok && next.Type() == token.Error {
		r.errorToken(next)
		p.consumeToken()
		next, ok = p.unsafeLowLevelLookahead()
	}

	return next, ok
}

func (p *simpleParser) consumeToken() {
	_ = p.scanner.Next()
}

func (p *simpleParser) parseSQLStatement(r reporter) (stmt *ast.SQLStmt) {
	stmt = &ast.SQLStmt{}

	if next, ok := p.lookaheadWithType(r, token.KeywordExplain); ok {
		stmt.Explain = next
		p.consumeToken()

		if next, ok := p.lookaheadWithType(r, token.KeywordQuery); ok {
			stmt.Query = next
			p.consumeToken()

			if next, ok := p.lookaheadWithType(r, token.KeywordPlan); ok {
				stmt.Plan = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordPlan)
				// At this point, just assume that 'QUERY' was a mistake. Don't
				// abort. It's very unlikely that 'PLAN' occurs somewhere, so
				// assume that the user meant to input 'EXPLAIN <statement>'
				// instead of 'EXPLAIN QUERY PLAN <statement>'.
			}
		}
	}

	// according to the grammar, these are the tokens that initiate a statement
	p.searchNext(r, token.StatementSeparator, token.EOF, token.KeywordAlter, token.KeywordAnalyze, token.KeywordAttach, token.KeywordBegin, token.KeywordCommit, token.KeywordCreate, token.KeywordDelete, token.KeywordDetach, token.KeywordDrop, token.KeywordInsert, token.KeywordPragma, token.KeywordReindex, token.KeywordRelease, token.KeywordRollback, token.KeywordSavepoint, token.KeywordSelect, token.KeywordUpdate, token.KeywordVacuum)

	next, ok := p.unsafeLowLevelLookahead()
	if !ok {
		r.incompleteStatement()
		return
	}

	// lookahead processing to check what the statement ahead is
	switch next.Type() {
	case token.KeywordAlter:
		stmt.AlterTableStmt = p.parseAlterTableStmt(r)
	case token.StatementSeparator:
		r.incompleteStatement()
		p.consumeToken()
	case token.KeywordPragma:
		// we don't support pragmas, as we don't need them yet
		r.unsupportedConstruct(next)
		p.skipUntil(token.StatementSeparator, token.EOF)
	default:
		r.unsupportedConstruct(next)
		p.skipUntil(token.StatementSeparator, token.EOF)
	}

	p.searchNext(r, token.StatementSeparator, token.EOF)
	next, ok = p.unsafeLowLevelLookahead()
	if !ok {
		return
	}
	if next.Type() == token.StatementSeparator {
		// if there's a statement separator, consume this token and get the next
		// token, so that it can be checked if that next token is an EOF
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
	}
	if next.Type() == token.EOF {
		p.consumeToken()
		return
	}
	return
}

func (p *simpleParser) parseAlterTableStmt(r reporter) (stmt *ast.AlterTableStmt) {
	stmt = &ast.AlterTableStmt{}

	p.searchNext(r, token.KeywordAlter)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Alter = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordTable {
		stmt.Table = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordTable)
		// do not consume anything, report that there is no 'TABLE' keyword, and
		// proceed as if we would have received it
	}

	schemaOrTableName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaOrTableName.Type() != token.Literal {
		r.unexpectedToken(token.Literal)
		return
	}
	p.consumeToken() // consume the literal

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "." {
		// schemaOrTableName is schema name
		stmt.SchemaName = schemaOrTableName
		stmt.Period = next
		p.consumeToken()

		tableName, ok := p.lookahead(r)
		if !ok {
			return
		}
		if tableName.Type() == token.Literal {
			stmt.TableName = tableName
			p.consumeToken()
		}
	} else {
		stmt.TableName = schemaOrTableName
	}
	switch next.Type() {
	case token.KeywordRename:
		stmt.Rename = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordTo:
			stmt.To = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}
			stmt.NewTableName = next
			p.consumeToken()
		case token.KeywordColumn:
			stmt.Column = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}

			fallthrough
		case token.Literal:
			stmt.ColumnName = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.KeywordTo {
				r.unexpectedToken(token.KeywordTo)
				p.consumeToken()
				return
			}

			stmt.To = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}

			stmt.NewColumnName = next
			p.consumeToken()
		default:
			r.unexpectedToken(token.KeywordTo, token.KeywordColumn, token.Literal)
		}
	case token.KeywordAdd:
		stmt.Add = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordColumn:
			stmt.Column = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}
			fallthrough
		case token.Literal:
			stmt.ColumnDef = p.parseColumnDef(r)
		default:
			r.unexpectedToken(token.KeywordColumn, token.Literal)
		}
	}

	return
}

func (p *simpleParser) parseColumnDef(r reporter) (def *ast.ColumnDef) {
	def = &ast.ColumnDef{}

	if next, ok := p.lookaheadWithType(r, token.Literal); ok {
		def.ColumnName = next
		p.consumeToken()

		if next, ok = p.lookahead(r); ok && next.Type() == token.Literal {
			def.TypeName = p.parseTypeName(r)
		}

		for {
			next, ok = p.optionalLookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordConstraint ||
				next.Type() == token.KeywordPrimary ||
				next.Type() == token.KeywordNot ||
				next.Type() == token.KeywordUnique ||
				next.Type() == token.KeywordCheck ||
				next.Type() == token.KeywordDefault ||
				next.Type() == token.KeywordCollate ||
				next.Type() == token.KeywordGenerated ||
				next.Type() == token.KeywordReferences {
				def.ColumnConstraint = append(def.ColumnConstraint, p.parseColumnConstraint(r))
			} else {
				break
			}
		}
	}
	return
}

func (p *simpleParser) parseTypeName(r reporter) (name *ast.TypeName) {
	name = &ast.TypeName{}

	// one or more name
	if next, ok := p.lookaheadWithType(r, token.Literal); ok {
		name.Name = append(name.Name, next)
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
	}
	for {
		if next, ok := p.lookaheadWithType(r, token.Literal); ok {
			name.Name = append(name.Name, next)
			p.consumeToken()
		} else {
			break
		}
	}

	if next, ok := p.lookaheadWithType(r, token.Delimiter); ok {
		if next.Value() == "(" {
			name.LeftParen = next
			p.consumeToken()

			name.SignedNumber1 = p.parseSignedNumber(r)
		} else {
			r.unexpectedToken(token.Delimiter)
		}
	} else {
		return
	}

	if next, ok := p.lookaheadWithType(r, token.Delimiter); ok {
		switch next.Value() {
		case ",":
			name.Comma = next
			p.consumeToken()

			name.SignedNumber2 = p.parseSignedNumber(r)
			next, ok = p.lookaheadWithType(r, token.Delimiter)
			if !ok {
				return
			}
			fallthrough
		case ")":
			name.RightParen = next
			p.consumeToken()
		}
	} else {
		return
	}

	return
}

func (p *simpleParser) parseSignedNumber(r reporter) (num *ast.SignedNumber) {
	num = &ast.SignedNumber{}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.UnaryOperator:
		num.Sign = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() != token.Literal {
			r.unexpectedToken(token.Literal)
			return
		}
		fallthrough
	case token.Literal:
		num.NumericLiteral = next
		p.consumeToken()
	default:
		r.unexpectedToken(token.UnaryOperator, token.Literal)
		return
	}
	return
}

func (p *simpleParser) parseColumnConstraint(r reporter) (constr *ast.ColumnConstraint) {
	constr = &ast.ColumnConstraint{}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordConstraint {
		constr.Constraint = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			constr.Name = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
			// report that the token was unexpected, but continue as if the
			// missing literal token was present
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordPrimary:
		// PRIMARY
		constr.Primary = next
		p.consumeToken()

		// KEY
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordKey {
			constr.Key = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordKey)
		}

		// ASC, DESC
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAsc {
			constr.Asc = next
			p.consumeToken()
		} else if next.Type() == token.KeywordDesc {
			constr.Desc = next
			p.consumeToken()
		}

		// conflict clause
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			constr.ConflictClause = p.parseConflictClause(r)
		}

		// AUTOINCREMENT
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAutoincrement {
			constr.Autoincrement = next
			p.consumeToken()
		}

	case token.KeywordNot:
		// NOT
		constr.Not = next
		p.consumeToken()

		// NULL
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordNull {
			constr.Null = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordNull)
		}

		// conflict clause
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			constr.ConflictClause = p.parseConflictClause(r)
		}

	case token.KeywordUnique:
		// UNIQUE
		constr.Unique = next
		p.consumeToken()

		// conflict clause
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			constr.ConflictClause = p.parseConflictClause(r)
		}

	case token.KeywordCheck:
		// CHECK
		constr.Check = next
		p.consumeToken()

		// left paren
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			constr.LeftParen = next
			p.consumeToken()
		} else {
			r.unexpectedSingleRuneToken(token.Delimiter, '(')
			// assume that the opening paren has been omitted, report the
			// error but proceed as if it was found
		}

		// expr
		constr.Expr = p.parseExpression(r)

		// right paren
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == ")" {
			constr.RightParen = next
			p.consumeToken()
		} else {
			r.unexpectedSingleRuneToken(token.Delimiter, ')')
			// assume that the opening paren has been omitted, report the
			// error but proceed as if it was found
		}

	case token.KeywordDefault:
		constr.Default = next
		p.consumeToken()

	case token.KeywordCollate:
		constr.Collate = next
		p.consumeToken()

	case token.KeywordGenerated:
		constr.Generated = next
		p.consumeToken()

	case token.KeywordReferences:
		constr.ForeignKeyClause = p.parseForeignKeyClause(r)
	default:
		r.unexpectedToken(token.KeywordPrimary, token.KeywordNot, token.KeywordUnique, token.KeywordCheck, token.KeywordDefault, token.KeywordCollate, token.KeywordGenerated, token.KeywordReferences)
	}

	return
}

func (p *simpleParser) parseForeignKeyClause(r reporter) (clause *ast.ForeignKeyClause) {
	panic("implement me")
}

func (p *simpleParser) parseConflictClause(r reporter) (clause *ast.ConflictClause) {
	clause = &ast.ConflictClause{}

	next, ok := p.optionalLookahead(r)
	if !ok {
		return
	}

	// ON
	if next.Type() == token.KeywordOn {
		clause.On = next
		p.consumeToken()
	} else {
		// if there's no 'ON' token, the empty production is assumed, which is
		// why no error is reported here
		return
	}

	// CONFLICT
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordConflict {
		clause.Conflict = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordConflict)
		return
	}

	// ROLLBACK, ABORT, FAIL, IGNORE, REPLACE
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordRollback:
		clause.Rollback = next
		p.consumeToken()
	case token.KeywordAbort:
		clause.Abort = next
		p.consumeToken()
	case token.KeywordFail:
		clause.Fail = next
		p.consumeToken()
	case token.KeywordIgnore:
		clause.Ignore = next
		p.consumeToken()
	case token.KeywordReplace:
		clause.Replace = next
		p.consumeToken()
	default:
		r.unexpectedToken(token.KeywordRollback, token.KeywordAbort, token.KeywordFail, token.KeywordIgnore, token.KeywordReplace)
	}
	return
}

func (p *simpleParser) parseExpression(r reporter) (expr *ast.Expr) {
	panic("implement me")
}
