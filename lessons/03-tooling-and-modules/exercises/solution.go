//go:build solution

package exercises

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var count int

func Increment() int {
	count++
	return count
}

func IsExported(name string) bool {
	r, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(r)
}

func SHA256Hex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func Bump(version, part string) (string, error) {
	parts := strings.Split(strings.TrimPrefix(version, "v"), ".")
	if !strings.HasPrefix(version, "v") || len(parts) != 3 {
		return "", fmt.Errorf("malformed version %q", version)
	}
	var n [3]int
	for i, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil || v < 0 {
			return "", fmt.Errorf("malformed version %q", version)
		}
		n[i] = v
	}
	switch part {
	case "major":
		n = [3]int{n[0] + 1, 0, 0}
	case "minor":
		n = [3]int{n[0], n[1] + 1, 0}
	case "patch":
		n[2]++
	default:
		return "", fmt.Errorf("unknown part %q", part)
	}
	return fmt.Sprintf("v%d.%d.%d", n[0], n[1], n[2]), nil
}

func Platform() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}
