package behavioral

import "fmt"

// Observer: notifies subscribers when state changes.
// Channels are idiomatic for goroutine-based observers; interfaces suit sync use.

type EventType string

const (
	EventCreated EventType = "created"
	EventDeleted EventType = "deleted"
)

type Event struct {
	Type    EventType
	Payload any
}

type Subscriber interface {
	OnEvent(Event)
}

type EventBus struct {
	subscribers map[EventType][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{subscribers: make(map[EventType][]Subscriber)}
}

func (b *EventBus) Subscribe(t EventType, s Subscriber) {
	b.subscribers[t] = append(b.subscribers[t], s)
}

func (b *EventBus) Publish(e Event) {
	for _, s := range b.subscribers[e.Type] {
		s.OnEvent(e)
	}
}

type AuditLog struct{}

func (a *AuditLog) OnEvent(e Event) {
	fmt.Printf("audit: %s — %v\n", e.Type, e.Payload)
}

func ExampleObserver() {
	bus := NewEventBus()
	bus.Subscribe(EventCreated, &AuditLog{})
	bus.Publish(Event{Type: EventCreated, Payload: "user:42"})
	// Output:
	// audit: created — user:42
}
