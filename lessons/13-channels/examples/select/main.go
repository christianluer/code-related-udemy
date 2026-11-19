// select with even/odd/quit channels, as in the course.
package main

import "fmt"

func main() {
	even, odd, quit := make(chan int), make(chan int), make(chan bool)

	go send(even, odd, quit)

	for {
		select {
		case v := <-even:
			fmt.Println("from even:", v)
		case v := <-odd:
			fmt.Println("from odd: ", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("quit closed (comma ok):", i, ok)
				return
			}
			fmt.Println("from quit:", i)
		}
	}
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := range 10 {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
