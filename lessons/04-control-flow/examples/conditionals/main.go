// if / else, comparison and logical operators, switch, init and the idioms.
package main

import (
	"fmt"
	"math/rand/v2"
)

func init() {
	fmt.Println("init runs before main: this is where initialization for my program occurs")
}

func main() {
	x := rand.IntN(251) // 0..250 inclusive
	fmt.Println("x =", x)

	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}

	// tagless switch: same logic, often easier to read
	switch {
	case x <= 100:
		fmt.Println("switch: low")
	case x <= 200:
		fmt.Println("switch: medium")
	default:
		fmt.Println("switch: high")
	}

	// switch on a value, multiple values per case, fallthrough
	switch d := x % 3; d {
	case 0:
		fmt.Println("divisible by 3")
		fallthrough
	case 1:
		fmt.Println("(0 falls through to here; 1 lands here directly)")
	default:
		fmt.Println("remainder 2")
	}

	fmt.Println("\nlogical operators:")
	t1, t2, f := true, true, false
	fmt.Println("true && true  =", t1 && t2)
	fmt.Println("true && false =", t1 && f)
	fmt.Println("true || true  =", t1 || t2)
	fmt.Println("true || false =", t1 || f)
	fmt.Println("!true         =", !t1)

	// statement; statement idiom
	if z := rand.IntN(5); z == 3 {
		fmt.Println("\nz is 3")
	} else {
		fmt.Println("\nz is", z, "(z is not visible after this if/else)")
	}

	// comma ok idiom
	m := map[string]int{"James": 42, "Moneypenny": 32}
	if age, ok := m["Q"]; !ok {
		fmt.Println("Q is not in the map, age is the zero value:", age)
	}

	// Is rand.IntN(n) inclusive? Find out empirically with n = 3.
	seen := map[int]bool{}
	for range 1000 {
		seen[rand.IntN(3)] = true
	}
	fmt.Println("rand.IntN(3) produced:", seen, "so 0 is included and 3 is excluded")
}
