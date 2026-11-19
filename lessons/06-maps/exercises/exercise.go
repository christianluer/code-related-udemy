//go:build !solution

// Package exercises contains the exercises for lesson 06.
//
// Replace each TODO, then run:  go run ./cmd/learn check 6
package exercises

// Favorites returns a map keyed by "last_first" name whose values are that
// person's favorite things (HOE #49):
//
//	bond_james        shaken, not stirred | martinis | fast cars
//	moneypenny_jenny  intelligence | literature | computer science
//	no_dr             cats | ice cream | sunsets
func Favorites() map[string][]string {
	// TODO
	return nil
}

// AddFavorite appends thing to person's favorites, creating the entry if
// the person is not in the map yet. (HOE #50)
func AddFavorite(m map[string][]string, person, thing string) {
	// TODO
}

// Remove deletes key from m and reports whether it was present. (HOE #51)
func Remove(m map[string][]string, key string) bool {
	// TODO: comma ok + delete
	return false
}

// WordFrequency counts how many times each word occurs, case-insensitively
// ("The" and "the" are the same word). (HOE #52)
func WordFrequency(words []string) map[string]int {
	// TODO
	return nil
}

// SortedKeys returns the keys of m in alphabetical order.
func SortedKeys(m map[string]int) []string {
	// TODO
	return nil
}

// TopWords returns the n most frequent words, most frequent first. Ties are
// broken alphabetically. If n > len(freq), return all words.
func TopWords(freq map[string]int, n int) []string {
	// TODO
	return nil
}
