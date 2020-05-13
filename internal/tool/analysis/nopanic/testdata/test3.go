package main

func testFn0() {
	testFn1()
}

func testFn1() int {
	m := 1
	panic("foo: fail") // want `panic is disallowed without recover` `panic is disallowed inside main Package`
	m = 2
	return m
}
