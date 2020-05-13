package main

func main() {
	panic(nil) // want `panic is disallowed inside main Package` `panic is disallowed without recover`
}
