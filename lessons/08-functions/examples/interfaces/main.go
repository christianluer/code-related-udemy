// Methods, interfaces & polymorphism, the Stringer interface and a logging wrapper.
package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type shape interface {
	area() float64
}

type square struct{ length, width float64 }
type circle struct{ radius float64 }

func (s square) area() float64 { return s.length * s.width }
func (c circle) area() float64 { return math.Pi * math.Pow(c.radius, 2) }

func info(s shape) {
	fmt.Printf("%T has area %.2f\n", s, s.area())
}

// Sayer: Dog and Cat both implement it, so both are also of type Sayer.
type Sayer interface{ Say() string }
type Dog struct{}
type Cat struct{}

func (Dog) Say() string { return "Woof!" }
func (Cat) Say() string { return "Meow!" }

// Stringer: control how fmt prints your type.
type book struct{ title string }

func (b book) String() string { return "📖 " + b.title }

type count int

func (c count) String() string { return "count is " + strconv.Itoa(int(c)) }

// logInfo is a wrapper that accepts anything that is a fmt.Stringer.
func logInfo(s fmt.Stringer) { log.Println("LOG FROM 138:", s.String()) }

func main() {
	info(square{2, 3})
	info(circle{1})

	for _, s := range []Sayer{Dog{}, Cat{}} {
		fmt.Println(s.Say())
	}

	b := book{"The Go Programming Language"}
	var n count = 42
	fmt.Println(b)
	fmt.Println(n)
	logInfo(b)
	logInfo(n)

	// Type assertion & type switch: recover the concrete type.
	var s shape = circle{2}
	if c, ok := s.(circle); ok {
		fmt.Println("it's a circle with radius", c.radius)
	}
	for _, v := range []any{42, "hi", b, 3.14} {
		switch x := v.(type) {
		case int:
			fmt.Println("int", x)
		case fmt.Stringer:
			fmt.Println("Stringer", x)
		default:
			fmt.Printf("something else: %T\n", x)
		}
	}
}
