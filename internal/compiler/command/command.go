package command

// Command describes a data structure that can be used by the executor to
// manipulate the database.
type Command interface {
	Type() Type
}
