// Hello gophers: package main, func main, the fmt package and the blank identifier.
package main

import "fmt"

func main() {
	fmt.Println("Hello, gophers 😃!")

	// Println is variadic: it takes any number of arguments of any type.
	fmt.Println("one", 2, 3.0, true, 'x')

	// Println returns the number of bytes written and an error.
	n, err := fmt.Println("count my bytes")
	fmt.Println("bytes written:", n, "error:", err)

	// Use the blank identifier to throw away values you don't need.
	_, _ = fmt.Println("I ignore both return values")

	// Printf formats with "verbs": %v value, %T type, %q quoted, %d decimal.
	fmt.Printf("%v is of type %T\n", 42, 42)
	fmt.Printf("%q is of type %T\n", "gopher", "gopher")
}
