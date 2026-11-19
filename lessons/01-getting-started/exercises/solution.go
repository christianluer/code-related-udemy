//go:build solution

package exercises

import (
	"fmt"
	"unicode/utf8"
)

func Greeting() string {
	return "Hello, gophers 😃🌴!"
}

func RawLiteral() string {
	return `Hello\nGophers`
}

func ByteAndRuneCount(s string) (bytes, runes int) {
	return len(s), utf8.RuneCountInString(s)
}

func CodePoints(s string) []string {
	var out []string
	for _, r := range s {
		out = append(out, fmt.Sprintf("%U", r))
	}
	return out
}
