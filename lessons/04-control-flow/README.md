# 04 · Control flow

Course section: *Control Flow*, hands-on exercises #23 to #39.

Control flow is the order in which statements execute. It comes in three forms: **sequence**, **conditional**, and **loop**.
Program execution starts by initializing package `main` (including every `func init()`), then
calls `func main()`. When main returns, the program exits and does **not** wait for other goroutines.

## Under the hood: runtime, stack, heap, compiler

- The **Go runtime** schedules goroutines, collects garbage, allocates memory, manages channels, and
  gives every goroutine its own growable **stack**.
- The **stack** is LIFO, per function call, and fast. The **heap** holds longer-lived values and is managed by the GC.
- The **compiler** (`gc`) parses, type-checks, optimizes, and generates machine code.

## if, comparison & logical operators

```go
if x < 10 {
} else if x == 10 {
} else {
}
```

Comparison operators: `==  !=  <  <=  >  >=`. Logical operators: `&&` AND, `||` OR, `!` NOT.
Both `&&` and `||` short-circuit.

## The "statement; statement" and "comma ok" idioms

```go
if z := 2 * rand.Intn(5); z >= 4 { ... }  // z only exists inside the if/else: narrow scope
if age, ok := m["Q"]; !ok { fmt.Println("Q is not in the map") }
```

## switch

```go
switch {                    // no expression: like an if/else chain
case x < 0:
	fmt.Println("negative")
case x == 0, x == 1:        // several values in one case
	fmt.Println("tiny")
	fallthrough             // run the next case body too (rare!)
default:
	fmt.Println("other")
}
switch day { case "sat", "sun": ... }  // with an expression
```

Cases don't fall through by default, so there's no `break` needed.

## select (preview of concurrency)

`select` is a switch over channel operations. It waits until one case can proceed, and picks
randomly if several are ready. More in lesson 13.

## for: the only loop keyword

```go
for i := 0; i < 10; i++ { }   // init; condition; post
for x < 100 { }               // condition only (a "while" loop)
for { if done { break } }     // infinite loop: use break
for i := range 10 { }         // range over an int (Go 1.22+)
for i, v := range slice { }   // range over slice, array, string, map, channel
```

`break` leaves the loop and `continue` skips to the next iteration. Nested loops run the inner
loop completely for every iteration of the outer one.

## Modulus

`%` gives the remainder: `11 % 3 == 2`. `x % 2 == 0` means even.

## Random numbers

`math/rand/v2` gives `rand.IntN(n)`, which returns `[0, n)`: it includes 0 and excludes n.
In the older `math/rand`, the same function is `rand.Intn`.

## init

`func init()` is **niladic** (it takes no arguments and returns nothing). It runs before main. A file can have
several init funcs.

## Exercises

| Exercise | Course |
|---|---|
| `Bucket` if/else ranges | HOE #24, #25 |
| `DescribePair` logical operators | HOE #27, #28 |
| `Odds` continue + modulus | HOE #33 |
| `FirstMultiple` break | HOE #31, #32 |
| `Grid` nested loops | HOE #34 |
| `IndexValues` for range | HOE #35 |
| `AgeOf` comma ok idiom | HOE #36, #37 |
