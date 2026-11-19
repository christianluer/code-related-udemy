// Generics: type constraints, type sets, ~underlying types, cmp/slices/maps.
package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// Without generics: one func per type.
func addI(a, b int) int         { return a + b }
func addF(a, b float64) float64 { return a + b }

// Type constraint inline.
func addT[T int | float64](a, b T) T { return a + b }

// Type set interface with ~ (underlying types).
type myNumbers interface {
	~int | ~float64
}

func addN[T myNumbers](a, b T) T { return a + b }

type myAlias int

// cmp.Ordered: anything that supports <.
func maxOf[T cmp.Ordered](xs ...T) T {
	m := xs[0]
	for _, x := range xs[1:] {
		if x > m {
			m = x
		}
	}
	return m
}

// Generic type.
type Pair[K comparable, V any] struct {
	Key K
	Val V
}

func (p Pair[K, V]) String() string { return fmt.Sprintf("%v=%v", p.Key, p.Val) }

func main() {
	fmt.Println(addI(1, 2), addF(1.5, 2.5))
	fmt.Println(addT(1, 2), addT(1.5, 2.5))

	var n myAlias = 42
	fmt.Println(addN(n, 1), "<- myAlias works thanks to ~int")
	// addT(n, 1) // compile error: myAlias does not satisfy int | float64

	fmt.Println(maxOf(3, 9, 2), maxOf("apple", "pear", "fig"), maxOf(1.5, -2.0))

	p := Pair[string, int]{"answer", 42}
	fmt.Println(p)

	// The standard library is full of generic helpers.
	xs := []int{5, 2, 8, 1}
	slices.Sort(xs)
	fmt.Println(xs, slices.Contains(xs, 8), slices.Index(xs, 5), slices.Max(xs))

	m := map[string]int{"b": 2, "a": 1, "c": 3}
	fmt.Println(slices.Sorted(maps.Keys(m)))
}
