package compiler

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser"
)

const (
	fuzzCorpusDir = "testdata/fuzz/corpus"
)

// TestFuzzCorpus runs the current fuzzing corpus.
func TestFuzzCorpus(t *testing.T) {
	assert := assert.New(t)

	corpusFiles, err := filepath.Glob(filepath.Join(fuzzCorpusDir, "*"))
	assert.NoError(err)

	for _, corpusFile := range corpusFiles {
		t.Run(filepath.Base(corpusFile), _TestCorpusFile(corpusFile))
	}
}

func _TestCorpusFile(file string) func(*testing.T) {
	return func(t *testing.T) {
		assert := assert.New(t)

		data, err := ioutil.ReadFile(file)
		assert.NoError(err)
		content := string(data)
		parser := parser.New(content)
		compiler := New()

		stmt, errs, ok := parser.Next()
		assert.True(ok)
		if !ok || len(errs) != 0 {
			return
		}

		cmd, err := compiler.Compile(stmt)
		// We don't care about what the compilation result is or if there was an
		// error. Fuzzing just ensures, that the compiler doesn't panic for any
		// input. We just need to make sure, that the compilation result and
		// compilation errors are mutually exclusive.
		if err == nil {
			assert.NotNil(cmd)
		} else {
			assert.Nil(cmd)
		}
	}
}
