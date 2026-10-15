package concurrency

import (
	"context"
	"fmt"
	"sync"
)

// Worker Pool: fixed N goroutines pull jobs from a shared channel.
// Bounds concurrency — prevents spawning unbounded goroutines for large workloads.

type Job struct {
	ID    int
	Input int
}

type Result struct {
	JobID  int
	Output int
}

func WorkerPool(ctx context.Context, workers int, jobs <-chan Job) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				select {
				case <-ctx.Done():
					return
				case results <- Result{JobID: job.ID, Output: job.Input * job.Input}:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

func ExampleWorkerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan Job, 5)
	for i := 1; i <= 5; i++ {
		jobs <- Job{ID: i, Input: i}
	}
	close(jobs)

	results := WorkerPool(ctx, 3, jobs)
	total := 0
	for r := range results {
		total += r.Output
	}
	fmt.Println(total) // 1+4+9+16+25 = 55
	// Output:
	// 55
}
