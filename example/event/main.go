package main

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended/event"
)


var (	UserCreatedEventID event.EventID = 1
	OrderPlacedEventID event.EventID = 2
)


type UserCreatedEvent struct {
	UserID string
	Name   string
}

 
type OrderPlacedEvent struct {
	OrderID  string
	Amount   float64
	Currency string
}

func (e UserCreatedEvent) GetEventID() event.EventID {
	return UserCreatedEventID
}

func (e OrderPlacedEvent) GetEventID() event.EventID {
	return OrderPlacedEventID
}

func main() {
	bus := event.NewEventBus()

	// Subscribe with a type-safe adapter
	bus.Subscribe(UserCreatedEventID, func(e event.Event) {
		if ue, ok := e.(UserCreatedEvent); ok {
			fmt.Println("User Created:", ue.UserID, ue.Name)
		}
	})

	bus.Subscribe(OrderPlacedEventID, func(e event.Event) {
		if oe, ok := e.(OrderPlacedEvent); ok {
			fmt.Println("Order Placed 1:", oe.OrderID, oe.Amount, oe.Currency)
		}
	})

	bus.Subscribe(OrderPlacedEventID, func(e event.Event) {
		if oe, ok := e.(OrderPlacedEvent); ok {
			fmt.Println("Order Placed 2:", oe.OrderID)
		}
	})

	// Publish
	bus.Publish(UserCreatedEvent{UserID: "123", Name: "Luigi"})
	bus.Publish(OrderPlacedEvent{OrderID: "A001", Amount: 42.0, Currency: "EUR"})
}
