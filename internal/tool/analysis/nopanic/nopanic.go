// Package nopanic implements an analyzer that checks if somewhere in the
// source, there is a panic.
package nopanic

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer implements the analyzer that checks for panics.
var Analyzer = &analysis.Analyzer{
	Name: "nopanic",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// Doc is the documentation string that is shown on the command line if help is
// requested.
const Doc = "check if there is any panic in the code"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{
		(*ast.CallExpr)(nil),
	}, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		if callExprIsPanic(call) {
			pass.Reportf(call.Pos(), "panic is disallowed in this location")
		}
	})
	return nil, nil
}

func callExprIsPanic(call *ast.CallExpr) bool {
	ident, ok := call.Fun.(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == "panic"
}
