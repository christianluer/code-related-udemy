//go:build solution

package exercises

import (
	"fmt"
	"math"
)

func Describe(v any) string {
	return fmt.Sprintf("%v (%T)", v, v)
}

func Formats(n int) (dec, bin, hex string) {
	return fmt.Sprintf("%d", n), fmt.Sprintf("%b", n), fmt.Sprintf("%#x", n)
}

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func Limits() (int8, uint8) {
	return math.MaxInt8, math.MaxUint8
}

func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}

func OnesCount(n uint) int {
	count := 0
	for n > 0 {
		count += int(n & 1)
		n >>= 1
	}
	return count
}
