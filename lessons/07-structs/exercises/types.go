package exercises

// Person stores a person's name and favorite ice cream flavors. (HOE #53)
type Person struct {
	First   string
	Last    string
	Flavors []string
}

// Engine describes a vehicle's engine. (HOE #55)
type Engine struct {
	Electric bool
}

// Vehicle embeds an Engine, so Engine's fields are promoted. (HOE #55)
type Vehicle struct {
	Engine
	Make  string
	Model string
	Doors int
	Color string
}
