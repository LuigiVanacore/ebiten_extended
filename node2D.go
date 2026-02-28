package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type Node2D struct {
	Node
	localTransform transform.Transform
	worldTransform transform.Transform
	isDirty        bool
}

func NewNode2D(name string) *Node2D {
	return &Node2D{Node: *NewNode(name), isDirty: true}
}

func (s *Node2D) GetTransform() transform.Transform {
	return s.localTransform
}

func (s *Node2D) SetTransform(transform transform.Transform) {
	s.localTransform = transform
	s.MarkDirty()
}

func (s *Node2D) SetPosition(x, y float64) {
	s.localTransform.SetPosition(x, y)
	s.MarkDirty()
}

func (s *Node2D) GetPosition() math2D.Vector2D {
	return s.localTransform.GetPosition()
}

func (s *Node2D) SetRotation(rotation float64) {
	s.localTransform.SetRotation(rotation)
	s.MarkDirty()
}

func (s *Node2D) GetRotation() float64 {
	return s.localTransform.GetRotation()
}

func (s *Node2D) SetScale(x, y float64) {
	s.localTransform.Scale(x, y)
	s.MarkDirty()
}

func (s *Node2D) MarkDirty() {
	if s.isDirty {
		return
	}
	s.isDirty = true
	s.Node.MarkDirty()
}

// GetWorldTransform returns the cached transform from root to this node (root * ... * parent * self).
// Avoids iterating and allocating if already clean.
func (b *Node2D) GetWorldTransform() transform.Transform {
	if !b.isDirty {
		return b.worldTransform
	}

	// Compute new world transform from parent if possible
	world := transform.Transform{}
	parent := b.GetParent()
	if parentTransformable, ok := parent.(transform.Transformable); ok {
		world = parentTransformable.GetWorldTransform()
	}

	// Apply self local transform to the parent's world matrix
	world.Concat(b.localTransform)

	// Cache
	b.worldTransform = world
	b.isDirty = false

	return b.worldTransform
}

func (b *Node2D) GetWorldPosition() math2D.Vector2D {
	worldTransform := b.GetWorldTransform()
	return worldTransform.GetPosition()
}
