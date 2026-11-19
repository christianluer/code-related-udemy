// Scope: universe, package, file, function and block scope, plus shadowing.
package main

import "fmt" // file scope: only this file sees the fmt import

// package scope: visible in every file of package main
var x = "package-level x"

const planet = "Earth"

func main() {
	fmt.Println(x, "|", planet)

	x := "function-level x" // shadows the package-level x
	fmt.Println(x)

	{
		x := "block-level x" // shadows again, only inside these braces
		y := 42
		fmt.Println(x, y)
	}
	// fmt.Println(y) // compile error: y is out of scope here

	fmt.Println(x)

	// The "statement; statement" idiom keeps a variable's scope tiny.
	if z := len(x); z > 10 {
		fmt.Println("z only exists in this if/else:", z)
	}

	fmt.Println(show())
}

func show() string {
	return "inside show(), x is still: " + x // package scope, not main's x
}
