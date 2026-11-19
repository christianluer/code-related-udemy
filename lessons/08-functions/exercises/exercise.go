//go:build !solution

// Package exercises contains the exercises for lesson 08.
//
// Replace each TODO, then run:  go run ./cmd/learn check 8
package exercises

import (
	"io"
	"time"
)

// Divmod returns the quotient and the remainder of a / b. (HOE #58)
func Divmod(a, b int) (int, int) {
	// TODO
	return 0, 0
}

// Sum returns the sum of any number of ints. (HOE #59)
func Sum(nums ...int) int {
	// TODO
	return 0
}

// SumSlice returns the sum of xi by UNFURLING it into Sum. (HOE #59)
func SumSlice(xi []int) int {
	// TODO: call Sum
	return 0
}

// DeferOrder must defer three funcs that append "one", "two" and "three"
// to log (in that order), then append "body" and return.
// The caller must receive []string{"body", "three", "two", "one"}. (HOE #60)
func DeferOrder() (log []string) {
	// TODO: deferred funcs can modify a named result
	return nil
}

// Speak returns "I am <First> and I am <Age> years old". (HOE #61)
func (p Person) Speak() string {
	// TODO
	return ""
}

// String makes Person a fmt.Stringer: Person{"James", 32} prints as "James (32)".
func (p Person) String() string {
	// TODO
	return ""
}

// Area of a square is length * width. (HOE #62)
func (s Square) Area() float64 {
	// TODO
	return 0
}

// Area of a circle is π r². (HOE #62)
func (c Circle) Area() float64 {
	// TODO: math.Pi and math.Pow
	return 0
}

// Info returns the area of any Shape, formatted like "area: 6.00". (HOE #62)
func Info(s Shape) string {
	// TODO
	return ""
}

// WriteGreeting writes "Hello, <name>!\n" to w and returns any write error.
// It works with files, buffers, os.Stdout... anything that is an io.Writer.
func WriteGreeting(w io.Writer, name string) error {
	// TODO
	return nil
}

// Add returns a + b.
func Add(a, b int) int {
	// TODO
	return 0
}

// Subtract returns a - b.
func Subtract(a, b int) int {
	// TODO
	return 0
}

// DoMath calls f with a and b and returns the result. (HOE #64)
func DoMath(a, b int, f func(int, int) int) int {
	// TODO
	return 0
}

// Greet looks up a user in ds and returns "Hello, <Name>". If the lookup
// fails, return the error wrapped with fmt.Errorf and %w. (HOE #67)
func Greet(ds Datastore, id int) (string, error) {
	// TODO
	return "", nil
}

// FortyTwo returns a func which returns 42. (HOE #70)
func FortyTwo() func() int {
	// TODO
	return nil
}

// Squared returns n * n. (HOE #71)
func Squared(n int) int {
	// TODO
	return 0
}

// PrintSquare uses the callback f to return "the square of <n> is <f(n)>". (HOE #71)
func PrintSquare(f func(int) int, n int) string {
	// TODO
	return ""
}

// Counter returns a func that returns 1, 2, 3... on successive calls.
// Each counter returned by Counter is independent. (HOE #72: closure)
func Counter() func() int {
	// TODO
	return nil
}

// Factorial returns n! using recursion. Factorial(0) == 1.
func Factorial(n int) int {
	// TODO
	return 0
}

// Timed is a wrapper: it runs f and returns how long f took. (HOE #73)
func Timed(f func()) time.Duration {
	// TODO
	return 0
}
