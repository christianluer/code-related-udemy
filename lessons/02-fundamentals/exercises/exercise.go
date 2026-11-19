//go:build !solution

// Package exercises contains the exercises for lesson 02.
//
// Replace each TODO, then run:  go run ./cmd/learn check 2
package exercises

// Describe returns the value and type of v, formatted like "42 (int)".
// (HOE #11: print verbs to show value and type.)
func Describe(v any) string {
	// TODO: fmt.Sprintf with %v and %T
	return ""
}

// Formats returns n as decimal, binary and hexadecimal (with a 0x prefix).
//
//	Formats(747) == "747", "1011101011", "0x2eb"
//
// (HOE #12)
func Formats(n int) (dec, bin, hex string) {
	// TODO
	return "", "", ""
}

// TODO: (HOE #8, #9) use iota and bit shifting so that
// KB = 1024, MB = 1024 KB, GB = 1024 MB, ... EB = 1024 PB.
const (
	KB = 0
	MB = 0
	GB = 0
	TB = 0
	PB = 0
	EB = 0
)

// TODO: (HOE #8) use iota so that Sunday == 0, Monday == 1, ... Saturday == 6.
const (
	Sunday    Weekday = 0
	Monday    Weekday = 0
	Tuesday   Weekday = 0
	Wednesday Weekday = 0
	Thursday  Weekday = 0
	Friday    Weekday = 0
	Saturday  Weekday = 0
)

// Limits returns the largest int8 and the largest uint8. (HOE #13)
func Limits() (int8, uint8) {
	// TODO
	return 0, 0
}

// Average returns the mean of nums as a float64, or 0 for an empty slice.
// Average([]int{1, 2}) must be 1.5, not 1.
func Average(nums []int) float64 {
	// TODO: you will need a conversion
	return 0
}

// OnesCount returns how many bits are set to 1 in n, using & and >>.
// OnesCount(0b1011) == 3. (HOE #7: bitwise ops & bit shifting)
func OnesCount(n uint) int {
	// TODO
	return 0
}
