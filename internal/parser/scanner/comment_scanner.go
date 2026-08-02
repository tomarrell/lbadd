package scanner

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// TryScanComment attempts to scan a SQL comment starting at the current
// position of src.  It returns (tok, true) on success and (nil, false) when
// the current position is not the start of a comment.
//
// Supported forms (per https://www.sqlite.org/lang_comment.html):
//
//	-- text until end of line (or EOF)
//	/* text, may span lines, ends at first */
func TryScanComment(src []rune, pos int) (kind token.Kind, end int, ok bool) {
	if pos >= len(src) {
		return 0, 0, false
	}

	// Single-line comment: starts with --
	if pos+1 < len(src) && src[pos] == '-' && src[pos+1] == '-' {
		end = pos + 2
		for end < len(src) && src[end] != '\n' {
			end++
		}
		return token.SingleLineComment, end, true
	}

	// Multi-line comment: starts with /*
	if pos+1 < len(src) && src[pos] == '/' && src[pos+1] == '*' {
		end = pos + 2
		for end+1 < len(src) {
			if src[end] == '*' && src[end+1] == '/' {
				end += 2 // consume closing */
				return token.MultiLineComment, end, true
			}
			end++
		}
		// Unterminated multi-line comment — consume to EOF and return Unknown
		return token.Unknown, len(src), true
	}

	return 0, 0, false
}
