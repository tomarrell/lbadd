package main

type db struct {
	tree *btree
}

func newDB() *db {
	return &db{}
}
