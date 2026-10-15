package behavioral

import (
	"fmt"
	"iter"
)

// Iterator: Go's range keyword and iter.Seq (1.23+) are the idiomatic forms.
// Use a custom iterator when you need lazy evaluation over a custom collection.

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) { s.items = append(s.items, item) }
func (s *Stack[T]) Len() int    { return len(s.items) }

// All returns an iter.Seq that yields items top-to-bottom.
func (s *Stack[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s.items) - 1; i >= 0; i-- {
			if !yield(s.items[i]) {
				return
			}
		}
	}
}

func ExampleIterator() {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Output:
	// 3
	// 2
	// 1
}
