package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/event"
	"github.com/hajimehoshi/ebiten/v2"
)

type Collider struct {
	ebiten_extended.Node2D
	collisionShape           CollisionShape
	mask                     CollisionMask
	isWorldCoordinateUpdated bool

	OnCollisionEnter *event.Event
	OnCollisionStay  *event.Event
	OnCollisionExit  *event.Event
}

func NewCollider(shape CollisionShape, mask CollisionMask) *Collider {
	c := &Collider{
		collisionShape:   shape,
		mask:             mask,
		OnCollisionEnter: &event.Event{},
		OnCollisionStay:  &event.Event{},
		OnCollisionExit:  &event.Event{},
	}
	return c
}

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

func (c *Collider) CanCollideWith(collider *Collider) bool {
	return c.mask.IsCollidable(collider.GetCollisionMask())
}

func (c *Collider) IsColliding(collider *Collider) bool {
	if c.collisionShape == nil || collider.collisionShape == nil {
		return false
	}
	leftWorldTransf := c.GetWorldTransform()
	rightWorldTransf := collider.GetWorldTransform()
	c.collisionShape.UpdateTransform(leftWorldTransf)
	collider.collisionShape.UpdateTransform(rightWorldTransf)
	return c.collisionShape.IsColliding(collider.collisionShape)
}
