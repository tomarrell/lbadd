package scanner

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
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

func Test_specificTokenSequence(t *testing.T) {
	input := "SelEct FroMs"
	desiredOutput := []string{
		"SelEct",
		"FroMs",
	}
	desiredOutputType := []token.Type{
		token.KeywordSelect,
		token.KeywordFrom,
	}
	testSingleString(input, desiredOutput, desiredOutputType, t)
}

func Test_randomTokenSequence(t *testing.T) {
	size := 5
	randomSequence, randomSequenceArray := generateRandomKeywords(size)
	var scannerOutput []string
	scanner := New([]rune(randomSequence))
	for scanner.HasNext() {
		token := scanner.Next()
		if token != nil {
			scannerOutput = append(scannerOutput, token.Value())
		}
	}
	for i, k := range scannerOutput {
		if k != randomSequenceArray[i] {
			t.Fail()
		}
	}

	if t.Failed() {
		fmt.Printf("Expected %v, obtained %v\n", randomSequenceArray, scannerOutput)
	} else {
		fmt.Printf("Tested on: %v\n", randomSequenceArray)
	}
}

func Test_scanOperator(t *testing.T) {
	desiredOutput := []string{
		"||",
		"*",
		"/",
		"%",
		"+",
		"-",
		"<<",
		">>",
		"&",
		"|",
		"<",
		"<=",
		">=",
		"=",
		"==",
		"!=",
		"<>",
		"~",
	}
	desiredOutputType := []token.Type{
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.UnaryOperator,
		token.UnaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.BinaryOperator,
		token.UnaryOperator,
	}

	var scannerOutput []token.Type
	for i := range desiredOutput {
		scanner := New([]rune(desiredOutput[i]))
		for scanner.HasNext() {
			nextToken := scanner.Next()
			if nextToken.Type() == token.Error {
				t.Errorf("Error token recieved - %v\n", nextToken.Value())
			} else {
				scannerOutput = append(scannerOutput, nextToken.Type())
			}
		}
	}

	for i, op := range desiredOutputType {
		if op != scannerOutput[i] {
			t.Errorf("Expected %v, obtained %v\n", op, scannerOutput[i])
		}
	}

	if t.Failed() {
		fmt.Printf("Expected %v, obtained %v\n", desiredOutputType, scannerOutput)
	} else {
		fmt.Printf("Tested on: %v\n", desiredOutput)
	}
}

func Test_multipleSequences(t *testing.T) {
	inputSequences := []string{
		"SELECT    *      FROM      users",
		"SELECT       FROM || & +7 59 \"foobar\"",
	}
	desiredOutput := [][]string{
		[]string{
			"SELECT",
			"*",
			"FROM",
			"users",
		},
		[]string{
			"SELECT",
			"FROM",
			"||",
			"&",
			"+",
			"7",
			"59",
			"\"foobar\"",
		},
	}
	desiredOutputType := [][]token.Type{
		[]token.Type{
			token.KeywordSelect,
			token.BinaryOperator,
			token.KeywordFrom,
			token.Literal,
		},
		[]token.Type{
			token.KeywordSelect,
			token.KeywordFrom,
			token.BinaryOperator,
			token.BinaryOperator,
			token.UnaryOperator,
			token.Literal,
			token.Literal,
			token.Literal,
		},
	}
	for i, k := range inputSequences {
		testSingleString(k, desiredOutput[i], desiredOutputType[i], t)
	}
}

func generateRandomKeywords(size int) (string, []string) {
	var randKeywordArray []string
	randomKeywords := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		keyword := keywordSlice[rand.Intn(len(keywordSlice))]
		randKeywordArray = append(randKeywordArray, keyword)
		randomKeywords += keyword
		randomKeywords += " "
	}
	return randomKeywords, randKeywordArray
}

func testSingleString(input string, desiredOutput []string, desiredOutputType []token.Type, t *testing.T) {
	var scannerOutput []string
	var scannerOutputType []token.Type
	scanner := New([]rune(input))
	for scanner.HasNext() {
		nextToken := scanner.Next()
		if nextToken != nil {
			if nextToken.Type() != token.Error {
				scannerOutput = append(scannerOutput, nextToken.Value())
				scannerOutputType = append(scannerOutputType, nextToken.Type())
			} else {
				t.Errorf("Error token recieved - %v\n", nextToken.Value())
			}
		}
	}

	for i, k := range desiredOutputType {
		if k != scannerOutputType[i] {
			t.Errorf("Mismatch in desired (%v) and output (%v) token types\n", k, scannerOutputType[i])
		}
		if desiredOutput[i] != scannerOutput[i] {
			t.Errorf("Mismatch in desired (%v) and output (%v) tokens\n", desiredOutput[i], scannerOutput[i])
		}
		fmt.Printf("%v %v %v %v\n", k, scannerOutputType[i], desiredOutput[i], scannerOutput[i])
	}

	if t.Failed() {
		fmt.Printf("Expected %v, obtained %v\n", desiredOutput, scannerOutput)
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
