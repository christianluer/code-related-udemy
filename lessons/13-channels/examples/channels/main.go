// Channels: blocking, buffering, directions, range, close and comma ok.
package main

import "fmt"

func main() {
	// Unbuffered: the send must happen in another goroutine.
	c := make(chan int)
	go func() { c <- 42 }()
	fmt.Println("received", <-c)

	// Buffered: a send succeeds without a receiver while there's room.
	b := make(chan int, 2)
	b <- 1
	b <- 2
	fmt.Println("buffered:", <-b, <-b, "len/cap:", len(b), cap(b))

	// Directional channels in func params.
	nums := make(chan int)
	go send(nums)
	receive(nums)

	// The types.
	var general chan int = make(chan int)
	var sendOnly chan<- int = general // general -> specific: ok
	var recvOnly <-chan int = general
	fmt.Printf("%T | %T | %T\n", general, sendOnly, recvOnly)
	// general = sendOnly // compile error: specific -> general

	// Comma ok.
	done := make(chan int, 1)
	done <- 7
	close(done)
	v, ok := <-done
	fmt.Println("v, ok =", v, ok)
	v, ok = <-done
	fmt.Println("v, ok =", v, ok, "(closed and drained)")
}

func send(out chan<- int) {
	for i := range 5 {
		out <- i
	}
	close(out)
}

func receive(in <-chan int) {
	for v := range in { // stops when send closes the channel
		fmt.Println("range got", v)
	}
}
