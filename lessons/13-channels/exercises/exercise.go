//go:build !solution

// Package exercises contains the exercises for lesson 13.
//
// Replace each TODO, then run:  go run ./cmd/learn check 13
//
// The stubs return closed channels so the tests fail instead of hanging.
package exercises

import (
	"context"
	"time"
)

func closed() chan int {
	c := make(chan int)
	close(c)
	return c
}

// Generate returns a channel that yields 0, 1, ... n-1 and is then closed.
// The values are sent from a goroutine. (HOE #6)
func Generate(n int) <-chan int {
	// TODO
	return closed()
}

// Collect receives every value from c until it is closed. (HOE #3)
func Collect(c <-chan int) []int {
	// TODO: for range over the channel
	return nil
}

// Buffered returns a closed channel with capacity len(values) that already
// holds every value. Don't use a goroutine. (HOE #1)
func Buffered(values ...int) <-chan int {
	// TODO
	return closed()
}

// Receive does a single receive from c using the comma ok idiom. (HOE #5)
func Receive(c <-chan int) (int, bool) {
	// TODO
	return 0, false
}

// FirstOf returns whichever string arrives first on a or b. (HOE #4)
func FirstOf(a, b <-chan string) string {
	// TODO: select
	return ""
}

// WithTimeout returns the next value from c, or (0, false) if nothing
// arrives within d.
func WithTimeout(c <-chan int, d time.Duration) (int, bool) {
	// TODO: select + time.After
	return 0, false
}

// Producers launches g goroutines that each send the numbers 0..n-1 on ONE
// shared channel, and returns all g*n values (in any order). (HOE #7)
func Producers(g, n int) []int {
	// TODO
	return nil
}

// FanIn merges all inputs onto one channel, which is closed once every
// input is closed. (Fan in)
func FanIn(inputs ...<-chan int) <-chan int {
	// TODO
	return closed()
}

// FanOut computes f(jobs[i]) for every job using `workers` goroutines and
// returns the results in the same order as jobs. (Fan out)
func FanOut(jobs []int, workers int, f func(int) int) []int {
	// TODO
	return nil
}

// CountUntilCancel returns a channel yielding 0, 1, 2, ... forever, until ctx
// is cancelled. Then the channel must be closed and the goroutine must exit.
// (Context)
func CountUntilCancel(ctx context.Context) <-chan int {
	// TODO
	return closed()
}
