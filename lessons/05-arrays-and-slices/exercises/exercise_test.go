package exercises

import (
	"slices"
	"testing"
)

func fortyTwos() []int { return []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} }

func TestSquares(t *testing.T) {
	if got, want := Squares(), [5]int{0, 1, 4, 9, 16}; got != want {
		t.Errorf("Squares() = %v, want %v", got, want)
	}
}

func TestSliceUp(t *testing.T) {
	a, b, c, d := SliceUp(fortyTwos())
	checks := []struct {
		name      string
		got, want []int
	}{
		{"a", a, []int{42, 43, 44, 45, 46}},
		{"b", b, []int{47, 48, 49, 50, 51}},
		{"c", c, []int{44, 45, 46, 47, 48}},
		{"d", d, []int{43, 44, 45, 46, 47}},
	}
	for _, ch := range checks {
		if !slices.Equal(ch.got, ch.want) {
			t.Errorf("%s = %v, want %v", ch.name, ch.got, ch.want)
		}
	}
}

func TestExtend(t *testing.T) {
	got := Extend(fortyTwos(), 52, []int{53, 54, 55})
	want := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55}
	if !slices.Equal(got, want) {
		t.Errorf("Extend = %v, want %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	x := fortyTwos()
	got := Delete(x, 3, 6)
	if want := []int{42, 43, 44, 48, 49, 50, 51}; !slices.Equal(got, want) {
		t.Errorf("Delete(x, 3, 6) = %v, want %v", got, want)
	}
	if !slices.Equal(x, fortyTwos()) {
		t.Errorf("Delete modified the caller's slice: x is now %v", x)
	}
}

func TestCleanStates(t *testing.T) {
	got := CleanStates(States)
	if len(got) != 50 || cap(got) != 50 {
		t.Fatalf("len %d cap %d, want len 50 cap 50 (use make with the right capacity)", len(got), cap(got))
	}
	if got[0] != "Alabama" || got[49] != "Wyoming" || got[29] != "New Jersey" {
		t.Errorf("got %q, %q, %q; want trimmed names like \"Alabama\"", got[0], got[29], got[49])
	}
}

func TestFlatten(t *testing.T) {
	records := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."},
	}
	want := []string{"James", "Bond", "Shaken, not stirred", "Miss", "Moneypenny", "I'm 008."}
	if got := Flatten(records); !slices.Equal(got, want) {
		t.Errorf("Flatten = %v, want %v", got, want)
	}
}

func TestClone(t *testing.T) {
	x := []int{1, 2, 3}
	c := Clone(x)
	if !slices.Equal(c, x) {
		t.Fatalf("Clone(%v) = %v", x, c)
	}
	c[0] = 100
	if x[0] != 1 {
		t.Errorf("changing the clone changed the original: x = %v", x)
	}
}
