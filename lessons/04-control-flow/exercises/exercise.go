//go:build !solution

// Package exercises contains the exercises for lesson 04.
//
// Replace each TODO, then run:  go run ./cmd/learn check 4
package exercises

// Bucket describes where x falls (HOE #24):
//
//	0..100   -> "between 0 and 100"
//	101..200 -> "between 101 and 200"
//	201..250 -> "between 201 and 250"
//	anything else -> "out of range"
func Bucket(x int) string {
	// TODO
	return ""
}

// DescribePair returns the FIRST matching description, checked in this order (HOE #27):
//
//	"x and y are both less than 4"
//	"x and y are both greater than 6"
//	"x is greater than or equal to 4 and less than or equal to 6"
//	"y is not 5"
//	"none of the previous cases were met"
func DescribePair(x, y int) string {
	// TODO: if/else if or a tagless switch (HOE #28)
	return ""
}

// Odds returns every odd number from 0 to n inclusive.
// Use the modulus operator and the continue statement. (HOE #33)
func Odds(n int) []int {
	// TODO
	return nil
}

// FirstMultiple returns the first number >= start that is divisible by k (k > 0).
// Use a loop with a break (or return) statement. (HOE #31, #32)
func FirstMultiple(start, k int) int {
	// TODO
	return 0
}

// Grid returns every coordinate of an n x n grid using nested loops, in order:
//
//	Grid(2) == []string{"0,0", "0,1", "1,0", "1,1"}
//
// (HOE #34)
func Grid(n int) []string {
	// TODO
	return nil
}

// IndexValues returns "index:value" for every element of xi using for range.
//
//	IndexValues([]int{42, 43}) == []string{"0:42", "1:43"}
//
// (HOE #35)
func IndexValues(xi []int) []string {
	// TODO
	return nil
}

// AgeOf looks up name in m using the comma ok idiom (HOE #37):
//
//	found     -> "James is 42"
//	not found -> "Q is not in the map"
func AgeOf(m map[string]int, name string) string {
	// TODO
	return ""
}
