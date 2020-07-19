package ast

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

type (
	// SQLStmt as in the SQLite grammar.
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
		ReIndexStmt            *ReIndexStmt
		ReleaseStmt            *ReleaseStmt
		RollbackStmt           *RollbackStmt
		SavepointStmt          *SavepointStmt
		SelectStmt             *SelectStmt
		UpdateStmt             *UpdateStmt
		UpdateStmtLimited      *UpdateStmtLimited
		VacuumStmt             *VacuumStmt
	}

	// AlterTableStmt as in the SQLite grammar.
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

	// AnalyzeStmt as in the SQLite grammar.
	AnalyzeStmt struct {
		Analyze          token.Token
		SchemaName       token.Token
		TableOrIndexName token.Token
		Period           token.Token
	}

	// AttachStmt as in the SQLite grammar.
	AttachStmt struct {
		Attach     token.Token
		Database   token.Token
		Expr       *Expr
		As         token.Token
		SchemaName token.Token
	}

	// BeginStmt as in the SQLite grammar.
	BeginStmt struct {
		Begin       token.Token
		Deferred    token.Token
		Immediate   token.Token
		Exclusive   token.Token
		Transaction token.Token
	}

	// CommitStmt as in the SQLite grammar.
	CommitStmt struct {
		Commit      token.Token
		End         token.Token
		Transaction token.Token
	}

	// CreateIndexStmt as in the SQLite grammar.
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
		IndexedColumns []*IndexedColumn
		RightParen     token.Token
		Where          token.Token
		Expr           *Expr
	}

	// CreateTableStmt as in the SQLite grammar.
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

	// CreateTriggerStmt as in the SQLite grammar.
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

	// CreateViewStmt as in the SQLite grammar.
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

	// CreateVirtualTableStmt as in the SQLite grammar.
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

	// DeleteStmt as in the SQLite grammar.
	DeleteStmt struct {
		WithClause         *WithClause
		Delete             token.Token
		From               token.Token
		QualifiedTableName *QualifiedTableName
		Where              token.Token
		Expr               *Expr
	}

	// DeleteStmtLimited as in the SQLite grammar.
	DeleteStmtLimited struct {
		*DeleteStmt
		Order        token.Token
		By           token.Token
		OrderingTerm []*OrderingTerm
		Limit        token.Token
		Expr1        *Expr
		Offset       token.Token
		Comma        token.Token
		Expr2        *Expr
	}

	// DetachStmt as in the SQLite grammar.
	DetachStmt struct {
		Detach     token.Token
		Database   token.Token
		SchemaName token.Token
	}

	// DropIndexStmt as in the SQLite grammar.
	DropIndexStmt struct {
		Drop       token.Token
		Index      token.Token
		If         token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		IndexName  token.Token
	}

	// DropTableStmt as in the SQLite grammar.
	DropTableStmt struct {
		Drop       token.Token
		Table      token.Token
		If         token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		TableName  token.Token
	}

	// DropTriggerStmt as in the SQLite grammar.
	DropTriggerStmt struct {
		Drop        token.Token
		Trigger     token.Token
		If          token.Token
		Exists      token.Token
		SchemaName  token.Token
		Period      token.Token
		TriggerName token.Token
	}

	// DropViewStmt as in the SQLite grammar.
	DropViewStmt struct {
		Drop       token.Token
		View       token.Token
		If         token.Token
		Exists     token.Token
		SchemaName token.Token
		Period     token.Token
		ViewName   token.Token
	}

	// QualifiedTableName as in the SQLite grammar.
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

	// InsertStmt as in the SQLite grammar.
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
		LeftParen                token.Token
		ColumnName               []token.Token
		RightParen               token.Token
		Values                   token.Token
		SelectStmt               *SelectStmt
		Default                  token.Token
		ParenthesizedExpressions []*ParenthesizedExpressions
		UpsertClause             *UpsertClause
	}

	// ParenthesizedExpressions as in the SQLite grammar.
	ParenthesizedExpressions struct {
		LeftParen  token.Token
		Exprs      []*Expr
		RightParen token.Token
	}

	// ReIndexStmt as in the SQLite grammar.
	ReIndexStmt struct {
		ReIndex          token.Token
		CollationName    token.Token
		SchemaName       token.Token
		Period           token.Token
		TableOrIndexName token.Token
	}

	// SavepointStmt as in the SQLite grammar.
	SavepointStmt struct {
		Savepoint     token.Token
		SavepointName token.Token
	}

	// ReleaseStmt as in the SQLite grammar.
	ReleaseStmt struct {
		Release       token.Token
		Savepoint     token.Token
		SavepointName token.Token
	}

	// RollbackStmt as in the SQLite grammar.
	RollbackStmt struct {
		Rollback      token.Token
		Transaction   token.Token
		To            token.Token
		Savepoint     token.Token
		SavepointName token.Token
	}

	// SelectStmt as in the SQLite grammar.
	SelectStmt struct {
		WithClause   *WithClause
		SelectCore   []*SelectCore
		Order        token.Token
		By           token.Token
		OrderingTerm []*OrderingTerm
		Limit        token.Token
		Expr1        *Expr
		Offset       token.Token
		Comma        token.Token
		Expr2        *Expr
	}

	// SelectCore as in the SQLite grammar.
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

	// UpdateStmt as in the SQLite grammar.
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

	// UpdateSetter as in the SQLite grammar.
	UpdateSetter struct {
		ColumnName     token.Token
		ColumnNameList *ColumnNameList
		Assign         token.Token
		Expr           *Expr
	}

	// UpdateStmtLimited as in the SQLite grammar.
	UpdateStmtLimited struct {
		*UpdateStmt
		Order        token.Token
		By           token.Token
		OrderingTerm []*OrderingTerm
		Limit        token.Token
		Expr1        *Expr
		Offset       token.Token
		Comma        token.Token
		Expr2        *Expr
	}

	// UpsertClause as in the SQLite grammar.
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

	// VacuumStmt as in the SQLite grammar.
	VacuumStmt struct {
		Vacuum     token.Token
		SchemaName token.Token
		Into       token.Token
		Filename   token.Token
	}

	// WithClause as in the SQLite grammar.
	WithClause struct {
		With         token.Token
		Recursive    token.Token
		RecursiveCte []*RecursiveCte
	}

	// RecursiveCte as in the SQLite grammar.
	RecursiveCte struct {
		CteTableName *CteTableName
		As           token.Token
		LeftParen    token.Token
		SelectStmt   *SelectStmt
		RightParen   token.Token
	}

	// CteTableName as in the SQLite grammar.
	CteTableName struct {
		TableName  token.Token
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
	}

	// Expr as in the SQLite grammar.
	Expr struct {
		LiteralValue   token.Token
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
		TypeName       *TypeName
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
		And            token.Token
		In             token.Token
		SelectStmt     *SelectStmt
		TableFunction  token.Token
		Exists         token.Token
		Case           token.Token
		WhenThenClause []*WhenThenClause
		Else           token.Token
		End            token.Token
		RaiseFunction  *RaiseFunction
	}

	// WhenThenClause as in the SQLite grammar
	WhenThenClause struct {
		When  token.Token
		Expr1 *Expr
		Then  token.Token
		Expr2 *Expr
	}

	// FilterClause as in the SQLite grammar.
	FilterClause struct {
		Filter     token.Token
		LeftParen  token.Token
		Where      token.Token
		Expr       *Expr
		RightParen token.Token
	}

	// OverClause as in the SQLite grammar.
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

	// RaiseFunction as in the SQLite grammar.
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

	// OrderingTerm as in the SQLite grammar.
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

	// ResultColumn as in the SQLite grammar.
	ResultColumn struct {
		Expr        *Expr
		As          token.Token
		ColumnAlias token.Token
		Asterisk    token.Token
		TableName   token.Token
		Period      token.Token
	}

	// NamedWindow as in the SQLite grammar.
	NamedWindow struct {
		WindowName token.Token
		As         token.Token
		WindowDefn *WindowDefn
	}

	// WindowDefn as in the SQLite grammar.
	WindowDefn struct {
		LeftParen      token.Token
		BaseWindowName token.Token
		Partition      token.Token
		By1            token.Token
		Expr           []*Expr
		Order          token.Token
		By2            token.Token
		OrderingTerm   []*OrderingTerm
		FrameSpec      *FrameSpec
		RightParen     token.Token
	}

	// FrameSpec as in the SQLite grammar.
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

	// TableConstraint as in the SQLite grammar.
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

	// ForeignKeyClause as in the SQLite grammar.
	ForeignKeyClause struct {
		References           token.Token
		ForeignTable         token.Token
		LeftParen            token.Token
		ColumnName           []token.Token
		RightParen           token.Token
		ForeignKeyClauseCore []*ForeignKeyClauseCore
		Not                  token.Token
		Deferrable           token.Token
		Initially            token.Token
		Deferred             token.Token
		Immediate            token.Token
	}

	// ForeignKeyClauseCore as in the SQLite grammar.
	ForeignKeyClauseCore struct {
		On       token.Token
		Delete   token.Token
		Update   token.Token
		Set      token.Token
		Null     token.Token
		Default  token.Token
		Cascade  token.Token
		Restrict token.Token
		No       token.Token
		Action   token.Token
		Match    token.Token
		Name     token.Token
	}

	// CommonTableExpression as in the SQLite grammar.
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

	// CompoundOperator as in the SQLite grammar.
	CompoundOperator struct {
		Union     token.Token
		All       token.Token
		Intersect token.Token
		Except    token.Token
	}

	// TableOrSubquery as in the SQLite grammar.
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

	// JoinClause as in the SQLite grammar.
	JoinClause struct {
		TableOrSubquery *TableOrSubquery
		JoinClausePart  []*JoinClausePart
	}

	// JoinClausePart as in the SQLite grammar.
	JoinClausePart struct {
		JoinOperator    *JoinOperator
		TableOrSubquery *TableOrSubquery
		JoinConstraint  *JoinConstraint
	}

	// JoinConstraint as in the SQLite grammar.
	JoinConstraint struct {
		On         token.Token
		Expr       *Expr
		Using      token.Token
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
	}

	// JoinOperator as in the SQLite grammar.
	JoinOperator struct {
		Comma   token.Token
		Natural token.Token
		Left    token.Token
		Outer   token.Token
		Inner   token.Token
		Cross   token.Token
		Join    token.Token
	}

	// ColumnDef as in the SQLite grammar.
	ColumnDef struct {
		ColumnName       token.Token
		TypeName         *TypeName
		ColumnConstraint []*ColumnConstraint
	}

	// ColumnConstraint as in the SQLite grammar.
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
		SignedNumber     *SignedNumber
		LiteralValue     token.Token
		Collate          token.Token
		CollationName    token.Token
		ForeignKeyClause *ForeignKeyClause
		Generated        token.Token
		Always           token.Token
		As               token.Token
		Stored           token.Token
		Virtual          token.Token
	}

	// ColumnNameList as in the SQLite grammar.
	ColumnNameList struct {
		LeftParen  token.Token
		ColumnName []token.Token
		RightParen token.Token
	}

	// ConflictClause as in the SQLite grammar.
	ConflictClause struct {
		On       token.Token
		Conflict token.Token
		Rollback token.Token
		Abort    token.Token
		Fail     token.Token
		Ignore   token.Token
		Replace  token.Token
	}

	// TypeName as in the SQLite grammar.
	TypeName struct {
		Name          []token.Token
		LeftParen     token.Token
		SignedNumber1 *SignedNumber
		Comma         token.Token
		SignedNumber2 *SignedNumber
		RightParen    token.Token
	}

	// SignedNumber as in the SQLite grammar.
	SignedNumber struct {
		Sign           token.Token
		NumericLiteral token.Token
	}

	// IndexedColumn as in the SQLite grammar.
	IndexedColumn struct {
		ColumnName    token.Token
		Expr          *Expr
		Collate       token.Token
		CollationName token.Token
		Asc           token.Token
		Desc          token.Token
	}
)
