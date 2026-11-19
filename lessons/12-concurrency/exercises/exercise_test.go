package exercises

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunAll(t *testing.T) {
	var done atomic.Int32
	slow := func() {
		time.Sleep(30 * time.Millisecond)
		done.Add(1)
	}
	start := time.Now()
	RunAll(slow, slow, slow, slow, slow)
	if n := done.Load(); n != 5 {
		t.Fatalf("RunAll returned before every func finished: %d of 5 done", n)
	}
	if el := time.Since(start); el > 120*time.Millisecond {
		t.Errorf("RunAll took %v; the funcs should run concurrently (about 30ms)", el)
	}
}

func TestMethodSets(t *testing.T) {
	var h any = Person{Name: "James"}
	if _, ok := h.(Human); ok {
		t.Error("a Person VALUE implements Human; Speak should have a pointer receiver")
	}
	p := &Person{Name: "James"}
	if got, want := SaySomething(p), "Hi, I'm James"; got != want {
		t.Errorf("SaySomething(&person) = %q, want %q", got, want)
	}
}

func hammer(inc func()) {
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				inc()
			}
		}()
	}
	wg.Wait()
}

func TestMutexCounter(t *testing.T) {
	var c MutexCounter
	hammer(c.Inc)
	if got := c.Value(); got != 10000 {
		t.Errorf("Value() = %d, want 10000", got)
	}
}

func TestAtomicCounter(t *testing.T) {
	var c AtomicCounter
	hammer(c.Inc)
	if got := c.Value(); got != 10000 {
		t.Errorf("Value() = %d, want 10000", got)
	}
}

func TestParallelSum(t *testing.T) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i + 1
	}
	for _, w := range []int{1, 3, 8, 2000} {
		if got := ParallelSum(nums, w); got != 500500 {
			t.Errorf("ParallelSum(1..1000, %d workers) = %d, want 500500", w, got)
		}
	}
	if got := ParallelSum(nil, 4); got != 0 {
		t.Errorf("ParallelSum(nil, 4) = %d, want 0", got)
	}
}
