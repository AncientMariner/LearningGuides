package behavioral

import "fmt"

// Chain of Responsibility: passes a request along a chain of handlers.
// Middleware (e.g. http.Handler wrapping) is the canonical Go form.

type LogLevel int

const (
	LogLevelUnknown LogLevel = iota // 0 = unset/invalid
	LogLevelInfo                    // 1
	LogLevelWarn                    // 2
	LogLevelError                   // 3
)

type Handler interface {
	SetNext(Handler)
	Handle(level LogLevel, message string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) { b.next = h }

func (b *BaseHandler) handleNext(level LogLevel, message string) {
	if b.next != nil {
		b.next.Handle(level, message)
	}
}

type InfoHandler struct{ BaseHandler }

func (h *InfoHandler) Handle(level LogLevel, message string) {
	if level == LogLevelInfo {
		fmt.Printf("INFO: %s\n", message)
		return
	}
	h.handleNext(level, message)
}

type WarnHandler struct{ BaseHandler }

func (h *WarnHandler) Handle(level LogLevel, message string) {
	if level == LogLevelWarn {
		fmt.Printf("WARN: %s\n", message)
		return
	}
	h.handleNext(level, message)
}

type ErrorHandler struct{ BaseHandler }

func (h *ErrorHandler) Handle(level LogLevel, message string) {
	if level == LogLevelError {
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

	info.Handle(LogLevelInfo, "starting up")
	info.Handle(LogLevelWarn, "low memory")
	info.Handle(LogLevelError, "disk full")
	// Output:
	// INFO: starting up
	// WARN: low memory
	// ERROR: disk full
}
