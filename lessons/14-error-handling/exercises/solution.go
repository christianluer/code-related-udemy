//go:build solution

package exercises

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadConfig(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}
	return b, nil
}

func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse age %q: %w", s, err)
	}
	if n < 0 || n > 150 {
		return 0, fmt.Errorf("age %d: %w", n, ErrInvalidAge)
	}
	return n, nil
}

func (c CustomErr) Error() string {
	return "custom error: " + c.Info
}

func Describe(err error) string {
	return err.Error()
}

func (s *SqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", s.Lat, s.Long, s.Err)
}

func (s *SqrtError) Unwrap() error {
	return s.Err
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &SqrtError{
			Lat:  "50.2289 N",
			Long: "99.4656 W",
			Err:  fmt.Errorf("square root of negative number: %v", f),
		}
	}
	return math.Sqrt(f), nil
}

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

func ValidateAll(values []string, check func(string) error) error {
	var errs []error
	for _, v := range values {
		if err := check(v); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
