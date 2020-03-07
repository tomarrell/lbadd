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
		Query string
		Stmt  *ast.SQLStmt
	}{
		{
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
			"ATTACH DATABASE bro AS newBro",
			&ast.SQLStmt{
				AttachStmt: &ast.AttachStmt{
					Attach:   token.New(1, 1, 0, 6, token.KeywordAttach, "ATTACH"),
					Database: token.New(1, 8, 7, 8, token.KeywordDatabase, "DATABASE"),
					Expr: &ast.Expr{
						LiteralValue: token.New(1, 1, 1, 1, token.Literal, "bro"),
					},
					As:         token.New(1, 1, 1, 1, token.KeywordAs, "AS"),
					SchemaName: token.New(1, 1, 1, 1, token.Literal, "newBro"),
				},
			},
		},
	}
	for _, input := range inputs {
		t.Run(input.Query[0:11], func(t *testing.T) {
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
