package parser

import (
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func (p *simpleParser) parseSQLStatement(r reporter) (stmt *ast.SQLStmt) {
	stmt = &ast.SQLStmt{}

	if next, ok := p.lookahead(r); ok && next.Type() == token.KeywordExplain {
		stmt.Explain = next
		p.consumeToken()

		if next, ok := p.lookahead(r); ok && next.Type() == token.KeywordQuery {
			stmt.Query = next
			p.consumeToken()

			if next, ok := p.lookahead(r); ok && next.Type() == token.KeywordPlan {
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
	p.searchNext(r, token.StatementSeparator, token.EOF, token.KeywordAlter, token.KeywordAnalyze, token.KeywordAttach, token.KeywordBegin, token.KeywordCommit, token.KeywordCreate, token.KeywordDelete, token.KeywordDetach, token.KeywordDrop, token.KeywordEnd, token.KeywordInsert, token.KeywordPragma, token.KeywordReIndex, token.KeywordRelease, token.KeywordReplace, token.KeywordRollback, token.KeywordSavepoint, token.KeywordSelect, token.KeywordUpdate, token.KeywordVacuum, token.KeywordValues, token.KeywordWith)

	next, ok := p.unsafeLowLevelLookahead()
	if !ok {
		r.incompleteStatement()
		return
	}

	// lookahead processing to check what the statement ahead is
	switch next.Type() {
	case token.KeywordAlter:
		stmt.AlterTableStmt = p.parseAlterTableStmt(r)
	case token.KeywordAnalyze:
		stmt.AnalyzeStmt = p.parseAnalyzeStmt(r)
	case token.KeywordAttach:
		stmt.AttachStmt = p.parseAttachDatabaseStmt(r)
	case token.KeywordBegin:
		stmt.BeginStmt = p.parseBeginStmt(r)
	case token.KeywordCommit:
		stmt.CommitStmt = p.parseCommitStmt(r)
	case token.KeywordCreate:
		p.parseCreateStmts(stmt, r)
	case token.KeywordDelete:
		p.parseDeleteStmts(stmt, nil, r)
	case token.KeywordDetach:
		stmt.DetachStmt = p.parseDetachDatabaseStmt(r)
	case token.KeywordDrop:
		p.parseDropStmts(stmt, r)
	case token.KeywordEnd:
		stmt.CommitStmt = p.parseCommitStmt(r)
	case token.KeywordInsert:
		stmt.InsertStmt = p.parseInsertStmt(nil, r)
	case token.KeywordReIndex:
		stmt.ReIndexStmt = p.parseReIndexStmt(r)
	case token.KeywordRelease:
		stmt.ReleaseStmt = p.parseReleaseStmt(r)
	case token.KeywordReplace:
		stmt.InsertStmt = p.parseInsertStmt(nil, r)
	case token.KeywordRollback:
		stmt.RollbackStmt = p.parseRollbackStmt(r)
	case token.KeywordSavepoint:
		stmt.SavepointStmt = p.parseSavepointStmt(r)
	case token.KeywordSelect:
		stmt.SelectStmt = p.parseSelectStmt(nil, r)
	case token.KeywordUpdate:
		p.parseUpdateStmts(stmt, nil, r)
	case token.KeywordVacuum:
		stmt.VacuumStmt = p.parseVacuumStmt(r)
	case token.KeywordValues:
		stmt.SelectStmt = p.parseSelectStmt(nil, r)
	case token.KeywordWith:
		p.parseWithClauseBeginnerStmts(stmt, r)
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
	if next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		p.consumeToken()
		return
	}
	return
}

// parseAlterTableStmt parses the alter-table stmt as defined in:
// https://sqlite.org/lang_altertable.html
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

// parseColumnDef parses the column-def stmt as defined in:
// https://sqlite.org/syntax/column-def.html
func (p *simpleParser) parseColumnDef(r reporter) (def *ast.ColumnDef) {
	def = &ast.ColumnDef{}

	if next, ok := p.lookahead(r); ok && next.Type() == token.Literal {
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
				next.Type() == token.KeywordAs ||
				next.Type() == token.KeywordReferences {
				def.ColumnConstraint = append(def.ColumnConstraint, p.parseColumnConstraint(r))
			} else {
				break
			}
		}
	}
	return
}

// parseTypeName parses the type-name stmt as defined in:
// https://sqlite.org/syntax/type-name.html
func (p *simpleParser) parseTypeName(r reporter) (name *ast.TypeName) {
	name = &ast.TypeName{}

	// one or more name
	if next, ok := p.lookahead(r); ok && next.Type() == token.Literal {
		name.Name = append(name.Name, next)
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
	}
	for {
		if next, ok := p.lookahead(r); ok && next.Type() == token.Literal {
			name.Name = append(name.Name, next)
			p.consumeToken()
		} else {
			break
		}
	}

	if next, ok := p.lookahead(r); ok && next.Type() == token.Delimiter && next.Value() == "(" {
		name.LeftParen = next
		p.consumeToken()

		name.SignedNumber1 = p.parseSignedNumber(r)

	} else {
		return
	}

	if next, ok := p.lookahead(r); ok && next.Type() == token.Delimiter {
		switch next.Value() {
		case ",":
			name.Comma = next
			p.consumeToken()

			name.SignedNumber2 = p.parseSignedNumber(r)
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Delimiter {
				r.unexpectedToken(token.Delimiter)
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

// parseSignedNumber parses the signed-number stmt as defined in:
// https://sqlite.org/syntax/signed-number.html
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
		if next.Type() != token.LiteralNumeric {
			r.unexpectedToken(token.LiteralNumeric)
			return
		}
		fallthrough
	case token.LiteralNumeric:
		num.NumericLiteral = next
		p.consumeToken()
	default:
		r.unexpectedToken(token.UnaryOperator, token.LiteralNumeric)
		return
	}
	return
}

// parseColumnConstraint parses the column-constraint stmt as defined in:
// https://sqlite.org/syntax/column-constraint.html
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
		if constr.Expr == nil {
			r.expectedExpression()
		}

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
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.UnaryOperator || next.Type() == token.LiteralNumeric {
			constr.SignedNumber = p.parseSignedNumber(r)
		}
		// Only either of the 3 cases can exist, thus, if one of the cases is found, return.
		if constr.SignedNumber != nil {
			return
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			constr.LiteralValue = next
			p.consumeToken()
		}
		if constr.LiteralValue != nil {
			return
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			constr.LeftParen = next
			p.consumeToken()
			constr.Expr = p.parseExpression(r)
			if constr.Expr == nil {
				r.expectedExpression()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Delimiter && next.Value() == ")" {
				constr.RightParen = next
				p.consumeToken()
			} else {
				r.unexpectedSingleRuneToken(')')
			}
		}
	case token.KeywordCollate:
		constr.Collate = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			constr.CollationName = next
			p.consumeToken()
		}
	case token.KeywordReferences:
		constr.ForeignKeyClause = p.parseForeignKeyClause(r)
	case token.KeywordGenerated:
		constr.Generated = next
		p.consumeToken()

		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAlways {
			constr.Always = next
			p.consumeToken()
		}
		fallthrough
	case token.KeywordAs:
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAs {
			constr.As = next
			p.consumeToken()
		}
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			constr.LeftParen = next
			p.consumeToken()
			constr.Expr = p.parseExpression(r)
			if constr.Expr == nil {
				r.expectedExpression()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Delimiter && next.Value() == ")" {
				constr.RightParen = next
				p.consumeToken()
			} else {
				r.unexpectedSingleRuneToken(')')
			}
		}
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordStored {
			constr.Stored = next
			p.consumeToken()
		}
		if next.Type() == token.KeywordVirtual {
			constr.Virtual = next
			p.consumeToken()
		}
	default:
		r.unexpectedToken(token.KeywordPrimary, token.KeywordNot, token.KeywordUnique, token.KeywordCheck, token.KeywordDefault, token.KeywordCollate, token.KeywordGenerated, token.KeywordAs, token.KeywordReferences)
	}
	return
}

// parseForeignKeyClause parses the foreign-key-clause stmt as defined in:
// https://sqlite.org/syntax/foreign-key-clause.html
func (p *simpleParser) parseForeignKeyClause(r reporter) (clause *ast.ForeignKeyClause) {
	clause = &ast.ForeignKeyClause{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordReferences {
		clause.References = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			clause.ForeignTable = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			clause.LeftParen = next
			p.consumeToken()
			for {
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.Literal {
					clause.ColumnName = append(clause.ColumnName, next)
					p.consumeToken()
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Value() == "," {
						p.consumeToken()
					}
					if next.Value() == ")" {
						clause.RightParen = next
						p.consumeToken()
						break
					}
				} else {
					break
				}
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		for {
			if next.Type() == token.KeywordOn || next.Type() == token.KeywordMatch {
				clause.ForeignKeyClauseCore = append(clause.ForeignKeyClauseCore, p.parseForeignKeyClauseCore(r))
			} else {
				break
			}
			next, ok = p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				return
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordNot {
			clause.Not = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordDeferrable {
			clause.Deferrable = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordInitially {
			clause.Initially = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordImmediate {
				clause.Immediate = next
				p.consumeToken()
			} else if next.Type() == token.KeywordDeferred {
				clause.Deferred = next
				p.consumeToken()
			}
		}
	} else {
		return
	}
	return
}

func (p *simpleParser) parseForeignKeyClauseCore(r reporter) (stmt *ast.ForeignKeyClauseCore) {
	stmt = &ast.ForeignKeyClauseCore{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOn {
		stmt.On = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordDelete {
			stmt.Delete = next
			p.consumeToken()
		} else if next.Type() == token.KeywordUpdate {
			stmt.Update = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordSet:
			stmt.Set = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordNull {
				stmt.Null = next
				p.consumeToken()
			} else if next.Type() == token.KeywordDefault {
				stmt.Default = next
				p.consumeToken()
			}
		case token.KeywordCascade:
			stmt.Cascade = next
			p.consumeToken()
		case token.KeywordRestrict:
			stmt.Restrict = next
			p.consumeToken()
		case token.KeywordNo:
			stmt.No = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordAction {
				stmt.Action = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordAction)
			}
		}

	} else if next.Type() == token.KeywordMatch {
		stmt.Match = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.Name = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}
	return

}

// parseConflictClause parses conflict-clause stmt as defined in:
// https://sqlite.org/syntax/conflict-clause.html
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
		// why no error is reported here.
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

// parseExpression parses expr as defined in:
// https://sqlite.org/syntax/expr.html
// parseExprX or parseExprXSubY are the helper functions that parse line X and sub line Y in the spec.
// (bind-parameter is removed and neglected while counting line numbers)
// parseExprXHelper functions are helper functions for parseExprX, mainly to
// avoid code duplication and suffice alternate paths possible.
func (p *simpleParser) parseExpression(r reporter) (expr *ast.Expr) {
	expr = &ast.Expr{}
	// The following rules being Left Recursive, have been converted to remove it.
	// Details of the conversion follow above the implementations.
	// S - is the starting production rule for expr.
	// S -> SX | Y is converted to S -> YS' and S' -> XS'.

	// S -> (literal) S' and S -> (schema.table.column) S'.
	literal, ok := p.lookahead(r)
	if !ok {
		return
	}
	if literal.Type() == token.Literal || literal.Type() == token.LiteralNumeric {
		expr.LiteralValue = literal
		p.consumeToken()
		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Value() == "." {
			return p.parseExpr2(literal, nil, nil, r)
		} else if next.Type() == token.Delimiter && next.Value() == "(" {
			return p.parseExpr5(literal, r)
		} else {
			returnExpr := p.parseExprRecursive(&ast.Expr{LiteralValue: literal}, r)
			if returnExpr != nil {
				expr = returnExpr
			}
			return
		}
	}

	// S -> (unary op) S'.
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.UnaryOperator {
		expr.UnaryOperator = next
		p.consumeToken()
		expr.Expr1 = p.parseExpression(r)
		if expr.Expr1 == nil {
			r.expectedExpression()
		}

		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		returnExpr := p.parseExprRecursive(expr, r)
		if returnExpr != nil {
			expr = returnExpr
		}
		return
	}

	// S -> (parenth. expr) S'.
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		expr.LeftParen = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if !(next.Type() == token.KeywordWith || next.Type() == token.KeywordSelect || next.Type() == token.KeywordValues) {
			for {
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Value() == "," {
					p.consumeToken()
				}
				if next.Type() == token.Delimiter && next.Value() == ")" {
					expr.RightParen = next
					p.consumeToken()

					next, ok := p.optionalLookahead(r)
					if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
						return
					}

					returnExpr := p.parseExprRecursive(expr, r)
					if returnExpr != nil {
						expr = returnExpr
					}
					return
				}
				expression := p.parseExpression(r)
				if expression != nil {
					expr.Expr = append(expr.Expr, expression)
				} else {
					break
				}
			}
		} else {
			expr.SelectStmt = p.parseSelectStmt(nil, r)
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Delimiter {
				expr.RightParen = next
				p.consumeToken()
				return
			}
		}
	}

	// S -> (CAST) S'.
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordCast {
		expr.Cast = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			expr.LeftParen = next
			p.consumeToken()

			expr.Expr1 = p.parseExpression(r)
			if expr.Expr1 == nil {
				r.expectedExpression()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordAs {
				expr.As = next
				p.consumeToken()

				expr.TypeName = p.parseTypeName(r)

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.Delimiter && next.Value() == ")" {
					expr.RightParen = next
					p.consumeToken()

					next, ok := p.optionalLookahead(r)
					if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
						return
					}
					returnExpr := p.parseExprRecursive(expr, r)
					if returnExpr != nil {
						expr = returnExpr
					}
				}
			}
		} else {
			r.unexpectedToken(token.Delimiter)
		}
		return
	}

	// S -> (NOT EXISTS) S'.
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordNot {
		expr.Not = next
		p.consumeToken()
	}
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordExists {
		expr.Exists = next
		p.consumeToken()
	}
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		expr.LeftParen = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordSelect || next.Type() == token.KeywordWith || next.Type() == token.KeywordValues {
			expr.SelectStmt = p.parseSelectStmt(nil, r)

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Delimiter && next.Value() == ")" {
				expr.RightParen = next
				p.consumeToken()

				next, ok := p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return
				}
				returnExpr := p.parseExprRecursive(expr, r)
				if returnExpr != nil {
					expr = returnExpr
				}
				return
			}
		} else {
			r.unexpectedToken(token.KeywordSelect, token.KeywordValues, token.KeywordWith)
		}
	}

	// S -> (CASE) S'.
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordCase {
		expr.Case = next
		p.consumeToken()
		// This is an optional expression.
		expr.Expr1 = p.parseExpression(r)

		for {
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordEnd || next.Type() == token.KeywordElse {
				break
			}
			expr.WhenThenClause = append(expr.WhenThenClause, p.parseWhenThenClause(r))
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordElse {
			expr.Else = next
			p.consumeToken()
			expr.Expr2 = p.parseExpression(r)
			if expr.Expr2 == nil {
				r.expectedExpression()
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordEnd {
			expr.End = next
			p.consumeToken()

			next, ok := p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				return
			}
			returnExpr := p.parseExprRecursive(expr, r)
			if returnExpr != nil {
				expr = returnExpr
			}
		} else {
			r.unexpectedToken(token.KeywordEnd)
		}
		return
	}

	// S -> (raise-function) S'.
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordRaise {
		raiseFunction := p.parseRaiseFunction(r)
		expr.RaiseFunction = raiseFunction

		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		returnExpr := p.parseExprRecursive(expr, r)
		if returnExpr != nil {
			expr = returnExpr
		}
		return
	}

	// expr can safely assigned to nil as every possible type of expr returns
	// in all the above cases from their respective functions.
	expr = nil
	return
}

// parseExprRecursive will get the smaller expr and will be asked to return a bigger expr,
// IF it exists.
func (p *simpleParser) parseExprRecursive(expr *ast.Expr, r reporter) *ast.Expr {
	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return nil
	}
	switch next.Type() {
	case token.BinaryOperator, token.UnaryOperator:
		return p.parseExpr4(expr, r)
	case token.KeywordCollate:
		return p.parseExpr8(expr, r)
	case token.KeywordNot:
		tokenNot := next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return nil
		}
		switch next.Type() {
		case token.KeywordLike:
			return p.parseExpr9(expr, tokenNot, r)
		case token.KeywordGlob:
			return p.parseExpr9(expr, tokenNot, r)
		case token.KeywordRegexp:
			return p.parseExpr9(expr, tokenNot, r)
		case token.KeywordMatch:
			return p.parseExpr9(expr, tokenNot, r)
		case token.KeywordNull:
			return p.parseExpr10(expr, tokenNot, r)
		case token.KeywordBetween:
			return p.parseExpr12(expr, tokenNot, r)
		case token.KeywordIn:
			return p.parseExpr13(expr, tokenNot, r)
		}
	case token.KeywordLike:
		return p.parseExpr9(expr, nil, r)
	case token.KeywordGlob:
		return p.parseExpr9(expr, nil, r)
	case token.KeywordRegexp:
		return p.parseExpr9(expr, nil, r)
	case token.KeywordMatch:
		return p.parseExpr9(expr, nil, r)
	case token.KeywordIsnull:
		return p.parseExpr10(expr, nil, r)
	case token.KeywordNotnull:
		return p.parseExpr10(expr, nil, r)
	case token.KeywordIs:
		return p.parseExpr11(expr, r)
	case token.KeywordBetween:
		return p.parseExpr12(expr, nil, r)
	case token.KeywordIn:
		return p.parseExpr13(expr, nil, r)
	}
	return nil
}

// parseExprBeginWithLiteral parses possible expressions that begin with a literal.
// A nil is returned if it turns out not to be an expression.
func (p *simpleParser) parseExprBeginWithLiteral(literal token.Token, r reporter) (expr *ast.Expr) {
	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return nil
	}
	if next.Value() == "." {
		return p.parseExpr2(literal, nil, nil, r)
	} else if next.Type() == token.Delimiter && next.Value() == "(" {
		return p.parseExpr5(literal, r)
	} else {
		returnExpr := p.parseExprRecursive(&ast.Expr{LiteralValue: literal}, r)
		return returnExpr
	}
}

// parseExpr2 parses S' -> (schema.table.column clause) S'.
func (p *simpleParser) parseExpr2(schemaOrTableName, period, tableOrColName token.Token, r reporter) (expr *ast.Expr) {
	expr = &ast.Expr{}
	if period == nil && tableOrColName == nil {
		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Value() == "." {
			period := next
			p.consumeToken()
			tableOrColumnName, ok := p.lookahead(r)
			if !ok {
				return
			}
			if tableOrColumnName.Type() != token.Literal {
				return
			}
			p.consumeToken()
			expr = p.parseExpr2Helper(schemaOrTableName, period, tableOrColumnName, r)
			returnExpr := p.parseExprRecursive(expr, r)
			if returnExpr != nil {
				expr = returnExpr
			}
			return
		}
		return
	}
	expr = p.parseExpr2Helper(schemaOrTableName, period, tableOrColName, r)
	returnExpr := p.parseExprRecursive(expr, r)
	if returnExpr != nil {
		expr = returnExpr
	}
	return
}

// parseExpr2Helper parses the line 2 of expr based on alternate paths encountered while parsing.
func (p *simpleParser) parseExpr2Helper(schemaOrTableName, period, tableOrColName token.Token, r reporter) (expr *ast.Expr) {
	expr = &ast.Expr{}
	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		expr.TableName = schemaOrTableName
		expr.Period1 = period
		expr.ColumnName = tableOrColName
		return
	}
	if next.Value() == "." {
		p.consumeToken()
		expr.Period2 = next
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			p.consumeToken()
			expr.ColumnName = next
			expr.TableName = tableOrColName
			expr.SchemaName = schemaOrTableName
			expr.Period1 = period
		} else {
			r.unexpectedToken(token.Literal)
		}
	} else {
		expr.TableName = schemaOrTableName
		expr.Period1 = period
		expr.ColumnName = tableOrColName
	}
	return
}

// parseExpr4 parses S' -> (binary-op) S'.
func (p *simpleParser) parseExpr4(expr *ast.Expr, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr
	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	if next.Type() == token.BinaryOperator || next.Value() == "+" || next.Value() == "-" {
		exprParent.BinaryOperator = next
		p.consumeToken()
		exprParent.Expr2 = p.parseExpression(r)
		if exprParent.Expr2 == nil {
			r.expectedExpression()
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr5 parses S' -> (function-name) S'.
func (p *simpleParser) parseExpr5(functionName token.Token, r reporter) (expr *ast.Expr) {
	expr = &ast.Expr{}
	expr.FunctionName = functionName
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		expr.LeftParen = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordDistinct:
			expr.Distinct = next
			p.consumeToken()
			expr.Expr = p.parseExprSequence(r)
		case token.BinaryOperator:
			expr.Asterisk = next
			p.consumeToken()
		case token.Delimiter:
			if next.Value() == ")" {
				expr.RightParen = next
				p.consumeToken()
			} else {
				r.unexpectedSingleRuneToken(')')
			}
		default:
			expr.Expr = p.parseExprSequence(r)
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		// 3 possiblities, Filter OR over clause OR a ')'
		switch next.Type() {
		case token.KeywordFilter:
			expr.FilterClause = p.parseFilterClause(r)
			next, ok = p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				return
			}
			if next.Type() == token.KeywordOver {
				expr.OverClause = p.parseOverClause(r)
			}
		case token.KeywordOver:
			expr.OverClause = p.parseOverClause(r)
		default:
			// Check whether it was already recorded before.
			if expr.RightParen == nil {
				if next.Type() == token.Delimiter && next.Value() == ")" {
					expr.RightParen = next
					p.consumeToken()
				} else {
					r.unexpectedSingleRuneToken(')')
				}
			}
		}

		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		returnExpr := p.parseExprRecursive(expr, r)
		if returnExpr != nil {
			expr = returnExpr
		}
	}
	return
}

// parseExprSequence parses a sequence of exprs separated by ",".
func (p *simpleParser) parseExprSequence(r reporter) (exprs []*ast.Expr) {
	exprs = []*ast.Expr{}
	for {
		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == ")" {
			return
		}
		if next.Value() == "," {
			p.consumeToken()
		}
		exprVal := p.parseExpression(r)
		if exprVal == nil {
			break
		} else {
			exprs = append(exprs, exprVal)
		}
	}
	return
}

// parseExpr8 parses S' -> (COLLATE) S' | epsilon.
func (p *simpleParser) parseExpr8(expr *ast.Expr, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr
	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	if next.Type() == token.KeywordCollate {
		exprParent.Collate = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return nil
		}
		if next.Type() == token.Literal {
			exprParent.CollationName = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr9 parses S' -> (NOT-LIKE,GLOB,REGEXP,MATCH) S' | epsilon.
func (p *simpleParser) parseExpr9(expr *ast.Expr, tokenNot token.Token, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr
	exprParent.Not = tokenNot

	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	switch next.Type() {
	case token.KeywordLike:
		exprParent.Like = next
		p.consumeToken()
	case token.KeywordGlob:
		exprParent.Glob = next
		p.consumeToken()
	case token.KeywordRegexp:
		exprParent.Regexp = next
		p.consumeToken()
	case token.KeywordMatch:
		exprParent.Match = next
		p.consumeToken()
	}

	exprParent.Expr2 = p.parseExpression(r)
	if exprParent.Expr2 == nil {
		r.expectedExpression()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	if next.Type() == token.KeywordEscape {
		exprParent.Escape = next
		p.consumeToken()
		exprParent.Expr3 = p.parseExpression(r)
		if exprParent.Expr3 == nil {
			r.expectedExpression()
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr10 parses S' -> (ISNULL,NOTNULL,NOT NULL) S' | epsilon.
func (p *simpleParser) parseExpr10(expr *ast.Expr, tokenNot token.Token, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr
	exprParent.Not = tokenNot

	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	switch next.Type() {
	case token.KeywordIsnull:
		exprParent.Isnull = next
		p.consumeToken()
	case token.KeywordNotnull:
		exprParent.Notnull = next
		p.consumeToken()
	case token.KeywordNull:
		exprParent.Null = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr11 parses S' -> (IS NOT) S' | epsilon.
func (p *simpleParser) parseExpr11(expr *ast.Expr, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr

	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	if next.Type() == token.KeywordIs {
		exprParent.Is = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return nil
		}
		if next.Type() == token.KeywordNot {
			exprParent.Not = next
			p.consumeToken()
		}

		exprParent.Expr2 = p.parseExpression(r)
		if exprParent.Expr2 == nil {
			r.expectedExpression()
		}
	} else {
		r.unexpectedToken(token.KeywordIs)
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr12 parses S' -> (NOT BETWEEN) S' | epsilon.
func (p *simpleParser) parseExpr12(expr *ast.Expr, tokenNot token.Token, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr
	exprParent.Not = tokenNot

	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	if next.Type() == token.KeywordBetween {
		exprParent.Between = next
		p.consumeToken()

		exprParent.Expr2 = p.parseExpression(r)
		if exprParent.Expr2 == nil {
			r.expectedExpression()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return nil
		}
		if next.Type() == token.KeywordAnd {
			exprParent.And = next
			p.consumeToken()

			exprParent.Expr3 = p.parseExpression(r)
			if exprParent.Expr3 == nil {
				r.expectedExpression()
			}
		} else {
			r.unexpectedToken(token.KeywordAnd)
		}
	} else {
		r.unexpectedToken(token.KeywordBetween)
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr13 parses S' -> (NOT IN) S' | epsilon.
func (p *simpleParser) parseExpr13(expr *ast.Expr, tokenNot token.Token, r reporter) *ast.Expr {
	exprParent := &ast.Expr{}
	exprParent.Expr1 = expr

	exprParent.Not = tokenNot

	next, ok := p.lookahead(r)
	if !ok {
		return nil
	}
	if next.Type() == token.KeywordIn {
		exprParent.In = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return nil
		}
		switch next.Type() {
		case token.Delimiter:
			if next.Value() == "(" {
				exprParent.LeftParen = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return nil
				}
				if next.Type() == token.KeywordWith || next.Type() == token.KeywordSelect || next.Type() == token.KeywordValues {
					exprParent.SelectStmt = p.parseSelectStmt(nil, r)
				} else if next.Type() == token.Delimiter && next.Value() == ")" {
					exprParent.RightParen = next
					p.consumeToken()
				} else {
					exprParent.Expr = p.parseExprSequence(r)
				}

				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return exprParent
				}
				if next.Type() == token.Delimiter && next.Value() == ")" {
					exprParent.RightParen = next
					p.consumeToken()
				}
			} else {
				r.unexpectedSingleRuneToken('(')
			}
		case token.Literal:
			schemaOrTableName := next
			p.consumeToken()

			periodOrDelimiter, ok := p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				p.parseExpr13Sub2(exprParent, nil, nil, schemaOrTableName, r)
				return exprParent
			}
			if periodOrDelimiter.Value() == "." {
				p.consumeToken()
				literal, ok := p.lookahead(r)
				if !ok {
					return nil
				}
				if literal.Type() == token.Literal {
					p.consumeToken()
					next, ok = p.optionalLookahead(r)
					if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
						p.parseExpr13Sub2(exprParent, schemaOrTableName, periodOrDelimiter, literal, r)
						return exprParent
					}
					if next.Type() == token.Delimiter && next.Value() == "(" {
						p.parseExpr13Sub3(exprParent, schemaOrTableName, periodOrDelimiter, literal, r)
					} else {
						p.parseExpr13Sub2(exprParent, schemaOrTableName, periodOrDelimiter, literal, r)
					}
				} else {
					r.unexpectedToken(token.Literal)
				}
			} else if periodOrDelimiter.Type() == token.Delimiter && periodOrDelimiter.Value() == "(" {
				p.parseExpr13Sub3(exprParent, nil, nil, schemaOrTableName, r)
			} else {
				p.parseExpr13Sub2(exprParent, nil, nil, schemaOrTableName, r)
			}
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return exprParent
	}
	resultExpr := p.parseExprRecursive(exprParent, r)
	if resultExpr != nil {
		return resultExpr
	}
	return exprParent
}

// parseExpr13Sub2 parses the sub-line 2 of line 13 in expr.
func (p *simpleParser) parseExpr13Sub2(exprParent *ast.Expr, schemaName, period, tableName token.Token, r reporter) {
	exprParent.SchemaName = schemaName
	exprParent.Period1 = period
	exprParent.TableName = tableName
}

// parseExpr13Sub3 parses the sub-line 3 of line 13 in expr.
func (p *simpleParser) parseExpr13Sub3(exprParent *ast.Expr, schemaName, period, tableFunctionName token.Token, r reporter) {
	exprParent.SchemaName = schemaName
	exprParent.Period1 = period
	exprParent.TableFunction = tableFunctionName

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		exprParent.LeftParen = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == ")" {
			exprParent.RightParen = next
			p.consumeToken()
		} else {
			exprParent.Expr = p.parseExprSequence(r)
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == ")" {
			exprParent.RightParen = next
			p.consumeToken()
		}
	}
}

// parseWhenThenClause parses the when-then clause as defined in statement.go
func (p *simpleParser) parseWhenThenClause(r reporter) (stmt *ast.WhenThenClause) {
	stmt = &ast.WhenThenClause{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordWhen {
		stmt.When = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordWhen)
	}

	stmt.Expr1 = p.parseExpression(r)
	if stmt.Expr1 == nil {
		r.expectedExpression()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordThen {
		stmt.Then = next
		p.consumeToken()
	}

	stmt.Expr2 = p.parseExpression(r)
	if stmt.Expr2 == nil {
		r.expectedExpression()
	}

	return
}

// parseAttachDatabaseStmt parses a single ATTACH statement as defined in the spec:
// https://sqlite.org/lang_attach.html
func (p *simpleParser) parseAttachDatabaseStmt(r reporter) (stmt *ast.AttachStmt) {
	stmt = &ast.AttachStmt{}
	p.searchNext(r, token.KeywordAttach)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Attach = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordDatabase {
		stmt.Database = next
		p.consumeToken()
	}
	stmt.Expr = p.parseExpression(r)
	if stmt.Expr == nil {
		r.expectedExpression()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordAs {
		stmt.As = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordAs)
		return
	}

	schemaName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaName.Type() != token.Literal {
		r.unexpectedToken(token.Literal)
		return
	}
	stmt.SchemaName = schemaName
	p.consumeToken()
	return
}

// parseDetachDatabaseStmt parses a single DETACH statement as defined in spec:
// https://sqlite.org/lang_detach.html
func (p *simpleParser) parseDetachDatabaseStmt(r reporter) (stmt *ast.DetachStmt) {
	stmt = &ast.DetachStmt{}
	p.searchNext(r, token.KeywordDetach)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Detach = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordDatabase {
		stmt.Database = next
		p.consumeToken()
	}

	schemaName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaName.Type() != token.Literal {
		r.unexpectedToken(token.Literal)
		return
	}
	stmt.SchemaName = schemaName
	p.consumeToken()
	return
}

// parseVacuumStmt parses a single VACUUM statement as defined in the spec:
// https://sqlite.org/lang_vacuum.html
func (p *simpleParser) parseVacuumStmt(r reporter) (stmt *ast.VacuumStmt) {
	stmt = &ast.VacuumStmt{}
	p.searchNext(r, token.KeywordVacuum)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Vacuum = next
	p.consumeToken()

	// optionalLookahead is used because, the lookahead function
	// always looks for the next "real" token and not EOF.
	// Since Just "VACUUM" is a valid statement, we have to accept
	// the fact that there can be no tokens after the first keyword.
	// Same logic is applied for the next INTO keyword check too.
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordInto {
		stmt.Into = next
		p.consumeToken()
		fileName, ok := p.lookahead(r)
		if !ok {
			return
		}
		if fileName.Type() != token.Literal {
			r.unexpectedToken(token.Literal)
			return
		}
		stmt.Filename = fileName
		p.consumeToken()
	}
	return
}

// parseAnalyzeStmt parses a single ANALYZE statement as defined in the spec:
// https://sqlite.org/lang_analyze.html
func (p *simpleParser) parseAnalyzeStmt(r reporter) (stmt *ast.AnalyzeStmt) {
	stmt = &ast.AnalyzeStmt{}
	p.searchNext(r, token.KeywordAnalyze)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Analyze = next
	p.consumeToken()

	// optionalLookahead is used, because ANALYZE alone is a valid statement
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		stmt.TableOrIndexName = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
		return
	}

	period, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	// Since if there is a period, it means there is definitely an
	// existance of a literal, we need a more restrictive condition.
	// Thus we reject if we dont find a literal.
	if period.Value() == "." {
		stmt.Period = period
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() != token.Literal {
			r.unexpectedToken(token.Literal)
		}
		stmt.TableOrIndexName = next
		p.consumeToken()
	}
	return
}

// parseBeginStmt parses a single BEGIN statement as defined in the spec:
// https://sqlite.org/lang_transaction.html
func (p *simpleParser) parseBeginStmt(r reporter) (stmt *ast.BeginStmt) {
	stmt = &ast.BeginStmt{}
	p.searchNext(r, token.KeywordBegin)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Begin = next
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	switch next.Type() {
	case token.KeywordDeferred:
		stmt.Deferred = next
		p.consumeToken()
	case token.KeywordImmediate:
		stmt.Immediate = next
		p.consumeToken()
	case token.KeywordExclusive:
		stmt.Exclusive = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordTransaction {
		stmt.Transaction = next
		p.consumeToken()
	}
	return
}

// parseCommitStmt parses a single COMMIT statement as defined in the spec:
// https://sqlite.org/lang_transaction.html
func (p *simpleParser) parseCommitStmt(r reporter) (stmt *ast.CommitStmt) {
	stmt = &ast.CommitStmt{}
	p.searchNext(r, token.KeywordCommit, token.KeywordEnd)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordCommit {
		stmt.Commit = next
	} else if next.Type() == token.KeywordEnd {
		stmt.End = next
	}
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordTransaction {
		stmt.Transaction = next
		p.consumeToken()
	}
	return
}

// parseRollbackStmt parses a single ROLLBACK statement as defined in the spec:
// https://sqlite.org/lang_transaction.html
func (p *simpleParser) parseRollbackStmt(r reporter) (stmt *ast.RollbackStmt) {
	stmt = &ast.RollbackStmt{}
	p.searchNext(r, token.KeywordRollback)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Rollback = next
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordTransaction {
		stmt.Transaction = next
		p.consumeToken()
	}

	// If the keyword TRANSACTION exists in the statement, we need to
	// check whether TO also exists. Out of TRANSACTION and TO, each not
	// existing and existing, we have the following logic.
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordTo {
		stmt.To = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordSavepoint {
		stmt.Savepoint = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SavepointName = next
	}
	p.consumeToken()
	return
}

// parseCreateStmts parses the multiple variations of CREATE stmts.
// The variations are CREATE INDEX,TABLE,TRIGGER and VIEW.
func (p *simpleParser) parseCreateStmts(stmt *ast.SQLStmt, r reporter) {
	p.searchNext(r, token.KeywordCreate)
	createToken, ok := p.lookahead(r)
	if !ok {
		return
	}
	p.consumeToken()
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordIndex:
		stmt.CreateIndexStmt = p.parseCreateIndexStmt(createToken, r)
	case token.KeywordUnique:
		stmt.CreateIndexStmt = p.parseCreateIndexStmt(createToken, r)
	case token.KeywordTable:
		stmt.CreateTableStmt = p.parseCreateTableStmt(createToken, nil, nil, r)
	case token.KeywordTrigger:
		stmt.CreateTriggerStmt = p.parseCreateTriggerStmt(stmt, createToken, nil, nil, r)
	case token.KeywordView:
		stmt.CreateViewStmt = p.parseCreateViewStmt(createToken, nil, nil, r)
	case token.KeywordTemp:
		tempToken := next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordTable:
			stmt.CreateTableStmt = p.parseCreateTableStmt(createToken, tempToken, nil, r)
		case token.KeywordTrigger:
			stmt.CreateTriggerStmt = p.parseCreateTriggerStmt(stmt, createToken, tempToken, nil, r)
		case token.KeywordView:
			stmt.CreateViewStmt = p.parseCreateViewStmt(createToken, tempToken, nil, r)
		}
	case token.KeywordTemporary:
		temporaryToken := next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordTable:
			stmt.CreateTableStmt = p.parseCreateTableStmt(createToken, nil, temporaryToken, r)
		case token.KeywordTrigger:
			stmt.CreateTriggerStmt = p.parseCreateTriggerStmt(stmt, createToken, nil, temporaryToken, r)
		case token.KeywordView:
			stmt.CreateViewStmt = p.parseCreateViewStmt(createToken, nil, temporaryToken, r)
		}
	case token.KeywordVirtual:
		stmt.CreateVirtualTableStmt = p.parseCreateVirtualTableStmt(createToken, r)
	}
}

// parseCreateIndexStmt parses a single CREATE INDEX statement as defined in the spec:
// https://sqlite.org/lang_createindex.html
func (p *simpleParser) parseCreateIndexStmt(createToken token.Token, r reporter) (stmt *ast.CreateIndexStmt) {
	stmt = &ast.CreateIndexStmt{}
	stmt.Create = createToken

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordUnique {
		stmt.Unique = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordIndex {
		stmt.Index = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordIf {
		stmt.If = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordNot {
			stmt.Not = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordExists {
				stmt.Exists = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordExists)
			}
		} else {
			r.unexpectedToken(token.KeywordNot)
		}
	}

	schemaNameOrIndexName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaNameOrIndexName.Type() == token.Literal {
		// This is the case where there might not be a schemaName
		// We assume that its the table name in the beginning and assign it
		stmt.IndexName = schemaNameOrIndexName
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		// If we find the signs of there being a SchemaName,
		// we over-write the previous value
		if next.Value() == "." {
			stmt.SchemaName = schemaNameOrIndexName
			stmt.Period = next
			p.consumeToken()
			indexName, ok := p.lookahead(r)
			if !ok {
				return
			}
			stmt.IndexName = indexName
			p.consumeToken()
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOn {
		stmt.On = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.TableName = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "(" {
		stmt.LeftParen = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		for next.Value() != ")" {
			stmt.IndexedColumns = append(stmt.IndexedColumns, p.parseIndexedColumn(r))
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
		}
		if next.Value() == ")" {
			stmt.RightParen = next
			p.consumeToken()
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordWhere {
		stmt.Where = next
		p.consumeToken()
		stmt.Expr = p.parseExpression(r)
		if stmt.Expr == nil {
			r.expectedExpression()
		}
	}
	return
}

// parseIndexedColumn parses the indexed-column statement as defined in:
// https://sqlite.org/syntax/indexed-column.html
func (p *simpleParser) parseIndexedColumn(r reporter) (stmt *ast.IndexedColumn) {
	stmt = &ast.IndexedColumn{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.ColumnName = next
		p.consumeToken()
	} else {
		stmt.Expr = p.parseExpression(r)
		if stmt.Expr == nil {
			r.expectedExpression()
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordCollate {
		stmt.Collate = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.CollationName = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordAsc {
		stmt.Asc = next
		p.consumeToken()
	}
	if next.Type() == token.KeywordDesc {
		stmt.Desc = next
		p.consumeToken()
	}
	return
}

// parseCreateTableStmt parses create-table stmt as defined in:
// https://sqlite.org/lang_createtable.html
func (p *simpleParser) parseCreateTableStmt(createToken, tempToken, temporaryToken token.Token, r reporter) (stmt *ast.CreateTableStmt) {
	stmt = &ast.CreateTableStmt{}
	stmt.Create = createToken
	stmt.Temp = tempToken
	stmt.Temporary = temporaryToken
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordTable {
		stmt.Table = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordTable)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordIf {
		stmt.If = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordNot {
			stmt.Not = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordExists {
				stmt.Exists = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordExists)
			}
		} else {
			r.unexpectedToken(token.KeywordNot)
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		stmt.TableName = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "." {
		stmt.Period = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.TableName = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	} else {
		stmt.SchemaName = nil
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.Delimiter:
		stmt.LeftParen = next
		p.consumeToken()
		for {
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.ColumnDef = append(stmt.ColumnDef, p.parseColumnDef(r))
			} else {
				break
			}
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				for {
					if next.Type() == token.KeywordConstraint ||
						next.Type() == token.KeywordPrimary ||
						next.Type() == token.KeywordUnique ||
						next.Type() == token.KeywordCheck ||
						next.Type() == token.KeywordForeign {
						stmt.TableConstraint = append(stmt.TableConstraint, p.parseTableConstraint(r))
					} else {
						break
					}
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Value() == "," {
						p.consumeToken()
					}
					if next.Value() == ")" {
						stmt.RightParen = next
						break
					}
				}
			}
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
				break
			}
		}
	case token.KeywordAs:
		stmt.As = next
		p.consumeToken()
		stmt.SelectStmt = p.parseSelectStmt(nil, r)
	}
	return
}

// parseCreateTriggerStmt parses create-trigger stmts as defined in:
// https://sqlite.org/lang_createtrigger.html
func (p *simpleParser) parseCreateTriggerStmt(sqlStmt *ast.SQLStmt, createToken, tempToken, temporaryToken token.Token, r reporter) (stmt *ast.CreateTriggerStmt) {
	stmt = &ast.CreateTriggerStmt{}
	stmt.Create = createToken
	stmt.Temp = tempToken
	stmt.Temporary = temporaryToken
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordTrigger {
		stmt.Trigger = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordNot {
				stmt.Not = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordExists {
					stmt.Exists = next
					p.consumeToken()
				} else {
					r.unexpectedToken(token.KeywordExists)
				}
			} else {
				r.unexpectedToken(token.KeywordNot)
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.TriggerName = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.TriggerName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		} else {
			stmt.SchemaName = nil
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordBefore:
			stmt.Before = next
			p.consumeToken()
		case token.KeywordAfter:
			stmt.After = next
			p.consumeToken()
		case token.KeywordInstead:
			stmt.Instead = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordOf {
				stmt.Of1 = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordOf)
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordDelete:
			stmt.Delete = next
			p.consumeToken()
		case token.KeywordInsert:
			stmt.Insert = next
			p.consumeToken()
		case token.KeywordUpdate:
			stmt.Update = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordOf {
				stmt.Of2 = next
				p.consumeToken()
				for {
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.Literal {
						stmt.ColumnName = append(stmt.ColumnName, next)
						p.consumeToken()
					}
					if next.Value() == "," {
						p.consumeToken()
					}
					if next.Type() == token.KeywordOn {
						break
					}
				}
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			stmt.On = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.TableName = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordFor {
					stmt.For = next
					p.consumeToken()

					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.KeywordEach {
						stmt.Each = next
						p.consumeToken()

						next, ok = p.lookahead(r)
						if !ok {
							return
						}
						if next.Type() == token.KeywordRow {
							stmt.Row = next
							p.consumeToken()
						} else {
							r.unexpectedToken(token.KeywordRow)
						}
					} else {
						r.unexpectedToken(token.KeywordEach)
					}
				}

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordWhen {
					stmt.When = next
					p.consumeToken()
					stmt.Expr = p.parseExpression(r)
					if stmt.Expr == nil {
						r.expectedExpression()
					}
				}

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordBegin {
					stmt.Begin = next
					p.consumeToken()
				}

				for {
					next, ok = p.optionalLookahead(r)
					if !ok || next.Type() == token.EOF {
						return
					}
					if next.Type() == token.KeywordEnd {
						stmt.End = next
						p.consumeToken()
						return
					}
					if next.Value() == ";" {
						p.consumeToken()
					}
					// Any of the 4 statements can exist according to the FSM of the parser.
					// First, we have to check for existance of "WITH" and parse it through the
					// helper function for stmts beginning with with-clauses
					// The other case can check for the existance of the leading keyword and
					// parse the respective stmts.
					if next.Type() == token.KeywordWith {
						p.parseWithClauseBeginnerStmts(sqlStmt, r)
						if sqlStmt.UpdateStmt != nil {
							stmt.UpdateStmt = append(stmt.UpdateStmt, sqlStmt.UpdateStmt)
							sqlStmt.UpdateStmt = nil
						} else if sqlStmt.InsertStmt != nil {
							stmt.InsertStmt = append(stmt.InsertStmt, sqlStmt.InsertStmt)
							sqlStmt.InsertStmt = nil
						} else if sqlStmt.DeleteStmt != nil {
							stmt.DeleteStmt = append(stmt.DeleteStmt, sqlStmt.DeleteStmt)
							sqlStmt.DeleteStmt = nil
						} else if sqlStmt.SelectStmt != nil {
							stmt.SelectStmt = append(stmt.SelectStmt, sqlStmt.SelectStmt)
							sqlStmt.SelectStmt = nil
						}
					} else {
						if next.Type() == token.KeywordUpdate {
							stmt.UpdateStmt = append(stmt.UpdateStmt, p.parseUpdateStmt(nil, nil, r))
						} else if next.Type() == token.KeywordDelete {
							stmt.DeleteStmt = append(stmt.DeleteStmt, p.parseDeleteStmt(nil, nil, r))
						} else if next.Type() == token.KeywordSelect {
							stmt.SelectStmt = append(stmt.SelectStmt, p.parseSelectStmt(nil, r))
						} else if next.Type() == token.KeywordInsert || next.Type() == token.KeywordReplace {
							stmt.InsertStmt = append(stmt.InsertStmt, p.parseInsertStmt(nil, r))
						}
					}
				}
			} else {
				r.unexpectedToken(token.Literal)
			}
		}
	} else {
		r.unexpectedToken(token.KeywordTrigger)
	}
	return
}

// createViewStmt parses create-view stmts as defined in:
// https://sqlite.org/lang_createview.html
func (p *simpleParser) parseCreateViewStmt(createToken, tempToken, temporaryToken token.Token, r reporter) (stmt *ast.CreateViewStmt) {
	stmt = &ast.CreateViewStmt{}
	stmt.Create = createToken
	stmt.Temp = tempToken
	stmt.Temporary = temporaryToken
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordView {
		stmt.View = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordNot {
				stmt.Not = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordExists {
					stmt.Exists = next
					p.consumeToken()
				} else {
					r.unexpectedToken(token.KeywordExists)
				}
			} else {
				r.unexpectedToken(token.KeywordNot)
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.ViewName = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.ViewName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		} else {
			stmt.SchemaName = nil
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			stmt.LeftParen = next
			p.consumeToken()

			for {
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.Literal {
					stmt.ColumnName = append(stmt.ColumnName, next)
					p.consumeToken()
				} else if next.Value() == "," {
					p.consumeToken()
				} else if next.Type() == token.Delimiter && next.Value() == ")" {
					stmt.RightParen = next
					p.consumeToken()
					break
				} else {
					r.unexpectedToken(token.Literal, token.Delimiter)
					break
				}
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAs {
			stmt.As = next
			p.consumeToken()
			stmt.SelectStmt = p.parseSelectStmt(nil, r)
		}
	} else {
		r.unexpectedToken(token.KeywordView)
	}
	return
}

// parseCreateVirtualTableStmt parses create-virtual-stmts as defined in:
// https://sqlite.org/lang_createvtab.html
func (p *simpleParser) parseCreateVirtualTableStmt(createToken token.Token, r reporter) (stmt *ast.CreateVirtualTableStmt) {
	stmt = &ast.CreateVirtualTableStmt{}
	stmt.Create = createToken

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordVirtual {
		stmt.Virtual = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordTable {
			stmt.Table = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordNot {
				stmt.Not = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordExists {
					stmt.Exists = next
					p.consumeToken()
				} else {
					r.unexpectedToken(token.KeywordExists)
				}
			} else {
				r.unexpectedToken(token.KeywordNot)
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.TableName = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.TableName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		} else {
			stmt.SchemaName = nil
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordUsing {
			stmt.Using = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.ModuleName = next
				p.consumeToken()

				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return
				}
				if next.Type() == token.Delimiter && next.Value() == "(" {
					stmt.LeftParen = next
					p.consumeToken()

					for {
						next, ok = p.lookahead(r)
						if !ok {
							return
						}
						if next.Type() == token.Literal {
							stmt.ModuleArgument = append(stmt.ModuleArgument, next)
							p.consumeToken()
						} else if next.Type() == token.Delimiter && next.Value() == ")" {
							stmt.RightParen = next
							p.consumeToken()
							return
						} else if next.Value() == "," {
							p.consumeToken()
						} else {
							break
						}
					}
				}
			} else {
				r.unexpectedToken(token.Literal)
			}
		}
	} else {
		r.unexpectedToken(token.KeywordVirtual)
	}
	return
}

func (p *simpleParser) parseDeleteStmtHelper(withClause *ast.WithClause, r reporter) (deleteStmt *ast.DeleteStmt) {
	deleteStmt = &ast.DeleteStmt{}

	if withClause != nil {
		deleteStmt.WithClause = withClause
	} else {
		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordWith {
			deleteStmt.WithClause = p.parseWithClause(r)
		}
	}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordDelete {
		deleteStmt.Delete = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordDelete)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordFrom {
		deleteStmt.From = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordFrom)
	}
	deleteStmt.QualifiedTableName = p.parseQualifiedTableName(r)

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordWhere {
		deleteStmt.Where = next
		p.consumeToken()
		deleteStmt.Expr = p.parseExpression(r)
		if deleteStmt.Expr == nil {
			r.expectedExpression()
		}
	}
	return
}

func (p *simpleParser) parseDeleteStmts(sqlStmt *ast.SQLStmt, withClause *ast.WithClause, r reporter) {
	deleteStmt := p.parseDeleteStmtHelper(withClause, r)

	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		sqlStmt.DeleteStmt = deleteStmt
		return
	}
	if next.Type() == token.KeywordOrder || next.Type() == token.KeywordLimit {
		sqlStmt.DeleteStmtLimited = p.parseDeleteStmtLimited(deleteStmt, r)
	}
}

// parseDeleteStmt parses the DELETE statement as defined in:
// https://sqlite.org/lang_delete.html
func (p *simpleParser) parseDeleteStmt(deleteStmt *ast.DeleteStmt, withClause *ast.WithClause, r reporter) *ast.DeleteStmt {
	if deleteStmt != nil {
		return deleteStmt
	}
	return p.parseDeleteStmtHelper(withClause, r)
}

// parseDeleteStmt parses the delete-stmt-limited statement as defined in:
// https://sqlite.org/lang_delete.html
func (p *simpleParser) parseDeleteStmtLimited(deleteStmt *ast.DeleteStmt, r reporter) (stmt *ast.DeleteStmtLimited) {
	stmt = &ast.DeleteStmtLimited{}
	stmt.DeleteStmt = deleteStmt

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOrder {
		stmt.Order = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By = next
			p.consumeToken()

			for {
				stmt.OrderingTerm = append(stmt.OrderingTerm, p.parseOrderingTerm(r))
				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return
				}
				if next.Value() == "," {
					p.consumeToken()
				} else {
					break
				}
			}
		} else {
			r.unexpectedToken(token.KeywordBy)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordLimit {
		stmt.Limit = next
		p.consumeToken()
		stmt.Expr1 = p.parseExpression(r)
		if stmt.Expr1 == nil {
			r.expectedExpression()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordOffset {
			stmt.Offset = next
			p.consumeToken()
		} else if next.Value() == "," {
			stmt.Comma = next
			p.consumeToken()
		}
		stmt.Expr2 = p.parseExpression(r)
		if stmt.Expr2 == nil {
			r.expectedExpression()
		}
	}
	return
}

// parseWithClause parses the WithClause as defined in:
// https://sqlite.org/syntax/with-clause.html
func (p *simpleParser) parseWithClause(r reporter) (withClause *ast.WithClause) {
	withClause = &ast.WithClause{}
	p.searchNext(r, token.KeywordWith)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	withClause.With = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordRecursive {
		withClause.Recursive = next
		p.consumeToken()
	}
	for {
		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.Literal {
			withClause.RecursiveCte = append(withClause.RecursiveCte, p.parseRecursiveCte(r))
		}
		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Value() == "," {
			p.consumeToken()
		} else {
			break
		}
	}
	return
}

func (p *simpleParser) parseRecursiveCte(r reporter) (recursiveCte *ast.RecursiveCte) {
	recursiveCte = &ast.RecursiveCte{}
	recursiveCte.CteTableName = p.parseCteTableName(r)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordAs {
		recursiveCte.As = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordAs)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "(" {
		recursiveCte.LeftParen = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Delimiter)
	}
	recursiveCte.SelectStmt = p.parseSelectStmt(nil, r)
	next, ok = p.lookahead(r)
	if !ok {
		return
	}

	if next.Value() == ")" {
		recursiveCte.RightParen = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Delimiter)
	}
	return
}

// parseCteTableName parses the cte-table-name stmt as defined in:
//https://sqlite.org/syntax/cte-table-name.html
func (p *simpleParser) parseCteTableName(r reporter) (cteTableName *ast.CteTableName) {
	cteTableName = &ast.CteTableName{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		cteTableName.TableName = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Value() == "(" {
		cteTableName.LeftParen = next
		p.consumeToken()
		for {
			columnName, ok := p.lookahead(r)
			if !ok {
				return
			}
			if columnName.Type() == token.Literal {
				cteTableName.ColumnName = append(cteTableName.ColumnName, columnName)
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
				break
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				cteTableName.RightParen = next
				p.consumeToken()
				break
			}
		}
	}
	return
}

// parseSelectStmt parses the select stmt as defined in:
// https://sqlite.org/syntax/select-stmt.html
func (p *simpleParser) parseSelectStmt(withClause *ast.WithClause, r reporter) (stmt *ast.SelectStmt) {
	stmt = &ast.SelectStmt{}

	// parseSelect can be called from withClauseBeginnerStmts or otherwise.
	if withClause != nil {
		stmt.WithClause = withClause
	} else {
		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordWith {
			stmt.WithClause = p.parseWithClause(r)
		}
	}
	var selectCore *ast.SelectCore

	// Keep looping and searching for the select core until its exhausted.
	// We are sure that a select core starts here as its the type of stmt we expect.
	for {
		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordSelect || next.Type() == token.KeywordValues {
			if selectCore != nil {
				// Operated on previous selectCores.
				// If there's no compounding in this statement;
				// strict rule, thus breaking flow.
				if selectCore.CompoundOperator == nil {
					r.unexpectedToken(token.KeywordUnion, token.KeywordIntersect, token.KeywordExcept)
					break
				}
			}
			selectCore = p.parseSelectCore(r)
			stmt.SelectCore = append(stmt.SelectCore, selectCore)
		} else {
			break
		}
	}

	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordOrder {
		stmt.Order = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By = next
			p.consumeToken()

		} else {
			r.unexpectedToken(token.KeywordBy)
		}
		for {
			stmt.OrderingTerm = append(stmt.OrderingTerm, p.parseOrderingTerm(r))
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			} else {
				break
			}
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordLimit {
		stmt.Limit = next
		p.consumeToken()
		stmt.Expr1 = p.parseExpression(r)
		if stmt.Expr1 == nil {
			r.expectedExpression()
		}

		if stmt.Expr1 == nil {
			r.expectedExpression()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordOffset {
			stmt.Offset = next
			p.consumeToken()
			stmt.Expr2 = p.parseExpression(r)
			if stmt.Expr2 == nil {
				r.expectedExpression()
			}
		}
		if next.Value() == "," {
			stmt.Comma = next
			p.consumeToken()
			stmt.Expr2 = p.parseExpression(r)
			if stmt.Expr2 == nil {
				r.expectedExpression()
			}
		}
	}
	return
}

// parseQualifiedTableName parses qualified-table-name as defined in:
// https://sqlite.org/syntax/qualified-table-name.html
func (p *simpleParser) parseQualifiedTableName(r reporter) (stmt *ast.QualifiedTableName) {
	stmt = &ast.QualifiedTableName{}
	schemaOrTableName, ok := p.lookahead(r)
	if !ok {
		return
	}
	// We expect that the first literal can be either the schema name
	// or the table name. When we confirm the existance of a period, we
	// re-assign the table name and when we confirm that there is no
	// period, we reset the schema name value to nil.
	if schemaOrTableName.Type() == token.Literal {
		stmt.TableName = schemaOrTableName
		p.consumeToken()
	}
	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Value() == "." {
		stmt.Period = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = schemaOrTableName
			stmt.TableName = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordAs {
		stmt.As = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.Alias = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordIndexed {
		stmt.Indexed = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.IndexName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		} else {
			r.unexpectedToken(token.KeywordBy)
		}

	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordNot {
		stmt.Not = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIndexed {
			stmt.Indexed = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordIndexed)
		}
	}
	return
}

// parseSelectCore parses a core block of the select statement.
// The select stmt gets multiple select core stmts, where,
// each select core can be either starting with a "SELECT" keyword
// or a "VALUES" keyword. The compound operator belongs to the select core
// stmt which is right before it and not the one after.
// Due to a bit of controversy in the table-or-subquery and the join-clause
// stmts after FROM in SELECT core, all statements are parsed into join-clause.
func (p *simpleParser) parseSelectCore(r reporter) (stmt *ast.SelectCore) {
	stmt = &ast.SelectCore{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordSelect {
		stmt.Select = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordDistinct {
			stmt.Distinct = next
			p.consumeToken()
		} else if next.Type() == token.KeywordAll {
			stmt.All = next
			p.consumeToken()
		}

		for {
			stmt.ResultColumn = append(stmt.ResultColumn, p.parseResultColumn(r))
			next, ok = p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			} else {
				break
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordFrom {
			stmt.From = next
			p.consumeToken()
			stmt.JoinClause = p.parseJoinClause(r)
			if stmt.JoinClause == nil {
				for {
					stmt.TableOrSubquery = append(stmt.TableOrSubquery, p.parseTableOrSubquery(r))
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Value() == "," {
						p.consumeToken()
					} else {
						break
					}
				}
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordWhere {
			stmt.Where = next
			p.consumeToken()
			stmt.Expr1 = p.parseExpression(r)
			if stmt.Expr1 == nil {
				r.expectedExpression()
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordGroup {
			stmt.Group = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordBy {
				stmt.By = next
				p.consumeToken()
			}
			for {
				expression := p.parseExpression(r)
				if expression != nil {
					stmt.Expr2 = append(stmt.Expr2, expression)
				}
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Value() == "," {
					p.consumeToken()
				} else {
					break
				}
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordHaving {
				stmt.Having = next
				p.consumeToken()
				stmt.Expr3 = p.parseExpression(r)
				if stmt.Expr3 == nil {
					r.expectedExpression()
				}
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordWindow {
			// Consuming the keyword window
			stmt.Window = next
			p.consumeToken()
			for {
				stmt.NamedWindow = append(stmt.NamedWindow, p.parseNamedWindow(r))

				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return
				}
				if next.Value() == "," {
					p.consumeToken()
				} else {
					break
				}
			}
		}
	}
	if next.Type() == token.KeywordValues {
		stmt.Values = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			for {
				parExpr := p.parseParenthesizedExpression(r)
				if parExpr != nil {
					stmt.ParenthesizedExpressions = append(stmt.ParenthesizedExpressions, parExpr)
				}
				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					if stmt.ParenthesizedExpressions == nil {
						r.expectedExpression()
					}
					return
				}
				// Do not allow nil exprs.
				if next.Value() == "," && parExpr != nil {
					p.consumeToken()
				} else {
					if len(stmt.ParenthesizedExpressions) == 0 {
						r.expectedExpression()
					}
					break
				}
			}
		} else {
			r.unexpectedSingleRuneToken('(')
		}
	}

	// Checking whether there is a token that leads to a part of the statement
	// ensures that stmt.CompoundOperator is nil, instead of an assigned empty value.
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordUnion || next.Type() == token.KeywordIntersect || next.Type() == token.KeywordExcept {
		stmt.CompoundOperator = p.parseCompoundOperator(r)
	}
	return
}

// parseResultColumn parses result-column as defined in:
// https://sqlite.org/syntax/result-column.html
func (p *simpleParser) parseResultColumn(r reporter) (stmt *ast.ResultColumn) {
	stmt = &ast.ResultColumn{}
	tableNameOrAsteriskOrExpr, ok := p.lookahead(r)
	if !ok {
		return
	}
	switch tableNameOrAsteriskOrExpr.Type() {
	case token.BinaryOperator:
		stmt.Asterisk = tableNameOrAsteriskOrExpr
		p.consumeToken()
	case token.Literal:
		p.consumeToken()
		period, ok := p.optionalLookahead(r)
		if !ok || period.Type() == token.EOF || period.Type() == token.StatementSeparator {
			// If the statement ends on a literal, its an expression of form literal.
			stmt.Expr = &ast.Expr{LiteralValue: tableNameOrAsteriskOrExpr}
			return
		}
		if period.Value() == "." {
			p.consumeToken()
			next, ok := p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				p.consumeToken()
				stmt.Expr = p.parseExpr2(tableNameOrAsteriskOrExpr, period, next, r)
			} else if next.Value() == "*" {
				stmt.Asterisk = next
				p.consumeToken()
				stmt.TableName = tableNameOrAsteriskOrExpr
				stmt.Period = period
			}
		} else {
			// Conditions for recursive expressions or single expressions.
			if recExpr := p.parseExprRecursive(&ast.Expr{LiteralValue: tableNameOrAsteriskOrExpr}, r); recExpr != nil {
				stmt.Expr = recExpr
			} else {
				if singleExpr := p.parseExprBeginWithLiteral(tableNameOrAsteriskOrExpr, r); singleExpr != nil {
					stmt.Expr = singleExpr
				} else {
					stmt.Expr = &ast.Expr{LiteralValue: tableNameOrAsteriskOrExpr}
				}
			}
		}
		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordAs {
			stmt.As = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.Literal {
			stmt.ColumnAlias = next
			p.consumeToken()
		}
	default:
		stmt.Expr = p.parseExpression(r)
		if stmt.Expr == nil {
			r.expectedExpression()
		}

		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordAs {
			stmt.As = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.Literal {
			stmt.ColumnAlias = next
			p.consumeToken()
		}
	}
	return
}

func (p *simpleParser) parseNamedWindow(r reporter) (stmt *ast.NamedWindow) {
	stmt = &ast.NamedWindow{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.WindowName = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAs {
			stmt.As = next
			p.consumeToken()
			stmt.WindowDefn = p.parseWindowDefn(r)
		} else {
			r.unexpectedToken(token.KeywordAs)
		}
	}
	return
}

// parseWindowDefn parses window-defn as defined in:
// https://sqlite.org/syntax/window-defn.html
func (p *simpleParser) parseWindowDefn(r reporter) (stmt *ast.WindowDefn) {
	stmt = &ast.WindowDefn{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "(" {
		stmt.LeftParen = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.BaseWindowName = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordPartition {
		stmt.Partition = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By1 = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordBy)
		}
		for {
			expression := p.parseExpression(r)
			if expression != nil {
				stmt.Expr = append(stmt.Expr, expression)
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			} else {
				break
			}
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOrder {
		stmt.Order = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By2 = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordBy)
		}
		// possible area of error
		// When there are statements that arent parsed by ordering term, they might loop here forever.
		// Fix might be to check for expr existing in the beginning, which is not possible due to the
		// vastness of expr.
		for {
			stmt.OrderingTerm = append(stmt.OrderingTerm, p.parseOrderingTerm(r))
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			} else {
				break
			}
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordRange || next.Type() == token.KeywordRows || next.Type() == token.KeywordGroups {
		stmt.FrameSpec = p.parseFrameSpec(r)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == ")" {
		stmt.RightParen = next
		p.consumeToken()
	}
	return
}

// parsesOrderingTerm parses ordering-term as defined in:
// https://sqlite.org/syntax/ordering-term.html
func (p *simpleParser) parseOrderingTerm(r reporter) (stmt *ast.OrderingTerm) {
	stmt = &ast.OrderingTerm{}
	stmt.Expr = p.parseExpression(r)
	if stmt.Expr == nil {
		r.expectedExpression()
	}

	// Since Expr can take in COLLATE and collation-name, it has been omitted and pushed to expr.
	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordAsc {
		stmt.Asc = next
		p.consumeToken()
	}
	if next.Type() == token.KeywordDesc {
		stmt.Desc = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordNulls {
		stmt.Nulls = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordFirst {
			stmt.First = next
			p.consumeToken()
		}
		if next.Type() == token.KeywordLast {
			stmt.Last = next
			p.consumeToken()
		}
	}
	return
}

// parseFrameSpec parses frame-spec as defined in:
// https://sqlite.org/syntax/frame-spec.html
func (p *simpleParser) parseFrameSpec(r reporter) (stmt *ast.FrameSpec) {
	stmt = &ast.FrameSpec{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordRange:
		stmt.Range = next
		p.consumeToken()
	case token.KeywordRows:
		stmt.Rows = next
		p.consumeToken()
	case token.KeywordGroups:
		stmt.Groups = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordBetween:
		stmt.Between = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordUnbounded:
			stmt.Unbounded1 = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordPreceding {
				stmt.Preceding1 = next
				p.consumeToken()
			}
		case token.KeywordCurrent:
			stmt.Current1 = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordRow {
				// This is set as Row1 because this is one of the either paths
				// taken by the FSM. Other path has 2x ROW's.
				stmt.Row1 = next
				p.consumeToken()
			}
		default:
			stmt.Expr1 = p.parseExpression(r)
			if stmt.Expr1 == nil {
				r.expectedExpression()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordPreceding {
				stmt.Preceding1 = next
				p.consumeToken()
			}
			if next.Type() == token.KeywordFollowing {
				stmt.Following1 = next
				p.consumeToken()
			}
		}
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAnd {
			stmt.And = next
			p.consumeToken()
		}
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordUnbounded:
			stmt.Unbounded2 = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordFollowing {
				stmt.Following2 = next
				p.consumeToken()
			}
		case token.KeywordCurrent:
			stmt.Current2 = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordRow {
				stmt.Row2 = next
				p.consumeToken()
			}
		default:
			stmt.Expr2 = p.parseExpression(r)
			if stmt.Expr2 == nil {
				r.expectedExpression()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordPreceding {
				stmt.Preceding2 = next
				p.consumeToken()
			}
			if next.Type() == token.KeywordFollowing {
				stmt.Following2 = next
				p.consumeToken()
			}
		}
	case token.KeywordUnbounded:
		stmt.Unbounded1 = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordPreceding {
			stmt.Preceding1 = next
			p.consumeToken()
		}
	case token.KeywordCurrent:
		stmt.Current1 = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordRow {
			stmt.Row1 = next
			p.consumeToken()
		}
	default:
		stmt.Expr1 = p.parseExpression(r)
		if stmt.Expr1 == nil {
			r.expectedExpression()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordPreceding {
			stmt.Preceding1 = next
			p.consumeToken()
		}
	}
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordExclude {
		stmt.Exclude = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordNo:
			stmt.No = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordOthers {
				stmt.Others = next
				p.consumeToken()
			}
		case token.KeywordCurrent:
			stmt.Current3 = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordRow {
				stmt.Row3 = next
				p.consumeToken()
			}
		case token.KeywordGroup:
			stmt.Group = next
			p.consumeToken()
		case token.KeywordTies:
			stmt.Ties = next
			p.consumeToken()
		}
	}
	return
}

func (p *simpleParser) parseParenthesizedExpression(r reporter) (stmt *ast.ParenthesizedExpressions) {
	stmt = &ast.ParenthesizedExpressions{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "(" {
		stmt.LeftParen = next
		p.consumeToken()
		for {
			expr := p.parseExpression(r)
			if expr != nil {
				stmt.Exprs = append(stmt.Exprs, expr)
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
				break
			}
			// Do not allow nil exprs.
			if next.Value() == "," && expr != nil {
				p.consumeToken()
			} else {
				r.expectedExpression()
				break
			}
		}
	}
	// Minimum of one expr must exist.
	if len(stmt.Exprs) == 0 {
		stmt = nil
	}
	return
}

// parseCompoundOperator parses compound-operator as defined in:
// https://sqlite.org/syntax/compound-operator.html
func (p *simpleParser) parseCompoundOperator(r reporter) (stmt *ast.CompoundOperator) {
	stmt = &ast.CompoundOperator{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordUnion {
		stmt.Union = next
		p.consumeToken()
		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordAll {
			stmt.All = next
			p.consumeToken()
		}
	}
	if next.Type() == token.KeywordIntersect {
		stmt.Intersect = next
		p.consumeToken()
	}
	if next.Type() == token.KeywordExcept {
		stmt.Except = next
		p.consumeToken()
	}
	return
}

// parsetableOrSubquery parses table-or-subquery as defined in:
// https://sqlite.org/syntax/table-or-subquery.html
func (p *simpleParser) parseTableOrSubquery(r reporter) (stmt *ast.TableOrSubquery) {
	stmt = &ast.TableOrSubquery{}
	schemaOrTableNameOrLeftPar, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaOrTableNameOrLeftPar.Type() == token.Literal {
		stmt.TableName = schemaOrTableNameOrLeftPar
		p.consumeToken()

		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		var tableNameOrTableFunctionName token.Token
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			tableNameOrTableFunctionName, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.SchemaName = schemaOrTableNameOrLeftPar
				stmt.TableName = tableNameOrTableFunctionName
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordAs || next.Type() == token.KeywordIndexed || next.Type() == token.Literal || next.Type() == token.KeywordNot {
			if next.Type() == token.KeywordAs {
				stmt.As = next
				p.consumeToken()
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.Literal {
					stmt.TableAlias = next
					p.consumeToken()
				} else {
					r.unexpectedToken(token.Literal)
				}
			}
			if next.Type() == token.Literal {
				stmt.TableAlias = next
				p.consumeToken()
			}

			next, ok = p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				return
			}
			if next.Type() == token.KeywordIndexed {
				stmt.Indexed = next
				p.consumeToken()
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordBy {
					stmt.By = next
					p.consumeToken()
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.Literal {
						stmt.IndexName = next
						p.consumeToken()
					} else {
						r.unexpectedToken(token.Literal)
					}
				} else {
					r.unexpectedToken(token.KeywordBy)
				}
			}
			if next.Type() == token.KeywordNot {
				stmt.Not = next
				p.consumeToken()
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordIndexed {
					stmt.Indexed = next
					p.consumeToken()
				}
			}
		} else {
			if next.Value() == "(" {
				stmt.TableFunctionName = tableNameOrTableFunctionName
				for {
					// Since this rule allows an open bracket, we need to check whether an expr
					// exists before we allow it to look for an expresion to avoid null ptr errors.
					next, ok := p.lookahead(r)
					if !ok {
						return
					}
					if next.Value() == ")" {
						stmt.RightParen = next
						p.consumeToken()
						break
					}
					expression := p.parseExpression(r)
					if expression != nil {
						stmt.Expr = append(stmt.Expr, expression)
					}

					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Value() == "," {
						p.consumeToken()
					}
				}

				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return
				}
				if next.Type() == token.KeywordAs {
					stmt.As = next
					p.consumeToken()
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.Literal {
						stmt.TableAlias = next
						p.consumeToken()
					}
				}
				if next.Type() == token.Literal {
					stmt.TableAlias = next
					p.consumeToken()
				}
			}
		}
	} else if schemaOrTableNameOrLeftPar.Value() == "(" {
		stmt.SelectStmt = p.parseSelectStmt(nil, r)
		if stmt.SelectStmt == nil {
			stmt.JoinClause = p.parseJoinClause(r)
			if stmt.JoinClause == nil {
				for {
					next, ok := p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.Delimiter || next.Type() == token.Literal {
						stmt.TableOrSubquery = append(stmt.TableOrSubquery, p.parseTableOrSubquery(r))
					} else {
						break
					}

					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Value() == "," {
						p.consumeToken()
					}
					if next.Value() == ")" {
						stmt.RightParen = next
						p.consumeToken()
						break
					}
				}
			}
			next, ok := p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
				return
			}
		} else {
			next, ok := p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordAs {
				stmt.As = next
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.TableAlias = next
				p.consumeToken()
			}
		}
	}
	return
}

// parseJoinClause parses join-clause as defined in:
// https://sqlite.org/syntax/join-clause.html
func (p *simpleParser) parseJoinClause(r reporter) (stmt *ast.JoinClause) {
	stmt = &ast.JoinClause{}
	stmt.TableOrSubquery = p.parseTableOrSubquery(r)

	for {
		next, ok := p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordNatural ||
			next.Type() == token.KeywordJoin ||
			next.Value() == "," ||
			next.Type() == token.KeywordLeft ||
			next.Type() == token.KeywordInner ||
			next.Type() == token.KeywordCross {
			stmt.JoinClausePart = append(stmt.JoinClausePart, p.parseJoinClausePart(r))
		} else {
			break
		}
	}
	return
}

func (p *simpleParser) parseJoinClausePart(r reporter) (stmt *ast.JoinClausePart) {
	stmt = &ast.JoinClausePart{}
	stmt.JoinOperator = p.parseJoinOperator(r)
	stmt.TableOrSubquery = p.parseTableOrSubquery(r)
	// This check for existance of join constraint is necessary to return a nil
	// value of join constraint, before an empty value is assigned to it
	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if !(next.Type() == token.KeywordOn || next.Type() == token.KeywordUsing) {
		return
	}
	stmt.JoinConstraint = p.parseJoinConstraint(r)
	return
}

// parseJoinConstraint parses join-constraint as defined in:
// https://sqlite.org/syntax/join-constraint.html
func (p *simpleParser) parseJoinConstraint(r reporter) (stmt *ast.JoinConstraint) {
	stmt = &ast.JoinConstraint{}
	next, ok := p.optionalLookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOn {
		stmt.On = next
		p.consumeToken()
		stmt.Expr = p.parseExpression(r)
		if stmt.Expr == nil {
			r.expectedExpression()
		}
	}
	if next.Type() == token.KeywordUsing {
		stmt.Using = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Value() == "(" {
			stmt.LeftParen = next
			p.consumeToken()
		}
		for {
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.ColumnName = append(stmt.ColumnName, next)
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			} else if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
				break
			}
		}
	}
	return
}

// parseJoinOperator parses join-operator as defined in:
// https://sqlite.org/syntax/join-operator.html
func (p *simpleParser) parseJoinOperator(r reporter) (stmt *ast.JoinOperator) {
	stmt = &ast.JoinOperator{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "," {
		stmt.Comma = next
		p.consumeToken()
		return
	}
	if next.Type() == token.KeywordJoin {
		stmt.Join = next
		p.consumeToken()
		return
	}
	if next.Type() == token.KeywordNatural {
		stmt.Natural = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordLeft:
		stmt.Left = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOuter {
			stmt.Outer = next
			p.consumeToken()
		}
	case token.KeywordInner:
		stmt.Inner = next
		p.consumeToken()
	case token.KeywordCross:
		stmt.Cross = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordJoin {
		stmt.Join = next
		p.consumeToken()
	}
	return
}

func (p *simpleParser) parseTableConstraint(r reporter) (stmt *ast.TableConstraint) {
	stmt = &ast.TableConstraint{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordConstraint {
		stmt.Constraint = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.Name = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordPrimary || next.Type() == token.KeywordUnique {
		if next.Type() == token.KeywordPrimary {
			stmt.Primary = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordKey {
				stmt.Key = next
				p.consumeToken()
			}
		} else {
			stmt.Unique = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Value() == "(" {
			stmt.LeftParen = next
			p.consumeToken()
		}
		for {
			stmt.IndexedColumn = append(stmt.IndexedColumn, p.parseIndexedColumn(r))
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			}
			if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
				break
			}
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordOn {
			stmt.ConflictClause = p.parseConflictClause(r)
		}
	} else {
		if next.Type() == token.KeywordCheck {
			stmt.Check = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "(" {
				stmt.LeftParen = next
				p.consumeToken()
			}
			stmt.Expr = p.parseExpression(r)
			if stmt.Expr == nil {
				r.expectedExpression()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
			}
		}
		if next.Type() == token.KeywordForeign {
			stmt.Foreign = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordKey {
				stmt.Key = next
				p.consumeToken()
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Value() == "(" {
					stmt.LeftParen = next
					p.consumeToken()
				}
				for {
					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.Literal {
						stmt.ColumnName = append(stmt.ColumnName, next)
						p.consumeToken()
					}
					if next.Value() == "," {
						p.consumeToken()
					}
					if next.Value() == ")" {
						stmt.RightParen = next
						p.consumeToken()
						break
					}
				}
				stmt.ForeignKeyClause = p.parseForeignKeyClause(r)
			} else {
				r.unexpectedToken(token.KeywordKey)
			}
		}
	}
	return
}

// parseInsertStmt parses insert-stmt as defined in:
// https://sqlite.org/lang_insert.html
func (p *simpleParser) parseInsertStmt(withClause *ast.WithClause, r reporter) (stmt *ast.InsertStmt) {
	stmt = &ast.InsertStmt{}

	if withClause != nil {
		stmt.WithClause = withClause
	} else {
		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordWith {
			stmt.WithClause = p.parseWithClause(r)
		}
	}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordInsert {
		stmt.Insert = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOr {
			stmt.Or = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			switch next.Type() {
			case token.KeywordReplace:
				stmt.Replace = next
			case token.KeywordRollback:
				stmt.Rollback = next
			case token.KeywordAbort:
				stmt.Abort = next
			case token.KeywordFail:
				stmt.Fail = next
			case token.KeywordIgnore:
				stmt.Ignore = next
			}
			p.consumeToken()
		}
	} else if next.Type() == token.KeywordReplace {
		stmt.Replace = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordInsert, token.KeywordReplace)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordInto {
		stmt.Into = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		stmt.TableName = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "." {
		stmt.Period = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.TableName = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	} else {
		stmt.SchemaName = nil
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordAs {
		stmt.As = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.Alias = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		stmt.LeftParen = next
		p.consumeToken()
		for {
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.ColumnName = append(stmt.ColumnName, next)
				p.consumeToken()
			}
			if next.Value() == "," {
				p.consumeToken()
			}
			if next.Type() == token.Delimiter && next.Value() == ")" {
				stmt.RightParen = next
				p.consumeToken()
				break
			}
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordValues {
		stmt.SelectStmt = p.parseSelectStmt(nil, r)
		// Since VALUES and parenthesized expressions can be parsed on its own in insert-stmt and also in select-stmt,
		// the way of distinction of which goes where is the following.
		// If there cant be multiple values of select core AND nothing that the select-stmt parses exists after the
		// VALUES and parenthesized expressions clause, we move the stmts to insert-stmt's variables.
		if stmt.SelectStmt.Order == nil && stmt.SelectStmt.Limit == nil && !(len(stmt.SelectStmt.SelectCore) > 1) {
			stmt.Values = stmt.SelectStmt.SelectCore[0].Values
			stmt.ParenthesizedExpressions = stmt.SelectStmt.SelectCore[0].ParenthesizedExpressions
			stmt.SelectStmt = nil
		}
	} else if next.Type() == token.KeywordDefault {
		stmt.Default = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordValues {
			stmt.Values = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordValues)
		}
	} else if next.Type() == token.KeywordSelect || next.Type() == token.KeywordWith {
		stmt.SelectStmt = p.parseSelectStmt(nil, r)
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordOn {
		stmt.UpsertClause = p.parseUpsertClause(r)
	} else {
		r.unexpectedToken(token.KeywordOn)
	}

	return
}

// parseUpsertClause parses upsert-clause as defined in:
// https://sqlite.org/syntax/upsert-clause.html
func (p *simpleParser) parseUpsertClause(r reporter) (clause *ast.UpsertClause) {
	clause = &ast.UpsertClause{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOn {
		clause.On = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordConflict {
			clause.Conflict = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordConflict)
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			clause.LeftParen = next
			p.consumeToken()
			for {
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Value() == "," {
					p.consumeToken()
				} else if next.Type() == token.Delimiter && next.Value() == ")" {
					clause.RightParen = next
					p.consumeToken()
					break
				} else {
					clause.IndexedColumn = append(clause.IndexedColumn, p.parseIndexedColumn(r))
				}
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordWhere {
			clause.Where1 = next
			p.consumeToken()
			clause.Expr1 = p.parseExpression(r)
			if clause.Expr1 == nil {
				r.expectedExpression()
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordDo {
			clause.Do = next
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordNothing {
			clause.Nothing = next
			p.consumeToken()
			return
		}
		if next.Type() == token.KeywordUpdate {
			clause.Update = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordSet {
				clause.Set = next
				p.consumeToken()

			} else {
				r.unexpectedToken(token.KeywordSet)
			}
		}

		for {
			next, ok = p.optionalLookahead(r)
			if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
				return
			}
			if next.Type() == token.KeywordWhere {
				break
			}
			if next.Value() == "," {
				p.consumeToken()
			}
			clause.UpdateSetter = append(clause.UpdateSetter, p.parseUpdateSetter(r))
		}

		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordWhere {
			clause.Where2 = next
			p.consumeToken()
			clause.Expr2 = p.parseExpression(r)
			if clause.Expr2 == nil {
				r.expectedExpression()
			}
		}
	}
	return
}

// parseUpdateSetter parses the update-setter clause as defined in statement.go
func (p *simpleParser) parseUpdateSetter(r reporter) (stmt *ast.UpdateSetter) {
	stmt = &ast.UpdateSetter{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		stmt.ColumnNameList = p.parseColumnNameList(r)
	} else if next.Type() == token.Literal {
		stmt.ColumnName = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Delimiter, token.Literal)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "=" {
		stmt.Assign = next
		p.consumeToken()
		stmt.Expr = p.parseExpression(r)
		if stmt.Expr == nil {
			r.expectedExpression()
		}
	} else {
		r.unexpectedSingleRuneToken('=')
	}
	return
}

// parseColumnNameList parses column-name-list as defined in:
// https://sqlite.org/syntax/column-name-list.html
func (p *simpleParser) parseColumnNameList(r reporter) (stmt *ast.ColumnNameList) {
	stmt = &ast.ColumnNameList{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Delimiter && next.Value() == "(" {
		stmt.LeftParen = next
		p.consumeToken()
	} else {
		r.unexpectedSingleRuneToken('(')
	}
	for {
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == ")" {
			stmt.RightParen = next
			p.consumeToken()
			break
		}
		if next.Type() == token.Literal {
			stmt.ColumnName = append(stmt.ColumnName, next)
			p.consumeToken()
		}
		if next.Value() == "," {
			p.consumeToken()
		}
	}
	return
}

// parseUpdateStmtHelper parses the entire update statement.
// This function is separated in order to avoid code duplication.
func (p *simpleParser) parseUpdateStmtHelper(withClause *ast.WithClause, r reporter) (updateStmt *ast.UpdateStmt) {
	updateStmt = &ast.UpdateStmt{}

	if withClause != nil {
		updateStmt.WithClause = withClause
	} else {
		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordWith {
			updateStmt.WithClause = p.parseWithClause(r)
		}
	}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordUpdate {
		updateStmt.Update = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOr {
			updateStmt.Or = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			switch next.Type() {
			case token.KeywordRollback:
				updateStmt.Rollback = next
			case token.KeywordAbort:
				updateStmt.Abort = next
			case token.KeywordReplace:
				updateStmt.Replace = next
			case token.KeywordFail:
				updateStmt.Fail = next
			case token.KeywordIgnore:
				updateStmt.Ignore = next
			}
			p.consumeToken()
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			updateStmt.QualifiedTableName = p.parseQualifiedTableName(r)
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordSet {
				updateStmt.Set = next
				p.consumeToken()

				for {
					next, ok = p.optionalLookahead(r)
					if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
						return
					}
					if next.Type() == token.KeywordWhere || next.Type() == token.KeywordOrder || next.Type() == token.KeywordLimit {
						break
					}
					if next.Value() == "," {
						p.consumeToken()
					}
					updateStmt.UpdateSetter = append(updateStmt.UpdateSetter, p.parseUpdateSetter(r))
				}

				next, ok = p.optionalLookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordWhere {
					updateStmt.Where = next
					p.consumeToken()
					updateStmt.Expr = p.parseExpression(r)
					if updateStmt.Expr == nil {
						r.expectedExpression()
					}
				}

			} else {
				r.unexpectedToken(token.KeywordSet)
			}
		} else {
			r.unexpectedToken(token.Literal)
		}
	} else {
		r.unexpectedToken(token.KeywordUpdate)
	}
	return
}

// parseUpdateStmts parses all different variants of the Update stmt.
// This includes with or without the WithClause, and the "Limited" version of the stmt.
func (p *simpleParser) parseUpdateStmts(sqlStmt *ast.SQLStmt, withClause *ast.WithClause, r reporter) {
	updateStmt := p.parseUpdateStmtHelper(withClause, r)

	next, ok := p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		sqlStmt.UpdateStmt = updateStmt
	}
	if next.Type() == token.KeywordOrder || next.Type() == token.KeywordLimit {
		sqlStmt.UpdateStmtLimited = p.parseUpdateStmtLimited(updateStmt, r)
	}
}

// parseUpdateStmt parses update-stmt as defined in:
// https://sqlite.org/lang_update.html
func (p *simpleParser) parseUpdateStmt(updateStmt *ast.UpdateStmt, withClause *ast.WithClause, r reporter) (stmt *ast.UpdateStmt) {
	if updateStmt != nil {
		return updateStmt
	}
	return p.parseUpdateStmtHelper(withClause, r)
}

// parseUpdateStmtLimited parses update-stmt-limited as defined in:
// https://sqlite.org/syntax/update-stmt-limited.html
func (p *simpleParser) parseUpdateStmtLimited(updateStmt *ast.UpdateStmt, r reporter) (stmt *ast.UpdateStmtLimited) {
	stmt = &ast.UpdateStmtLimited{}
	stmt.UpdateStmt = updateStmt

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOrder {
		stmt.Order = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By = next
			p.consumeToken()

			for {
				stmt.OrderingTerm = append(stmt.OrderingTerm, p.parseOrderingTerm(r))
				next, ok = p.optionalLookahead(r)
				if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
					return
				}
				if next.Value() == "," {
					p.consumeToken()
				} else {
					break
				}
			}
		} else {
			r.unexpectedToken(token.KeywordBy)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.KeywordLimit {
		stmt.Limit = next
		p.consumeToken()
		stmt.Expr1 = p.parseExpression(r)
		if stmt.Expr1 == nil {
			r.expectedExpression()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			return
		}
		if next.Type() == token.KeywordOffset {
			stmt.Offset = next
			p.consumeToken()
		} else if next.Value() == "," {
			stmt.Comma = next
			p.consumeToken()
		}
		stmt.Expr2 = p.parseExpression(r)
		if stmt.Expr2 == nil {
			r.expectedExpression()
		}
	}
	return
}

// parseSavepointStmt parses the savepoint-stmt as defined in:
// https://sqlite.org/lang_savepoint.html
func (p *simpleParser) parseSavepointStmt(r reporter) (stmt *ast.SavepointStmt) {
	stmt = &ast.SavepointStmt{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordSavepoint {
		stmt.Savepoint = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordSavepoint)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SavepointName = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
	}
	return
}

// parseReleaseStmt parses the release-stmt as defined in:
// https://sqlite.org/lang_savepoint.html
func (p *simpleParser) parseReleaseStmt(r reporter) (stmt *ast.ReleaseStmt) {
	stmt = &ast.ReleaseStmt{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordRelease {
		stmt.Release = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordSavepoint)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordSavepoint {
		stmt.Savepoint = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SavepointName = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
	}
	return
}

// parseReIndexStmt parses the release-stmt as defined in:
// https://sqlite.org/lang_reindex.html
func (p *simpleParser) parseReIndexStmt(r reporter) (stmt *ast.ReIndexStmt) {
	stmt = &ast.ReIndexStmt{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordReIndex {
		stmt.ReIndex = next
		p.consumeToken()
	}

	collationOrSchemaName, ok := p.optionalLookahead(r)
	if !ok || collationOrSchemaName.Type() == token.EOF || collationOrSchemaName.Type() == token.StatementSeparator {
		return
	}
	if collationOrSchemaName.Type() == token.Literal {
		stmt.CollationName = collationOrSchemaName
		stmt.SchemaName = collationOrSchemaName
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		stmt.SchemaName = nil
		return
	}
	if next.Value() == "." {
		stmt.CollationName = nil
		stmt.Period = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
		return
	}
	if next.Type() == token.Literal {
		stmt.TableOrIndexName = next
		p.consumeToken()
	}
	return
}

// parseDropIndexStmt parses drop-index stmts as defined in:
// https://sqlite.org/lang_dropindex.html
func (p *simpleParser) parseDropIndexStmt(dropToken token.Token, r reporter) (stmt *ast.DropIndexStmt) {
	stmt = &ast.DropIndexStmt{}

	stmt.Drop = dropToken

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordIndex {
		stmt.Index = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordExists {
				stmt.Exists = next
				p.consumeToken()
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.IndexName = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			stmt.SchemaName = nil
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.IndexName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		}
	} else {
		r.unexpectedToken(token.KeywordIndex)
	}
	return
}

// parseDropTableStmt parses drop-index stmts as defined in:
// https://sqlite.org/lang_droptable.html
func (p *simpleParser) parseDropTableStmt(dropToken token.Token, r reporter) (stmt *ast.DropTableStmt) {
	stmt = &ast.DropTableStmt{}

	stmt.Drop = dropToken

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordTable {
		stmt.Table = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordExists {
				stmt.Exists = next
				p.consumeToken()
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.TableName = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			stmt.SchemaName = nil
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.TableName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		}
	} else {
		r.unexpectedToken(token.KeywordTable)
	}
	return
}

// parseDropTriggerStmt parses drop-index stmts as defined in:
// https://sqlite.org/lang_droptrigger.html
func (p *simpleParser) parseDropTriggerStmt(dropToken token.Token, r reporter) (stmt *ast.DropTriggerStmt) {
	stmt = &ast.DropTriggerStmt{}

	stmt.Drop = dropToken

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordTrigger {
		stmt.Trigger = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordExists {
				stmt.Exists = next
				p.consumeToken()
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.TriggerName = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			stmt.SchemaName = nil
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.TriggerName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		}
	} else {
		r.unexpectedToken(token.KeywordTrigger)
	}
	return
}

// parseDropViewStmt parses drop-index stmts as defined in:
// https://sqlite.org/lang_dropview.html
func (p *simpleParser) parseDropViewStmt(dropToken token.Token, r reporter) (stmt *ast.DropViewStmt) {
	stmt = &ast.DropViewStmt{}

	stmt.Drop = dropToken

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordView {
		stmt.View = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIf {
			stmt.If = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordExists {
				stmt.Exists = next
				p.consumeToken()
			}
		}

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.SchemaName = next
			stmt.ViewName = next
			p.consumeToken()
		}

		next, ok = p.optionalLookahead(r)
		if !ok || next.Type() == token.EOF || next.Type() == token.StatementSeparator {
			stmt.SchemaName = nil
			return
		}
		if next.Value() == "." {
			stmt.Period = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.ViewName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		}
	} else {
		r.unexpectedToken(token.KeywordView)
	}
	return
}

// parseFilterClause parses the filter-clause as defined in:
// https://sqlite.org/syntax/filter-clause.html
func (p *simpleParser) parseFilterClause(r reporter) (clause *ast.FilterClause) {
	clause = &ast.FilterClause{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordFilter {
		clause.Filter = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			clause.LeftParen = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordWhere {
				clause.Where = next
				p.consumeToken()

				clause.Expr = p.parseExpression(r)
				if clause.Expr == nil {
					r.expectedExpression()
				}

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.Delimiter && next.Value() == ")" {
					clause.RightParen = next
					p.consumeToken()
				} else {
					r.unexpectedSingleRuneToken(')')
				}
			} else {
				r.unexpectedToken(token.KeywordWhere)
			}
		} else {
			r.unexpectedSingleRuneToken('(')
		}
	}
	return
}

// parseOverClause parses over-clause as defined in:
// https://sqlite.org/syntax/over-clause.html
func (p *simpleParser) parseOverClause(r reporter) (clause *ast.OverClause) {
	clause = &ast.OverClause{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordOver {
		clause.Over = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			clause.WindowName = next
			p.consumeToken()
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			clause.LeftParen = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				clause.BaseWindowName = next
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordPartition {
				clause.Partition = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordBy {
					clause.By = next
					p.consumeToken()

					for {
						next, ok = p.lookahead(r)
						if !ok {
							return
						}
						if next.Value() == "," {
							p.consumeToken()
						}
						if next.Type() == token.KeywordOrder || next.Type() == token.KeywordRange || next.Type() == token.KeywordRows || next.Type() == token.KeywordGroups || next.Value() == ")" {
							break
						}
						expression := p.parseExpression(r)
						if expression != nil {
							clause.Expr = append(clause.Expr, expression)
						} else {
							break
						}
					}
				} else {
					r.unexpectedToken(token.KeywordBy)
				}
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordOrder {
				clause.Order = next
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.KeywordBy {
					clause.By = next
					p.consumeToken()

					for {
						next, ok = p.lookahead(r)
						if !ok {
							return
						}
						if next.Value() == "," {
							p.consumeToken()
						}
						if next.Type() == token.KeywordRange || next.Type() == token.KeywordRows || next.Type() == token.KeywordGroups || next.Value() == ")" {
							break
						}
						clause.OrderingTerm = append(clause.OrderingTerm, p.parseOrderingTerm(r))
					}
				} else {
					r.unexpectedToken(token.KeywordBy)
				}
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordRange || next.Type() == token.KeywordRows || next.Type() == token.KeywordGroups {
				clause.FrameSpec = p.parseFrameSpec(r)
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Delimiter && next.Value() == ")" {
				clause.RightParen = next
				p.consumeToken()
			}
		}
	} else {
		r.unexpectedToken(token.KeywordOver)
	}
	return
}

// parseRaiseFunction parses raise-function as defined in:
// https://sqlite.org/syntax/raise-function.html
func (p *simpleParser) parseRaiseFunction(r reporter) (stmt *ast.RaiseFunction) {
	stmt = &ast.RaiseFunction{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordRaise {
		stmt.Raise = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			stmt.LeftParen = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordIgnore {
				stmt.Ignore = next
				p.consumeToken()
				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Type() == token.Delimiter && next.Value() == ")" {
					stmt.RightParen = next
					p.consumeToken()
				} else {
					r.unexpectedSingleRuneToken(')')
				}
			}

			if next.Type() == token.KeywordRollback || next.Type() == token.KeywordAbort || next.Type() == token.KeywordFail {
				if next.Type() == token.KeywordRollback {
					stmt.Rollback = next
				} else if next.Type() == token.KeywordAbort {
					stmt.Abort = next
				} else if next.Type() == token.KeywordFail {
					stmt.Fail = next
				} else {
					r.unexpectedToken(token.KeywordRollback, token.KeywordAbort, token.KeywordFail)
				}
				p.consumeToken()

				next, ok = p.lookahead(r)
				if !ok {
					return
				}
				if next.Value() == "," {
					stmt.Comma = next
					p.consumeToken()

					next, ok = p.lookahead(r)
					if !ok {
						return
					}
					if next.Type() == token.Literal {
						stmt.ErrorMessage = next
						p.consumeToken()

						next, ok = p.lookahead(r)
						if !ok {
							return
						}
						if next.Type() == token.Delimiter && next.Value() == ")" {
							stmt.RightParen = next
							p.consumeToken()
						} else {
							r.unexpectedSingleRuneToken(')')
						}
					} else {
						r.unexpectedToken(token.Literal)
					}
				} else {
					r.unexpectedSingleRuneToken(',')
				}
			}
		} else {
			r.unexpectedToken(token.Delimiter)
		}
	} else {
		r.unexpectedToken(token.KeywordRaise)
		stmt = nil
	}
	return
}

// parseWithClauseBeginnerStmts parses all different statements that begin with a WithClause.
// This includes Delete, Insert, Update and Select stmts.
func (p *simpleParser) parseWithClauseBeginnerStmts(stmt *ast.SQLStmt, r reporter) {
	withClause := &ast.WithClause{}
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordWith {
		withClause = p.parseWithClause(r)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordDelete:
		p.parseDeleteStmts(stmt, withClause, r)
	case token.KeywordInsert:
		stmt.InsertStmt = p.parseInsertStmt(withClause, r)
	case token.KeywordUpdate:
		p.parseUpdateStmts(stmt, withClause, r)
	case token.KeywordSelect:
		stmt.SelectStmt = p.parseSelectStmt(withClause, r)
	}
}

// parseDropStmts parses the multiple variations of DROP stmt.
// The variations are DROP INDEX,TABLE,TRIGGER and VIEW.
func (p *simpleParser) parseDropStmts(stmt *ast.SQLStmt, r reporter) {
	dropToken, ok := p.lookahead(r)
	if !ok {
		return
	}
	if dropToken.Type() == token.KeywordDrop {
		p.consumeToken()
		next, ok := p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordIndex:
			stmt.DropIndexStmt = p.parseDropIndexStmt(dropToken, r)
		case token.KeywordTable:
			stmt.DropTableStmt = p.parseDropTableStmt(dropToken, r)
		case token.KeywordTrigger:
			stmt.DropTriggerStmt = p.parseDropTriggerStmt(dropToken, r)
		case token.KeywordView:
			stmt.DropViewStmt = p.parseDropViewStmt(dropToken, r)
		}
	} else {
		r.unexpectedToken(token.KeywordDrop)
	}
}
