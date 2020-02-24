package main

import (
	"github.com/tomarrell/lbadd/internal/tool/analysis/nopanic"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	// write an analyzer in a sub-package of this package and add it here as an
	// argument in multichecker.Main(...).
	multichecker.Main(
		nopanic.Analyzer,
	)
}
