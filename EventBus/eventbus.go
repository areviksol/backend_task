package eventbus

import (
	"sync"
)

// EventBus struct manages event subscriptions and notifications
type EventBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

// NewEventBus creates a new instance of EventBus
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

// Subscribe adds a new channel to the list of subscribers for a given event
func (eb *EventBus) Subscribe(event string) chan interface{} {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	ch := make(chan interface{}, 1)
	eb.subscribers[event] = append(eb.subscribers[event], ch)
	return ch
}

// Publish notifies all subscribers of a given event
func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, ch := range eb.subscribers[event] {
		ch <- data
	}
}
