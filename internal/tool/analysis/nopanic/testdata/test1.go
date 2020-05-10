package main

func main() {
	panic(nil) // want `panic is disallowed without recover`
}
