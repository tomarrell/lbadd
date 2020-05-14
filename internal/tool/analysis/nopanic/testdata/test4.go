package main

import "fmt"

func testFn4() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f")
			panic("throwing another error") // want `panic is disallowed inside main Package` `panic is not allowed inside recover`
		}
	}()
	g(0)
}

func g(i int) {
	if i > 3 {
		panic("error panic") // want `panic is disallowed inside main Package`
	}
	g(i + 1)
}
