//go:build solution

package exercises

import (
	"fmt"
	"io"
	"math"
	"time"
)

func Divmod(a, b int) (int, int) {
	return a / b, a % b
}

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func SumSlice(xi []int) int {
	return Sum(xi...)
}

func DeferOrder() (log []string) {
	defer func() { log = append(log, "one") }()
	defer func() { log = append(log, "two") }()
	defer func() { log = append(log, "three") }()
	log = append(log, "body")
	return log
}

func (p Person) Speak() string {
	return fmt.Sprintf("I am %s and I am %d years old", p.First, p.Age)
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.First, p.Age)
}

func (s Square) Area() float64 {
	return s.Length * s.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func Info(s Shape) string {
	return fmt.Sprintf("area: %.2f", s.Area())
}

func WriteGreeting(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "Hello, %s!\n", name)
	return err
}

func Add(a, b int) int      { return a + b }
func Subtract(a, b int) int { return a - b }

func DoMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func Greet(ds Datastore, id int) (string, error) {
	u, err := ds.GetUser(id)
	if err != nil {
		return "", fmt.Errorf("greet user %d: %w", id, err)
	}
	return "Hello, " + u.Name, nil
}

func FortyTwo() func() int {
	return func() int { return 42 }
}

func Squared(n int) int {
	return n * n
}

func PrintSquare(f func(int) int, n int) string {
	return fmt.Sprintf("the square of %d is %d", n, f(n))
}

func Counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Timed(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
