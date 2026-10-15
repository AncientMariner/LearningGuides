package structural

import "fmt"

// Bridge: decouples abstraction from implementation so both can vary independently.

type Renderer interface {
	RenderCircle(radius float64)
}

type VectorRenderer struct{}
type RasterRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing vector circle with radius %.2f\n", radius)
}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing raster circle with radius %.2f\n", radius)
}

type Circle struct {
	renderer Renderer
	radius   float64
}

func NewCircle(r Renderer, radius float64) *Circle {
	return &Circle{renderer: r, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float64) {
	c.radius *= factor
}

func ExampleBridge() {
	c := NewCircle(&VectorRenderer{}, 5)
	c.Draw()
	c.Resize(2)
	c.Draw()
	// Output:
	// Drawing vector circle with radius 5.00
	// Drawing vector circle with radius 10.00
}
