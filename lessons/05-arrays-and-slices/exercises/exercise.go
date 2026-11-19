//go:build !solution

// Package exercises contains the exercises for lesson 05.
//
// Replace each TODO, then run:  go run ./cmd/learn check 5
package exercises

// Squares returns an ARRAY of 5 ints where each index position holds its
// square: [0 1 4 9 16]. (HOE #40, #42)
func Squares() [5]int {
	// TODO
	return [5]int{}
}

// SliceUp takes x = [42 43 44 45 46 47 48 49 50 51] and uses SLICING to return
// (HOE #44):
//
//	a = [42 43 44 45 46]
//	b = [47 48 49 50 51]
//	c = [44 45 46 47 48]
//	d = [43 44 45 46 47]
func SliceUp(x []int) (a, b, c, d []int) {
	// TODO
	return nil, nil, nil, nil
}

// Extend appends single and then every element of many to x, and returns
// the result. (HOE #45: append a value, then a slice using ...)
func Extend(x []int, single int, many []int) []int {
	// TODO
	return x
}

// Delete returns a new slice with the elements x[i:j] removed, using append
// and slicing. The caller's x must NOT be modified. (HOE #46)
//
//	Delete([]int{42, 43, 44, 45, 46, 47, 48}, 3, 6) == [42 43 44 48]
func Delete(x []int, i, j int) []int {
	// TODO
	return x
}

// CleanStates returns names with surrounding spaces trimmed.
// Build the result with make and append so that the underlying array is
// created only ONCE: len and cap of the result must both equal len(names).
// (HOE #47, try it with the States variable in data.go)
func CleanStates(names []string) []string {
	// TODO: make([]string, 0, ...) + append + strings.TrimSpace
	return nil
}

// Flatten ranges over the records, then over each record's fields, and
// returns every field in order. (HOE #48)
func Flatten(records [][]string) []string {
	// TODO
	return nil
}

// Clone returns a copy of x that does not share x's underlying array.
func Clone(x []int) []int {
	// TODO: returning x itself shares the array; that's the bug to fix
	return x
}
