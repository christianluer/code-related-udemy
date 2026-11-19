// Checking errors, wrapping with %w, errors.Is / errors.As, custom errors, Join.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"strconv"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

type sqrtError struct {
	lat, long string
	err       error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func (se sqrtError) Unwrap() error { return se.err }

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{"50.2289 N", "99.4656 W", ErrNorgateMath}
	}
	return math.Sqrt(f), nil
}

func main() {
	// 1. Check errors.
	_, err := os.Open("no-such-file.txt")
	if err != nil {
		fmt.Println("fmt:", err)
		log.Println("log:", err)
	}

	// 2. Wrap with context.
	wrapped := fmt.Errorf("loading settings: %w", err)
	fmt.Println(wrapped)
	fmt.Println("errors.Is(fs.ErrNotExist):", errors.Is(wrapped, fs.ErrNotExist))
	var pe *fs.PathError
	if errors.As(wrapped, &pe) {
		fmt.Println("errors.As found *fs.PathError for path:", pe.Path)
	}

	// 3. Custom error type with extra info.
	if _, err := sqrt(-10); err != nil {
		fmt.Println(err)
		fmt.Println("is ErrNorgateMath?", errors.Is(err, ErrNorgateMath))
		var se sqrtError
		if errors.As(err, &se) {
			fmt.Println("lat/long:", se.lat, se.long)
		}
	}

	// 4. strconv errors are typed too.
	if _, err := strconv.Atoi("forty-two"); err != nil {
		var ne *strconv.NumError
		fmt.Println(err, "| NumError?", errors.As(err, &ne))
	}

	// 5. Join several errors.
	joined := errors.Join(errors.New("name is empty"), errors.New("age is negative"))
	fmt.Printf("joined:\n%v\n", joined)
}
