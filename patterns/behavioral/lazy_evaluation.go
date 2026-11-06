package behavioral

import (
	"fmt"
	"sync"
)

// Lazy Evaluation: defer expensive computation until first use.
// sync.Once for lazy init; closures for lazy values.

type LazyValue[T any] struct {
	compute func() T
	once    sync.Once
	value   T
}

func NewLazy[T any](compute func() T) *LazyValue[T] {
	return &LazyValue[T]{compute: compute}
}

func (l *LazyValue[T]) Get() T {
	l.once.Do(func() {
		l.value = l.compute()
		fmt.Println("computed")
	})
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
