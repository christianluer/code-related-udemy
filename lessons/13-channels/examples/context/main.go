// context: WithCancel, WithTimeout, WithDeadline and WithValue.
package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

func main() {
	// WithCancel: stop a generator when we have enough.
	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Print(n, " ")
		if n == 5 {
			break
		}
	}
	cancel() // the generator goroutine sees ctx.Done() and exits: no leak
	fmt.Println("\ncancelled:", ctx.Err())

	// WithTimeout: give slow work a budget.
	ctx, cancel = context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	fmt.Println("slow work:", slowWork(ctx))

	// WithDeadline is the same idea with an absolute time.
	dl, cancel2 := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
	defer cancel2()
	d, _ := dl.Deadline()
	fmt.Println("deadline in about", time.Until(d).Round(time.Minute))

	// WithValue: request-scoped data (user IDs, trace IDs...), not optional params.
	vctx := context.WithValue(context.Background(), ctxKey("userID"), 42)
	fmt.Println("userID from context:", vctx.Value(ctxKey("userID")))
}

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := 1; ; n++ {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func slowWork(ctx context.Context) error {
	select {
	case <-time.After(time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
