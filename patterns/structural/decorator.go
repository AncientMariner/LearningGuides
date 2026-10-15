package structural

import "fmt"

// Decorator: wraps an interface to add behaviour without modifying the original.
// Middleware chains (e.g. http.Handler) are the canonical Go use case.

type TextProcessor interface {
	Process(text string) string
}

type PlainText struct{}

func (p *PlainText) Process(text string) string { return text }

type UpperCaseDecorator struct {
	inner TextProcessor
}

func NewUpperCase(inner TextProcessor) TextProcessor {
	return &UpperCaseDecorator{inner: inner}
}

func (u *UpperCaseDecorator) Process(text string) string {
	result := u.inner.Process(text)
	upper := make([]byte, len(result))
	for i := range result {
		c := result[i]
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		upper[i] = c
	}
	return string(upper)
}

type PrefixDecorator struct {
	inner  TextProcessor
	prefix string
}

func NewPrefix(inner TextProcessor, prefix string) TextProcessor {
	return &PrefixDecorator{inner: inner, prefix: prefix}
}

func (p *PrefixDecorator) Process(text string) string {
	return fmt.Sprintf("%s%s", p.prefix, p.inner.Process(text))
}

func ExampleDecorator() {
	var p TextProcessor = &PlainText{}
	p = NewUpperCase(p)
	p = NewPrefix(p, ">> ")
	fmt.Println(p.Process("hello"))
	// Output:
	// >> HELLO
}
