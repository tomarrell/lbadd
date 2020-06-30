package message

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// ConvertCommandToMessage converts a command.Command to a message.Command
func ConvertCommandToMessage(cmd command.Command) Message {
	switch c := cmd.(type) {
	case command.Scan:
		return ConvertCommandToMessageScan(c)
	case command.Select:
		return ConvertCommandToMessageSelect(c)
	case command.Project:
		return ConvertCommandToMessageProject(c)
	case command.Delete:
		return ConvertCommandToMessageDelete(c)
	case command.DropIndex:
		return ConvertCommandToMessageDrop(c)
	case command.DropTable:
		return ConvertCommandToMessageDrop(c)
	case command.DropTrigger:
		return ConvertCommandToMessageDrop(c)
	case command.DropView:
		return ConvertCommandToMessageDrop(c)
	case command.Update:
		return ConvertCommandToMessageUpdate(c)
	case command.Join:
		return ConvertCommandToMessageJoin(c)
	case command.Limit:
		return ConvertCommandToMessageLimit(c)
	case command.Insert:
		return ConvertCommandToMessageInsert(c)
	}
	return nil
}

// ConvertCommandToMessageTable converts a command.Table to a SimpleTable
func ConvertCommandToMessageTable(cmd *command.Table) *SimpleTable {
	simpleTable := &SimpleTable{}

	return simpleTable
}

// ConvertCommandToMessageScan converts a Command type to a Command_Scan type.
func ConvertCommandToMessageScan(cmd command.Command) *Command_Scan {
	msgCmdScan := &Command_Scan{}
	// msgCmdScan.Table = ConvertCommandToMessageTable(cmd.(*command.Scan).Table)
	return msgCmdScan
}

// ConvertCommandToMessageSelect converts a Command type to a Command_Select type.
func ConvertCommandToMessageSelect(command command.Command) *Command_Select {
	return command.(*Command_Select)
}

// ConvertCommandToMessageProject converts a Command type to a Command_Project type.
func ConvertCommandToMessageProject(command command.Command) *Command_Project {
	return command.(*Command_Project)
}

// ConvertCommandToMessageDelete converts a Command type to a Command_Delete type.
func ConvertCommandToMessageDelete(command command.Command) *Command_Delete {
	return command.(*Command_Delete)
}

// ConvertCommandToMessageDrop converts a Command type to a CommandDrop type.
func ConvertCommandToMessageDrop(command command.Command) *CommandDrop {
	return command.(*CommandDrop)
}

// ConvertCommandToMessageUpdate converts a Command type to a Command_Update type.
func ConvertCommandToMessageUpdate(command command.Command) *Command_Update {
	return command.(*Command_Update)
}

// ConvertCommandToMessageJoin converts a Command type to a Command_Join type.
func ConvertCommandToMessageJoin(command command.Command) *Command_Join {
	return command.(*Command_Join)
}

// ConvertCommandToMessageLimit converts a Command type to a Command_Limit type.
func ConvertCommandToMessageLimit(command command.Command) *Command_Limit {
	return command.(*Command_Limit)
}

// ConvertCommandToMessageInsert converts a Command type to a Command_Insert type.
func ConvertCommandToMessageInsert(command command.Command) *Command_Insert {
	return command.(*Command_Insert)
}

// ConvertMessageToCommand converts a message.Command to a command.Command
func ConvertMessageToCommand(msg Message) command.Command {
	switch m := msg.(type) {
	case *Command_Scan:
		return ConvertMessageToCommandScan(m)
	case *Command_Select:
		return ConvertMessageToCommandSelect(m)
	case *Command_Project:
		return ConvertMessageToCommandProject(m)
	case *Command_Delete:
		return ConvertMessageToCommandDelete(m)
	case *CommandDrop:
		switch m.Target {
		case 0:
			return ConvertMessageToCommandDropTable(m)
		case 1:
			return ConvertMessageToCommandDropView(m)
		case 2:
			return ConvertMessageToCommandDropIndex(m)
		case 3:
			return ConvertMessageToCommandDropTrigger(m)
		}
	case *Command_Update:
		return ConvertCommandToMessageUpdate(m)
	case *Command_Join:
		return ConvertCommandToMessageJoin(m)
	case *Command_Limit:
		return ConvertCommandToMessageLimit(m)
	case *Command_Insert:
		return ConvertCommandToMessageInsert(m)
	}
	return nil
}

// ConvertMessageToCommandTable converts a message.SimpleTable to a command.Table
func ConvertMessageToCommandTable(msg *SimpleTable) command.Table {
	cmdTable := command.SimpleTable{}
	cmdTable.Schema = msg.Schema
	cmdTable.Table = msg.Table
	cmdTable.Alias = msg.Alias
	cmdTable.Indexed = msg.Indexed
	cmdTable.Index = msg.Index
	return cmdTable
}

// ConvertMessageToCommandScan converts a message.Command_Scan to a command.Scan
func ConvertMessageToCommandScan(msg *Command_Scan) command.Scan {
	cmdScan := command.Scan{}
	cmdScan.Table = ConvertMessageToCommandTable(msg.Table)
	return cmdScan
}

// ConvertMessageToCommandLiteralExpr converts a message.Expr to a command.LiteralExpr
func ConvertMessageToCommandLiteralExpr(msg *Expr) command.LiteralExpr {
	literalExpr := command.LiteralExpr{}
	literalExpr.Value = msg.GetLiteral().GetValue()
	return literalExpr
}

// ConvertMessageToCommandConstantBooleanExpr converts a message.Expr to a command.ConstantBooleanExpr
func ConvertMessageToCommandConstantBooleanExpr(msg *Expr) command.ConstantBooleanExpr {
	constantBooleanExpr := command.ConstantBooleanExpr{}
	constantBooleanExpr.Value = msg.GetConstant().GetValue()
	return constantBooleanExpr
}

// ConvertMessageToCommandUnaryExpr converts a message.Expr to a command.UnaryExpr
func ConvertMessageToCommandUnaryExpr(msg *Expr) command.UnaryExpr {
	unaryExpr := command.UnaryExpr{}
	unaryExpr.Operator = msg.GetUnary().GetOperator()
	unaryExpr.Value = ConvertMessageToCommandExpr(msg.GetUnary().GetValue())
	return unaryExpr
}

// ConvertMessageToCommandBinaryExpr converts a message.Expr to a command.BinaryExpr
func ConvertMessageToCommandBinaryExpr(msg *Expr) command.BinaryExpr {
	binaryExpr := command.BinaryExpr{}
	binaryExpr.Operator = msg.GetBinary().GetOperator()
	binaryExpr.Left = ConvertMessageToCommandExpr(msg.GetBinary().GetLeft())
	binaryExpr.Right = ConvertMessageToCommandExpr(msg.GetBinary().GetRight())
	return binaryExpr
}

// ConvertMessageToCommandExprSlice converts a []*message.Expr to []command.Expr
func ConvertMessageToCommandExprSlice(msg []*Expr) []command.Expr {
	// ConvertTODO
	return []command.Expr{}
}

// ConvertMessageToCommandFunctionExpr converts a message.Expr to a command.FunctionExpr
func ConvertMessageToCommandFunctionExpr(msg *Expr) command.FunctionExpr {
	functionExpr := command.FunctionExpr{}
	functionExpr.Name = msg.GetFunc().GetName()
	functionExpr.Distinct = msg.GetFunc().GetDistinct()
	functionExpr.Args = ConvertMessageToCommandExprSlice(msg.GetFunc().GetArgs())
	return functionExpr
}

// ConvertMessageToCommandEqualityExpr converts a message.Expr to a command.EqualityExpr
func ConvertMessageToCommandEqualityExpr(msg *Expr) command.EqualityExpr {
	equalityExpr := command.EqualityExpr{}
	equalityExpr.Left = ConvertMessageToCommandExpr(msg.GetEquality().GetLeft())
	equalityExpr.Right = ConvertMessageToCommandExpr(msg.GetEquality().GetRight())
	equalityExpr.Invert = msg.GetEquality().Invert
	return equalityExpr
}

// ConvertMessageToCommandRangeExpr converts a message.Expr to a command.RangeExpr
func ConvertMessageToCommandRangeExpr(msg *Expr) command.RangeExpr {
	rangeExpr := command.RangeExpr{}
	rangeExpr.Needle = ConvertMessageToCommandExpr(msg.GetRange().GetNeedle())
	rangeExpr.Lo = ConvertMessageToCommandExpr(msg.GetRange().GetLo())
	rangeExpr.Hi = ConvertMessageToCommandExpr(msg.GetRange().GetHi())
	return rangeExpr
}

// ConvertMessageToCommandExpr converts a message.Expr to a command.Expr
func ConvertMessageToCommandExpr(msg *Expr) command.Expr {
	switch msg.Kind() {
	case KindLiteralExpr:
		return ConvertMessageToCommandLiteralExpr(msg)
	case KindConstantBooleanExpr:
		return ConvertMessageToCommandConstantBooleanExpr(msg)
	case KindUnaryExpr:
		return ConvertMessageToCommandUnaryExpr(msg)
	case KindBinaryExpr:
		return ConvertMessageToCommandBinaryExpr(msg)
	case KindFunctionExpr:
		return ConvertMessageToCommandFunctionExpr(msg)
	case KindEqualityExpr:
		return ConvertMessageToCommandEqualityExpr(msg)
	case KindRangeExpr:
		return ConvertMessageToCommandRangeExpr(msg)
	}
	return nil
}

// ConvertMessageToCommandListScan converts a message.List to a command.Scan
func ConvertMessageToCommandListScan(msg *List) command.Scan {
	cmdScan := command.Scan{}
	cmdScan.Table = ConvertMessageToCommandTable(msg.GetScan().GetTable())
	return cmdScan
}

// ConvertMessageToCommandListSelect converts a message.List to a command.Select
func ConvertMessageToCommandListSelect(msg *List) command.Select {
	cmdSelect := command.Select{}
	cmdSelect.Filter = ConvertMessageToCommandExpr(msg.GetSelect().GetFilter())
	cmdSelect.Input = ConvertMessageToCommandList(msg.GetSelect().GetInput())
	return cmdSelect
}

// ConvertMessageToCommandListProject converts a message.List to a command.Project
func ConvertMessageToCommandListProject(msg *List) command.Project {
	cmdProject := command.Project{}
	cmdProject.Cols = ConvertMessageToCommandCols(msg.GetProject().GetCols())
	cmdProject.Input = ConvertMessageToCommandList(msg.GetProject().GetInput())
	return cmdProject
}

// ConvertMessageToCommandListJoin converts a message.List to a command.Join
func ConvertMessageToCommandListJoin(msg *List) command.Join {
	cmdJoin := command.Join{}
	cmdJoin.Natural = msg.GetJoin().GetNatural()
	cmdJoin.Type = ConvertMessageToCommandJoinType(msg.GetJoin().GetType())
	cmdJoin.Filter = ConvertMessageToCommandExpr(msg.GetJoin().GetFilter())
	cmdJoin.Left = ConvertMessageToCommandList(msg.GetJoin().GetLeft())
	cmdJoin.Right = ConvertMessageToCommandList(msg.GetJoin().GetRight())
	return cmdJoin
}

// ConvertMessageToCommandListLimit converts a message.List to a command.Limit
func ConvertMessageToCommandListLimit(msg *List) command.Limit {
	cmdLimit := command.Limit{}
	cmdLimit.Limit = ConvertMessageToCommandExpr(msg.GetLimit().GetLimit())
	cmdLimit.Input = ConvertMessageToCommandList(msg.GetLimit().GetInput())
	return cmdLimit
}

// ConvertMessageToCommandListOffset converts a message.List to a command.Offset
func ConvertMessageToCommandListOffset(msg *List) command.Offset {
	cmdOffset := command.Offset{}
	cmdOffset.Offset = ConvertMessageToCommandExpr(msg.GetOffset().GetOffset())
	cmdOffset.Input = ConvertMessageToCommandList(msg.GetDistinct().GetInput())
	return cmdOffset
}

// ConvertMessageToCommandListDistinct converts a message.List to a command.Distinct
func ConvertMessageToCommandListDistinct(msg *List) command.Distinct {
	cmdDistinct := command.Distinct{}
	cmdDistinct.Input = ConvertMessageToCommandList(msg.GetDistinct().GetInput())
	return cmdDistinct
}

// ConvertMessageToCommandExprRepeatedSlice converts a message.RepeatedExpr to a [][]command.Expr
func ConvertMessageToCommandExprRepeatedSlice(msg []*RepeatedExpr) [][]command.Expr {
	cmdRepeatedExprSlice := [][]command.Expr{}
	for i := range msg {
		cmdRepeatedExpr := []command.Expr{}
		for j := range msg[i].Expr {
			cmdRepeatedExpr = append(cmdRepeatedExpr, ConvertMessageToCommandExpr(msg[i].Expr[j]))
		}
		cmdRepeatedExprSlice = append(cmdRepeatedExprSlice, cmdRepeatedExpr)
	}
	return cmdRepeatedExprSlice
}

// ConvertMessageToCommandListValues converts a message.List to a command.Values
func ConvertMessageToCommandListValues(msg *List) command.Values {
	cmdValues := command.Values{}
	cmdValues.Values = ConvertMessageToCommandExprRepeatedSlice(msg.GetValues().GetExpr())
	return cmdValues
}

// ConvertMessageToCommandList converts a message.List to a command.List
func ConvertMessageToCommandList(msg *List) command.List {
	switch msg.List.(type) {
	case *List_Scan:
		return ConvertMessageToCommandListScan(msg)
	case *List_Select:
		return ConvertMessageToCommandListSelect(msg)
	case *List_Project:
		return ConvertMessageToCommandListProject(msg)
	case *List_Join:
		return ConvertMessageToCommandListJoin(msg)
	case *List_Limit:
		return ConvertMessageToCommandListLimit(msg)
	case *List_Offset:
		return ConvertMessageToCommandListOffset(msg)
	case *List_Distinct:
		return ConvertMessageToCommandListDistinct(msg)
	case *List_Values:
		return ConvertMessageToCommandListValues(msg)
	}
	return nil
}

// ConvertMessageToCommandSelect converts a message.Command_Select to a command.Select
func ConvertMessageToCommandSelect(msg *Command_Select) command.Select {
	cmdSelect := command.Select{}
	cmdSelect.Filter = ConvertMessageToCommandExpr(msg.GetFilter())
	cmdSelect.Input = ConvertMessageToCommandList(msg.GetInput())
	return cmdSelect
}

// ConvertMessageToCommandCol converts a message.Column to a command.Column
func ConvertMessageToCommandCol(msg *Column) command.Column {
	cmdCol := command.Column{}
	cmdCol.Table = msg.GetTable()
	cmdCol.Column = ConvertMessageToCommandExpr(msg.GetColumn())
	cmdCol.Alias = msg.GetAlias()
	return cmdCol
}

// ConvertMessageToCommandCols converts a []message.Column to a []command.Column
func ConvertMessageToCommandCols(msg []*Column) []command.Column {
	cmdCols := []command.Column{}
	for i := range msg {
		cmdCols = append(cmdCols, ConvertMessageToCommandCol(msg[i]))
	}
	return cmdCols
}

// ConvertMessageToCommandProject converts a message.Command_Project to a command.Project
func ConvertMessageToCommandProject(msg *Command_Project) command.Project {
	cmdProject := command.Project{}
	cmdProject.Cols = ConvertMessageToCommandCols(msg.GetCols())
	cmdProject.Input = ConvertMessageToCommandList(msg.GetInput())
	return cmdProject
}

// ConvertMessageToCommandDelete converts a message.Command_Delete to a command.Delete
func ConvertMessageToCommandDelete(msg *Command_Delete) command.Delete {
	cmdDelete := command.Delete{}
	cmdDelete.Filter = ConvertMessageToCommandExpr(msg.GetFilter())
	cmdDelete.Table = ConvertMessageToCommandTable(msg.GetTable())
	return cmdDelete
}

// ConvertMessageToCommandDropTable converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropTable(msg *CommandDrop) command.DropTable {
	cmdDropTable := command.DropTable{}
	cmdDropTable.IfExists = msg.GetIfExists()
	cmdDropTable.Schema = msg.GetSchema()
	cmdDropTable.Name = msg.GetName()
	return cmdDropTable
}

// ConvertMessageToCommandDropView converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropView(msg *CommandDrop) command.DropView {
	cmdDropView := command.DropView{}
	cmdDropView.IfExists = msg.GetIfExists()
	cmdDropView.Schema = msg.GetSchema()
	cmdDropView.Name = msg.GetName()
	return cmdDropView
}

// ConvertMessageToCommandDropIndex converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropIndex(msg *CommandDrop) command.DropIndex {
	cmdDropIndex := command.DropIndex{}
	cmdDropIndex.IfExists = msg.GetIfExists()
	cmdDropIndex.Schema = msg.GetName()
	cmdDropIndex.Name = msg.GetName()
	return cmdDropIndex
}

// ConvertMessageToCommandDropTrigger converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropTrigger(msg *CommandDrop) command.DropTrigger {
	cmdDropTrigger := command.DropTrigger{}
	cmdDropTrigger.IfExists = msg.GetIfExists()
	cmdDropTrigger.Schema = msg.GetSchema()
	cmdDropTrigger.Name = msg.GetName()
	return cmdDropTrigger
}

// ConvertMessageToCommandUpdateOr converts a message.UpdateOr to command.UpdateOr
func ConvertMessageToCommandUpdateOr(msg UpdateOr) command.UpdateOr {
	return command.UpdateOr(msg.Number())
}

// ConvertMessageToCommandUpdate converts a message.Command_Update to a command.Update
func ConvertMessageToCommandUpdate(msg *Command_Update) command.Update {
	cmdUpdate := command.Update{}
	cmdUpdate.UpdateOr = ConvertMessageToCommandUpdateOr(msg.GetUpdateOr())
	cmdUpdate.Filter = ConvertMessageToCommandExpr(msg.GetFilter())
	cmdUpdate.Table = ConvertMessageToCommandTable(msg.GetTable())
	return cmdUpdate
}

// ConvertMessageToCommandJoinType converts a message.JoinType to a command.JoinType
func ConvertMessageToCommandJoinType(msg JoinType) command.JoinType {
	return command.JoinType(msg.Number())
}

// ConvertMessageToCommandJoin converts a message.Command_Join to a command.Join
func ConvertMessageToCommandJoin(msg *Command_Join) command.Join {
	cmdJoin := command.Join{}
	cmdJoin.Natural = msg.Natural
	cmdJoin.Type = ConvertMessageToCommandJoinType(msg.GetType())
	cmdJoin.Filter = ConvertMessageToCommandExpr(msg.GetFilter())
	cmdJoin.Left = ConvertMessageToCommandList(msg.GetLeft())
	cmdJoin.Right = ConvertMessageToCommandList(msg.GetRight())
	return cmdJoin
}

// ConvertMessageToCommandLimit converts a message.Command_Limit to a command.Limit
func ConvertMessageToCommandLimit(msg *Command_Limit) command.Limit {
	cmdLimit := command.Limit{}
	cmdLimit.Limit = ConvertMessageToCommandExpr(msg.GetLimit())
	cmdLimit.Input = ConvertMessageToCommandList(msg.GetInput())
	return cmdLimit
}

// ConvertMessageToCommandInsertOr converts a message.InsertOr to command.InsertOr
func ConvertMessageToCommandInsertOr(msg InsertOr) command.InsertOr {
	return command.InsertOr(msg.Number())
}

// ConvertMessageToCommandInsert converts a message.Command_Insert to a command.Insert
func ConvertMessageToCommandInsert(msg *Command_Insert) command.Insert {
	cmdInsert := command.Insert{}
	cmdInsert.InsertOr = ConvertMessageToCommandInsertOr(msg.GetInsertOr())
	cmdInsert.Table = ConvertMessageToCommandTable(msg.GetTable())
	cmdInsert.Cols = ConvertMessageToCommandCols(msg.GetCols())
	cmdInsert.DefaultValues = msg.GetDefaultValues()
	cmdInsert.Input = ConvertMessageToCommandList(msg.GetInput())
	return cmdInsert
}
