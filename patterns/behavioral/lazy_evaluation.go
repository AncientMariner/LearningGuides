package behavioral

import "fmt"

// Lazy Evaluation: defer expensive computation until first use.
// sync.Once for lazy init; closures for lazy values.

type LazyValue[T any] struct {
	compute func() T
	value   T
	done    bool
}

func NewLazy[T any](compute func() T) *LazyValue[T] {
	return &LazyValue[T]{compute: compute}
}

func (l *LazyValue[T]) Get() T {
	if !l.done {
		l.value = l.compute()
		l.done = true
		fmt.Println("computed")
	}
	return l.value
}

func ExampleLazyEvaluation() {
	lv := NewLazy(func() int { return 42 })
	fmt.Println(lv.Get())
	fmt.Println(lv.Get()) // not computed again
	// Output:
	// computed
	// 42
	// 42
}
