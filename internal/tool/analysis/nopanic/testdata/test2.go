package main

func testFn1() {
	testFn2()
}

func testFn2() {
	panic("panic called") // want `panic is disallowed without recover` `panic is disallowed inside main Package`
}
