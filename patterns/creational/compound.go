package creational

import "fmt"

// Compound (MVC): combines Decorator, Factory, and Observer.
// Model notifies observers; View renders; Controller mediates.

// Model

type ProductModel struct {
	Name      string
	observers []func(string)
}

func (m *ProductModel) SetName(name string) {
	m.Name = name
	for _, obs := range m.observers {
		obs(name)
	}
}

func (m *ProductModel) Subscribe(fn func(string)) {
	m.observers = append(m.observers, fn)
}

// View

type ProductView struct{}

func (v *ProductView) Render(name string) {
	fmt.Printf("View: product name is %q\n", name)
}

// Controller

type ProductController struct {
	model *ProductModel
	view  *ProductView
}

func NewProductController(model *ProductModel, view *ProductView) *ProductController {
	c := &ProductController{model: model, view: view}
	model.Subscribe(view.Render)
	return c
}

func (c *ProductController) UpdateName(name string) {
	c.model.SetName(name)
}

func ExampleCompound() {
	model := &ProductModel{}
	view := &ProductView{}
	controller := NewProductController(model, view)
	controller.UpdateName("Widget")
	// Output:
	// View: product name is "Widget"
}
