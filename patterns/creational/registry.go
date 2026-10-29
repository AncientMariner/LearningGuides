package creational

import "fmt"

// Registry: type-safe factory registry using generics.
// New types register themselves — no central switch to modify.
// Useful when you have many types or need external packages to register their own.

type Creator[T any] func() T

type Registry[T any] struct {
	creators map[string]Creator[T]
}

func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{creators: make(map[string]Creator[T])}
}

func (r *Registry[T]) Register(kind string, fn Creator[T]) {
	r.creators[kind] = fn
}

func (r *Registry[T]) Create(kind string) (T, error) {
	fn, ok := r.creators[kind]
	if !ok {
		var zero T
		return zero, fmt.Errorf("unknown kind: %s", kind)
	}
	return fn(), nil
}

func ExampleRegistry() {
	reg := NewRegistry[Animal]()
	reg.Register("dog", func() Animal { return Dog{} })
	reg.Register("cat", func() Animal { return Cat{} })

	a, _ := reg.Create("dog")
	b, _ := reg.Create("cat")
	fmt.Println(a.Speak())
	fmt.Println(b.Speak())

	_, err := reg.Create("fish")
	fmt.Println(err)
	// Output:
	// Woof
	// Meow
	// unknown kind: fish
}
