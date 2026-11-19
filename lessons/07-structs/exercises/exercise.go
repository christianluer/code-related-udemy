//go:build !solution

// Package exercises contains the exercises for lesson 07.
//
// Replace each TODO, then run:  go run ./cmd/learn check 7
package exercises

// NewPerson returns a Person with the given names and flavors. (HOE #53)
func NewPerson(first, last string, flavors ...string) Person {
	// TODO: composite literal
	return Person{}
}

// ByLastName stores each person in a map keyed by last name. (HOE #54)
func ByLastName(people []Person) map[string]Person {
	// TODO
	return nil
}

// NewVehicle builds a Vehicle with a composite literal. (HOE #55)
func NewVehicle(manufacturer, model string, doors int, color string, electric bool) Vehicle {
	// TODO
	return Vehicle{}
}

// IsElectric reports whether v's engine is electric.
// Use the PROMOTED field, not v.Engine.Electric.
func IsElectric(v Vehicle) bool {
	// TODO
	return false
}

// NewProfile returns an anonymous struct. Friends must be an initialized
// (non-nil) map so the caller can add to it. (HOE #56)
func NewProfile(first string, drinks ...string) struct {
	First     string
	Friends   map[string]int
	FavDrinks []string
} {
	// TODO: fill these in
	return struct {
		First     string
		Friends   map[string]int
		FavDrinks []string
	}{}
}

// FullNames returns "First Last" for every person, in order.
func FullNames(people []Person) []string {
	// TODO
	return nil
}
