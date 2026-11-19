//go:build solution

package exercises

import "strings"

func Squares() [5]int {
	var a [5]int
	for i := range a {
		a[i] = i * i
	}
	return a
}

func SliceUp(x []int) (a, b, c, d []int) {
	return x[:5], x[5:], x[2:7], x[1:6]
}

func Extend(x []int, single int, many []int) []int {
	x = append(x, single)
	x = append(x, many...)
	return x
}

func Delete(x []int, i, j int) []int {
	out := make([]int, 0, len(x)-(j-i))
	out = append(out, x[:i]...)
	return append(out, x[j:]...)
}

func CleanStates(names []string) []string {
	out := make([]string, 0, len(names))
	for _, n := range names {
		out = append(out, strings.TrimSpace(n))
	}
	return out
}

func Flatten(records [][]string) []string {
	var out []string
	for _, rec := range records {
		for _, field := range rec {
			out = append(out, field)
		}
	}
	return out
}

func Clone(x []int) []int {
	out := make([]int, len(x))
	copy(out, x)
	return out
}
