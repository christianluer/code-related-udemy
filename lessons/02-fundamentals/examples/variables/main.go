// Variables, zero values, the blank identifier, conversion and constants.
package main

import (
	"fmt"
	"time"
)

// Package-level declarations must use var / const, never :=.
var packageLevel = "visible to every file in this package"

const typed int = 7
const untyped = 7 // takes on whatever type the context needs

func main() {
	// Zero values.
	var b bool
	var i int
	var f float64
	var s string
	var p *int
	var sl []int
	var m map[string]int
	fmt.Printf("bool=%v int=%v float64=%v string=%q pointer=%v slice=%v(nil? %v) map=%v\n",
		b, i, f, s, p, sl, sl == nil, m)

	// Short declaration and multiple initialization.
	x, y, name := 42, 3.14, "gopher"
	fmt.Printf("%v %T | %v %T | %v %T\n", x, x, y, y, name, name)

	// Specificity: var when you need a type other than the default.
	var small int8 = 100
	var precise float32 = 1.5
	fmt.Printf("%v %T, %v %T\n", small, small, precise, precise)

	// Blank identifier: discard values you don't need.
	_, err := fmt.Sscan("12", &i)
	fmt.Println("scanned:", i, "err:", err)

	// Conversion, not casting.
	sum := 10
	count := 4
	avg := float64(sum) / float64(count)
	fmt.Println("average:", avg, "integer division:", sum/count)

	// Untyped constants are flexible; typed ones are not.
	var asFloat float64 = untyped // ok
	// var nope float64 = typed   // compile error: cannot use typed (int) as float64
	fmt.Println(asFloat, typed, 2*time.Second, untyped*time.Millisecond)

	fmt.Println(packageLevel)
}
