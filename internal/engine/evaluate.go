package engine

import (
	"fmt"
	"sort"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func (e Engine) evaluate(ctx ExecutionContext, c command.Command) (Table, error) {
	switch cmd := c.(type) {
	case command.List:
		return e.evaluateList(ctx, cmd)
	}
	return Table{}, ErrUnimplemented(c)
}

func (e Engine) evaluateList(ctx ExecutionContext, l command.List) (Table, error) {
	switch list := l.(type) {
	case command.Values:
		values, err := e.evaluateValues(ctx, list)
		if err != nil {
			return Table{}, fmt.Errorf("values: %w", err)
		}
		return values, nil
	case command.Scan:
		scanned, err := e.evaluateScan(ctx, list)
		if err != nil {
			return Table{}, fmt.Errorf("scan: %w", err)
		}
		return scanned, nil
	case command.Project:
		return e.evaluateProjection(ctx, list)
	}
	return Table{}, ErrUnimplemented(l)
}

func (e Engine) evaluateProjection(ctx ExecutionContext, proj command.Project) (Table, error) {
	origin, err := e.evaluateList(ctx, proj.Input)
	if err != nil {
		return Table{}, fmt.Errorf("list: %w", err)
	}

	if len(proj.Cols) == 0 {
		e.log.Debug().
			Str("ctx", ctx.String()).
			Msg("projection filters all columns")
		return EmptyTable, nil
	}

	var expectedColumnNames []string
	aliases := make(map[string]string)
	for _, col := range proj.Cols {
		// evaluate the column name
		colNameExpr, err := e.evaluateExpression(ctx, col.Column)
		if err != nil {
			return Table{}, fmt.Errorf("eval column name: %w", err)
		}
		var colName string
		if colNameExpr.Is(types.String) {
			colName = colNameExpr.(types.StringValue).Value
		} else {
			casted, err := types.String.Cast(colNameExpr)
			if err != nil {
				return Table{}, fmt.Errorf("cannot cast %v to %v: %w", colNameExpr.Type(), types.String, err)
			}
			colName = casted.(types.StringValue).Value
		}
		if col.Table != "" {
			colName = col.Table + "." + colName
		}
		if col.Alias != "" {
			aliases[colName] = col.Alias
		}

		expectedColumnNames = append(expectedColumnNames, colName)
	}

	// check if the table actually has all expected columns
	for _, expectedCol := range expectedColumnNames {
		if !origin.HasColumn(expectedCol) {
			return Table{}, ErrNoSuchColumn(expectedCol)
		}
	}

	// apply aliases
	for i, col := range origin.Cols {
		if alias, ok := aliases[col.QualifiedName]; ok {
			origin.Cols[i].Alias = alias
		}
	}

	sort.Strings(expectedColumnNames)

	var toRemove []string
	for _, col := range origin.Cols {
		searchResult := sort.SearchStrings(expectedColumnNames, col.QualifiedName)
		if searchResult == len(expectedColumnNames) || expectedColumnNames[searchResult] != col.QualifiedName {
			toRemove = append(toRemove, col.QualifiedName)
		}
	}

	for _, toRemoveCol := range toRemove {
		origin = origin.RemoveColumnByQualifiedName(toRemoveCol)
	}

	return origin, nil
}

func (e Engine) evaluateValues(ctx ExecutionContext, v command.Values) (tbl Table, err error) {
	var colCnt int
	for _, values := range v.Values {
		rowValues := make([]types.Value, len(values))
		colCnt = len(values)
		for x, value := range values {
			internalValue, err := e.evaluateExpression(ctx, value)
			if err != nil {
				return Table{}, fmt.Errorf("expr: %w", err)
			}
			rowValues[x] = internalValue
		}
		tbl.Rows = append(tbl.Rows, Row{
			Values: rowValues,
		})
	}

	for i := 1; i <= colCnt; i++ {
		tbl.Cols = append(tbl.Cols, Col{
			QualifiedName: fmt.Sprintf("column%d", i),
			Type:          tbl.Rows[0].Values[i-1].Type(),
		})
	}

	return
}

func (e Engine) evaluateScan(ctx ExecutionContext, s command.Scan) (Table, error) {
	defer e.profiler.Enter(EvtFullTableScan(s.Table.QualifiedName())).Exit()

	switch table := s.Table.(type) {
	case command.SimpleTable:
		return e.scanSimpleTable(ctx, table)
	default:
		return Table{}, ErrUnimplemented(fmt.Sprintf("scan %T", table))
	}
}
