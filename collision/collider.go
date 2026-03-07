package collision

import (
	"errors"
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/event"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

// DrawDebug draws the collision shape outline for debugging.
// Call when Engine.IsDebug() is true; typically from your game's Draw or a debug overlay.
func (c *Collider) DrawDebug(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	pos := c.GetWorldPosition()
	debugColor := color.RGBA{0, 255, 0, 180}

	switch s := c.collisionShape.(type) {
	case *CollisionCircle:
		r := s.circle.GetRadius()
		vector.StrokeCircle(target, float32(pos.X()), float32(pos.Y()), float32(r), 2, debugColor, true)
	case *CollisionRect:
		sz := s.rectangle.GetSize()
		tlX := pos.X() - sz.X()/2
		tlY := pos.Y() - sz.Y()/2
		vector.StrokeRect(target, float32(tlX), float32(tlY), float32(sz.X()), float32(sz.Y()), 2, debugColor, true)
	case *CollisionOrientedRect:
		or := s.rectangle
		rot := c.GetRotation()
		center := pos
		he := or.GetHalfExtended()
		// Draw 4 edges of oriented rect
		for i := 0; i < 4; i++ {
			c1 := OrientedRectangleCorner(math2D.NewOrientedRectangle(center, he, rot), i)
			c2 := OrientedRectangleCorner(math2D.NewOrientedRectangle(center, he, rot), (i+1)%4)
			vector.StrokeLine(target, float32(c1.X()), float32(c1.Y()), float32(c2.X()), float32(c2.Y()), 2, debugColor, true)
		}
	case *CollisionPolygon:
		verts := polygonWorldVertices(s.vertices, pos, c.GetRotation())
		for i := 0; i < len(verts); i++ {
			next := (i + 1) % len(verts)
			vector.StrokeLine(target, float32(verts[i].X()), float32(verts[i].Y()), float32(verts[next].X()), float32(verts[next].Y()), 2, debugColor, true)
		}
	}
}

// CanCollideWith returns true if this collider's mask allows collision with the other participant.
func (c *Collider) CanCollideWith(other CollisionParticipant) bool {
	return c.mask.IsCollidable(other.GetCollisionMask())
}
