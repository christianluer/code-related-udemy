package exercises

// Person is used by the method and Stringer exercises. (HOE #61)
type Person struct {
	First string
	Age   int
}

// Shape is anything that has an Area method. (HOE #62)
type Shape interface {
	Area() float64
}

// Square is a rectangle, really; the course calls it a square.
type Square struct {
	Length float64
	Width  float64
}

// Circle has a radius.
type Circle struct {
	Radius float64
}

// User is a record in a database. (HOE #67)
type User struct {
	ID   int
	Name string
}

// Datastore is the behavior our code needs from a database. Because it's an
// interface, tests can pass in a fake (a "mock") instead of a real database.
type Datastore interface {
	GetUser(id int) (User, error)
}
