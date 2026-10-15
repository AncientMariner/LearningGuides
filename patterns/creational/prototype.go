package creational

import "fmt"

// Prototype: Go structs copy by value; Clone() is used for pointer types
// or when deep copy of nested fields is needed.

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Age     int
	Address *Address
}

func (p *Person) Clone() *Person {
	addrCopy := *p.Address
	return &Person{
		Name:    p.Name,
		Age:     p.Age,
		Address: &addrCopy,
	}
}

func ExamplePrototype() {
	original := &Person{Name: "Alice", Age: 30, Address: &Address{Street: "1 Main St", City: "Springfield"}}
	clone := original.Clone()
	clone.Name = "Bob"
	clone.Address.City = "Shelbyville"

	fmt.Println(original.Name, original.Address.City)
	fmt.Println(clone.Name, clone.Address.City)
	// Output:
	// Alice Springfield
	// Bob Shelbyville
}
