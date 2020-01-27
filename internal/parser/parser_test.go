package parser

import (
	"flag"
	"os"
	"testing"

	"github.com/TimSatke/golden"
)

var (
	update bool
)

func TestMain(m *testing.M) {
	flag.BoolVar(&update, "update", false, "enable to record tests and write results to disk as base for future comparisons")
	flag.Parse()

	os.Exit(m.Run())
}

func TestParserGolden(t *testing.T) {
	inputs := []struct {
		Name  string
		Query string
	}{
		{"empty", ""},
	}
	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			p := New(input.Query)

			for {
				stmt, errs, ok := p.Next()
				if !ok {
					break
				}

				g := golden.New(t)
				g.ShouldUpdate = update
				g.AssertStruct(input.Name+"_ast", stmt)
				g.AssertStruct(input.Name+"_errs", errs)
			}
		})
	}
}
