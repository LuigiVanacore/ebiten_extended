package collision

import (
	"errors"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/event"
	"github.com/hajimehoshi/ebiten/v2"
)

// Collider is a Node2D with a collision shape and mask. Subscribe to CollisionEnter(),
// CollisionStay(), and CollisionExit() for lifecycle events; add to a CollisionManager and
// run CheckCollision each frame.
type Collider struct {
	ebiten_extended.Node2D
	collisionShape           CollisionShape
	mask                     CollisionMask
	isWorldCoordinateUpdated bool

	onCollisionEnter *event.Event[*Collider]
	onCollisionStay  *event.Event[*Collider]
	onCollisionExit  *event.Event[*Collider]
}

// NewCollider returns a new collider with the given name, shape, and mask.
// Add it to a CollisionManager with AddCollider.
func NewCollider(name string, shape CollisionShape, mask CollisionMask) (*Collider, error) {
	if shape == nil {
		return nil, errors.New("collision: NewCollider shape must not be nil")
	}
	c := &Collider{
		Node2D:           *ebiten_extended.NewNode2D(name),
		collisionShape:   shape,
		mask:             mask,
		onCollisionEnter: &event.Event[*Collider]{},
		onCollisionStay:  &event.Event[*Collider]{},
		onCollisionExit:  &event.Event[*Collider]{},
	}
	return c, nil
}

// CollisionEnter returns the event fired once when another collider first overlaps this one.
func (c *Collider) CollisionEnter() *event.Event[*Collider] { return c.onCollisionEnter }

// CollisionStay returns the event fired every frame while another collider continues to overlap.
func (c *Collider) CollisionStay() *event.Event[*Collider] { return c.onCollisionStay }

// CollisionExit returns the event fired once when an overlapping collider stops touching.
func (c *Collider) CollisionExit() *event.Event[*Collider] { return c.onCollisionExit }

func (c *Collider) IsWorldCoordinateUpdated() bool {
	return c.isWorldCoordinateUpdated
}

func (c *Collider) SetWorldCoordinateUpdated(flag bool) {
	c.isWorldCoordinateUpdated = flag
}

func (c *Collider) GetCollisionMask() CollisionMask {
	return c.mask
}

func (c *Collider) SetCollisionMask(mask CollisionMask) {
	c.mask = mask
}

func (c *Collider) GetShape() CollisionShape {
	return c.collisionShape
}

func (c *Collider) DrawDebug(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	// Optional: draw collision shape outline for debugging
}

// CanCollideWith returns true if this collider's mask allows collision with the other participant.
func (c *Collider) CanCollideWith(other CollisionParticipant) bool {
	return c.mask.IsCollidable(other.GetCollisionMask())
}
