package lbadd

// The set of data types that are can be used for columns
type columnType int

const (
	columnTypeInvalid columnType = iota
	columnTypeInt
	columnTypeFloat
	columnTypeBool
	columnTypeString
	columnTypeDate
)

// A single column within a table
type column struct {
	dataType   columnType
	name       string
	isNullable bool
}
