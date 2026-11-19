// Goroutines and sync.WaitGroup.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("OS/ARCH:   ", runtime.GOOS, runtime.GOARCH)
	fmt.Println("CPUs:      ", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go foo(&wg)
	go bar(&wg)
	fmt.Println("Goroutines:", runtime.NumGoroutine(), "(main + foo + bar)")

	wg.Wait()
	fmt.Println("both finished; Goroutines:", runtime.NumGoroutine())

	// Launch many goroutines.
	start := time.Now()
	results := make([]int, 10)
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond) // pretend to work
			results[i] = i * i                // each goroutine owns its own slot: no race
		}()
	}
	wg.Wait()
	fmt.Println(results, "in", time.Since(start).Round(time.Millisecond), "(not 500ms!)")
}

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 3 {
		fmt.Println("foo:", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func bar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 3 {
		fmt.Println("bar:", i)
		time.Sleep(10 * time.Millisecond)
	}
}
