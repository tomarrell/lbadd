package compiler

import "testing"

func TestCompileSelect(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"simple select",
			"SELECT * FROM myTable WHERE true"},
		{"simple select",
			"SELECT name FROM myTable WHERE true"},
		{"select distinct",
			"SELECT DISTINCT * FROM myTable WHERE true"},
		{"select with implicit join",
			"SELECT * FROM a, b WHERE true"},
		{"select with explicit join",
			"SELECT * FROM a JOIN b WHERE true"},
		{"select with implicit and explicit join",
			"SELECT * FROM a, b JOIN c WHERE true"},
		{"select expression",
			"SELECT name, amount * price AS total_price FROM items JOIN prices"},
	}
	for _, test := range tests {
		RunGolden(t, test.input, test.name)
	}
}
