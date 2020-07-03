package message

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// ConvertCommandToMessage converts a command.Command to a message.Message.
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

// ConvertCommandToMessageTable converts a command.Table to a SimpleTable.
func ConvertCommandToMessageTable(cmd command.Table) *SimpleTable {
	simpleTable := &SimpleTable{
		Schema:  cmd.(*command.SimpleTable).Schema,
		Table:   cmd.(*command.SimpleTable).Table,
		Alias:   cmd.(*command.SimpleTable).Alias,
		Indexed: cmd.(*command.SimpleTable).Indexed,
		Index:   cmd.(*command.SimpleTable).Index,
	}
	return simpleTable
}

// ConvertCommandToMessageScan converts a Command type to a Command_Scan type.
func ConvertCommandToMessageScan(cmd command.Command) *Command_Scan {
	msgCmdScan := &Command_Scan{
		Table: ConvertCommandToMessageTable(cmd.(*command.Scan).Table),
	}
	return msgCmdScan
}

// ConvertCommandToMessageLiteralExpr converts a command.Expr to a message.Expr_Literal.
func ConvertCommandToMessageLiteralExpr(cmd *command.LiteralExpr) *Expr_Literal {
	msgExprLiteral := &Expr_Literal{
		&LiteralExpr{
			Value: cmd.Value,
		},
	}
	return msgExprLiteral
}

// ConvertCommandToMessageConstantBooleanExpr converts a command.Expr to a message.Expr_Constant.
func ConvertCommandToMessageConstantBooleanExpr(cmd *command.ConstantBooleanExpr) *Expr_Constant {
	msgExprConstant := &Expr_Constant{
		&ConstantBooleanExpr{
			Value: cmd.Value,
		},
	}
	return msgExprConstant
}

// ConvertCommandToMessageUnaryExpr converts a command.Expr to a message.Expr_Unary.
func ConvertCommandToMessageUnaryExpr(cmd *command.UnaryExpr) *Expr_Unary {
	msgExprUnary := &Expr_Unary{
		&UnaryExpr{
			Operator: cmd.Operator,
			Value:    ConvertCommandToMessageExpr(cmd.Value),
		},
	}
	return msgExprUnary
}

// ConvertCommandToMessageBinaryExpr converts a command.Expr to a message.Expr_Binary.
func ConvertCommandToMessageBinaryExpr(cmd *command.BinaryExpr) *Expr_Binary {
	msgExprBinary := &Expr_Binary{
		&BinaryExpr{
			Operator: cmd.Operator,
			Left:     ConvertCommandToMessageExpr(cmd.Left),
			Right:    ConvertCommandToMessageExpr(cmd.Right),
		},
	}
	return msgExprBinary
}

// ConvertCommandToMessageRepeatedExpr converts a []command.Expr to a message.Expr.
func ConvertCommandToMessageRepeatedExpr(cmd []command.Expr) []*Expr {
	msgRepeatedExpr := []*Expr{}
	for i := range cmd {
		msgRepeatedExpr = append(msgRepeatedExpr, ConvertCommandToMessageExpr(cmd[i]))
	}
	return msgRepeatedExpr
}

// ConvertCommandToMessageFunctionalExpr converts a command.Expr to a message.Expr_Func.
func ConvertCommandToMessageFunctionalExpr(cmd *command.FunctionExpr) *Expr_Func {
	msgExprFunc := &Expr_Func{
		&FunctionExpr{
			Name:     cmd.Name,
			Distinct: cmd.Distinct,
			Args:     ConvertCommandToMessageRepeatedExpr(cmd.Args),
		},
	}
	return msgExprFunc
}

// ConvertCommandToMessageEqualityExpr converts a command.Expr to a message.Expr_Equality.
func ConvertCommandToMessageEqualityExpr(cmd *command.EqualityExpr) *Expr_Equality {
	msgExprEquality := &Expr_Equality{
		&EqualityExpr{
			Left:   ConvertCommandToMessageExpr(cmd.Left),
			Right:  ConvertCommandToMessageExpr(cmd.Right),
			Invert: cmd.Invert,
		},
	}
	return msgExprEquality
}

// ConvertCommandToMessageRangeExpr converts a command.Expr to a message.Expr_Range.
func ConvertCommandToMessageRangeExpr(cmd *command.RangeExpr) *Expr_Range {
	msgExprRange := &Expr_Range{
		&RangeExpr{
			Needle: ConvertCommandToMessageExpr(cmd.Needle),
			Lo:     ConvertCommandToMessageExpr(cmd.Lo),
			Hi:     ConvertCommandToMessageExpr(cmd.Hi),
			Invert: cmd.Invert,
		},
	}
	return msgExprRange
}

// ConvertCommandToMessageExpr converts command.Expr to a message.Expr.
func ConvertCommandToMessageExpr(cmd command.Expr) *Expr {
	msgExpr := &Expr{}
	switch c := cmd.(type) {
	case *command.LiteralExpr:
		msgExpr.Expr = ConvertCommandToMessageLiteralExpr(c)
	case *command.ConstantBooleanExpr:
		msgExpr.Expr = ConvertCommandToMessageConstantBooleanExpr(c)
	case *command.UnaryExpr:
		msgExpr.Expr = ConvertCommandToMessageUnaryExpr(c)
	case *command.BinaryExpr:
		msgExpr.Expr = ConvertCommandToMessageBinaryExpr(c)
	case *command.FunctionExpr:
		msgExpr.Expr = ConvertCommandToMessageFunctionalExpr(c)
	case *command.EqualityExpr:
		msgExpr.Expr = ConvertCommandToMessageEqualityExpr(c)
	case *command.RangeExpr:
		msgExpr.Expr = ConvertCommandToMessageRangeExpr(c)
	}
	return msgExpr
}

// ConvertCommandToMessageListScan converts a command.Scan to a message.List_Scan.
func ConvertCommandToMessageListScan(cmd *command.Scan) *List_Scan {
	msgListScan := &List_Scan{
		&Command_Scan{
			Table: ConvertCommandToMessageTable(cmd.Table),
		},
	}
	return msgListScan
}

// ConvertCommandToMessageListSelect converts a command.Select to a message.List_Select.
func ConvertCommandToMessageListSelect(cmd *command.Select) *List_Select {
	msgListSelect := &List_Select{
		&Command_Select{
			Filter: ConvertCommandToMessageExpr(cmd.Filter),
			Input:  ConvertCommandToMessageList(cmd.Input),
		},
	}
	return msgListSelect
}

// ConvertCommandToMessageListProject converts a command.Project to a message.List_Project.
func ConvertCommandToMessageListProject(cmd *command.Project) *List_Project {
	msgListProject := &List_Project{
		&Command_Project{
			Cols:  ConvertCommandToMessageColSlice(cmd.Cols),
			Input: ConvertCommandToMessageList(cmd.Input),
		},
	}
	return msgListProject
}

// ConvertCommandToMessageListJoin converts a command.Join to a message.List_Join.
func ConvertCommandToMessageListJoin(cmd *command.Join) *List_Join {
	msgListJoin := &List_Join{
		&Command_Join{
			Natural: cmd.Natural,
			Type:    ConvertCommandToMessageJoinType(cmd.Type),
			Filter:  ConvertCommandToMessageExpr(cmd.Filter),
			Left:    ConvertCommandToMessageList(cmd.Left),
			Right:   ConvertCommandToMessageList(cmd.Right),
		},
	}
	return msgListJoin
}

// ConvertCommandToMessageListLimit converts a command.Limit to a message.List_Limit.
func ConvertCommandToMessageListLimit(cmd *command.Limit) *List_Limit {
	msgListLimit := &List_Limit{
		&Command_Limit{
			Limit: ConvertCommandToMessageExpr(cmd.Limit),
			Input: ConvertCommandToMessageList(cmd.Input),
		},
	}
	return msgListLimit
}

// ConvertCommandToMessageListOffset converts a command.Offset to a message.List_Offset.
func ConvertCommandToMessageListOffset(cmd *command.Offset) *List_Offset {
	msgListOffset := &List_Offset{
		&Command_Offset{
			Offset: ConvertCommandToMessageExpr(cmd.Offset),
			Input:  ConvertCommandToMessageList(cmd.Input),
		},
	}
	return msgListOffset
}

// ConvertCommandToMessageListDistinct converts a command.Distinct to a message.List_Distinct.
func ConvertCommandToMessageListDistinct(cmd *command.Distinct) *List_Distinct {
	msgListDistinct := &List_Distinct{
		&Command_Distinct{
			Input: ConvertCommandToMessageList(cmd.Input),
		},
	}
	return msgListDistinct
}

// ConvertCommandToMessageRepeatedExprSlice converts a [][]command.Expr to a [][]message.Expr.
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

// ConvertCommandToMessageListValues converts a command.Values to a message.List_Values.
func ConvertCommandToMessageListValues(cmd *command.Values) *List_Values {
	msgListValues := &List_Values{
		&Command_Values{
			Expr: ConvertCommandToMessageRepeatedExprSlice(cmd.Values),
		},
	}
	return msgListValues
}

// ConvertCommandToMessageList converts
func ConvertCommandToMessageList(cmd command.List) *List {
	msgList := &List{}
	switch c := cmd.(type) {
	case *command.Scan:
		msgList.List = ConvertCommandToMessageListScan(c)
	case *command.Select:
		msgList.List = ConvertCommandToMessageListSelect(c)
	case *command.Project:
		msgList.List = ConvertCommandToMessageListProject(c)
	case *command.Join:
		msgList.List = ConvertCommandToMessageListJoin(c)
	case *command.Limit:
		msgList.List = ConvertCommandToMessageListLimit(c)
	case *command.Offset:
		msgList.List = ConvertCommandToMessageListOffset(c)
	case *command.Distinct:
		msgList.List = ConvertCommandToMessageListDistinct(c)
	case *command.Values:
		msgList.List = ConvertCommandToMessageListValues(c)
	}
	return msgList
}

// ConvertCommandToMessageSelect converts a Command type to a Command_Select type.
func ConvertCommandToMessageSelect(cmd command.Command) *Command_Select {
	msgCmdSelect := &Command_Select{
		Filter: ConvertCommandToMessageExpr(cmd.(*command.Select).Filter),
		Input:  ConvertCommandToMessageList(cmd.(*command.Select).Input),
	}
	return msgCmdSelect
}

// ConvertCommandToMessageCol converts command.Column to a message.Column.
func ConvertCommandToMessageCol(cmd command.Column) *Column {
	msgCol := &Column{
		Table:  cmd.Table,
		Column: ConvertCommandToMessageExpr(cmd.Column),
		Alias:  cmd.Alias,
	}
	return msgCol
}

// ConvertCommandToMessageColSlice converts []command.Column to a []message.Column.
func ConvertCommandToMessageColSlice(cmd []command.Column) []*Column {
	msgCols := []*Column{}
	for i := range cmd {
		msgCols = append(msgCols, ConvertCommandToMessageCol(cmd[i]))
	}
	return msgCols
}

// ConvertCommandToMessageProject converts a Command type to a Command_Project type.
func ConvertCommandToMessageProject(cmd command.Command) *Command_Project {
	msgCmdProject := &Command_Project{
		Cols:  ConvertCommandToMessageColSlice(cmd.(*command.Project).Cols),
		Input: ConvertCommandToMessageList(cmd.(*command.Project).Input),
	}
	return msgCmdProject
}

// ConvertCommandToMessageDelete converts a Command type to a Command_Delete type.
func ConvertCommandToMessageDelete(cmd command.Command) *Command_Delete {
	msgCmdDelete := &Command_Delete{
		Table:  ConvertCommandToMessageTable(cmd.(*command.Delete).Table),
		Filter: ConvertCommandToMessageExpr(cmd.(*command.Delete).Filter),
	}
	return msgCmdDelete
}

// ConvertCommandToMessageDrop converts a Command type to a CommandDrop type.
func ConvertCommandToMessageDrop(cmd command.Command) *CommandDrop {
	msgCmdDrop := &CommandDrop{}
	switch cmd := cmd.(type) {
	case *command.DropTable:
		msgCmdDrop.Target = DropTarget_Table
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	case *command.DropView:
		msgCmdDrop.Target = DropTarget_View
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	case *command.DropIndex:
		msgCmdDrop.Target = DropTarget_Index
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	case *command.DropTrigger:
		msgCmdDrop.Target = DropTarget_Trigger
		msgCmdDrop.IfExists = cmd.IfExists
		msgCmdDrop.Schema = cmd.Schema
		msgCmdDrop.Name = cmd.Name
	}
	return msgCmdDrop
}

// ConvertCommandToMessageUpdateOr converts a command.Update or to a message.UpdateOr.
// Returns -1 if the UpdateOr type doesn't match.
func ConvertCommandToMessageUpdateOr(cmd command.UpdateOr) UpdateOr {
	switch cmd {
	case command.UpdateOrUnknown:
		return UpdateOr_UpdateOrUnknown
	case command.UpdateOrRollback:
		return UpdateOr_UpdateOrRollback
	case command.UpdateOrAbort:
		return UpdateOr_UpdateOrAbort
	case command.UpdateOrReplace:
		return UpdateOr_UpdateOrReplace
	case command.UpdateOrFail:
		return UpdateOr_UpdateOrFail
	case command.UpdateOrIgnore:
		return UpdateOr_UpdateOrIgnore
	}
	return -1
}

// ConvertCommandToMessageUpdateSetterLiteral converts a command.Literal to a message.UpdateSetter_Literal.
func ConvertCommandToMessageUpdateSetterLiteral(cmd command.LiteralExpr) *UpdateSetter_Literal {
	msgUpdateSetterLiteral := &UpdateSetter_Literal{
		&LiteralExpr{
			Value: cmd.Value,
		},
	}
	return msgUpdateSetterLiteral
}

// ConvertCommandToMessageUpdateSetterConstant converts a command.Constant to a message.UpdateSetter_Constant.
func ConvertCommandToMessageUpdateSetterConstant(cmd command.ConstantBooleanExpr) *UpdateSetter_Constant {
	msgUpdateSetterConstant := &UpdateSetter_Constant{
		&ConstantBooleanExpr{
			Value: cmd.Value,
		},
	}
	return msgUpdateSetterConstant
}

// ConvertCommandToMessageUpdateSetterUnary converts a command.Unary to a message.UpdateSetter_Unary.
func ConvertCommandToMessageUpdateSetterUnary(cmd command.UnaryExpr) *UpdateSetter_Unary {
	msgUpdateSetterUnary := &UpdateSetter_Unary{
		&UnaryExpr{
			Operator: cmd.Operator,
			Value:    ConvertCommandToMessageExpr(cmd.Value),
		},
	}
	return msgUpdateSetterUnary
}

// ConvertCommandToMessageUpdateSetterBinary converts a command.Binary to a message.UpdateSetter_Binary.
func ConvertCommandToMessageUpdateSetterBinary(cmd command.BinaryExpr) *UpdateSetter_Binary {
	msgUpdateSetterBinary := &UpdateSetter_Binary{
		&BinaryExpr{
			Operator: cmd.Operator,
			Left:     ConvertCommandToMessageExpr(cmd.Left),
			Right:    ConvertCommandToMessageExpr(cmd.Right),
		},
	}

	return msgUpdateSetterBinary
}

// ConvertCommandToMessageUpdateSetterFunc converts a command.Func to a message.UpdateSetter_Func.
func ConvertCommandToMessageUpdateSetterFunc(cmd command.FunctionExpr) *UpdateSetter_Func {
	msgUpdateSetterFunc := &UpdateSetter_Func{
		&FunctionExpr{
			Name:     cmd.Name,
			Distinct: cmd.Distinct,
			Args:     ConvertCommandToMessageRepeatedExpr(cmd.Args),
		},
	}
	return msgUpdateSetterFunc
}

// ConvertCommandToMessageUpdateSetterEquality converts a command.Equality to a message.UpdateSetter_Equality.
func ConvertCommandToMessageUpdateSetterEquality(cmd command.EqualityExpr) *UpdateSetter_Equality {
	msgUpdateSetterEquality := &UpdateSetter_Equality{
		&EqualityExpr{
			Left:   ConvertCommandToMessageExpr(cmd.Left),
			Right:  ConvertCommandToMessageExpr(cmd.Right),
			Invert: cmd.Invert,
		},
	}
	return msgUpdateSetterEquality
}

// ConvertCommandToMessageUpdateSetterRange converts a command.Range to a message.UpdateSetter_Range.
func ConvertCommandToMessageUpdateSetterRange(cmd command.RangeExpr) *UpdateSetter_Range {
	msgUpdateSetterRange := &UpdateSetter_Range{
		&RangeExpr{
			Needle: ConvertCommandToMessageExpr(cmd.Needle),
			Lo:     ConvertCommandToMessageExpr(cmd.Lo),
			Hi:     ConvertCommandToMessageExpr(cmd.Hi),
			Invert: cmd.Invert,
		},
	}

	return msgUpdateSetterRange
}

// ConvertCommandToMessageUpdateSetter converts a command.UpdateSetter to a message.UpdateSetter.
func ConvertCommandToMessageUpdateSetter(cmd command.UpdateSetter) *UpdateSetter {
	msgUpdateSetter := &UpdateSetter{}
	msgUpdateSetter.Cols = cmd.Cols
	switch val := cmd.Value.(type) {
	case command.LiteralExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterLiteral(val)
	case command.ConstantBooleanExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterConstant(val)
	case command.UnaryExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterUnary(val)
	case command.BinaryExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterBinary(val)
	case command.FunctionExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterFunc(val)
	case command.EqualityExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterEquality(val)
	case command.RangeExpr:
		msgUpdateSetter.Value = ConvertCommandToMessageUpdateSetterRange(val)
	}
	return msgUpdateSetter
}

// ConvertCommandToMessageUpdateSetterSlice converts a []command.UpdateSetter to a []message.UpdateSetter.
func ConvertCommandToMessageUpdateSetterSlice(cmd []command.UpdateSetter) []*UpdateSetter {
	msgUpdateSetterSlice := []*UpdateSetter{}
	for i := range cmd {
		msgUpdateSetterSlice = append(msgUpdateSetterSlice, ConvertCommandToMessageUpdateSetter(cmd[i]))
	}
	return msgUpdateSetterSlice
}

// ConvertCommandToMessageUpdate converts a Command type to a Command_Update type.
func ConvertCommandToMessageUpdate(cmd command.Command) *Command_Update {
	msgCmdUpdate := &Command_Update{
		UpdateOr: ConvertCommandToMessageUpdateOr(cmd.(*command.Update).UpdateOr),
		Table:    ConvertCommandToMessageTable(cmd.(*command.Update).Table),
		Updates:  ConvertCommandToMessageUpdateSetterSlice(cmd.(*command.Update).Updates),
		Filter:   ConvertCommandToMessageExpr(cmd.(*command.Update).Filter),
	}

	return msgCmdUpdate
}

// ConvertCommandToMessageJoinType converts command.JoinType to message.JoinType.
// It returns -1 on not finding a valid JoinType.
func ConvertCommandToMessageJoinType(cmd command.JoinType) JoinType {
	switch cmd {
	case command.JoinUnknown:
		return JoinType_JoinUnknown
	case command.JoinLeft:
		return JoinType_JoinLeft
	case command.JoinLeftOuter:
		return JoinType_JoinLeftOuter
	case command.JoinInner:
		return JoinType_JoinInner
	case command.JoinCross:
		return JoinType_JoinCross
	}
	return -1
}

// ConvertCommandToMessageJoin converts a Command type to a Command_Join type.
func ConvertCommandToMessageJoin(cmd command.Command) *Command_Join {
	msgCmdJoin := &Command_Join{
		Natural: cmd.(*command.Join).Natural,
		Type:    ConvertCommandToMessageJoinType(cmd.(*command.Join).Type),
		Filter:  ConvertCommandToMessageExpr(cmd.(*command.Join).Filter),
		Left:    ConvertCommandToMessageList(cmd.(*command.Join).Left),
		Right:   ConvertCommandToMessageList(cmd.(*command.Join).Right),
	}
	return msgCmdJoin
}

// ConvertCommandToMessageLimit converts a Command type to a Command_Limit type.
func ConvertCommandToMessageLimit(cmd command.Command) *Command_Limit {
	msgCmdLimit := &Command_Limit{
		Limit: ConvertCommandToMessageExpr(cmd.(*command.Limit).Limit),
		Input: ConvertCommandToMessageList(cmd.(*command.Limit).Input),
	}
	return msgCmdLimit
}

// ConvertCommandToMessageInsertOr converts command.InsertOr to a message.InsertOr.
// It returns -1 on not finding the right InsertOr type.
func ConvertCommandToMessageInsertOr(cmd command.InsertOr) InsertOr {
	switch cmd {
	case command.InsertOrUnknown:
		return InsertOr_InsertOrUnknown
	case command.InsertOrReplace:
		return InsertOr_InsertOrReplace
	case command.InsertOrRollback:
		return InsertOr_InsertOrRollback
	case command.InsertOrAbort:
		return InsertOr_InsertOrAbort
	case command.InsertOrFail:
		return InsertOr_InsertOrFail
	case command.InsertOrIgnore:
		return InsertOr_InsertOrIgnore
	}
	return -1
}

// ConvertCommandToMessageInsert converts a Command type to a Command_Insert type.
func ConvertCommandToMessageInsert(cmd command.Command) *Command_Insert {
	msgCmdInsert := &Command_Insert{
		InsertOr:      ConvertCommandToMessageInsertOr(cmd.(*command.Insert).InsertOr),
		Table:         ConvertCommandToMessageTable(cmd.(*command.Insert).Table),
		Cols:          ConvertCommandToMessageColSlice(cmd.(*command.Insert).Cols),
		DefaultValues: cmd.(*command.Insert).DefaultValues,
		Input:         ConvertCommandToMessageList(cmd.(*command.Insert).Input),
	}
	return msgCmdInsert
}

// ConvertMessageToCommand converts a message.Command to a command.Command.
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

// ConvertMessageToCommandTable converts a message.SimpleTable to a command.Table.
func ConvertMessageToCommandTable(msg *SimpleTable) command.Table {
	cmdTable := &command.SimpleTable{
		Schema:  msg.Schema,
		Table:   msg.Table,
		Alias:   msg.Alias,
		Indexed: msg.Indexed,
		Index:   msg.Index,
	}
	return cmdTable
}

// ConvertMessageToCommandScan converts a message.Command_Scan to a command.Scan.
func ConvertMessageToCommandScan(msg *Command_Scan) *command.Scan {
	cmdScan := &command.Scan{
		Table: ConvertMessageToCommandTable(msg.Table),
	}
	return cmdScan
}

// ConvertMessageToCommandLiteralExpr converts a message.Expr to a command.LiteralExpr.
func ConvertMessageToCommandLiteralExpr(msg *Expr) *command.LiteralExpr {
	literalExpr := &command.LiteralExpr{
		Value: msg.GetLiteral().GetValue(),
	}
	return literalExpr
}

// ConvertMessageToCommandConstantBooleanExpr converts a message.Expr to a command.ConstantBooleanExpr.
func ConvertMessageToCommandConstantBooleanExpr(msg *Expr) *command.ConstantBooleanExpr {
	constantBooleanExpr := &command.ConstantBooleanExpr{
		Value: msg.GetConstant().GetValue(),
	}
	return constantBooleanExpr
}

// ConvertMessageToCommandUnaryExpr converts a message.Expr to a command.UnaryExpr.
func ConvertMessageToCommandUnaryExpr(msg *Expr) *command.UnaryExpr {
	unaryExpr := &command.UnaryExpr{
		Operator: msg.GetUnary().GetOperator(),
		Value:    ConvertMessageToCommandExpr(msg.GetUnary().GetValue()),
	}
	return unaryExpr
}

// ConvertMessageToCommandBinaryExpr converts a message.Expr to a command.BinaryExpr.
func ConvertMessageToCommandBinaryExpr(msg *Expr) *command.BinaryExpr {
	binaryExpr := &command.BinaryExpr{
		Operator: msg.GetBinary().GetOperator(),
		Left:     ConvertMessageToCommandExpr(msg.GetBinary().GetLeft()),
		Right:    ConvertMessageToCommandExpr(msg.GetBinary().GetRight()),
	}
	return binaryExpr
}

// ConvertMessageToCommandExprSlice converts a []*message.Expr to []command.Expr.
func ConvertMessageToCommandExprSlice(msg []*Expr) []command.Expr {
	msgExprSlice := []command.Expr{}
	for i := range msg {
		msgExprSlice = append(msgExprSlice, ConvertMessageToCommandExpr(msg[i]))
	}
	return msgExprSlice
}

// ConvertMessageToCommandFunctionExpr converts a message.Expr to a command.FunctionExpr.
func ConvertMessageToCommandFunctionExpr(msg *Expr) *command.FunctionExpr {
	functionExpr := &command.FunctionExpr{

		Name:     msg.GetFunc().GetName(),
		Distinct: msg.GetFunc().GetDistinct(),
		Args:     ConvertMessageToCommandExprSlice(msg.GetFunc().GetArgs()),
	}
	return functionExpr
}

// ConvertMessageToCommandEqualityExpr converts a message.Expr to a command.EqualityExpr.
func ConvertMessageToCommandEqualityExpr(msg *Expr) *command.EqualityExpr {
	equalityExpr := &command.EqualityExpr{
		Left:   ConvertMessageToCommandExpr(msg.GetEquality().GetLeft()),
		Right:  ConvertMessageToCommandExpr(msg.GetEquality().GetRight()),
		Invert: msg.GetEquality().Invert,
	}
	return equalityExpr
}

// ConvertMessageToCommandRangeExpr converts a message.Expr to a command.RangeExpr.
func ConvertMessageToCommandRangeExpr(msg *Expr) *command.RangeExpr {
	rangeExpr := &command.RangeExpr{
		Needle: ConvertMessageToCommandExpr(msg.GetRange().GetNeedle()),
		Lo:     ConvertMessageToCommandExpr(msg.GetRange().GetLo()),
		Hi:     ConvertMessageToCommandExpr(msg.GetRange().GetHi()),
	}
	return rangeExpr
}

// ConvertMessageToCommandExpr converts a message.Expr to a command.Expr.
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

// ConvertMessageToCommandListScan converts a message.List to a command.Scan.
func ConvertMessageToCommandListScan(msg *List) *command.Scan {
	cmdScan := &command.Scan{
		Table: ConvertMessageToCommandTable(msg.GetScan().GetTable()),
	}
	return cmdScan
}

// ConvertMessageToCommandListSelect converts a message.List to a command.Select.
func ConvertMessageToCommandListSelect(msg *List) *command.Select {
	cmdSelect := &command.Select{
		Filter: ConvertMessageToCommandExpr(msg.GetSelect().GetFilter()),
		Input:  ConvertMessageToCommandList(msg.GetSelect().GetInput()),
	}
	return cmdSelect
}

// ConvertMessageToCommandListProject converts a message.List to a command.Project.
func ConvertMessageToCommandListProject(msg *List) *command.Project {
	cmdProject := &command.Project{
		Cols:  ConvertMessageToCommandCols(msg.GetProject().GetCols()),
		Input: ConvertMessageToCommandList(msg.GetProject().GetInput()),
	}
	return cmdProject
}

// ConvertMessageToCommandListJoin converts a message.List to a command.Join.
func ConvertMessageToCommandListJoin(msg *List) *command.Join {
	cmdJoin := &command.Join{
		Natural: msg.GetJoin().GetNatural(),
		Type:    ConvertMessageToCommandJoinType(msg.GetJoin().GetType()),
		Filter:  ConvertMessageToCommandExpr(msg.GetJoin().GetFilter()),
		Left:    ConvertMessageToCommandList(msg.GetJoin().GetLeft()),
		Right:   ConvertMessageToCommandList(msg.GetJoin().GetRight()),
	}
	return cmdJoin
}

// ConvertMessageToCommandListLimit converts a message.List to a command.Limit.
func ConvertMessageToCommandListLimit(msg *List) *command.Limit {
	cmdLimit := &command.Limit{
		Limit: ConvertMessageToCommandExpr(msg.GetLimit().GetLimit()),
		Input: ConvertMessageToCommandList(msg.GetLimit().GetInput()),
	}
	return cmdLimit
}

// ConvertMessageToCommandListOffset converts a message.List to a command.Offset.
func ConvertMessageToCommandListOffset(msg *List) *command.Offset {
	cmdOffset := &command.Offset{
		Offset: ConvertMessageToCommandExpr(msg.GetOffset().GetOffset()),
		Input:  ConvertMessageToCommandList(msg.GetDistinct().GetInput()),
	}
	return cmdOffset
}

// ConvertMessageToCommandListDistinct converts a message.List to a command.Distinct.
func ConvertMessageToCommandListDistinct(msg *List) *command.Distinct {
	cmdDistinct := &command.Distinct{
		Input: ConvertMessageToCommandList(msg.GetDistinct().GetInput()),
	}
	return cmdDistinct
}

// ConvertMessageToCommandExprRepeatedSlice converts a message.RepeatedExpr to a [][]command.Expr.
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

// ConvertMessageToCommandListValues converts a message.List to a command.Values.
func ConvertMessageToCommandListValues(msg *List) command.Values {
	cmdValues := command.Values{
		Values: ConvertMessageToCommandExprRepeatedSlice(msg.GetValues().GetExpr()),
	}
	return cmdValues
}

// ConvertMessageToCommandList converts a message.List to a command.List.
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
	cmdSelect := &command.Select{
		Filter: ConvertMessageToCommandExpr(msg.GetFilter()),
		Input:  ConvertMessageToCommandList(msg.GetInput()),
	}
	return cmdSelect
}

// ConvertMessageToCommandCol converts a message.Column to a command.Column
func ConvertMessageToCommandCol(msg *Column) command.Column {
	cmdCol := command.Column{
		Table:  msg.GetTable(),
		Column: ConvertMessageToCommandExpr(msg.GetColumn()),
		Alias:  msg.GetAlias(),
	}
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
	cmdDelete := command.Delete{
		Filter: ConvertMessageToCommandExpr(msg.GetFilter()),
		Table:  ConvertMessageToCommandTable(msg.GetTable()),
	}
	return cmdDelete
}

// ConvertMessageToCommandDropTable converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropTable(msg *CommandDrop) command.DropTable {
	cmdDropTable := command.DropTable{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
	return cmdDropTable
}

// ConvertMessageToCommandDropView converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropView(msg *CommandDrop) command.DropView {
	cmdDropView := command.DropView{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
	return cmdDropView
}

// ConvertMessageToCommandDropIndex converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropIndex(msg *CommandDrop) command.DropIndex {
	cmdDropIndex := command.DropIndex{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
	return cmdDropIndex
}

// ConvertMessageToCommandDropTrigger converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropTrigger(msg *CommandDrop) command.DropTrigger {
	cmdDropTrigger := command.DropTrigger{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
	return cmdDropTrigger
}

// ConvertMessageToCommandUpdateOr converts a message.UpdateOr to command.UpdateOr
func ConvertMessageToCommandUpdateOr(msg UpdateOr) command.UpdateOr {
	return command.UpdateOr(msg.Number())
}

// ConvertMessageToCommandUpdateSetterLiteralExpr converts message.LiteralExpr to command.Expr
func ConvertMessageToCommandUpdateSetterLiteralExpr(msg *LiteralExpr) command.Expr {
	cmdExpr := command.LiteralExpr{
		Value: msg.Value,
	}
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterConstantExpr converts message.ConstantBooleanExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterConstantExpr(msg *ConstantBooleanExpr) command.Expr {
	cmdExpr := command.ConstantBooleanExpr{
		Value: msg.Value,
	}
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterUnaryExpr converts message.UnaryExpr to command.Expr
func ConvertMessageToCommandUpdateSetterUnaryExpr(msg *UnaryExpr) command.Expr {
	cmdExpr := command.UnaryExpr{
		Operator: msg.Operator,
		Value:    ConvertMessageToCommandBinaryExpr(msg.Value),
	}
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterBinaryExpr converts message.BinaryExpr to command.Expr
func ConvertMessageToCommandUpdateSetterBinaryExpr(msg *BinaryExpr) command.Expr {
	cmdExpr := command.BinaryExpr{
		Operator: msg.Operator,
		Left:     ConvertMessageToCommandBinaryExpr(msg.Left),
		Right:    ConvertMessageToCommandBinaryExpr(msg.Right),
	}
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterFuncExpr converts message.FunctionExpr tp command.Expr
func ConvertMessageToCommandUpdateSetterFuncExpr(msg *FunctionExpr) command.Expr {
	cmdExpr := command.FunctionExpr{
		Name:     msg.Name,
		Distinct: msg.Distinct,
		Args:     ConvertMessageToCommandExprSlice(msg.Args),
	}
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterEqualityExpr converts message.EqualityExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterEqualityExpr(msg *EqualityExpr) command.Expr {
	cmdExpr := command.EqualityExpr{
		Left:   ConvertMessageToCommandBinaryExpr(msg.Left),
		Right:  ConvertMessageToCommandBinaryExpr(msg.Right),
		Invert: msg.Invert,
	}
	return cmdExpr
}

// ConvertMessageToCommandUpdateSetterRangeExpr converts a message.RangeExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterRangeExpr(msg *RangeExpr) command.Expr {
	cmdExpr := command.RangeExpr{
		Needle: ConvertMessageToCommandBinaryExpr(msg.Needle),
		Lo:     ConvertMessageToCommandBinaryExpr(msg.Lo),
		Hi:     ConvertMessageToCommandBinaryExpr(msg.Hi),
		Invert: msg.Invert,
	}
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
	cmdUpdate := command.Update{
		UpdateOr: ConvertMessageToCommandUpdateOr(msg.GetUpdateOr()),
		Updates:  ConvertMessageToCommandUpdateSetterSlice(msg.GetUpdates()),
		Table:    ConvertMessageToCommandTable(msg.GetTable()),
		Filter:   ConvertMessageToCommandExpr(msg.GetFilter()),
	}
	return cmdUpdate
}

// ConvertMessageToCommandJoinType converts a message.JoinType to a command.JoinType
func ConvertMessageToCommandJoinType(msg JoinType) command.JoinType {
	return command.JoinType(msg.Number())
}

// ConvertMessageToCommandJoin converts a message.Command_Join to a command.Join
func ConvertMessageToCommandJoin(msg *Command_Join) command.Join {
	cmdJoin := command.Join{
		Natural: msg.Natural,
		Type:    ConvertMessageToCommandJoinType(msg.GetType()),
		Filter:  ConvertMessageToCommandExpr(msg.GetFilter()),
		Left:    ConvertMessageToCommandList(msg.GetLeft()),
		Right:   ConvertMessageToCommandList(msg.GetRight()),
	}
	return cmdJoin
}

// ConvertMessageToCommandLimit converts a message.Command_Limit to a command.Limit
func ConvertMessageToCommandLimit(msg *Command_Limit) command.Limit {
	cmdLimit := command.Limit{
		Limit: ConvertMessageToCommandExpr(msg.GetLimit()),
		Input: ConvertMessageToCommandList(msg.GetInput()),
	}
	return cmdLimit
}

// ConvertMessageToCommandInsertOr converts a message.InsertOr to command.InsertOr
func ConvertMessageToCommandInsertOr(msg InsertOr) command.InsertOr {
	return command.InsertOr(msg.Number())
}

// ConvertMessageToCommandInsert converts a message.Command_Insert to a command.Insert
func ConvertMessageToCommandInsert(msg *Command_Insert) command.Insert {
	cmdInsert := command.Insert{
		InsertOr:      ConvertMessageToCommandInsertOr(msg.GetInsertOr()),
		Table:         ConvertMessageToCommandTable(msg.GetTable()),
		Cols:          ConvertMessageToCommandCols(msg.GetCols()),
		DefaultValues: msg.GetDefaultValues(),
		Input:         ConvertMessageToCommandList(msg.GetInput()),
	}
	return cmdInsert
}
