package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_Is(t *testing.T) {
	assert := assert.New(t)

	v := NewString("foobar")
	assert.True(v.Is(String)) // Is must yield the same result as .Type() ==
	assert.True(v.Is(v.Type()))
	assert.Equal(String, v.Type())
}
