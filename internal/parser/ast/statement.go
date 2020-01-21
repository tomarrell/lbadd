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
		PragmaStmt             *PragmaStmt
		ReindexStmt            *ReindexStmt
		ReleaseStmt            *ReleaseStmt
		RollbackStmt           *RollbackStmt
		SavepointStmt          *SavepointStmt
		SelectStmt             *SelectStmt
		UpdateStmt             *UpdateStmt
		UpdateStmtLimited      *UpdateStmtLimited
		VacuumStmt             *VacuumStmt
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

	RollbackStmt struct {
		Rollback      token.Token
		Transaction   token.Token
		To            token.Token
		Savepoint     token.Token
		SavepointName token.Token
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
		WithClause         *WithClause
		Delete             token.Token
		From               token.Token
		QualifiedTableName *QualifiedTableName
		Where              token.Token
		Expr1              *Expr
		Order              token.Token
		By                 token.Token
		OrderingTerm       []*OrderingTerm
		Limit              token.Token
		Expr2              *Expr
		Offset             token.Token
		Comma              token.Token
		Expr3              *Expr
	}

	DetachStmt struct {
		Detach     token.Token
		Database   token.Token
		SchemaName token.Token
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
		SignedNumber     *SignedNumber
		LiteralValue     token.Token
		Collate          token.Token
		CollationName    token.Token
		ForeignKeyClause *ForeignKeyClause
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
		SignedNumber1 *SignedNumber
		Comma         token.Token
		SignedNumber2 *SignedNumber
		RightParen    token.Token
	}

	SignedNumber struct {
		Plus           token.Token
		Minus          token.Token
		NumericLiteral token.Token
	}
)
