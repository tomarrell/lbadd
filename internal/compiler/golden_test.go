package compiler

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomarrell/lbadd/internal/parser"
)

var (
	update = flag.Bool("update", false, "update fixture files")
)

func TestMain(m *testing.M) {
	flag.Parse()

	os.Exit(m.Run())
}

func RunGolden(t *testing.T, input string) {
	t.Helper()
	t.Run("", func(t *testing.T) {
		t.Helper()
		runGolden(t, input)
	})
}

func runGolden(t *testing.T, input string) {
	t.Logf("testcase:\nname: %v\ninput: \"%v\"", t.Name(), input)

	require := require.New(t)

	c := &simpleCompiler{}
	p := parser.New(input)
	stmt, errs, ok := p.Next()
	require.Len(errs, 0)
	require.True(ok, "expected at least one statement that can be parsed")

	got, err := c.Compile(stmt)
	require.NoError(err)

	gotGoString := fmt.Sprintf("%#v", got)
	gotString := got.String()
	gotFull := gotGoString + "\n\nString:\n" + gotString

	testFilePath := "testdata/" + t.Name() + ".golden"

	if *update {
		t.Logf("overwriting golden file %v", testFilePath)
		err := os.MkdirAll(filepath.Dir(testFilePath), 0777)
		require.NoError(err)
		err = ioutil.WriteFile(testFilePath, []byte(gotFull), 0666)
		require.NoError(err)
		t.Fail()
	} else {
		data, err := ioutil.ReadFile(testFilePath)
		require.NoError(err)
		require.Equal(string(data), gotFull)
	}
}
