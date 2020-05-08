package main

import "context"

type Foo interface {
	MyFunc1(context.Context)   // valid
	MyFunc2(_ context.Context) // want `unused context.Context argument`
}
