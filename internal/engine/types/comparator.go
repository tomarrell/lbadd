package types

// Comparator is the interface that wraps the basic compare method. The compare
// method compares the left and right value as follows. -1 if left<right, 0 if
// left==right, 1 if left>right. What exectly is considered to be <, ==, > is up
// to the implementation. By definition, the NULL value is smaller than any
// other value. When comparing NULL to another NULL value, and both NULLs have
// the same type, the result is undefined, however, no error must be returned.
type Comparator interface {
	// Compare compares the given to values left and right as follows. -1 if
	// left<right, 0 if left==right, 1 if left>right. However, NULL<any, so if
	// the left value is NULL, and is comparable to the right value (same type),
	// this will return -1 and no error. NULL~NULL is undefined.
	Compare(left, right Value) (int, error)
}
