package exercises

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"slices"
	"testing"
	"time"
)

func TestMultipleReturns(t *testing.T) {
	if q, r := Divmod(17, 5); q != 3 || r != 2 {
		t.Errorf("Divmod(17, 5) = (%d, %d), want (3, 2)", q, r)
	}
}

func TestVariadic(t *testing.T) {
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1, 2, 3) = %d, want 6", got)
	}
	if got := SumSlice([]int{10, 20, 30, 40}); got != 100 {
		t.Errorf("SumSlice([10 20 30 40]) = %d, want 100", got)
	}
}

func TestDefer(t *testing.T) {
	want := []string{"body", "three", "two", "one"}
	if got := DeferOrder(); !slices.Equal(got, want) {
		t.Errorf("DeferOrder() = %v, want %v", got, want)
	}
}

func TestMethods(t *testing.T) {
	p := Person{First: "James", Age: 32}
	if got, want := p.Speak(), "I am James and I am 32 years old"; got != want {
		t.Errorf("Speak() = %q, want %q", got, want)
	}
}

func TestStringer(t *testing.T) {
	var s fmt.Stringer = Person{First: "James", Age: 32}
	if got := fmt.Sprint(s); got != "James (32)" {
		t.Errorf("fmt.Sprint(person) = %q, want %q", got, "James (32)")
	}
}

func TestInterfaces(t *testing.T) {
	if got := (Square{Length: 2, Width: 3}).Area(); got != 6 {
		t.Errorf("Square{2, 3}.Area() = %v, want 6", got)
	}
	if got := (Circle{Radius: 2}).Area(); math.Abs(got-4*math.Pi) > 1e-9 {
		t.Errorf("Circle{2}.Area() = %v, want %v", got, 4*math.Pi)
	}
	shapes := []Shape{Square{Length: 2, Width: 3}, Circle{Radius: 1}}
	want := []string{"area: 6.00", "area: 3.14"}
	for i, s := range shapes {
		if got := Info(s); got != want[i] {
			t.Errorf("Info(%T) = %q, want %q", s, got, want[i])
		}
	}
}

type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, errors.New("disk full") }

func TestWriter(t *testing.T) {
	var b bytes.Buffer
	if err := WriteGreeting(&b, "Gopher"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := b.String(); got != "Hello, Gopher!\n" {
		t.Errorf("wrote %q, want %q", got, "Hello, Gopher!\n")
	}
	if err := WriteGreeting(brokenWriter{}, "x"); err == nil {
		t.Error("WriteGreeting should return the writer's error")
	}
}

func TestDoMath(t *testing.T) {
	if got := DoMath(42, 16, Add); got != 58 {
		t.Errorf("DoMath(42, 16, Add) = %d, want 58", got)
	}
	if got := DoMath(42, 16, Subtract); got != 26 {
		t.Errorf("DoMath(42, 16, Subtract) = %d, want 26", got)
	}
	if got := DoMath(6, 7, func(a, b int) int { return a * b }); got != 42 {
		t.Errorf("DoMath(6, 7, multiply) = %d, want 42", got)
	}
}

type mockDB struct {
	users map[int]User
}

var errNotFound = errors.New("not found")

func (m mockDB) GetUser(id int) (User, error) {
	u, ok := m.users[id]
	if !ok {
		return User{}, errNotFound
	}
	return u, nil
}

func TestMockDB(t *testing.T) {
	db := mockDB{users: map[int]User{1: {ID: 1, Name: "Alice"}}}
	got, err := Greet(db, 1)
	if err != nil || got != "Hello, Alice" {
		t.Errorf("Greet(db, 1) = (%q, %v), want (\"Hello, Alice\", nil)", got, err)
	}
	_, err = Greet(db, 2)
	if !errors.Is(err, errNotFound) {
		t.Errorf("Greet(db, 2) error = %v, want it to wrap errNotFound", err)
	}
}

func TestFuncExpression(t *testing.T) {
	f := FortyTwo()
	if f == nil {
		t.Fatal("FortyTwo() returned nil")
	}
	if got := f(); got != 42 {
		t.Errorf("FortyTwo()() = %d, want 42", got)
	}
}

func TestCallback(t *testing.T) {
	if got := Squared(4); got != 16 {
		t.Errorf("Squared(4) = %d, want 16", got)
	}
	if got, want := PrintSquare(Squared, 4), "the square of 4 is 16"; got != want {
		t.Errorf("PrintSquare(Squared, 4) = %q, want %q", got, want)
	}
	fake := func(int) int { return 99 }
	if got, want := PrintSquare(fake, 3), "the square of 3 is 99"; got != want {
		t.Errorf("PrintSquare must call the callback: got %q, want %q", got, want)
	}
}

func TestClosure(t *testing.T) {
	a, b := Counter(), Counter()
	if a == nil || b == nil {
		t.Fatal("Counter() returned nil")
	}
	got := []int{a(), a(), a(), b(), a(), b()}
	if want := []int{1, 2, 3, 1, 4, 2}; !slices.Equal(got, want) {
		t.Errorf("calls returned %v, want %v (counters must be independent)", got, want)
	}
}

func TestRecursion(t *testing.T) {
	cases := map[int]int{0: 1, 1: 1, 5: 120, 10: 3628800}
	for in, want := range cases {
		if got := Factorial(in); got != want {
			t.Errorf("Factorial(%d) = %d, want %d", in, got, want)
		}
	}
}

func TestWrapper(t *testing.T) {
	called := false
	d := Timed(func() {
		time.Sleep(20 * time.Millisecond)
		called = true
	})
	if !called {
		t.Fatal("Timed did not call f")
	}
	if d < 20*time.Millisecond || d > 2*time.Second {
		t.Errorf("Timed returned %v, want about 20ms", d)
	}
}
