package scanner

import (
	"bytes"
	"math/rand"
	"time"
	"unicode"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ token.Token = (*genTok)(nil)

var rng = rand.New(rand.NewSource(time.Now().Unix()))

type genTok struct {
	offset int
	value  string
	typ    token.Type
}

func (t genTok) Line() int {
	return 1
}

func (t genTok) Col() int {
	return t.offset + 1
}

func (t genTok) Offset() int {
	return t.offset
}

func (t genTok) Length() int {
	return len(t.value)
}

func (t genTok) Type() token.Type {
	return t.typ
}

func (t genTok) Value() string {
	return t.value
}

func generateEOF(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.EOF,
	}
}
func generateStatementSeparator(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.StatementSeparator,
		value:  ";",
	}
}
func generateKeywordAbort(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAbort,
		value:  caseShuffle("Abort"),
	}
}
func generateKeywordAction(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAction,
		value:  caseShuffle("Action"),
	}
}
func generateKeywordAdd(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAdd,
		value:  caseShuffle("Add"),
	}
}
func generateKeywordAfter(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAfter,
		value:  caseShuffle("After"),
	}
}
func generateKeywordAll(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAll,
		value:  caseShuffle("All"),
	}
}
func generateKeywordAlter(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAlter,
		value:  caseShuffle("Alter"),
	}
}
func generateKeywordAnalyze(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAnalyze,
		value:  caseShuffle("Analyze"),
	}
}
func generateKeywordAnd(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAnd,
		value:  caseShuffle("And"),
	}
}
func generateKeywordAs(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAs,
		value:  caseShuffle("As"),
	}
}
func generateKeywordAsc(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAsc,
		value:  caseShuffle("Asc"),
	}
}
func generateKeywordAttach(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAttach,
		value:  caseShuffle("Attach"),
	}
}
func generateKeywordAutoincrement(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordAutoincrement,
		value:  caseShuffle("Autoincrement"),
	}
}
func generateKeywordBefore(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordBefore,
		value:  caseShuffle("Before"),
	}
}
func generateKeywordBegin(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordBegin,
		value:  caseShuffle("Begin"),
	}
}
func generateKeywordBetween(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordBetween,
		value:  caseShuffle("Between"),
	}
}
func generateKeywordBy(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordBy,
		value:  caseShuffle("By"),
	}
}
func generateKeywordCascade(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCascade,
		value:  caseShuffle("Cascade"),
	}
}
func generateKeywordCase(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCase,
		value:  caseShuffle("Case"),
	}
}
func generateKeywordCast(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCast,
		value:  caseShuffle("Cast"),
	}
}
func generateKeywordCheck(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCheck,
		value:  caseShuffle("Check"),
	}
}
func generateKeywordCollate(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCollate,
		value:  caseShuffle("Collate"),
	}
}
func generateKeywordColumn(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordColumn,
		value:  caseShuffle("Column"),
	}
}
func generateKeywordCommit(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCommit,
		value:  caseShuffle("Commit"),
	}
}
func generateKeywordConflict(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordConflict,
		value:  caseShuffle("Conflict"),
	}
}
func generateKeywordConstraint(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordConstraint,
		value:  caseShuffle("Constraint"),
	}
}
func generateKeywordCreate(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCreate,
		value:  caseShuffle("Create"),
	}
}
func generateKeywordCross(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCross,
		value:  caseShuffle("Cross"),
	}
}
func generateKeywordCurrent(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCurrent,
		value:  caseShuffle("Current"),
	}
}
func generateKeywordCurrentDate(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCurrentDate,
		value:  caseShuffle("CurrentDate"),
	}
}
func generateKeywordCurrentTime(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCurrentTime,
		value:  caseShuffle("CurrentTime"),
	}
}
func generateKeywordCurrentTimestamp(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordCurrentTimestamp,
		value:  caseShuffle("CurrentTimestamp"),
	}
}
func generateKeywordDatabase(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDatabase,
		value:  caseShuffle("Database"),
	}
}
func generateKeywordDefault(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDefault,
		value:  caseShuffle("Default"),
	}
}
func generateKeywordDeferrable(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDeferrable,
		value:  caseShuffle("Deferrable"),
	}
}
func generateKeywordDeferred(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDeferred,
		value:  caseShuffle("Deferred"),
	}
}
func generateKeywordDelete(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDelete,
		value:  caseShuffle("Delete"),
	}
}
func generateKeywordDesc(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDesc,
		value:  caseShuffle("Desc"),
	}
}
func generateKeywordDetach(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDetach,
		value:  caseShuffle("Detach"),
	}
}
func generateKeywordDistinct(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDistinct,
		value:  caseShuffle("Distinct"),
	}
}
func generateKeywordDo(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDo,
		value:  caseShuffle("Do"),
	}
}
func generateKeywordDrop(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordDrop,
		value:  caseShuffle("Drop"),
	}
}
func generateKeywordEach(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordEach,
		value:  caseShuffle("Each"),
	}
}
func generateKeywordElse(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordElse,
		value:  caseShuffle("Else"),
	}
}
func generateKeywordEnd(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordEnd,
		value:  caseShuffle("End"),
	}
}
func generateKeywordEscape(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordEscape,
		value:  caseShuffle("Escape"),
	}
}
func generateKeywordExcept(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordExcept,
		value:  caseShuffle("Except"),
	}
}
func generateKeywordExclude(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordExclude,
		value:  caseShuffle("Exclude"),
	}
}
func generateKeywordExclusive(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordExclusive,
		value:  caseShuffle("Exclusive"),
	}
}
func generateKeywordExists(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordExists,
		value:  caseShuffle("Exists"),
	}
}
func generateKeywordExplain(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordExplain,
		value:  caseShuffle("Explain"),
	}
}
func generateKeywordFail(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFail,
		value:  caseShuffle("Fail"),
	}
}
func generateKeywordFilter(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFilter,
		value:  caseShuffle("Filter"),
	}
}
func generateKeywordFirst(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFirst,
		value:  caseShuffle("First"),
	}
}
func generateKeywordFollowing(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFollowing,
		value:  caseShuffle("Following"),
	}
}
func generateKeywordFor(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFor,
		value:  caseShuffle("For"),
	}
}
func generateKeywordForeign(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordForeign,
		value:  caseShuffle("Foreign"),
	}
}
func generateKeywordFrom(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFrom,
		value:  caseShuffle("From"),
	}
}
func generateKeywordFull(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordFull,
		value:  caseShuffle("Full"),
	}
}
func generateKeywordGenerated(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordGenerated,
		value:  caseShuffle("Generated"),
	}
}
func generateKeywordGlob(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordGlob,
		value:  caseShuffle("Glob"),
	}
}
func generateKeywordGroup(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordGroup,
		value:  caseShuffle("Group"),
	}
}
func generateKeywordGroups(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordGroups,
		value:  caseShuffle("Groups"),
	}
}
func generateKeywordHaving(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordHaving,
		value:  caseShuffle("Having"),
	}
}
func generateKeywordIf(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIf,
		value:  caseShuffle("If"),
	}
}
func generateKeywordIgnore(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIgnore,
		value:  caseShuffle("Ignore"),
	}
}
func generateKeywordImmediate(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordImmediate,
		value:  caseShuffle("Immediate"),
	}
}
func generateKeywordIn(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIn,
		value:  caseShuffle("In"),
	}
}
func generateKeywordIndex(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIndex,
		value:  caseShuffle("Index"),
	}
}
func generateKeywordIndexed(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIndexed,
		value:  caseShuffle("Indexed"),
	}
}
func generateKeywordInitially(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordInitially,
		value:  caseShuffle("Initially"),
	}
}
func generateKeywordInner(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordInner,
		value:  caseShuffle("Inner"),
	}
}
func generateKeywordInsert(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordInsert,
		value:  caseShuffle("Insert"),
	}
}
func generateKeywordInstead(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordInstead,
		value:  caseShuffle("Instead"),
	}
}
func generateKeywordIntersect(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIntersect,
		value:  caseShuffle("Intersect"),
	}
}
func generateKeywordInto(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordInto,
		value:  caseShuffle("Into"),
	}
}
func generateKeywordIs(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIs,
		value:  caseShuffle("Is"),
	}
}
func generateKeywordIsnull(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordIsnull,
		value:  caseShuffle("Isnull"),
	}
}
func generateKeywordJoin(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordJoin,
		value:  caseShuffle("Join"),
	}
}
func generateKeywordKey(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordKey,
		value:  caseShuffle("Key"),
	}
}
func generateKeywordLast(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordLast,
		value:  caseShuffle("Last"),
	}
}
func generateKeywordLeft(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordLeft,
		value:  caseShuffle("Left"),
	}
}
func generateKeywordLike(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordLike,
		value:  caseShuffle("Like"),
	}
}
func generateKeywordLimit(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordLimit,
		value:  caseShuffle("Limit"),
	}
}
func generateKeywordMatch(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordMatch,
		value:  caseShuffle("Match"),
	}
}
func generateKeywordNatural(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNatural,
		value:  caseShuffle("Natural"),
	}
}
func generateKeywordNo(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNo,
		value:  caseShuffle("No"),
	}
}
func generateKeywordNot(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNot,
		value:  caseShuffle("Not"),
	}
}
func generateKeywordNothing(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNothing,
		value:  caseShuffle("Nothing"),
	}
}
func generateKeywordNotnull(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNotnull,
		value:  caseShuffle("Notnull"),
	}
}
func generateKeywordNull(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNull,
		value:  caseShuffle("Null"),
	}
}
func generateKeywordNulls(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordNulls,
		value:  caseShuffle("Nulls"),
	}
}
func generateKeywordOf(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOf,
		value:  caseShuffle("Of"),
	}
}
func generateKeywordOffset(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOffset,
		value:  caseShuffle("Offset"),
	}
}
func generateKeywordOn(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOn,
		value:  caseShuffle("On"),
	}
}
func generateKeywordOr(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOr,
		value:  caseShuffle("Or"),
	}
}
func generateKeywordOrder(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOrder,
		value:  caseShuffle("Order"),
	}
}
func generateKeywordOthers(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOthers,
		value:  caseShuffle("Others"),
	}
}
func generateKeywordOuter(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOuter,
		value:  caseShuffle("Outer"),
	}
}
func generateKeywordOver(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordOver,
		value:  caseShuffle("Over"),
	}
}
func generateKeywordPartition(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordPartition,
		value:  caseShuffle("Partition"),
	}
}
func generateKeywordPlan(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordPlan,
		value:  caseShuffle("Plan"),
	}
}
func generateKeywordPragma(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordPragma,
		value:  caseShuffle("Pragma"),
	}
}
func generateKeywordPreceding(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordPreceding,
		value:  caseShuffle("Preceding"),
	}
}
func generateKeywordPrimary(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordPrimary,
		value:  caseShuffle("Primary"),
	}
}
func generateKeywordQuery(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordQuery,
		value:  caseShuffle("Query"),
	}
}
func generateKeywordRaise(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRaise,
		value:  caseShuffle("Raise"),
	}
}
func generateKeywordRange(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRange,
		value:  caseShuffle("Range"),
	}
}
func generateKeywordRecursive(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRecursive,
		value:  caseShuffle("Recursive"),
	}
}
func generateKeywordReferences(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordReferences,
		value:  caseShuffle("References"),
	}
}
func generateKeywordRegexp(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRegexp,
		value:  caseShuffle("Regexp"),
	}
}
func generateKeywordReIndex(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordReIndex,
		value:  caseShuffle("ReIndex"),
	}
}
func generateKeywordRelease(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRelease,
		value:  caseShuffle("Release"),
	}
}
func generateKeywordRename(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRename,
		value:  caseShuffle("Rename"),
	}
}
func generateKeywordReplace(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordReplace,
		value:  caseShuffle("Replace"),
	}
}
func generateKeywordRestrict(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRestrict,
		value:  caseShuffle("Restrict"),
	}
}
func generateKeywordRight(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRight,
		value:  caseShuffle("Right"),
	}
}
func generateKeywordRollback(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRollback,
		value:  caseShuffle("Rollback"),
	}
}
func generateKeywordRow(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRow,
		value:  caseShuffle("Row"),
	}
}
func generateKeywordRows(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordRows,
		value:  caseShuffle("Rows"),
	}
}
func generateKeywordSavepoint(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordSavepoint,
		value:  caseShuffle("Savepoint"),
	}
}
func generateKeywordSelect(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordSelect,
		value:  caseShuffle("Select"),
	}
}
func generateKeywordSet(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordSet,
		value:  caseShuffle("Set"),
	}
}
func generateKeywordTable(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTable,
		value:  caseShuffle("Table"),
	}
}
func generateKeywordTemp(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTemp,
		value:  caseShuffle("Temp"),
	}
}
func generateKeywordTemporary(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTemporary,
		value:  caseShuffle("Temporary"),
	}
}
func generateKeywordThen(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordThen,
		value:  caseShuffle("Then"),
	}
}
func generateKeywordTies(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTies,
		value:  caseShuffle("Ties"),
	}
}
func generateKeywordTo(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTo,
		value:  caseShuffle("To"),
	}
}
func generateKeywordTransaction(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTransaction,
		value:  caseShuffle("Transaction"),
	}
}
func generateKeywordTrigger(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordTrigger,
		value:  caseShuffle("Trigger"),
	}
}
func generateKeywordUnbounded(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordUnbounded,
		value:  caseShuffle("Unbounded"),
	}
}
func generateKeywordUnion(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordUnion,
		value:  caseShuffle("Union"),
	}
}
func generateKeywordUnique(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordUnique,
		value:  caseShuffle("Unique"),
	}
}
func generateKeywordUpdate(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordUpdate,
		value:  caseShuffle("Update"),
	}
}
func generateKeywordUsing(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordUsing,
		value:  caseShuffle("Using"),
	}
}
func generateKeywordVacuum(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordVacuum,
		value:  caseShuffle("Vacuum"),
	}
}
func generateKeywordValues(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordValues,
		value:  caseShuffle("Values"),
	}
}
func generateKeywordView(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordView,
		value:  caseShuffle("View"),
	}
}
func generateKeywordVirtual(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordVirtual,
		value:  caseShuffle("Virtual"),
	}
}
func generateKeywordWhen(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordWhen,
		value:  caseShuffle("When"),
	}
}
func generateKeywordWhere(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordWhere,
		value:  caseShuffle("Where"),
	}
}
func generateKeywordWindow(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordWindow,
		value:  caseShuffle("Window"),
	}
}
func generateKeywordWith(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordWith,
		value:  caseShuffle("With"),
	}
}
func generateKeywordWithout(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.KeywordWithout,
		value:  caseShuffle("Without"),
	}
}

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ -_=+[]{};:'\\|,<.>/?!@#$%^&*()§±`~¡™£¢∞§¶•ªº–≠§“‘…æ«≤≥÷`∑´®†¥¨ˆøπåß∂ƒ©˙∆˚¬≈ç√∫˜µ")

func generateLiteral(offset int) token.Token {
	var buf bytes.Buffer
	limit := rng.Intn(40)
	for i := 0; i < limit; i++ {
		buf.WriteRune(alphabet[rng.Intn(len(alphabet))])
	}

	return genTok{
		offset: offset,
		typ:    token.Literal,
		value:  "\"" + buf.String() + "\"",
	}
}

var unaryOperators = []string{"-", "+", "~"}

func generateUnaryOperator(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.UnaryOperator,
		value:  unaryOperators[rng.Intn(len(unaryOperators))],
	}
}

var binaryOperators = []string{"||", "*", "/", "%", "+", "-", "<<", ">>", "&", "|", "<", "<=", ">", ">=", "=", "==", "!=", "<>"}

func generateBinaryOperator(offset int) token.Token {
	return genTok{
		offset: offset,
		typ:    token.BinaryOperator,
		value:  binaryOperators[rng.Intn(len(binaryOperators))],
	}
}

func caseShuffle(s string) string {

	var buf bytes.Buffer
	for _, x := range s {
		n := rng.Intn(2)
		r := unicode.ToLower(x)
		if n == 0 {
			r = unicode.ToUpper(x)
		}
		_, _ = buf.WriteRune(r)
	}

	return buf.String()
}

func generateTokenType() token.Type {
	typ := token.Type(rng.Intn(int(token.Delimiter)))
	if typ == token.Error || typ == token.Unknown || typ == token.EOF {
		typ = token.StatementSeparator
	}
	return typ
}

func generateTokenForType(offset int, typ token.Type) token.Token {
	switch typ {
	case token.StatementSeparator:
		return generateStatementSeparator(offset)
	case token.KeywordAbort:
		return generateKeywordAbort(offset)
	case token.KeywordAction:
		return generateKeywordAction(offset)
	case token.KeywordAdd:
		return generateKeywordAdd(offset)
	case token.KeywordAfter:
		return generateKeywordAfter(offset)
	case token.KeywordAll:
		return generateKeywordAll(offset)
	case token.KeywordAlter:
		return generateKeywordAlter(offset)
	case token.KeywordAnalyze:
		return generateKeywordAnalyze(offset)
	case token.KeywordAnd:
		return generateKeywordAnd(offset)
	case token.KeywordAs:
		return generateKeywordAs(offset)
	case token.KeywordAsc:
		return generateKeywordAsc(offset)
	case token.KeywordAttach:
		return generateKeywordAttach(offset)
	case token.KeywordAutoincrement:
		return generateKeywordAutoincrement(offset)
	case token.KeywordBefore:
		return generateKeywordBefore(offset)
	case token.KeywordBegin:
		return generateKeywordBegin(offset)
	case token.KeywordBetween:
		return generateKeywordBetween(offset)
	case token.KeywordBy:
		return generateKeywordBy(offset)
	case token.KeywordCascade:
		return generateKeywordCascade(offset)
	case token.KeywordCase:
		return generateKeywordCase(offset)
	case token.KeywordCast:
		return generateKeywordCast(offset)
	case token.KeywordCheck:
		return generateKeywordCheck(offset)
	case token.KeywordCollate:
		return generateKeywordCollate(offset)
	case token.KeywordColumn:
		return generateKeywordColumn(offset)
	case token.KeywordCommit:
		return generateKeywordCommit(offset)
	case token.KeywordConflict:
		return generateKeywordConflict(offset)
	case token.KeywordConstraint:
		return generateKeywordConstraint(offset)
	case token.KeywordCreate:
		return generateKeywordCreate(offset)
	case token.KeywordCross:
		return generateKeywordCross(offset)
	case token.KeywordCurrent:
		return generateKeywordCurrent(offset)
	case token.KeywordCurrentDate:
		return generateKeywordCurrentDate(offset)
	case token.KeywordCurrentTime:
		return generateKeywordCurrentTime(offset)
	case token.KeywordCurrentTimestamp:
		return generateKeywordCurrentTimestamp(offset)
	case token.KeywordDatabase:
		return generateKeywordDatabase(offset)
	case token.KeywordDefault:
		return generateKeywordDefault(offset)
	case token.KeywordDeferrable:
		return generateKeywordDeferrable(offset)
	case token.KeywordDeferred:
		return generateKeywordDeferred(offset)
	case token.KeywordDelete:
		return generateKeywordDelete(offset)
	case token.KeywordDesc:
		return generateKeywordDesc(offset)
	case token.KeywordDetach:
		return generateKeywordDetach(offset)
	case token.KeywordDistinct:
		return generateKeywordDistinct(offset)
	case token.KeywordDo:
		return generateKeywordDo(offset)
	case token.KeywordDrop:
		return generateKeywordDrop(offset)
	case token.KeywordEach:
		return generateKeywordEach(offset)
	case token.KeywordElse:
		return generateKeywordElse(offset)
	case token.KeywordEnd:
		return generateKeywordEnd(offset)
	case token.KeywordEscape:
		return generateKeywordEscape(offset)
	case token.KeywordExcept:
		return generateKeywordExcept(offset)
	case token.KeywordExclude:
		return generateKeywordExclude(offset)
	case token.KeywordExclusive:
		return generateKeywordExclusive(offset)
	case token.KeywordExists:
		return generateKeywordExists(offset)
	case token.KeywordExplain:
		return generateKeywordExplain(offset)
	case token.KeywordFail:
		return generateKeywordFail(offset)
	case token.KeywordFilter:
		return generateKeywordFilter(offset)
	case token.KeywordFirst:
		return generateKeywordFirst(offset)
	case token.KeywordFollowing:
		return generateKeywordFollowing(offset)
	case token.KeywordFor:
		return generateKeywordFor(offset)
	case token.KeywordForeign:
		return generateKeywordForeign(offset)
	case token.KeywordFrom:
		return generateKeywordFrom(offset)
	case token.KeywordFull:
		return generateKeywordFull(offset)
	case token.KeywordGenerated:
		return generateKeywordGenerated(offset)
	case token.KeywordGlob:
		return generateKeywordGlob(offset)
	case token.KeywordGroup:
		return generateKeywordGroup(offset)
	case token.KeywordGroups:
		return generateKeywordGroups(offset)
	case token.KeywordHaving:
		return generateKeywordHaving(offset)
	case token.KeywordIf:
		return generateKeywordIf(offset)
	case token.KeywordIgnore:
		return generateKeywordIgnore(offset)
	case token.KeywordImmediate:
		return generateKeywordImmediate(offset)
	case token.KeywordIn:
		return generateKeywordIn(offset)
	case token.KeywordIndex:
		return generateKeywordIndex(offset)
	case token.KeywordIndexed:
		return generateKeywordIndexed(offset)
	case token.KeywordInitially:
		return generateKeywordInitially(offset)
	case token.KeywordInner:
		return generateKeywordInner(offset)
	case token.KeywordInsert:
		return generateKeywordInsert(offset)
	case token.KeywordInstead:
		return generateKeywordInstead(offset)
	case token.KeywordIntersect:
		return generateKeywordIntersect(offset)
	case token.KeywordInto:
		return generateKeywordInto(offset)
	case token.KeywordIs:
		return generateKeywordIs(offset)
	case token.KeywordIsnull:
		return generateKeywordIsnull(offset)
	case token.KeywordJoin:
		return generateKeywordJoin(offset)
	case token.KeywordKey:
		return generateKeywordKey(offset)
	case token.KeywordLast:
		return generateKeywordLast(offset)
	case token.KeywordLeft:
		return generateKeywordLeft(offset)
	case token.KeywordLike:
		return generateKeywordLike(offset)
	case token.KeywordLimit:
		return generateKeywordLimit(offset)
	case token.KeywordMatch:
		return generateKeywordMatch(offset)
	case token.KeywordNatural:
		return generateKeywordNatural(offset)
	case token.KeywordNo:
		return generateKeywordNo(offset)
	case token.KeywordNot:
		return generateKeywordNot(offset)
	case token.KeywordNothing:
		return generateKeywordNothing(offset)
	case token.KeywordNotnull:
		return generateKeywordNotnull(offset)
	case token.KeywordNull:
		return generateKeywordNull(offset)
	case token.KeywordNulls:
		return generateKeywordNulls(offset)
	case token.KeywordOf:
		return generateKeywordOf(offset)
	case token.KeywordOffset:
		return generateKeywordOffset(offset)
	case token.KeywordOn:
		return generateKeywordOn(offset)
	case token.KeywordOr:
		return generateKeywordOr(offset)
	case token.KeywordOrder:
		return generateKeywordOrder(offset)
	case token.KeywordOthers:
		return generateKeywordOthers(offset)
	case token.KeywordOuter:
		return generateKeywordOuter(offset)
	case token.KeywordOver:
		return generateKeywordOver(offset)
	case token.KeywordPartition:
		return generateKeywordPartition(offset)
	case token.KeywordPlan:
		return generateKeywordPlan(offset)
	case token.KeywordPragma:
		return generateKeywordPragma(offset)
	case token.KeywordPreceding:
		return generateKeywordPreceding(offset)
	case token.KeywordPrimary:
		return generateKeywordPrimary(offset)
	case token.KeywordQuery:
		return generateKeywordQuery(offset)
	case token.KeywordRaise:
		return generateKeywordRaise(offset)
	case token.KeywordRange:
		return generateKeywordRange(offset)
	case token.KeywordRecursive:
		return generateKeywordRecursive(offset)
	case token.KeywordReferences:
		return generateKeywordReferences(offset)
	case token.KeywordRegexp:
		return generateKeywordRegexp(offset)
	case token.KeywordReIndex:
		return generateKeywordReIndex(offset)
	case token.KeywordRelease:
		return generateKeywordRelease(offset)
	case token.KeywordRename:
		return generateKeywordRename(offset)
	case token.KeywordReplace:
		return generateKeywordReplace(offset)
	case token.KeywordRestrict:
		return generateKeywordRestrict(offset)
	case token.KeywordRight:
		return generateKeywordRight(offset)
	case token.KeywordRollback:
		return generateKeywordRollback(offset)
	case token.KeywordRow:
		return generateKeywordRow(offset)
	case token.KeywordRows:
		return generateKeywordRows(offset)
	case token.KeywordSavepoint:
		return generateKeywordSavepoint(offset)
	case token.KeywordSelect:
		return generateKeywordSelect(offset)
	case token.KeywordSet:
		return generateKeywordSet(offset)
	case token.KeywordTable:
		return generateKeywordTable(offset)
	case token.KeywordTemp:
		return generateKeywordTemp(offset)
	case token.KeywordTemporary:
		return generateKeywordTemporary(offset)
	case token.KeywordThen:
		return generateKeywordThen(offset)
	case token.KeywordTies:
		return generateKeywordTies(offset)
	case token.KeywordTo:
		return generateKeywordTo(offset)
	case token.KeywordTransaction:
		return generateKeywordTransaction(offset)
	case token.KeywordTrigger:
		return generateKeywordTrigger(offset)
	case token.KeywordUnbounded:
		return generateKeywordUnbounded(offset)
	case token.KeywordUnion:
		return generateKeywordUnion(offset)
	case token.KeywordUnique:
		return generateKeywordUnique(offset)
	case token.KeywordUpdate:
		return generateKeywordUpdate(offset)
	case token.KeywordUsing:
		return generateKeywordUsing(offset)
	case token.KeywordVacuum:
		return generateKeywordVacuum(offset)
	case token.KeywordValues:
		return generateKeywordValues(offset)
	case token.KeywordView:
		return generateKeywordView(offset)
	case token.KeywordVirtual:
		return generateKeywordVirtual(offset)
	case token.KeywordWhen:
		return generateKeywordWhen(offset)
	case token.KeywordWhere:
		return generateKeywordWhere(offset)
	case token.KeywordWindow:
		return generateKeywordWindow(offset)
	case token.KeywordWith:
		return generateKeywordWith(offset)
	case token.KeywordWithout:
		return generateKeywordWithout(offset)
	case token.Literal:
		return generateLiteral(offset)
	case token.UnaryOperator:
		return generateUnaryOperator(offset)
	case token.BinaryOperator:
		return generateBinaryOperator(offset)
	default:
	}
	return generateStatementSeparator(offset)
}

func generateToken(offset int) token.Token {
	if rng.Intn(2) == 0 {
		return generateTokenForType(offset, token.Literal)
	}

	typ := generateTokenType()
	return generateTokenForType(offset, typ)
}

func generateScannerInputAndExpectedOutput() (scannerInput string, scannerOutput []token.Token) {

	var buf bytes.Buffer

	amountOfTokens := rng.Intn(200)
	currentOffset := 0
	for i := 0; i < amountOfTokens; i++ {
		// generate token
		tok := generateToken(currentOffset)
		currentOffset += tok.Length() - 1

		// append to results
		buf.WriteString(tok.Value())
		scannerOutput = append(scannerOutput, tok)

		// generate whitespace
		whitespaces := rng.Intn(5) + 1
		currentOffset += whitespaces
		for i := 0; i < whitespaces; i++ {
			_, _ = buf.WriteRune(' ')
		}
	}

	// EOF token
	scannerOutput = append(scannerOutput, generateEOF(currentOffset))

	scannerInput = buf.String()
	return
}
