//go:build !solution

// Package exercises contains the exercises for lesson 14.
//
// Replace each TODO, then run:  go run ./cmd/learn check 14
package exercises

import "os"

// ReadConfig returns the contents of path. Instead of ignoring the error with
// the blank identifier, check it and WRAP it with the path for context, so that
// errors.Is(err, fs.ErrNotExist) still works for the caller. (HOE #1, #2)
func ReadConfig(path string) ([]byte, error) {
	// TODO
	b, _ := os.ReadFile(path)
	return b, nil
}

// ParseAge converts s to an age.
//   - if s isn't a number, return an error that wraps strconv's error
//   - if the age is < 0 or > 150, return an error that wraps ErrInvalidAge
func ParseAge(s string) (int, error) {
	// TODO: strconv.Atoi, fmt.Errorf with %w
	return 0, nil
}

// Error makes CustomErr implement the built-in error interface.
// It must return "custom error: <Info>". (HOE #3)
func (c CustomErr) Error() string {
	// TODO
	return ""
}

// Describe takes any value of type error and returns its message. (HOE #3)
func Describe(err error) string {
	// TODO
	return ""
}

// Error returns "math error: <Lat> <Long> <Err>". (HOE #4)
func (s *SqrtError) Error() string {
	// TODO
	return ""
}

// Unwrap lets errors.Is / errors.As look inside a SqrtError.
func (s *SqrtError) Unwrap() error {
	// TODO
	return nil
}

// Sqrt returns the square root of f. For negative numbers it returns a
// *SqrtError with Lat "50.2289 N", Long "99.4656 W" and a descriptive Err.
// (HOE #4)
func Sqrt(f float64) (float64, error) {
	// TODO: math.Sqrt
	return 0, nil
}

// SafeDivide returns a / b. Dividing by zero panics in Go; SafeDivide must
// RECOVER from that panic and return it as an error instead.
// Don't check b == 0 yourself; practice recover. (Recover)
func SafeDivide(a, b int) (result int, err error) {
	// TODO: defer + recover
	return 0, nil
}

// ValidateAll runs check on every value and returns all the failures combined
// with errors.Join, or nil if every check passed.
func ValidateAll(values []string, check func(string) error) error {
	// TODO
	return nil
}
