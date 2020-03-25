package securefs_test

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/database/storage/securefs"
)

func mustRead(t *testing.T, r io.Reader) []byte {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		assert.NoError(t, err)
	}
	return data
}

func TestSecureFs_FileOperations(t *testing.T) {
	assert := assert.New(t)

	underlying := afero.NewMemMapFs()
	fs := securefs.New(underlying)

	filename := "myfile.dat"
	content := "hello, world!"
	modContent := "hello, World!"

	file, err := fs.Create(filename)
	assert.NoError(err)
	assert.Equal(filename, file.Name())

	n, err := file.WriteString(content)
	assert.NoError(err)
	assert.Equal(len(content), n)

	underlyingFile, err := underlying.Open(filename)
	assert.NoError(err)
	assert.Equal(content, string(mustRead(t, underlyingFile)))
	assert.NoError(underlyingFile.Close())

	n, err = file.WriteAt([]byte("W"), 7)
	assert.Equal(1, n)
	assert.NoError(err)

	underlyingFile, err = underlying.Open(filename)
	assert.NoError(err)
	assert.Equal(modContent, string(mustRead(t, underlyingFile)))
	assert.NoError(underlyingFile.Close())
}
