package scanner

import (
	"testing"
)

func Test_hasNext(t *testing.T) {
	for _, k := range keywordsWithA {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithB {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithC {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithD {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithE {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithF {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithG {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithH {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithI {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithJ {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithK {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithL {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithM {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithN {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithO {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithP {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithQ {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithR {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithS {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithT {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithU {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithV {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
	for _, k := range keywordsWithW {
		input := k
		scanner := New([]rune(input))
		for scanner.HasNext() {
			scanner.Next()
		}
	}
}
