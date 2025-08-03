package event

import (
	"testing"
)

// mockEvent implements the Event interface for testing.
type mockEvent struct {
	id EventID
}

func (e *mockEvent) GetEventID() EventID {
	return e.id
}

func TestSubscribeAndPublish(t *testing.T) {
	bus := NewEventBus()
	eventID := EventID(1)
	called := false

	handler := func(e Event) {
		called = true
	}

	bus.Subscribe(eventID, handler)
	bus.Publish(&mockEvent{id: eventID})

	if !called {
		t.Errorf("Expected handler to be called on publish")
	}
}

func TestUnsubscribe(t *testing.T) {
	bus := NewEventBus()
	eventID := EventID(2)
	called := false

	handler := func(e Event) {
		called = true
	}

	index := bus.Subscribe(eventID, handler)
	bus.Unsubscribe(eventID, index)
	bus.Publish(&mockEvent{id: eventID})

	if called {
		t.Errorf("Handler should not be called after unsubscribe")
	}
}

func TestUnsubscribeAll(t *testing.T) {
	bus := NewEventBus()
	eventID := EventID(3)
	called := false

	handler := func(e Event) {
		called = true
	}

	bus.Subscribe(eventID, handler)
	bus.UnsubscribeAll(eventID)
	bus.Publish(&mockEvent{id: eventID})

	if called {
		t.Errorf("Handler should not be called after UnsubscribeAll")
	}
}

func TestClear(t *testing.T) {
	bus := NewEventBus()
	eventID := EventID(4)
	called := false

	handler := func(e Event) {
		called = true
	}

	bus.Subscribe(eventID, handler)
	bus.Clear()
	bus.Publish(&mockEvent{id: eventID})

	if called {
		t.Errorf("Handler should not be called after Clear")
	}
}

func TestGetSubscribers(t *testing.T) {
	bus := NewEventBus()
	eventID := EventID(5)

	handler1 := func(e Event) {}
	handler2 := func(e Event) {}

	bus.Subscribe(eventID, handler1)
	bus.Subscribe(eventID, handler2)

	subs := bus.GetSubscribers(eventID)
	if len(subs) != 2 {
		t.Errorf("Expected 2 subscribers, got %d", len(subs))
	}
}

func TestUnsubscribeInvalidIndex(t *testing.T) {
	bus := NewEventBus()
	eventID := EventID(6)

	handler := func(e Event) {}

	bus.Subscribe(eventID, handler)
	// Try to unsubscribe with an invalid index (out of range)
	bus.Unsubscribe(eventID, IndexSubscriber(10))
	// Should not panic or remove the handler
	if len(bus.GetSubscribers(eventID)) != 1 {
		t.Errorf("Handler should not be removed with invalid index")
	}
}
