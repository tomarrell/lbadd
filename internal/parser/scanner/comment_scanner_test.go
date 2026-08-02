package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func TestTryScanComment_SingleLine(t *testing.T) {
	src := []rune("-- this is a comment\nSELECT 1")
	kind, end, ok := TryScanComment(src, 0)
	assert.True(t, ok)
	assert.Equal(t, token.SingleLineComment, kind)
	// end should point to the newline (not included in comment)
	assert.Equal(t, 20, end)
	assert.Equal(t, '\n', src[end])
}

func TestTryScanComment_SingleLine_EOF(t *testing.T) {
	src := []rune("-- comment with no newline")
	kind, end, ok := TryScanComment(src, 0)
	assert.True(t, ok)
	assert.Equal(t, token.SingleLineComment, kind)
	assert.Equal(t, len(src), end)
}

func TestTryScanComment_MultiLine(t *testing.T) {
	src := []rune("/* multi\nline */ SELECT")
	kind, end, ok := TryScanComment(src, 0)
	assert.True(t, ok)
	assert.Equal(t, token.MultiLineComment, kind)
	// end points just past the closing */
	assert.Equal(t, 16, end)
	assert.Equal(t, ' ', src[end])
}

func TestTryScanComment_MultiLine_Unterminated(t *testing.T) {
	src := []rune("/* unterminated comment")
	kind, _, ok := TryScanComment(src, 0)
	assert.True(t, ok)
	assert.Equal(t, token.Unknown, kind)
}

func TestTryScanComment_NotAComment(t *testing.T) {
	src := []rune("SELECT 1")
	_, _, ok := TryScanComment(src, 0)
	assert.False(t, ok)
}

func TestTryScanComment_DivisionNotComment(t *testing.T) {
	// A single '/' is not a comment
	src := []rune("/2")
	_, _, ok := TryScanComment(src, 0)
	assert.False(t, ok)
}

func TestTryScanComment_MinusNotComment(t *testing.T) {
	// A single '-' is subtraction, not a comment
	src := []rune("-2")
	_, _, ok := TryScanComment(src, 0)
	assert.False(t, ok)
}
