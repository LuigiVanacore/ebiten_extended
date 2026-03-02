package transform

// Transformable is implemented by nodes that have a local and world transform (e.g. Node2D).
// The engine uses it for drawing and collision.
type Transformable interface {
	GetTransform() Transform
	SetTransform(transform Transform)
	GetWorldTransform() Transform
}
