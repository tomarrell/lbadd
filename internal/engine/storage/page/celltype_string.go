// Code generated by "stringer -type=CellType"; DO NOT EDIT.

package page

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CellTypeUnknown-0]
	_ = x[CellTypeRecord-1]
	_ = x[CellTypePointer-2]
}

const _CellType_name = "CellTypeUnknownCellTypeRecordCellTypePointer"

var _CellType_index = [...]uint8{0, 15, 29, 44}

func (i CellType) String() string {
	if i >= CellType(len(_CellType_index)-1) {
		return "CellType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CellType_name[_CellType_index[i]:_CellType_index[i+1]]
}
