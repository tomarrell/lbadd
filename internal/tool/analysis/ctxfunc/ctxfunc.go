// Package ctxfunc implements an analyzer that checks if a context argument is
// always the first parameter to a function, and that it is named 'ctx'.
package ctxfunc

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer implements the analyzer that checks all context arguments.
var Analyzer = &analysis.Analyzer{
	Name: "ctxfunc",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// Doc is the documentation string that is shown on the command line if help is
// requested.
const Doc = "check if there is any context parameter in the code, that is not the first argument to a function or that's not named 'ctx'"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{
		(*ast.FuncType)(nil),
	}, func(n ast.Node) {
		fn := n.(*ast.FuncType)
		checkFunction(fn, pass)
	})
	return nil, nil
}

func checkFunction(fn *ast.FuncType, pass *analysis.Pass) {
	foundContext := false
	for i, arg := range fn.Params.List {
		argType := pass.TypesInfo.TypeOf(arg.Type)
		if namedType, ok := argType.(*types.Named); ok {
			if namedType.String() == "context.Context" { // we found a context.Context argument
				n := len(arg.Names)
				if n < 1 {
					continue
				}
				if n > 1 || foundContext {
					pass.Reportf(arg.Pos(), "more than one context.Context argument")
					return
				}
				foundContext = true

				if i != 0 { // context.Context argument must be first argument
					// this is actually covered by go-lint
					pass.Reportf(arg.Pos(), "context.Context should be the first parameter of a function")
					return
				}

				// there is a single context.Context argument in the first
				// position, now check if it's named 'ctx'
				if arg.Names[0].String() == "_" {
					pass.Reportf(arg.Pos(), "unused context.Context argument")
					return
				} else if arg.Names[0].String() != "ctx" {
					pass.Reportf(arg.Names[0].Pos(), "context.Context argument should be named 'ctx'")
					return
				}
			}
		}
	}
}
