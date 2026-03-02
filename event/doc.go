// Package event provides a generic typed event (signal/slot) system.
//
// Event[T] holds listeners and invokes them when Emit(arg) is called. Use Connect
// to add a callback and Disconnect or a connection's disposal to remove it. IsEmpty
// and NumConnections report listener count. The collision package uses Event[*Collider]
// for OnCollisionEnter, OnCollisionStay, OnCollisionExit.
package event
