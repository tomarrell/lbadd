package main

func testFn01() {
	defer func() {
		if err := recover(); err != nil {
			panic(err) // want `panic is disallowed inside main Package`
		}
	}()
	panic("") // want `panic is disallowed inside main Package`
}
