//go:build solution

package exercises

func NewPerson(first, last string, flavors ...string) Person {
	return Person{First: first, Last: last, Flavors: flavors}
}

func ByLastName(people []Person) map[string]Person {
	m := make(map[string]Person, len(people))
	for _, p := range people {
		m[p.Last] = p
	}
	return m
}

func NewVehicle(manufacturer, model string, doors int, color string, electric bool) Vehicle {
	return Vehicle{
		Engine: Engine{Electric: electric},
		Make:   manufacturer,
		Model:  model,
		Doors:  doors,
		Color:  color,
	}
}

func IsElectric(v Vehicle) bool {
	return v.Electric
}

func NewProfile(first string, drinks ...string) struct {
	First     string
	Friends   map[string]int
	FavDrinks []string
} {
	return struct {
		First     string
		Friends   map[string]int
		FavDrinks []string
	}{
		First:     first,
		Friends:   map[string]int{},
		FavDrinks: drinks,
	}
}

func FullNames(people []Person) []string {
	out := make([]string, 0, len(people))
	for _, p := range people {
		out = append(out, p.First+" "+p.Last)
	}
	return out
}
