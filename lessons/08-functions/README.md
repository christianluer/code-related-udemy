# 08 · Functions, methods & interfaces

Course section: *Functions in the go programming language*, hands-on exercises #57 to #73.

Functions give you **modular code**: abstraction, reusability (DRY: Don't Repeat Yourself), and
readability. **Parameters** are declared in the func signature. **Arguments** are passed when you call it.

## Syntax

```go
func (r receiver) identifier(parameters) (returns) { code }

func greet() { }                                        // no params, no returns
func double(x int) int { return x * 2 }                 // 1 param, 1 return
func divmod(a, b int) (int, int) { return a / b, a % b } // multiple returns
func named() (x int, err error) { return }              // named returns (Todd doesn't like them)
```

**Everything in Go is pass by value.** The func receives a copy of each argument.

## Variadic parameters & unfurling

```go
func sum(nums ...int) int { /* nums is a []int */ }
sum(1, 2, 3)
xs := []int{4, 5}
sum(xs...)        // unfurl a slice into a variadic param
```

## defer

A deferred call runs when the surrounding func returns: after a return statement, at the end of the
body, or while panicking. Multiple defers run **LIFO** (last in, first out). Arguments are evaluated
immediately. Use it for cleanup: `f, _ := os.Open(x); defer f.Close()`.

## Methods

A method is a func attached to a TYPE via a **receiver**:

```go
type circle struct{ radius float64 }
func (c circle) area() float64 { return math.Pi * c.radius * c.radius }   // value receiver
func (c *circle) scale(f float64) { c.radius *= f }                       // pointer receiver: can mutate
```

## Interfaces & polymorphism

An interface defines a **set of methods**. Any type that has those methods **implicitly**
implements it (there's no `implements` keyword). A value can therefore be of more than one type, and
that is polymorphism.

```go
type shape interface { area() float64 }
func info(s shape) { fmt.Println(s.area()) }  // works for circles, squares, anything with area()
```

- The **empty interface** `interface{}` (alias `any`) is implemented by every type.
- `fmt.Stringer` is `String() string`. Implement it and `fmt.Println(v)` uses it.
- `io.Writer` is `Write(p []byte) (n int, err error)`. Files, `bytes.Buffer`, HTTP responses,
  and `os.Stdout` all implement it, so one func can write to any of them.
- A **string** is immutable bytes and a **[]byte** is mutable. Convert with `[]byte(s)` and `string(b)`.
- Interfaces make **mock testing** easy: depend on an interface, then pass a fake in your tests.

## Funcs are first-class citizens

```go
func() { fmt.Println("anonymous, self-executing") }()  // anonymous func / func literal / lambda
f := func(x int) int { return x * x }                   // func expression: assign a func to a variable
func makeCounter() func() int {                         // returning a func + closure
	n := 0
	return func() int { n++; return n }                 // n is "enclosed" and lives on
}
func doMath(a, b int, op func(int, int) int) int {      // callback: pass a func as an argument
	return op(a, b)
}
```

- **Closure**: one scope enclosing another. Variables in the outer scope are reachable in the inner
  scope. Closures help limit the scope of variables.
- **Recursion**: a func that calls itself. It needs a base case: `factorial(n) = n * factorial(n-1)`.
- **Wrapper**: a func that wraps another to add behavior (logging, timing, auth, error handling).
  **Callback**: a func passed in to be called later.

## Testing preview

Test files end in `_test.go`, live in the same package, and use `func TestXxx(t *testing.T)`.
Run them with `go test`. Report failures with `t.Errorf("got %v, want %v", got, want)`. See lesson 15.

## Exercises

| Exercise | Course |
|---|---|
| `Divmod` multiple returns | HOE #58 |
| `Sum` / `SumSlice` variadic + unfurling | HOE #59 |
| `DeferOrder` defer LIFO | HOE #60 |
| `Person.Speak` method | HOE #61 |
| `Person.String` Stringer | stringer interface |
| `Square/Circle.Area`, `Info` interfaces | HOE #62 |
| `WriteGreeting` io.Writer | writer interface |
| `DoMath`, `Add`, `Subtract` | HOE #64 |
| `Greet` mock-testing a database | HOE #67 |
| `FortyTwo` returning a func | HOE #69, #70 |
| `Squared` / `PrintSquare` callback | HOE #71 |
| `Counter` closure | HOE #72 |
| `Factorial` recursion | Recursion |
| `Timed` wrapper | HOE #73 |
