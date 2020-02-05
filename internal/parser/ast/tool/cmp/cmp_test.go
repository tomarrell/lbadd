package cmp

import (
	"flag"
	"os"
	"testing"

	"github.com/TimSatke/golden"
	"github.com/davecgh/go-spew/spew"
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var (
	update bool
)

func TestMain(m *testing.M) {
	flag.BoolVar(&update, "update", false, "enable to record tests and write results to disk as base for future comparisons")
	flag.Parse()

	os.Exit(m.Run())
}

func TestCompareAST(t *testing.T) {
	left := &ast.SQLStmt{
		AlterTableStmt: &ast.AlterTableStmt{
			Alter:        token.New(1, 1, 0, 5, token.KeywordAlter, "ALTER"),
			Table:        token.New(1, 7, 6, 5, token.KeywordTable, "TABLE"),
			TableName:    token.New(1, 13, 12, 5, token.Literal, "users"),
			Rename:       token.New(1, 19, 18, 6, token.KeywordRename, "RENAME"),
			To:           token.New(1, 26, 25, 2, token.KeywordTo, "TO"),
			NewTableName: token.New(1, 29, 28, 6, token.Literal, "admins"),
		},
	}
	right := &ast.SQLStmt{
		AlterTableStmt: &ast.AlterTableStmt{
			Alter:        token.New(1, 1, 0, 5, token.KeywordAlter, "alter"),
			Table:        token.New(1, 7, 6, 5, token.KeywordTable, "table"),
			TableName:    token.New(1, 12, 11, 5, token.Literal, "users"),
			Rename:       token.New(1, 19, 18, 6, token.KeywordRename, "rename"),
			To:           token.New(1, 26, 25, 2, token.KeywordTo, "to"),
			NewTableName: token.New(1, 29, 28, 6, token.Literal, "admins"),
		},
	}
	output := spew.Sdump(CompareAST(left, right))

	g := golden.New(t)
	g.ShouldUpdate = update
	g.Assert(t.Name(), []byte(output))
}
