//go:build solution

package exercises

func AddOne(p *int) {
	*p++
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func ChangeFirstValue(p Person, s string) Person {
	p.First = s
	return p
}

func ChangeFirstPointer(p *Person, s string) {
	p.First = s
}

// Birthday needs a pointer receiver to mutate; Years follows for consistency.
func (k *Kid) Birthday() {
	k.Age++
}

func (k *Kid) Years() int {
	return k.Age
}

func NewYoungin(name string, age int) Youngin {
	return &Kid{Name: name, Age: age}
}

func NewInt(v int) *int {
	x := v
	return &x
}

func DoubleAll(xs []int) {
	for i := range xs {
		xs[i] *= 2
	}
}
