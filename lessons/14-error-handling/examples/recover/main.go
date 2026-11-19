// defer, panic and recover (from the "Defer, Panic, and Recover" blog post).
package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")

	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(1, 0))
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // never runs
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}
