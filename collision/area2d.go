package collision

import (
	"errors"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/event"
)

// Area2D is a sensor/trigger: it detects overlaps but does not block movement.
// Subscribe to BodyEntered(), BodyStay(), BodyExited() for lifecycle events.
type Area2D struct {
	ebiten_extended.Node2D
	shape CollisionShape
	mask  CollisionMask

	onBodyEntered *event.Event[Area2DBodyEvent]
	onBodyStay    *event.Event[Area2DBodyEvent]
	onBodyExited  *event.Event[Area2DBodyEvent]
}

// NewArea2D creates an Area2D with the given name, shape, and mask.
func NewArea2D(name string, shape CollisionShape, mask CollisionMask) (*Area2D, error) {
	if shape == nil {
		return nil, errors.New("collision: NewArea2D shape must not be nil")
	}
	return &Area2D{
		Node2D:        *ebiten_extended.NewNode2D(name),
		shape:         shape,
		mask:          mask,
		onBodyEntered: &event.Event[Area2DBodyEvent]{},
		onBodyStay:    &event.Event[Area2DBodyEvent]{},
		onBodyExited:  &event.Event[Area2DBodyEvent]{},
	}, nil
}

// BodyEntered returns the event fired once when a collider first enters the area.
func (a *Area2D) BodyEntered() *event.Event[Area2DBodyEvent] { return a.onBodyEntered }

// BodyStay returns the event fired every frame while a collider remains inside the area.
func (a *Area2D) BodyStay() *event.Event[Area2DBodyEvent] { return a.onBodyStay }

// BodyExited returns the event fired once when a collider leaves the area.
func (a *Area2D) BodyExited() *event.Event[Area2DBodyEvent] { return a.onBodyExited }

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
