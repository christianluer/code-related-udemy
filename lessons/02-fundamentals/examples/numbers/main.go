// Numeral systems, Printf verbs, signed/unsigned limits, iota and bitwise ops.
package main

import (
	"fmt"
	"math"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	for _, n := range []int{747, 911, 90210} {
		fmt.Printf("%6d  bin %-18b  hex %#-7x  oct %#o\n", n, n, n, n)
	}

	fmt.Println("\nLimits:")
	fmt.Println("int8 :", math.MinInt8, "to", math.MaxInt8)
	fmt.Println("uint8:", 0, "to", math.MaxUint8)
	fmt.Println("int64:", math.MinInt64, "to", math.MaxInt64)

	var u uint8 = 255
	u++ // wraps around
	fmt.Println("255 + 1 as uint8 =", u)

	fmt.Println("\nByte sizes with iota:")
	fmt.Printf("KB %d\t%b\nMB %d\t%b\nGB %d\t%b\nTB %d\n", KB, KB, MB, MB, GB, GB, TB)

	fmt.Println("\nBitwise:")
	a, b := 0b1100, 0b1010
	fmt.Printf("a      %04b\n", a)
	fmt.Printf("b      %04b\n", b)
	fmt.Printf("a & b  %04b  AND\n", a&b)
	fmt.Printf("a | b  %04b  OR\n", a|b)
	fmt.Printf("a ^ b  %04b  XOR\n", a^b)
	fmt.Printf("a &^ b %04b  AND NOT\n", a&^b)
	fmt.Printf("1 << 3 = %d, 64 >> 2 = %d\n", 1<<3, 64>>2)
}
