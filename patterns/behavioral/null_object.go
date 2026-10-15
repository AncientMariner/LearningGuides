package behavioral

import "fmt"

// Null Object: provides a no-op implementation so callers skip nil checks.

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(message string) { fmt.Println(message) }

// NullLogger silently discards all messages.
type NullLogger struct{}

func (l *NullLogger) Log(_ string) {}

func NewLogger(enabled bool) Logger {
	if enabled {
		return &ConsoleLogger{}
	}
	return &NullLogger{}
}

func ExampleNullObject() {
	active := NewLogger(true)
	silent := NewLogger(false)
	active.Log("hello")
	silent.Log("this is dropped")
	active.Log("world")
	// Output:
	// hello
	// world
}
