package column

//go:generate stringer -type=BaseType

// BaseType is the base type of a column, in unparameterized form. To
// parameterize a base type, use a (column.Type).
type BaseType uint16

var (
	parameterCount = map[BaseType]uint8{
		Decimal: 2,
		Varchar: 1,
	}
)

// NumParameters returns the amount of parameters, that the base type supports.
// For example, this is 1 for the VARCHAR type, and 2 for the DECIMAL type.
func (t BaseType) NumParameters() uint8 {
	return parameterCount[t] // zero is default value
}

// Supported base types.
const (
	Unknown BaseType = iota
	Decimal
	Varchar
)

// Type describes a type that consists of a base type and zero, one or two
// number parameters.
type Type interface {
	// BaseType returns the base type of this column type. Depending on the base
	// type, IsParameterized implies different constellations. Some base types
	// support only one parameter, some support two. If it supports one or two
	// parameters can be determined by calling NumParameters() of the BaseType.
	BaseType() BaseType
	IsParameterized() bool
	FirstParameter() float64
	SecondParameter() float64
}
