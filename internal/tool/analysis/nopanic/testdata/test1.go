package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f")
			panic("throwing another error") // want `panic is disallowed inside main Package` `panic is not allowed inside recover`
		}
	}()
	panic("error new changes") // want `panic is disallowed inside main Package`
	testFn1()
	testFn3()
	testFn4()
}
