// select: a switch for channel communication (a preview of lesson 13).
package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)

	go func() { time.Sleep(50 * time.Millisecond); fast <- "fast" }()
	go func() { time.Sleep(200 * time.Millisecond); slow <- "slow" }()

	for range 2 {
		select {
		case msg := <-fast:
			fmt.Println("received", msg)
		case msg := <-slow:
			fmt.Println("received", msg)
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}
}
