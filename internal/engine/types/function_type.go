package types

var (
	// Function is the function type. Functions are not comparable. The name of
	// this type is "Function".
	Function = FunctionType{
		typ: typ{
			name: "Function",
		},
	}
)

// FunctionType is a non-comparable type.
type FunctionType struct {
	typ
}
