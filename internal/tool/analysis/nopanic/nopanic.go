// Package nopanic implements an analyzer that checks if somewhere in the
// source, there is a panic.
package nopanic

import (
	"fmt"
	"go/ast"
	"go/token"

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
		(*ast.File)(nil),
	}, func(n ast.Node) {
		file := n.(*ast.File)
		checkPanicInMainPkg(file, pass)
	})

	// If no Panic inside the given package, don't run other panic analyzer
	if countNumberofPanic == 0 {
		return nil, nil
	}

	})
	return nil, nil
}

func checkPanicInMainPkg(file *ast.File, pass *analysis.Pass) {
	unresolvedIdents := file.Unresolved
	panicsInPackage := make([]token.Pos, 0, 10)
	recoversInPackage := make([]token.Pos, 0, 10)
	pkgName := file.Name.Name
	isPkgMain := false
	if pkgName == "main" {
		isPkgMain = true
	}
	for _, v := range unresolvedIdents {
		if v.Name == "panic" {
			panicsInPackage = append(panicsInPackage, v.Pos())
		}
		if v.Name == "recover" {
			recoversInPackage = append(recoversInPackage, v.Pos())
		}
	}

	if isPkgMain {
		for _, pos := range panicsInPackage {
			pass.Reportf(pos, "panic is disallowed inside main Package")
		}
	}
	if len(panicsInPackage) > 0 && len(recoversInPackage) == 0 {
		for _, pos := range panicsInPackage {
			pass.Reportf(pos, "panic is disallowed without recover")
		}
	}
}


func callExprIsPanic(call *ast.CallExpr) bool {
	ident, ok := call.Fun.(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == "panic"
}
