// Benchmark string concatenation three ways, from a regular program, using
// testing.Benchmark. In real projects put benchmarks in _test.go files and
// run them with `go test -bench . -benchmem`.
package main

import (
	"fmt"
	"strings"
	"testing"
)

var words = strings.Fields(strings.Repeat("the quick brown fox jumps over the lazy dog ", 50))

func concat(xs []string) string {
	s := ""
	for i, w := range xs {
		if i > 0 {
			s += " "
		}
		s += w
	}
	return s
}

func join(xs []string) string { return strings.Join(xs, " ") }

func builder(xs []string) string {
	var sb strings.Builder
	for i, w := range xs {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(w)
	}
	return sb.String()
}

func main() {
	fmt.Printf("joining %d words\n\n", len(words))
	for _, c := range []struct {
		name string
		fn   func([]string) string
	}{
		{"+= concat", concat},
		{"strings.Join", join},
		{"strings.Builder", builder},
	} {
		r := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				c.fn(words)
			}
		})
		fmt.Printf("%-16s %10d ns/op %8d B/op %5d allocs/op\n",
			c.name, r.NsPerOp(), r.AllocedBytesPerOp(), r.AllocsPerOp())
	}
}
