# 02 · The fundamentals of Go

Course section: *The fundamentals of Go*, hands-on exercises #2 to #13.

## Variables: declare, assign, initialize

```go
var x int        // declare: holds a VALUE of TYPE int, initialized to its zero value
x = 42           // assign
var y int = 43   // initialize: declare + assign
z := 44          // short declaration operator: only inside funcs, type is inferred
a, b := 1, "two" // multiple initialization
var f float64 = 3 // use var when you need a specific type
```

General guideline: use `var` for the zero value, `:=` most of the time, and `var x T = v` when
you need a specific type. Go is **statically typed**: a variable holds a VALUE of one TYPE forever.
Unused variables and imports are **compile errors** ("code pollution"). Use `_` to discard values.

## Zero values

| type | zero value |
|---|---|
| bool | `false` |
| ints | `0` |
| floats | `0.0` |
| string | `""` |
| pointers, funcs, interfaces, slices, channels, maps | `nil` |

## fmt.Printf verbs

| verb | shows | `42` prints |
|---|---|---|
| `%v` | value | `42` |
| `%T` | type | `int` |
| `%d` | decimal | `42` |
| `%b` | binary | `101010` |
| `%x` / `%#x` | hex | `2a` / `0x2a` |
| `%o` / `%O` | octal | `52` / `0o52` |
| `%q` | quoted | `'*'` |
| `%U` | unicode | `U+002A` |
| `%08b` | padded | `00101010` |

## Numeral systems

- decimal is base 10 (0-9), binary is base 2 (0-1), hexadecimal is base 16 (0-9, A-F).
- `911` is `0b1110001111` in binary and `0x38F` in hex (the Porsche 911 license plate).

## Values, types, conversion

- A **conversion** `T(x)` changes the type of an expression: `float64(i)`. Go has
  conversion, **not casting**, and never mixes numeric types implicitly.
- **Constants** are created at compile time and can be typed (`const x int = 1`) or untyped
  (`const x = 1`). An untyped constant takes the type its context needs, so `1e6 * time.Second` is legal.
- **Scope** is lexical, using blocks: universe → package → file (imports) → function → block.
  Keep scope as narrow as possible.

## Built-in, aggregate, composite types

- **Basic/built-in**: bool, numeric (int, int8..64, uint8..64, float32/64, complex, byte, rune), string.
- **Aggregate**: array, slice, struct (many values together).
- **Composite**: struct, which composes values of different types.
- Use `int` unless you have a reason not to. It is 64 bits on 64-bit systems.

## iota & bitwise ops

```go
const (
	_  = iota             // 0, thrown away
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // the expression repeats with iota = 2: 1 << 20
)
```

Bitwise operators: `&` AND, `|` OR, `^` XOR, `&^` AND NOT, `<<` left shift, `>>` right shift.
`x << n` multiplies by 2ⁿ.

## Signed vs unsigned

`int8` ranges from -128 to 127 and `uint8` (a byte) from 0 to 255. See the `math` package for
`math.MaxInt8` etc. Overflowing a variable wraps around.

## Exercises

| Exercise | Course |
|---|---|
| `Describe` uses `%v` and `%T` | HOE #11 |
| `Formats` prints decimal, binary, hex | HOE #12 |
| `KB`..`EB` constants with iota | HOE #8, #9 |
| `Sunday`..`Saturday` with iota | HOE #8 |
| `Limits` max int8 / uint8 | HOE #13 |
| `Average` needs a conversion | Values, types, conversion |
| `OnesCount` bitwise ops | HOE #7 |
