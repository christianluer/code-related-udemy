package exercises

// Number is a type set constraint. Thanks to ~, named types such as
// `type Celsius float64` satisfy it too.
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Stack is a generic last-in, first-out stack.
type Stack[T any] struct {
	items []T
}
