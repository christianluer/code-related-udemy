// Arrays and slices: literals, range, append, slicing, deleting, make, 2D.
package main

import "fmt"

func main() {
	// Arrays: fixed size, size is part of the type.
	flavors := [...]string{"Almond Biscotti Café", "Banana Pudding", "Coffee (GF)", "Lemon Bar"}
	fmt.Printf("%v\nlen %d, type %T\n\n", flavors, len(flavors), flavors)

	// Slice literal + range.
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range xi {
		fmt.Printf("%d:%d ", i, v)
	}
	fmt.Printf("\nlen %d cap %d type %T\n\n", len(xi), cap(xi), xi)

	// Slicing.
	fmt.Println("xi[:5] ", xi[:5])
	fmt.Println("xi[5:] ", xi[5:])
	fmt.Println("xi[2:7]", xi[2:7])
	fmt.Println("xi[1:6]", xi[1:6])

	// Append: one value, several values, a whole slice.
	xi = append(xi, 52)
	xi = append(xi, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	xi = append(xi, y...)
	fmt.Println("\nappended:", xi)

	// Delete 45, 46, 47 with append + slicing.
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println("deleted: ", x)

	// make: watch the capacity grow when we exceed it.
	s := make([]int, 0, 2)
	for i := range 6 {
		s = append(s, i)
		fmt.Printf("len %d cap %d %v\n", len(s), cap(s), s)
	}

	// Multidimensional slices.
	records := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."},
	}
	for i, rec := range records {
		fmt.Println("record", i)
		for j, field := range rec {
			fmt.Printf("\t%d: %s\n", j, field)
		}
	}
}
