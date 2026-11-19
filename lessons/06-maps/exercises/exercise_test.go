package exercises

import (
	"maps"
	"slices"
	"testing"
)

func TestFavorites(t *testing.T) {
	want := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	got := Favorites()
	if !maps.EqualFunc(got, want, slices.Equal[[]string]) {
		t.Errorf("Favorites() = %v, want %v", got, want)
	}
}

func TestAddFavorite(t *testing.T) {
	m := map[string][]string{"no_dr": {"cats"}}
	AddFavorite(m, "no_dr", "sunsets")
	AddFavorite(m, "fleming_ian", "steaks")
	AddFavorite(m, "fleming_ian", "espionage")
	if got := m["no_dr"]; !slices.Equal(got, []string{"cats", "sunsets"}) {
		t.Errorf("m[no_dr] = %v, want [cats sunsets]", got)
	}
	if got := m["fleming_ian"]; !slices.Equal(got, []string{"steaks", "espionage"}) {
		t.Errorf("m[fleming_ian] = %v, want [steaks espionage]", got)
	}
}

func TestRemove(t *testing.T) {
	m := map[string][]string{"a": {"x"}, "b": nil}
	if !Remove(m, "a") {
		t.Error("Remove(m, \"a\") = false, want true")
	}
	if _, ok := m["a"]; ok {
		t.Error("\"a\" is still in the map")
	}
	if !Remove(m, "b") {
		t.Error("Remove(m, \"b\") = false, want true (the key exists even though its value is nil)")
	}
	if Remove(m, "zzz") {
		t.Error("Remove(m, \"zzz\") = true, want false")
	}
}

func TestWordFrequency(t *testing.T) {
	words := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "In", "MY"}
	got := WordFrequency(words)
	if got["my"] != 3 || got["in"] != 2 || got["father"] != 1 || len(got) != 8 {
		t.Errorf("WordFrequency = %v, want my:3 in:2 father:1 and 8 distinct words", got)
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]int{"pear": 1, "apple": 2, "fig": 3, "banana": 4}
	if got, want := SortedKeys(m), []string{"apple", "banana", "fig", "pear"}; !slices.Equal(got, want) {
		t.Errorf("SortedKeys = %v, want %v", got, want)
	}
}

func TestTopWords(t *testing.T) {
	freq := map[string]int{"the": 9, "a": 5, "of": 5, "gatsby": 2, "egg": 1}
	if got, want := TopWords(freq, 3), []string{"the", "a", "of"}; !slices.Equal(got, want) {
		t.Errorf("TopWords(freq, 3) = %v, want %v", got, want)
	}
	if got := TopWords(freq, 99); len(got) != 5 {
		t.Errorf("TopWords(freq, 99) returned %d words, want all 5", len(got))
	}
}
