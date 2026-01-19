package events

import (
	"context"
	"sync"
)

// Event interface for all domain events
type Event interface {
	Name() string
}

// EventHandler is a function that handles events
type EventHandler func(ctx context.Context, event Event) error

// EventBus manages event subscriptions and publishing
type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

// NewEventBus creates a new EventBus
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

// Subscribe adds a handler for specific event types
func (eb *EventBus) Subscribe(eventName string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
}

// Publish sends an event to all subscribers
func (eb *EventBus) Publish(ctx context.Context, event Event) error {
	eb.mu.RLock()
	handlers := eb.handlers[event.Name()]
	eb.mu.RUnlock()

	for _, handler := range handlers {
		if err := handler(ctx, event); err != nil {
			return err
		}
	}
	return nil
}

// PublishAsync sends an event asynchronously (fire and forget)
func (eb *EventBus) PublishAsync(ctx context.Context, event Event) {
	go func() {
		eb.mu.RLock()
		handlers := eb.handlers[event.Name()]
		eb.mu.RUnlock()

		for _, handler := range handlers {
			_ = handler(ctx, event)
		}
	}()
}

// UserCreatedEvent is published when a new user is created
type UserCreatedEvent struct {
	UserID   string
	Email    string
	Role     string
	TenantID string
}

func (e UserCreatedEvent) Name() string { return "user.created" }

// UserUpdatedEvent is published when a user is updated
type UserUpdatedEvent struct {
	UserID string
}

func (e UserUpdatedEvent) Name() string { return "user.updated" }
