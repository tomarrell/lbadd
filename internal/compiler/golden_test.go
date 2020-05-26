package compiler

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser"
)

var (
	record = flag.Bool("record", false, "record golden tests")
)

func TestMain(m *testing.M) {
	flag.Parse()

	os.Exit(m.Run())
}

func RunGolden(t *testing.T, input, testName string) {
	t.Run(testName, func(t *testing.T) {
		t.Helper()
		runGolden(t, input)
	})
}

func runGolden(t *testing.T, input string) {
	t.Helper()
	assert := assert.New(t)

	c := &simpleCompiler{}
	p := parser.New(input)
	stmt, errs, ok := p.Next()
	assert.Len(errs, 0)
	assert.True(ok)

	got, err := c.Compile(stmt)
	assert.NoError(err)

	gotString := got.String()
	testFilePath := "testdata/golden/" + filepath.Base(t.Name()) + ".golden"

	if *record {
		t.Logf("overwriting golden file %v", testFilePath)
		err := os.MkdirAll(filepath.Dir(testFilePath), 0777)
		assert.NoError(err)
		err = ioutil.WriteFile(testFilePath, []byte(gotString), 0666)
		assert.NoError(err)
		t.Fail()
	} else {
		data, err := ioutil.ReadFile(testFilePath)
		assert.NoError(err)
		assert.Equal(string(data), gotString)
	}
}
