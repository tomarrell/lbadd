package main

import "context"

func main() {
	// this is so that staticcheck does not complain about the unused functions
	_ = testFnB0
	_ = testFnB1
	_ = testFn0
	_ = testFn1
	_ = testFn2
	_ = testFn3
	_ = testFn4
	_ = testFn5
	_ = testFn6
	_ = testFn7
	_ = testFn8
}

// functions without body

func testFn0(ctx context.Context, s string, i int)                // valid
func testFn1(ctx, ctx2 context.Context)                           // want `more than one context.Context argument`
func testFn2(ctx context.Context, ctx2 context.Context)           // want `more than one context.Context argument`
func testFn3(ctx1 context.Context, ctx2 context.Context)          // want `context.Context argument should be named 'ctx'`
func testFn4(ctx1 context.Context)                                // want `context.Context argument should be named 'ctx'`
func testFn5(s string, ctx context.Context)                       // want `context.Context should be the first parameter of a function`
func testFn6(s string, ctx1 context.Context)                      // want `context.Context should be the first parameter of a function`
func testFn7(ctx context.Context, s string, ctx2 context.Context) // want `more than one context.Context argument`
func testFn8(_ context.Context, s string)                         // want `unused context.Context argument`

// functions with body

func testFnB0(ctx context.Context, s string, i int) {} // valid
func testFnB1(_ context.Context)                    {} // want `unused context.Context argument`
