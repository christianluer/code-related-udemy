// Checksums: change a single character and the SHA-256 hash changes completely.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := []byte("The quick brown fox jumps over the lazy dog")
	b := []byte("The quick brown fox jumps over the lazy cog")

	fmt.Printf("%s\n  %x\n", a, sha256.Sum256(a))
	fmt.Printf("%s\n  %x\n", b, sha256.Sum256(b))

	fmt.Println("\nSame thing at the terminal:")
	fmt.Println("  echo -n 'The quick brown fox jumps over the lazy dog' | shasum -a 256")
}
