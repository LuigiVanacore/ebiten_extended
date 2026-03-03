package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// CollisionParticipant is implemented by Collider, Area2D, and RigidBody2D.
// It provides the data needed for overlap detection and event dispatch.
type CollisionParticipant interface {
	GetWorldPosition() math2D.Vector2D
	GetWorldTransform() transform.Transform
	GetShape() CollisionShape
	GetID() uint64
	GetCollisionMask() CollisionMask
	CanCollideWith(other CollisionParticipant) bool
}

// Body is a minimal interface for entities that can trigger Area2D events (Collider or RigidBody2D).
type Body interface {
	GetWorldPosition() math2D.Vector2D
	GetShape() CollisionShape
	GetID() uint64
}

// Area2DBodyEvent is emitted when a body enters or exits an Area2D.
type Area2DBodyEvent struct {
	Body Body
}
