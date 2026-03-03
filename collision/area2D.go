package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/event"
)

// Area2D is a sensor/trigger: it detects overlaps but does not block movement.
// Subscribe to OnBodyEntered, OnBodyStay, OnBodyExited for lifecycle events.
type Area2D struct {
	ebiten_extended.Node2D
	shape CollisionShape
	mask  CollisionMask

	OnBodyEntered *event.Event[Area2DBodyEvent]
	OnBodyStay    *event.Event[Area2DBodyEvent]
	OnBodyExited  *event.Event[Area2DBodyEvent]
}

// NewArea2D creates an Area2D with the given shape and mask. Panics if shape is nil.
func NewArea2D(shape CollisionShape, mask CollisionMask) *Area2D {
	if shape == nil {
		panic("collision: NewArea2D shape must not be nil")
	}
	return &Area2D{
		shape:         shape,
		mask:          mask,
		OnBodyEntered: &event.Event[Area2DBodyEvent]{},
		OnBodyStay:    &event.Event[Area2DBodyEvent]{},
		OnBodyExited:  &event.Event[Area2DBodyEvent]{},
	}
}

func (a *Area2D) GetShape() CollisionShape {
	return a.shape
}

func (a *Area2D) GetCollisionMask() CollisionMask {
	return a.mask
}

func (a *Area2D) SetCollisionMask(mask CollisionMask) {
	a.mask = mask
}

func (a *Area2D) CanCollideWith(other CollisionParticipant) bool {
	return a.mask.IsCollidable(other.GetCollisionMask())
}
