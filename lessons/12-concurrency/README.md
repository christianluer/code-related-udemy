# 12 · Concurrency

Course section: *Concurrency*, hands-on exercises #1 to #6.

> Concurrency is the composition of independently executing processes. Parallelism is the
> simultaneous execution of (possibly related) computations. **Concurrency is about *dealing* with
> lots of things at once. Parallelism is about *doing* lots of things at once.** (Rob Pike)

The first Intel dual-core CPU shipped in 2006, and Google started Go in 2007 to make the most of multiple cores.

## Goroutines

Put `go` in front of a func or method call and it runs concurrently:

```go
go doSomething()
go func() { ... }()
```

A goroutine is a function executing concurrently with other goroutines in the same address space.
Goroutines are cheap (a few KB of stack) and scheduled by the Go runtime onto OS threads.
`main` does **not** wait for them: when main returns, the program exits.

```go
runtime.NumCPU()        // logical CPUs
runtime.NumGoroutine()  // goroutines currently running
runtime.Gosched()       // yield the processor
```

## sync.WaitGroup

```go
var wg sync.WaitGroup
for i := range 10 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		work(i)
	}()
}
wg.Wait() // block until the counter reaches zero
```

Go 1.25 adds `wg.Go(func() { ... })`, which does the Add/Done for you.

## Race conditions

When two goroutines access the same variable and at least one writes, without synchronization,
you get a **race condition** and unpredictable results.

```bash
go run -race ./lessons/12-concurrency/examples/race
go test -race ./...
```

Fix them with one of these:

- a **mutex**: `var mu sync.Mutex; mu.Lock(); counter++; mu.Unlock()`. Use `sync.RWMutex` for many readers.
- **atomics**: `atomic.AddInt64(&counter, 1)`, or the typed `atomic.Int64{}.Add(1)`.
- **channels**: don't share memory, communicate (lesson 13).

## Method sets revisited

| receivers | values that have the method |
|---|---|
| `(t T)` | `T` and `*T` |
| `(t *T)` | `*T` only |

"The method set of a type determines the INTERFACES that the type implements" (Go spec).

## GOOS & GOARCH

`runtime.GOOS` and `runtime.GOARCH` tell you the platform you built for (HOE #6). Cross compile with
`GOOS=linux GOARCH=arm64 go build`.

## Exercises

| Exercise | Course |
|---|---|
| `RunAll` uses a WaitGroup | HOE #1 |
| `SaySomething` + `*Person` method sets | HOE #2 |
| `MutexCounter` fixes the race with a mutex | HOE #3, #4 |
| `AtomicCounter` fixes the race with atomic | HOE #5 |
| `ParallelSum` splits work over goroutines | concurrency vs parallelism |

`learn check 12` runs the tests with `-race`, so racy code fails even when the numbers look right.
