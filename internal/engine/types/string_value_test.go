package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValue_Is(t *testing.T) {
	assert := assert.New(t)

	v := NewString("foobar")
	assert.True(v.Is(String))
}

func TestStringValue_Type(t *testing.T) {
	assert := assert.New(t)

	v := NewString("foobar")
	assert.Equal(String, v.Type())
}

func TestStringValue_String(t *testing.T) {
	assert := assert.New(t)

	baseStr := "foobar"
	v := NewString(baseStr)
	assert.Equal(baseStr, v.String())
}
