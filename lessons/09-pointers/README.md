# 09 · Pointers

Course section: *Pointers*, hands-on exercises #74 to #77.

All values are stored in memory, and every memory location has an **address** (think post office boxes).
A **pointer** is a memory address.

```go
x := 42
p := &x          // & "address of": p is of type *int
fmt.Println(p)   // 0xc000012345 (an address)
fmt.Println(*p)  // * "dereference": 42, the value at that address
*p = 21          // change the value at the address, so x is now 21
fmt.Printf("%T %T\n", x, p) // int *int
```

`*int` in a **type** means "pointer to int". `*p` in an **expression** means "the value p points to".

## Everything is pass by value

A func always receives a **copy**. But when the copy is a pointer, or a slice/map header,
it still refers to the same underlying data:

```go
func addOne(v int)   { v++ }    // changes a copy: the caller sees nothing
func addOneP(v *int) { *v++ }   // changes the caller's value
func setFirst(s []int) { s[0] = 100 } // caller sees it: the copied header points to the same array
```

Reference-like types: pointers, slices, maps, channels, funcs, interfaces. Their zero value is `nil`.

## Value semantics vs pointer semantics

| | value semantics | pointer semantics |
|---|---|---|
| what moves | a copy of the data | a copy of the address (shared memory) |
| pros | simpler, safer, isolated bugs, fewer races, stays on the **stack** | no big copies, can mutate |
| use for | built-ins (numeric, string, bool), reference types (slices, maps...), small structs | large structs (≥ 64 bytes rule of thumb), mutation, `json.Unmarshal(&v)` |

Be **consistent**: once a type has one pointer-receiver method, give all its methods pointer receivers.
"The majority of bugs we get in software have to do with the mutation of memory" (Bill Kennedy).

## Stack, heap & escape analysis

Values that can't be proven to stay inside a function **escape to the heap**: for example a returned
pointer to a local, or a variable captured by a closure. Heap memory is garbage collected and slower to allocate.

```bash
go build -gcflags=-m ./lessons/09-pointers/examples/escape   # "moved to heap: x"
```

It's fine to return `&x` from a func in Go. The compiler moves `x` to the heap for you.

## Method sets

| receiver | the method is in the method set of |
|---|---|
| `func (t T) M()` | `T` **and** `*T` |
| `func (t *T) M()` | `*T` only |

So if an interface needs a method with a pointer receiver, only a `*T` satisfies it.
Calling `v.M()` on an **addressable** value `v` still works because Go writes `(&v).M()` for you,
but `T{}.M()` does not (a literal isn't addressable).

## Exercises

| Exercise | Course |
|---|---|
| `AddOne`, `Swap` dereference | HOE #74, #75 |
| `ChangeFirstValue` / `ChangeFirstPointer` | HOE #77 |
| `Kid` methods + `NewYoungin` method sets | HOE #76 |
| `NewInt` returns a pointer to a local (escapes) | stack & heap |
| `DoubleAll` mutates through a slice | pass by value & mutability |
