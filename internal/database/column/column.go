package column

// Column describes a database column, that consists of a type and multiple
// attributes, such as nullability, if it is a primary key etc.
type Column interface {
	Type() Type
	IsNullable() bool
	IsPrimaryKey() bool
	ShouldAutoincrement() bool
	// extend this as we add support for more things, such as default values,
	// uniqueness etc.
}
