//go:build !solution

// Package exercises contains the exercises for lesson 10.
//
// Replace each TODO, then run:  go run ./cmd/learn check 10
package exercises

import "cmp"

// Sum adds up xs. It works for any Number, including named types.
func Sum[T Number](xs []T) T {
	// TODO
	var zero T
	return zero
}

// Max returns the largest element and true, or the zero value and false
// when xs is empty.
func Max[T cmp.Ordered](xs ...T) (T, bool) {
	// TODO
	var zero T
	return zero, false
}

// Contains reports whether v is in xs.
func Contains[T comparable](xs []T, v T) bool {
	// TODO
	return false
}

// Map returns a new slice holding f(x) for every x in xs.
//
//	Map([]int{1, 2}, strconv.Itoa) == []string{"1", "2"}
func Map[T, U any](xs []T, f func(T) U) []U {
	// TODO
	return nil
}

// Filter returns the elements of xs for which keep returns true.
func Filter[T any](xs []T, keep func(T) bool) []T {
	// TODO
	return nil
}

// Push adds v to the top of the stack.
func (s *Stack[T]) Push(v T) {
	// TODO
}

// Pop removes and returns the top element, or the zero value and false
// when the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	// TODO
	var zero T
	return zero, false
}

// Len returns the number of elements on the stack.
func (s *Stack[T]) Len() int {
	// TODO
	return -1
}
