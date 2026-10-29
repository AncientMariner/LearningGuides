package creational

import "fmt"

// Factory Method: constructor functions (NewXxx) are the idiomatic Go form.

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string { return "Woof" }
func (c Cat) Speak() string { return "Meow" }

func Speak[T Animal](a T) string {
	return a.Speak()
}

func NewAnimal(kind string) (Animal, error) {
	switch kind {
	case "dog":
		return Dog{}, nil
	case "cat":
		return Cat{}, nil
	default:
		return nil, fmt.Errorf("unknown animal: %s", kind)
	}
}

func ExampleFactory() {
	dog, _ := NewAnimal("dog")
	cat, _ := NewAnimal("cat")
	Speak(&Dog{})
	Speak(&Cat{})
	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())
	// Output:
	// Woof
	// Meow
}
