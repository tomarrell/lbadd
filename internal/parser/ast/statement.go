package ast

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// Statement
type (
	SQLStmt struct {
		Explain token.Token
		Query   token.Token
		Plan    token.Token

		AlterTableStmt         *AlterTableStmt
		AnalyzeStmt            *AnalyzeStmt
		AttachStmt             *AttachStmt
		BeginStmt              *BeginStmt
		CommitStmt             *CommitStmt
		CreateIndexStmt        *CreateIndexStmt
		CreateTableStmt        *CreateTableStmt
		CreateTriggerStmt      *CreateTriggerStmt
		CreateViewStmt         *CreateViewStmt
		CreateVirtualTableStmt *CreateVirtualTableStmt
		DeleteStmt             *DeleteStmt
		DeleteStmtLimited      *DeleteStmtLimited
		DetachStmt             *DetachStmt
		DropIndexStmt          *DropIndexStmt
		DropTableStmt          *DropTableStmt
		DropTriggerStmt        *DropTriggerStmt
		DropViewStmt           *DropViewStmt
		InsertStmt             *InsertStmt
		// not supported
		// PragmaStmt             *PragmaStmt
		ReindexStmt       *ReindexStmt
		ReleaseStmt       *ReleaseStmt
		RollbackStmt      *RollbackStmt
		SavepointStmt     *SavepointStmt
		SelectStmt        *SelectStmt
		UpdateStmt        *UpdateStmt
		UpdateStmtLimited *UpdateStmtLimited
		VacuumStmt        *VacuumStmt
	}

	AlterTableStmt struct {
		Alter         token.Token
		Table         token.Token
		SchemaName    token.Token
		Period        token.Token
		TableName     token.Token
		Rename        token.Token
		To            token.Token
		NewTableName  token.Token
		Column        token.Token
		ColumnName    token.Token
		NewColumnName token.Token
		Add           token.Token
		ColumnDef     *ColumnDef
	}

	AnalyzeStmt struct {
		Analyze          token.Token
		SchemaName       token.Token
		TableOrIndexName token.Token
		Period           token.Token
	}

	AttachStmt struct {
		Attach     token.Token
		Database   token.Token
		Expr       *Expr
		As         token.Token
		SchemaName token.Token
	}

	BeginStmt struct {
		Begin       token.Token
		Deferred    token.Token
		Immediate   token.Token
		Exclusive   token.Token
		Transaction token.Token
	}

	CommitStmt struct {
		Commit      token.Token
		End         token.Token
		Transaction token.Token
	}

	CreateIndexStmt struct {
		Create         token.Token
		Unique         token.Token
		Index          token.Token
		If             token.Token
		Not            token.Token
		Exists         token.Token
		SchemaName     token.Token
		Period         token.Token
		IndexName      token.Token
		On             token.Token
		TableName      token.Token
		LeftParen      token.Token
		IndexedColumns []token.Token
		RightParen     token.Token
		Where          token.Token
		Expr           *Expr
	}

	CreateTableStmt struct {
		Create          token.Token
		Temp            token.Token
		Temporary       token.Token
		Table           token.Token
		If              token.Token
		Not             token.Token
		Exists          token.Token
		SchemaName      token.Token
		Period          token.Token
		TableName       token.Token
		LeftParen       token.Token
		ColumnDef       []*ColumnDef
		TableConstraint []*TableConstraint
		RightParen      token.Token
		Without         token.Token
		Rowid           token.Token
		As              token.Token
		SelectStmt      *SelectStmt
	}

	CreateTriggerStmt struct {
		Create      token.Token
		Temp        token.Token
		Temporary   token.Token
		Trigger     token.Token
		If          token.Token
		Not         token.Token
		Exists      token.Token
		SchemaName  token.Token
		Period      token.Token
		TriggerName token.Token
		Before      token.Token
		After       token.Token
		Instead     token.Token
		Of1         token.Token
		Delete      token.Token
		Insert      token.Token
		Update      token.Token
		Of2         token.Token
		ColumnName  []token.Token
		On          token.Token
		TableName   token.Token
		For         token.Token
		Each        token.Token
		Row         token.Token
		When        token.Token
		Expr        *Expr
		Begin       token.Token
		UpdateStmt  []*UpdateStmt
		InsertStmt  []*InsertStmt
		DeleteStmt  []*DeleteStmt
		SelectStmt  []*SelectStmt
		End         token.Token
	}

	CreateViewStmt struct {
		Create     token.Token
		Temp       token.Token
		Temporary  token.Token
		View       token.Token
		If         token.Token
		Not        token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		ViewName   token.Token
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
		As         token.Token
		SelectStmt *SelectStmt
	}

	CreateVirtualTableStmt struct {
		Create         token.Token
		Virtual        token.Token
		Table          token.Token
		If             token.Token
		Not            token.Token
		Exists         token.Token
		SchemaName     token.Token
		Period         token.Token
		TableName      token.Token
		Using          token.Token
		ModuleName     token.Token
		LeftParen      token.Token
		ModuleArgument []token.Token
		RightParen     token.Token
	}

	DeleteStmt struct {
		WithClause         *WithClause
		Delete             token.Token
		From               token.Token
		QualifiedTableName *QualifiedTableName
		Where              token.Token
		Expr               *Expr
	}

	DeleteStmtLimited struct {
		*DeleteStmt
		Order        token.Token
		By           token.Token
		OrderingTerm []*OrderingTerm
		Limit        token.Token
		Expr2        *Expr
		Offset       token.Token
		Comma        token.Token
		Expr3        *Expr
	}

	DetachStmt struct {
		Detach     token.Token
		Database   token.Token
		SchemaName token.Token
	}

	DropIndexStmt struct {
		Drop       token.Token
		Index      token.Token
		If         token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		IndexName  token.Token
	}

	DropTableStmt struct {
		Drop       token.Token
		Table      token.Token
		If         token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		TableName  token.Token
	}

	DropTriggerStmt struct {
		Drop        token.Token
		Trigger     token.Token
		If          token.Token
		Exists      token.Token
		SchemaName  token.Token
		Period      token.Token
		TriggerName token.Token
	}

	DropViewStmt struct {
		Drop       token.Token
		View       token.Token
		If         token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		ViewName   token.Token
	}

	QualifiedTableName struct {
		SchemaName token.Token
		Period     token.Token
		TableName  token.Token
		As         token.Token
		Alias      token.Token
		Not        token.Token
		Indexed    token.Token
		By         token.Token
		IndexName  token.Token
	}

	InsertStmt struct {
		WithClause               *WithClause
		Insert                   token.Token
		Or                       token.Token
		Replace                  token.Token
		Rollback                 token.Token
		Abort                    token.Token
		Fail                     token.Token
		Ignore                   token.Token
		Into                     token.Token
		SchemaName               token.Token
		Period                   token.Token
		TableName                token.Token
		As                       token.Token
		Alias                    token.Token
		LeftParen1               token.Token
		ColumnName               []token.Token
		RightParen2              token.Token
		Default                  token.Token
		Values                   token.Token
		SelectStmt               *SelectStmt
		ParenthesizedExpressions *[]ParenthesizedExpressions
		UpsertClause             *UpsertClause
	}

	ParenthesizedExpressions struct {
		LeftParen  token.Token
		Exprs      []*Expr
		RightParen token.Token
	}

	ReindexStmt struct {
		Reindex          token.Token
		CollationName    token.Token
		SchemaName       token.Token
		Period           token.Token
		TableOrIndexName token.Token
	}

	SavepointStmt struct {
		Savepoint     token.Token
		SavepointName token.Token
	}

	ReleaseStmt struct {
		Release       token.Token
		Savepoint     token.Token
		SavepointName token.Token
	}

	RollbackStmt struct {
		Rollback      token.Token
		Transaction   token.Token
		To            token.Token
		Savepoint     token.Token
		SavepointName token.Token
	}

	SelectStmt struct {
		With                  token.Token
		Recursive             token.Token
		CommonTableExpression []*CommonTableExpression
		SelectCore            []*SelectCore
		Order                 token.Token
		By                    token.Token
		OrderingTerm          []*OrderingTerm
		Limit                 token.Token
		Expr1                 *Expr
		Offset                token.Token
		Comma                 token.Token
		Expr2                 *Expr
	}

	SelectCore struct {
		Select                   token.Token
		Distinct                 token.Token
		All                      token.Token
		ResultColumn             []*ResultColumn
		From                     token.Token
		JoinClause               *JoinClause
		TableOrSubquery          []*TableOrSubquery
		Where                    token.Token
		Expr1                    *Expr
		Group                    token.Token
		By                       token.Token
		Expr2                    []*Expr
		Having                   token.Token
		Expr3                    *Expr
		Window                   token.Token
		NamedWindow              []*NamedWindow
		Values                   token.Token
		ParenthesizedExpressions []*ParenthesizedExpressions
		CompoundOperator         *CompoundOperator
	}

	UpdateStmt struct {
		WithClause         *WithClause
		Update             token.Token
		Or                 token.Token
		Rollback           token.Token
		Abort              token.Token
		Replace            token.Token
		Fail               token.Token
		Ignore             token.Token
		QualifiedTableName *QualifiedTableName
		Set                token.Token
		UpdateSetter       []*UpdateSetter
		Where              token.Token
		Expr               *Expr
	}

	UpdateSetter struct {
		ColumnName     token.Token
		ColumnNameList *ColumnNameList
		Assign         token.Token
		Expr           *Expr
	}

	UpdateStmtLimited struct {
		*UpdateStmt
		Order        token.Token
		By           token.Token
		OrderingTerm []token.Token
		Limit        token.Token
		Expr1        *Expr
		Offset       token.Token
		Comma        token.Token
		Expr2        *Expr
	}

	UpsertClause struct {
		On            token.Token
		Conflict      token.Token
		LeftParen     token.Token
		IndexedColumn []*IndexedColumn
		RightParen    token.Token
		Where1        token.Token
		Expr1         *Expr
		Do            token.Token
		Nothing       token.Token
		Update        token.Token
		Set           token.Token
		UpdateSetter  []*UpdateSetter
		Where2        token.Token
		Expr2         *Expr
	}

	VacuumStmt struct {
		Vacuum     token.Token
		SchemaName token.Token
		Into       token.Token
		Filename   token.Token
	}

	WithClause struct {
		With         token.Token
		Recursive    token.Token
		RecursiveCte []*RecursiveCte
	}

	RecursiveCte struct {
		CteTableName *CteTableName
		As           token.Token
		LeftParen    token.Token
		SelectStmt   *SelectStmt
		RightParen   token.Token
	}

	CteTableName struct {
		TableName  token.Token
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
	}
)

// Other
type (
	Expr struct {
		LiteralValue   token.Token
		BindParameter  token.Token
		SchemaName     token.Token
		Period1        token.Token
		TableName      token.Token
		Period2        token.Token
		ColumnName     token.Token
		UnaryOperator  token.Token
		Expr1          *Expr
		BinaryOperator token.Token
		Expr2          *Expr
		FunctionName   token.Token
		LeftParen      token.Token
		Asterisk       token.Token
		Distinct       token.Token
		Expr           []*Expr
		RightParen     token.Token
		FilterClause   *FilterClause
		OverClause     *OverClause
		Cast           token.Token
		As             token.Token
		TypeName       token.Token
		Collate        token.Token
		CollationName  token.Token
		Not            token.Token
		Like           token.Token
		Glob           token.Token
		Regexp         token.Token
		Match          token.Token
		Escape         token.Token
		Expr3          *Expr
		Isnull         token.Token
		Notnull        token.Token
		Null           token.Token
		Is             token.Token
		Between        token.Token
		In             token.Token
		SelectStmt     *SelectStmt
		TableFunction  token.Token
		Exists         token.Token
		Case           token.Token
		When           token.Token
		Then           token.Token
		Else           token.Token
		Expr4          *Expr
		End            token.Token
		RaiseFunction  *RaiseFunction
	}

	FilterClause struct {
		Filter     token.Token
		LeftParen  token.Token
		Where      token.Token
		Expr       *Expr
		RightParen token.Token
	}

	OverClause struct {
		Over           token.Token
		WindowName     token.Token
		LeftParen      token.Token
		BaseWindowName token.Token
		Partition      token.Token
		By             token.Token
		Expr           []*Expr
		Order          token.Token
		OrderingTerm   []*OrderingTerm
		FrameSpec      *FrameSpec
		RightParen     token.Token
	}

	RaiseFunction struct {
		Raise        token.Token
		LeftParen    token.Token
		Ignore       token.Token
		Rollback     token.Token
		Abort        token.Token
		Fail         token.Token
		Comma        token.Token
		ErrorMessage token.Token
		RightParen   token.Token
	}

	OrderingTerm struct {
		Expr          *Expr
		Collate       token.Token
		CollationName token.Token
		Asc           token.Token
		Desc          token.Token
		Nulls         token.Token
		First         token.Token
		Last          token.Token
	}

	ResultColumn struct {
		Expr        *Expr
		As          token.Token
		ColumnAlias token.Token
		Asterisk    token.Token
		TableName   token.Token
		Period      token.Token
	}
)

// Window
type (
	NamedWindow struct {
		WindowName token.Token
		As         token.Token
		WindowDefn *WindowDefn
	}

	WindowDefn struct {
		LeftParen      token.Token
		BaseWindowName token.Token
		Partition      token.Token
		By             token.Token
		Expr           []*Expr
		Order          token.Token
		OrderingTerm   []*OrderingTerm
		FrameSpec      *FrameSpec
		RightParen     token.Token
	}

	FrameSpec struct {
		Range      token.Token
		Rows       token.Token
		Groups     token.Token
		Between    token.Token
		Unbounded1 token.Token
		Preceding1 token.Token
		Expr1      *Expr
		Current1   token.Token
		Row1       token.Token
		Following1 token.Token
		And        token.Token
		Expr2      *Expr
		Preceding2 token.Token
		Current2   token.Token
		Row2       token.Token
		Following2 token.Token
		Unbounded2 token.Token
		Exclude    token.Token
		No         token.Token
		Others     token.Token
		Current3   token.Token
		Row3       token.Token
		Group      token.Token
		Ties       token.Token
	}
)

// Table
type (
	TableConstraint struct {
		Constraint       token.Token
		Name             token.Token
		Primary          token.Token
		Key              token.Token
		Unique           token.Token
		LeftParen        token.Token
		RightParen       token.Token
		IndexedColumn    []*IndexedColumn
		ConflictClause   *ConflictClause
		Check            token.Token
		Expr             *Expr
		Foreign          token.Token
		ColumnName       []token.Token
		ForeignKeyClause *ForeignKeyClause
	}

	ForeignKeyClause struct {
		References   token.Token
		ForeignTable token.Token
		LeftParen    token.Token
		ColumnName   []token.Token
		RightParen   token.Token
		On           token.Token
		Delete       token.Token
		Update       token.Token
		Set          token.Token
		Null         token.Token
		Default      token.Token
		Cascade      token.Token
		Restrict     token.Token
		No           token.Token
		Action       token.Token
		Match        token.Token
		Name         token.Token
		Not          token.Token
		Deferrable   token.Token
		Initially    token.Token
		Deferred     token.Token
		Immediate    token.Token
	}

	CommonTableExpression struct {
		TableName   token.Token
		LeftParen1  token.Token
		ColumnName  []token.Token
		RightParen1 token.Token
		As          token.Token
		LeftParen2  token.Token
		SelectStmt  *SelectStmt
		RightParen2 token.Token
	}

	CompoundOperator struct {
		Union     token.Token
		All       token.Token
		Intersect token.Token
		Except    token.Token
	}

	TableOrSubquery struct {
		SchemaName        token.Token
		Period            token.Token
		TableName         token.Token
		As                token.Token
		TableAlias        token.Token
		Not               token.Token
		Indexed           token.Token
		By                token.Token
		IndexName         token.Token
		TableFunctionName token.Token
		LeftParen         token.Token
		Expr              []*Expr
		RightParen        token.Token
		JoinClause        *JoinClause
		TableOrSubquery   []*TableOrSubquery
		SelectStmt        *SelectStmt
	}
)

// Join
type (
	JoinClause struct {
		TableOrSubquery *TableOrSubquery
		JoinClausePart  *JoinClausePart
	}

	JoinClausePart struct {
		JoinOperator    *JoinOperator
		TableOrSubquery *TableOrSubquery
		JoinConstraint  *JoinConstraint
	}

	JoinConstraint struct {
		On         token.Token
		Expr       *Expr
		Using      token.Token
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
	}

	JoinOperator struct {
		Comma   token.Token
		Natural token.Token
		Left    token.Token
		Outer   token.Token
		Inner   token.Token
		Cross   token.Token
		Join    token.Token
	}
)

// Column
type (
	ColumnDef struct {
		ColumnName       token.Token
		TypeName         *TypeName
		ColumnConstraint []*ColumnConstraint
	}

	ColumnConstraint struct {
		Constraint       token.Token
		Name             token.Token
		Primary          token.Token
		Key              token.Token
		Asc              token.Token
		Desc             token.Token
		ConflictClause   *ConflictClause
		Autoincrement    token.Token
		Not              token.Token
		Null             token.Token
		Unique           token.Token
		Check            token.Token
		LeftParen        token.Token
		Expr             *Expr
		RightParen       token.Token
		Default          token.Token
		SignedNumber     token.Token
		LiteralValue     token.Token
		Collate          token.Token
		CollationName    token.Token
		ForeignKeyClause *ForeignKeyClause
	}

	ColumnNameList struct {
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
	}

	ConflictClause struct {
		On       token.Token
		Conflict token.Token
		Rollback token.Token
		Abort    token.Token
		Fail     token.Token
		Ignore   token.Token
		Replace  token.Token
	}

	TypeName struct {
		Name          []token.Token
		LeftParen     token.Token
		SignedNumber1 token.Token
		Comma         token.Token
		SignedNumber2 token.Token
		RightParen    token.Token
	}

	IndexedColumn struct {
		ColumnName    token.Token
		Expr          *Expr
		Collate       token.Token
		CollationName token.Token
		Asc           token.Token
		Desc          token.Token
	}
)
