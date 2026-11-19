package exercises

import (
	"slices"
	"testing"
)

func TestAddOne(t *testing.T) {
	x := 41
	AddOne(&x)
	if x != 42 {
		t.Errorf("after AddOne(&x), x = %d, want 42", x)
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("after Swap, a=%d b=%d, want a=2 b=1", a, b)
	}
}

func TestChangeFirstValue(t *testing.T) {
	orig := Person{First: "James"}
	got := ChangeFirstValue(orig, "Jenny")
	if got.First != "Jenny" {
		t.Errorf("returned First = %q, want %q", got.First, "Jenny")
	}
	if orig.First != "James" {
		t.Errorf("the original was modified: %q", orig.First)
	}
}

func TestChangeFirstPointer(t *testing.T) {
	p := Person{First: "James"}
	ChangeFirstPointer(&p, "Q")
	if p.First != "Q" {
		t.Errorf("p.First = %q, want %q", p.First, "Q")
	}
}

func TestMethodSets(t *testing.T) {
	y := NewYoungin("Henry", 15)
	if y == nil {
		t.Fatal("NewYoungin returned nil")
	}
	y.Birthday()
	y.Birthday()
	if got := y.Years(); got != 17 {
		t.Errorf("after two birthdays Years() = %d, want 17 (does Birthday change a copy?)", got)
	}
	k := Kid{Name: "Padget", Age: 13}
	k.Birthday()
	if k.Age != 14 {
		t.Errorf("k.Birthday() on an addressable Kid: Age = %d, want 14", k.Age)
	}
}

func TestNewInt(t *testing.T) {
	a := NewInt(5)
	b := NewInt(5)
	if a == nil || b == nil {
		t.Fatal("NewInt returned nil")
	}
	if *a != 5 {
		t.Errorf("*NewInt(5) = %d", *a)
	}
	*a = 99
	if *b != 5 || a == b {
		t.Error("each NewInt call must return a pointer to its own int")
	}
}

func TestDoubleAll(t *testing.T) {
	xs := []int{1, 2, 3}
	DoubleAll(xs)
	if want := []int{2, 4, 6}; !slices.Equal(xs, want) {
		t.Errorf("after DoubleAll, xs = %v, want %v", xs, want)
	}
}
