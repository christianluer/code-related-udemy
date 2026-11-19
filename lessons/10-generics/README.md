# 10 · Generics

Course section: *Generics*. Reading: [the generics tutorial](https://go.dev/doc/tutorial/generics),
[Why generics](https://go.dev/blog/why-generics), [Intro to generics](https://go.dev/blog/intro-generics).

Generics let you write DRY code that works for many types while staying type-safe.

## Type parameters and constraints

```go
func addI(a, b int) int         { return a + b }
func addF(a, b float64) float64 { return a + b }

// one generic func instead of two
func add[T int | float64](a, b T) T { return a + b }

add(1, 2)          // T inferred as int
add[float64](1, 2) // explicit
```

## Type set interfaces

An interface can describe a **set of types** to use as a constraint:

```go
type myNumbers interface {
	~int | ~float64   // ~ means "any type whose UNDERLYING type is int or float64"
}
type myAlias int      // underlying type int, so it satisfies ~int (but not int)
func sum[T myNumbers](xs ...T) T { ... }
```

Type-set interfaces (those containing `|` or `~T`) can only be used as constraints, not as regular variable types.

## Built-in and standard constraints

- `any`: every type.
- `comparable`: types that support `==` and `!=` (map keys).
- `cmp.Ordered` (standard library): types that support `< <= > >=`. Previously this lived in
  `golang.org/x/exp/constraints`.
- The `slices` and `maps` packages are built on generics: `slices.Sort`, `slices.Contains`,
  `slices.Index`, `maps.Keys`...

## Generic types

```go
type Stack[T any] struct{ items []T }
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
var s Stack[string]
```

## Concrete type vs interface type

A **concrete type** (int, string, a struct...) has a specific memory layout and you can create values
of it directly. An **interface type** describes behavior, and you can't create a value of it directly,
only values of concrete types that satisfy it.

Two kinds of polymorphism: **runtime** polymorphism through interfaces (`any`, `io.Reader`), and
**compile-time** polymorphism through generics. Generics can occasionally be slower than hand-written
code, so measure before optimizing.

## Exercises

| Exercise | Topic |
|---|---|
| `Sum[T Number]` | type set constraint with `~` |
| `Max[T cmp.Ordered]` | standard constraint |
| `Contains[T comparable]` | comparable |
| `Map[T, U any]` | two type params |
| `Filter[T any]` | callbacks + generics |
| `Stack[T]` Push/Pop/Len | generic type with methods |
