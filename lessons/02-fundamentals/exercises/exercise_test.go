package exercises

import "testing"

func TestDescribe(t *testing.T) {
	cases := map[any]string{
		42:       "42 (int)",
		"gopher": "gopher (string)",
		3.5:      "3.5 (float64)",
		true:     "true (bool)",
	}
	for in, want := range cases {
		if got := Describe(in); got != want {
			t.Errorf("Describe(%#v) = %q, want %q", in, got, want)
		}
	}
}

func TestFormats(t *testing.T) {
	cases := []struct {
		n             int
		dec, bin, hex string
	}{
		{747, "747", "1011101011", "0x2eb"},
		{911, "911", "1110001111", "0x38f"},
		{90210, "90210", "10110000001100010", "0x16062"},
	}
	for _, c := range cases {
		d, b, h := Formats(c.n)
		if d != c.dec || b != c.bin || h != c.hex {
			t.Errorf("Formats(%d) = (%q, %q, %q), want (%q, %q, %q)", c.n, d, b, h, c.dec, c.bin, c.hex)
		}
	}
}

func TestByteSizes(t *testing.T) {
	got := []int64{KB, MB, GB, TB, PB, EB}
	names := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	want := int64(1024)
	for i, g := range got {
		if g != want {
			t.Errorf("%s = %d, want %d", names[i], g, want)
		}
		want *= 1024
	}
}

func TestWeekdays(t *testing.T) {
	days := []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}
	for i, d := range days {
		if int(d) != i {
			t.Errorf("day %d has value %d, want %d", i, d, i)
		}
	}
}

func TestLimits(t *testing.T) {
	i, u := Limits()
	if i != 127 || u != 255 {
		t.Errorf("Limits() = (%d, %d), want (127, 255)", i, u)
	}
}

func TestAverage(t *testing.T) {
	cases := []struct {
		in   []int
		want float64
	}{
		{[]int{1, 2}, 1.5},
		{[]int{2, 4, 9}, 5},
		{nil, 0},
	}
	for _, c := range cases {
		if got := Average(c.in); got != c.want {
			t.Errorf("Average(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestOnesCount(t *testing.T) {
	cases := map[uint]int{0: 0, 1: 1, 0b1011: 3, 255: 8, 1 << 40: 1}
	for in, want := range cases {
		if got := OnesCount(in); got != want {
			t.Errorf("OnesCount(%b) = %d, want %d", in, got, want)
		}
	}
}
