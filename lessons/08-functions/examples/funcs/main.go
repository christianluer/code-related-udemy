// Function syntax, multiple returns, variadic params, unfurling and defer.
package main

import "fmt"

func main() {
	greet()
	fmt.Println(double(21))
	q, r := divmod(17, 5)
	fmt.Println("17 / 5 =", q, "remainder", r)

	fmt.Println("sum() =", sum())
	fmt.Println("sum(1, 2, 3) =", sum(1, 2, 3))
	xs := []int{4, 5, 6, 7}
	fmt.Println("sum(xs...) =", sum(xs...))

	fmt.Println("\ndefer demo:")
	deferDemo()

	fmt.Println("\nnamed result modified by defer:", triple(5))
}

func greet() { fmt.Println("hello from greet") }

func double(x int) int { return x * 2 }

func divmod(a, b int) (int, int) { return a / b, a % b }

func sum(nums ...int) int {
	fmt.Printf("  nums is %v of type %T\n", nums, nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func deferDemo() {
	defer fmt.Println("  deferred 1 (runs last)")
	defer fmt.Println("  deferred 2")
	x := 10
	defer fmt.Println("  deferred 3 sees x =", x, "(arguments are evaluated immediately)")
	x = 99
	fmt.Println("  body, x =", x)
}

func triple(n int) (result int) {
	defer func() { result *= 3 }()
	return n
}
