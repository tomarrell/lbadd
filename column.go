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

var columnNames []string = []string{"invalid", "integer", "float", "boolean", "string", "datetime"}

func (c columnType) String() string {
	if int(c) > len(columnNames)-1 || c < 0 {
		return columnNames[0]
	}

	return columnNames[c]
}

func parseColumnType(str string) columnType {
	for i, v := range columnNames {
		if str == v {
			return columnType(i)
		}
	}

	return 0
}

// A single column within a table
type column struct {
	dataType   columnType
	name       string
	isNullable bool
}
