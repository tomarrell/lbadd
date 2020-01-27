package scanner

import (
	"fmt"
	"testing"
)

func Test_hasNext(t *testing.T) {
	for _, k := range keywordSlice {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
}

func Test_tokenSequence(t *testing.T) {
	input := "SelEc FroM"
	var desiredOutput = []string{"SelEct", "FroM"}
	var output []string
	scanner := New([]rune(input))
	for scanner.HasNext() {
		token := scanner.Next()
		if token != nil {
			output = append(output, token.Value())
		}
	}

	for i, k := range desiredOutput {
		if k != output[i] {
			t.Fail()
		}
	}

	if t.Failed() {
		fmt.Printf("Expected %v, obtained %v\n", desiredOutput, output)
	}
}
