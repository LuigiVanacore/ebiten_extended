package event

import "testing"

// testConnection is a simple implementation of the connection interface for tests.
type testConnection struct {
	disposed bool
}

func (c *testConnection) IsDisposed() bool { return c.disposed }

// TestConnectAndEmit verifies that a connected handler is called when Emit is invoked.
func TestConnectAndEmit(t *testing.T) {
	var got int
	var called bool

	var e Event[int]
	e.Connect(nil, func(v int) {
		called = true
		got = v
	})

	e.Emit(42)

	if !called {
		t.Fatalf("expected handler to be called")
	}
	if got != 42 {
		t.Fatalf("expected value 42, got %d", got)
	}
	if e.NumConnections() != 1 {
		t.Fatalf("expected 1 connection, got %d", e.NumConnections())
	}
}

// TestDisconnect ensures that Disconnect prevents further calls to a handler.
func TestDisconnect(t *testing.T) {
	var called bool
	var e Event[int]

	conn := &testConnection{}
	e.Connect(conn, func(v int) {
		called = true
	})

	e.Disconnect(conn)
	e.Emit(1)

	if called {
		t.Fatalf("handler should not be called after Disconnect")
	}
}

// TestReset clears all handlers and leaves the event empty.
func TestReset(t *testing.T) {
	var called bool
	var e Event[int]

	e.Connect(nil, func(v int) {
		called = true
	})

	e.Reset()

	if !e.IsEmpty() || e.NumConnections() != 0 {
		t.Fatalf("expected event to be empty after Reset, got NumConnections=%d", e.NumConnections())
	}

	e.Emit(10)
	if called {
		t.Fatalf("handler should not be called after Reset")
	}
}

// TestForward wires one event into another using Forward and ensures the second event fires.
func TestForward(t *testing.T) {
	var src Event[int]
	var dst Event[int]

	var called bool
	dst.Connect(nil, func(v int) {
		called = true
	})

	src.Forward(nil, &dst)
	src.Emit(5)

	if !called {
		t.Fatalf("expected forwarded event handler to be called")
	}
}
