// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package message

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KindUnknown-0]
	_ = x[KindTestMessage-1]
	_ = x[KindAppendEntriesRequest-2]
	_ = x[KindAppendEntriesResponse-3]
	_ = x[KindFollowerLocationListRequest-4]
	_ = x[KindFollowerLocationListResponse-5]
	_ = x[KindLeaderLocationRequest-6]
	_ = x[KindLeaderLocationResponse-7]
	_ = x[KindRequestVoteRequest-8]
	_ = x[KindRequestVoteResponse-9]
	_ = x[KindLogAppendRequest-10]
	_ = x[KindCommand-11]
	_ = x[KindCommandScan-12]
	_ = x[KindCommandSelect-13]
	_ = x[KindCommandProject-14]
	_ = x[KindCommandDelete-15]
	_ = x[KindCommandDrop-16]
	_ = x[KindCommandUpdate-17]
	_ = x[KindCommandJoin-18]
	_ = x[KindCommandLimit-19]
	_ = x[KindCommandInsert-20]
	_ = x[KindExpr-21]
	_ = x[KindLiteralExpr-22]
	_ = x[KindConstantBooleanExpr-23]
	_ = x[KindUnaryExpr-24]
	_ = x[KindBinaryExpr-25]
	_ = x[KindFunctionExpr-26]
	_ = x[KindEqualityExpr-27]
	_ = x[KindRangeExpr-28]
}

const _Kind_name = "KindUnknownKindTestMessageKindAppendEntriesRequestKindAppendEntriesResponseKindFollowerLocationListRequestKindFollowerLocationListResponseKindLeaderLocationRequestKindLeaderLocationResponseKindRequestVoteRequestKindRequestVoteResponseKindLogAppendRequestKindCommandKindCommandScanKindCommandSelectKindCommandProjectKindCommandDeleteKindCommandDropKindCommandUpdateKindCommandJoinKindCommandLimitKindCommandInsertKindExprKindLiteralExprKindConstantBooleanExprKindUnaryExprKindBinaryExprKindFunctionExprKindEqualityExprKindRangeExpr"

var _Kind_index = [...]uint16{0, 11, 26, 50, 75, 106, 138, 163, 189, 211, 234, 254, 265, 280, 297, 315, 332, 347, 364, 379, 395, 412, 420, 435, 458, 471, 485, 501, 517, 530}

func (i Kind) String() string {
	if i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
