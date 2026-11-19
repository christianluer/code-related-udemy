// Every flavor of for loop, break, continue, nesting, range and modulus.
package main

import "fmt"

func main() {
	fmt.Print("init; cond; post: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	fmt.Print("\ncondition only:   ")
	n := 1
	for n < 100 {
		n *= 2
		fmt.Print(n, " ")
	}

	fmt.Print("\ninfinite + break: ")
	k := 0
	for {
		if k == 4 {
			break
		}
		fmt.Print(k, " ")
		k++
	}

	fmt.Print("\ncontinue (odds):  ")
	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}

	fmt.Println("\n\nnested loops:")
	for i := range 3 {
		fmt.Printf("outer %d:", i)
		for j := range 3 {
			fmt.Printf(" (%d,%d)", i, j)
		}
		fmt.Println()
	}

	fmt.Println("\nrange over a slice:")
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index %d value %d\n", i, v)
	}

	fmt.Println("\nrange over a map (random order!):")
	m := map[string]int{"James": 42, "Moneypenny": 32}
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("\nrange over a string gives runes:")
	for i, r := range "héllo" {
		fmt.Printf("%d:%c ", i, r)
	}
	fmt.Println()

	// Labeled break: exit an outer loop from inside an inner one.
outer:
	for i := range 5 {
		for j := range 5 {
			if i*j == 6 {
				fmt.Println("\nfound i*j == 6 at", i, j)
				break outer
			}
		}
	}
}
