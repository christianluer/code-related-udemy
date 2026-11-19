//go:build solution

package exercises

import (
	"slices"
	"strings"
)

// Years converts human years to dog years (1 human year = 7 dog years).
func Years(human int) int {
	return human * 7
}

// Count returns the number of words in s (words are separated by whitespace).
func Count(s string) int {
	return len(strings.Fields(s))
}

// UseCount returns how many times each word occurs in s (case-sensitive).
func UseCount(s string) map[string]int {
	m := map[string]int{}
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

// CenteredAvg returns the average of xs after dropping one smallest and one
// largest value. It returns 0 if xs has fewer than 3 elements.
func CenteredAvg(xs []int) float64 {
	if len(xs) < 3 {
		return 0
	}
	sorted := slices.Clone(xs)
	slices.Sort(sorted)
	sum := 0
	for _, v := range sorted[1 : len(sorted)-1] {
		sum += v
	}
	return float64(sum) / float64(len(sorted)-2)
}
