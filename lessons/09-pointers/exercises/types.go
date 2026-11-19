package exercises

// Person has a first name. (HOE #77)
type Person struct {
	First string
}

// Youngin must not be changed. (HOE #76)
type Youngin interface {
	Birthday()
	Years() int
}

// Kid should implement Youngin.
type Kid struct {
	Name string
	Age  int
}
