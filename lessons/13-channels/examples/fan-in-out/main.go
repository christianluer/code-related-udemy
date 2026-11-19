// Fan in (many channels -> one) and fan out (one job queue -> many throttled workers).
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	fmt.Println("FAN IN")
	c := fanIn(boring("Joe"), boring("Ann"))
	for range 6 {
		fmt.Println(<-c)
	}

	fmt.Println("\nFAN OUT (throttled to 3 workers)")
	jobs := make(chan int)
	results := make(chan string)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				time.Sleep(time.Duration(rand.IntN(40)) * time.Millisecond)
				results <- fmt.Sprintf("worker %d did job %d", w, j)
			}
		}()
	}
	go func() {
		for j := range 9 {
			jobs <- j
		}
		close(jobs)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	for r := range results {
		fmt.Println(r)
	}
}

// boring is Rob Pike's generator: it returns a receive-only channel.
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.IntN(30)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-a:
				c <- s
			case s := <-b:
				c <- s
			}
		}
	}()
	return c
}
