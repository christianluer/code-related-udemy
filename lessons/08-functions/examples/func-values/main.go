// Funcs as values: anonymous funcs, func expressions, returning funcs,
// callbacks, closures, recursion and wrappers.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Anonymous self-executing funcs.
	func() { fmt.Println("anonymous func ran") }()
	func(s string) { fmt.Println("anonymous func with param:", s) }("hi")

	// Func expression: assign a func to a variable.
	multiply := func(a, b int) int { return a * b }
	fmt.Printf("multiply(5, 10) = %d, type %T\n", multiply(5, 10), multiply)

	// Returning a func.
	f := fortyTwo()
	fmt.Println("returned func gives", f())

	// Callbacks.
	fmt.Println("doMath(42, 16, add) =", doMath(42, 16, add))
	fmt.Println("doMath(42, 16, subtract) =", doMath(42, 16, subtract))
	fmt.Println("doMath with a literal =", doMath(3, 4, func(a, b int) int { return a * b }))
	fmt.Println("evens:", filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 }))

	// Closure: each counter has its own enclosed n.
	a, b := incrementor(), incrementor()
	fmt.Println("a:", a(), a(), a(), "b:", b())

	// Recursion.
	fmt.Println("factorial(5) =", factorial(5))

	// Wrappers.
	timed(func() { time.Sleep(30 * time.Millisecond) })
	logged("greet", func() { fmt.Println("  Hello, World!") })
}

func fortyTwo() func() int { return func() int { return 42 } }

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }

func doMath(a, b int, op func(int, int) int) int { return op(a, b) }

func filter(xs []int, keep func(int) bool) []int {
	var out []int
	for _, x := range xs {
		if keep(x) {
			out = append(out, x)
		}
	}
	return out
}

func incrementor() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func timed(fn func()) {
	start := time.Now()
	fn()
	fmt.Println("elapsed:", time.Since(start).Round(time.Millisecond))
}

func logged(name string, fn func()) {
	fmt.Println("starting", name)
	fn()
	fmt.Println("finished", name)
}
