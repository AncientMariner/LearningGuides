package behavioral

import (
	"fmt"
	"sort"
)

// Strategy: swap algorithms at runtime.
// In Go, passing a func is often enough; use an interface when the strategy
// carries state or has multiple methods.

type SortStrategy func([]int)

func BubbleSort(data []int) {
	n := len(data)
	for i := range n {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func StdSort(data []int) { sort.Ints(data) }

type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter { return &Sorter{strategy: strategy} }

func (s *Sorter) SetStrategy(strategy SortStrategy) { s.strategy = strategy }

func (s *Sorter) Sort(data []int) {
	s.strategy(data)
	fmt.Println("sorted:", data)
}

func ExampleStrategy() {
	sorter := NewSorter(BubbleSort)
	sorter.Sort([]int{3, 1, 2})
	sorter.SetStrategy(StdSort)
	sorter.Sort([]int{9, 5, 7})
	// Output:
	// sorted: [1 2 3]
	// sorted: [5 7 9]
}
