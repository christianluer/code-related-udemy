package exercises

import (
	"runtime"
	"testing"
)

func TestIncrement(t *testing.T) {
	first := Increment()
	second := Increment()
	third := Increment()
	if second != first+1 || third != second+1 || first < 1 {
		t.Errorf("Increment() returned %d, %d, %d; want consecutive values starting at 1", first, second, third)
	}
}

func TestIsExported(t *testing.T) {
	cases := map[string]bool{
		"Println": true, "println": false, "X": true, "x": false,
		"_Hidden": false, "Ñandú": true, "élan": false, "": false,
	}
	for in, want := range cases {
		if got := IsExported(in); got != want {
			t.Errorf("IsExported(%q) = %v, want %v", in, got, want)
		}
	}
}

func TestSHA256Hex(t *testing.T) {
	cases := map[string]string{
		"hello": "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		"":      "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}
	for in, want := range cases {
		if got := SHA256Hex([]byte(in)); got != want {
			t.Errorf("SHA256Hex(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestBump(t *testing.T) {
	ok := []struct{ in, part, want string }{
		{"v1.2.3", "patch", "v1.2.4"},
		{"v1.2.3", "minor", "v1.3.0"},
		{"v1.2.3", "major", "v2.0.0"},
		{"v0.9.9", "minor", "v0.10.0"},
	}
	for _, c := range ok {
		got, err := Bump(c.in, c.part)
		if err != nil || got != c.want {
			t.Errorf("Bump(%q, %q) = (%q, %v), want (%q, nil)", c.in, c.part, got, err, c.want)
		}
	}
	bad := [][2]string{{"1.2.3", "patch"}, {"v1.2", "patch"}, {"v1.x.3", "minor"}, {"v1.2.3", "huge"}}
	for _, c := range bad {
		if _, err := Bump(c[0], c[1]); err == nil {
			t.Errorf("Bump(%q, %q) should return an error", c[0], c[1])
		}
	}
}

func TestPlatform(t *testing.T) {
	want := runtime.GOOS + "/" + runtime.GOARCH
	if got := Platform(); got != want {
		t.Errorf("Platform() = %q, want %q", got, want)
	}
}
