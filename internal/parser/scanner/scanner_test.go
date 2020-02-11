package scanner

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func Test_hasNext(t *testing.T) {
	for k := range keywordMap {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
}

func Test_specificTokenSequence(t *testing.T) {
	input := "SelEct FroMs"
	desiredOutputToken := []token.Token{
		token.New(1, 1, 0, 6, token.KeywordSelect, "SelEct"),
		token.New(1, 8, 7, 5, token.Literal, "FroMs"),
	}
	testSingleString(input, desiredOutputToken, t)
}

func Test_scanOperator(t *testing.T) {
	desiredOutputToken := []token.Token{
		token.New(1, 1, 0, 2, token.BinaryOperator, "||"),
		token.New(1, 1, 0, 1, token.BinaryOperator, "*"),
		token.New(1, 1, 0, 1, token.BinaryOperator, "/"),
		token.New(1, 1, 0, 1, token.BinaryOperator, "%"),
		token.New(1, 1, 0, 1, token.UnaryOperator, "+"),
		token.New(1, 1, 0, 1, token.UnaryOperator, "-"),
		token.New(1, 1, 0, 2, token.BinaryOperator, "<<"),
		token.New(1, 1, 0, 2, token.BinaryOperator, ">>"),
		token.New(1, 1, 0, 1, token.BinaryOperator, "&"),
		token.New(1, 1, 0, 1, token.BinaryOperator, "|"),
		token.New(1, 1, 0, 1, token.BinaryOperator, "<"),
		token.New(1, 1, 0, 2, token.BinaryOperator, "<="),
		token.New(1, 1, 0, 2, token.BinaryOperator, ">="),
		token.New(1, 1, 0, 1, token.BinaryOperator, "="),
		token.New(1, 1, 0, 2, token.BinaryOperator, "=="),
		token.New(1, 1, 0, 2, token.BinaryOperator, "!="),
		token.New(1, 1, 0, 2, token.BinaryOperator, "<>"),
		token.New(1, 1, 0, 1, token.UnaryOperator, "~"),
	}

	var scannerOutput []token.Token
	for i := range desiredOutputToken {
		scanner := New([]rune(desiredOutputToken[i].Value()))
		for scanner.HasNext() {
			nextToken := scanner.Next()
			if nextToken.Type() == token.Error {
				t.Errorf("Error token recieved - %v\n", nextToken.Value())
			} else {
				scannerOutput = append(scannerOutput, nextToken)
			}
		}
	}

	for i, op := range desiredOutputToken {
		if op != scannerOutput[i] {
			t.Errorf("Expected column %v, obtained %v\n", op.Col(), scannerOutput[i].Col())
			t.Errorf("Expected length %v, obtained %v\n", op.Length(), scannerOutput[i].Length())
			t.Errorf("Expected line %v, obtained %v\n", op.Line(), scannerOutput[i].Line())
			t.Errorf("Expected offset %v, obtained %v\n", op.Offset(), scannerOutput[i].Offset())
			t.Errorf("Expected type %v, obtained %v\n", op.Type(), scannerOutput[i].Type())
			t.Errorf("Expected value %v, obtained %v\n", op.Value(), scannerOutput[i].Value())
		}
	}

	if t.Failed() {
		fmt.Printf("Expected tokens %v, obtained %v\n", desiredOutputToken, scannerOutput)
	} else {
		fmt.Printf("Tested on: %v\n", desiredOutputToken)
	}
}

func Test_multipleSequences(t *testing.T) {
	inputSequences := []string{
		"SELECT    *      FROM      users",
		"SELECT      FROM || & +7 59 \"foobar\"",
	}
	desiredOutputToken := [][]token.Token{
		[]token.Token{
			token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
			token.New(1, 11, 10, 1, token.BinaryOperator, "*"),
			token.New(1, 18, 17, 4, token.KeywordFrom, "FROM"),
			token.New(1, 28, 27, 5, token.Literal, "users"),
		},
		[]token.Token{
			token.New(1, 1, 0, 6, token.KeywordSelect, "SELECT"),
			token.New(1, 13, 12, 4, token.KeywordFrom, "FROM"),
			token.New(1, 18, 17, 2, token.BinaryOperator, "||"),
			token.New(1, 21, 20, 1, token.BinaryOperator, "&"),
			token.New(1, 23, 22, 1, token.UnaryOperator, "+"),
			token.New(1, 24, 23, 1, token.Literal, "7"),
			token.New(1, 26, 25, 2, token.Literal, "59"),
			token.New(1, 29, 28, 8, token.Literal, "\"foobar\""),
		},
	}
	for i, k := range inputSequences {
		testSingleString(k, desiredOutputToken[i], t)
	}
}

func testSingleString(input string, desiredOutputToken []token.Token, t *testing.T) {
	var scannerOutput []token.Token
	scanner := New([]rune(input))
	for scanner.HasNext() {
		nextToken := scanner.Next()
		if nextToken != nil {
			if nextToken.Type() != token.Error {
				scannerOutput = append(scannerOutput, nextToken)
			} else {
				t.Errorf("Error token recieved - %v\n", nextToken.Value())
			}
		}
	}
	for i, op := range desiredOutputToken {
		if op != scannerOutput[i] {
			t.Errorf("Expected column %v, obtained %v\n", op.Col(), scannerOutput[i].Col())
			t.Errorf("Expected length %v, obtained %v\n", op.Length(), scannerOutput[i].Length())
			t.Errorf("Expected line %v, obtained %v\n", op.Line(), scannerOutput[i].Line())
			t.Errorf("Expected offset %v, obtained %v\n", op.Offset(), scannerOutput[i].Offset())
			t.Errorf("Expected type %v, obtained %v\n", op.Type(), scannerOutput[i].Type())
			t.Errorf("Expected value %v, obtained %v\n", op.Value(), scannerOutput[i].Value())
		}
	}

	if t.Failed() {
		fmt.Printf("Expected tokens%v, obtained %v\n", desiredOutputToken, scannerOutput)
	} else {
		fmt.Printf("Tested on: %v\n", input)
	}
}

func stringSplit(input string) []string {
	intermediateSlice := strings.Split(input, " ")
	var outSlice []string
	for _, k := range intermediateSlice {
		if k != "" {
			outSlice = append(outSlice, k)
		}
	}
	return outSlice
}
