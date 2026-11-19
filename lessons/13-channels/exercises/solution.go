//go:build solution

package exercises

import (
	"context"
	"sync"
	"time"
)

func Generate(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := range n {
			c <- i
		}
	}()
	return c
}

func Collect(c <-chan int) []int {
	var out []int
	for v := range c {
		out = append(out, v)
	}
	return out
}

func Buffered(values ...int) <-chan int {
	c := make(chan int, len(values))
	for _, v := range values {
		c <- v
	}
	close(c)
	return c
}

func Receive(c <-chan int) (int, bool) {
	v, ok := <-c
	return v, ok
}

func FirstOf(a, b <-chan string) string {
	select {
	case s := <-a:
		return s
	case s := <-b:
		return s
	}
}

func WithTimeout(c <-chan int, d time.Duration) (int, bool) {
	select {
	case v := <-c:
		return v, true
	case <-time.After(d):
		return 0, false
	}
}

func Producers(g, n int) []int {
	c := make(chan int)
	var wg sync.WaitGroup
	for range g {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range n {
				c <- i
			}
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return Collect(c)
}

func FanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func FanOut(jobs []int, workers int, f func(int) int) []int {
	results := make([]int, len(jobs))
	idx := make(chan int)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idx {
				results[i] = f(jobs[i])
			}
		}()
	}
	for i := range jobs {
		idx <- i
	}
	close(idx)
	wg.Wait()
	return results
}

func CountUntilCancel(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
