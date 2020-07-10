package engine

const (
	// EvtEvaluate is the event 'evaluate'. This is used for every call to
	// 'Evaluate' (exported), and is used to measure the total amount of time it
	// takes the database to evaluate a command.
	EvtEvaluate Evt = "evaluate"
	// EvtCompare is the event 'compare'. This is used for every comparison
	// (with 'cmp') of two values that is performed.
	EvtCompare Evt = "compare"
)

// Evt is an event this engine uses.
type Evt string

func (e Evt) String() string { return string(e) }

// ParameterizedEvt is a parameterized event, which will be printed in the form
// Name[Param].
type ParameterizedEvt struct {
	// Name is the name of the event.
	Name string
	// Param is the parameterization of the event.
	Param string
}

func (e ParameterizedEvt) String() string {
	return e.Name + "[" + e.Param + "]"
}

// EvtFullTableScan creates an event 'full table scan[table=<tableName>]'.
func EvtFullTableScan(tableName string) ParameterizedEvt {
	return ParameterizedEvt{
		Name:  "full table scan",
		Param: "table=" + tableName,
	}
}
