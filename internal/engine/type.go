package engine

type (
	// Comparator is the interface that wraps the basic compare method. The
	// compare method compares the left and right value as follows. -1 if
	// left<right, 0 if left==right, 1 if left>right. What exectly is considered
	// to be <, ==, > is up to the implementation.
	Comparator interface {
		Compare(Value, Value) (int, error)
	}

	// Type is a data type that consists of a type descriptor and a comparator.
	// The comparator forces types to define relations between two values of
	// this type. A type is only expected to be able to handle values of its own
	// type.
	Type interface {
		TypeDescriptor
		Comparator
	}
)
