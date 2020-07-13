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
	case *command.Scan:
		return ConvertCommandScanToMessageScan(c)
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
func ConvertCommandScanToMessageScan(cmd *command.Scan) (*Command_Scan, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	return &Command_Scan{
		Table: table,
	}, nil
}

// ConvertCommandLiteralExprToMessageLiteralExpr converts a command.Expr to a message.Expr_Literal.
func ConvertCommandLiteralExprToMessageLiteralExpr(cmd *command.LiteralExpr) (*Expr_Literal, error) {
	return &Expr_Literal{
		&LiteralExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandToMessageConstantBooleanExpr converts a command.Expr to a message.Expr_Constant.
func ConvertCommandToMessageConstantBooleanExpr(cmd *command.ConstantBooleanExpr) (*Expr_Constant, error) {
	return &Expr_Constant{
		&ConstantBooleanExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandToMessageUnaryExpr converts a command.Expr to a message.Expr_Unary.
func ConvertCommandToMessageUnaryExpr(cmd *command.UnaryExpr) (*Expr_Unary, error) {
	val, err := ConvertCommandToMessageExpr(cmd.Value)
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

// ConvertCommandToMessageBinaryExpr converts a command.Expr to a message.Expr_Binary.
func ConvertCommandToMessageBinaryExpr(cmd *command.BinaryExpr) (*Expr_Binary, error) {
	left, err := ConvertCommandToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandToMessageExpr(cmd.Right)
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

// ConvertCommandToMessageRepeatedExpr converts a []command.Expr to a message.Expr.
func ConvertCommandToMessageRepeatedExpr(cmd []command.Expr) ([]*Expr, error) {
	msgRepeatedExpr := []*Expr{}
	for i := range cmd {
		expr, err := ConvertCommandToMessageExpr(cmd[i])
		if err != nil {
			return nil, err
		}
		msgRepeatedExpr = append(msgRepeatedExpr, expr)
	}
	return msgRepeatedExpr, nil
}

// ConvertCommandToMessageFunctionalExpr converts a command.Expr to a message.Expr_Func.
func ConvertCommandToMessageFunctionalExpr(cmd *command.FunctionExpr) (*Expr_Func, error) {
	args, err := ConvertCommandToMessageRepeatedExpr(cmd.Args)
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

// ConvertCommandToMessageEqualityExpr converts a command.Expr to a message.Expr_Equality.
func ConvertCommandToMessageEqualityExpr(cmd *command.EqualityExpr) (*Expr_Equality, error) {
	left, err := ConvertCommandToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandToMessageExpr(cmd.Right)
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

// ConvertCommandToMessageRangeExpr converts a command.Expr to a message.Expr_Range.
func ConvertCommandToMessageRangeExpr(cmd *command.RangeExpr) (*Expr_Range, error) {
	needle, err := ConvertCommandToMessageExpr(cmd.Needle)
	if err != nil {
		return nil, err
	}
	lo, err := ConvertCommandToMessageExpr(cmd.Lo)
	if err != nil {
		return nil, err
	}
	hi, err := ConvertCommandToMessageExpr(cmd.Hi)
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

// ConvertCommandToMessageExpr converts command.Expr to a message.Expr.
func ConvertCommandToMessageExpr(cmd command.Expr) (*Expr, error) {
	var err error
	msgExpr := &Expr{}
	switch c := cmd.(type) {
	case *command.LiteralExpr:
		msgExpr.Expr, err = ConvertCommandLiteralExprToMessageLiteralExpr(c)
		if err != nil {
			return nil, err
		}
	case *command.ConstantBooleanExpr:
		msgExpr.Expr, err = ConvertCommandToMessageConstantBooleanExpr(c)
		if err != nil {
			return nil, err
		}
	case *command.UnaryExpr:
		msgExpr.Expr, err = ConvertCommandToMessageUnaryExpr(c)
		if err != nil {
			return nil, err
		}
	case *command.BinaryExpr:
		msgExpr.Expr, err = ConvertCommandToMessageBinaryExpr(c)
		if err != nil {
			return nil, err
		}
	case *command.FunctionExpr:
		msgExpr.Expr, err = ConvertCommandToMessageFunctionalExpr(c)
		if err != nil {
			return nil, err
		}
	case *command.EqualityExpr:
		msgExpr.Expr, err = ConvertCommandToMessageEqualityExpr(c)
		if err != nil {
			return nil, err
		}
	case *command.RangeExpr:
		msgExpr.Expr, err = ConvertCommandToMessageRangeExpr(c)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnknownCommandKind
	}
	return msgExpr, nil
}

// ConvertCommandToMessageListScan converts a command.Scan to a message.List_Scan.
func ConvertCommandToMessageListScan(cmd *command.Scan) (*List_Scan, error) {
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

// ConvertCommandToMessageListSelect converts a command.Select to a message.List_Select.
func ConvertCommandToMessageListSelect(cmd *command.Select) (*List_Select, error) {
	filter, err := ConvertCommandToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.Input)
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

// ConvertCommandToMessageListProject converts a command.Project to a message.List_Project.
func ConvertCommandToMessageListProject(cmd *command.Project) (*List_Project, error) {
	input, err := ConvertCommandToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	cols, err := ConvertCommandToMessageColSlice(cmd.Cols)
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

// ConvertCommandToMessageListJoin converts a command.Join to a message.List_Join.
func ConvertCommandToMessageListJoin(cmd *command.Join) (*List_Join, error) {
	filter, err := ConvertCommandToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	left, err := ConvertCommandToMessageList(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandToMessageList(cmd.Right)
	if err != nil {
		return nil, err
	}
	return &List_Join{
		&Command_Join{
			Natural: cmd.Natural,
			Type:    ConvertCommandToMessageJoinType(cmd.Type),
			Filter:  filter,
			Left:    left,
			Right:   right,
		},
	}, nil
}

// ConvertCommandToMessageListLimit converts a command.Limit to a message.List_Limit.
func ConvertCommandToMessageListLimit(cmd *command.Limit) (*List_Limit, error) {
	limit, err := ConvertCommandToMessageExpr(cmd.Limit)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.Input)
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

// ConvertCommandToMessageListOffset converts a command.Offset to a message.List_Offset.
func ConvertCommandToMessageListOffset(cmd *command.Offset) (*List_Offset, error) {
	offset, err := ConvertCommandToMessageExpr(cmd.Offset)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.Input)
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

// ConvertCommandToMessageListDistinct converts a command.Distinct to a message.List_Distinct.
func ConvertCommandToMessageListDistinct(cmd *command.Distinct) (*List_Distinct, error) {
	input, err := ConvertCommandToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &List_Distinct{
		&Command_Distinct{
			Input: input,
		},
	}, nil
}

// ConvertCommandToMessageRepeatedExprSlice converts a [][]command.Expr to a [][]message.Expr.
func ConvertCommandToMessageRepeatedExprSlice(cmd [][]command.Expr) ([]*RepeatedExpr, error) {
	msgRepeatedExprSlice := []*RepeatedExpr{}
	for i := range cmd {
		msgRepeatedExpr := &RepeatedExpr{}
		for j := range cmd[i] {
			expr, err := ConvertCommandToMessageExpr(cmd[i][j])
			if err != nil {
				return nil, err
			}
			msgRepeatedExpr.Expr = append(msgRepeatedExpr.Expr, expr)
		}
		msgRepeatedExprSlice = append(msgRepeatedExprSlice, msgRepeatedExpr)
	}
	return msgRepeatedExprSlice, nil
}

// ConvertCommandToMessageListValues converts a command.Values to a message.List_Values.
func ConvertCommandToMessageListValues(cmd *command.Values) (*List_Values, error) {
	exprSlice, err := ConvertCommandToMessageRepeatedExprSlice(cmd.Values)
	if err != nil {
		return nil, err
	}
	return &List_Values{
		&Command_Values{
			Expr: exprSlice,
		},
	}, nil
}

// ConvertCommandToMessageList converts
func ConvertCommandToMessageList(cmd command.List) (*List, error) {
	var err error
	msgList := &List{}
	switch c := cmd.(type) {
	case *command.Scan:
		msgList.List, err = ConvertCommandToMessageListScan(c)
		if err != nil {
			return nil, err
		}
	case *command.Select:
		msgList.List, err = ConvertCommandToMessageListSelect(c)
		if err != nil {
			return nil, err
		}
	case *command.Project:
		msgList.List, err = ConvertCommandToMessageListProject(c)
		if err != nil {
			return nil, err
		}
	case *command.Join:
		msgList.List, err = ConvertCommandToMessageListJoin(c)
		if err != nil {
			return nil, err
		}
	case *command.Limit:
		msgList.List, err = ConvertCommandToMessageListLimit(c)
		if err != nil {
			return nil, err
		}
	case *command.Offset:
		msgList.List, err = ConvertCommandToMessageListOffset(c)
		if err != nil {
			return nil, err
		}
	case *command.Distinct:
		msgList.List, err = ConvertCommandToMessageListDistinct(c)
		if err != nil {
			return nil, err
		}
	case *command.Values:
		msgList.List, err = ConvertCommandToMessageListValues(c)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnknownCommandKind
	}
	return msgList, nil
}

// ConvertCommandToMessageSelect converts a Command type to a Command_Select type.
func ConvertCommandToMessageSelect(cmd *command.Select) (*Command_Select, error) {
	filter, err := ConvertCommandToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Select{
		Filter: filter,
		Input:  input,
	}, nil
}

// ConvertCommandToMessageCol converts command.Column to a message.Column.
func ConvertCommandToMessageCol(cmd command.Column) (*Column, error) {
	column, err := ConvertCommandToMessageExpr(cmd.Column)
	if err != nil {
		return nil, err
	}
	return &Column{
		Table:  cmd.Table,
		Column: column,
		Alias:  cmd.Alias,
	}, nil
}

// ConvertCommandToMessageColSlice converts []command.Column to a []message.Column.
func ConvertCommandToMessageColSlice(cmd []command.Column) ([]*Column, error) {
	msgCols := []*Column{}
	for i := range cmd {
		col, err := ConvertCommandToMessageCol(cmd[i])
		if err != nil {
			return nil, err
		}
		msgCols = append(msgCols, col)
	}
	return msgCols, nil
}

// ConvertCommandToMessageProject converts a Command type to a Command_Project type.
func ConvertCommandToMessageProject(cmd command.Command) (*Command_Project, error) {
	cols, err := ConvertCommandToMessageColSlice(cmd.(*command.Project).Cols)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.(*command.Project).Input)
	if err != nil {
		return nil, err
	}
	return &Command_Project{
		Cols:  cols,
		Input: input,
	}, nil
}

// ConvertCommandToMessageDelete converts a Command type to a Command_Delete type.
func ConvertCommandToMessageDelete(cmd *command.Delete) (*Command_Delete, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	filter, err := ConvertCommandToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	return &Command_Delete{
		Table:  table,
		Filter: filter,
	}, nil
}

// ConvertCommandToMessageDrop converts a Command type to a CommandDrop type.
func ConvertCommandToMessageDrop(cmd command.Command) (*CommandDrop, error) {
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

// ConvertCommandToMessageUpdateOr converts a command.Update or to a message.UpdateOr.
// Returns -1 if the UpdateOr type doesn't match.
func ConvertCommandToMessageUpdateOr(cmd command.UpdateOr) (UpdateOr, error) {
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

// ConvertCommandToMessageUpdateSetterLiteral converts a command.Literal to a message.UpdateSetter_Literal.
func ConvertCommandToMessageUpdateSetterLiteral(cmd command.LiteralExpr) (*UpdateSetter_Literal, error) {
	return &UpdateSetter_Literal{
		&LiteralExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandToMessageUpdateSetterConstant converts a command.Constant to a message.UpdateSetter_Constant.
func ConvertCommandToMessageUpdateSetterConstant(cmd command.ConstantBooleanExpr) (*UpdateSetter_Constant, error) {

	return &UpdateSetter_Constant{
		&ConstantBooleanExpr{
			Value: cmd.Value,
		},
	}, nil
}

// ConvertCommandToMessageUpdateSetterUnary converts a command.Unary to a message.UpdateSetter_Unary.
func ConvertCommandToMessageUpdateSetterUnary(cmd command.UnaryExpr) (*UpdateSetter_Unary, error) {
	val, err := ConvertCommandToMessageExpr(cmd.Value)
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

// ConvertCommandToMessageUpdateSetterBinary converts a command.Binary to a message.UpdateSetter_Binary.
func ConvertCommandToMessageUpdateSetterBinary(cmd command.BinaryExpr) (*UpdateSetter_Binary, error) {
	left, err := ConvertCommandToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandToMessageExpr(cmd.Right)
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

// ConvertCommandToMessageUpdateSetterFunc converts a command.Func to a message.UpdateSetter_Func.
func ConvertCommandToMessageUpdateSetterFunc(cmd command.FunctionExpr) (*UpdateSetter_Func, error) {
	repExpr, err := ConvertCommandToMessageRepeatedExpr(cmd.Args)
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

// ConvertCommandToMessageUpdateSetterEquality converts a command.Equality to a message.UpdateSetter_Equality.
func ConvertCommandToMessageUpdateSetterEquality(cmd command.EqualityExpr) (*UpdateSetter_Equality, error) {
	left, err := ConvertCommandToMessageExpr(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandToMessageExpr(cmd.Right)
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

// ConvertCommandToMessageUpdateSetterRange converts a command.Range to a message.UpdateSetter_Range.
func ConvertCommandToMessageUpdateSetterRange(cmd command.RangeExpr) (*UpdateSetter_Range, error) {
	needle, err := ConvertCommandToMessageExpr(cmd.Needle)
	if err != nil {
		return nil, err
	}
	lo, err := ConvertCommandToMessageExpr(cmd.Lo)
	if err != nil {
		return nil, err
	}
	hi, err := ConvertCommandToMessageExpr(cmd.Hi)
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

// ConvertCommandToMessageUpdateSetter converts a command.UpdateSetter to a message.UpdateSetter.
func ConvertCommandToMessageUpdateSetter(cmd command.UpdateSetter) (*UpdateSetter, error) {
	var err error
	msgUpdateSetter := &UpdateSetter{}
	msgUpdateSetter.Cols = cmd.Cols
	switch val := cmd.Value.(type) {
	case command.LiteralExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterLiteral(val)
		if err != nil {
			return nil, err
		}
	case command.ConstantBooleanExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterConstant(val)
		if err != nil {
			return nil, err
		}
	case command.UnaryExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterUnary(val)
		if err != nil {
			return nil, err
		}
	case command.BinaryExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterBinary(val)
		if err != nil {
			return nil, err
		}
	case command.FunctionExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterFunc(val)
		if err != nil {
			return nil, err
		}
	case command.EqualityExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterEquality(val)
		if err != nil {
			return nil, err
		}
	case command.RangeExpr:
		msgUpdateSetter.Value, err = ConvertCommandToMessageUpdateSetterRange(val)
		if err != nil {
			return nil, err
		}
	}
	return msgUpdateSetter, nil
}

// ConvertCommandToMessageUpdateSetterSlice converts a []command.UpdateSetter to a []message.UpdateSetter.
func ConvertCommandToMessageUpdateSetterSlice(cmd []command.UpdateSetter) ([]*UpdateSetter, error) {
	msgUpdateSetterSlice := []*UpdateSetter{}
	for i := range cmd {
		updateSetter, err := ConvertCommandToMessageUpdateSetter(cmd[i])
		if err != nil {
			return nil, err
		}
		msgUpdateSetterSlice = append(msgUpdateSetterSlice, updateSetter)
	}
	return msgUpdateSetterSlice, nil
}

// ConvertCommandToMessageUpdate converts a Command type to a Command_Update type.
func ConvertCommandToMessageUpdate(cmd command.Command) (*Command_Update, error) {
	updateOr, err := ConvertCommandToMessageUpdateOr(cmd.(*command.Update).UpdateOr)
	if err != nil {
		return nil, err
	}
	table, err := ConvertCommandTableToMessageTable(cmd.(*command.Update).Table)
	if err != nil {
		return nil, err
	}
	updates, err := ConvertCommandToMessageUpdateSetterSlice(cmd.(*command.Update).Updates)
	if err != nil {
		return nil, err
	}
	filter, err := ConvertCommandToMessageExpr(cmd.(*command.Update).Filter)
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
func ConvertCommandToMessageJoin(cmd *command.Join) (*Command_Join, error) {
	filter, err := ConvertCommandToMessageExpr(cmd.Filter)
	if err != nil {
		return nil, err
	}
	left, err := ConvertCommandToMessageList(cmd.Left)
	if err != nil {
		return nil, err
	}
	right, err := ConvertCommandToMessageList(cmd.Right)
	if err != nil {
		return nil, err
	}

	return &Command_Join{
		Natural: cmd.Natural,
		Type:    ConvertCommandToMessageJoinType(cmd.Type),
		Filter:  filter,
		Left:    left,
		Right:   right,
	}, nil
}

// ConvertCommandToMessageLimit converts a Command type to a Command_Limit type.
func ConvertCommandToMessageLimit(cmd *command.Limit) (*Command_Limit, error) {
	limit, err := ConvertCommandToMessageExpr(cmd.Limit)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Limit{
		Limit: limit,
		Input: input,
	}, nil
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
func ConvertCommandToMessageInsert(cmd *command.Insert) (*Command_Insert, error) {
	table, err := ConvertCommandTableToMessageTable(cmd.Table)
	if err != nil {
		return nil, err
	}
	colSlice, err := ConvertCommandToMessageColSlice(cmd.Cols)
	if err != nil {
		return nil, err
	}
	input, err := ConvertCommandToMessageList(cmd.Input)
	if err != nil {
		return nil, err
	}
	return &Command_Insert{
		InsertOr:      ConvertCommandToMessageInsertOr(cmd.InsertOr),
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
	return &command.SimpleTable{
		Schema:  msg.Schema,
		Table:   msg.Table,
		Alias:   msg.Alias,
		Indexed: msg.Indexed,
		Index:   msg.Index,
	}
}

// ConvertMessageToCommandScan converts a message.Command_Scan to a command.Scan.
func ConvertMessageToCommandScan(msg *Command_Scan) *command.Scan {
	return &command.Scan{
		Table: ConvertMessageToCommandTable(msg.Table),
	}
}

// ConvertMessageToCommandLiteralExpr converts a message.Expr to a command.LiteralExpr.
func ConvertMessageToCommandLiteralExpr(msg *Expr) *command.LiteralExpr {
	return &command.LiteralExpr{
		Value: msg.GetLiteral().GetValue(),
	}
}

// ConvertMessageToCommandConstantBooleanExpr converts a message.Expr to a command.ConstantBooleanExpr.
func ConvertMessageToCommandConstantBooleanExpr(msg *Expr) *command.ConstantBooleanExpr {
	return &command.ConstantBooleanExpr{
		Value: msg.GetConstant().GetValue(),
	}
}

// ConvertMessageToCommandUnaryExpr converts a message.Expr to a command.UnaryExpr.
func ConvertMessageToCommandUnaryExpr(msg *Expr) *command.UnaryExpr {
	return &command.UnaryExpr{
		Operator: msg.GetUnary().GetOperator(),
		Value:    ConvertMessageToCommandExpr(msg.GetUnary().GetValue()),
	}
}

// ConvertMessageToCommandBinaryExpr converts a message.Expr to a command.BinaryExpr.
func ConvertMessageToCommandBinaryExpr(msg *Expr) *command.BinaryExpr {
	return &command.BinaryExpr{
		Operator: msg.GetBinary().GetOperator(),
		Left:     ConvertMessageToCommandExpr(msg.GetBinary().GetLeft()),
		Right:    ConvertMessageToCommandExpr(msg.GetBinary().GetRight()),
	}
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
	return &command.FunctionExpr{
		Name:     msg.GetFunc().GetName(),
		Distinct: msg.GetFunc().GetDistinct(),
		Args:     ConvertMessageToCommandExprSlice(msg.GetFunc().GetArgs()),
	}
}

// ConvertMessageToCommandEqualityExpr converts a message.Expr to a command.EqualityExpr.
func ConvertMessageToCommandEqualityExpr(msg *Expr) *command.EqualityExpr {
	return &command.EqualityExpr{
		Left:   ConvertMessageToCommandExpr(msg.GetEquality().GetLeft()),
		Right:  ConvertMessageToCommandExpr(msg.GetEquality().GetRight()),
		Invert: msg.GetEquality().Invert,
	}
}

// ConvertMessageToCommandRangeExpr converts a message.Expr to a command.RangeExpr.
func ConvertMessageToCommandRangeExpr(msg *Expr) *command.RangeExpr {
	return &command.RangeExpr{
		Needle: ConvertMessageToCommandExpr(msg.GetRange().GetNeedle()),
		Lo:     ConvertMessageToCommandExpr(msg.GetRange().GetLo()),
		Hi:     ConvertMessageToCommandExpr(msg.GetRange().GetHi()),
	}
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
	return &command.Scan{
		Table: ConvertMessageToCommandTable(msg.GetScan().GetTable()),
	}
}

// ConvertMessageToCommandListSelect converts a message.List to a command.Select.
func ConvertMessageToCommandListSelect(msg *List) *command.Select {
	return &command.Select{
		Filter: ConvertMessageToCommandExpr(msg.GetSelect().GetFilter()),
		Input:  ConvertMessageToCommandList(msg.GetSelect().GetInput()),
	}
}

// ConvertMessageToCommandListProject converts a message.List to a command.Project.
func ConvertMessageToCommandListProject(msg *List) *command.Project {
	return &command.Project{
		Cols:  ConvertMessageToCommandCols(msg.GetProject().GetCols()),
		Input: ConvertMessageToCommandList(msg.GetProject().GetInput()),
	}
}

// ConvertMessageToCommandListJoin converts a message.List to a command.Join.
func ConvertMessageToCommandListJoin(msg *List) *command.Join {
	return &command.Join{
		Natural: msg.GetJoin().GetNatural(),
		Type:    ConvertMessageToCommandJoinType(msg.GetJoin().GetType()),
		Filter:  ConvertMessageToCommandExpr(msg.GetJoin().GetFilter()),
		Left:    ConvertMessageToCommandList(msg.GetJoin().GetLeft()),
		Right:   ConvertMessageToCommandList(msg.GetJoin().GetRight()),
	}
}

// ConvertMessageToCommandListLimit converts a message.List to a command.Limit.
func ConvertMessageToCommandListLimit(msg *List) *command.Limit {
	return &command.Limit{
		Limit: ConvertMessageToCommandExpr(msg.GetLimit().GetLimit()),
		Input: ConvertMessageToCommandList(msg.GetLimit().GetInput()),
	}
}

// ConvertMessageToCommandListOffset converts a message.List to a command.Offset.
func ConvertMessageToCommandListOffset(msg *List) *command.Offset {
	return &command.Offset{
		Offset: ConvertMessageToCommandExpr(msg.GetOffset().GetOffset()),
		Input:  ConvertMessageToCommandList(msg.GetDistinct().GetInput()),
	}
}

// ConvertMessageToCommandListDistinct converts a message.List to a command.Distinct.
func ConvertMessageToCommandListDistinct(msg *List) *command.Distinct {
	return &command.Distinct{
		Input: ConvertMessageToCommandList(msg.GetDistinct().GetInput()),
	}
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
	return command.Values{
		Values: ConvertMessageToCommandExprRepeatedSlice(msg.GetValues().GetExpr()),
	}
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
	return &command.Select{
		Filter: ConvertMessageToCommandExpr(msg.GetFilter()),
		Input:  ConvertMessageToCommandList(msg.GetInput()),
	}
}

// ConvertMessageToCommandCol converts a message.Column to a command.Column
func ConvertMessageToCommandCol(msg *Column) command.Column {
	return command.Column{
		Table:  msg.GetTable(),
		Column: ConvertMessageToCommandExpr(msg.GetColumn()),
		Alias:  msg.GetAlias(),
	}
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
	return &command.Project{
		Cols:  ConvertMessageToCommandCols(msg.GetCols()),
		Input: ConvertMessageToCommandList(msg.GetInput()),
	}
}

// ConvertMessageToCommandDelete converts a message.Command_Delete to a command.Delete
func ConvertMessageToCommandDelete(msg *Command_Delete) command.Delete {
	return command.Delete{
		Filter: ConvertMessageToCommandExpr(msg.GetFilter()),
		Table:  ConvertMessageToCommandTable(msg.GetTable()),
	}
}

// ConvertMessageToCommandDropTable converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropTable(msg *CommandDrop) command.DropTable {
	return command.DropTable{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageToCommandDropView converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropView(msg *CommandDrop) command.DropView {
	return command.DropView{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageToCommandDropIndex converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropIndex(msg *CommandDrop) command.DropIndex {
	return command.DropIndex{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageToCommandDropTrigger converts a message.CommandDrop to a command.Drop
func ConvertMessageToCommandDropTrigger(msg *CommandDrop) command.DropTrigger {
	return command.DropTrigger{
		IfExists: msg.GetIfExists(),
		Schema:   msg.GetSchema(),
		Name:     msg.GetName(),
	}
}

// ConvertMessageToCommandUpdateOr converts a message.UpdateOr to command.UpdateOr
func ConvertMessageToCommandUpdateOr(msg UpdateOr) command.UpdateOr {
	return command.UpdateOr(msg.Number())
}

// ConvertMessageToCommandUpdateSetterLiteralExpr converts message.LiteralExpr to command.Expr
func ConvertMessageToCommandUpdateSetterLiteralExpr(msg *LiteralExpr) command.Expr {
	return command.LiteralExpr{
		Value: msg.Value,
	}
}

// ConvertMessageToCommandUpdateSetterConstantExpr converts message.ConstantBooleanExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterConstantExpr(msg *ConstantBooleanExpr) command.Expr {
	return command.ConstantBooleanExpr{
		Value: msg.Value,
	}
}

// ConvertMessageToCommandUpdateSetterUnaryExpr converts message.UnaryExpr to command.Expr
func ConvertMessageToCommandUpdateSetterUnaryExpr(msg *UnaryExpr) command.Expr {
	return command.UnaryExpr{
		Operator: msg.Operator,
		Value:    ConvertMessageToCommandBinaryExpr(msg.Value),
	}
}

// ConvertMessageToCommandUpdateSetterBinaryExpr converts message.BinaryExpr to command.Expr
func ConvertMessageToCommandUpdateSetterBinaryExpr(msg *BinaryExpr) command.Expr {
	return command.BinaryExpr{
		Operator: msg.Operator,
		Left:     ConvertMessageToCommandBinaryExpr(msg.Left),
		Right:    ConvertMessageToCommandBinaryExpr(msg.Right),
	}
}

// ConvertMessageToCommandUpdateSetterFuncExpr converts message.FunctionExpr tp command.Expr
func ConvertMessageToCommandUpdateSetterFuncExpr(msg *FunctionExpr) command.Expr {
	return command.FunctionExpr{
		Name:     msg.Name,
		Distinct: msg.Distinct,
		Args:     ConvertMessageToCommandExprSlice(msg.Args),
	}
}

// ConvertMessageToCommandUpdateSetterEqualityExpr converts message.EqualityExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterEqualityExpr(msg *EqualityExpr) command.Expr {
	return command.EqualityExpr{
		Left:   ConvertMessageToCommandBinaryExpr(msg.Left),
		Right:  ConvertMessageToCommandBinaryExpr(msg.Right),
		Invert: msg.Invert,
	}
}

// ConvertMessageToCommandUpdateSetterRangeExpr converts a message.RangeExpr to a command.Expr
func ConvertMessageToCommandUpdateSetterRangeExpr(msg *RangeExpr) command.Expr {
	return command.RangeExpr{
		Needle: ConvertMessageToCommandBinaryExpr(msg.Needle),
		Lo:     ConvertMessageToCommandBinaryExpr(msg.Lo),
		Hi:     ConvertMessageToCommandBinaryExpr(msg.Hi),
		Invert: msg.Invert,
	}
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
	return command.Update{
		UpdateOr: ConvertMessageToCommandUpdateOr(msg.GetUpdateOr()),
		Updates:  ConvertMessageToCommandUpdateSetterSlice(msg.GetUpdates()),
		Table:    ConvertMessageToCommandTable(msg.GetTable()),
		Filter:   ConvertMessageToCommandExpr(msg.GetFilter()),
	}
}

// ConvertMessageToCommandJoinType converts a message.JoinType to a command.JoinType
func ConvertMessageToCommandJoinType(msg JoinType) command.JoinType {
	return command.JoinType(msg.Number())
}

// ConvertMessageToCommandJoin converts a message.Command_Join to a command.Join
func ConvertMessageToCommandJoin(msg *Command_Join) command.Join {
	return command.Join{
		Natural: msg.Natural,
		Type:    ConvertMessageToCommandJoinType(msg.GetType()),
		Filter:  ConvertMessageToCommandExpr(msg.GetFilter()),
		Left:    ConvertMessageToCommandList(msg.GetLeft()),
		Right:   ConvertMessageToCommandList(msg.GetRight()),
	}
}

// ConvertMessageToCommandLimit converts a message.Command_Limit to a command.Limit
func ConvertMessageToCommandLimit(msg *Command_Limit) command.Limit {
	return command.Limit{
		Limit: ConvertMessageToCommandExpr(msg.GetLimit()),
		Input: ConvertMessageToCommandList(msg.GetInput()),
	}
}

// ConvertMessageToCommandInsertOr converts a message.InsertOr to command.InsertOr
func ConvertMessageToCommandInsertOr(msg InsertOr) command.InsertOr {
	return command.InsertOr(msg.Number())
}

// ConvertMessageToCommandInsert converts a message.Command_Insert to a command.Insert
func ConvertMessageToCommandInsert(msg *Command_Insert) command.Insert {
	return command.Insert{
		InsertOr:      ConvertMessageToCommandInsertOr(msg.GetInsertOr()),
		Table:         ConvertMessageToCommandTable(msg.GetTable()),
		Cols:          ConvertMessageToCommandCols(msg.GetCols()),
		DefaultValues: msg.GetDefaultValues(),
		Input:         ConvertMessageToCommandList(msg.GetInput()),
	}
}
