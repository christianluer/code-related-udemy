// Package dog turns puppy sounds into grown-up dog sounds.
package dog

import "strings"

// WhenGrownUp makes any sound LOUD.
func WhenGrownUp(s string) string {
	return strings.ToUpper(s)
}
