package message

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// ConvertCommandToMessage converts a command.Command to a message.Command
func ConvertCommandToMessage(cmd command.Command) Message {
	switch c := cmd.(type) {
	case *command.Scan:
		return ConvertCommandToMessageScan(c)
	case *command.Select:
		return ConvertCommandToMessageSelect(c)
	case *command.Project:
		return ConvertCommandToMessageProject(c)
	case *command.Delete:
		return ConvertCommandToMessageDelete(c)
	case *command.DropIndex:
		return ConvertCommandToMessageDrop(c)
	case *command.DropTable:
		return ConvertCommandToMessageDrop(c)
	case *command.DropTrigger:
		return ConvertCommandToMessageDrop(c)
	case *command.DropView:
		return ConvertCommandToMessageDrop(c)
	case *command.Update:
		return ConvertCommandToMessageUpdate(c)
	case *command.Join:
		return ConvertCommandToMessageJoin(c)
	case *command.Limit:
		return ConvertCommandToMessageLimit(c)
	case *command.Insert:
		return ConvertCommandToMessageInsert(c)
	}
	return nil
}

// ConvertCommandToMessageTable converts a command.Table to a SimpleTable
func ConvertCommandToMessageTable(cmd command.Table) *SimpleTable {
	simpleTable := &SimpleTable{}
	simpleTable.Schema = cmd.(*command.SimpleTable).Schema
	simpleTable.Table = cmd.(*command.SimpleTable).Table
	simpleTable.Alias = cmd.(*command.SimpleTable).Alias
	simpleTable.Indexed = cmd.(*command.SimpleTable).Indexed
	simpleTable.Index = cmd.(*command.SimpleTable).Index
	return simpleTable
}

// ConvertCommandToMessageScan converts a Command type to a Command_Scan type.
func ConvertCommandToMessageScan(cmd command.Command) *Command_Scan {
	msgCmdScan := &Command_Scan{}
	msgCmdScan.Table = ConvertCommandToMessageTable(cmd.(*command.Scan).Table)
	return msgCmdScan
}

// ConvertCommandToMessageLiteralExpr converts a command.Expr to a message.Expr_Literal
func ConvertCommandToMessageLiteralExpr(cmd *command.LiteralExpr) *Expr_Literal {
	msgExprLiteral := &Expr_Literal{&LiteralExpr{}}
	msgExprLiteral.Literal.Value = cmd.Value
	return msgExprLiteral
}

// ConvertCommandToMessageConstantBooleanExpr converts a command.Expr to a message.Expr_Constant
func ConvertCommandToMessageConstantBooleanExpr(cmd *command.ConstantBooleanExpr) *Expr_Constant {
	msgExprConstant := &Expr_Constant{&ConstantBooleanExpr{}}
	msgExprConstant.Constant.Value = cmd.Value
	return msgExprConstant
}

// ConvertCommandToMessageUnaryExpr converts a command.Expr to a message.Expr_Unary
func ConvertCommandToMessageUnaryExpr(cmd *command.UnaryExpr) *Expr_Unary {
	msgExprUnary := &Expr_Unary{&UnaryExpr{}}
	msgExprUnary.Unary.Operator = cmd.Operator
	msgExprUnary.Unary.Value = ConvertCommandToMessageExpr(cmd.Value)
	return msgExprUnary
}

// ConvertCommandToMessageBinaryExpr converts a command.Expr to a message.Expr_Binary
func ConvertCommandToMessageBinaryExpr(cmd *command.BinaryExpr) *Expr_Binary {
	msgExprBinary := &Expr_Binary{&BinaryExpr{}}
	msgExprBinary.Binary.Operator = cmd.Operator
	msgExprBinary.Binary.Left = ConvertCommandToMessageExpr(cmd.Left)
	msgExprBinary.Binary.Right = ConvertCommandToMessageExpr(cmd.Right)
	return msgExprBinary
}

// ConvertCommandToMessageRepeatedExpr converts a []command.Expr to a message.Expr
func ConvertCommandToMessageRepeatedExpr(cmd []command.Expr) []*Expr {
	msgRepeatedExpr := []*Expr{}
	for i := range cmd {
		msgRepeatedExpr = append(msgRepeatedExpr, ConvertCommandToMessageExpr(cmd[i]))
	}
	return msgRepeatedExpr
}

// ConvertCommandToMessageFunctionalExpr converts a command.Expr to a message.Expr_CFExpr_Func
func ConvertCommandToMessageFunctionalExpr(cmd *command.FunctionExpr) *Expr_Func {
	msgExprFunc := &Expr_Func{&FunctionExpr{}}
	msgExprFunc.Func.Name = cmd.Name
	msgExprFunc.Func.Distinct = cmd.Distinct
	msgExprFunc.Func.Args = ConvertCommandToMessageRepeatedExpr(cmd.Args)
	return msgExprFunc
}

// ConvertCommandToMessageEqualityExpr converts a command.Expr to a message.Expr_Equality
func ConvertCommandToMessageEqualityExpr(cmd *command.EqualityExpr) *Expr_Equality {
	msgExprEquality := &Expr_Equality{&EqualityExpr{}}
	msgExprEquality.Equality.Left = ConvertCommandToMessageExpr(cmd.Left)
	msgExprEquality.Equality.Right = ConvertCommandToMessageExpr(cmd.Right)
	msgExprEquality.Equality.Invert = cmd.Invert
	return msgExprEquality
}

// ConvertCommandToMessageRangeExpr converts a command.Expr to a message.Expr_Range
func ConvertCommandToMessageRangeExpr(cmd *command.RangeExpr) *Expr_Range {
	msgExprRange := &Expr_Range{&RangeExpr{}}
	msgExprRange.Range.Needle = ConvertCommandToMessageExpr(cmd.Needle)
	msgExprRange.Range.Lo = ConvertCommandToMessageExpr(cmd.Lo)
	msgExprRange.Range.Hi = ConvertCommandToMessageExpr(cmd.Hi)
	msgExprRange.Range.Invert = cmd.Invert
	return msgExprRange
}

// ConvertCommandToMessageExpr converts command.Expr to a message.Expr
func ConvertCommandToMessageExpr(cmd command.Expr) *Expr {
	msgExpr := &Expr{}
	switch cmd := cmd.(type) {
	case *command.LiteralExpr:
		msgExpr.Expr = ConvertCommandToMessageLiteralExpr(cmd)
	case *command.ConstantBooleanExpr:
		msgExpr.Expr = ConvertCommandToMessageConstantBooleanExpr(cmd)
	case *command.UnaryExpr:
		msgExpr.Expr = ConvertCommandToMessageUnaryExpr(cmd)
	case *command.BinaryExpr:
		msgExpr.Expr = ConvertCommandToMessageBinaryExpr(cmd)
	case *command.FunctionExpr:
		msgExpr.Expr = ConvertCommandToMessageFunctionalExpr(cmd)
	case *command.EqualityExpr:
		msgExpr.Expr = ConvertCommandToMessageEqualityExpr(cmd)
	case *command.RangeExpr:
		msgExpr.Expr = ConvertCommandToMessageRangeExpr(cmd)
	}
	return msgExpr
}

// ConvertCommandToMessageListScan converts a command.Scan to a message.List_Scan
func ConvertCommandToMessageListScan(cmd *command.Scan) *List_Scan {
	msgListScan := &List_Scan{&Command_Scan{}}
	msgListScan.Scan.Table = ConvertCommandToMessageTable(cmd.Table)
	return msgListScan
}

// ConvertCommandToMessageListSelect converts a command.Select to a message.List_Select
func ConvertCommandToMessageListSelect(cmd *command.Select) *List_Select {
	msgListSelect := &List_Select{&Command_Select{}}
	msgListSelect.Select.Filter = ConvertCommandToMessageExpr(cmd.Filter)
	msgListSelect.Select.Input = ConvertCommandToMessageList(cmd.Input)
	return msgListSelect
}

// ConvertCommandToMessageListProject converts a command.Project to a message.List_Project
func ConvertCommandToMessageListProject(cmd *command.Project) *List_Project {
	msgListProject := &List_Project{&Command_Project{}}
	msgListProject.Project.Cols = ConvertCommandToMessageCols(cmd.Cols)
	msgListProject.Project.Input = ConvertCommandToMessageList(cmd.Input)
	return msgListProject
}

// ConvertCommandToMessageListJoin converts a command.Join to a message.List_Join
func ConvertCommandToMessageListJoin(cmd *command.Join) *List_Join {
	msgListJoin := &List_Join{&Command_Join{}}
	msgListJoin.Join.Natural = cmd.Natural
	msgListJoin.Join.Type = ConvertCommandToMessageJoinType(cmd.Type)
	msgListJoin.Join.Filter = ConvertCommandToMessageExpr(cmd.Filter)
	msgListJoin.Join.Left = ConvertCommandToMessageList(cmd.Left)
	msgListJoin.Join.Right = ConvertCommandToMessageList(cmd.Right)
	return msgListJoin
}

// ConvertCommandToMessageListLimit converts a command.Limit to a message.List_Limit
func ConvertCommandToMessageListLimit(cmd *command.Limit) *List_Limit {
	msgListLimit := &List_Limit{&Command_Limit{}}
	msgListLimit.Limit.Limit = ConvertCommandToMessageExpr(cmd.Limit)
	msgListLimit.Limit.Input = ConvertCommandToMessageList(cmd.Input)
	return msgListLimit
}

// ConvertCommandToMessageListOffset converts a command.Offset to a message.List_Offset
func ConvertCommandToMessageListOffset(cmd *command.Offset) *List_Offset {
	msgListOffset := &List_Offset{&Command_Offset{}}
	msgListOffset.Offset.Offset = ConvertCommandToMessageExpr(cmd.Offset)
	msgListOffset.Offset.Input = ConvertCommandToMessageList(cmd.Input)
	return msgListOffset
}

// ConvertCommandToMessageListDistinct converts a command.Distinct to a message.List_Distinct
func ConvertCommandToMessageListDistinct(cmd *command.Distinct) *List_Distinct {
	msgListDistinct := &List_Distinct{&Command_Distinct{}}
	msgListDistinct.Distinct.Input = ConvertCommandToMessageList(cmd.Input)
	return msgListDistinct
}

// ConvertCommandToMessageRepeatedExprSlice converts a [][]command.Expr to a [][]message.Expr
func ConvertCommandToMessageRepeatedExprSlice(cmd [][]command.Expr) []*RepeatedExpr {
	msgRepeatedExprSlice := []*RepeatedExpr{}
	for i := range cmd {
		msgRepeatedExpr := &RepeatedExpr{}
		for j := range cmd[i] {
			msgRepeatedExpr.Expr = append(msgRepeatedExpr.Expr, ConvertCommandToMessageExpr(cmd[i][j]))
		}
		msgRepeatedExprSlice = append(msgRepeatedExprSlice, msgRepeatedExpr)
	}
	return msgRepeatedExprSlice
}

// ConvertCommandToMessageListValues converts a command.Values to a message.List_Values
func ConvertCommandToMessageListValues(cmd *command.Values) *List_Values {
	msgListValues := &List_Values{}
	msgListValues.Values.Expr = ConvertCommandToMessageRepeatedExprSlice(cmd.Values)
	return msgListValues
}

// ConvertCommandToMessageList converts
func ConvertCommandToMessageList(cmd command.List) *List {
	msgList := &List{}
	switch cmd := cmd.(type) {
	case *command.Scan:
		msgList.List = ConvertCommandToMessageListScan(cmd)
	case *command.Select:
		msgList.List = ConvertCommandToMessageListSelect(cmd)
	case *command.Project:
		msgList.List = ConvertCommandToMessageListProject(cmd)
	case *command.Join:
		msgList.List = ConvertCommandToMessageListJoin(cmd)
	case *command.Limit:
		msgList.List = ConvertCommandToMessageListLimit(cmd)
	case *command.Offset:
		msgList.List = ConvertCommandToMessageListOffset(cmd)
	case *command.Distinct:
		msgList.List = ConvertCommandToMessageListDistinct(cmd)
	case *command.Values:
		msgList.List = ConvertCommandToMessageListValues(cmd)
	}
	return msgList
}

// ConvertCommandToMessageSelect converts a Command type to a Command_Select type.
func ConvertCommandToMessageSelect(cmd command.Command) *Command_Select {
	msgCmdSelect := &Command_Select{}
	msgCmdSelect.Filter = ConvertCommandToMessageExpr(cmd.(*command.Select).Filter)
	msgCmdSelect.Input = ConvertCommandToMessageList(cmd.(*command.Select).Input)
	return msgCmdSelect
}

// ConvertCommandToMessageCol converts command.Column to a message.Column
func ConvertCommandToMessageCol(cmd command.Column) *Column {
	msgCol := &Column{}
	msgCol.Table = cmd.Table
	msgCol.Column = ConvertCommandToMessageExpr(cmd.Column)
	msgCol.Alias = cmd.Alias
	return msgCol
}

// ConvertCommandToMessageCols converts []command.Column to a []message.Column
func ConvertCommandToMessageCols(cmd []command.Column) []*Column {
	msgCols := []*Column{}
	for i := range cmd {
		msgCols = append(msgCols, ConvertCommandToMessageCol(cmd[i]))
	}
	return msgCols
}

// ConvertCommandToMessageProject converts a Command type to a Command_Project type.
func ConvertCommandToMessageProject(cmd command.Command) *Command_Project {
	msgCmdProject := &Command_Project{}
	msgCmdProject.Cols = ConvertCommandToMessageCols(cmd.(*command.Project).Cols)
	msgCmdProject.Input = ConvertCommandToMessageList(cmd.(*command.Project).Input)
	return msgCmdProject
}

// ConvertCommandToMessageDelete converts a Command type to a Command_Delete type.
func ConvertCommandToMessageDelete(cmd command.Command) *Command_Delete {
	msgCmdDelete := &Command_Delete{}
	msgCmdDelete.Table = ConvertCommandToMessageTable(cmd.(*command.Delete).Table)
	msgCmdDelete.Filter = ConvertCommandToMessageExpr(cmd.(*command.Delete).Filter)
	return msgCmdDelete
}

// ConvertCommandToMessageDrop converts a Command type to a CommandDrop type.
func ConvertCommandToMessageDrop(cmd command.Command) *CommandDrop {
	msgCmdDrop := &CommandDrop{}
	switch cmd := cmd.(type) {
	case *command.DropTable:
		msgCmdDrop.Target = 0
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	case *command.DropView:
		msgCmdDrop.Target = 1
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	case *command.DropIndex:
		msgCmdDrop.Target = 2
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	case *command.DropTrigger:
		msgCmdDrop.Target = 3
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	}
	return msgCmdDrop
}

// ConvertCommandToMessageUpdateOr converts a command.Update or to a message.UpdateOr
func ConvertCommandToMessageUpdateOr(cmd command.UpdateOr) UpdateOr {
	switch cmd {
	case 0:
		return UpdateOr_UpdateOrUnknown
	case 1:
		return UpdateOr_UpdateOrRollback
	case 2:
		return UpdateOr_UpdateOrAbort
	case 3:
		return UpdateOr_UpdateOrReplace
	case 4:
		return UpdateOr_UpdateOrFail
	case 5:
		return UpdateOr_UpdateOrIgnore
	}
	return -1
}

// ConvertCommandToMessageUpdateSetterLiteral converts a command.Literal to a message.UpdateSetter_Literal
func ConvertCommandToMessageUpdateSetterLiteral(cmd command.LiteralExpr) *UpdateSetter_Literal {
	msgUpdateSetterLiteral := &UpdateSetter_Literal{&LiteralExpr{}}
	msgUpdateSetterLiteral.Literal.Value = cmd.Value
	return msgUpdateSetterLiteral
}

// ConvertCommandToMessageUpdateSetterConstant converts a command.Constant to a message.UpdateSetter_Constant
func ConvertCommandToMessageUpdateSetterConstant(cmd command.ConstantBooleanExpr) *UpdateSetter_Constant {
	msgUpdateSetterConstant := &UpdateSetter_Constant{&ConstantBooleanExpr{}}
	msgUpdateSetterConstant.Constant.Value = cmd.Value
	return msgUpdateSetterConstant
}

// ConvertCommandToMessageUpdateSetterUnary converts a command.Unary to a message.UpdateSetter_Unary
func ConvertCommandToMessageUpdateSetterUnary(cmd command.UnaryExpr) *UpdateSetter_Unary {
	msgUpdateSetterUnary := &UpdateSetter_Unary{&UnaryExpr{}}
	msgUpdateSetterUnary.Unary.Operator = cmd.Operator
	msgUpdateSetterUnary.Unary.Value = ConvertCommandToMessageExpr(cmd.Value)
	return msgUpdateSetterUnary
}

// ConvertCommandToMessageUpdateSetterBinary converts a command.Binary to a message.UpdateSetter_Binary
func ConvertCommandToMessageUpdateSetterBinary(cmd command.BinaryExpr) *UpdateSetter_Binary {
	msgUpdateSetterBinary := &UpdateSetter_Binary{&BinaryExpr{}}
	msgUpdateSetterBinary.Binary.Operator = cmd.Operator
	msgUpdateSetterBinary.Binary.Left = ConvertCommandToMessageExpr(cmd.Left)
	msgUpdateSetterBinary.Binary.Right = ConvertCommandToMessageExpr(cmd.Right)
	return msgUpdateSetterBinary
}

// ConvertCommandToMessageUpdateSetterFunc converts a command.Func to a message.UpdateSetter_Func
func ConvertCommandToMessageUpdateSetterFunc(cmd command.FunctionExpr) *UpdateSetter_Func {
	msgUpdateSetterFunc := &UpdateSetter_Func{&FunctionExpr{}}
	msgUpdateSetterFunc.Func.Name = cmd.Name
	msgUpdateSetterFunc.Func.Distinct = cmd.Distinct
	msgUpdateSetterFunc.Func.Args = ConvertCommandToMessageRepeatedExpr(cmd.Args)
	return msgUpdateSetterFunc
}

// ConvertCommandToMessageUpdateSetterEquality converts a command.Equality to a message.UpdateSetter_Equality
func ConvertCommandToMessageUpdateSetterEquality(cmd command.EqualityExpr) *UpdateSetter_Equality {
	msgUpdateSetterEquality := &UpdateSetter_Equality{&EqualityExpr{}}
	msgUpdateSetterEquality.Equality.Left = ConvertCommandToMessageExpr(cmd.Left)
	msgUpdateSetterEquality.Equality.Right = ConvertCommandToMessageExpr(cmd.Right)
	msgUpdateSetterEquality.Equality.Invert = cmd.Invert
	return msgUpdateSetterEquality
}

// ConvertCommandToMessageUpdateSetterRange converts a command.Range to a message.UpdateSetter_Range
func ConvertCommandToMessageUpdateSetterRange(cmd command.RangeExpr) *UpdateSetter_Range {
	msgUpdateSetterRange := &UpdateSetter_Range{&RangeExpr{}}
	msgUpdateSetterRange.Range.Needle = ConvertCommandToMessageExpr(cmd.Needle)
	msgUpdateSetterRange.Range.Lo = ConvertCommandToMessageExpr(cmd.Lo)
	msgUpdateSetterRange.Range.Hi = ConvertCommandToMessageExpr(cmd.Hi)
	msgUpdateSetterRange.Range.Invert = cmd.Invert
	return msgUpdateSetterRange
}

// ConvertCommandToMessageUpdateSetter converts a command.UpdateSetter to a message.UpdateSetter
func ConvertCommandToMessageUpdateSetter(cmd command.UpdateSetter) *UpdateSetter {
	msgUpdateSetter := &UpdateSetter{}
	msgUpdateSetter.Cols = cmd.Cols
	switch cmd.Value.(type) {
	case command.LiteralExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterLiteral(cmd.Value.(command.LiteralExpr))
	case command.ConstantBooleanExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterConstant(cmd.Value.(command.ConstantBooleanExpr))
	case command.UnaryExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterUnary(cmd.Value.(command.UnaryExpr))
	case command.BinaryExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterBinary(cmd.Value.(command.BinaryExpr))
	case command.FunctionExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterFunc(cmd.Value.(command.FunctionExpr))
	case command.EqualityExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterEquality(cmd.Value.(command.EqualityExpr))
	case command.RangeExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterRange(cmd.Value.(command.RangeExpr))
	}
	return msgUpdateSetter
}

// ConvertCommandToMessageUpdateSetterSlice converts a []command.UpdateSetter to a []message.UpdateSetter
func ConvertCommandToMessageUpdateSetterSlice(cmd []command.UpdateSetter) []*UpdateSetter {
	msgUpdateSetterSlice := []*UpdateSetter{}
	for i := range cmd {
		msgUpdateSetterSlice = append(msgUpdateSetterSlice, ConvertCommandToMessageUpdateSetter(cmd[i]))
	}
	return msgUpdateSetterSlice
}

// ConvertCommandToMessageUpdate converts a Command type to a Command_Update type.
func ConvertCommandToMessageUpdate(cmd command.Command) *Command_Update {
	msgCmdUpdate := &Command_Update{}
	msgCmdUpdate.UpdateOr = ConvertCommandToMessageUpdateOr(cmd.(*command.Update).UpdateOr)
	msgCmdUpdate.Table = ConvertCommandToMessageTable(cmd.(*command.Update).Table)
	msgCmdUpdate.Updates = ConvertCommandToMessageUpdateSetterSlice(cmd.(*command.Update).Updates)
	msgCmdUpdate.Filter = ConvertCommandToMessageExpr(cmd.(*command.Update).Filter)
	return msgCmdUpdate
}

// ConvertCommandToMessageJoinType converts command.JoinType to message.JoinType
func ConvertCommandToMessageJoinType(cmd command.JoinType) JoinType {
	switch cmd {
	case 0:
		return JoinType_JoinUnknown
	case 1:
		return JoinType_JoinLeft
	case 2:
		return JoinType_JoinLeftOuter
	case 3:
		return JoinType_JoinInner
	case 4:
		return JoinType_JoinCross
	}
	return -1
}

// ConvertCommandToMessageJoin converts a Command type to a Command_Join type.
func ConvertCommandToMessageJoin(cmd command.Command) *Command_Join {
	msgCmdJoin := &Command_Join{}
	msgCmdJoin.Natural = cmd.(*command.Join).Natural
	msgCmdJoin.Type = ConvertCommandToMessageJoinType(cmd.(*command.Join).Type)
	msgCmdJoin.Filter = ConvertCommandToMessageExpr(cmd.(*command.Join).Filter)
	msgCmdJoin.Left = ConvertCommandToMessageList(cmd.(*command.Join).Left)
	msgCmdJoin.Right = ConvertCommandToMessageList(cmd.(*command.Join).Right)
	return msgCmdJoin
}

// ConvertCommandToMessageLimit converts a Command type to a Command_Limit type.
func ConvertCommandToMessageLimit(cmd command.Command) *Command_Limit {
	msgCmdLimit := &Command_Limit{}
	msgCmdLimit.Limit = ConvertCommandToMessageExpr(cmd.(*command.Limit).Limit)
	msgCmdLimit.Input = ConvertCommandToMessageList(cmd.(*command.Limit).Input)
	return msgCmdLimit
}

// ConvertCommandToMessageInsertOr converts command.InsertOr to a message.InsertOr
func ConvertCommandToMessageInsertOr(cmd command.InsertOr) InsertOr {
	switch cmd {
	case 0:
		return InsertOr_InsertOrUnknown
	case 1:
		return InsertOr_InsertOrReplace
	case 2:
		return InsertOr_InsertOrRollback
	case 3:
		return InsertOr_InsertOrAbort
	case 4:
		return InsertOr_InsertOrFail
	case 5:
		return InsertOr_InsertOrIgnore
	}
	return -1
}

// ConvertCommandToMessageInsert converts a Command type to a Command_Insert type.
func ConvertCommandToMessageInsert(cmd command.Command) *Command_Insert {
	msgCmdInsert := &Command_Insert{}
	msgCmdInsert.InsertOr = ConvertCommandToMessageInsertOr(cmd.(*command.Insert).InsertOr)
	msgCmdInsert.Table = ConvertCommandToMessageTable(cmd.(*command.Insert).Table)
	msgCmdInsert.Cols = ConvertCommandToMessageCols(cmd.(*command.Insert).Cols)
	msgCmdInsert.DefaultValues = cmd.(*command.Insert).DefaultValues
	msgCmdInsert.Input = ConvertCommandToMessageList(cmd.(*command.Insert).Input)
	return msgCmdInsert
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
		return ConvertMessageToCommandUpdate(m)
	case *Command_Join:
		return ConvertMessageToCommandJoin(m)
	case *Command_Limit:
		return ConvertMessageToCommandLimit(m)
	case *Command_Insert:
		return ConvertMessageToCommandInsert(m)
	}
	return nil
}

// ConvertMessageToCommandTable converts a message.SimpleTable to a command.Table
func ConvertMessageToCommandTable(msg *SimpleTable) command.Table {
	cmdTable := &command.SimpleTable{}
	cmdTable.Schema = msg.Schema
	cmdTable.Table = msg.Table
	cmdTable.Alias = msg.Alias
	cmdTable.Indexed = msg.Indexed
	cmdTable.Index = msg.Index
	return cmdTable
}

// ConvertMessageToCommandScan converts a message.Command_Scan to a command.Scan
func ConvertMessageToCommandScan(msg *Command_Scan) *command.Scan {
	cmdScan := &command.Scan{}
	cmdScan.Table = ConvertMessageToCommandTable(msg.Table)
	return cmdScan
}

// ConvertMessageToCommandLiteralExpr converts a message.Expr to a command.LiteralExpr
func ConvertMessageToCommandLiteralExpr(msg *Expr) *command.LiteralExpr {
	literalExpr := &command.LiteralExpr{}
	literalExpr.Value = msg.GetLiteral().GetValue()
	return literalExpr
}

// ConvertMessageToCommandConstantBooleanExpr converts a message.Expr to a command.ConstantBooleanExpr
func ConvertMessageToCommandConstantBooleanExpr(msg *Expr) *command.ConstantBooleanExpr {
	constantBooleanExpr := &command.ConstantBooleanExpr{}
	constantBooleanExpr.Value = msg.GetConstant().GetValue()
	return constantBooleanExpr
}

// ConvertMessageToCommandUnaryExpr converts a message.Expr to a command.UnaryExpr
func ConvertMessageToCommandUnaryExpr(msg *Expr) *command.UnaryExpr {
	unaryExpr := &command.UnaryExpr{}
	unaryExpr.Operator = msg.GetUnary().GetOperator()
	unaryExpr.Value = ConvertMessageToCommandExpr(msg.GetUnary().GetValue())
	return unaryExpr
}

// ConvertMessageToCommandBinaryExpr converts a message.Expr to a command.BinaryExpr
func ConvertMessageToCommandBinaryExpr(msg *Expr) *command.BinaryExpr {
	binaryExpr := &command.BinaryExpr{}
	binaryExpr.Operator = msg.GetBinary().GetOperator()
	binaryExpr.Left = ConvertMessageToCommandExpr(msg.GetBinary().GetLeft())
	binaryExpr.Right = ConvertMessageToCommandExpr(msg.GetBinary().GetRight())
	return binaryExpr
}

// ConvertMessageToCommandExprSlice converts a []*message.Expr to []command.Expr
func ConvertMessageToCommandExprSlice(msg []*Expr) []command.Expr {
	msgExprSlice := []command.Expr{}
	for i := range msg {
		msgExprSlice = append(msgExprSlice, ConvertMessageToCommandExpr(msg[i]))
	}
	return msgExprSlice
}

// ConvertMessageToCommandFunctionExpr converts a message.Expr to a command.FunctionExpr
func ConvertMessageToCommandFunctionExpr(msg *Expr) *command.FunctionExpr {
	functionExpr := &command.FunctionExpr{}
	functionExpr.Name = msg.GetFunc().GetName()
	functionExpr.Distinct = msg.GetFunc().GetDistinct()
	functionExpr.Args = ConvertMessageToCommandExprSlice(msg.GetFunc().GetArgs())
	return functionExpr
}

// ConvertMessageToCommandEqualityExpr converts a message.Expr to a command.EqualityExpr
func ConvertMessageToCommandEqualityExpr(msg *Expr) *command.EqualityExpr {
	equalityExpr := &command.EqualityExpr{}
	equalityExpr.Left = ConvertMessageToCommandExpr(msg.GetEquality().GetLeft())
	equalityExpr.Right = ConvertMessageToCommandExpr(msg.GetEquality().GetRight())
	equalityExpr.Invert = msg.GetEquality().Invert
	return equalityExpr
}

// ConvertMessageToCommandRangeExpr converts a message.Expr to a command.RangeExpr
func ConvertMessageToCommandRangeExpr(msg *Expr) *command.RangeExpr {
	rangeExpr := &command.RangeExpr{}
	rangeExpr.Needle = ConvertMessageToCommandExpr(msg.GetRange().GetNeedle())
	rangeExpr.Lo = ConvertMessageToCommandExpr(msg.GetRange().GetLo())
	rangeExpr.Hi = ConvertMessageToCommandExpr(msg.GetRange().GetHi())
	return rangeExpr
}

// ConvertMessageToCommandExpr converts a message.Expr to a command.Expr
func ConvertMessageToCommandExpr(msg *Expr) command.Expr {
	if msg == nil {
		return nil
	}
	switch msg.Expr.(type) {
	case *Expr_Literal:
		return ConvertMessageToCommandLiteralExpr(msg)
	case *Expr_Constant:
		return ConvertMessageToCommandConstantBooleanExpr(msg)
	case *Expr_Unary:
		return ConvertMessageToCommandUnaryExpr(msg)
	case *Expr_Binary:
		return ConvertMessageToCommandBinaryExpr(msg)
	case *Expr_Func:
		return ConvertMessageToCommandFunctionExpr(msg)
	case *Expr_Equality:
		return ConvertMessageToCommandEqualityExpr(msg)
	case *Expr_Range:
		return ConvertMessageToCommandRangeExpr(msg)
	}
	return nil
}

// ConvertMessageToCommandListScan converts a message.List to a command.Scan
func ConvertMessageToCommandListScan(msg *List) *command.Scan {
	cmdScan := &command.Scan{}
	cmdScan.Table = ConvertMessageToCommandTable(msg.GetScan().GetTable())
	return cmdScan
}

// ConvertMessageToCommandListSelect converts a message.List to a command.Select
func ConvertMessageToCommandListSelect(msg *List) *command.Select {
	cmdSelect := &command.Select{}
	cmdSelect.Filter = ConvertMessageToCommandExpr(msg.GetSelect().GetFilter())
	cmdSelect.Input = ConvertMessageToCommandList(msg.GetSelect().GetInput())
	return cmdSelect
}

// ConvertMessageToCommandListProject converts a message.List to a command.Project
func ConvertMessageToCommandListProject(msg *List) *command.Project {
	cmdProject := &command.Project{}
	cmdProject.Cols = ConvertMessageToCommandCols(msg.GetProject().GetCols())
	cmdProject.Input = ConvertMessageToCommandList(msg.GetProject().GetInput())
	return cmdProject
}

// ConvertMessageToCommandListJoin converts a message.List to a command.Join
func ConvertMessageToCommandListJoin(msg *List) *command.Join {
	cmdJoin := &command.Join{}
	cmdJoin.Natural = msg.GetJoin().GetNatural()
	cmdJoin.Type = ConvertMessageToCommandJoinType(msg.GetJoin().GetType())
	cmdJoin.Filter = ConvertMessageToCommandExpr(msg.GetJoin().GetFilter())
	cmdJoin.Left = ConvertMessageToCommandList(msg.GetJoin().GetLeft())
	cmdJoin.Right = ConvertMessageToCommandList(msg.GetJoin().GetRight())
	return cmdJoin
}

// ConvertMessageToCommandListLimit converts a message.List to a command.Limit
func ConvertMessageToCommandListLimit(msg *List) *command.Limit {
	cmdLimit := &command.Limit{}
	cmdLimit.Limit = ConvertMessageToCommandExpr(msg.GetLimit().GetLimit())
	cmdLimit.Input = ConvertMessageToCommandList(msg.GetLimit().GetInput())
	return cmdLimit
}

// ConvertMessageToCommandListOffset converts a message.List to a command.Offset
func ConvertMessageToCommandListOffset(msg *List) *command.Offset {
	cmdOffset := &command.Offset{}
	cmdOffset.Offset = ConvertMessageToCommandExpr(msg.GetOffset().GetOffset())
	cmdOffset.Input = ConvertMessageToCommandList(msg.GetDistinct().GetInput())
	return cmdOffset
}

// ConvertMessageToCommandListDistinct converts a message.List to a command.Distinct
func ConvertMessageToCommandListDistinct(msg *List) *command.Distinct {
	cmdDistinct := &command.Distinct{}
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
	if msg == nil {
		return nil
	}
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
func ConvertMessageToCommandSelect(msg *Command_Select) *command.Select {
	cmdSelect := &command.Select{}
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
func ConvertMessageToCommandProject(msg *Command_Project) *command.Project {
	cmdProject := &command.Project{}
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
	cmdDropIndex.Schema = msg.GetSchema()
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

// ConvertMessageToCommandUpdateSetterLiteralExpr converts message.LiteralExpr to command.Expr
func ConvertMessageToCommandUpdateSetterLiteralExpr(msg *LiteralExpr) command.Expr {
	cmdExpr := command.LiteralExpr{}
	cmdExpr.Value = msg.Value
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterConstantExpr converts message.ConstantBooleanExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterConstantExpr(msg *ConstantBooleanExpr) command.Expr {
	cmdExpr := command.ConstantBooleanExpr{}
	cmdExpr.Value = msg.Value
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterUnaryExpr converts message.UnaryExpr to command.Expr
func ConvertMessageToCommandUpdateSetterUnaryExpr(msg *UnaryExpr) command.Expr {
	cmdExpr := command.UnaryExpr{}
	cmdExpr.Operator = msg.Operator
	cmdExpr.Value = ConvertMessageToCommandBinaryExpr(msg.Value)
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterBinaryExpr converts message.BinaryExpr to command.Expr
func ConvertMessageToCommandUpdateSetterBinaryExpr(msg *BinaryExpr) command.Expr {
	cmdExpr := command.BinaryExpr{}
	cmdExpr.Operator = msg.Operator
	cmdExpr.Left = ConvertMessageToCommandBinaryExpr(msg.Left)
	cmdExpr.Right = ConvertMessageToCommandBinaryExpr(msg.Right)
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterFuncExpr converts message.FunctionExpr tp command.Expr
func ConvertMessageToCommandUpdateSetterFuncExpr(msg *FunctionExpr) command.Expr {
	cmdExpr := command.FunctionExpr{}
	cmdExpr.Name = msg.Name
	cmdExpr.Distinct = msg.Distinct
	cmdExpr.Args = ConvertMessageToCommandExprSlice(msg.Args)
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterEqualityExpr converts message.EqualityExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterEqualityExpr(msg *EqualityExpr) command.Expr {
	cmdExpr := command.EqualityExpr{}
	cmdExpr.Left = ConvertMessageToCommandBinaryExpr(msg.Left)
	cmdExpr.Right = ConvertMessageToCommandBinaryExpr(msg.Right)
	cmdExpr.Invert = msg.Invert
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterRangeExpr converts a message.RangeExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterRangeExpr(msg *RangeExpr) command.Expr {
	cmdExpr := command.RangeExpr{}
	cmdExpr.Needle = ConvertMessageToCommandBinaryExpr(msg.Needle)
	cmdExpr.Lo = ConvertMessageToCommandBinaryExpr(msg.Lo)
	cmdExpr.Hi = ConvertMessageToCommandBinaryExpr(msg.Hi)
	cmdExpr.Invert = msg.Invert
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetter converts a message.UpdateSetter to a command.UpdateSetter.
func ConvertMessageToCommandUpdateSetter(msg *UpdateSetter) command.UpdateSetter {
	cmdUpdateSetter := command.UpdateSetter{}
	cmdUpdateSetter.Cols = msg.Cols
	switch msg.Value.(type) {
	case *UpdateSetter_Literal:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterLiteralExpr(msg.GetLiteral())
	case *UpdateSetter_Constant:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterConstantExpr(msg.GetConstant())
	case *UpdateSetter_Unary:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterUnaryExpr(msg.GetUnary())
	case *UpdateSetter_Binary:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterBinaryExpr(msg.GetBinary())
	case *UpdateSetter_Func:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterFuncExpr(msg.GetFunc())
	case *UpdateSetter_Equality:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterEqualityExpr(msg.GetEquality())
	case *UpdateSetter_Range:
		cmdUpdateSetter.Value = ConvertMessageToCommandUpdateSetterRangeExpr(msg.GetRange())
	}
	return cmdUpdateSetter
}

// ConvertMessageToCommandUpdateSetterSlice converts a []message.UpdateSetter to a []command.UpdateSetter.
func ConvertMessageToCommandUpdateSetterSlice(msg []*UpdateSetter) []command.UpdateSetter {
	cmdUpdateSetterSlice := []command.UpdateSetter{}
	for i := range msg {
		cmdUpdateSetterSlice = append(cmdUpdateSetterSlice, ConvertMessageToCommandUpdateSetter(msg[i]))
	}
	return cmdUpdateSetterSlice
}

// ConvertMessageToCommandUpdate converts a message.Command_Update to a command.Update
func ConvertMessageToCommandUpdate(msg *Command_Update) command.Update {
	cmdUpdate := command.Update{}
	cmdUpdate.UpdateOr = ConvertMessageToCommandUpdateOr(msg.GetUpdateOr())
	cmdUpdate.Updates = ConvertMessageToCommandUpdateSetterSlice(msg.GetUpdates())
	cmdUpdate.Table = ConvertMessageToCommandTable(msg.GetTable())
	cmdUpdate.Filter = ConvertMessageToCommandExpr(msg.GetFilter())
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
