// +build gofuzz

package parser

import (
	"time"

	"github.com/tomarrell/lbadd/internal/parser/ast"
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
	parser := New(input)
stmts:
	for {
		res := make(chan result)
		go waitForParseResult(parser, res)
		select {
		case <-time.After(5 * time.Second):
			panic("timeout after 5s")
		case result := <-res:
			if !result.ok {
				break stmts
			}
			if len(result.errs) != 0 {
				return DataNotInteresting
			}
			if result.stmt == nil {
				panic("ok, no errors, but also no statement")
			}
		}
	}
	return DataInteresting
}

type result struct {
	stmt *ast.SQLStmt
	errs []error
	ok   bool
}

func waitForParseResult(parser Parser, ch chan<- result) {
	stmt, errs, ok := parser.Next()
	ch <- result{stmt, errs, ok}
}
