package exercises

import (
	"context"
	"slices"
	"sync/atomic"
	"testing"
	"time"
)

// recvTimeout receives from c or fails the test after a second.
func recvTimeout(t *testing.T, c <-chan int) (int, bool) {
	t.Helper()
	select {
	case v, ok := <-c:
		return v, ok
	case <-time.After(time.Second):
		t.Fatal("timed out waiting on the channel (is something blocked?)")
		return 0, false
	}
}

func drain(t *testing.T, c <-chan int) []int {
	t.Helper()
	var out []int
	for {
		v, ok := recvTimeout(t, c)
		if !ok {
			return out
		}
		out = append(out, v)
	}
}

func TestGenerate(t *testing.T) {
	if got, want := drain(t, Generate(5)), []int{0, 1, 2, 3, 4}; !slices.Equal(got, want) {
		t.Errorf("Generate(5) yielded %v, want %v", got, want)
	}
	if got := drain(t, Generate(100)); len(got) != 100 {
		t.Errorf("Generate(100) yielded %d values", len(got))
	}
}

func TestCollect(t *testing.T) {
	c := make(chan int, 3)
	c <- 7
	c <- 8
	c <- 9
	close(c)
	if got := Collect(c); !slices.Equal(got, []int{7, 8, 9}) {
		t.Errorf("Collect = %v, want [7 8 9]", got)
	}
}

func TestBuffered(t *testing.T) {
	c := Buffered(4, 5, 6)
	if cap(c) != 3 {
		t.Errorf("cap = %d, want 3", cap(c))
	}
	if got := drain(t, c); !slices.Equal(got, []int{4, 5, 6}) {
		t.Errorf("Buffered(4, 5, 6) yielded %v", got)
	}
}

func TestReceive(t *testing.T) {
	c := make(chan int, 1)
	c <- 42
	if v, ok := Receive(c); v != 42 || !ok {
		t.Errorf("Receive(open chan with 42) = (%d, %v), want (42, true)", v, ok)
	}
	close(c)
	if v, ok := Receive(c); v != 0 || ok {
		t.Errorf("Receive(closed chan) = (%d, %v), want (0, false)", v, ok)
	}
}

func TestFirstOf(t *testing.T) {
	never := make(chan string)
	soon := make(chan string, 1)
	soon <- "b wins"
	if got := FirstOf(never, soon); got != "b wins" {
		t.Errorf("FirstOf = %q, want %q", got, "b wins")
	}
	a := make(chan string, 1)
	a <- "a wins"
	if got := FirstOf(a, never); got != "a wins" {
		t.Errorf("FirstOf = %q, want %q", got, "a wins")
	}
}

func TestWithTimeout(t *testing.T) {
	c := make(chan int, 1)
	c <- 5
	if v, ok := WithTimeout(c, time.Second); v != 5 || !ok {
		t.Errorf("WithTimeout(ready) = (%d, %v), want (5, true)", v, ok)
	}
	start := time.Now()
	if _, ok := WithTimeout(make(chan int), 30*time.Millisecond); ok {
		t.Error("WithTimeout(empty) should report false")
	}
	if el := time.Since(start); el < 30*time.Millisecond {
		t.Errorf("WithTimeout gave up after %v, before the 30ms timeout", el)
	}
}

func TestProducers(t *testing.T) {
	got := Producers(10, 10)
	if len(got) != 100 {
		t.Fatalf("Producers(10, 10) returned %d values, want 100", len(got))
	}
	counts := map[int]int{}
	for _, v := range got {
		counts[v]++
	}
	for i := range 10 {
		if counts[i] != 10 {
			t.Errorf("value %d appeared %d times, want 10", i, counts[i])
		}
	}
}

func filled(vals ...int) <-chan int {
	c := make(chan int, len(vals))
	for _, v := range vals {
		c <- v
	}
	close(c)
	return c
}

func TestFanIn(t *testing.T) {
	got := drain(t, FanIn(filled(1, 2, 3), filled(10, 20), filled()))
	slices.Sort(got)
	if want := []int{1, 2, 3, 10, 20}; !slices.Equal(got, want) {
		t.Errorf("FanIn yielded %v, want %v", got, want)
	}
}

func TestFanOut(t *testing.T) {
	var running, peak atomic.Int32
	slowSquare := func(n int) int {
		cur := running.Add(1)
		for {
			p := peak.Load()
			if cur <= p || peak.CompareAndSwap(p, cur) {
				break
			}
		}
		time.Sleep(10 * time.Millisecond)
		running.Add(-1)
		return n * n
	}
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := FanOut(jobs, 4, slowSquare)
	if want := []int{1, 4, 9, 16, 25, 36, 49, 64}; !slices.Equal(got, want) {
		t.Errorf("FanOut = %v, want %v", got, want)
	}
	if p := peak.Load(); p < 2 || p > 4 {
		t.Errorf("peak concurrency was %d, want between 2 and 4 workers", p)
	}
}

func TestCountUntilCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	c := CountUntilCancel(ctx)
	for want := range 5 {
		v, ok := recvTimeout(t, c)
		if !ok || v != want {
			t.Fatalf("got (%d, %v), want (%d, true)", v, ok, want)
		}
	}
	cancel()
	deadline := time.After(time.Second)
	for {
		select {
		case _, ok := <-c:
			if !ok {
				return
			}
		case <-deadline:
			t.Fatal("channel was not closed after cancel()")
		}
	}
}
