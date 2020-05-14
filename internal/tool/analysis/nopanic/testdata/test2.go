package main

func testFn1() {
	testFn2()
}

func testFn2() {
	panic(nil) // want `panic is disallowed without recover` `panic is disallowed inside main Package` `panic is not allowed without error`
}
