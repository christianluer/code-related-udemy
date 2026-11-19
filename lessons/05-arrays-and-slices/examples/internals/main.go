// Slice internals: pointer + len + cap, and the shared underlying array.
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	b := a[1:3] // shares a's underlying array

	fmt.Printf("a=%v len=%d cap=%d\n", a, len(a), cap(a))
	fmt.Printf("b=%v len=%d cap=%d (cap runs to the end of a's array)\n", b, len(b), cap(b))

	b[0] = 100
	fmt.Println("after b[0] = 100, a =", a, "(a changed too!)")

	b = append(b, 200) // still has capacity, so it overwrites a[3]
	fmt.Println("after append(b, 200), a =", a)

	// Use make + copy for an independent slice with a new underlying array.
	c := make([]int, len(a))
	copy(c, a)
	c[0] = -1
	fmt.Println("copy c =", c, "a is unchanged =", a)

	// Passing a slice to a func copies the header, not the data.
	double(a)
	fmt.Println("after double(a):", a)

	// But append inside a func doesn't change the caller's len.
	grow(a)
	fmt.Println("after grow(a):  ", a, "(len unchanged)")
}

func double(xs []int) {
	for i := range xs {
		xs[i] *= 2
	}
}

func grow(xs []int) {
	xs = append(xs, 999)
	_ = xs
}
