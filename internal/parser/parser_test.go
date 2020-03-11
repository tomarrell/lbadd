package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func TestSingleStatementParse(t *testing.T) {
	inputs := []struct {
		Name  string
		Query string
		Stmt  *ast.SQLStmt
	}{
		{
			"alter rename table",
			"ALTER TABLE users RENAME TO admins",
			&ast.SQLStmt{
				AlterTableStmt: &ast.AlterTableStmt{
					Alter:        token.New(1, 1, 0, 5, token.KeywordAlter, "ALTER"),
					Table:        token.New(1, 7, 6, 5, token.KeywordTable, "TABLE"),
					TableName:    token.New(1, 13, 12, 5, token.Literal, "users"),
					Rename:       token.New(1, 19, 18, 6, token.KeywordRename, "RENAME"),
					To:           token.New(1, 26, 25, 2, token.KeywordTo, "TO"),
					NewTableName: token.New(1, 29, 28, 6, token.Literal, "admins"),
				},
			},
		},
		{
			"alter rename column",
			"ALTER TABLE users RENAME COLUMN name TO username",
			&ast.SQLStmt{
				AlterTableStmt: &ast.AlterTableStmt{
					Alter:         token.New(1, 1, 0, 5, token.KeywordAlter, "ALTER"),
					Table:         token.New(1, 7, 6, 5, token.KeywordTable, "TABLE"),
					TableName:     token.New(1, 13, 12, 5, token.Literal, "users"),
					Rename:        token.New(1, 19, 18, 6, token.KeywordRename, "RENAME"),
					Column:        token.New(1, 26, 25, 6, token.KeywordColumn, "COLUMN"),
					ColumnName:    token.New(1, 33, 32, 4, token.Literal, "name"),
					To:            token.New(1, 38, 37, 2, token.KeywordTo, "TO"),
					NewColumnName: token.New(1, 41, 40, 8, token.Literal, "username"),
				},
			},
		},
		{
			"alter rename column implicit",
			"ALTER TABLE users RENAME name TO username",
			&ast.SQLStmt{
				AlterTableStmt: &ast.AlterTableStmt{
					Alter:         token.New(1, 1, 0, 5, token.KeywordAlter, "ALTER"),
					Table:         token.New(1, 7, 6, 5, token.KeywordTable, "TABLE"),
					TableName:     token.New(1, 13, 12, 5, token.Literal, "users"),
					Rename:        token.New(1, 19, 18, 6, token.KeywordRename, "RENAME"),
					ColumnName:    token.New(1, 26, 25, 4, token.Literal, "name"),
					To:            token.New(1, 31, 30, 2, token.KeywordTo, "TO"),
					NewColumnName: token.New(1, 34, 33, 8, token.Literal, "username"),
				},
			},
		},
		{
			"alter add column with two constraints",
			"ALTER TABLE users ADD COLUMN foo VARCHAR(15) CONSTRAINT pk PRIMARY KEY AUTOINCREMENT CONSTRAINT nn NOT NULL",
			&ast.SQLStmt{
				AlterTableStmt: &ast.AlterTableStmt{
					Alter:     token.New(1, 1, 0, 5, token.KeywordAlter, "ALTER"),
					Table:     token.New(1, 7, 6, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 13, 12, 5, token.Literal, "users"),
					Add:       token.New(1, 19, 18, 3, token.KeywordAdd, "ADD"),
					Column:    token.New(1, 23, 22, 6, token.KeywordColumn, "COLUMN"),
					ColumnDef: &ast.ColumnDef{
						ColumnName: token.New(1, 30, 29, 3, token.Literal, "foo"),
						TypeName: &ast.TypeName{
							Name: []token.Token{
								token.New(1, 34, 33, 7, token.Literal, "VARCHAR"),
							},
							LeftParen: token.New(1, 41, 40, 1, token.Delimiter, "("),
							SignedNumber1: &ast.SignedNumber{
								NumericLiteral: token.New(1, 42, 41, 2, token.Literal, "15"),
							},
							RightParen: token.New(1, 44, 43, 1, token.Delimiter, ")"),
						},
						ColumnConstraint: []*ast.ColumnConstraint{
							{
								Constraint:    token.New(1, 46, 45, 10, token.KeywordConstraint, "CONSTRAINT"),
								Name:          token.New(1, 57, 56, 2, token.Literal, "pk"),
								Primary:       token.New(1, 60, 59, 7, token.KeywordPrimary, "PRIMARY"),
								Key:           token.New(1, 68, 67, 3, token.KeywordKey, "KEY"),
								Autoincrement: token.New(1, 72, 71, 13, token.KeywordAutoincrement, "AUTOINCREMENT"),
							},
							{
								Constraint: token.New(1, 86, 85, 10, token.KeywordConstraint, "CONSTRAINT"),
								Name:       token.New(1, 97, 96, 2, token.Literal, "nn"),
								Not:        token.New(1, 100, 99, 3, token.KeywordNot, "NOT"),
								Null:       token.New(1, 104, 103, 4, token.KeywordNull, "NULL"),
							},
						},
					},
				},
			},
		},
		{
			"alter add column implicit with two constraints",
			"ALTER TABLE users ADD foo VARCHAR(15) CONSTRAINT pk PRIMARY KEY AUTOINCREMENT CONSTRAINT nn NOT NULL",
			&ast.SQLStmt{
				AlterTableStmt: &ast.AlterTableStmt{
					Alter:     token.New(1, 1, 0, 5, token.KeywordAlter, "ALTER"),
					Table:     token.New(1, 7, 6, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 13, 12, 5, token.Literal, "users"),
					Add:       token.New(1, 19, 18, 3, token.KeywordAdd, "ADD"),
					ColumnDef: &ast.ColumnDef{
						ColumnName: token.New(1, 23, 22, 3, token.Literal, "foo"),
						TypeName: &ast.TypeName{
							Name: []token.Token{
								token.New(1, 27, 26, 7, token.Literal, "VARCHAR"),
							},
							LeftParen: token.New(1, 34, 33, 1, token.Delimiter, "("),
							SignedNumber1: &ast.SignedNumber{
								NumericLiteral: token.New(1, 35, 34, 2, token.Literal, "15"),
							},
							RightParen: token.New(1, 37, 36, 1, token.Delimiter, ")"),
						},
						ColumnConstraint: []*ast.ColumnConstraint{
							{
								Constraint:    token.New(1, 39, 38, 10, token.KeywordConstraint, "CONSTRAINT"),
								Name:          token.New(1, 50, 49, 2, token.Literal, "pk"),
								Primary:       token.New(1, 53, 52, 7, token.KeywordPrimary, "PRIMARY"),
								Key:           token.New(1, 61, 60, 3, token.KeywordKey, "KEY"),
								Autoincrement: token.New(1, 65, 64, 13, token.KeywordAutoincrement, "AUTOINCREMENT"),
							},
							{
								Constraint: token.New(1, 79, 78, 10, token.KeywordConstraint, "CONSTRAINT"),
								Name:       token.New(1, 90, 89, 2, token.Literal, "nn"),
								Not:        token.New(1, 93, 92, 3, token.KeywordNot, "NOT"),
								Null:       token.New(1, 97, 96, 4, token.KeywordNull, "NULL"),
							},
						},
					},
				},
			},
		},
		{
			"attach database",
			"ATTACH DATABASE myDb AS newDb",
			&ast.SQLStmt{
				AttachStmt: &ast.AttachStmt{
					Attach:   token.New(1, 1, 0, 6, token.KeywordAttach, "ATTACH"),
					Database: token.New(1, 8, 7, 8, token.KeywordDatabase, "DATABASE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 17, 16, 4, token.Literal, "myDb"),
					},
					As:         token.New(1, 22, 21, 2, token.KeywordAs, "AS"),
					SchemaName: token.New(1, 25, 24, 5, token.Literal, "newDb"),
				},
			},
		},
		{
			"attach schema",
			"ATTACH mySchema AS newSchema",
			&ast.SQLStmt{
				AttachStmt: &ast.AttachStmt{
					Attach: token.New(1, 1, 0, 6, token.KeywordAttach, "ATTACH"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 8, 7, 8, token.Literal, "mySchema"),
					},
					As:         token.New(1, 17, 16, 2, token.KeywordAs, "AS"),
					SchemaName: token.New(1, 20, 19, 9, token.Literal, "newSchema"),
				},
			},
		},
		{
			"DETACH with DATABASE",
			"DETACH DATABASE newDb",
			&ast.SQLStmt{
				DetachStmt: &ast.DetachStmt{
					Detach:     token.New(1, 1, 0, 6, token.KeywordDetach, "DETACH"),
					Database:   token.New(1, 8, 7, 8, token.KeywordDatabase, "DATABASE"),
					SchemaName: token.New(1, 17, 16, 5, token.Literal, "newDb"),
				},
			},
		},
		{
			"DETACH without DATABASE",
			"DETACH newSchema",
			&ast.SQLStmt{
				DetachStmt: &ast.DetachStmt{
					Detach:     token.New(1, 1, 0, 6, token.KeywordDetach, "DETACH"),
					SchemaName: token.New(1, 8, 7, 9, token.Literal, "newSchema"),
				},
			},
		},
		{
			"vacuum",
			"VACUUM",
			&ast.SQLStmt{
				VacuumStmt: &ast.VacuumStmt{
					Vacuum: token.New(1, 1, 0, 6, token.KeywordVacuum, "VACUUM"),
				},
			},
		},
		{
			"VACUUM with schema-name",
			"VACUUM mySchema",
			&ast.SQLStmt{
				VacuumStmt: &ast.VacuumStmt{
					Vacuum:     token.New(1, 1, 0, 6, token.KeywordVacuum, "VACUUM"),
					SchemaName: token.New(1, 8, 7, 8, token.Literal, "mySchema"),
				},
			},
		},
		{
			"VACUUM with INTO",
			"VACUUM INTO newFile",
			&ast.SQLStmt{
				VacuumStmt: &ast.VacuumStmt{
					Vacuum:   token.New(1, 1, 0, 6, token.KeywordVacuum, "VACUUM"),
					Into:     token.New(1, 8, 7, 4, token.KeywordInto, "INTO"),
					Filename: token.New(1, 13, 12, 7, token.Literal, "newFile"),
				},
			},
		},
		{
			"VACUUM with schema-name and INTO",
			"VACUUM mySchema INTO newFile",
			&ast.SQLStmt{
				VacuumStmt: &ast.VacuumStmt{
					Vacuum:     token.New(1, 1, 0, 6, token.KeywordVacuum, "VACUUM"),
					SchemaName: token.New(1, 8, 7, 8, token.Literal, "mySchema"),
					Into:       token.New(1, 17, 16, 4, token.KeywordInto, "INTO"),
					Filename:   token.New(1, 22, 21, 7, token.Literal, "newFile"),
				},
			},
		},
		{
			"analyze",
			"ANALYZE",
			&ast.SQLStmt{
				AnalyzeStmt: &ast.AnalyzeStmt{
					Analyze: token.New(1, 1, 0, 7, token.KeywordAnalyze, "ANALYZE"),
				},
			},
		},
		{
			"ANALYZE with schema-name/table-or-index-name",
			"ANALYZE mySchemaOrTableOrIndex",
			&ast.SQLStmt{
				AnalyzeStmt: &ast.AnalyzeStmt{
					Analyze:          token.New(1, 1, 0, 7, token.KeywordAnalyze, "ANALYZE"),
					SchemaName:       token.New(1, 9, 8, 22, token.Literal, "mySchemaOrTableOrIndex"),
					TableOrIndexName: token.New(1, 9, 8, 22, token.Literal, "mySchemaOrTableOrIndex"),
				},
			},
		},
		{
			"ANALYZE with schema-name/table-or-index-name elaborated",
			"ANALYZE mySchemaOrTableOrIndex.specificAttr",
			&ast.SQLStmt{
				AnalyzeStmt: &ast.AnalyzeStmt{
					Analyze:          token.New(1, 1, 0, 7, token.KeywordAnalyze, "ANALYZE"),
					SchemaName:       token.New(1, 9, 8, 22, token.Literal, "mySchemaOrTableOrIndex"),
					Period:           token.New(1, 31, 30, 1, token.Literal, "."),
					TableOrIndexName: token.New(1, 32, 31, 12, token.Literal, "specificAttr"),
				},
			},
		},
		{
			"begin",
			"BEGIN",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin: token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
				},
			},
		},
		{
			"BEGIN with DEFERRED",
			"BEGIN DEFERRED",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:    token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Deferred: token.New(1, 7, 6, 8, token.KeywordDeferred, "DEFERRED"),
				},
			},
		},
		{
			"BEGIN with IMMEDIATE",
			"BEGIN IMMEDIATE",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:     token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Immediate: token.New(1, 7, 6, 9, token.KeywordImmediate, "IMMEDIATE"),
				},
			},
		},
		{
			"BEGIN with EXCLUSIVE",
			"BEGIN EXCLUSIVE",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:     token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Exclusive: token.New(1, 7, 6, 9, token.KeywordExclusive, "EXCLUSIVE"),
				},
			},
		},
		{
			"BEGIN with TRANSACTION",
			"BEGIN TRANSACTION",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:       token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Transaction: token.New(1, 7, 6, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"BEGIN with DEFERRED and TRANSACTION",
			"BEGIN DEFERRED TRANSACTION",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:       token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Deferred:    token.New(1, 7, 6, 8, token.KeywordDeferred, "DEFERRED"),
					Transaction: token.New(1, 16, 15, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"BEGIN with IMMEDIATE and TRANSACTION",
			"BEGIN IMMEDIATE TRANSACTION",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:       token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Immediate:   token.New(1, 7, 6, 9, token.KeywordImmediate, "IMMEDIATE"),
					Transaction: token.New(1, 17, 16, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"BEGIN with EXCLUSIVE and TRANSACTION",
			"BEGIN EXCLUSIVE TRANSACTION",
			&ast.SQLStmt{
				BeginStmt: &ast.BeginStmt{
					Begin:       token.New(1, 1, 0, 5, token.KeywordBegin, "BEGIN"),
					Exclusive:   token.New(1, 7, 6, 9, token.KeywordExclusive, "EXCLUSIVE"),
					Transaction: token.New(1, 17, 16, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"commit",
			"COMMIT",
			&ast.SQLStmt{
				CommitStmt: &ast.CommitStmt{
					Commit: token.New(1, 1, 0, 6, token.KeywordCommit, "COMMIT"),
				},
			},
		},
		{
			"end",
			"END",
			&ast.SQLStmt{
				CommitStmt: &ast.CommitStmt{
					End: token.New(1, 1, 0, 3, token.KeywordEnd, "END"),
				},
			},
		},
		{
			"COMMIT with TRANSACTION",
			"COMMIT TRANSACTION",
			&ast.SQLStmt{
				CommitStmt: &ast.CommitStmt{
					Commit:      token.New(1, 1, 0, 6, token.KeywordCommit, "COMMIT"),
					Transaction: token.New(1, 8, 7, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"END with TRANSACTION",
			"END TRANSACTION",
			&ast.SQLStmt{
				CommitStmt: &ast.CommitStmt{
					End:         token.New(1, 1, 0, 3, token.KeywordEnd, "END"),
					Transaction: token.New(1, 5, 4, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"rollback",
			"ROLLBACK",
			&ast.SQLStmt{
				RollbackStmt: &ast.RollbackStmt{
					Rollback: token.New(1, 1, 0, 8, token.KeywordRollback, "ROLLBACK"),
				},
			},
		},
		{
			"ROLLBACK with TRANSACTION",
			"ROLLBACK TRANSACTION",
			&ast.SQLStmt{
				RollbackStmt: &ast.RollbackStmt{
					Rollback:    token.New(1, 1, 0, 8, token.KeywordRollback, "ROLLBACK"),
					Transaction: token.New(1, 10, 9, 11, token.KeywordTransaction, "TRANSACTION"),
				},
			},
		},
		{
			"ROLLBACK with TRANSACTION and TO",
			"ROLLBACK TRANSACTION TO mySavePoint",
			&ast.SQLStmt{
				RollbackStmt: &ast.RollbackStmt{
					Rollback:      token.New(1, 1, 0, 8, token.KeywordRollback, "ROLLBACK"),
					Transaction:   token.New(1, 10, 9, 11, token.KeywordTransaction, "TRANSACTION"),
					To:            token.New(1, 22, 21, 2, token.KeywordTo, "TO"),
					SavepointName: token.New(1, 25, 24, 11, token.Literal, "mySavePoint"),
				},
			},
		},
		{
			"ROLLBACK with TRANSACTION, TO and SAVEPOINT",
			"ROLLBACK TRANSACTION TO SAVEPOINT mySavePoint",
			&ast.SQLStmt{
				RollbackStmt: &ast.RollbackStmt{
					Rollback:      token.New(1, 1, 0, 8, token.KeywordRollback, "ROLLBACK"),
					Transaction:   token.New(1, 10, 9, 11, token.KeywordTransaction, "TRANSACTION"),
					To:            token.New(1, 22, 21, 2, token.KeywordTo, "TO"),
					Savepoint:     token.New(1, 25, 24, 9, token.KeywordSavepoint, "SAVEPOINT"),
					SavepointName: token.New(1, 35, 34, 11, token.Literal, "mySavePoint"),
				},
			},
		},
		{
			"ROLLBACK with TO",
			"ROLLBACK TO mySavePoint",
			&ast.SQLStmt{
				RollbackStmt: &ast.RollbackStmt{
					Rollback:      token.New(1, 1, 0, 8, token.KeywordRollback, "ROLLBACK"),
					To:            token.New(1, 10, 9, 2, token.KeywordTo, "TO"),
					SavepointName: token.New(1, 13, 12, 11, token.Literal, "mySavePoint"),
				},
			},
		},
		{
			"ROLLBACK with TO and SAVEPOINT",
			"ROLLBACK TO SAVEPOINT mySavePoint",
			&ast.SQLStmt{
				RollbackStmt: &ast.RollbackStmt{
					Rollback:      token.New(1, 1, 0, 8, token.KeywordRollback, "ROLLBACK"),
					To:            token.New(1, 10, 9, 2, token.KeywordTo, "TO"),
					Savepoint:     token.New(1, 13, 12, 9, token.KeywordSavepoint, "SAVEPOINT"),
					SavepointName: token.New(1, 23, 22, 11, token.Literal, "mySavePoint"),
				},
			},
		},
	}
	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			assert := assert.New(t)

			p := New(input.Query)

			stmt, errs, ok := p.Next()
			assert.True(ok, "expected exactly one statement")
			for _, err := range errs {
				assert.Nil(err)
			}

			opts := []cmp.Option{
				cmp.Comparer(func(t1, t2 token.Token) bool {
					if t1 == t2 {
						return true
					}
					if (t1 == nil && t2 != nil) ||
						(t1 != nil && t2 == nil) {
						return false
					}
					return t1.Line() == t2.Line() &&
						t1.Col() == t2.Col() &&
						t1.Offset() == t2.Offset() &&
						t1.Length() == t2.Length() &&
						t1.Type() == t2.Type() &&
						t1.Value() == t2.Value()
				}),
			}
			t.Log(cmp.Diff(input.Stmt, stmt, opts...))
			assert.True(cmp.Equal(input.Stmt, stmt, opts...))

			_, _, ok = p.Next()
			assert.False(ok, "expected only one statement")
		})
	}
}
