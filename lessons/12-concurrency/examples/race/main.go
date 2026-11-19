// A race condition, then two fixes: a mutex and package atomic.
//
// Run with the race detector to see it caught:
//
//	go run -race ./lessons/12-concurrency/examples/race
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const gs = 100

func main() {
	fmt.Println("racy:  ", racy(), "(should be", gs, "- run with -race!)")
	fmt.Println("mutex: ", withMutex())
	fmt.Println("atomic:", withAtomic())
}

func racy() int {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(gs)
	for range gs {
		go func() {
			v := counter
			runtime.Gosched() // yield: makes the race much more likely
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

func withMutex() int {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gs)
	for range gs {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

func withAtomic() int64 {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(gs)
	for range gs {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return atomic.LoadInt64(&counter)
}
