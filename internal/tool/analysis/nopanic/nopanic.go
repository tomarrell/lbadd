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

	inspect.Preorder([]ast.Node{
		(*ast.FuncDecl)(nil),
	}, func(n ast.Node) {
		fun := n.(*ast.FuncDecl)
		checkPanicInRecoverAndRecoverWithoutDefer(fun.Body, pass)
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

func checkPanicInRecoverAndRecoverWithoutDefer(block *ast.BlockStmt, pass *analysis.Pass) {
	list := block.List
	for _, v := range list {
		switch v.(type) {
		case *ast.IfStmt:
			if cnd, ok := v.(*ast.IfStmt); ok {
				if cnd.Init != nil {
					initStmt := cnd.Init.(*ast.AssignStmt)
					rhsExp := initStmt.Rhs
					if callExpr, ok := rhsExp[0].(*ast.CallExpr); ok {
						if exp, ok := callExpr.Fun.(*ast.Ident); ok {
							if exp.Name == "recover" {
								pass.Reportf(initStmt.TokPos, "recover is disallowed without defer")
							}
						}
					}
				}
			}
		case *ast.DeferStmt:
			dfrStmt := v.(*ast.DeferStmt)
			callExpr := dfrStmt.Call
			internalfun := callExpr.Fun
			blkStmt, ok := internalfun.(*ast.FuncLit)
			if !ok {
				continue
			}
			deferBlkList := blkStmt.Body.List

			for _, k := range deferBlkList {
				switch k.(type) {
				case *ast.IfStmt:
					if cnd, ok := k.(*ast.IfStmt); ok {
						if cnd.Init != nil {
							initStmt := cnd.Init.(*ast.AssignStmt)
							rhsExp := initStmt.Rhs
							callExpr := rhsExp[0].(*ast.CallExpr)
							exp, ok := callExpr.Fun.(*ast.Ident)
							if !ok {
								continue
							}
							if exp.Name == "recover" {
								ifBodyList := cnd.Body.List
								for _, b := range ifBodyList {
									switch b.(type) {
									case *ast.ExprStmt:
										if exprStmt, ok := b.(*ast.ExprStmt); ok {
											callExpr := exprStmt.X.(*ast.CallExpr)
											if exp, ok := callExpr.Fun.(*ast.Ident); ok {
												if exp.Name == "panic" {
													fmt.Printf("\n typ2 inside defer report")
													pass.Reportf(exp.NamePos, "panic is not allowed inside recover")
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
