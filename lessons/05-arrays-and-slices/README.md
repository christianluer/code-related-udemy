# 05 · Grouping data: arrays & slices

Course section: *Grouping data values - array & slice*, hands-on exercises #40 to #48.

## Four ways to group values

| | what | size |
|---|---|---|
| **array** | numbered sequence of values of the same type | fixed, and part of the type: `[5]int` ≠ `[6]int` |
| **slice** | built on top of an array; values of the same type | grows; has a **len** and a **cap** |
| **map** | key/value store, unordered | grows (lesson 06) |
| **struct** | values of *different* types together | fixed (lesson 07) |

Arrays are mostly a building block for Go's internals. **Use slices.**

## Arrays

```go
var a [3]int                    // [0 0 0]
b := [3]string{"a", "b", "c"}   // array literal
c := [...]int{1, 2, 3, 4}       // the compiler counts: [4]int
fmt.Println(len(c))
```

Arrays are **values**: assigning one or passing it to a func copies every element.

## Slices

```go
xs := []int{42, 43, 44}           // slice literal (composite literal)
for i, v := range xs { }          // range
xs[0]                             // access by index
xs = append(xs, 45, 46)           // append returns a (possibly new) slice: always reassign
xs = append(xs, other...)         // append a whole slice by "unfurling" it
xs[1:3]                           // slicing: from index 1 up to, not including, 3
xs[:2], xs[2:], xs[:]
xs = append(xs[:2], xs[4:]...)    // delete elements 2 and 3
s := make([]int, 0, 50)           // make([]T, len, cap): pre-size the underlying array
grid := [][]string{{"a", "b"}, {"c"}} // multidimensional: a slice of slices
```

## Slice internals

A slice is a tiny struct with three fields: a **pointer** to an underlying array, a **len**, and a **cap**.

- Slicing does **not** copy. The new slice shares the underlying array, so writes show through.
- `append` writes into the existing array if there is capacity. When there isn't, it allocates a new,
  bigger array (roughly doubling) and copies.
- Passing a slice to a func copies the 3-field header, not the data, so the func can modify elements.
- Want an independent copy? Use `copy(dst, src)` or `slices.Clone(src)`.
- If you know the final size, use `make([]T, 0, n)` so the underlying array is created only once.

## Exercises

| Exercise | Course |
|---|---|
| `Squares` array + range | HOE #40, #42 |
| `SliceUp` slicing | HOE #43, #44 |
| `Extend` append & unfurling | HOE #45 |
| `Delete` delete with append | HOE #46 |
| `CleanStates` make + append without reallocating | HOE #47 |
| `Flatten` [][]string | HOE #48 |
| `Clone` independent copy | slice internals |
