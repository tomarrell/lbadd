package test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler"
	"github.com/tomarrell/lbadd/internal/engine"
	"github.com/tomarrell/lbadd/internal/engine/storage"
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

		// try to parse the input
		p := parser.New(content)
		stmt, errs, ok := p.Next()
		if !ok || len(errs) != 0 {
			return
		}

		// compile the statement
		c := compiler.New()
		cmd, err := c.Compile(stmt)
		if err != nil {
			return
		}

		// create a new im-memory db file if none is set
		fs := afero.NewMemMapFs()
		f, err := fs.Create("mydbfile")
		assert.NoError(err)

		dbFile, err := storage.Create(f)
		assert.NoError(err)
		defer func() { _ = dbFile.Close() }()

		// fire up the engine
		e, err := engine.New(dbFile)
		assert.NoError(err)

		result, err := e.Evaluate(cmd)
		if err != nil {
			return
		}
		_ = result
	}
}
