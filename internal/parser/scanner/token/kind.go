// Code generated. DO NOT EDIT.
package token

// Kind represents a token kind/type in the SQL scanner.
type Kind uint

const (
	Unknown Kind = iota
	EOF
	Whitespace

	// Comment token kinds (resolves issue #200 -- SQL comments per SQLite spec)
	// https://www.sqlite.org/lang_comment.html
	SingleLineComment // -- comment until end of line
	MultiLineComment  // /* comment */

	// Literals
	LiteralInteger
	LiteralFloat
	LiteralString

	// Identifiers
	Identifier

	// Keywords (subset — extend as grammar requires)
	KeywordSelect
	KeywordFrom
	KeywordWhere
	KeywordInsert
	KeywordInto
	KeywordValues
	KeywordCreate
	KeywordTable
	KeywordDrop
	KeywordUpdate
	KeywordSet
	KeywordDelete
	KeywordAnd
	KeywordOr
	KeywordNot
	KeywordNull
	KeywordIs
	KeywordIn
	KeywordLike
	KeywordBetween
	KeywordOrder
	KeywordBy
	KeywordGroup
	KeywordHaving
	KeywordLimit
	KeywordOffset
	KeywordDistinct
	KeywordAll
	KeywordAs
	KeywordOn
	KeywordJoin
	KeywordInner
	KeywordLeft
	KeywordRight
	KeywordFull
	KeywordOuter
	KeywordCross
	KeywordNatural
	KeywordUnion
	KeywordIntersect
	KeywordExcept
	KeywordCase
	KeywordWhen
	KeywordThen
	KeywordElse
	KeywordEnd
	KeywordIf
	KeywordExists
	KeywordPrimary
	KeywordKey
	KeywordForeign
	KeywordReferences
	KeywordUnique
	KeywordCheck
	KeywordDefault
	KeywordConstraint
	KeywordIndex
	KeywordView
	KeywordTrigger
	KeywordBegin
	KeywordCommit
	KeywordRollback
	KeywordTransaction
	KeywordSavepoint
	KeywordRelease
	KeywordReindex
	KeywordVacuum
	KeywordAnalyze
	KeywordExplain
	KeywordPlan
	KeywordPragma
	KeywordAttach
	KeywordDetach
	KeywordDatabase
	KeywordTemp
	KeywordTemporary
	KeywordWith
	KeywordRecursive
	KeywordReplace
	KeywordConflict
	KeywordFail
	KeywordIgnore
	KeywordAbort
	KeywordRollbackTo
	KeywordDeferrable
	KeywordInitially
	KeywordDeferred
	KeywordImmediate
	KeywordExclusive
	KeywordMatch
	KeywordGlob
	KeywordRegexp
	KeywordEscape
	KeywordOf
	KeywordRaise
	KeywordBefore
	KeywordAfter
	KeywordInstead
	KeywordFor
	KeywordEach
	KeywordRow
	KeywordStatement
	KeywordNew
	KeywordOld
	KeywordCast
	KeywordColumn
	KeywordRowid
	KeywordVirtual
	KeywordUsing
	KeywordCurrentTime
	KeywordCurrentDate
	KeywordCurrentTimestamp

	// Operators and punctuation
	Plus
	Minus
	Asterisk
	Slash
	Percent
	Equals
	NotEquals
	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
	LeftParen
	RightParen
	Comma
	Semicolon
	Dot
	BitwiseAnd
	BitwiseOr
	BitwiseNot
	BitwiseShiftLeft
	BitwiseShiftRight
	Concatenation
)
