package events

import (
	"github.com/stretchr/testify/mock"
)

type EventHandler func(payload any)

type EventBus struct {
	mock.Mock
	subscribers map[string][]EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventHandler),
	}
}

func (e *EventBus) Publish(eventName string, payload any) {
	e.Called(eventName)
	
	if handlers, ok := e.subscribers[eventName]; ok {
		for _, handler := range handlers {
			handler(payload)
		}
	}
}

func (e *EventBus) Subscribe(eventName string, handler EventHandler) {
	e.subscribers[eventName] = append(e.subscribers[eventName], handler)
}