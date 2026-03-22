package physics

import (
	"errors"

	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/collision"
	"github.com/LuigiVanacore/ludum/math2d"
)

// RigidBody2D is a physics body with velocity, gravity, and collision shape.
// It does not overlap with other RigidBody2D (PhysicsWorld resolves collisions).
// Static bodies (e.g. floor, walls) do not move when colliding.
// Kinematic bodies are moved by code, not by physics; they push dynamic bodies but are not pushed.
// Friction (0-1) reduces sliding; Restitution (0-1) controls bounce.
type RigidBody2D struct {
	ludum.Node2D
	velocity     math2d.Vector2D
	mass         float64
	UsesGravity  bool
	GravityScale float64
	Static       bool // if true, body does not move on collision
	Kinematic    bool // if true, position is moved by code only; pushes dynamic bodies but is not pushed
	Friction     float64
	Restitution  float64
	shape        collision.CollisionShape
	mask         collision.CollisionMask
}

// NewRigidBody2D creates a RigidBody2D with the given name, collision shape and mask.
// Gravity is enabled by default (UsesGravity = true).
func NewRigidBody2D(name string, shape collision.CollisionShape, mask collision.CollisionMask) (*RigidBody2D, error) {
	if shape == nil {
		return nil, errors.New("physics: NewRigidBody2D shape must not be nil")
	}
	rb := &RigidBody2D{
		Node2D:       *ludum.NewNode2D(name),
		mass:         1,
		UsesGravity:  true,
		GravityScale: 1,
		Friction:     0.5,
		Restitution:  0,
		shape:        shape,
		mask:         mask,
	}
	return rb, nil
}

func (r *RigidBody2D) GetVelocity() math2d.Vector2D {
	return r.velocity
}

func (r *RigidBody2D) SetVelocity(v math2d.Vector2D) {
	r.velocity = v
}

func (r *RigidBody2D) ApplyImpulse(v math2d.Vector2D) {
	r.velocity = math2d.AddVectors(r.velocity, v)
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
