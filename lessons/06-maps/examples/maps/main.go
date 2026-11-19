// Maps: create, add, access, len, range, delete, comma ok, sorted output.
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}
	fmt.Println("the age of Henry was", am["Henry"])
	fmt.Println(am, "len:", len(am))

	am["Lucas"] = 28 // add
	am["Todd"] = 43  // update

	fmt.Println("\nrange (random order):")
	for k, v := range am {
		fmt.Println(" ", k, v)
	}

	fmt.Println("\nsorted by key:")
	for _, k := range slices.Sorted(maps.Keys(am)) {
		fmt.Println(" ", k, am[k])
	}

	delete(am, "Lucas")
	delete(am, "I don't exist") // no error

	fmt.Println("\nlooking up a missing key gives the zero value:", am["Lucas"])
	if v, ok := am["Lucas"]; !ok {
		fmt.Println("comma ok says Lucas is gone; v =", v)
	}
	if v, ok := am["Todd"]; ok {
		fmt.Println("Todd is present with", v)
	}

	// map[string][]string
	favs := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
	}
	favs["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for _, who := range slices.Sorted(maps.Keys(favs)) {
		fmt.Println(who)
		for i, thing := range favs[who] {
			fmt.Printf("\t%d %s\n", i, thing)
		}
	}
}
