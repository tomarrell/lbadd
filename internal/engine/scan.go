package engine

import "github.com/tomarrell/lbadd/internal/compiler/command"

func (e Engine) scanSimpleTable(ctx ExecutionContext, table command.SimpleTable) (Table, error) {
	tableName := table.QualifiedName()

	// only perform scan if not already scanned
	if table, alreadyScanned := ctx.getScannedTable(tableName); alreadyScanned {
		return table, nil
	}

	// TODO: load table from the database file

	ctx.putScannedTable(table.QualifiedName(), Table{})
	return Table{}, ErrUnimplemented("scan simple table")
}
