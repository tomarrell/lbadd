package message

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// ConvertCommandToMessage converts a command.Command to a message.Message.
func ConvertCommandToMessage(cmd command.Command) (Message, error) {
	if cmd == nil {
		return nil, nil
	}
	switch c := cmd.(type) {
	case command.Scan:
		return ConvertCommandScanToMessageScan(c)
	case command.Select:
		return ConvertCommandSelectToMessageSelect(c)
	case command.Project:
		return ConvertCommandProjectToMessageProject(c)
	case command.Delete:
		return ConvertCommandDeleteToMessageDelete(c)
	case *command.DropIndex:
		return ConvertCommandDropToMessageDrop(c)
	case *command.DropTable:
		return ConvertCommandDropToMessageDrop(c)
	case *command.DropTrigger:
		return ConvertCommandDropToMessageDrop(c)
	case *command.DropView:
		return ConvertCommandDropToMessageDrop(c)
	case command.Update:
		return ConvertCommandUpdateToMessageUpdate(c)
	case command.Join:
		return ConvertCommandJoinToMessageJoin(c)
	case command.Limit:
		return ConvertCommandLimitToMessageLimit(c)
	case command.Insert:
		return ConvertCommandInsertToMessageInsert(c)
	}
	return nil, ErrUnknownCommandKind
}

// ConvertCommandTableToMessageTable converts a command.Table to a SimpleTable.
func ConvertCommandTableToMessageTable(cmd command.Table) (*SimpleTable, error) {
	if cmd == nil {
		//TODO
		return nil, nil
	}
	return &SimpleTable{
		Schema:  cmd.(*command.SimpleTable).Schema,
		Table:   cmd.(*command.SimpleTable).Table,
		Alias:   cmd.(*command.SimpleTable).Alias,
		Indexed: cmd.(*command.SimpleTable).Indexed,
		Index:   cmd.(*command.SimpleTable).Index,
	}, nil
}

// ConvertCommandScanToMessageScan converts a Command type to a Command_Scan type.
func ConvertCommandScanToMessageScan(cmd command.Scan) (*Command_Scan, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	return &Command_Scan{
		Table: table,
	}, nil
}

// ConvertCommandLiteralExprToMessageLiteralExpr converts a command.Expr to a message.Expr_Literal.
func ConvertCommandLiteralExprToMessageLiteralExpr(cmd command.LiteralExpr) (*Expr_Literal, error) {
	return &Expr_Literal{
		&LiteralExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandConstantBooleanExprToMessageConstantBooleanExpr converts a command.Expr to a message.Expr_Constant.
func ConvertCommandConstantBooleanExprToMessageConstantBooleanExpr(cmd command.ConstantBooleanExpr) (*Expr_Constant, error) {
	return &Expr_Constant{
		&ConstantBooleanExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandUnaryExprToMessageUnaryExpr converts a command.Expr to a message.Expr_Unary.
func ConvertCommandUnaryExprToMessageUnaryExpr(cmd command.UnaryExpr) (*Expr_Unary, error) {
	val, err := ConvertCommandExprToMessageExpr(cmd.Value)
	if err != nil {
		return nil, err
	}
	return &Expr_Unary{
		&UnaryExpr{
			Operator: cmd.Operator,
			Value:    val,
		},
	}, nil
}

// ConvertCommandBinaryExprToMessageBinaryExpr converts a command.Expr to a message.Expr_Binary.
func ConvertCommandBinaryExprToMessageBinaryExpr(cmd command.BinaryExpr) (*Expr_Binary, error) {
	left, err := ConvertCommandExprToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandExprToMessageExpr(cmd.Right)
	if err != nil {
		return nil, err
	}
	return &Expr_Binary{
		&BinaryExpr{
			Operator: cmd.Operator,
			Left:     left,
			Right:    right,
		},
	}, nil
}

// ConvertCommandRepeatedExprToMessageRepeatedExpr converts a []command.Expr to a message.Expr.
func ConvertCommandRepeatedExprToMessageRepeatedExpr(cmd []command.Expr) ([]*Expr, error) {
	msgRepeatedExpr := []*Expr{}
	for i := range cmd {
		expr, err := ConvertCommandExprToMessageExpr(cmd[i])
		if err != nil {
			return nil, err
		}
		msgRepeatedExpr = append(msgRepeatedExpr, expr)
	}
	return msgRepeatedExpr, nil
}

// ConvertCommandFunctionalExprToMessageFunctionalExpr converts a command.Expr to a message.Expr_Func.
func ConvertCommandFunctionalExprToMessageFunctionalExpr(cmd command.FunctionExpr) (*Expr_Func, error) {
	args, err := ConvertCommandRepeatedExprToMessageRepeatedExpr(cmd.Args)
	if err != nil {
		return nil, err
	}
	return &Expr_Func{
		&FunctionExpr{
			Name:     cmd.Name,
			Distinct: cmd.Distinct,
			Args:     args,
		},
	}, nil
}

// ConvertCommandEqualityExprToMessageEqualityExpr converts a command.Expr to a message.Expr_Equality.
func ConvertCommandEqualityExprToMessageEqualityExpr(cmd command.EqualityExpr) (*Expr_Equality, error) {
	left, err := ConvertCommandExprToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandExprToMessageExpr(cmd.Right)
	if err != nil {
		return nil, err
	}
	return &Expr_Equality{
		&EqualityExpr{
			Left:   left,
			Right:  right,
			Invert: cmd.Invert,
		},
	}, nil
}

// ConvertCommandRangeExprToMessageRangeExpr converts a command.Expr to a message.Expr_Range.
func ConvertCommandRangeExprToMessageRangeExpr(cmd command.RangeExpr) (*Expr_Range, error) {
	needle, err := ConvertCommandExprToMessageExpr(cmd.Needle)
	if err != nil {
		return nil, err
	}
	lo, err := ConvertCommandExprToMessageExpr(cmd.Lo)
	if err != nil {
		return nil, err
	}
	hi, err := ConvertCommandExprToMessageExpr(cmd.Hi)
	if err != nil {
		return nil, err
	}
	return &Expr_Range{
		&RangeExpr{
			Needle: needle,
			Lo:     lo,
			Hi:     hi,
			Invert: cmd.Invert,
		},
	}, nil
}

// ConvertCommandExprToMessageExpr converts command.Expr to a message.Expr.
func ConvertCommandExprToMessageExpr(cmd command.Expr) (*Expr, error) {
	var err error
	msgExpr := &Expr{}
	switch c := cmd.(type) {
	case command.LiteralExpr:
		msgExpr.Expr, err = ConvertCommandLiteralExprToMessageLiteralExpr(c)
		if err != nil {
			return nil, err
		}
	case command.ConstantBooleanExpr:
		msgExpr.Expr, err = ConvertCommandConstantBooleanExprToMessageConstantBooleanExpr(c)
		if err != nil {
			return nil, err
		}
	case command.UnaryExpr:
		msgExpr.Expr, err = ConvertCommandUnaryExprToMessageUnaryExpr(c)
		if err != nil {
			return nil, err
		}
	case command.BinaryExpr:
		msgExpr.Expr, err = ConvertCommandBinaryExprToMessageBinaryExpr(c)
		if err != nil {
			return nil, err
		}
	case command.FunctionExpr:
		msgExpr.Expr, err = ConvertCommandFunctionalExprToMessageFunctionalExpr(c)
		if err != nil {
			return nil, err
		}
	case command.EqualityExpr:
		msgExpr.Expr, err = ConvertCommandEqualityExprToMessageEqualityExpr(c)
		if err != nil {
			return nil, err
		}
	case command.RangeExpr:
		msgExpr.Expr, err = ConvertCommandRangeExprToMessageRangeExpr(c)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnknownCommandKind
	}
	return msgExpr, nil
}

// ConvertCommandListScanToMessageListScan converts a command.Scan to a message.List_Scan.
func ConvertCommandListScanToMessageListScan(cmd command.Scan) (*List_Scan, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	return &List_Scan{
		&Command_Scan{
			Table: table,
		},
	}, nil
}

// ConvertCommandListSelectToMessageListSelect converts a command.Select to a message.List_Select.
func ConvertCommandListSelectToMessageListSelect(cmd command.Select) (*List_Select, error) {
	filter, err := ConvertCommandExprToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &List_Select{
		&Command_Select{
			Filter: filter,
			Input:  input,
		},
	}, nil
}

// ConvertCommandListProjectToMessageListProject converts a command.Project to a message.List_Project.
func ConvertCommandListProjectToMessageListProject(cmd command.Project) (*List_Project, error) {
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	cols, err := ConvertCommandColSliceToMessageColSlice(cmd.Cols)
	if err != nil {
		return nil, err
	}
	return &List_Project{
		&Command_Project{
			Cols:  cols,
			Input: input,
		},
	}, nil
}

// ConvertCommandListJoinToMessageListJoin converts a command.Join to a message.List_Join.
func ConvertCommandListJoinToMessageListJoin(cmd command.Join) (*List_Join, error) {
	filter, err := ConvertCommandExprToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	left, err := ConvertCommandListToMessageList(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandListToMessageList(cmd.Right)
	if err != nil {
		return nil, err
	}
	return &List_Join{
		&Command_Join{
			Natural: cmd.Natural,
			Type:    ConvertCommandJoinTypeToMessageJoinType(cmd.Type),
			Filter:  filter,
			Left:    left,
			Right:   right,
		},
	}, nil
}

// ConvertCommandListLimitToMessageListLimit converts a command.Limit to a message.List_Limit.
func ConvertCommandListLimitToMessageListLimit(cmd command.Limit) (*List_Limit, error) {
	limit, err := ConvertCommandExprToMessageExpr(cmd.Limit)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &List_Limit{
		&Command_Limit{
			Limit: limit,
			Input: input,
		},
	}, nil
}

// ConvertCommandListOffsetToMessageListOffset converts a command.Offset to a message.List_Offset.
func ConvertCommandListOffsetToMessageListOffset(cmd command.Offset) (*List_Offset, error) {
	offset, err := ConvertCommandExprToMessageExpr(cmd.Offset)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &List_Offset{
		&Command_Offset{
			Offset: offset,
			Input:  input,
		},
	}, nil
}

// ConvertCommandListDistinctToMessageListDistinct converts a command.Distinct to a message.List_Distinct.
func ConvertCommandListDistinctToMessageListDistinct(cmd command.Distinct) (*List_Distinct, error) {
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &List_Distinct{
		&Command_Distinct{
			Input: input,
		},
	}, nil
}

// ConvertCommandRepeatedExprToMessageRepeatedExprSlice converts a [][]command.Expr to a [][]message.Expr.
func ConvertCommandRepeatedExprToMessageRepeatedExprSlice(cmd [][]command.Expr) ([]*RepeatedExpr, error) {
	msgRepeatedExprSlice := []*RepeatedExpr{}
	for i := range cmd {
		msgRepeatedExpr := &RepeatedExpr{}
		for j := range cmd[i] {
			expr, err := ConvertCommandExprToMessageExpr(cmd[i][j])
			if err != nil {
				return nil, err
			}
			msgRepeatedExpr.Expr = append(msgRepeatedExpr.Expr, expr)
		}
		msgRepeatedExprSlice = append(msgRepeatedExprSlice, msgRepeatedExpr)
	}
	return msgRepeatedExprSlice, nil
}

// ConvertCommandListValuesToMessageListValues converts a command.Values to a message.List_Values.
func ConvertCommandListValuesToMessageListValues(cmd command.Values) (*List_Values, error) {
	exprSlice, err := ConvertCommandRepeatedExprToMessageRepeatedExprSlice(cmd.Values)
	if err != nil {
		return nil, err
	}
	return &List_Values{
		&Command_Values{
			Expr: exprSlice,
		},
	}, nil
}

// ConvertCommandListToMessageList converts a command.List to a message.List.
func ConvertCommandListToMessageList(cmd command.List) (*List, error) {
	var err error
	msgList := &List{}
	switch c := cmd.(type) {
	case command.Scan:
		msgList.List, err = ConvertCommandListScanToMessageListScan(c)
		if err != nil {
			return nil, err
		}
	case command.Select:
		msgList.List, err = ConvertCommandListSelectToMessageListSelect(c)
		if err != nil {
			return nil, err
		}
	case command.Project:
		msgList.List, err = ConvertCommandListProjectToMessageListProject(c)
		if err != nil {
			return nil, err
		}
	case command.Join:
		msgList.List, err = ConvertCommandListJoinToMessageListJoin(c)
		if err != nil {
			return nil, err
		}
	case command.Limit:
		msgList.List, err = ConvertCommandListLimitToMessageListLimit(c)
		if err != nil {
			return nil, err
		}
	case command.Offset:
		msgList.List, err = ConvertCommandListOffsetToMessageListOffset(c)
		if err != nil {
			return nil, err
		}
	case command.Distinct:
		msgList.List, err = ConvertCommandListDistinctToMessageListDistinct(c)
		if err != nil {
			return nil, err
		}
	case command.Values:
		msgList.List, err = ConvertCommandListValuesToMessageListValues(c)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnknownCommandKind
	}
	return msgList, nil
}

// ConvertCommandSelectToMessageSelect converts a Command type to a Command_Select type.
func ConvertCommandSelectToMessageSelect(cmd command.Select) (*Command_Select, error) {
	filter, err := ConvertCommandExprToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Select{
		Filter: filter,
		Input:  input,
	}, nil
}

// ConvertCommandColToMessageCol converts command.Column to a message.Column.
func ConvertCommandColToMessageCol(cmd command.Column) (*Column, error) {
	column, err := ConvertCommandExprToMessageExpr(cmd.Column)
	if err != nil {
		return nil, err
	}
	return &Column{
		Table:  cmd.Table,
		Column: column,
		Alias:  cmd.Alias,
	}, nil
}

// ConvertCommandColSliceToMessageColSlice converts []command.Column to a []message.Column.
func ConvertCommandColSliceToMessageColSlice(cmd []command.Column) ([]*Column, error) {
	msgCols := []*Column{}
	for i := range cmd {
		col, err := ConvertCommandColToMessageCol(cmd[i])
		if err != nil {
			return nil, err
		}
		msgCols = append(msgCols, col)
	}
	return msgCols, nil
}

// ConvertCommandProjectToMessageProject converts a Command type to a Command_Project type.
func ConvertCommandProjectToMessageProject(cmd command.Project) (*Command_Project, error) {
	cols, err := ConvertCommandColSliceToMessageColSlice(cmd.Cols)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Project{
		Cols:  cols,
		Input: input,
	}, nil
}

// ConvertCommandDeleteToMessageDelete converts a Command type to a Command_Delete type.
func ConvertCommandDeleteToMessageDelete(cmd command.Delete) (*Command_Delete, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	filter, err := ConvertCommandExprToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	return &Command_Delete{
		Table:  table,
		Filter: filter,
	}, nil
}

// ConvertCommandDropToMessageDrop converts a Command type to a CommandDrop type.
func ConvertCommandDropToMessageDrop(cmd command.Command) (*CommandDrop, error) {
	if cmd == nil {
		return nil, ErrNilCommand
	}
	msgCmdDrop := &CommandDrop{}
	switch c := cmd.(type) {
	case *command.DropTable:
		msgCmdDrop.Target = DropTarget_Table
		msgCmdDrop.IfExists = c.IfExists
		msgCmdDrop.Schema = c.Schema
		msgCmdDrop.Name = c.Name
	case *command.DropView:
		msgCmdDrop.Target = DropTarget_View
		msgCmdDrop.IfExists = c.IfExists
		msgCmdDrop.Schema = c.Schema
		msgCmdDrop.Name = c.Name
	case *command.DropIndex:
		msgCmdDrop.Target = DropTarget_Index
		msgCmdDrop.IfExists = c.IfExists
		msgCmdDrop.Schema = c.Schema
		msgCmdDrop.Name = c.Name
	case *command.DropTrigger:
		msgCmdDrop.Target = DropTarget_Trigger
		msgCmdDrop.IfExists = c.IfExists
		msgCmdDrop.Schema = c.Schema
		msgCmdDrop.Name = c.Name
	}
	return msgCmdDrop, nil
}

// ConvertCommandUpdateOrToMessageUpdateOr converts a command.Update or to a message.UpdateOr.
// Returns -1 if the UpdateOr type doesn't match.
func ConvertCommandUpdateOrToMessageUpdateOr(cmd command.UpdateOr) (UpdateOr, error) {
	switch cmd {
	case command.UpdateOrUnknown:
		return UpdateOr_UpdateOrUnknown, nil
	case command.UpdateOrRollback:
		return UpdateOr_UpdateOrRollback, nil
	case command.UpdateOrAbort:
		return UpdateOr_UpdateOrAbort, nil
	case command.UpdateOrReplace:
		return UpdateOr_UpdateOrReplace, nil
	case command.UpdateOrFail:
		return UpdateOr_UpdateOrFail, nil
	case command.UpdateOrIgnore:
		return UpdateOr_UpdateOrIgnore, nil
	}
	return -1, ErrUnknownCommandKind
}

// ConvertCommandUpdateSetterLiteralToMessageUpdateSetterLiteral converts
// a command.Literal to a message.UpdateSetter_Literal.
func ConvertCommandUpdateSetterLiteralToMessageUpdateSetterLiteral(
	cmd command.LiteralExpr,
) (*UpdateSetter_Literal, error) {
	return &UpdateSetter_Literal{
		&LiteralExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandUpdateSetterConstantToMessageUpdateSetterConstant converts
// a command.Constant to a message.UpdateSetter_Constant.
func ConvertCommandUpdateSetterConstantToMessageUpdateSetterConstant(
	cmd command.ConstantBooleanExpr,
) (*UpdateSetter_Constant, error) {

	return &UpdateSetter_Constant{
		&ConstantBooleanExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandUpdateSetterUnaryToMessageUpdateSetterUnary converts
//  a command.Unary to a message.UpdateSetter_Unary.
func ConvertCommandUpdateSetterUnaryToMessageUpdateSetterUnary(
	cmd command.UnaryExpr,
) (*UpdateSetter_Unary, error) {
	val, err := ConvertCommandExprToMessageExpr(cmd.Value)
	if err != nil {
		return nil, err
	}
	return &UpdateSetter_Unary{
		&UnaryExpr{
			Operator: cmd.Operator,
			Value:    val,
		},
	}, nil
}

// ConvertCommandUpdateSetterBinaryToMessageUpdateSetterBinary converts
//  a command.Binary to a message.UpdateSetter_Binary.
func ConvertCommandUpdateSetterBinaryToMessageUpdateSetterBinary(
	cmd command.BinaryExpr,
) (*UpdateSetter_Binary, error) {
	left, err := ConvertCommandExprToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandExprToMessageExpr(cmd.Right)
	if err != nil {
		return nil, err
	}
	return &UpdateSetter_Binary{
		&BinaryExpr{
			Operator: cmd.Operator,
			Left:     left,
			Right:    right,
		},
	}, nil
}

// ConvertCommandUpdateSetterFuncToMessageUpdateSetterFunc converts
// a command.Func to a message.UpdateSetter_Func.
func ConvertCommandUpdateSetterFuncToMessageUpdateSetterFunc(
	cmd command.FunctionExpr,
) (*UpdateSetter_Func, error) {
	repExpr, err := ConvertCommandRepeatedExprToMessageRepeatedExpr(cmd.Args)
	if err != nil {
		return nil, err
	}
	return &UpdateSetter_Func{
		&FunctionExpr{
			Name:     cmd.Name,
			Distinct: cmd.Distinct,
			Args:     repExpr,
		},
	}, nil
}

// ConvertCommandUpdateSetterEqualityToMessageUpdateSetterEquality converts
// a command.Equality to a message.UpdateSetter_Equality.
func ConvertCommandUpdateSetterEqualityToMessageUpdateSetterEquality(
	cmd command.EqualityExpr,
) (*UpdateSetter_Equality, error) {
	left, err := ConvertCommandExprToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandExprToMessageExpr(cmd.Right)
	if err != nil {
		return nil, err
	}
	return &UpdateSetter_Equality{
		&EqualityExpr{
			Left:   left,
			Right:  right,
			Invert: cmd.Invert,
		},
	}, nil
}

// ConvertCommandUpdateSetterRangeToMessageUpdateSetterRange converts
// a command.Range to a message.UpdateSetter_Range.
func ConvertCommandUpdateSetterRangeToMessageUpdateSetterRange(
	cmd command.RangeExpr,
) (*UpdateSetter_Range, error) {
	needle, err := ConvertCommandExprToMessageExpr(cmd.Needle)
	if err != nil {
		return nil, err
	}
	lo, err := ConvertCommandExprToMessageExpr(cmd.Lo)
	if err != nil {
		return nil, err
	}
	hi, err := ConvertCommandExprToMessageExpr(cmd.Hi)
	if err != nil {
		return nil, err
	}
	return &UpdateSetter_Range{
		&RangeExpr{
			Needle: needle,
			Lo:     lo,
			Hi:     hi,
			Invert: cmd.Invert,
		},
	}, nil
}

// ConvertCommandUpdateSetterToMessageUpdateSetter converts
// a command.UpdateSetter to a message.UpdateSetter.
func ConvertCommandUpdateSetterToMessageUpdateSetter(
	cmd command.UpdateSetter,
) (*UpdateSetter, error) {
	var err error
	msgUpdateSetter := &UpdateSetter{}
	msgUpdateSetter.Cols = cmd.Cols
	switch val := cmd.Value.(type) {
	case command.LiteralExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterLiteralToMessageUpdateSetterLiteral(val)
		if err != nil {
			return nil, err
		}
	case command.ConstantBooleanExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterConstantToMessageUpdateSetterConstant(val)
		if err != nil {
			return nil, err
		}
	case command.UnaryExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterUnaryToMessageUpdateSetterUnary(val)
		if err != nil {
			return nil, err
		}
	case command.BinaryExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterBinaryToMessageUpdateSetterBinary(val)
		if err != nil {
			return nil, err
		}
	case command.FunctionExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterFuncToMessageUpdateSetterFunc(val)
		if err != nil {
			return nil, err
		}
	case command.EqualityExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterEqualityToMessageUpdateSetterEquality(val)
		if err != nil {
			return nil, err
		}
	case command.RangeExpr:
		msgUpdateSetter.Value, err = ConvertCommandUpdateSetterRangeToMessageUpdateSetterRange(val)
		if err != nil {
			return nil, err
		}
	}
	return msgUpdateSetter, nil
}

// ConvertCommandUpdateSetterSliceToMessageUpdateSetterSlice converts
// a []command.UpdateSetter to a []message.UpdateSetter.
func ConvertCommandUpdateSetterSliceToMessageUpdateSetterSlice(
	cmd []command.UpdateSetter,
) ([]*UpdateSetter, error) {
	msgUpdateSetterSlice := []*UpdateSetter{}
	for i := range cmd {
		updateSetter, err := ConvertCommandUpdateSetterToMessageUpdateSetter(cmd[i])
		if err != nil {
			return nil, err
		}
		msgUpdateSetterSlice = append(msgUpdateSetterSlice, updateSetter)
	}
	return msgUpdateSetterSlice, nil
}

// ConvertCommandUpdateToMessageUpdate converts a Command type to a Command_Update type.
func ConvertCommandUpdateToMessageUpdate(cmd command.Update) (*Command_Update, error) {
	updateOr, err := ConvertCommandUpdateOrToMessageUpdateOr(cmd.UpdateOr)
	if err != nil {
		return nil, err
	}
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	updates, err := ConvertCommandUpdateSetterSliceToMessageUpdateSetterSlice(cmd.Updates)
	if err != nil {
		return nil, err
	}
	filter, err := ConvertCommandExprToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	return &Command_Update{
		UpdateOr: updateOr,
		Table:    table,
		Updates:  updates,
		Filter:   filter,
	}, nil
}

// ConvertCommandJoinTypeToMessageJoinType converts command.JoinType to message.JoinType.
// It returns -1 on not finding a valid JoinType.
func ConvertCommandJoinTypeToMessageJoinType(cmd command.JoinType) JoinType {
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

// ConvertCommandJoinToMessageJoin converts a Command type to a Command_Join type.
func ConvertCommandJoinToMessageJoin(cmd command.Join) (*Command_Join, error) {
	filter, err := ConvertCommandExprToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	left, err := ConvertCommandListToMessageList(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandListToMessageList(cmd.Right)
	if err != nil {
		return nil, err
	}

	return &Command_Join{
		Natural: cmd.Natural,
		Type:    ConvertCommandJoinTypeToMessageJoinType(cmd.Type),
		Filter:  filter,
		Left:    left,
		Right:   right,
	}, nil
}

// ConvertCommandLimitToMessageLimit converts a Command type to a Command_Limit type.
func ConvertCommandLimitToMessageLimit(cmd command.Limit) (*Command_Limit, error) {
	limit, err := ConvertCommandExprToMessageExpr(cmd.Limit)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Limit{
		Limit: limit,
		Input: input,
	}, nil
}

// ConvertCommandInsertOrToMessageInsertOr converts command.InsertOr to a message.InsertOr.
// It returns -1 on not finding the right InsertOr type.
func ConvertCommandInsertOrToMessageInsertOr(cmd command.InsertOr) InsertOr {
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

// ConvertCommandInsertToMessageInsert converts a Command type to a Command_Insert type.
func ConvertCommandInsertToMessageInsert(cmd command.Insert) (*Command_Insert, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	colSlice, err := ConvertCommandColSliceToMessageColSlice(cmd.Cols)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandListToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Insert{
		InsertOr:      ConvertCommandInsertOrToMessageInsertOr(cmd.InsertOr),
		Table:         table,
		Cols:          colSlice,
		DefaultValues: cmd.DefaultValues,
		Input:         input,
	}, nil
}

// ConvertMessageToCommand converts a message.Command to a command.Command.
func ConvertMessageToCommand(msg Message) command.Command {
	switch m := msg.(type) {
	case *Command_Scan:
		return ConvertMessageScanToCommandScan(m)
	case *Command_Select:
		return ConvertMessageSelectToCommandSelect(m)
	case *Command_Project:
		return ConvertMessageProjectToCommandProject(m)
	case *Command_Delete:
		return ConvertMessageDeleteToCommandDelete(m)
	case *CommandDrop:
		switch m.Target {
		case 0:
			return ConvertMessageDropToCommandDropTable(m)
		case 1:
			return ConvertMessageDropToCommandDropView(m)
		case 2:
			return ConvertMessageDropToCommandDropIndex(m)
		case 3:
			return ConvertMessageDropToCommandDropTrigger(m)
		}
	case *Command_Update:
		return ConvertMessageUpdateToCommandUpdate(m)
	case *Command_Join:
		return ConvertMessageJoinToCommandJoin(m)
	case *Command_Limit:
		return ConvertMessageLimitToCommandLimit(m)
	case *Command_Insert:
		return ConvertMessageInsertToCommandInsert(m)
	}
	return nil
}

// ConvertMessageTableToCommandTable converts a message.SimpleTable to a command.Table.
func ConvertMessageTableToCommandTable(msg *SimpleTable) command.Table {
	return &command.SimpleTable{
		Schema:  msg.Schema,
		Table:   msg.Table,
		Alias:   msg.Alias,
		Indexed: msg.Indexed,
		Index:   msg.Index,
	}
}

// ConvertMessageScanToCommandScan converts a message.Command_Scan to a command.Scan.
func ConvertMessageScanToCommandScan(msg *Command_Scan) command.Scan {
	return command.Scan{
		Table: ConvertMessageTableToCommandTable(msg.Table),
	}
}

// ConvertMessageLiteralExprToCommandLiteralExpr converts a message.Expr to a command.LiteralExpr.
func ConvertMessageLiteralExprToCommandLiteralExpr(msg *Expr) command.LiteralExpr {
	return command.LiteralExpr{
		Value: msg.GetLiteral().GetValue(),
	}
}

// ConvertMessageBooleanExprToCommandConstantBooleanExpr converts a message.Expr to a command.ConstantBooleanExpr.
func ConvertMessageBooleanExprToCommandConstantBooleanExpr(msg *Expr) command.ConstantBooleanExpr {
	return command.ConstantBooleanExpr{
		Value: msg.GetConstant().GetValue(),
	}
}

// ConvertMessageUnaryExprToCommandUnaryExpr converts a message.Expr to a command.UnaryExpr.
func ConvertMessageUnaryExprToCommandUnaryExpr(msg *Expr) command.UnaryExpr {
	return command.UnaryExpr{
		Operator: msg.GetUnary().GetOperator(),
		Value:    ConvertMessageExprToCommandExpr(msg.GetUnary().GetValue()),
	}
}

// ConvertMessageBinaryExprToCommandBinaryExpr converts a message.Expr to a command.BinaryExpr.
func ConvertMessageBinaryExprToCommandBinaryExpr(msg *Expr) command.BinaryExpr {
	return command.BinaryExpr{
		Operator: msg.GetBinary().GetOperator(),
		Left:     ConvertMessageExprToCommandExpr(msg.GetBinary().GetLeft()),
		Right:    ConvertMessageExprToCommandExpr(msg.GetBinary().GetRight()),
	}
}

// ConvertMessageExprSliceToCommandExprSlice converts a []*message.Expr to []command.Expr.
func ConvertMessageExprSliceToCommandExprSlice(msg []*Expr) []command.Expr {
	msgExprSlice := []command.Expr{}
	for i := range msg {
		msgExprSlice = append(msgExprSlice, ConvertMessageExprToCommandExpr(msg[i]))
	}
	return msgExprSlice
}

// ConvertMessageFunctionalExprToCommandFunctionExpr converts a message.Expr to a command.FunctionExpr.
func ConvertMessageFunctionalExprToCommandFunctionExpr(msg *Expr) command.FunctionExpr {
	return command.FunctionExpr{
		Name:     msg.GetFunc().GetName(),
		Distinct: msg.GetFunc().GetDistinct(),
		Args:     ConvertMessageExprSliceToCommandExprSlice(msg.GetFunc().GetArgs()),
	}
}

// ConvertMessageEqualityExprToCommandEqualityExpr converts a message.Expr to a command.EqualityExpr.
func ConvertMessageEqualityExprToCommandEqualityExpr(msg *Expr) command.EqualityExpr {
	return command.EqualityExpr{
		Left:   ConvertMessageExprToCommandExpr(msg.GetEquality().GetLeft()),
		Right:  ConvertMessageExprToCommandExpr(msg.GetEquality().GetRight()),
		Invert: msg.GetEquality().Invert,
	}
}

// ConvertMessageRangeExprToCommandRangeExpr converts a message.Expr to a command.RangeExpr.
func ConvertMessageRangeExprToCommandRangeExpr(msg *Expr) command.RangeExpr {
	return command.RangeExpr{
		Needle: ConvertMessageExprToCommandExpr(msg.GetRange().GetNeedle()),
		Lo:     ConvertMessageExprToCommandExpr(msg.GetRange().GetLo()),
		Hi:     ConvertMessageExprToCommandExpr(msg.GetRange().GetHi()),
	}
}

// ConvertMessageExprToCommandExpr converts a message.Expr to a command.Expr.
func ConvertMessageExprToCommandExpr(msg *Expr) command.Expr {
	if msg == nil {
		return nil
	}
	switch msg.Expr.(type) {
	case *Expr_Literal:
		return ConvertMessageLiteralExprToCommandLiteralExpr(msg)
	case *Expr_Constant:
		return ConvertMessageBooleanExprToCommandConstantBooleanExpr(msg)
	case *Expr_Unary:
		return ConvertMessageUnaryExprToCommandUnaryExpr(msg)
	case *Expr_Binary:
		return ConvertMessageBinaryExprToCommandBinaryExpr(msg)
	case *Expr_Func:
		return ConvertMessageFunctionalExprToCommandFunctionExpr(msg)
	case *Expr_Equality:
		return ConvertMessageEqualityExprToCommandEqualityExpr(msg)
	case *Expr_Range:
		return ConvertMessageRangeExprToCommandRangeExpr(msg)
	}
	return nil
}

// ConvertMessageListScanToCommandListScan converts a message.List to a command.Scan.
func ConvertMessageListScanToCommandListScan(msg *List) command.Scan {
	return command.Scan{
		Table: ConvertMessageTableToCommandTable(msg.GetScan().GetTable()),
	}
}

// ConvertMessageListSelectToCommandListSelect converts a message.List to a command.Select.
func ConvertMessageListSelectToCommandListSelect(msg *List) command.Select {
	return command.Select{
		Filter: ConvertMessageExprToCommandExpr(msg.GetSelect().GetFilter()),
		Input:  ConvertMessageListToCommandList(msg.GetSelect().GetInput()),
	}
}

// ConvertMessageListProjectToCommandListProject converts a message.List to a command.Project.
func ConvertMessageListProjectToCommandListProject(msg *List) command.Project {
	return command.Project{
		Cols:  ConvertMessageColsToCommandCols(msg.GetProject().GetCols()),
		Input: ConvertMessageListToCommandList(msg.GetProject().GetInput()),
	}
}

// ConvertMessageListJoinToCommandListJoin converts a message.List to a command.Join.
func ConvertMessageListJoinToCommandListJoin(msg *List) command.Join {
	return command.Join{
		Natural: msg.GetJoin().GetNatural(),
		Type:    ConvertMessageJoinTypeToCommandJoinType(msg.GetJoin().GetType()),
		Filter:  ConvertMessageExprToCommandExpr(msg.GetJoin().GetFilter()),
		Left:    ConvertMessageListToCommandList(msg.GetJoin().GetLeft()),
		Right:   ConvertMessageListToCommandList(msg.GetJoin().GetRight()),
	}
}

// ConvertMessageListLimitToCommandListLimit converts a message.List to a command.Limit.
func ConvertMessageListLimitToCommandListLimit(msg *List) command.Limit {
	return command.Limit{
		Limit: ConvertMessageExprToCommandExpr(msg.GetLimit().GetLimit()),
		Input: ConvertMessageListToCommandList(msg.GetLimit().GetInput()),
	}
}

// ConvertMessageListOffsetToCommandListOffset converts a message.List to a command.Offset.
func ConvertMessageListOffsetToCommandListOffset(msg *List) command.Offset {
	return command.Offset{
		Offset: ConvertMessageExprToCommandExpr(msg.GetOffset().GetOffset()),
		Input:  ConvertMessageListToCommandList(msg.GetDistinct().GetInput()),
	}
}

// ConvertMessageListDistinctToCommandListDistinct converts a message.List to a command.Distinct.
func ConvertMessageListDistinctToCommandListDistinct(msg *List) command.Distinct {
	return command.Distinct{
		Input: ConvertMessageListToCommandList(msg.GetDistinct().GetInput()),
	}
}

// ConvertMessageExprToCommandExprRepeatedSlice converts a message.RepeatedExpr to a [][]command.Expr.
func ConvertMessageExprToCommandExprRepeatedSlice(msg []*RepeatedExpr) [][]command.Expr {
	cmdRepeatedExprSlice := [][]command.Expr{}
	for i := range msg {
		cmdRepeatedExpr := []command.Expr{}
		for j := range msg[i].Expr {
			cmdRepeatedExpr = append(cmdRepeatedExpr, ConvertMessageExprToCommandExpr(msg[i].Expr[j]))
		}
		cmdRepeatedExprSlice = append(cmdRepeatedExprSlice, cmdRepeatedExpr)
	}
	return cmdRepeatedExprSlice
}

// ConvertMessageListValuesToCommandListValues converts a message.List to a command.Values.
func ConvertMessageListValuesToCommandListValues(msg *List) command.Values {
	return command.Values{
		Values: ConvertMessageExprToCommandExprRepeatedSlice(msg.GetValues().GetExpr()),
	}
}

// ConvertMessageListToCommandList converts a message.List to a command.List.
func ConvertMessageListToCommandList(msg *List) command.List {
	if msg == nil {
		return nil
	}
	switch msg.List.(type) {
	case *List_Scan:
		return ConvertMessageListScanToCommandListScan(msg)
	case *List_Select:
		return ConvertMessageListSelectToCommandListSelect(msg)
	case *List_Project:
		return ConvertMessageListProjectToCommandListProject(msg)
	case *List_Join:
		return ConvertMessageListJoinToCommandListJoin(msg)
	case *List_Limit:
		return ConvertMessageListLimitToCommandListLimit(msg)
	case *List_Offset:
		return ConvertMessageListOffsetToCommandListOffset(msg)
	case *List_Distinct:
		return ConvertMessageListDistinctToCommandListDistinct(msg)
	case *List_Values:
		return ConvertMessageListValuesToCommandListValues(msg)
	}
	return nil
}

// ConvertMessageSelectToCommandSelect converts a message.Command_Select to a command.Select.
func ConvertMessageSelectToCommandSelect(msg *Command_Select) command.Select {
	return command.Select{
		Filter: ConvertMessageExprToCommandExpr(msg.GetFilter()),
		Input:  ConvertMessageListToCommandList(msg.GetInput()),
	}
}

// ConvertMessageColToCommandCol converts a message.Column to a command.Column.
func ConvertMessageColToCommandCol(msg *Column) command.Column {
	return command.Column{
		Table:  msg.GetTable(),
		Column: ConvertMessageExprToCommandExpr(msg.GetColumn()),
		Alias:  msg.GetAlias(),
	}
}

// ConvertMessageColsToCommandCols converts a []message.Column to a []command.Column.
func ConvertMessageColsToCommandCols(msg []*Column) []command.Column {
	cmdCols := []command.Column{}
	for i := range msg {
		cmdCols = append(cmdCols, ConvertMessageColToCommandCol(msg[i]))
	}
	return cmdCols
}

// ConvertMessageProjectToCommandProject converts a message.Command_Project to a command.Project.
func ConvertMessageProjectToCommandProject(msg *Command_Project) command.Project {
	return command.Project{
		Cols:  ConvertMessageColsToCommandCols(msg.GetCols()),
		Input: ConvertMessageListToCommandList(msg.GetInput()),
	}
}

// ConvertMessageDeleteToCommandDelete converts a message.Command_Delete to a command.Delete.
func ConvertMessageDeleteToCommandDelete(msg *Command_Delete) command.Delete {
	return command.Delete{
		Filter: ConvertMessageExprToCommandExpr(msg.GetFilter()),
		Table:  ConvertMessageTableToCommandTable(msg.GetTable()),
	}
}

// ConvertMessageDropToCommandDropTable converts a message.CommandDrop to a command.Drop.
func ConvertMessageDropToCommandDropTable(msg *CommandDrop) command.DropTable {
	return command.DropTable{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageDropToCommandDropView converts a message.CommandDrop to a command.Drop.
func ConvertMessageDropToCommandDropView(msg *CommandDrop) command.DropView {
	return command.DropView{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageDropToCommandDropIndex converts a message.CommandDrop to a command.Drop.
func ConvertMessageDropToCommandDropIndex(msg *CommandDrop) command.DropIndex {
	return command.DropIndex{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageDropToCommandDropTrigger converts a message.CommandDrop to a command.Drop.
func ConvertMessageDropToCommandDropTrigger(msg *CommandDrop) command.DropTrigger {
	return command.DropTrigger{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageUpdateOrToCommandUpdateOr converts a message.UpdateOr to command.UpdateOr.
func ConvertMessageUpdateOrToCommandUpdateOr(msg UpdateOr) command.UpdateOr {
	return command.UpdateOr(msg.Number())
}

// ConvertMessageUpdateSetterLiteralExprToCommandUpdateSetterLiteralExpr converts
// a message.LiteralExpr to command.Expr.
func ConvertMessageUpdateSetterLiteralExprToCommandUpdateSetterLiteralExpr(
	msg *LiteralExpr,
) command.Expr {
	return command.LiteralExpr{
		Value: msg.Value,
	}
}

// ConvertMessageUpdateSetterConstantExprToCommandUpdateSetterConstantExpr converts
// a message.ConstantBooleanExpr to a command.Expr.
func ConvertMessageUpdateSetterConstantExprToCommandUpdateSetterConstantExpr(
	msg *ConstantBooleanExpr,
) command.Expr {
	return command.ConstantBooleanExpr{
		Value: msg.Value,
	}
}

// ConvertMessageUpdateSetterUnaryExprToCommandUpdateSetterUnaryExpr converts
// a message.UnaryExpr to command.Expr.
func ConvertMessageUpdateSetterUnaryExprToCommandUpdateSetterUnaryExpr(
	msg *UnaryExpr,
) command.Expr {
	return command.UnaryExpr{
		Operator: msg.Operator,
		Value:    ConvertMessageBinaryExprToCommandBinaryExpr(msg.Value),
	}
}

// ConvertMessageUpdateSetterBinaryExprToCommandUpdateSetterBinaryExpr converts
// a message.BinaryExpr to command.Expr.
func ConvertMessageUpdateSetterBinaryExprToCommandUpdateSetterBinaryExpr(
	msg *BinaryExpr,
) command.Expr {
	return command.BinaryExpr{
		Operator: msg.Operator,
		Left:     ConvertMessageBinaryExprToCommandBinaryExpr(msg.Left),
		Right:    ConvertMessageBinaryExprToCommandBinaryExpr(msg.Right),
	}
}

// ConvertMessageUpdateSetterFuncExprToCommandUpdateSetterFuncExpr converts
// a message.FunctionExpr tp command.Expr.
func ConvertMessageUpdateSetterFuncExprToCommandUpdateSetterFuncExpr(
	msg *FunctionExpr,
) command.Expr {
	return command.FunctionExpr{
		Name:     msg.Name,
		Distinct: msg.Distinct,
		Args:     ConvertMessageExprSliceToCommandExprSlice(msg.Args),
	}
}

// ConvertMessageUpdateEqualityExprToCommandUpdateSetterEqualityExpr converts
// a message.EqualityExpr to a command.Expr.
func ConvertMessageUpdateEqualityExprToCommandUpdateSetterEqualityExpr(
	msg *EqualityExpr,
) command.Expr {
	return command.EqualityExpr{
		Left:   ConvertMessageBinaryExprToCommandBinaryExpr(msg.Left),
		Right:  ConvertMessageBinaryExprToCommandBinaryExpr(msg.Right),
		Invert: msg.Invert,
	}
}

// ConvertMessageUpdateSetterRangeExprToCommandUpdateSetterRangeExpr converts
// a message.RangeExpr to a command.Expr.
func ConvertMessageUpdateSetterRangeExprToCommandUpdateSetterRangeExpr(
	msg *RangeExpr,
) command.Expr {
	return command.RangeExpr{
		Needle: ConvertMessageBinaryExprToCommandBinaryExpr(msg.Needle),
		Lo:     ConvertMessageBinaryExprToCommandBinaryExpr(msg.Lo),
		Hi:     ConvertMessageBinaryExprToCommandBinaryExpr(msg.Hi),
		Invert: msg.Invert,
	}
}

// ConvertMessageUpdateSetterToCommandUpdateSetter converts
// a message.UpdateSetter to a command.UpdateSetter.
func ConvertMessageUpdateSetterToCommandUpdateSetter(msg *UpdateSetter) command.UpdateSetter {
	cmdUpdateSetter := command.UpdateSetter{}
	cmdUpdateSetter.Cols = msg.Cols
	switch msg.Value.(type) {
	case *UpdateSetter_Literal:
		cmdUpdateSetter.Value = ConvertMessageUpdateSetterLiteralExprToCommandUpdateSetterLiteralExpr(msg.GetLiteral())
	case *UpdateSetter_Constant:
		cmdUpdateSetter.Value = ConvertMessageUpdateSetterConstantExprToCommandUpdateSetterConstantExpr(msg.GetConstant())
	case *UpdateSetter_Unary:
		cmdUpdateSetter.Value = ConvertMessageUpdateSetterUnaryExprToCommandUpdateSetterUnaryExpr(msg.GetUnary())
	case *UpdateSetter_Binary:
		cmdUpdateSetter.Value = ConvertMessageUpdateSetterBinaryExprToCommandUpdateSetterBinaryExpr(msg.GetBinary())
	case *UpdateSetter_Func:
		cmdUpdateSetter.Value = ConvertMessageUpdateSetterFuncExprToCommandUpdateSetterFuncExpr(msg.GetFunc())
	case *UpdateSetter_Equality:
		cmdUpdateSetter.Value = ConvertMessageUpdateEqualityExprToCommandUpdateSetterEqualityExpr(msg.GetEquality())
	case *UpdateSetter_Range:
		cmdUpdateSetter.Value = ConvertMessageUpdateSetterRangeExprToCommandUpdateSetterRangeExpr(msg.GetRange())
	}
	return cmdUpdateSetter
}

// ConvertMessageUpdateSetterSliceToCommandUpdateSetterSlice converts
// a []message.UpdateSetter to a []command.UpdateSetter.
func ConvertMessageUpdateSetterSliceToCommandUpdateSetterSlice(
	msg []*UpdateSetter,
) []command.UpdateSetter {
	cmdUpdateSetterSlice := []command.UpdateSetter{}
	for i := range msg {
		cmdUpdateSetterSlice = append(cmdUpdateSetterSlice, ConvertMessageUpdateSetterToCommandUpdateSetter(msg[i]))
	}
	return cmdUpdateSetterSlice
}

// ConvertMessageUpdateToCommandUpdate converts a message.Command_Update to a command.Update.
func ConvertMessageUpdateToCommandUpdate(msg *Command_Update) command.Update {
	return command.Update{
		UpdateOr: ConvertMessageUpdateOrToCommandUpdateOr(msg.GetUpdateOr()),
		Updates:  ConvertMessageUpdateSetterSliceToCommandUpdateSetterSlice(msg.GetUpdates()),
		Table:    ConvertMessageTableToCommandTable(msg.GetTable()),
		Filter:   ConvertMessageExprToCommandExpr(msg.GetFilter()),
	}
}

// ConvertMessageJoinTypeToCommandJoinType converts a message.JoinType to a command.JoinType.
func ConvertMessageJoinTypeToCommandJoinType(msg JoinType) command.JoinType {
	return command.JoinType(msg.Number())
}

// ConvertMessageJoinToCommandJoin converts a message.Command_Join to a command.Join.
func ConvertMessageJoinToCommandJoin(msg *Command_Join) command.Join {
	return command.Join{
		Natural: msg.Natural,
		Type:    ConvertMessageJoinTypeToCommandJoinType(msg.GetType()),
		Filter:  ConvertMessageExprToCommandExpr(msg.GetFilter()),
		Left:    ConvertMessageListToCommandList(msg.GetLeft()),
		Right:   ConvertMessageListToCommandList(msg.GetRight()),
	}
}

// ConvertMessageLimitToCommandLimit converts a message.Command_Limit to a command.Limit.
func ConvertMessageLimitToCommandLimit(msg *Command_Limit) command.Limit {
	return command.Limit{
		Limit: ConvertMessageExprToCommandExpr(msg.GetLimit()),
		Input: ConvertMessageListToCommandList(msg.GetInput()),
	}
}

// ConvertMessageInsertOrToCommandInsertOr converts a message.InsertOr to command.InsertOr.
func ConvertMessageInsertOrToCommandInsertOr(msg InsertOr) command.InsertOr {
	return command.InsertOr(msg.Number())
}

// ConvertMessageInsertToCommandInsert converts a message.Command_Insert to a command.Insert
func ConvertMessageInsertToCommandInsert(msg *Command_Insert) command.Insert {
	return command.Insert{
		InsertOr:      ConvertMessageInsertOrToCommandInsertOr(msg.GetInsertOr()),
		Table:         ConvertMessageTableToCommandTable(msg.GetTable()),
		Cols:          ConvertMessageColsToCommandCols(msg.GetCols()),
		DefaultValues: msg.GetDefaultValues(),
		Input:         ConvertMessageListToCommandList(msg.GetInput()),
	}
}
