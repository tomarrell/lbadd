// +build gofuzz

package compiler

import (
	"github.com/tomarrell/lbadd/internal/parser"
)

const (
	// DataNotInteresting indicates, that the input was not interesting, meaning
	// that the input was not valid and the parser handled detected and returned
	// an error. The input will still be added to the corpus.
	DataNotInteresting int = 0
	// DataInteresting indicates a valid parser input. The fuzzer should keep it
	// and modify it further.
	DataInteresting int = 1
	// Skip indicates, that the data must not be added to the corpus. You
	// probably shouldn't use it.
	Skip int = -1
)

func Fuzz(data []byte) int {
	input := string(data)
	parser := parser.New(input)
	compiler := New()

	stmt, errs, ok := parser.Next()
	if !ok || len(errs) != 0 {
		// no statement at all or parse errors are not interesting, we need
		// something we can compile
		return Skip
	}
	if _, _, ok = parser.Next(); ok {
		// more than one statement is valid, but not interesting
		return DataNotInteresting
	}

	cmd, err := compiler.Compile(stmt)
	if err != nil {
		// compile error, not interesting, we want something that compiles or
		// crashes the compiler
		return DataNotInteresting
	}
	if cmd == nil {
		panic("no compile errors, but also no command")
	}

	return DataInteresting
}
