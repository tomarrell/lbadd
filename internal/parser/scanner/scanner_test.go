package scanner

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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
	input := "SelEct FroM"
	var desiredOutput = []string{"SelEct", "FroM"}
	var scannerOutput []string
	scanner := New([]rune(input))
	for scanner.HasNext() {
		token := scanner.Next()
		if token != nil {
			scannerOutput = append(scannerOutput, token.Value())
		}
	}
	for i, k := range scannerOutput {
		if k != desiredOutput[i] {
			t.Fail()
		}
	}

	if t.Failed() {
		fmt.Printf("Expected %v, obtained %v\n", desiredOutput, scannerOutput)
	} else {
		fmt.Printf("Tested on: %v\n", desiredOutput)
	}
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
