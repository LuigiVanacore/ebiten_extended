package physics

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

// RigidBody2D is a physics body with velocity, gravity, and collision shape.
// It does not overlap with other RigidBody2D (PhysicsWorld resolves collisions).
// Static bodies (e.g. floor, walls) do not move when colliding.
// Friction (0-1) reduces sliding; Restitution (0-1) controls bounce.
type RigidBody2D struct {
	ebiten_extended.Node2D
	velocity     math2D.Vector2D
	mass         float64
	UsesGravity  bool
	GravityScale float64
	Static       bool // if true, body does not move on collision
	Friction     float64 // 0=no friction, 1=full friction; combined with other body on collision
	Restitution  float64 // 0=no bounce, 1=full bounce; min of both bodies used
	shape        collision.CollisionShape
	mask         collision.CollisionMask
}

// NewRigidBody2D creates a RigidBody2D. Panics if shape is nil.
func NewRigidBody2D(shape collision.CollisionShape, mask collision.CollisionMask) *RigidBody2D {
	if shape == nil {
		panic("physics: NewRigidBody2D shape must not be nil")
	}
	rb := &RigidBody2D{
		Node2D:      *ebiten_extended.NewNode2D("rigidbody"),
		mass:        1,
		GravityScale: 1,
		Friction:    0.5,
		Restitution: 0,
		shape:       shape,
		mask:        mask,
	}
	return rb
}

func (r *RigidBody2D) GetVelocity() math2D.Vector2D {
	return r.velocity
}

func (r *RigidBody2D) SetVelocity(v math2D.Vector2D) {
	r.velocity = v
}

func (r *RigidBody2D) ApplyImpulse(v math2D.Vector2D) {
	r.velocity = math2D.AddVectors(r.velocity, v)
}

func (r *RigidBody2D) GetShape() collision.CollisionShape {
	return r.shape
}

func (r *RigidBody2D) GetCollisionMask() collision.CollisionMask {
	return r.mask
}

func (r *RigidBody2D) SetCollisionMask(mask collision.CollisionMask) {
	r.mask = mask
}

func (r *RigidBody2D) CanCollideWith(other collision.CollisionParticipant) bool {
	return r.mask.IsCollidable(other.GetCollisionMask())
}
