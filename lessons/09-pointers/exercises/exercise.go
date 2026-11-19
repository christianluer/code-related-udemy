//go:build !solution

// Package exercises contains the exercises for lesson 09.
//
// Replace each TODO, then run:  go run ./cmd/learn check 9
package exercises

// AddOne increments the int that p points to. (HOE #75: dereference)
func AddOne(p *int) {
	// TODO
}

// Swap exchanges the values that a and b point to.
func Swap(a, b *int) {
	// TODO
}

// ChangeFirstValue uses VALUE semantics: it returns a changed copy of p
// and leaves the caller's Person alone. (HOE #77)
func ChangeFirstValue(p Person, s string) Person {
	// TODO
	return Person{}
}

// ChangeFirstPointer uses POINTER semantics: it changes the caller's
// Person and returns nothing. (HOE #77)
func ChangeFirstPointer(p *Person, s string) {
	// TODO
}

// Birthday adds one to the kid's age. It must change the caller's Kid,
// so think about which receiver it needs. (HOE #76)
//
// TODO: fix the receiver, and keep both methods consistent.
func (k Kid) Birthday() {
	k.Age++
}

// Years returns the kid's age.
func (k Kid) Years() int {
	return k.Age
}

// NewYoungin returns a Youngin backed by a Kid. Choose carefully between
// returning Kid{...} and &Kid{...}. (HOE #76)
func NewYoungin(name string, age int) Youngin {
	// TODO
	return nil
}

// NewInt returns a pointer to a NEW int holding v. Changing *NewInt(v)
// must not affect anything else. (escape analysis)
func NewInt(v int) *int {
	// TODO
	return nil
}

// DoubleAll doubles every element of xs IN PLACE. No return value needed:
// the slice header is copied but the underlying array is shared.
func DoubleAll(xs []int) {
	// TODO
}
