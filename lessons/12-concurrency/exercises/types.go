package exercises

import "sync"

// Person is used by the method set exercise. (HOE #2)
type Person struct {
	Name string
}

// Human is implemented by anything with a Speak method.
type Human interface {
	Speak() string
}

// MutexCounter is a counter that is safe for concurrent use. (HOE #4)
type MutexCounter struct {
	mu sync.Mutex
	n  int
}

// AtomicCounter is a counter that is safe for concurrent use. (HOE #5)
type AtomicCounter struct {
	n int64
}
