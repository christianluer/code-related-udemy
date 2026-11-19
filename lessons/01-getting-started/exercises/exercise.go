//go:build !solution

// Package exercises contains the exercises for lesson 01.
//
// Replace each TODO, then run:  go run ./cmd/learn check 1
package exercises

// Greeting returns a friendly greeting that contains at least one emoji.
// (Hands-on exercise #1: print out some text with emojis.)
func Greeting() string {
	// TODO: return a greeting with an emoji in it, e.g. "Hello 🌴"
	return "Hello"
}

// RawLiteral returns the text  Hello\nGophers  where \n is NOT a newline,
// but the two characters backslash and n. Use a raw string literal.
func RawLiteral() string {
	// TODO: use backticks
	return "Hello\nGophers"
}

// ByteAndRuneCount returns how many bytes and how many runes are in s.
func ByteAndRuneCount(s string) (bytes, runes int) {
	// TODO: len() counts bytes; package unicode/utf8 can count runes
	return 0, 0
}

// CodePoints returns the Unicode code point of every rune in s,
// formatted like "U+0047" (hint: the %U verb).
//
//	CodePoints("Go") == []string{"U+0047", "U+006F"}
func CodePoints(s string) []string {
	// TODO
	return nil
}
