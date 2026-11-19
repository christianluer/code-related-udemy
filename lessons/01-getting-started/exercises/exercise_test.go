package exercises

import (
	"slices"
	"testing"
)

func TestGreeting(t *testing.T) {
	g := Greeting()
	for _, r := range g {
		if r >= 0x1F300 || (r >= 0x2600 && r <= 0x27BF) {
			return
		}
	}
	t.Errorf("Greeting() = %q, want it to contain an emoji", g)
}

func TestRawLiteral(t *testing.T) {
	want := "Hello\\nGophers"
	if got := RawLiteral(); got != want {
		t.Errorf("RawLiteral() = %q, want %q (backslash + n, not a newline)", got, want)
	}
}

func TestByteAndRuneCount(t *testing.T) {
	cases := []struct {
		in           string
		bytes, runes int
	}{
		{"Go", 2, 2},
		{"Hello, 世界", 13, 9},
		{"🌴", 4, 1},
		{"", 0, 0},
	}
	for _, c := range cases {
		b, r := ByteAndRuneCount(c.in)
		if b != c.bytes || r != c.runes {
			t.Errorf("ByteAndRuneCount(%q) = (%d, %d), want (%d, %d)", c.in, b, r, c.bytes, c.runes)
		}
	}
}

func TestCodePoints(t *testing.T) {
	got := CodePoints("Go世")
	want := []string{"U+0047", "U+006F", "U+4E16"}
	if !slices.Equal(got, want) {
		t.Errorf("CodePoints(%q) = %v, want %v", "Go世", got, want)
	}
}
