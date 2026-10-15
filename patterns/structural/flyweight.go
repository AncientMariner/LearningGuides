package structural

import "fmt"

// Flyweight: shares immutable state across many objects to reduce memory.
// sync.Pool handles pooling; here we show an intrinsic-state cache.

type TreeType struct {
	name    string
	color   string
	texture string
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf("Drawing %s tree at (%d,%d)\n", t.name, x, y)
}

type TreeTypeFactory struct {
	cache map[string]*TreeType
}

func NewTreeTypeFactory() *TreeTypeFactory {
	return &TreeTypeFactory{cache: make(map[string]*TreeType)}
}

func (f *TreeTypeFactory) Get(name, color, texture string) *TreeType {
	key := name + color + texture
	if t, ok := f.cache[key]; ok {
		return t
	}
	t := &TreeType{name: name, color: color, texture: texture}
	f.cache[key] = t
	return t
}

type Tree struct {
	x, y     int
	treeType *TreeType
}

func (t *Tree) Draw() { t.treeType.Draw(t.x, t.y) }

func ExampleFlyweight() {
	factory := NewTreeTypeFactory()
	oak := factory.Get("oak", "green", "rough")
	trees := []Tree{
		{x: 1, y: 2, treeType: oak},
		{x: 3, y: 4, treeType: oak},
		{x: 5, y: 6, treeType: factory.Get("oak", "green", "rough")},
	}
	for _, t := range trees {
		t.Draw()
	}
	// Output:
	// Drawing oak tree at (1,2)
	// Drawing oak tree at (3,4)
	// Drawing oak tree at (5,6)
}
