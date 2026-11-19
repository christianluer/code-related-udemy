//go:build solution

package exercises

import (
	"sort"
	"strings"
)

func Favorites() map[string][]string {
	return map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}
}

func AddFavorite(m map[string][]string, person, thing string) {
	m[person] = append(m[person], thing)
}

func Remove(m map[string][]string, key string) bool {
	_, ok := m[key]
	delete(m, key)
	return ok
}

func WordFrequency(words []string) map[string]int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[strings.ToLower(w)]++
	}
	return counts
}

func SortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func TopWords(freq map[string]int, n int) []string {
	keys := SortedKeys(freq)
	sort.SliceStable(keys, func(i, j int) bool { return freq[keys[i]] > freq[keys[j]] })
	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}
