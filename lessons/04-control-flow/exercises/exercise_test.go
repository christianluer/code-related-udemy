package exercises

import (
	"slices"
	"testing"
)

func TestBucket(t *testing.T) {
	cases := map[int]string{
		0: "between 0 and 100", 100: "between 0 and 100",
		101: "between 101 and 200", 200: "between 101 and 200",
		201: "between 201 and 250", 250: "between 201 and 250",
		-1: "out of range", 251: "out of range",
	}
	for in, want := range cases {
		if got := Bucket(in); got != want {
			t.Errorf("Bucket(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestDescribePair(t *testing.T) {
	cases := []struct {
		x, y int
		want string
	}{
		{1, 2, "x and y are both less than 4"},
		{7, 9, "x and y are both greater than 6"},
		{5, 0, "x is greater than or equal to 4 and less than or equal to 6"},
		{4, 9, "x is greater than or equal to 4 and less than or equal to 6"},
		{9, 3, "y is not 5"},
		{1, 7, "y is not 5"},
		{9, 5, "none of the previous cases were met"},
	}
	for _, c := range cases {
		if got := DescribePair(c.x, c.y); got != c.want {
			t.Errorf("DescribePair(%d, %d) = %q, want %q", c.x, c.y, got, c.want)
		}
	}
}

func TestOdds(t *testing.T) {
	if got, want := Odds(10), []int{1, 3, 5, 7, 9}; !slices.Equal(got, want) {
		t.Errorf("Odds(10) = %v, want %v", got, want)
	}
	if got, want := Odds(7), []int{1, 3, 5, 7}; !slices.Equal(got, want) {
		t.Errorf("Odds(7) = %v, want %v", got, want)
	}
}

func TestFirstMultiple(t *testing.T) {
	cases := [][3]int{{10, 7, 14}, {14, 7, 14}, {1, 1, 1}, {100, 33, 132}}
	for _, c := range cases {
		if got := FirstMultiple(c[0], c[1]); got != c[2] {
			t.Errorf("FirstMultiple(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}

func TestGrid(t *testing.T) {
	want := []string{"0,0", "0,1", "0,2", "1,0", "1,1", "1,2", "2,0", "2,1", "2,2"}
	if got := Grid(3); !slices.Equal(got, want) {
		t.Errorf("Grid(3) = %v, want %v", got, want)
	}
	if got := Grid(5); len(got) != 25 {
		t.Errorf("len(Grid(5)) = %d, want 25", len(got))
	}
}

func TestIndexValues(t *testing.T) {
	xi := []int{42, 43, 44, 45, 46, 47}
	want := []string{"0:42", "1:43", "2:44", "3:45", "4:46", "5:47"}
	if got := IndexValues(xi); !slices.Equal(got, want) {
		t.Errorf("IndexValues(%v) = %v, want %v", xi, got, want)
	}
}

func TestAgeOf(t *testing.T) {
	m := map[string]int{"James": 42, "Moneypenny": 32, "Baby": 0}
	cases := map[string]string{
		"James": "James is 42",
		"Baby":  "Baby is 0",
		"Q":     "Q is not in the map",
	}
	for in, want := range cases {
		if got := AgeOf(m, in); got != want {
			t.Errorf("AgeOf(m, %q) = %q, want %q", in, got, want)
		}
	}
}
