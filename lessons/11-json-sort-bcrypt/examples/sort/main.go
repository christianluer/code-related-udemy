// Sorting: built-ins, custom with sort.Slice, sort.Interface, and slices.SortFunc.
package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

	people := []person{{"James", 32}, {"Moneypenny", 27}, {"Q", 64}, {"M", 56}}

	sort.Sort(ByAge(people))
	fmt.Println("by age (sort.Interface):", people)

	sort.Slice(people, func(i, j int) bool { return people[i].First < people[j].First })
	fmt.Println("by name (sort.Slice):   ", people)

	slices.SortFunc(people, func(a, b person) int { return cmp.Compare(b.Age, a.Age) })
	fmt.Println("by age desc (slices):   ", people)
}
