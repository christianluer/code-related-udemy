package exercises

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

const quote = `The road to success is always under construction and the road is long`

// TestYears is a TABLE TEST with subtests.
func TestYears(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 7},
		{"ten", 10, 70},
		{"forty-two", 42, 294},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Years(tt.in); got != tt.want {
				t.Errorf("Years(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := map[string]int{
		"":                0,
		"one":             1,
		"  two   words  ": 2,
		quote:             13,
	}
	for in, want := range tests {
		if got := Count(in); got != want {
			t.Errorf("Count(%q) = %d, want %d", in, got, want)
		}
	}
}

func TestUseCount(t *testing.T) {
	got := UseCount(quote)
	if got["road"] != 2 || got["is"] != 2 || got["The"] != 1 || got["the"] != 1 || len(got) != 11 {
		keys := slices.Sorted(maps.Keys(got))
		t.Errorf("UseCount(quote) = %v (keys %v)", got, keys)
	}
}

func TestCenteredAvg(t *testing.T) {
	tests := []struct {
		in   []int
		want float64
	}{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{10, 0, 5}, 5},
		{[]int{1, 2, 3, 4}, 2.5},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{1, 2}, 0},
		{nil, 0},
	}
	for _, tt := range tests {
		in := slices.Clone(tt.in)
		if got := CenteredAvg(in); got != tt.want {
			t.Errorf("CenteredAvg(%v) = %v, want %v", tt.in, got, tt.want)
		}
		if !slices.Equal(in, tt.in) {
			t.Errorf("CenteredAvg reordered the caller's slice: %v", in)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

func ExampleCount() {
	fmt.Println(Count("The road to success is always under construction"))
	// Output: 8
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output: 6
}

func BenchmarkYears(b *testing.B) {
	for b.Loop() {
		Years(42)
	}
}

func BenchmarkCount(b *testing.B) {
	long := strings.Repeat(quote+" ", 100)
	for b.Loop() {
		Count(long)
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	xs := []int{9000, 4, 10, 8, 6, 12, 7, 33, 1}
	for b.Loop() {
		CenteredAvg(xs)
	}
}
