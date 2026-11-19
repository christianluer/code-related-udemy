//go:build !solution

// Package exercises contains the exercises for lesson 03.
//
// Replace each TODO, then run:  go run ./cmd/learn check 3
package exercises

// Increment adds one to a counter that lives at PACKAGE scope and returns
// the new value. The first call returns 1, the second 2, and so on. (HOE #14)
func Increment() int {
	// TODO: declare a package-level variable (outside this func) and use it here
	return 0
}

// IsExported reports whether name would be visible outside its package,
// i.e. whether it starts with an upper-case letter (any language).
//
//	IsExported("Println") == true
//	IsExported("println") == false
//	IsExported("Ñandú") == true
func IsExported(name string) bool {
	// TODO: packages unicode and unicode/utf8 will help
	return false
}

// SHA256Hex returns the SHA-256 checksum of data as a lowercase hex string,
// exactly what `shasum -a 256` prints.
func SHA256Hex(data []byte) string {
	// TODO: crypto/sha256
	return ""
}

// Bump increments one part ("major", "minor" or "patch") of a semantic version.
//
//	Bump("v1.2.3", "minor") == "v1.3.0"
//	Bump("v1.2.3", "major") == "v2.0.0"
//
// It returns an error for malformed versions or unknown parts.
func Bump(version, part string) (string, error) {
	// TODO: strings and strconv will help
	return "", nil
}

// Platform returns the operating system and architecture the program was
// built for, like "darwin/arm64" or "linux/amd64".
func Platform() string {
	// TODO: package runtime
	return ""
}
