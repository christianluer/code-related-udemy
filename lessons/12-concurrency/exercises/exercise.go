//go:build !solution

// Package exercises contains the exercises for lesson 12.
//
// Replace each TODO, then run:  go run ./cmd/learn check 12
// (the tests run with -race, so data races fail them)
package exercises

// RunAll runs every func in its own goroutine and returns only after ALL of
// them have finished. (HOE #1)
func RunAll(fns ...func()) {
	// TODO: sync.WaitGroup
}

// Speak returns "Hi, I'm <Name>". Use a POINTER receiver, so only *Person
// implements Human. (HOE #2)
//
// TODO: change the receiver.
func (p Person) Speak() string {
	return ""
}

// SaySomething calls Speak on any Human. (HOE #2)
func SaySomething(h Human) string {
	// TODO
	return ""
}

// Inc adds one to the counter. It is called from many goroutines at once.
// TODO: this is a data race. Protect it with c.mu. (HOE #3, #4)
func (c *MutexCounter) Inc() {
	c.n++
}

// Value returns the current count.
// TODO: reading also needs the lock.
func (c *MutexCounter) Value() int {
	return c.n
}

// Inc adds one to the counter.
// TODO: this is a data race. Fix it with package sync/atomic. (HOE #5)
func (c *AtomicCounter) Inc() {
	c.n++
}

// Value returns the current count.
// TODO: use an atomic load.
func (c *AtomicCounter) Value() int64 {
	return c.n
}

// ParallelSum adds up nums by splitting the work across `workers` goroutines.
// workers is at least 1 and may be larger than len(nums).
func ParallelSum(nums []int, workers int) int {
	// TODO
	return 0
}
