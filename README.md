# code-related-udemy · Learn To Code Go

An interactive, hands-on companion to Todd McLeod's **Learn To Code Go** course
(`course_outline.pdf`). Every section of the course is a lesson with notes,
runnable examples, test-driven exercises and a quiz — all driven by a small
terminal tutor.

```
ʕ◔ϖ◔ʔ   Learn To Code Go
```

## Quick start

Requires Go 1.24+.

```sh
go run ./cmd/learn          # interactive shell (type "help")
go run ./cmd/learn next     # what should I do now?
```

The typical loop for a lesson:

| Step | Command |
|------|---------|
| Read the notes | `read 4` |
| Run the examples | `run 4` (or `run 4 loops`) |
| Edit the exercises | `lessons/04-control-flow/exercises/exercise.go` |
| Re-test on every save | `watch 4` |
| Stuck? | `hint 4`, then `solution 4` as a last resort |
| Test yourself | `quiz 4` |

Lessons can be referenced by number or by part of their name (`check maps`).
If you omit the lesson, the first unfinished one is used.

### All commands

| Command | What it does |
|---------|--------------|
| `list` | lessons with progress bars and quiz scores |
| `next` | suggests the next lesson and the steps for it |
| `read N` | prints the lesson notes |
| `examples N` / `run N [name]` | lists / runs example programs |
| `check N` | runs the exercise tests (with the race detector) and records progress |
| `watch N` | re-runs `check` every time you save a `.go` file |
| `hint N` | hints for the exercises you haven't solved yet |
| `solution N` | shows the reference solution (asks first) |
| `quiz N` | multiple choice quiz; best score is saved |
| `verify` | proves every reference solution passes |
| `reset` | wipes your progress |

Progress lives in `.learn-progress.json` (git-ignored). Set `NO_COLOR=1` to
disable colours.

## Curriculum

| # | Lesson | Course sections covered |
|---|--------|-------------------------|
| 1 | Getting going with Go | Course introduction, Getting going with Go (strings, UTF-8, runes, bytes) |
| 2 | The fundamentals of Go | Variables, zero values, `fmt`, constants, `iota`, numeric types, bit shifting |
| 3 | Tooling, modules & packages | Go toolchain, scope, exported identifiers, modules, versioning, checksums |
| 4 | Control flow | `if`, `switch`, `for`, `range`, `select`, comma-ok idiom |
| 5 | Arrays & slices | Arrays, slicing, `append`, deleting, multidimensional slices, internals |
| 6 | Maps | Maps, comma-ok, delete, range, word frequency |
| 7 | Structs | Structs, embedding, anonymous structs, composition |
| 8 | Functions, methods & interfaces | Variadic funcs, defer, methods, interfaces, polymorphism, `io.Writer`, stringer, closures, callbacks, recursion, mocking |
| 9 | Pointers | Addresses, dereferencing, value vs pointer semantics, method sets, escape analysis |
| 10 | Generics | Type parameters, constraints, type sets, generic types |
| 11 | JSON, sort & bcrypt | `encoding/json`, `sort`, custom sorting, `bcrypt` |
| 12 | Concurrency | Goroutines, `WaitGroup`, method sets revisited, race conditions, mutex, atomic |
| 13 | Channels & context | Unbuffered/buffered channels, directional channels, range/close, select, comma-ok, fan-in, fan-out, `context` |
| 14 | Error handling | Errors as values, checking, printing/logging, `panic`/`recover`, custom errors, wrapping, `errors.Is/As/Join` |
| 15 | Docs, testing & benchmarking | Writing docs, `go doc`, examples, table tests, coverage, benchmarks |

## How the repo is laid out

```
cmd/learn/                 the terminal tutor (standard library only)
lessons/NN-topic/
  README.md                lesson notes
  lesson.json              title, hints and quiz
  examples/<name>/main.go  runnable programs (go run ./lessons/NN-topic/examples/<name>)
  exercises/
    exercise.go            YOUR code: TODO stubs          (//go:build !solution)
    solution.go            reference answers               (//go:build solution)
    exercise_test.go       the tests both are checked against
```

`exercise.go` and `solution.go` declare the same functions, and build tags
make sure only one of them is compiled at a time. That means you can use the
plain Go tools too:

```sh
go test ./lessons/04-control-flow/exercises/                 # your code
go test -tags solution ./lessons/04-control-flow/exercises/  # the reference
```

The stubs compile but every test fails until you implement them. The
concurrency lesson's counters deliberately contain data races, so they only
pass once the race detector is happy.

## Make targets

```sh
make learn     # start the tutor
make check N=4 # check one lesson
make verify    # all reference solutions pass (also run in CI)
make lint      # gofmt + go vet for both build variants
```
