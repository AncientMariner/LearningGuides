package creational

import "fmt"

// Abstract Factory: interface that returns interfaces.
// Allows swapping entire families of related objects (e.g. mock vs real).

type Button interface {
	Render() string
}

type Checkbox interface {
	Render() string
}

type UIFactory interface {
	NewButton() Button
	NewCheckbox() Checkbox
}

// Dark theme

type DarkButton struct{}
type DarkCheckbox struct{}

func (b DarkButton) Render() string   { return "dark button" }
func (c DarkCheckbox) Render() string { return "dark checkbox" }

type DarkUIFactory struct{}

func (f DarkUIFactory) NewButton() Button     { return DarkButton{} }
func (f DarkUIFactory) NewCheckbox() Checkbox { return DarkCheckbox{} }

// Light theme

type LightButton struct{}
type LightCheckbox struct{}

func (b LightButton) Render() string   { return "light button" }
func (c LightCheckbox) Render() string { return "light checkbox" }

type LightUIFactory struct{}

func (f LightUIFactory) NewButton() Button     { return LightButton{} }
func (f LightUIFactory) NewCheckbox() Checkbox { return LightCheckbox{} }

func ExampleAbstractFactory() {
	var factory UIFactory = DarkUIFactory{}
	fmt.Println(factory.NewButton().Render())
	fmt.Println(factory.NewCheckbox().Render())
	// Output:
	// dark button
	// dark checkbox
}
