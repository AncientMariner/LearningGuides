package behavioral

import "fmt"

// Memento: captures and restores an object's state without exposing internals.

type EditorMemento struct {
	content string
}

type Editor struct {
	content string
}

func (e *Editor) Type(text string) { e.content += text }
func (e *Editor) Content() string  { return e.content }

func (e *Editor) Save() *EditorMemento {
	return &EditorMemento{content: e.content}
}

func (e *Editor) Restore(m *EditorMemento) {
	e.content = m.content
}

type History struct {
	states []*EditorMemento
}

func (h *History) Push(m *EditorMemento) {
	h.states = append(h.states, m)
}

func (h *History) Pop() *EditorMemento {
	n := len(h.states)
	if n == 0 {
		return nil
	}
	m := h.states[n-1]
	h.states = h.states[:n-1]
	return m
}

func ExampleMemento() {
	editor := &Editor{}
	history := &History{}

	editor.Type("Hello")
	history.Push(editor.Save())

	editor.Type(", World")
	history.Push(editor.Save())

	editor.Type("!!!")
	fmt.Println(editor.Content())

	editor.Restore(history.Pop())
	fmt.Println(editor.Content())
	// Output:
	// Hello, World!!!
	// Hello, World
}
