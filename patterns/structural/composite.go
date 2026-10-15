package structural

import "fmt"

// Composite: treats individual objects and compositions uniformly.
// A slice of interface is the idiomatic Go form.

type Component interface {
	Name() string
	Display(indent string)
}

type Leaf struct {
	name string
}

func NewLeaf(name string) *Leaf { return &Leaf{name: name} }

func (l *Leaf) Name() string              { return l.name }
func (l *Leaf) Display(indent string)     { fmt.Println(indent + l.name) }

type Composite struct {
	name     string
	children []Component
}

func NewComposite(name string) *Composite { return &Composite{name: name} }

func (c *Composite) Name() string { return c.name }
func (c *Composite) Add(child Component) { c.children = append(c.children, child) }

func (c *Composite) Display(indent string) {
	fmt.Println(indent + c.name)
	for _, child := range c.children {
		child.Display(indent + "  ")
	}
}

func ExampleComposite() {
	root := NewComposite("root")
	root.Add(NewLeaf("leaf1"))
	sub := NewComposite("sub")
	sub.Add(NewLeaf("leaf2"))
	root.Add(sub)
	root.Display("")
	// Output:
	// root
	//   leaf1
	//   sub
	//     leaf2
}
