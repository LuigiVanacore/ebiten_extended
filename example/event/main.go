package main

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended/event"
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

func main() {
	userCreated := &event.Event[UserCreatedEvent]{}
	orderPlaced := &event.Event[OrderPlacedEvent]{}

	userCreated.Connect(nil, func(e UserCreatedEvent) {
		fmt.Println("User Created:", e.UserID, e.Name)
	})

	orderPlaced.Connect(nil, func(e OrderPlacedEvent) {
		fmt.Println("Order Placed 1:", e.OrderID, e.Amount, e.Currency)
	})
	orderPlaced.Connect(nil, func(e OrderPlacedEvent) {
		fmt.Println("Order Placed 2:", e.OrderID)
	})

	userCreated.Emit(UserCreatedEvent{UserID: "123", Name: "Luigi"})
	orderPlaced.Emit(OrderPlacedEvent{OrderID: "A001", Amount: 42.0, Currency: "EUR"})
}
