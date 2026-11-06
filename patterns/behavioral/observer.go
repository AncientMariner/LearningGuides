package behavioral

import (
	"fmt"
	"sync"
)

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
	mu          sync.RWMutex
	subscribers map[EventType][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{subscribers: make(map[EventType][]Subscriber)}
}

func (b *EventBus) Subscribe(t EventType, s Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[t] = append(b.subscribers[t], s)
}

func (b *EventBus) Publish(e Event) {
	b.mu.RLock()
	subs := b.subscribers[e.Type]
	b.mu.RUnlock()
	for _, s := range subs {
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
