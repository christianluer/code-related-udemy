# 15 · Documentation, testing & benchmarking

Course sections: *Writing documentation*, *Testing & Benchmarking*, *Farewell* (idiomatic Go).

## Reading documentation

```bash
go doc fmt                  # package docs
go doc fmt.Println          # one symbol
go doc strings.Builder.WriteString
go doc -u .                 # include unexported identifiers
go doc ./lessons/15-docs-and-testing/exercises Years
```

Use [pkg.go.dev](https://pkg.go.dev) for the standard library and every public module. It replaced
godoc.org: push your module, request it on pkg.go.dev, and your docs appear. To browse docs locally
like `godoc -http=:8080`, run `go run golang.org/x/pkgsite/cmd/pkgsite@latest` and open localhost:8080.

## Writing documentation

Doc comments are plain comments **directly above** a declaration, with no blank line in between, that
**begin with the name** being documented:

```go
// Package dog converts human years to dog years.
package dog

// Years converts human years to dog years (1 human year = 7 dog years).
func Years(human int) int { return human * 7 }
```

- The first sentence of a package comment appears in package lists.
- Put long package docs in their own `doc.go` file.
- Write complete sentences. Explain *why*, not *what*.

## Tests

- Tests live in a file ending in `_test.go`, in the same package as the code.
- Test funcs have the signature `func TestXxx(t *testing.T)`.
- Report failures with `t.Error` / `t.Errorf` (keep going) or `t.Fatal` (stop this test).
- Use the "expected / got" idiom: `t.Errorf("Years(%d) = %d, want %d", in, got, want)`.

```go
func TestYears(t *testing.T) {
	tests := []struct {                 // TABLE TEST: many cases, one loop
		name  string
		in    int
		want  int
	}{
		{"zero", 0, 0},
		{"ten", 10, 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // subtests: go test -run TestYears/ten
			if got := Years(tt.in); got != tt.want {
				t.Errorf("Years(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
```

## Example tests

Examples run as tests **and** show up in the docs:

```go
func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}
```

## Benchmarks

```go
func BenchmarkYears(b *testing.B) {
	for b.Loop() {          // Go 1.24+; older code uses: for i := 0; i < b.N; i++
		Years(10)
	}
}
```

## The commands

```bash
go test ./...                       # all tests
go test -v -run TestYears .         # verbose, one test
go test -bench . -benchmem          # benchmarks (don't forget the ".")
go test -cover                      # coverage percentage
go test -coverprofile c.out && go tool cover -html=c.out   # coverage in the browser
go test -race ./...                 # race detector
gofmt -l -w .   # format          go vet ./...   # suspicious constructs
```

`golint` is deprecated; use [staticcheck](https://staticcheck.dev) or `golangci-lint`.

## Idiomatic Go

- camelCase names, short receivers, `:=` where it reads well.
- Accept interfaces, return concrete types. Avoid globals.
- `defer` to release resources. `panic` only for truly exceptional cases.
- Handle every error. Write tests with the standard `testing` package.
- Readability first ("mechanical sympathy" when performance matters). Pass by value.

## Exercises

These are the course's BET exercises (**B**enchmarks, **E**xamples, **T**ests).

| Exercise | Course |
|---|---|
| `Years` + its doc comment | Writing docs HOE #1 (dog package) |
| `Count` / `UseCount` | Testing HOE #2 (quote package) |
| `CenteredAvg` | Testing HOE #3 (mymath, table test) |
| the `Example...` funcs in the test file must pass too | Example tests |

**Your turn:** create `exercises/my_test.go` with your own table test and benchmark, then run
`go test -bench . -cover ./lessons/15-docs-and-testing/exercises`.
