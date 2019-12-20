//go:generate stringer -type=step

package lbadd

type step int

const (
	stepInit step = iota
	stepSelectField
	stepSelectComma
	stepSelectFrom
	stepSelectTable
)
