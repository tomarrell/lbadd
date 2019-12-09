package main

type query struct {
	queryType  queryType
	tableName  string
	conditions []condition
	updates    map[string]string
	inserts    [][]string
	fields     []string
}

// The type of the parsed query
// e.g. SELECT, INSERT etc
type queryType int

const (
	queryUnknownType queryType = iota
	selectQuery
	updateQuery
	insertQuery
	deleteQuery
)

func (qt queryType) String() string {
	switch qt {
	case selectQuery:
		return "select"
	case updateQuery:
		return "update"
	case insertQuery:
		return "insert"
	case deleteQuery:
		return "deleteQuery"
	default:
		return "unknownOperator"
	}
}

// Operator
type operatorType int

const (
	unknownOperator operatorType = iota
	equal                        // =
	notEqual                     // !=
	greater                      // >
	lesser                       // <
	greaterOrEqual               // >=
	lesserOrEqual                // <=
)

func (ot operatorType) String() string {
	switch ot {
	case equal:
		return "equal"
	case notEqual:
		return "notEqual"
	case greater:
		return "greater"
	case lesser:
		return "lesser"
	case greaterOrEqual:
		return "greaterOrEqual"
	case lesserOrEqual:
		return "lesserOrEqual"
	default:
		return "unknownOperator"
	}
}

// Condition
type condition struct {
	lhs        string
	lhsIsField bool
	operator   operatorType
	rhs        string
	rhsIsField bool
}
