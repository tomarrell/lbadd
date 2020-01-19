package ast

type SqlStmt struct {
	Explain bool
	Query   bool
	Plan    bool

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
