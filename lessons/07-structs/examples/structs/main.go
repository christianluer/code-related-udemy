// Structs: literals, embedding & promotion, anonymous structs, composition.
package main

import "fmt"

type person struct {
	first   string
	last    string
	age     int
	flavors []string
}

type secretAgent struct {
	person        // embedded
	ltk    bool   // license to kill
	first  string // shadows person.first
}

// Engine and Car show method promotion through composition.
type Engine struct{ HP int }

func (e Engine) Start() string { return fmt.Sprintf("engine with %d hp started", e.HP) }

type Car struct {
	Engine
	Model string
}

// Duration-style named types self-document code.
type celsius float64

func main() {
	p1 := person{first: "James", last: "Bond", age: 32, flavors: []string{"vanilla", "mint"}}
	p2 := person{first: "Jenny", last: "Moneypenny", age: 27}

	fmt.Println(p1)
	fmt.Printf("%+v\n", p2)
	fmt.Printf("%#v\n\n", p2)

	for _, f := range p1.flavors {
		fmt.Println(p1.first, "likes", f)
	}

	sa := secretAgent{person: p1, ltk: true, first: "007"}
	fmt.Println("\npromoted last:", sa.last, "| shadowed first:", sa.first, "| inner first:", sa.person.first)

	c := Car{Engine: Engine{HP: 300}, Model: "DB5"}
	fmt.Println(c.Model, "->", c.Start()) // Start is promoted from Engine

	anon := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "James",
		friends:   map[string]int{"Moneypenny": 555, "Q": 777},
		favDrinks: []string{"martini", "water"},
	}
	fmt.Printf("\nanonymous: %+v\n", anon)

	// Map of structs keyed by last name.
	byLast := map[string]person{p1.last: p1, p2.last: p2}
	fmt.Println("Moneypenny's age:", byLast["Moneypenny"].age)

	var t celsius = 21.5
	fmt.Printf("%v is of type %T\n", t, t)

	// Comparing: structs of comparable fields can use ==.
	type point struct{ X, Y int }
	fmt.Println("point{1,2} == point{1,2}:", point{1, 2} == point{1, 2})
}
