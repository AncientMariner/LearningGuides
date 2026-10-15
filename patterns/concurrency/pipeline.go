package concurrency

import "context"

// Pipeline: chain of goroutines connected by channels.
// Each stage reads from an input channel, transforms, and writes to an output channel.
// Closing the done channel (via context) tears down all stages cleanly.

func Generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func ExamplePipeline() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := Generate(ctx, 2, 3, 4)
	squares := Square(ctx, nums)

	for v := range squares {
		_ = v // 4, 9, 16
	}
}
