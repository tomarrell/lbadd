package compiler

// Option is a functional option that can be applied to a compiler. If the
// option is applicable to the compiler, is determined by the compiler itself.
type Option func(*simpleCompiler)
