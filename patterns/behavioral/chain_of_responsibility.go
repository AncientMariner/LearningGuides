package behavioral

import "fmt"

// Chain of Responsibility: passes a request along a chain of handlers.
// Middleware (e.g. http.Handler wrapping) is the canonical Go form.

type Handler interface {
	SetNext(Handler)
	Handle(level int, message string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) { b.next = h }

func (b *BaseHandler) handleNext(level int, message string) {
	if b.next != nil {
		b.next.Handle(level, message)
	}
}

type InfoHandler struct{ BaseHandler }

func (h *InfoHandler) Handle(level int, message string) {
	if level == 1 {
		fmt.Printf("INFO: %s\n", message)
		return
	}
	h.handleNext(level, message)
}

type WarnHandler struct{ BaseHandler }

func (h *WarnHandler) Handle(level int, message string) {
	if level == 2 {
		fmt.Printf("WARN: %s\n", message)
		return
	}
	h.handleNext(level, message)
}

type ErrorHandler struct{ BaseHandler }

func (h *ErrorHandler) Handle(level int, message string) {
	if level == 3 {
		fmt.Printf("ERROR: %s\n", message)
		return
	}
	h.handleNext(level, message)
}

func ExampleChainOfResponsibility() {
	info := &InfoHandler{}
	warn := &WarnHandler{}
	errH := &ErrorHandler{}
	info.SetNext(warn)
	warn.SetNext(errH)

	info.Handle(1, "starting up")
	info.Handle(2, "low memory")
	info.Handle(3, "disk full")
	// Output:
	// INFO: starting up
	// WARN: low memory
	// ERROR: disk full
}
