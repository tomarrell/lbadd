// +build gofuzz

package parser

// Fuzz is used for fuzzy testing.
func Fuzz(data []byte) int {
	parser := New(string(data))
	var foundErrors bool
	for {
		stmt, errs, ok := parser.Next()
		if !ok {
			break
		}
		foundErrors = foundErrors || len(errs) > 0
		if stmt == nil {
			panic("stmt must never be nil")
		}
	}
	return 1
}
