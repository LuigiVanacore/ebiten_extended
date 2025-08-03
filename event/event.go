package event


type EventID uint64

type IndexSubscriber int

type Event interface {
	GetEventID() EventID
}


type EventBus struct {
	subscribers map[EventID][]func(Event)
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventID][]func(Event)),
	}
}

func (bus *EventBus) Subscribe(eventID EventID, handler func(Event)) IndexSubscriber {
	bus.subscribers[eventID] = append(bus.subscribers[eventID], handler)
	return (IndexSubscriber) (len(bus.subscribers[eventID]) - 1)
}

func (bus *EventBus) Publish(event Event) {
	if handlers, ok := bus.subscribers[event.GetEventID()]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

func (bus *EventBus) UnsubscribeAll(eventID EventID) {
	if _, ok := bus.subscribers[eventID]; ok {
		bus.subscribers[eventID] = nil
	}
}

func (bus *EventBus) Unsubscribe(eventID EventID, index IndexSubscriber) {
	if handlers, ok := bus.subscribers[eventID]; ok {
		if int(index) < len(handlers) {
			bus.subscribers[eventID] = append(handlers[:index], handlers[index+1:]...)
		}
	}

}
func (bus *EventBus) Clear() {
	bus.subscribers = make(map[EventID][]func(Event))
}

func (bus *EventBus) GetSubscribers(eventID EventID) []func(Event) {
	if handlers, ok := bus.subscribers[eventID]; ok {
		return handlers
	}
	return nil
}

