package compiler

import "testing"

func TestCompileGolden(t *testing.T) {
	t.Run("select", _TestCompileSelect)
	t.Run("delete", _TestCompileDelete)
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

func _TestCompileSelect(t *testing.T) {
	tests := []string{
		"SELECT * FROM myTable WHERE true",
		"SELECT name FROM myTable WHERE true",
		"SELECT DISTINCT * FROM myTable WHERE true",
		"SELECT * FROM a, b WHERE true",
		"SELECT * FROM a JOIN b WHERE true",
		"SELECT * FROM a, b JOIN c WHERE true",
		"SELECT name, amount * price AS total_price FROM items JOIN prices",
		"SELECT col1 FROM a, b NATURAL JOIN c, d, e LEFT OUTER JOIN f CROSS JOIN g, h, i",
	}
	for _, test := range tests {
		RunGolden(t, test)
	}
}
