// ASCII, Unicode and UTF-8: strings are bytes, runes are code points.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hi, 世界 🌴"

	fmt.Println("string:     ", s)
	fmt.Println("len (bytes):", len(s))
	fmt.Println("runes:      ", utf8.RuneCountInString(s))

	fmt.Println("\nbyte by byte (index loop):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%02x ", s[i])
	}
	fmt.Println()

	fmt.Println("\nrune by rune (range loop):")
	for i, r := range s {
		fmt.Printf("byte offset %2d  %c  %U  %d bytes\n", i, r, r, utf8.RuneLen(r))
	}

	// Conversions between the three views of text.
	fmt.Println("\n[]byte: ", []byte("Go"))
	fmt.Println("[]rune: ", []rune("Go"))
	fmt.Println("string: ", string([]byte{71, 111}))

	// 'A' is a rune literal (an int32), not a string.
	fmt.Printf("'A' = %d = %#x = %U = %c\n", 'A', 'A', 'A', 'A')
}
