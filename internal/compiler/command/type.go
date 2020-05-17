package command

//go:generate stringer -type=Type

// Type represents the type of a top level command. Nested commands may have a
// different type.
type Type uint32

// Known types of commands
const (
	TypeUnknown Type = iota
	TypeSelect
)
