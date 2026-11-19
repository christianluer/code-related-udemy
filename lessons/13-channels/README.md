# 13 · Channels & context

Course section: *Channels*, hands-on exercises #1 to #7.

Go's concurrency comes from **CSP**, Tony Hoare's *Communicating Sequential Processes*: independent
processes (goroutines) that interact only by passing messages (channels).

> **Do not communicate by sharing memory; instead, share memory by communicating.**

## Channel basics

```go
c := make(chan int)      // unbuffered
c <- 42                  // send
v := <-c                 // receive
b := make(chan int, 4)   // buffered: holds up to 4 values before a send blocks
close(c)                 // no more sends; receivers drain what's left, then get the zero value
```

**Channels block.** An unbuffered send waits until a receiver is ready. It's like runners in a relay race,
who must hand off the baton at the same moment. This code deadlocks:

```go
c := make(chan int)
c <- 42            // blocks forever: nobody is receiving
fmt.Println(<-c)
```

The fix is to send from another goroutine (`go func() { c <- 42 }()`) or to use a buffer of 1.

## Directional channels

Read the type left to right:

```go
chan int    // bidirectional: send and receive
chan<- int  // send-only
<-chan int  // receive-only
```

A general (bidirectional) channel converts to a specific one, but specific does not convert to general.
Use directional types in func params to document and enforce intent:
`func producer(out chan<- int)`, `func consumer(in <-chan int)`.

## range, select & comma ok

```go
for v := range c { }          // receive until c is closed

select {                      // wait on several channel ops at once
case v := <-even:
case v := <-odd:
case <-quit:
	return
case <-time.After(time.Second): // timeout
default:                      // don't block at all
}

v, ok := <-c                  // ok == false: c is closed and empty
```

## Patterns

- **Fan in**: take values from many channels and put them onto one.
- **Fan out**: split work across many goroutines (workers), optionally **throttled** to N workers.
- **Generator**: a func that returns a `<-chan T` and fills it from a goroutine.

## Context

Use `context.Context` to carry **cancellation**, **deadlines**, and **request-scoped values** across API
boundaries and goroutines. Servers use one per request.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ctx, cancel = context.WithTimeout(ctx, 2*time.Second)  // or WithDeadline(ctx, t)
ctx = context.WithValue(ctx, key, "value")

select {
case <-ctx.Done():  // closed when cancelled or timed out
	return ctx.Err() // context.Canceled or context.DeadlineExceeded
case out <- v:
}
```

Further reading: [go.dev/blog/context](https://go.dev/blog/context),
[Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide).

## Exercises

| Exercise | Course |
|---|---|
| `Generate` generator | HOE #6 |
| `Collect` range over a channel | HOE #3 |
| `Buffered` buffered channel | HOE #1 |
| `Receive` comma ok | HOE #5 |
| `FirstOf` / `WithTimeout` select | HOE #4 |
| `Producers` 10 goroutines × 10 values | HOE #7 |
| `FanIn` | Fan in |
| `FanOut` worker pool | Fan out |
| `CountUntilCancel` context | Context |
