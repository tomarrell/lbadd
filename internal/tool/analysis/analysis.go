package main

import (
	"github.com/tomarrell/lbadd/internal/tool/analysis/ctxfunc"
	"github.com/tomarrell/lbadd/internal/tool/analysis/nopanic"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/nilness"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unsafeptr"
)

func main() {
	// write an analyzer in a sub-package of this package and add it here as an
	// argument in multichecker.Main(...).
	multichecker.Main(
		nopanic.Analyzer,
		ctxfunc.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		copylock.Analyzer,
		errorsas.Analyzer,
		lostcancel.Analyzer,
		loopclosure.Analyzer,
		nilfunc.Analyzer,
		nilness.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		stdmethods.Analyzer,
		tests.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
	)
}
