package concurrency

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// Errgroup: run goroutines concurrently and collect the first error.
// Cancels the shared context on first failure so other goroutines can stop early.

func FetchAll(ctx context.Context, urls []string) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]string, len(urls))

	for i, url := range urls {
		i, url := i, url // capture loop vars (pre-Go 1.22)
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			// simulate fetch
			results[i] = fmt.Sprintf("body of %s", url)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func ExampleErrgroup() {
	results, err := FetchAll(context.Background(), []string{"a.com", "b.com"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, r := range results {
		fmt.Println(r)
	}
	// Output:
	// body of a.com
	// body of b.com
}
