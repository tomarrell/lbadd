package main

import "fmt"

func testFn3() {
	if err := recover(); err != nil { // want `recover is disallowed without defer`
		fmt.Printf("panic recovered")
	}
	panic("throw error") // want `panic is disallowed inside main Package`
}
