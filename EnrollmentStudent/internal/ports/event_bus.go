package ports

import "Enrollment/internal/application/events"

type EventBuss interface {
	Publish(event_name string, payload any)
	Subscribe(event_name string, handler events.EventHandler)
}