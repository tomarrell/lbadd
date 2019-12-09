package main

type parser struct {
	i               int
	sql             string
	step            step
	query           query
	err             error
	nextUpdateField string
}
