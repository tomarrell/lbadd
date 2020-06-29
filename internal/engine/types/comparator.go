package types

// Comparator is the interface that wraps the basic compare method. The
// compare method compares the left and right value as follows. -1 if
// left<right, 0 if left==right, 1 if left>right. What exectly is considered
// to be <, ==, > is up to the implementation.
type Comparator interface {
	// Compare compares the given to values left and right as follows. -1 if
	// left<right, 0 if left==right, 1 if left>right.
	Compare(left, right Value) (int, error)
}
