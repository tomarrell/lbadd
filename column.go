package lbadd

// The set of data types that are can be used for columns
type columnType int

const (
	columnTypeInvalid columnType = iota
	columnTypeInt
	columnTypeFloat
	columnTypeBool
	columnTypeString
	columnTypeDateTime
)

func (c columnType) String() string {
	switch c {
	case columnTypeInt:
		return "integer"
	case columnTypeFloat:
		return "float"
	case columnTypeBool:
		return "boolean"
	case columnTypeString:
		return "string"
	case columnTypeDateTime:
		return "datetime"
	default:
		return "invalid"
	}
}

// A single column within a table
type column struct {
	dataType   columnType
	name       string
	isNullable bool
}
