//go:build solution

package exercises

import (
	"sync"
	"sync/atomic"
)

func RunAll(fns ...func()) {
	var wg sync.WaitGroup
	for _, f := range fns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
}

func (p *Person) Speak() string {
	return "Hi, I'm " + p.Name
}

func SaySomething(h Human) string {
	return h.Speak()
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.n, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.n)
}

func ParallelSum(nums []int, workers int) int {
	results := make([]int, workers)
	chunk := (len(nums) + workers - 1) / workers
	var wg sync.WaitGroup
	for w := range workers {
		lo, hi := w*chunk, min((w+1)*chunk, len(nums))
		if lo >= hi {
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, n := range nums[lo:hi] {
				results[w] += n
			}
		}()
	}
	wg.Wait()
	total := 0
	for _, r := range results {
		total += r
	}
	return total
}
