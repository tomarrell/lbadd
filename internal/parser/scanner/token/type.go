package token

//go:generate stringer -type=Type

// Type is the type of a token.
type Type uint16

// All available token types.
const (
	Unknown Type = iota
	// Error indicates that a syntax error has been detected by the lexical
	// analyzer (scanner) and that the error should be printed. The parser also
	// should consider resetting its state.
	Error
	// EOF indicates that the lexical analyzer (scanner) has reached the end of
	// its input. After receiving this token, the parser can close the token
	// stream, as there will not be any more tokens. He also can start building
	// (if not already done) the AST, as he know knows of all tokens.
	EOF

	// StatementSeparator is the type of tokens that represent a single
	// semicolon, as a single semicolon separates multiple sql statements.
	StatementSeparator

	KeywordAbort
	KeywordAction
	KeywordAdd
	KeywordAfter
	KeywordAll
	KeywordAlter
	KeywordAlways
	KeywordAnalyze
	KeywordAnd
	KeywordAs
	KeywordAsc
	KeywordAttach
	KeywordAutoincrement
	KeywordBefore
	KeywordBegin
	KeywordBetween
	KeywordBy
	KeywordCascade
	KeywordCase
	KeywordCast
	KeywordCheck
	KeywordCollate
	KeywordColumn
	KeywordCommit
	KeywordConflict
	KeywordConstraint
	KeywordCreate
	KeywordCross
	KeywordCurrent
	KeywordCurrentDate
	KeywordCurrentTime
	KeywordCurrentTimestamp
	KeywordDatabase
	KeywordDefault
	KeywordDeferrable
	KeywordDeferred
	KeywordDelete
	KeywordDesc
	KeywordDetach
	KeywordDistinct
	KeywordDo
	KeywordDrop
	KeywordEach
	KeywordElse
	KeywordEnd
	KeywordEscape
	KeywordExcept
	KeywordExclude
	KeywordExclusive
	KeywordExists
	KeywordExplain
	KeywordFail
	KeywordFilter
	KeywordFirst
	KeywordFollowing
	KeywordFor
	KeywordForeign
	KeywordFrom
	KeywordFull
	KeywordGenerated
	KeywordGlob
	KeywordGroup
	KeywordGroups
	KeywordHaving
	KeywordIf
	KeywordIgnore
	KeywordImmediate
	KeywordIn
	KeywordIndex
	KeywordIndexed
	KeywordInitially
	KeywordInner
	KeywordInsert
	KeywordInstead
	KeywordIntersect
	KeywordInto
	KeywordIs
	KeywordIsnull
	KeywordJoin
	KeywordKey
	KeywordLast
	KeywordLeft
	KeywordLike
	KeywordLimit
	KeywordMatch
	KeywordNatural
	KeywordNo
	KeywordNot
	KeywordNothing
	KeywordNotnull
	KeywordNull
	KeywordNulls
	KeywordOf
	KeywordOffset
	KeywordOn
	KeywordOr
	KeywordOrder
	KeywordOthers
	KeywordOuter
	KeywordOver
	KeywordPartition
	KeywordPlan
	KeywordPragma
	KeywordPreceding
	KeywordPrimary
	KeywordQuery
	KeywordRaise
	KeywordRange
	KeywordRecursive
	KeywordReferences
	KeywordRegexp
	KeywordReIndex
	KeywordRelease
	KeywordRename
	KeywordReplace
	KeywordRestrict
	KeywordRight
	KeywordRollback
	KeywordRow
	KeywordRows
	KeywordSavepoint
	KeywordSelect
	KeywordSet
	KeywordStored
	KeywordTable
	KeywordTemp
	KeywordTemporary
	KeywordThen
	KeywordTies
	KeywordTo
	KeywordTransaction
	KeywordTrigger
	KeywordUnbounded
	KeywordUnion
	KeywordUnique
	KeywordUpdate
	KeywordUsing
	KeywordVacuum
	KeywordValues
	KeywordView
	KeywordVirtual
	KeywordWhen
	KeywordWhere
	KeywordWindow
	KeywordWith
	KeywordWithout

	Literal
	LiteralNumeric

	UnaryOperator
	BinaryOperator
	Delimiter
)
