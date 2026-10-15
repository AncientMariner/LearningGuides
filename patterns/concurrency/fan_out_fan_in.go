package concurrency

import (
	"context"
	"sync"
)

// Fan-out / Fan-in: distribute work across N goroutines, merge results into one channel.
// Fan-out: multiple goroutines read from the same input channel.
// Fan-in: merge multiple output channels into one.

func FanOut(ctx context.Context, in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, n)
	for i := range n {
		outs[i] = Square(ctx, in)
	}
	return outs
}

func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	forward := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			select {
			case merged <- v:
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go forward(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}

func ExampleFanOutFanIn() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := Generate(ctx, 1, 2, 3, 4, 5, 6)
	// fan out to 3 workers, fan in results
	workers := FanOut(ctx, in, 3)
	out := FanIn(ctx, workers...)

	results := make([]int, 0, 6)
	for v := range out {
		results = append(results, v)
	}
	// results contains squares of 1-6 in some order
	_ = results
}
