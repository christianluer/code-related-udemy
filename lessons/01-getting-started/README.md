# 01 · Getting going with Go

Course sections: *Course introduction*, *Getting going with Go*, hands-on exercise #1.

## Why Go?

- Built at Google (2006, open-sourced Nov 2009, v1.0 March 2012) by **Rob Pike** (Unix, UTF-8),
  **Robert Griesemer** and **Ken Thompson** (Unix, B, helped create C).
- Designed for the multi-core era: the first dual-core CPUs shipped in 2005.
- Goals: **efficient compilation, efficient execution, ease of programming**.
- Compiles to a single **static executable**. There is a garbage collector, but no virtual machine.
- Cross-compiles to other operating systems (`GOOS=linux go build`).
- Great for web services at scale, networking (http/tcp/udp), concurrency, systems, CLIs.
- Statically typed: a VALUE of a certain TYPE. Backward compatibility promise from v1 onward.
- It is "Go", but google for "golang". The mascot is the Gopher ʕ◔ϖ◔ʔ.

**Documentation worth bookmarking:** [go.dev/doc](https://go.dev/doc), the
[spec](https://go.dev/ref/spec), [Effective Go](https://go.dev/doc/effective_go),
[the standard library](https://pkg.go.dev/std), [Go by Example](https://gobyexample.com),
[A Tour of Go](https://go.dev/tour), [Go Proverbs](https://go-proverbs.github.io).

## Your first program

```go
package main // an executable must be package main...

import "fmt"

func main() { // ...and have a func main: the entry point
	fmt.Println("Hello, gophers 😃!")
}
```

- Semicolons are inserted by the compiler, so you don't write them.
- `package.Identifier`: `fmt.Println` reads as "from package fmt, use the Println func".
- `func Println(a ...any) (n int, err error)`: the `...any` is a **variadic parameter**.
  It takes as many arguments of any type as you like. `any` is an alias for `interface{}`, the empty interface,
  which every value satisfies.
- You can throw away return values with the **blank identifier** `_`.

## How computers work

Computers run on electricity, which has two discrete states, ON and OFF. A **coding scheme**
gives meaning to patterns of on/off (think Morse code). One on/off switch is a **bit**, and 8 bits make a **byte**.

## ASCII, Unicode & UTF-8

- **ASCII** is an early coding scheme: 128 characters in 7 bits.
- **Unicode** assigns a number (a *code point*, written `U+0041`) to every character in every language.
- **UTF-8** encodes code points as 1 to 4 bytes. It was invented by Rob Pike and Ken Thompson. Go source code is UTF-8.
- In Go, a `string` is a read-only sequence of **bytes**, a `byte` is an alias for `uint8`, and a
  `rune` is an alias for `int32` (one code point).
- `len(s)` counts **bytes**. `utf8.RuneCountInString(s)` counts **runes**. `for i, r := range s`
  walks runes.

## String literals

```go
interpreted := "tab:\t newline:\n quote:\""   // escape sequences are processed
raw := `C:\no\escapes\here
and it can span lines`                        // backticks: what you see is what you get
```

## Exercises (`exercises/exercise.go`)

| Exercise | Course |
|---|---|
| `Greeting` returns text containing an emoji | HOE #1 |
| `RawLiteral` uses a raw string literal | HOE #1 |
| `ByteAndRuneCount` shows bytes vs runes | ASCII/Unicode/UTF-8 |
| `CodePoints` formats runes as `U+XXXX` | ASCII/Unicode/UTF-8 |
