package exercises

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

type celsius float64

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum(ints) = %v, want 6", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4 {
		t.Errorf("Sum(float64s) = %v, want 4", got)
	}
	if got := Sum([]celsius{20, 1.5}); got != 21.5 {
		t.Errorf("Sum(celsius) = %v, want 21.5", got)
	}
}

func TestMax(t *testing.T) {
	if got, ok := Max(3, 9, 2); got != 9 || !ok {
		t.Errorf("Max(3, 9, 2) = (%v, %v), want (9, true)", got, ok)
	}
	if got, ok := Max("apple", "pear", "fig"); got != "pear" || !ok {
		t.Errorf("Max(strings) = (%q, %v), want (\"pear\", true)", got, ok)
	}
	if got, ok := Max(-5.5, -2.0); got != -2 || !ok {
		t.Errorf("Max(-5.5, -2.0) = (%v, %v), want (-2, true)", got, ok)
	}
	if _, ok := Max[int](); ok {
		t.Error("Max() of nothing should return ok == false")
	}
}

func TestContains(t *testing.T) {
	if !Contains([]string{"a", "b"}, "b") || Contains([]string{"a"}, "z") {
		t.Error("Contains is wrong for strings")
	}
	type point struct{ X, Y int }
	if !Contains([]point{{1, 2}, {3, 4}}, point{3, 4}) {
		t.Error("Contains is wrong for structs")
	}
}

func TestMap(t *testing.T) {
	if got := Map([]int{1, 2, 3}, strconv.Itoa); !slices.Equal(got, []string{"1", "2", "3"}) {
		t.Errorf("Map(ints, Itoa) = %v", got)
	}
	if got := Map([]string{"go", "gopher"}, func(s string) int { return len(s) }); !slices.Equal(got, []int{2, 6}) {
		t.Errorf("Map(strings, len) = %v", got)
	}
}

func TestFilter(t *testing.T) {
	evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	if !slices.Equal(evens, []int{2, 4, 6}) {
		t.Errorf("Filter evens = %v", evens)
	}
	gs := Filter([]string{"go", "rust", "gopher"}, func(s string) bool { return strings.HasPrefix(s, "go") })
	if !slices.Equal(gs, []string{"go", "gopher"}) {
		t.Errorf("Filter prefix = %v", gs)
	}
}

func TestStack(t *testing.T) {
	var s Stack[string]
	if s.Len() != 0 {
		t.Fatalf("empty stack Len() = %d", s.Len())
	}
	s.Push("a")
	s.Push("b")
	s.Push("c")
	if s.Len() != 3 {
		t.Fatalf("Len() = %d after 3 pushes", s.Len())
	}
	var got []string
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		got = append(got, v)
		if len(got) > 3 {
			break
		}
	}
	if !slices.Equal(got, []string{"c", "b", "a"}) {
		t.Errorf("popped %v, want [c b a]", got)
	}
}
