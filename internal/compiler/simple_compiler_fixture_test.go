package compiler

import "testing"

func TestCompileGolden(t *testing.T) {
	t.Run("select", _TestCompileSelect)
	t.Run("delete", _TestCompileDelete)
	t.Run("drop", _TestCompileDrop)
}

func _TestCompileDelete(t *testing.T) {
	tests := []string{
		"DELETE FROM myTable",
		"DELETE FROM mySchema.myTable",
		"DELETE FROM myTable WHERE col1 == col2",
	}
	for _, test := range tests {
		RunGolden(t, test)
	}
}

func _TestCompileDrop(t *testing.T) {
	tests := []string{
		"DROP TABLE myTable",
		"DROP TABLE IF EXISTS myTable",
		"DROP TABLE mySchema.myTable",
		"DROP TABLE IF EXISTS mySchema.myTable",
		"DROP VIEW myView",
		"DROP VIEW IF EXISTS myView",
		"DROP VIEW mySchema.myView",
		"DROP VIEW IF EXISTS mySchema.myView",
		"DROP INDEX myIndex",
		"DROP INDEX IF EXISTS myIndex",
		"DROP INDEX mySchema.myIndex",
		"DROP INDEX IF EXISTS mySchema.myIndex",
		"DROP TRIGGER myTrigger",
		"DROP TRIGGER IF EXISTS myTrigger",
		"DROP TRIGGER mySchema.myTrigger",
		"DROP TRIGGER IF EXISTS mySchema.myTrigger",
	}
	for _, test := range tests {
		RunGolden(t, test)
	}
}

func _TestCompileSelect(t *testing.T) {
	tests := []string{
		"SELECT * FROM myTable",
		"SELECT * FROM myTable WHERE true",
		"SELECT * FROM myTable LIMIT 5",
		"SELECT * FROM myTable LIMIT 5, 10",
		"SELECT * FROM myTable LIMIT 5 OFFSET 10",
		"SELECT DISTINCT * FROM myTable WHERE true",
		"SELECT * FROM a, b WHERE true",
		"SELECT * FROM a JOIN b WHERE true",
		"SELECT * FROM a, b JOIN c WHERE true",
		"SELECT name, amount * price AS total_price FROM items JOIN prices",
		"SELECT AVG(price) AS avg_price FROM items LEFT JOIN prices",
		"SELECT AVG(DISTINCT price) AS avg_price FROM items LEFT JOIN prices",
	}
	for _, test := range tests {
		RunGolden(t, test)
	}
}
