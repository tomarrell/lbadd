package engine

type Evt string

func (e Evt) String() string { return string(e) }

type ParameterizedEvt struct {
	Name  string
	Param string
}

func (e ParameterizedEvt) String() string {
	return e.Name + "[" + e.Param + "]"
}

const (
	EvtEvaluate Evt = "evaluate"
)

func EvtFullTableScan(tableName string) ParameterizedEvt {
	return ParameterizedEvt{
		Name:  "full table scan",
		Param: "table=" + tableName,
	}
}
