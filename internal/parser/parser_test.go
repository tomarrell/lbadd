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
								NumericLiteral: token.New(1, 42, 41, 2, token.LiteralNumeric, "15"),
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
								NumericLiteral: token.New(1, 35, 34, 2, token.LiteralNumeric, "15"),
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
		{
			"create index",
			"CREATE INDEX myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:     token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					IndexName: token.New(1, 14, 13, 7, token.Literal, "myIndex"),
					On:        token.New(1, 22, 21, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 25, 24, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 33, 32, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 34, 33, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 45, 44, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE",
			"CREATE UNIQUE INDEX myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:    token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:     token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					IndexName: token.New(1, 21, 20, 7, token.Literal, "myIndex"),
					On:        token.New(1, 29, 28, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 32, 31, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 40, 39, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 41, 40, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with IF NOT EXISTS",
			"CREATE INDEX IF NOT EXISTS myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:     token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					If:        token.New(1, 14, 13, 2, token.KeywordIf, "IF"),
					Not:       token.New(1, 17, 16, 3, token.KeywordNot, "NOT"),
					Exists:    token.New(1, 21, 20, 6, token.KeywordExists, "EXISTS"),
					IndexName: token.New(1, 28, 27, 7, token.Literal, "myIndex"),
					On:        token.New(1, 36, 35, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 39, 38, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 47, 46, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 48, 47, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE and IF NOT EXISTS",
			"CREATE UNIQUE INDEX IF NOT EXISTS myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:    token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:     token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					If:        token.New(1, 21, 20, 2, token.KeywordIf, "IF"),
					Not:       token.New(1, 24, 23, 3, token.KeywordNot, "NOT"),
					Exists:    token.New(1, 28, 27, 6, token.KeywordExists, "EXISTS"),
					IndexName: token.New(1, 35, 34, 7, token.Literal, "myIndex"),
					On:        token.New(1, 43, 42, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 46, 45, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 54, 53, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 55, 54, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 66, 65, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"create index with schema and index name",
			"CREATE INDEX mySchema.myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:      token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					SchemaName: token.New(1, 14, 13, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 22, 21, 1, token.Literal, "."),
					IndexName:  token.New(1, 23, 22, 7, token.Literal, "myIndex"),
					On:         token.New(1, 31, 30, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 34, 33, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 42, 41, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 43, 42, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 54, 53, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE with schema and index name",
			"CREATE UNIQUE INDEX mySchema.myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:     token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:      token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					SchemaName: token.New(1, 21, 20, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 29, 28, 1, token.Literal, "."),
					IndexName:  token.New(1, 30, 29, 7, token.Literal, "myIndex"),
					On:         token.New(1, 38, 37, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 41, 40, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 49, 48, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 50, 49, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 61, 60, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with IF NOT EXISTS with schema and index name",
			"CREATE INDEX IF NOT EXISTS mySchema.myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:      token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					If:         token.New(1, 14, 13, 2, token.KeywordIf, "IF"),
					Not:        token.New(1, 17, 16, 3, token.KeywordNot, "NOT"),
					Exists:     token.New(1, 21, 20, 6, token.KeywordExists, "EXISTS"),
					SchemaName: token.New(1, 28, 27, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 36, 35, 1, token.Literal, "."),
					IndexName:  token.New(1, 37, 36, 7, token.Literal, "myIndex"),
					On:         token.New(1, 45, 44, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 48, 47, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 56, 55, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 57, 56, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 68, 67, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE and IF NOT EXISTS with schema and index name",
			"CREATE UNIQUE INDEX IF NOT EXISTS mySchema.myIndex ON myTable (exprLiteral)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:     token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:      token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					If:         token.New(1, 21, 20, 2, token.KeywordIf, "IF"),
					Not:        token.New(1, 24, 23, 3, token.KeywordNot, "NOT"),
					Exists:     token.New(1, 28, 27, 6, token.KeywordExists, "EXISTS"),
					SchemaName: token.New(1, 35, 34, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 43, 42, 1, token.Literal, "."),
					IndexName:  token.New(1, 44, 43, 7, token.Literal, "myIndex"),
					On:         token.New(1, 52, 51, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 55, 54, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 63, 62, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 64, 63, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 75, 74, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with WHERE",
			"CREATE INDEX myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:     token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					IndexName: token.New(1, 14, 13, 7, token.Literal, "myIndex"),
					On:        token.New(1, 22, 21, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 25, 24, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 33, 32, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 34, 33, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 45, 44, 1, token.Delimiter, ")"),
					Where:      token.New(1, 47, 46, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 53, 52, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE and WHERE",
			"CREATE UNIQUE INDEX myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:    token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:     token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					IndexName: token.New(1, 21, 20, 7, token.Literal, "myIndex"),
					On:        token.New(1, 29, 28, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 32, 31, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 40, 39, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 41, 40, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
					Where:      token.New(1, 54, 53, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 60, 59, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with IF NOT EXISTS and WHERE",
			"CREATE INDEX IF NOT EXISTS myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:     token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					If:        token.New(1, 14, 13, 2, token.KeywordIf, "IF"),
					Not:       token.New(1, 17, 16, 3, token.KeywordNot, "NOT"),
					Exists:    token.New(1, 21, 20, 6, token.KeywordExists, "EXISTS"),
					IndexName: token.New(1, 28, 27, 7, token.Literal, "myIndex"),
					On:        token.New(1, 36, 35, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 39, 38, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 47, 46, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 48, 47, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
					Where:      token.New(1, 61, 60, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 67, 66, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE, IF NOT EXISTS and WHERE",
			"CREATE UNIQUE INDEX IF NOT EXISTS myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:    token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:     token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					If:        token.New(1, 21, 20, 2, token.KeywordIf, "IF"),
					Not:       token.New(1, 24, 23, 3, token.KeywordNot, "NOT"),
					Exists:    token.New(1, 28, 27, 6, token.KeywordExists, "EXISTS"),
					IndexName: token.New(1, 35, 34, 7, token.Literal, "myIndex"),
					On:        token.New(1, 43, 42, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 46, 45, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 54, 53, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 55, 54, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 66, 65, 1, token.Delimiter, ")"),
					Where:      token.New(1, 68, 67, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 74, 73, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"create index with schema and index name and WHERE",
			"CREATE INDEX mySchema.myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:      token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					SchemaName: token.New(1, 14, 13, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 22, 21, 1, token.Literal, "."),
					IndexName:  token.New(1, 23, 22, 7, token.Literal, "myIndex"),
					On:         token.New(1, 31, 30, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 34, 33, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 42, 41, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 43, 42, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 54, 53, 1, token.Delimiter, ")"),
					Where:      token.New(1, 56, 55, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 62, 61, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE, schema name, index name and WHERE",
			"CREATE UNIQUE INDEX mySchema.myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:     token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:      token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					SchemaName: token.New(1, 21, 20, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 29, 28, 1, token.Literal, "."),
					IndexName:  token.New(1, 30, 29, 7, token.Literal, "myIndex"),
					On:         token.New(1, 38, 37, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 41, 40, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 49, 48, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 50, 49, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 61, 60, 1, token.Delimiter, ")"),
					Where:      token.New(1, 63, 62, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 69, 68, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with IF NOT EXISTS,schema name, index name and WHERE",
			"CREATE INDEX IF NOT EXISTS mySchema.myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:      token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					If:         token.New(1, 14, 13, 2, token.KeywordIf, "IF"),
					Not:        token.New(1, 17, 16, 3, token.KeywordNot, "NOT"),
					Exists:     token.New(1, 21, 20, 6, token.KeywordExists, "EXISTS"),
					SchemaName: token.New(1, 28, 27, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 36, 35, 1, token.Literal, "."),
					IndexName:  token.New(1, 37, 36, 7, token.Literal, "myIndex"),
					On:         token.New(1, 45, 44, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 48, 47, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 56, 55, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 57, 56, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 68, 67, 1, token.Delimiter, ")"),
					Where:      token.New(1, 70, 69, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 76, 75, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE, IF NOT EXISTS, schema name, index name and WHERE",
			"CREATE UNIQUE INDEX IF NOT EXISTS mySchema.myIndex ON myTable (exprLiteral) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:     token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:      token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					If:         token.New(1, 21, 20, 2, token.KeywordIf, "IF"),
					Not:        token.New(1, 24, 23, 3, token.KeywordNot, "NOT"),
					Exists:     token.New(1, 28, 27, 6, token.KeywordExists, "EXISTS"),
					SchemaName: token.New(1, 35, 34, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 43, 42, 1, token.Literal, "."),
					IndexName:  token.New(1, 44, 43, 7, token.Literal, "myIndex"),
					On:         token.New(1, 52, 51, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 55, 54, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 63, 62, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 64, 63, 11, token.Literal, "exprLiteral"),
						},
					},
					RightParen: token.New(1, 75, 74, 1, token.Delimiter, ")"),
					Where:      token.New(1, 77, 76, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 83, 82, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with UNIQUE, IF NOT EXISTS, schema name, index name and WHERE with multiple indexedcolums",
			"CREATE UNIQUE INDEX IF NOT EXISTS mySchema.myIndex ON myTable (exprLiteral1,exprLiteral2,exprLiteral3) WHERE exprLiteral",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Unique:     token.New(1, 8, 7, 6, token.KeywordUnique, "UNIQUE"),
					Index:      token.New(1, 15, 14, 5, token.KeywordIndex, "INDEX"),
					If:         token.New(1, 21, 20, 2, token.KeywordIf, "IF"),
					Not:        token.New(1, 24, 23, 3, token.KeywordNot, "NOT"),
					Exists:     token.New(1, 28, 27, 6, token.KeywordExists, "EXISTS"),
					SchemaName: token.New(1, 35, 34, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 43, 42, 1, token.Literal, "."),
					IndexName:  token.New(1, 44, 43, 7, token.Literal, "myIndex"),
					On:         token.New(1, 52, 51, 2, token.KeywordOn, "ON"),
					TableName:  token.New(1, 55, 54, 7, token.Literal, "myTable"),
					LeftParen:  token.New(1, 63, 62, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 64, 63, 12, token.Literal, "exprLiteral1"),
						},
						&ast.IndexedColumn{
							ColumnName: token.New(1, 77, 76, 12, token.Literal, "exprLiteral2"),
						},
						&ast.IndexedColumn{
							ColumnName: token.New(1, 90, 89, 12, token.Literal, "exprLiteral3"),
						},
					},
					RightParen: token.New(1, 102, 101, 1, token.Delimiter, ")"),
					Where:      token.New(1, 104, 103, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 110, 109, 11, token.Literal, "exprLiteral"),
					},
				},
			},
		},
		{
			"CREATE INDEX with full fledged indexed columns and DESC",
			"CREATE INDEX myIndex ON myTable (exprLiteral COLLATE myCollation DESC)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:     token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					IndexName: token.New(1, 14, 13, 7, token.Literal, "myIndex"),
					On:        token.New(1, 22, 21, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 25, 24, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 33, 32, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName:    token.New(1, 34, 33, 11, token.Literal, "exprLiteral"),
							Collate:       token.New(1, 46, 45, 7, token.KeywordCollate, "COLLATE"),
							CollationName: token.New(1, 54, 53, 11, token.Literal, "myCollation"),
							Desc:          token.New(1, 66, 65, 4, token.KeywordDesc, "DESC"),
						},
					},
					RightParen: token.New(1, 70, 69, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"CREATE INDEX with indexed columns and ASC",
			"CREATE INDEX myIndex ON myTable (exprLiteral ASC)",
			&ast.SQLStmt{
				CreateIndexStmt: &ast.CreateIndexStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Index:     token.New(1, 8, 7, 5, token.KeywordIndex, "INDEX"),
					IndexName: token.New(1, 14, 13, 7, token.Literal, "myIndex"),
					On:        token.New(1, 22, 21, 2, token.KeywordOn, "ON"),
					TableName: token.New(1, 25, 24, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 33, 32, 1, token.Delimiter, "("),
					IndexedColumns: []*ast.IndexedColumn{
						&ast.IndexedColumn{
							ColumnName: token.New(1, 34, 33, 11, token.Literal, "exprLiteral"),
							Asc:        token.New(1, 46, 45, 3, token.KeywordAsc, "ASC"),
						},
					},
					RightParen: token.New(1, 49, 48, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			"DELETE basic",
			"DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					Delete: token.New(1, 1, 0, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 13, 12, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with WHERE and basic qualified table name",
			"DELETE FROM myTable WHERE myLiteral",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					Delete: token.New(1, 1, 0, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 13, 12, 7, token.Literal, "myTable"),
					},
					Where: token.New(1, 21, 20, 5, token.KeywordWhere, "WHERE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 27, 26, 9, token.Literal, "myLiteral"),
					},
				},
			},
		},
		{
			"DELETE with schema name and table name",
			"DELETE FROM mySchema.myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					Delete: token.New(1, 1, 0, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						SchemaName: token.New(1, 13, 12, 8, token.Literal, "mySchema"),
						Period:     token.New(1, 21, 20, 1, token.Literal, "."),
						TableName:  token.New(1, 22, 21, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with schema name, table name and AS",
			"DELETE FROM mySchema.myTable AS newSchemaTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					Delete: token.New(1, 1, 0, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						SchemaName: token.New(1, 13, 12, 8, token.Literal, "mySchema"),
						Period:     token.New(1, 21, 20, 1, token.Literal, "."),
						TableName:  token.New(1, 22, 21, 7, token.Literal, "myTable"),
						As:         token.New(1, 30, 29, 2, token.KeywordAs, "AS"),
						Alias:      token.New(1, 33, 32, 14, token.Literal, "newSchemaTable"),
					},
				},
			},
		},
		{
			"DELETE with schema name, table name, AS and INDEXED BY",
			"DELETE FROM mySchema.myTable AS newSchemaTable INDEXED BY myIndex",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					Delete: token.New(1, 1, 0, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						SchemaName: token.New(1, 13, 12, 8, token.Literal, "mySchema"),
						Period:     token.New(1, 21, 20, 1, token.Literal, "."),
						TableName:  token.New(1, 22, 21, 7, token.Literal, "myTable"),
						As:         token.New(1, 30, 29, 2, token.KeywordAs, "AS"),
						Alias:      token.New(1, 33, 32, 14, token.Literal, "newSchemaTable"),
						Indexed:    token.New(1, 48, 47, 7, token.KeywordIndexed, "INDEXED"),
						By:         token.New(1, 56, 55, 2, token.KeywordBy, "BY"),
						IndexName:  token.New(1, 59, 58, 7, token.Literal, "myIndex"),
					},
				},
			},
		},
		{
			"DELETE with schema name, table name and NOT INDEXED",
			"DELETE FROM mySchema.myTable NOT INDEXED",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					Delete: token.New(1, 1, 0, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 8, 7, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						SchemaName: token.New(1, 13, 12, 8, token.Literal, "mySchema"),
						Period:     token.New(1, 21, 20, 1, token.Literal, "."),
						TableName:  token.New(1, 22, 21, 7, token.Literal, "myTable"),
						Not:        token.New(1, 30, 29, 3, token.KeywordNot, "NOT"),
						Indexed:    token.New(1, 34, 33, 7, token.KeywordIndexed, "INDEXED"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic select stmt and basic cte-table-name",
			"WITH myTable AS (SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 26, 25, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 28, 27, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 35, 34, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 40, 39, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			`DELETE with "with clause" with RECURSIVE, basic select stmt and basic cte-table-name`,
			"WITH RECURSIVE myTable AS (SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With:      token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						Recursive: token.New(1, 6, 5, 9, token.KeywordRecursive, "RECURSIVE"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 16, 15, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 24, 23, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 27, 26, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 28, 27, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 35, 34, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 36, 35, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 38, 37, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 45, 44, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 50, 49, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			`DELETE with "with clause" with RECURSIVE, basic select stmt and cte-table-name with single col`,
			"WITH RECURSIVE myTable (myCol) AS (SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With:      token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						Recursive: token.New(1, 6, 5, 9, token.KeywordRecursive, "RECURSIVE"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 16, 15, 7, token.Literal, "myTable"),
									LeftParen: token.New(1, 24, 23, 1, token.Delimiter, "("),
									ColumnName: []token.Token{
										token.New(1, 25, 24, 5, token.Literal, "myCol"),
									},
									RightParen: token.New(1, 30, 29, 1, token.Delimiter, ")"),
								},
								As:        token.New(1, 32, 31, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 35, 34, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 36, 35, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 43, 42, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 44, 43, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 46, 45, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 53, 52, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 58, 57, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			`DELETE with "with clause" with RECURSIVE, basic select stmt and cte-table-name with multiple cols`,
			"WITH RECURSIVE myTable (myCol1,myCol2) AS (SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With:      token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						Recursive: token.New(1, 6, 5, 9, token.KeywordRecursive, "RECURSIVE"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 16, 15, 7, token.Literal, "myTable"),
									LeftParen: token.New(1, 24, 23, 1, token.Delimiter, "("),
									ColumnName: []token.Token{
										token.New(1, 25, 24, 6, token.Literal, "myCol1"),
										token.New(1, 32, 31, 6, token.Literal, "myCol2"),
									},
									RightParen: token.New(1, 38, 37, 1, token.Delimiter, ")"),
								},
								As:        token.New(1, 40, 39, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 43, 42, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 44, 43, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 51, 50, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 54, 53, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 61, 60, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 66, 65, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause,select stmt with WITH, basic common table expression and basic cte-table-name",
			"WITH myTable AS (WITH myTable AS (SELECT *) SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									With: token.New(1, 18, 17, 4, token.KeywordWith, "WITH"),
									CommonTableExpression: []*ast.CommonTableExpression{
										&ast.CommonTableExpression{
											TableName:  token.New(1, 23, 22, 7, token.Literal, "myTable"),
											As:         token.New(1, 31, 30, 2, token.KeywordAs, "AS"),
											LeftParen2: token.New(1, 34, 33, 1, token.Delimiter, "("),
											SelectStmt: &ast.SelectStmt{
												SelectCore: []*ast.SelectCore{
													&ast.SelectCore{
														Select: token.New(1, 35, 34, 6, token.KeywordSelect, "SELECT"),
														ResultColumn: []*ast.ResultColumn{
															&ast.ResultColumn{
																Asterisk: token.New(1, 42, 41, 1, token.BinaryOperator, "*"),
															},
														},
													},
												},
											},
											RightParen2: token.New(1, 43, 42, 1, token.Delimiter, ")"),
										},
									},
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 45, 44, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 52, 51, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 53, 52, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 55, 54, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 62, 61, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 67, 66, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause,select stmt with WITH, common table expression with single col and basic cte-table-name",
			"WITH myTable AS (WITH myTable (myCol) AS (SELECT *) SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									With: token.New(1, 18, 17, 4, token.KeywordWith, "WITH"),
									CommonTableExpression: []*ast.CommonTableExpression{
										&ast.CommonTableExpression{
											TableName:  token.New(1, 23, 22, 7, token.Literal, "myTable"),
											LeftParen1: token.New(1, 31, 30, 1, token.Delimiter, "("),
											ColumnName: []token.Token{
												token.New(1, 32, 31, 5, token.Literal, "myCol"),
											},
											RightParen1: token.New(1, 37, 36, 1, token.Delimiter, ")"),
											As:          token.New(1, 39, 38, 2, token.KeywordAs, "AS"),
											LeftParen2:  token.New(1, 42, 41, 1, token.Delimiter, "("),
											SelectStmt: &ast.SelectStmt{
												SelectCore: []*ast.SelectCore{
													&ast.SelectCore{
														Select: token.New(1, 43, 42, 6, token.KeywordSelect, "SELECT"),
														ResultColumn: []*ast.ResultColumn{
															&ast.ResultColumn{
																Asterisk: token.New(1, 50, 49, 1, token.BinaryOperator, "*"),
															},
														},
													},
												},
											},
											RightParen2: token.New(1, 51, 50, 1, token.Delimiter, ")"),
										},
									},
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 53, 52, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 60, 59, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 61, 60, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 63, 62, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 70, 69, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 75, 74, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause,select stmt with WITH and RECURSIVE, basic common table expression and basic cte-table-name",
			"WITH myTable AS (WITH RECURSIVE myTable AS (SELECT *) SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									With:      token.New(1, 18, 17, 4, token.KeywordWith, "WITH"),
									Recursive: token.New(1, 23, 22, 9, token.KeywordRecursive, "RECURSIVE"),
									CommonTableExpression: []*ast.CommonTableExpression{
										&ast.CommonTableExpression{
											TableName:  token.New(1, 33, 32, 7, token.Literal, "myTable"),
											As:         token.New(1, 41, 40, 2, token.KeywordAs, "AS"),
											LeftParen2: token.New(1, 44, 43, 1, token.Delimiter, "("),
											SelectStmt: &ast.SelectStmt{
												SelectCore: []*ast.SelectCore{
													&ast.SelectCore{
														Select: token.New(1, 45, 44, 6, token.KeywordSelect, "SELECT"),
														ResultColumn: []*ast.ResultColumn{
															&ast.ResultColumn{
																Asterisk: token.New(1, 52, 51, 1, token.BinaryOperator, "*"),
															},
														},
													},
												},
											},
											RightParen2: token.New(1, 53, 52, 1, token.Delimiter, ")"),
										},
									},
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 55, 54, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 62, 61, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 63, 62, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 65, 64, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 72, 71, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 77, 76, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause,select stmt with WITH, common table expression with multiple cols and basic cte-table-name",
			"WITH myTable AS (WITH myTable (myCol1,myCol2) AS (SELECT *) SELECT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									With: token.New(1, 18, 17, 4, token.KeywordWith, "WITH"),
									CommonTableExpression: []*ast.CommonTableExpression{
										&ast.CommonTableExpression{
											TableName:  token.New(1, 23, 22, 7, token.Literal, "myTable"),
											LeftParen1: token.New(1, 31, 30, 1, token.Delimiter, "("),
											ColumnName: []token.Token{
												token.New(1, 32, 31, 6, token.Literal, "myCol1"),
												token.New(1, 39, 38, 6, token.Literal, "myCol2"),
											},
											RightParen1: token.New(1, 45, 44, 1, token.Delimiter, ")"),
											As:          token.New(1, 47, 46, 2, token.KeywordAs, "AS"),
											LeftParen2:  token.New(1, 50, 49, 1, token.Delimiter, "("),
											SelectStmt: &ast.SelectStmt{
												SelectCore: []*ast.SelectCore{
													&ast.SelectCore{
														Select: token.New(1, 51, 50, 6, token.KeywordSelect, "SELECT"),
														ResultColumn: []*ast.ResultColumn{
															&ast.ResultColumn{
																Asterisk: token.New(1, 58, 57, 1, token.BinaryOperator, "*"),
															},
														},
													},
												},
											},
											RightParen2: token.New(1, 59, 58, 1, token.Delimiter, ")"),
										},
									},
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 61, 60, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 68, 67, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 69, 68, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 71, 70, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 78, 77, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 83, 82, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with DISTINCT and basic cte-table-name",
			"WITH myTable AS (SELECT DISTINCT *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select:   token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											Distinct: token.New(1, 25, 24, 8, token.KeywordDistinct, "DISTINCT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 34, 33, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 35, 34, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 37, 36, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 44, 43, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 49, 48, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with ALL and basic cte-table-name",
			"WITH myTable AS (SELECT ALL *) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											All:    token.New(1, 25, 24, 3, token.KeywordAll, "ALL"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 29, 28, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 30, 29, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 32, 31, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 39, 38, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 44, 43, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt's result column with table name and basic cte-table-name",
			"WITH myTable AS (SELECT myTable.*) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													TableName: token.New(1, 25, 24, 7, token.Literal, "myTable"),
													Period:    token.New(1, 32, 31, 1, token.Literal, "."),
													Asterisk:  token.New(1, 33, 32, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 34, 33, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 36, 35, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 43, 42, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 48, 47, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt's result column with expr and basic cte-table-name",
			"WITH myTable AS (SELECT myExpr) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Expr: &ast.Expr{
														LiteralValue: token.New(1, 25, 24, 6, token.Literal, "myExpr"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 31, 30, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 33, 32, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 40, 39, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 45, 44, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt's result column with expr with column-alias and basic cte-table-name",
			"WITH myTable AS (SELECT myExpr myColAlias) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Expr: &ast.Expr{
														LiteralValue: token.New(1, 25, 24, 6, token.Literal, "myExpr"),
													},
													ColumnAlias: token.New(1, 32, 31, 10, token.Literal, "myColAlias"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 42, 41, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 44, 43, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 51, 50, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 56, 55, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt's result column with expr with column-alias and AS and basic cte-table-name",
			"WITH myTable AS (SELECT myExpr AS myColAlias) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Expr: &ast.Expr{
														LiteralValue: token.New(1, 25, 24, 6, token.Literal, "myExpr"),
													},
													As:          token.New(1, 32, 31, 2, token.KeywordAs, "AS"),
													ColumnAlias: token.New(1, 35, 34, 10, token.Literal, "myColAlias"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 45, 44, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 47, 46, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 54, 53, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 59, 58, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with basic joinclause and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 7, token.Literal, "myTable"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 39, 38, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 41, 40, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 48, 47, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 53, 52, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1,myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Comma: token.New(1, 40, 39, 1, token.Delimiter, ","),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 41, 40, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 49, 48, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 51, 50, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 58, 57, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 63, 62, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's ON and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1,myTable2 ON myExpr) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Comma: token.New(1, 40, 39, 1, token.Delimiter, ","),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 41, 40, 8, token.Literal, "myTable2"),
													},
													JoinConstraint: &ast.JoinConstraint{
														On: token.New(1, 50, 49, 2, token.KeywordOn, "ON"),
														Expr: &ast.Expr{
															LiteralValue: token.New(1, 53, 52, 6, token.Literal, "myExpr"),
														},
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 61, 60, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 68, 67, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 73, 72, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's USING and single Col and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1,myTable2 USING (myCol)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Comma: token.New(1, 40, 39, 1, token.Delimiter, ","),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 41, 40, 8, token.Literal, "myTable2"),
													},
													JoinConstraint: &ast.JoinConstraint{
														Using:     token.New(1, 50, 49, 5, token.KeywordUsing, "USING"),
														LeftParen: token.New(1, 56, 55, 1, token.Delimiter, "("),
														ColumnName: []token.Token{
															token.New(1, 57, 56, 5, token.Literal, "myCol"),
														},
														RightParen: token.New(1, 62, 61, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 63, 62, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 65, 64, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 72, 71, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 77, 76, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's USING and multiple Cols and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1,myTable2 USING (myCol1,myCol2)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Comma: token.New(1, 40, 39, 1, token.Delimiter, ","),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 41, 40, 8, token.Literal, "myTable2"),
													},
													JoinConstraint: &ast.JoinConstraint{
														Using:     token.New(1, 50, 49, 5, token.KeywordUsing, "USING"),
														LeftParen: token.New(1, 56, 55, 1, token.Delimiter, "("),
														ColumnName: []token.Token{
															token.New(1, 57, 56, 6, token.Literal, "myCol1"),
															token.New(1, 64, 63, 6, token.Literal, "myCol2"),
														},
														RightParen: token.New(1, 70, 69, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 71, 70, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 73, 72, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 80, 79, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 85, 84, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint and JOIN and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1 JOIN myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Join: token.New(1, 41, 40, 4, token.KeywordJoin, "JOIN"),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 46, 45, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 54, 53, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 56, 55, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 63, 62, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 68, 67, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint,NATURAL and JOIN in join operator and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1 NATURAL JOIN myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Natural: token.New(1, 41, 40, 7, token.KeywordNatural, "NATURAL"),
														Join:    token.New(1, 49, 48, 4, token.KeywordJoin, "JOIN"),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 54, 53, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 62, 61, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 64, 63, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 71, 70, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 76, 75, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint, LEFT and JOIN in join operator and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1 LEFT JOIN myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Left: token.New(1, 41, 40, 4, token.KeywordLeft, "LEFT"),
														Join: token.New(1, 46, 45, 4, token.KeywordJoin, "JOIN"),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 51, 50, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 61, 60, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 68, 67, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 73, 72, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint, LEFT, OUTER and JOIN in join operator and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1 LEFT OUTER JOIN myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Left:  token.New(1, 41, 40, 4, token.KeywordLeft, "LEFT"),
														Outer: token.New(1, 46, 45, 5, token.KeywordOuter, "OUTER"),
														Join:  token.New(1, 52, 51, 4, token.KeywordJoin, "JOIN"),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 57, 56, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 65, 64, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 67, 66, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 74, 73, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 79, 78, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint,INNER and JOIN in join operator and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1 INNER JOIN myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Inner: token.New(1, 41, 40, 5, token.KeywordInner, "INNER"),
														Join:  token.New(1, 47, 46, 4, token.KeywordJoin, "JOIN"),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 52, 51, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 60, 59, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 62, 61, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 69, 68, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 74, 73, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with FROM with joinclause's basic join constraint,CROSS and JOIN in join operator and basic cte-table-name",
			"WITH myTable AS (SELECT * FROM myTable1 CROSS JOIN myTable2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											From: token.New(1, 27, 26, 4, token.KeywordFrom, "FROM"),
											JoinClause: &ast.JoinClause{
												TableOrSubquery: &ast.TableOrSubquery{
													TableName: token.New(1, 32, 31, 8, token.Literal, "myTable1"),
												},
												JoinClausePart: &ast.JoinClausePart{
													JoinOperator: &ast.JoinOperator{
														Cross: token.New(1, 41, 40, 5, token.KeywordCross, "CROSS"),
														Join:  token.New(1, 47, 46, 4, token.KeywordJoin, "JOIN"),
													},
													TableOrSubquery: &ast.TableOrSubquery{
														TableName: token.New(1, 52, 51, 8, token.Literal, "myTable2"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 60, 59, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 62, 61, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 69, 68, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 74, 73, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WHERE and basic cte-table-name",
			"WITH myTable AS (SELECT * WHERE myExpr) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Where: token.New(1, 27, 26, 5, token.KeywordWhere, "WHERE"),
											Expr1: &ast.Expr{
												LiteralValue: token.New(1, 33, 32, 6, token.Literal, "myExpr"),
											},
										},
									},
								},
								RightParen: token.New(1, 39, 38, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 41, 40, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 48, 47, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 53, 52, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with GROUP BY and single expr, and basic cte-table-name",
			"WITH myTable AS (SELECT * GROUP BY myExpr) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Group: token.New(1, 27, 26, 5, token.KeywordGroup, "GROUP"),
											By:    token.New(1, 33, 32, 2, token.KeywordBy, "BY"),
											Expr2: []*ast.Expr{
												&ast.Expr{
													LiteralValue: token.New(1, 36, 35, 6, token.Literal, "myExpr"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 42, 41, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 44, 43, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 51, 50, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 56, 55, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with GROUP BY and multiple expr, and basic cte-table-name",
			"WITH myTable AS (SELECT * GROUP BY myExpr1,myExpr2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Group: token.New(1, 27, 26, 5, token.KeywordGroup, "GROUP"),
											By:    token.New(1, 33, 32, 2, token.KeywordBy, "BY"),
											Expr2: []*ast.Expr{
												&ast.Expr{
													LiteralValue: token.New(1, 36, 35, 7, token.Literal, "myExpr1"),
												},
												&ast.Expr{
													LiteralValue: token.New(1, 44, 43, 7, token.Literal, "myExpr2"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 53, 52, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 60, 59, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 65, 64, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with GROUP BY, multiple expr and HAVING, and basic cte-table-name",
			"WITH myTable AS (SELECT * GROUP BY myExpr1,myExpr2 HAVING myExpr3) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Group: token.New(1, 27, 26, 5, token.KeywordGroup, "GROUP"),
											By:    token.New(1, 33, 32, 2, token.KeywordBy, "BY"),
											Expr2: []*ast.Expr{
												&ast.Expr{
													LiteralValue: token.New(1, 36, 35, 7, token.Literal, "myExpr1"),
												},
												&ast.Expr{
													LiteralValue: token.New(1, 44, 43, 7, token.Literal, "myExpr2"),
												},
											},
											Having: token.New(1, 52, 51, 6, token.KeywordHaving, "HAVING"),
											Expr3: &ast.Expr{
												LiteralValue: token.New(1, 59, 58, 7, token.Literal, "myExpr3"),
											},
										},
									},
								},
								RightParen: token.New(1, 66, 65, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 68, 67, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 75, 74, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 80, 79, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and basic WindowDefn and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS ()) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen:  token.New(1, 46, 45, 1, token.Delimiter, "("),
														RightParen: token.New(1, 47, 46, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 48, 47, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 50, 49, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 57, 56, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 62, 61, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basiWindowName, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (basicWindowName)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen:      token.New(1, 46, 45, 1, token.Delimiter, "("),
														BaseWindowName: token.New(1, 47, 46, 15, token.Literal, "basicWindowName"),
														RightParen:     token.New(1, 62, 61, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 63, 62, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 65, 64, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 72, 71, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 77, 76, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with PARTITION and single expr, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (PARTITION BY myExpr)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Partition: token.New(1, 47, 46, 9, token.KeywordPartition, "PARTITION"),
														By1:       token.New(1, 57, 56, 2, token.KeywordBy, "BY"),
														Expr: []*ast.Expr{
															&ast.Expr{
																LiteralValue: token.New(1, 60, 59, 6, token.Literal, "myExpr"),
															},
														},
														RightParen: token.New(1, 66, 65, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 67, 66, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 69, 68, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 76, 75, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 81, 80, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with PARTITION and multiple expr, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (PARTITION BY myExpr1,myExpr2)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Partition: token.New(1, 47, 46, 9, token.KeywordPartition, "PARTITION"),
														By1:       token.New(1, 57, 56, 2, token.KeywordBy, "BY"),
														Expr: []*ast.Expr{
															&ast.Expr{
																LiteralValue: token.New(1, 60, 59, 7, token.Literal, "myExpr1"),
															},
															&ast.Expr{
																LiteralValue: token.New(1, 68, 67, 7, token.Literal, "myExpr2"),
															},
														},
														RightParen: token.New(1, 75, 74, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 76, 75, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 78, 77, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 85, 84, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 90, 89, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 6, token.Literal, "myExpr"),
																},
															},
														},
														RightParen: token.New(1, 62, 61, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 63, 62, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 65, 64, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 72, 71, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 77, 76, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY and multiple basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr1,myExpr2)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 7, token.Literal, "myExpr1"),
																},
															},
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 64, 63, 7, token.Literal, "myExpr2"),
																},
															},
														},
														RightParen: token.New(1, 71, 70, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 72, 71, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 74, 73, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 81, 80, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 86, 85, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY, single basic ordering term and COLLATE, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr1 COLLATE myCollation)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 7, token.Literal, "myExpr1"),
																},
																Collate:       token.New(1, 64, 63, 7, token.KeywordCollate, "COLLATE"),
																CollationName: token.New(1, 72, 71, 11, token.Literal, "myCollation"),
															},
														},
														RightParen: token.New(1, 83, 82, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 84, 83, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 86, 85, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 93, 92, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 98, 97, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY, single basic ordering term and ASC, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr1 ASC)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 7, token.Literal, "myExpr1"),
																},
																Asc: token.New(1, 64, 63, 3, token.KeywordAsc, "ASC"),
															},
														},
														RightParen: token.New(1, 67, 66, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 68, 67, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 70, 69, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 77, 76, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 82, 81, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY, single basic ordering term and DESC, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr1 DESC)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 7, token.Literal, "myExpr1"),
																},
																Desc: token.New(1, 64, 63, 4, token.KeywordDesc, "DESC"),
															},
														},
														RightParen: token.New(1, 68, 67, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 69, 68, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 71, 70, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 78, 77, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 83, 82, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY, single basic ordering term and NULLS FIRST, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr1 NULLS FIRST)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 7, token.Literal, "myExpr1"),
																},
																Nulls: token.New(1, 64, 63, 5, token.KeywordNulls, "NULLS"),
																First: token.New(1, 70, 69, 5, token.KeywordFirst, "FIRST"),
															},
														},
														RightParen: token.New(1, 75, 74, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 76, 75, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 78, 77, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 85, 84, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 90, 89, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with ORDER BY, single basic ordering term and NULLS LAST, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ORDER BY myExpr1 NULLS LAST)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Order:     token.New(1, 47, 46, 5, token.KeywordOrder, "ORDER"),
														By2:       token.New(1, 53, 52, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 56, 55, 7, token.Literal, "myExpr1"),
																},
																Nulls: token.New(1, 64, 63, 5, token.KeywordNulls, "NULLS"),
																Last:  token.New(1, 70, 69, 4, token.KeywordLast, "LAST"),
															},
														},
														RightParen: token.New(1, 74, 73, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 75, 74, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 77, 76, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 84, 83, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 89, 88, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE UNBOUNDED PRECEDING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:      token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Unbounded1: token.New(1, 53, 52, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 63, 62, 9, token.KeywordPreceding, "PRECEDING"),
														},
														RightParen: token.New(1, 72, 71, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 73, 72, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 75, 74, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 82, 81, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 87, 86, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with ROWS and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (ROWS UNBOUNDED PRECEDING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Rows:       token.New(1, 47, 46, 4, token.KeywordRows, "ROWS"),
															Unbounded1: token.New(1, 52, 51, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 62, 61, 9, token.KeywordPreceding, "PRECEDING"),
														},
														RightParen: token.New(1, 71, 70, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 72, 71, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 74, 73, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 81, 80, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 86, 85, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with GROUPS, UNBOUNDED PRECEDING and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (GROUPS UNBOUNDED PRECEDING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Groups:     token.New(1, 47, 46, 6, token.KeywordGroups, "GROUPS"),
															Unbounded1: token.New(1, 54, 53, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 64, 63, 9, token.KeywordPreceding, "PRECEDING"),
														},
														RightParen: token.New(1, 73, 72, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 74, 73, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 76, 75, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 83, 82, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 88, 87, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE and expr PRECEDING and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE myLiteral PRECEDING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range: token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Expr1: &ast.Expr{
																LiteralValue: token.New(1, 53, 52, 9, token.Literal, "myLiteral"),
															},
															Preceding1: token.New(1, 63, 62, 9, token.KeywordPreceding, "PRECEDING"),
														},
														RightParen: token.New(1, 72, 71, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 73, 72, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 75, 74, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 82, 81, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 87, 86, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE and CURRENT ROW and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE CURRENT ROW)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:    token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Current1: token.New(1, 53, 52, 7, token.KeywordCurrent, "CURRENT"),
															Row1:     token.New(1, 61, 60, 3, token.KeywordRow, "ROW"),
														},
														RightParen: token.New(1, 64, 63, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 65, 64, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 67, 66, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 74, 73, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 79, 78, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with GROUPS, BETWEEN UNBOUNDED PRECEDING, AND, expr PRECEDING and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (GROUPS BETWEEN UNBOUNDED PRECEDING AND myLiteral PRECEDING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Groups:     token.New(1, 47, 46, 6, token.KeywordGroups, "GROUPS"),
															Between:    token.New(1, 54, 53, 7, token.KeywordBetween, "BETWEEN"),
															Unbounded1: token.New(1, 62, 61, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 72, 71, 9, token.KeywordPreceding, "PRECEDING"),
															And:        token.New(1, 82, 81, 3, token.KeywordAnd, "AND"),
															Expr2: &ast.Expr{
																LiteralValue: token.New(1, 86, 85, 9, token.Literal, "myLiteral"),
															},
															Preceding2: token.New(1, 96, 95, 9, token.KeywordPreceding, "PRECEDING"),
														},
														RightParen: token.New(1, 105, 104, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 106, 105, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 108, 107, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 115, 114, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 120, 119, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, BETWEEN and expr PRECEDING, AND, expr FOLLOWING and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE BETWEEN myLiteral PRECEDING AND myExpr FOLLOWING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:   token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Between: token.New(1, 53, 52, 7, token.KeywordBetween, "BETWEEN"),
															Expr1: &ast.Expr{
																LiteralValue: token.New(1, 61, 60, 9, token.Literal, "myLiteral"),
															},
															Preceding1: token.New(1, 71, 70, 9, token.KeywordPreceding, "PRECEDING"),
															And:        token.New(1, 81, 80, 3, token.KeywordAnd, "AND"),
															Expr2: &ast.Expr{
																LiteralValue: token.New(1, 85, 84, 6, token.Literal, "myExpr"),
															},
															Following2: token.New(1, 92, 91, 9, token.KeywordFollowing, "FOLLOWING"),
														},
														RightParen: token.New(1, 101, 100, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 102, 101, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 104, 103, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 111, 110, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 116, 115, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, BETWEEEN, CURRENT ROW, AND and UNBOUNDED FOLLOWING, and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:      token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Between:    token.New(1, 53, 52, 7, token.KeywordBetween, "BETWEEN"),
															Current1:   token.New(1, 61, 60, 7, token.KeywordCurrent, "CURRENT"),
															Row1:       token.New(1, 69, 68, 3, token.KeywordRow, "ROW"),
															And:        token.New(1, 73, 72, 3, token.KeywordAnd, "AND"),
															Unbounded2: token.New(1, 77, 76, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Following2: token.New(1, 87, 86, 9, token.KeywordFollowing, "FOLLOWING"),
														},
														RightParen: token.New(1, 96, 95, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 97, 96, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 99, 98, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 106, 105, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 111, 110, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, BETWEEEN, expr FOLLOWING, AND and CURRENT ROW, and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE BETWEEN myExpr FOLLOWING AND CURRENT ROW)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:   token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Between: token.New(1, 53, 52, 7, token.KeywordBetween, "BETWEEN"),
															Expr1: &ast.Expr{
																LiteralValue: token.New(1, 61, 60, 6, token.Literal, "myExpr"),
															},
															Following1: token.New(1, 68, 67, 9, token.KeywordFollowing, "FOLLOWING"),
															And:        token.New(1, 78, 77, 3, token.KeywordAnd, "AND"),
															Current2:   token.New(1, 82, 81, 7, token.KeywordCurrent, "CURRENT"),
															Row2:       token.New(1, 90, 89, 3, token.KeywordRow, "ROW"),
														},
														RightParen: token.New(1, 93, 92, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 94, 93, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 96, 95, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 103, 102, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 108, 107, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, EXCLUDE NO OTHERS and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE UNBOUNDED PRECEDING EXCLUDE NO OTHERS)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:      token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Unbounded1: token.New(1, 53, 52, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 63, 62, 9, token.KeywordPreceding, "PRECEDING"),
															Exclude:    token.New(1, 73, 72, 7, token.KeywordExclude, "EXCLUDE"),
															No:         token.New(1, 81, 80, 2, token.KeywordNo, "NO"),
															Others:     token.New(1, 84, 83, 6, token.KeywordOthers, "OTHERS"),
														},
														RightParen: token.New(1, 90, 89, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 91, 90, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 93, 92, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 100, 99, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 105, 104, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, EXCLUDE CURRENT ROW and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE UNBOUNDED PRECEDING EXCLUDE CURRENT ROW)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:      token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Unbounded1: token.New(1, 53, 52, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 63, 62, 9, token.KeywordPreceding, "PRECEDING"),
															Exclude:    token.New(1, 73, 72, 7, token.KeywordExclude, "EXCLUDE"),
															Current3:   token.New(1, 81, 80, 7, token.KeywordCurrent, "CURRENT"),
															Row3:       token.New(1, 89, 88, 3, token.KeywordRow, "ROW"),
														},
														RightParen: token.New(1, 92, 91, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 93, 92, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 95, 94, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 102, 101, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 107, 106, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, EXCLUDE GROUP and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE UNBOUNDED PRECEDING EXCLUDE GROUP)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:      token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Unbounded1: token.New(1, 53, 52, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 63, 62, 9, token.KeywordPreceding, "PRECEDING"),
															Exclude:    token.New(1, 73, 72, 7, token.KeywordExclude, "EXCLUDE"),
															Group:      token.New(1, 81, 80, 5, token.KeywordGroup, "GROUP"),
														},
														RightParen: token.New(1, 86, 85, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 87, 86, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 89, 88, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 96, 95, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 101, 100, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE, EXCLUDE TIES and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (RANGE UNBOUNDED PRECEDING EXCLUDE TIES)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														FrameSpec: &ast.FrameSpec{
															Range:      token.New(1, 47, 46, 5, token.KeywordRange, "RANGE"),
															Unbounded1: token.New(1, 53, 52, 9, token.KeywordUnbounded, "UNBOUNDED"),
															Preceding1: token.New(1, 63, 62, 9, token.KeywordPreceding, "PRECEDING"),
															Exclude:    token.New(1, 73, 72, 7, token.KeywordExclude, "EXCLUDE"),
															Ties:       token.New(1, 81, 80, 4, token.KeywordTies, "TIES"),
														},
														RightParen: token.New(1, 85, 84, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 86, 85, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 88, 87, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 95, 94, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 100, 99, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, select stmt with WINDOW and WindowDefn with basic frame spec with RANGE and CURRENT ROW and single basic ordering term, and basic cte-table-name",
			"WITH myTable AS (SELECT * WINDOW myWindow AS (PARTITION BY myExpr1 ORDER BY myExpr2 RANGE CURRENT ROW)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											Window: token.New(1, 27, 26, 6, token.KeywordWindow, "WINDOW"),
											NamedWindow: []*ast.NamedWindow{
												&ast.NamedWindow{
													WindowName: token.New(1, 34, 33, 8, token.Literal, "myWindow"),
													As:         token.New(1, 43, 42, 2, token.KeywordAs, "AS"),
													WindowDefn: &ast.WindowDefn{
														LeftParen: token.New(1, 46, 45, 1, token.Delimiter, "("),
														Partition: token.New(1, 47, 46, 9, token.KeywordPartition, "PARTITION"),
														By1:       token.New(1, 57, 56, 2, token.KeywordBy, "BY"),
														Expr: []*ast.Expr{
															&ast.Expr{
																LiteralValue: token.New(1, 60, 59, 7, token.Literal, "myExpr1"),
															},
														},
														Order: token.New(1, 68, 67, 5, token.KeywordOrder, "ORDER"),
														By2:   token.New(1, 74, 73, 2, token.KeywordBy, "BY"),
														OrderingTerm: []*ast.OrderingTerm{
															&ast.OrderingTerm{
																Expr: &ast.Expr{
																	LiteralValue: token.New(1, 77, 76, 7, token.Literal, "myExpr2"),
																},
															},
														},
														FrameSpec: &ast.FrameSpec{
															Range:    token.New(1, 85, 84, 5, token.KeywordRange, "RANGE"),
															Current1: token.New(1, 91, 90, 7, token.KeywordCurrent, "CURRENT"),
															Row1:     token.New(1, 99, 98, 3, token.KeywordRow, "ROW"),
														},
														RightParen: token.New(1, 102, 101, 1, token.Delimiter, ")"),
													},
												},
											},
										},
									},
								},
								RightParen: token.New(1, 103, 102, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 105, 104, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 112, 111, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 117, 116, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, VALUES with single expr with single set, and basic cte-table-name",
			"WITH myTable AS (VALUES (myExpr)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Values: token.New(1, 18, 17, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 25, 24, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 26, 25, 6, token.Literal, "myExpr"),
														},
													},
													RightParen: token.New(1, 32, 31, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 33, 32, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 35, 34, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 42, 41, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 47, 46, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, VALUES with multiple expr with single set, and basic cte-table-name",
			"WITH myTable AS (VALUES (myExpr1,myExpr2)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Values: token.New(1, 18, 17, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 25, 24, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 26, 25, 7, token.Literal, "myExpr1"),
														},
														&ast.Expr{
															LiteralValue: token.New(1, 34, 33, 7, token.Literal, "myExpr2"),
														},
													},
													RightParen: token.New(1, 41, 40, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 42, 41, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 44, 43, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 51, 50, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 56, 55, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, VALUES with multiple expr with multiple sets, and basic cte-table-name",
			"WITH myTable AS (VALUES (myExpr1,myExpr2),(myExpr1,myExpr2)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Values: token.New(1, 18, 17, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 25, 24, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 26, 25, 7, token.Literal, "myExpr1"),
														},
														&ast.Expr{
															LiteralValue: token.New(1, 34, 33, 7, token.Literal, "myExpr2"),
														},
													},
													RightParen: token.New(1, 41, 40, 1, token.Delimiter, ")"),
												},
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 43, 42, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 44, 43, 7, token.Literal, "myExpr1"),
														},
														&ast.Expr{
															LiteralValue: token.New(1, 52, 51, 7, token.Literal, "myExpr2"),
														},
													},
													RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 60, 59, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 62, 61, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 69, 68, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 74, 73, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic VALUES and basic SELECT with UNION compound operator, and basic cte-table-name",
			"WITH myTable AS (SELECT * UNION VALUES (myExpr1)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											CompoundOperator: &ast.CompoundOperator{
												Union: token.New(1, 27, 26, 5, token.KeywordUnion, "UNION"),
											},
										},
										&ast.SelectCore{

											Values: token.New(1, 33, 32, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 40, 39, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 41, 40, 7, token.Literal, "myExpr1"),
														},
													},
													RightParen: token.New(1, 48, 47, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 49, 48, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 51, 50, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 58, 57, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 63, 62, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic VALUES and basic SELECT with UNION ALL compound operator, and basic cte-table-name",
			"WITH myTable AS (SELECT * UNION ALL VALUES (myExpr1)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											CompoundOperator: &ast.CompoundOperator{
												Union: token.New(1, 27, 26, 5, token.KeywordUnion, "UNION"),
												All:   token.New(1, 33, 32, 3, token.KeywordAll, "ALL"),
											},
										},
										&ast.SelectCore{

											Values: token.New(1, 37, 36, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 44, 43, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 45, 44, 7, token.Literal, "myExpr1"),
														},
													},
													RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 53, 52, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 55, 54, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 62, 61, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 67, 66, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic VALUES and basic SELECT with INTERSECT compound operator, and basic cte-table-name",
			"WITH myTable AS (SELECT * INTERSECT VALUES (myExpr1)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											CompoundOperator: &ast.CompoundOperator{
												Intersect: token.New(1, 27, 26, 9, token.KeywordIntersect, "INTERSECT"),
											},
										},
										&ast.SelectCore{

											Values: token.New(1, 37, 36, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 44, 43, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 45, 44, 7, token.Literal, "myExpr1"),
														},
													},
													RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 53, 52, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 55, 54, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 62, 61, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 67, 66, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic VALUES and basic SELECT with EXCEPT compound operator, and basic cte-table-name",
			"WITH myTable AS (SELECT * EXCEPT VALUES (myExpr1)) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
											CompoundOperator: &ast.CompoundOperator{
												Except: token.New(1, 27, 26, 6, token.KeywordExcept, "EXCEPT"),
											},
										},
										&ast.SelectCore{

											Values: token.New(1, 34, 33, 6, token.KeywordValues, "VALUES"),
											ParenthesizedExpressions: []*ast.ParenthesizedExpressions{
												&ast.ParenthesizedExpressions{
													LeftParen: token.New(1, 41, 40, 1, token.Delimiter, "("),
													Exprs: []*ast.Expr{
														&ast.Expr{
															LiteralValue: token.New(1, 42, 41, 7, token.Literal, "myExpr1"),
														},
													},
													RightParen: token.New(1, 49, 48, 1, token.Delimiter, ")"),
												},
											},
										},
									},
								},
								RightParen: token.New(1, 50, 49, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 52, 51, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 59, 58, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 64, 63, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic SELECT with ORDER BY, and basic cte-table-name",
			"WITH myTable AS (SELECT * ORDER BY myLiteral) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
									Order: token.New(1, 27, 26, 5, token.KeywordOrder, "ORDER"),
									By:    token.New(1, 33, 32, 2, token.KeywordBy, "BY"),
									OrderingTerm: []*ast.OrderingTerm{
										&ast.OrderingTerm{
											Expr: &ast.Expr{
												LiteralValue: token.New(1, 36, 35, 9, token.Literal, "myLiteral"),
											},
										},
									},
								},
								RightParen: token.New(1, 45, 44, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 47, 46, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 54, 53, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 59, 58, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic SELECT with basic LIMIT with single Expr, and basic cte-table-name",
			"WITH myTable AS (SELECT * LIMIT myExpr1) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
									Limit: token.New(1, 27, 26, 5, token.KeywordLimit, "LIMIT"),
									Expr1: &ast.Expr{
										LiteralValue: token.New(1, 33, 32, 7, token.Literal, "myExpr1"),
									},
								},
								RightParen: token.New(1, 40, 39, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 42, 41, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 49, 48, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 54, 53, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic SELECT with LIMIT with multiple Expr with comma, and basic cte-table-name",
			"WITH myTable AS (SELECT * LIMIT myExpr1,myExpr2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
									Limit: token.New(1, 27, 26, 5, token.KeywordLimit, "LIMIT"),
									Expr1: &ast.Expr{
										LiteralValue: token.New(1, 33, 32, 7, token.Literal, "myExpr1"),
									},
									Comma: token.New(1, 40, 39, 1, token.Delimiter, ","),
									Expr2: &ast.Expr{
										LiteralValue: token.New(1, 41, 40, 7, token.Literal, "myExpr2"),
									},
								},
								RightParen: token.New(1, 48, 47, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 50, 49, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 57, 56, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 62, 61, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			"DELETE with basic with clause, basic SELECT with LIMIT with multiple Expr with OFFSET, and basic cte-table-name",
			"WITH myTable AS (SELECT * LIMIT myExpr1 OFFSET myExpr2) DELETE FROM myTable",
			&ast.SQLStmt{
				DeleteStmt: &ast.DeleteStmt{
					WithClause: &ast.WithClause{
						With: token.New(1, 1, 0, 4, token.KeywordWith, "WITH"),
						RecursiveCte: []*ast.RecursiveCte{
							&ast.RecursiveCte{
								CteTableName: &ast.CteTableName{
									TableName: token.New(1, 6, 5, 7, token.Literal, "myTable"),
								},
								As:        token.New(1, 14, 13, 2, token.KeywordAs, "AS"),
								LeftParen: token.New(1, 17, 16, 1, token.Delimiter, "("),
								SelectStmt: &ast.SelectStmt{
									SelectCore: []*ast.SelectCore{
										&ast.SelectCore{
											Select: token.New(1, 18, 17, 6, token.KeywordSelect, "SELECT"),
											ResultColumn: []*ast.ResultColumn{
												&ast.ResultColumn{
													Asterisk: token.New(1, 25, 24, 1, token.BinaryOperator, "*"),
												},
											},
										},
									},
									Limit: token.New(1, 27, 26, 5, token.KeywordLimit, "LIMIT"),
									Expr1: &ast.Expr{
										LiteralValue: token.New(1, 33, 32, 7, token.Literal, "myExpr1"),
									},
									Offset: token.New(1, 41, 40, 6, token.KeywordOffset, "OFFSET"),
									Expr2: &ast.Expr{
										LiteralValue: token.New(1, 48, 47, 7, token.Literal, "myExpr2"),
									},
								},
								RightParen: token.New(1, 55, 54, 1, token.Delimiter, ")"),
							},
						},
					},
					Delete: token.New(1, 57, 56, 6, token.KeywordDelete, "DELETE"),
					From:   token.New(1, 64, 63, 4, token.KeywordFrom, "FROM"),
					QualifiedTableName: &ast.QualifiedTableName{
						TableName: token.New(1, 69, 68, 7, token.Literal, "myTable"),
					},
				},
			},
		},
		{
			`CREATE TABLE basic with basic select`,
			"CREATE TABLE myTable AS SELECT *",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					As:        token.New(1, 22, 21, 2, token.KeywordAs, "AS"),
					SelectStmt: &ast.SelectStmt{
						SelectCore: []*ast.SelectCore{
							&ast.SelectCore{
								Select: token.New(1, 25, 24, 6, token.KeywordSelect, "SELECT"),
								ResultColumn: []*ast.ResultColumn{
									&ast.ResultColumn{
										Asterisk: token.New(1, 32, 31, 1, token.BinaryOperator, "*"),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			`CREATE TABLE with TEMP`,
			"CREATE TEMP TABLE myTable AS SELECT *",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Temp:      token.New(1, 8, 7, 4, token.KeywordTemp, "TEMP"),
					Table:     token.New(1, 13, 12, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 19, 18, 7, token.Literal, "myTable"),
					As:        token.New(1, 27, 26, 2, token.KeywordAs, "AS"),
					SelectStmt: &ast.SelectStmt{
						SelectCore: []*ast.SelectCore{
							&ast.SelectCore{
								Select: token.New(1, 30, 29, 6, token.KeywordSelect, "SELECT"),
								ResultColumn: []*ast.ResultColumn{
									&ast.ResultColumn{
										Asterisk: token.New(1, 37, 36, 1, token.BinaryOperator, "*"),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			`CREATE TABLE with TEMPORARY`,
			"CREATE TEMPORARY TABLE myTable AS SELECT *",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Temporary: token.New(1, 8, 7, 9, token.KeywordTemporary, "TEMPORARY"),
					Table:     token.New(1, 18, 17, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 24, 23, 7, token.Literal, "myTable"),
					As:        token.New(1, 32, 31, 2, token.KeywordAs, "AS"),
					SelectStmt: &ast.SelectStmt{
						SelectCore: []*ast.SelectCore{
							&ast.SelectCore{
								Select: token.New(1, 35, 34, 6, token.KeywordSelect, "SELECT"),
								ResultColumn: []*ast.ResultColumn{
									&ast.ResultColumn{
										Asterisk: token.New(1, 42, 41, 1, token.BinaryOperator, "*"),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			`CREATE TABLE with IF NOT EXISTS`,
			"CREATE TABLE IF NOT EXISTS myTable AS SELECT *",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					If:        token.New(1, 14, 13, 2, token.KeywordIf, "IF"),
					Not:       token.New(1, 17, 16, 3, token.KeywordNot, "NOT"),
					Exists:    token.New(1, 21, 20, 6, token.KeywordExists, "EXISTS"),
					TableName: token.New(1, 28, 27, 7, token.Literal, "myTable"),
					As:        token.New(1, 36, 35, 2, token.KeywordAs, "AS"),
					SelectStmt: &ast.SelectStmt{
						SelectCore: []*ast.SelectCore{
							&ast.SelectCore{
								Select: token.New(1, 39, 38, 6, token.KeywordSelect, "SELECT"),
								ResultColumn: []*ast.ResultColumn{
									&ast.ResultColumn{
										Asterisk: token.New(1, 46, 45, 1, token.BinaryOperator, "*"),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			`CREATE TABLE with schema and table name`,
			"CREATE TABLE mySchema.myTable AS SELECT *",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:     token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:      token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					SchemaName: token.New(1, 14, 13, 8, token.Literal, "mySchema"),
					Period:     token.New(1, 22, 21, 1, token.Literal, "."),
					TableName:  token.New(1, 23, 22, 7, token.Literal, "myTable"),
					As:         token.New(1, 31, 30, 2, token.KeywordAs, "AS"),
					SelectStmt: &ast.SelectStmt{
						SelectCore: []*ast.SelectCore{
							&ast.SelectCore{
								Select: token.New(1, 34, 33, 6, token.KeywordSelect, "SELECT"),
								ResultColumn: []*ast.ResultColumn{
									&ast.ResultColumn{
										Asterisk: token.New(1, 41, 40, 1, token.BinaryOperator, "*"),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def`,
			"CREATE TABLE myTable (myColumn)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
						},
					},
					RightParen: token.New(1, 31, 30, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with multiple basic column-def`,
			"CREATE TABLE myTable (myColumn1,myColumn2)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
						&ast.ColumnDef{
							ColumnName: token.New(1, 33, 32, 9, token.Literal, "myColumn2"),
						},
					},
					RightParen: token.New(1, 42, 41, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and basic table-constraint`,
			"CREATE TABLE myTable (myColumn1,CHECK (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Check:     token.New(1, 33, 32, 5, token.KeywordCheck, "CHECK"),
							LeftParen: token.New(1, 39, 38, 1, token.Delimiter, "("),
							Expr: &ast.Expr{
								LiteralValue: token.New(1, 40, 39, 6, token.Literal, "myExpr"),
							},
							RightParen: token.New(1, 46, 45, 1, token.Delimiter, ")"),
						},
					},
					RightParen: token.New(1, 47, 46, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and CONSTRAINT`,
			"CREATE TABLE myTable (myColumn1,CONSTRAINT myConstraint CHECK (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Constraint: token.New(1, 33, 32, 10, token.KeywordConstraint, "CONSTRAINT"),
							Name:       token.New(1, 44, 43, 12, token.Literal, "myConstraint"),
							Check:      token.New(1, 57, 56, 5, token.KeywordCheck, "CHECK"),
							LeftParen:  token.New(1, 63, 62, 1, token.Delimiter, "("),
							Expr: &ast.Expr{
								LiteralValue: token.New(1, 64, 63, 6, token.Literal, "myExpr"),
							},
							RightParen: token.New(1, 70, 69, 1, token.Delimiter, ")"),
						},
					},
					RightParen: token.New(1, 71, 70, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and PRIMARY KEY with single indexed-column`,
			"CREATE TABLE myTable (myColumn1,PRIMARY KEY (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Primary:   token.New(1, 33, 32, 7, token.KeywordPrimary, "PRIMARY"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 46, 45, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
						},
					},
					RightParen: token.New(1, 53, 52, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and PRIMARY KEY with single indexed-column and conflict-clause with ROLLBACK`,
			"CREATE TABLE myTable (myColumn1,PRIMARY KEY (myExpr) ON CONFLICT ROLLBACK)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Primary:   token.New(1, 33, 32, 7, token.KeywordPrimary, "PRIMARY"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 46, 45, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
							ConflictClause: &ast.ConflictClause{
								On:       token.New(1, 54, 53, 2, token.KeywordOn, "ON"),
								Conflict: token.New(1, 57, 56, 8, token.KeywordConflict, "CONFLICT"),
								Rollback: token.New(1, 66, 65, 8, token.KeywordRollback, "ROLLBACK"),
							},
						},
					},
					RightParen: token.New(1, 74, 73, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and PRIMARY KEY with single indexed-column and conflict-clause with ABORT`,
			"CREATE TABLE myTable (myColumn1,PRIMARY KEY (myExpr) ON CONFLICT ABORT)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Primary:   token.New(1, 33, 32, 7, token.KeywordPrimary, "PRIMARY"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 46, 45, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
							ConflictClause: &ast.ConflictClause{
								On:       token.New(1, 54, 53, 2, token.KeywordOn, "ON"),
								Conflict: token.New(1, 57, 56, 8, token.KeywordConflict, "CONFLICT"),
								Abort:    token.New(1, 66, 65, 5, token.KeywordAbort, "ABORT"),
							},
						},
					},
					RightParen: token.New(1, 71, 70, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and PRIMARY KEY with single indexed-column and conflict-clause with FAIL`,
			"CREATE TABLE myTable (myColumn1,PRIMARY KEY (myExpr) ON CONFLICT FAIL)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Primary:   token.New(1, 33, 32, 7, token.KeywordPrimary, "PRIMARY"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 46, 45, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
							ConflictClause: &ast.ConflictClause{
								On:       token.New(1, 54, 53, 2, token.KeywordOn, "ON"),
								Conflict: token.New(1, 57, 56, 8, token.KeywordConflict, "CONFLICT"),
								Fail:     token.New(1, 66, 65, 4, token.KeywordFail, "FAIL"),
							},
						},
					},
					RightParen: token.New(1, 70, 69, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and PRIMARY KEY with single indexed-column and conflict-clause with IGNORE`,
			"CREATE TABLE myTable (myColumn1,PRIMARY KEY (myExpr) ON CONFLICT IGNORE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Primary:   token.New(1, 33, 32, 7, token.KeywordPrimary, "PRIMARY"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 46, 45, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
							ConflictClause: &ast.ConflictClause{
								On:       token.New(1, 54, 53, 2, token.KeywordOn, "ON"),
								Conflict: token.New(1, 57, 56, 8, token.KeywordConflict, "CONFLICT"),
								Ignore:   token.New(1, 66, 65, 6, token.KeywordIgnore, "IGNORE"),
							},
						},
					},
					RightParen: token.New(1, 72, 71, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and PRIMARY KEY with single indexed-column and conflict-clause with REPLACE`,
			"CREATE TABLE myTable (myColumn1,PRIMARY KEY (myExpr) ON CONFLICT REPLACE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Primary:   token.New(1, 33, 32, 7, token.KeywordPrimary, "PRIMARY"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 46, 45, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 52, 51, 1, token.Delimiter, ")"),
							ConflictClause: &ast.ConflictClause{
								On:       token.New(1, 54, 53, 2, token.KeywordOn, "ON"),
								Conflict: token.New(1, 57, 56, 8, token.KeywordConflict, "CONFLICT"),
								Replace:  token.New(1, 66, 65, 7, token.KeywordReplace, "REPLACE"),
							},
						},
					},
					RightParen: token.New(1, 73, 72, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and UNIQUE`,
			"CREATE TABLE myTable (myColumn1,UNIQUE (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Unique:    token.New(1, 33, 32, 6, token.KeywordUnique, "UNIQUE"),
							LeftParen: token.New(1, 40, 39, 1, token.Delimiter, "("),
							IndexedColumn: []*ast.IndexedColumn{
								&ast.IndexedColumn{
									ColumnName: token.New(1, 41, 40, 6, token.Literal, "myExpr"),
								},
							},
							RightParen: token.New(1, 47, 46, 1, token.Delimiter, ")"),
						},
					},
					RightParen: token.New(1, 48, 47, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and basic foreign key clause`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
							},
						},
					},
					RightParen: token.New(1, 78, 77, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and multiple column name and basic foreign key clause`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol1,myCol2) REFERENCES myForeignTable)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 6, token.Literal, "myCol1"),
								token.New(1, 53, 52, 6, token.Literal, "myCol2"),
							},
							RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 61, 60, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 72, 71, 14, token.Literal, "myForeignTable"),
							},
						},
					},
					RightParen: token.New(1, 86, 85, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name, foreign key clause with single column name`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable (myNewCol))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								LeftParen:    token.New(1, 79, 78, 1, token.Delimiter, "("),
								ColumnName: []token.Token{
									token.New(1, 80, 79, 8, token.Literal, "myNewCol"),
								},
								RightParen: token.New(1, 88, 87, 1, token.Delimiter, ")"),
							},
						},
					},
					RightParen: token.New(1, 89, 88, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name, foreign key clause with mutiple column name`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable (myNewCol1,myNewCol2))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								LeftParen:    token.New(1, 79, 78, 1, token.Delimiter, "("),
								ColumnName: []token.Token{
									token.New(1, 80, 79, 9, token.Literal, "myNewCol1"),
									token.New(1, 90, 89, 9, token.Literal, "myNewCol2"),
								},
								RightParen: token.New(1, 99, 98, 1, token.Delimiter, ")"),
							},
						},
					},
					RightParen: token.New(1, 100, 99, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON DELETE SET NULL`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable ON DELETE SET NULL)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										On:     token.New(1, 79, 78, 2, token.KeywordOn, "ON"),
										Delete: token.New(1, 82, 81, 6, token.KeywordDelete, "DELETE"),
										Set:    token.New(1, 89, 88, 3, token.KeywordSet, "SET"),
										Null:   token.New(1, 93, 92, 4, token.KeywordNull, "NULL"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 97, 96, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON DELETE SET DEFAULT`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable ON DELETE SET DEFAULT)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										On:      token.New(1, 79, 78, 2, token.KeywordOn, "ON"),
										Delete:  token.New(1, 82, 81, 6, token.KeywordDelete, "DELETE"),
										Set:     token.New(1, 89, 88, 3, token.KeywordSet, "SET"),
										Default: token.New(1, 93, 92, 7, token.KeywordDefault, "DEFAULT"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 100, 99, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON DELETE CASCADE`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable ON DELETE CASCADE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										On:      token.New(1, 79, 78, 2, token.KeywordOn, "ON"),
										Delete:  token.New(1, 82, 81, 6, token.KeywordDelete, "DELETE"),
										Cascade: token.New(1, 89, 88, 7, token.KeywordCascade, "CASCADE"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 96, 95, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON DELETE RESTRICT`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable ON DELETE RESTRICT)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										On:       token.New(1, 79, 78, 2, token.KeywordOn, "ON"),
										Delete:   token.New(1, 82, 81, 6, token.KeywordDelete, "DELETE"),
										Restrict: token.New(1, 89, 88, 8, token.KeywordRestrict, "RESTRICT"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 97, 96, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON DELETE NO ACTION`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable ON DELETE NO ACTION)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										On:     token.New(1, 79, 78, 2, token.KeywordOn, "ON"),
										Delete: token.New(1, 82, 81, 6, token.KeywordDelete, "DELETE"),
										No:     token.New(1, 89, 88, 2, token.KeywordNo, "NO"),
										Action: token.New(1, 92, 91, 6, token.KeywordAction, "ACTION"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 98, 97, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON UPDATE NO ACTION`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable ON UPDATE NO ACTION)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										On:     token.New(1, 79, 78, 2, token.KeywordOn, "ON"),
										Update: token.New(1, 82, 81, 6, token.KeywordUpdate, "UPDATE"),
										No:     token.New(1, 89, 88, 2, token.KeywordNo, "NO"),
										Action: token.New(1, 92, 91, 6, token.KeywordAction, "ACTION"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 98, 97, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with variations of fkc core - ON UPDATE NO ACTION`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable MATCH myMatch)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										Match: token.New(1, 79, 78, 5, token.KeywordMatch, "MATCH"),
										Name:  token.New(1, 85, 84, 7, token.Literal, "myMatch"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 92, 91, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with multple fkc cores`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable MATCH myMatch ON DELETE NO ACTION)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								ForeignKeyClauseCore: []*ast.ForeignKeyClauseCore{
									&ast.ForeignKeyClauseCore{
										Match: token.New(1, 79, 78, 5, token.KeywordMatch, "MATCH"),
										Name:  token.New(1, 85, 84, 7, token.Literal, "myMatch"),
									},
									&ast.ForeignKeyClauseCore{
										On:     token.New(1, 93, 92, 2, token.KeywordOn, "ON"),
										Delete: token.New(1, 96, 95, 6, token.KeywordDelete, "DELETE"),
										No:     token.New(1, 103, 102, 2, token.KeywordNo, "NO"),
										Action: token.New(1, 106, 105, 6, token.KeywordAction, "ACTION"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 112, 111, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with DEFERRABLE`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable DEFERRABLE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								Deferrable:   token.New(1, 79, 78, 10, token.KeywordDeferrable, "DEFERRABLE"),
							},
						},
					},
					RightParen: token.New(1, 89, 88, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with NOT DEFERRABLE`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable NOT DEFERRABLE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								Not:          token.New(1, 79, 78, 3, token.KeywordNot, "NOT"),
								Deferrable:   token.New(1, 83, 82, 10, token.KeywordDeferrable, "DEFERRABLE"),
							},
						},
					},
					RightParen: token.New(1, 93, 92, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with DEFERRABLE INITIALLY DEFERRED`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable DEFERRABLE INITIALLY DEFERRED)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								Deferrable:   token.New(1, 79, 78, 10, token.KeywordDeferrable, "DEFERRABLE"),
								Initially:    token.New(1, 90, 89, 9, token.KeywordInitially, "INITIALLY"),
								Deferred:     token.New(1, 100, 99, 8, token.KeywordDeferred, "DEFERRED"),
							},
						},
					},
					RightParen: token.New(1, 108, 107, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single basic column-def and table-constraint and FOREIGN KEY and single column name and foreign key clause(fkc) with DEFERRABLE INITIALLY IMMEDIATE`,
			"CREATE TABLE myTable (myColumn1,FOREIGN KEY (myCol) REFERENCES myForeignTable DEFERRABLE INITIALLY IMMEDIATE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 9, token.Literal, "myColumn1"),
						},
					},
					TableConstraint: []*ast.TableConstraint{
						&ast.TableConstraint{
							Foreign:   token.New(1, 33, 32, 7, token.KeywordForeign, "FOREIGN"),
							Key:       token.New(1, 41, 40, 3, token.KeywordKey, "KEY"),
							LeftParen: token.New(1, 45, 44, 1, token.Delimiter, "("),
							ColumnName: []token.Token{
								token.New(1, 46, 45, 5, token.Literal, "myCol"),
							},
							RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
							ForeignKeyClause: &ast.ForeignKeyClause{
								References:   token.New(1, 53, 52, 10, token.KeywordReferences, "REFERENCES"),
								ForeignTable: token.New(1, 64, 63, 14, token.Literal, "myForeignTable"),
								Deferrable:   token.New(1, 79, 78, 10, token.KeywordDeferrable, "DEFERRABLE"),
								Initially:    token.New(1, 90, 89, 9, token.KeywordInitially, "INITIALLY"),
								Immediate:    token.New(1, 100, 99, 9, token.KeywordImmediate, "IMMEDIATE"),
							},
						},
					},
					RightParen: token.New(1, 109, 108, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint NOT NULL`,
			"CREATE TABLE myTable (myColumn CONSTRAINT myConstraint NOT NULL)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Constraint: token.New(1, 32, 31, 10, token.KeywordConstraint, "CONSTRAINT"),
									Name:       token.New(1, 43, 42, 12, token.Literal, "myConstraint"),
									Not:        token.New(1, 56, 55, 3, token.KeywordNot, "NOT"),
									Null:       token.New(1, 60, 59, 4, token.KeywordNull, "NULL"),
								},
							},
						},
					},
					RightParen: token.New(1, 64, 63, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint UNIQUE`,
			"CREATE TABLE myTable (myColumn CONSTRAINT myConstraint UNIQUE)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Constraint: token.New(1, 32, 31, 10, token.KeywordConstraint, "CONSTRAINT"),
									Name:       token.New(1, 43, 42, 12, token.Literal, "myConstraint"),
									Unique:     token.New(1, 56, 55, 6, token.KeywordUnique, "UNIQUE"),
								},
							},
						},
					},
					RightParen: token.New(1, 62, 61, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint CHECK`,
			"CREATE TABLE myTable (myColumn CONSTRAINT myConstraint CHECK (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Constraint: token.New(1, 32, 31, 10, token.KeywordConstraint, "CONSTRAINT"),
									Name:       token.New(1, 43, 42, 12, token.Literal, "myConstraint"),
									Check:      token.New(1, 56, 55, 5, token.KeywordCheck, "CHECK"),
									LeftParen:  token.New(1, 62, 61, 1, token.Delimiter, "("),
									Expr: &ast.Expr{
										LiteralValue: token.New(1, 63, 62, 6, token.Literal, "myExpr"),
									},
									RightParen: token.New(1, 69, 68, 1, token.Delimiter, ")"),
								},
							},
						},
					},
					RightParen: token.New(1, 70, 69, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint COLLATE`,
			"CREATE TABLE myTable (myColumn COLLATE myCollation)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Collate:       token.New(1, 32, 31, 7, token.KeywordCollate, "COLLATE"),
									CollationName: token.New(1, 40, 39, 11, token.Literal, "myCollation"),
								},
							},
						},
					},
					RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint fkc`,
			"CREATE TABLE myTable (myColumn REFERENCES myForeignTable)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									ForeignKeyClause: &ast.ForeignKeyClause{
										References:   token.New(1, 32, 31, 10, token.KeywordReferences, "REFERENCES"),
										ForeignTable: token.New(1, 43, 42, 14, token.Literal, "myForeignTable"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 57, 56, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint PRIMARY KEY basic`,
			"CREATE TABLE myTable (myColumn PRIMARY KEY)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Primary: token.New(1, 32, 31, 7, token.KeywordPrimary, "PRIMARY"),
									Key:     token.New(1, 40, 39, 3, token.KeywordKey, "KEY"),
								},
							},
						},
					},
					RightParen: token.New(1, 43, 42, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint PRIMARY KEY with ASC and AUTOINCREMENT`,
			"CREATE TABLE myTable (myColumn PRIMARY KEY ASC AUTOINCREMENT)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Primary:       token.New(1, 32, 31, 7, token.KeywordPrimary, "PRIMARY"),
									Key:           token.New(1, 40, 39, 3, token.KeywordKey, "KEY"),
									Asc:           token.New(1, 44, 43, 3, token.KeywordAsc, "ASC"),
									Autoincrement: token.New(1, 48, 47, 13, token.KeywordAutoincrement, "AUTOINCREMENT"),
								},
							},
						},
					},
					RightParen: token.New(1, 61, 60, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint PRIMARY KEY with DESC and AUTOINCREMENT`,
			"CREATE TABLE myTable (myColumn PRIMARY KEY DESC AUTOINCREMENT)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Primary:       token.New(1, 32, 31, 7, token.KeywordPrimary, "PRIMARY"),
									Key:           token.New(1, 40, 39, 3, token.KeywordKey, "KEY"),
									Desc:          token.New(1, 44, 43, 4, token.KeywordDesc, "DESC"),
									Autoincrement: token.New(1, 49, 48, 13, token.KeywordAutoincrement, "AUTOINCREMENT"),
								},
							},
						},
					},
					RightParen: token.New(1, 62, 61, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint GENERATED ALWAYS and AS`,
			"CREATE TABLE myTable (myColumn GENERATED ALWAYS AS (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Generated: token.New(1, 32, 31, 9, token.KeywordGenerated, "GENERATED"),
									Always:    token.New(1, 42, 41, 6, token.KeywordAlways, "ALWAYS"),
									As:        token.New(1, 49, 48, 2, token.KeywordAs, "AS"),
									LeftParen: token.New(1, 52, 51, 1, token.Delimiter, "("),
									Expr: &ast.Expr{
										LiteralValue: token.New(1, 53, 52, 6, token.Literal, "myExpr"),
									},
									RightParen: token.New(1, 59, 58, 1, token.Delimiter, ")"),
								},
							},
						},
					},
					RightParen: token.New(1, 60, 59, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint AS and STORED`,
			"CREATE TABLE myTable (myColumn AS (myExpr) STORED)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									As:        token.New(1, 32, 31, 2, token.KeywordAs, "AS"),
									LeftParen: token.New(1, 35, 34, 1, token.Delimiter, "("),
									Expr: &ast.Expr{
										LiteralValue: token.New(1, 36, 35, 6, token.Literal, "myExpr"),
									},
									RightParen: token.New(1, 42, 41, 1, token.Delimiter, ")"),
									Stored:     token.New(1, 44, 43, 6, token.KeywordStored, "STORED"),
								},
							},
						},
					},
					RightParen: token.New(1, 50, 49, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint AS and VIRTUAL`,
			"CREATE TABLE myTable (myColumn AS (myExpr) VIRTUAL)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									As:        token.New(1, 32, 31, 2, token.KeywordAs, "AS"),
									LeftParen: token.New(1, 35, 34, 1, token.Delimiter, "("),
									Expr: &ast.Expr{
										LiteralValue: token.New(1, 36, 35, 6, token.Literal, "myExpr"),
									},
									RightParen: token.New(1, 42, 41, 1, token.Delimiter, ")"),
									Virtual:    token.New(1, 44, 43, 7, token.KeywordVirtual, "VIRTUAL"),
								},
							},
						},
					},
					RightParen: token.New(1, 51, 50, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint DEFAULT and expr`,
			"CREATE TABLE myTable (myColumn DEFAULT (myExpr))",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Default:   token.New(1, 32, 31, 7, token.KeywordDefault, "DEFAULT"),
									LeftParen: token.New(1, 40, 39, 1, token.Delimiter, "("),
									Expr: &ast.Expr{
										LiteralValue: token.New(1, 41, 40, 6, token.Literal, "myExpr"),
									},
									RightParen: token.New(1, 47, 46, 1, token.Delimiter, ")"),
								},
							},
						},
					},
					RightParen: token.New(1, 48, 47, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint DEFAULT and positive signed number 1`,
			"CREATE TABLE myTable (myColumn DEFAULT +91)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Default: token.New(1, 32, 31, 7, token.KeywordDefault, "DEFAULT"),
									SignedNumber: &ast.SignedNumber{
										Sign:           token.New(1, 40, 39, 1, token.UnaryOperator, "+"),
										NumericLiteral: token.New(1, 41, 40, 2, token.LiteralNumeric, "91"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 43, 42, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint DEFAULT and negative signed number 1`,
			"CREATE TABLE myTable (myColumn DEFAULT -91)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Default: token.New(1, 32, 31, 7, token.KeywordDefault, "DEFAULT"),
									SignedNumber: &ast.SignedNumber{
										Sign:           token.New(1, 40, 39, 1, token.UnaryOperator, "-"),
										NumericLiteral: token.New(1, 41, 40, 2, token.LiteralNumeric, "91"),
									},
								},
							},
						},
					},
					RightParen: token.New(1, 43, 42, 1, token.Delimiter, ")"),
				},
			},
		},
		{
			`CREATE TABLE with single column-def with column constraint DEFAULT and negative signed number 1`,
			"CREATE TABLE myTable (myColumn DEFAULT myLiteral)",
			&ast.SQLStmt{
				CreateTableStmt: &ast.CreateTableStmt{
					Create:    token.New(1, 1, 0, 6, token.KeywordCreate, "CREATE"),
					Table:     token.New(1, 8, 7, 5, token.KeywordTable, "TABLE"),
					TableName: token.New(1, 14, 13, 7, token.Literal, "myTable"),
					LeftParen: token.New(1, 22, 21, 1, token.Delimiter, "("),
					ColumnDef: []*ast.ColumnDef{
						&ast.ColumnDef{
							ColumnName: token.New(1, 23, 22, 8, token.Literal, "myColumn"),
							ColumnConstraint: []*ast.ColumnConstraint{
								&ast.ColumnConstraint{
									Default:      token.New(1, 32, 31, 7, token.KeywordDefault, "DEFAULT"),
									LiteralValue: token.New(1, 40, 39, 9, token.Literal, "myLiteral"),
								},
							},
						},
					},
					RightParen: token.New(1, 49, 48, 1, token.Delimiter, ")"),
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
