// Pointers: &, *, pass by value, value vs pointer semantics, method sets.
package main

import "fmt"

type person struct{ first string }

func changeValue(p person, s string) person { p.first = s; return p }
func changePointer(p *person, s string)     { p.first = s }

type youngin interface{ birthday() }

type kid struct {
	name string
	age  int
}

func (k *kid) birthday() { k.age++ }

func main() {
	x := 42
	p := &x
	fmt.Printf("x=%v (%T)  &x=%v (%T)  *p=%v\n", x, x, p, p, *p)
	*p = 21
	fmt.Println("after *p = 21, x =", x)

	var np *int
	fmt.Println("zero value of a pointer:", np, "is nil?", np == nil)

	// Pass by value: a copy...
	a := 1
	addOne(a)
	fmt.Println("addOne(a) -> a =", a)
	// ...but the copy can be an address.
	addOnePtr(&a)
	fmt.Println("addOnePtr(&a) -> a =", a)

	// Slices and maps: the copied header still refers to the same data.
	xs := []int{1, 2, 3}
	m := map[string]int{"a": 1}
	mutate(xs, m)
	fmt.Println("after mutate:", xs, m)

	// Value semantics return a new value; pointer semantics share memory.
	p1 := person{"James"}
	p2 := changeValue(p1, "Jenny")
	fmt.Println("value semantics:  ", p1.first, "->", p2.first)
	changePointer(&p1, "Q")
	fmt.Println("pointer semantics:", p1.first)

	// Method sets: only *kid has birthday(), so only *kid is a youngin.
	k := kid{"Henry", 15}
	var y youngin = &k
	// var y youngin = k // compile error: kid does not implement youngin (pointer receiver)
	y.birthday()
	k.birthday() // fine: k is addressable, Go calls (&k).birthday()
	fmt.Println(k.name, "is now", k.age)
}

func addOne(v int)     { v++ }
func addOnePtr(v *int) { *v++ }

func mutate(xs []int, m map[string]int) {
	xs[0] = 100
	m["b"] = 2
}
