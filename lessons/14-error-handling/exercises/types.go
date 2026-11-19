package exercises

import "errors"

// ErrInvalidAge is a sentinel error: callers check for it with errors.Is.
var ErrInvalidAge = errors.New("invalid age")

// CustomErr is an error type that carries extra information. (HOE #3)
type CustomErr struct {
	Info string
}

// SqrtError reports where a bad square root happened. (HOE #4)
type SqrtError struct {
	Lat  string
	Long string
	Err  error
}
