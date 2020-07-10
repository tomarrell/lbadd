package engine

import "github.com/tomarrell/lbadd/internal/engine/types"

//go:generate stringer -type=cmpResult

type cmpResult uint8

const (
	cmpUncomparable cmpResult = iota
	// cmpEqual is returned if left==right.
	cmpEqual
	// cmpLessThan is returned if left<right.
	cmpLessThan
	// cmpGreaterThan is returned if left>right.
	cmpGreaterThan
)

// cmp compares two values. The result is to be interpreted as R(left, right) or
// left~right, meaning if e.g. cmpLessThan is returned, it is to be understood
// as left<right. If left and right cannot be compared, e.g. because they have
// different types, cmpUncomparable will be returned.
func (e Engine) cmp(left, right types.Value) cmpResult {
	defer e.profiler.Enter(EvtCompare).Exit()

	// types must be equal
	if !right.Is(left.Type()) {
		return cmpUncomparable
	}
	comparator, ok := left.Type().(types.Comparator)
	if !ok {
		return cmpUncomparable
	}
	res, err := comparator.Compare(left, right)
	if err != nil {
		return cmpUncomparable
	}
	switch res {
	case -1:
		return cmpLessThan
	case 0:
		return cmpEqual
	case 1:
		return cmpGreaterThan
	}
	return cmpUncomparable
}

// eq checks if left and right are equal. If left and right can't be compared
// according to (Engine).cmp, false is returned.
func (e Engine) eq(left, right types.Value) bool {
	return e.cmp(left, right) == cmpEqual
}

// lt checks if the left value is less than the right value. For the <= (less
// than or equal) relation, see (Engine).lteq.
func (e Engine) lt(left, right types.Value) bool {
	return e.cmp(left, right) == cmpLessThan
}

// gt checks if the left value is less than the right value. For the >= (greater
// than or equal) relation, see (Engine).gteq.
func (e Engine) gt(left, right types.Value) bool {
	return e.lt(right, left)
}

// lteq checks if the left value is smaller than or equal to the right value.
func (e Engine) lteq(left, right types.Value) bool {
	return e.eq(left, right) || e.lt(left, right)
}

// gteq checks if the right value is smaller than or equal to the left value.
func (e Engine) gteq(left, right types.Value) bool {
	return e.eq(left, right) || e.gt(left, right)
}
