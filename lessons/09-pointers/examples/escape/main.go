// Escape analysis. Build with:
//
//	go build -gcflags=-m ./lessons/09-pointers/examples/escape
//
// and look for "moved to heap" and "does not escape".
package main

import "fmt"

type big struct{ data [1024]int }

// stays on the stack: the value is copied out
func valueResult() big {
	var b big
	b.data[0] = 1
	return b
}

// escapes: we return the address of a local
func pointerResult() *big {
	var b big
	b.data[0] = 2
	return &b
}

// escapes: the closure captures n and outlives the call
func counter() func() int {
	n := 0
	return func() int { n++; return n }
}

func main() {
	v := valueResult()
	p := pointerResult()
	c := counter()
	c()
	fmt.Println(v.data[0], p.data[0], c())
}
