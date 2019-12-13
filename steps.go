//go:generate stringer -type=step

package main

type step int

const (
	stepInit step = iota
	stepSelectField
	stepSelectComma
	stepSelectFrom
	stepSelectTable
)
