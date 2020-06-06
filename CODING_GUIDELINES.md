# Coding Guidelines
This document describes our coding guidelines.
One of the main points of this project is, to keep code readable.
This is why we enforce the following rules.
If you are a contributor to this project, please follow these rules.
If you disagree with one or more of them, please open an issue and provide solid description why we should change them.
Many of these rules are enforced by static code analysis, which you can run locally with `make lint`.

## Files
When creating a new package `mypackage`, we aim for the following files in the folder.
* `doc.go` only consists of a package-documentation comment and a package declaration.
* `example_test.go` contains example functions that show how to use your package.
  Use of `package mypackage_test` is mandatory here. Be aware that example functions are testable in Go.
* `mypackage.go` contains the entry point of this package
  For ease of readability, no other file should contain an entry point to your package (except structs).
* `mypackage_test.go` contains tests for your package.
  If you like, you can use the `package mypackage_test` in that file, however, this is up to you.
* `error.go` contains sentinel errors.
  For possible content, see section _Sentinel errors_ below.
* Any other file in your package that is necessary.
  If you use a lot of structs with methods, try to place each struct in a separate file, that is named like the struct, but in lower-snake-case (`hello_world`).

## Code
We want to keep our code readable (that's one of the main points of this project) and clean.
Many of the things we expect from code, is ensured by our internal analyzers (`internal/tool/analysis`).
However, not everything may be covered by those.
This is why we rely on you to follow the below rules for our code.

### Ordering
We want to keep the following structure in any non-test `.go`-file.

1. package declaration (no package comment, use a file `doc.go` for that)
2. import declarations
3. const declarations
4. var declarations
5. interface declarations
6. struct declarations (followed by constructor functions)
7. exported methods
8. unexported methods
9. exported functions
10. unexported functions

If there are multiple independent structs in one file, you should probably split the file into two or more.

### Constructor functions
Use the `&struct` syntax for initialization whenever possible.
If you need to do previous initialization, do this before inline-initializing the struct in the return statement.
The example below describes a correct example of how to write a constructor.

```go
type Foo struct {
    id   [16]byte
    x    int
    y    int
    desc string
}
func NewFoo(desc string) *Foo {
    id := createID()
    return &Foo{
        id:   id,
        x:    0,
        y:    0,
        desc: desc
    }
}

```

### `context.Context`
If not in a performance critical layer, use `context.Context` and respect it.
Especially when writing network APIs, `context.Context` is a must.
When you use it, there are a few rules.

* a `context.Context` must always be the first argument in a method or function
* there must be only one `context.Context` argument
* the `context.Context` argument has to be named `ctx`
* the context has to be respected (implementation must abort as soon as possible when `<-ctx.Done()` returns)

### Sentinel errors
Sentinel errors are exported error values that you use for comparison.
```go
res, err := somepkg.Operation()
if err == somepkg.ErrTimeout {
    // timeout, maybe try again?
} else if err != nil {
    // some other error, handle accordingly
}
```
If you make use of sentinel errors in your package, **sentinel errors must be constant**.
Go itself doesn't follow this rule.
As an example, `io.EOF` is not a constant error.
Run a larger application after calling `io.EOF = nil` and see what happens.
This must not be possible in this codebase.
To do this, use the following snippet.

```go
// error.go
package mypackage

type Error string
func (e Error) Error() string { return string(e) } // implement error interface

const (
    ErrMyError Error = "my error"
    ErrTimeout Error = "timeout"
)
```

This way, an API user can easily check against `mypackage.ErrTimeout`, or even check if the error came from your package `myErr, ok := err.(mypackage.Error)`.
The latter is also, why we encourage this snippet of code in every package that uses sentinel errors.

### Forbidden functions
There are a few function calls that endanger the stability of the whole application.
If and only if there is a very good reason and that reason is documented sufficiently, then one of the following functions may be used.
The forbidden functions are generally functions, that are known to `panic`.
If a function documents that it `panic`s, if certain conditions are not met, **you must not use that function**.
A non-exhaustive list of these functions includes the following.

* builtin `panic` (exceptions: see below)
* `"log".log.Fatal{,f,ln}`
* `"log".log.Panic{,f,ln}`
* any other logging framework's `Fatal` or `Panic` functions (we only use `rs/zerolog`, but nevertheless)
* `os.Exit`
* `runtime.Goexit`

#### `panic`
There exists one circumstance, in which we allow the use of `panic` without documentation.
`panic` may be used to propagate an `error` inside a package.
This implies, that every path into this package has to be guarded by a deferred `recover()` call, and that deferred call must not `panic` himself (re-panic).
A typical use-case would be a locally fatal error in a parser, where you don't want your 2000 production rules to return 2 errors.
Nevertheless, don't use this unless necessary.
It is not necessary for packages with only a few functions.

## Comments
Of course we want comments.
However, make sure that they are wrapped neatly, and don't cause crazy long lines.
If you use VSCode, use the extension `Rewrap` with default settings, and hit `Alt/Option + Q` inside a comment to wrap if necessary.

## Committing changes
Whatever you work on, create a new branch for it.
If the task is long running, the branch should have a meaningful name, otherwise, something like `Username-patch-1` is sufficient.
After finishing your work, create a PR.
If you want early reviews and feedback, create a draft PR.
See the codeowners file, to see what reviewers make sense.

Before committing, run `make lint test` on your system.
If that fails, the CI build will probably fail as well.
If `make lint test` doesn't fail on your machine but in CI, that's even worse.
If you like, create a pre-commit hook, but don't commit that.
Be warned however, some tests may take a few seconds to complete.

## Reviews
Before anything is merged into `master`, there is at least one review approval needed.
Reviewers also have to check for violations with the rules described in this document.
Be constructive, use GitHubs suggestion feature if feasible.

# VSCode
Suggested plugins:
* `Go`
* `Rewrap`
* `vscode-icons`

Run with `gopls` latest stable version and default settings, and you are good.
If you like, use the following settings, which are used by some of us.
```json
"gopls": {
    "usePlaceholders": true,
    "deepCompletion": true,
    "staticcheck": true,
    "completeUnimported": true
},
```