// +build gofuzz

package parser

const (
	// DataNotInteresting indicates, that the input was not interesting, meaning
	// that the input was not valid and the parser handled detected and returned
	// an error.
	DataNotInteresting int = 0
	// DataInteresting indicates a valid parser input. The fuzzer should keep it
	// and modify it further.
	DataInteresting int = 1
	// Skip indicates, that the data must not be added to the corpus. You
	// probably shouldn't use it.
	Skip int = -1
)

func Fuzz(data []byte) int {
	input := string(data)
	parser := New(input)
	for {
		stmt, errs, ok := parser.Next()
		if !ok {
			break
		}
		if len(errs) != 0 {
			return DataNotInteresting
		}
		if stmt == nil {
			panic("ok, no errors, but also no statement")
		}
	}
	return DataInteresting
}
