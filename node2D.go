package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// Node2D represents a 2D scene graph node incorporating spatial transformations (position, rotation, scale).
type Node2D struct {
	Node
	localTransform transform.Transform
	worldTransform transform.Transform
	isDirty        bool
}

// NewNode2D creates and initializes a new Node2D with a given name, marking it as initially dirty.
func NewNode2D(name string) *Node2D {
	return &Node2D{Node: *NewNode(name), isDirty: true}
}

// AddChildren overrides Node.AddChildren to pass the Node2D itself as parent (not the embedded Node),
// so GetParent() returns a transform.Transformable when the parent is a Node2D.
func (n *Node2D) AddChildren(child SceneNode) {
	child.AttachParent(n)
	n.children = append(n.children, child)
}

// GetTransform retrieves the local spatial transform of this node.
func (s *Node2D) GetTransform() transform.Transform {
	return s.localTransform
}

// SetTransform updates the local spatial transform and flags the node as dirty to trigger recalculations.
func (s *Node2D) SetTransform(transform transform.Transform) {
	s.localTransform = transform
	s.MarkDirty()
}

// SetPosition modifies the local X and Y position of the node and marks it as dirty.
func (s *Node2D) SetPosition(x, y float64) {
	s.localTransform.SetPosition(math2D.NewVector2D(x, y))
	s.MarkDirty()
}

// GetPosition returns the current local 2D position vector of the node.
func (s *Node2D) GetPosition() math2D.Vector2D {
	return s.localTransform.GetPosition()
}

// SetRotation assigns the local rotation angle (in radians) of the node and marks it as dirty.
func (s *Node2D) SetRotation(rotation float64) {
	s.localTransform.SetRotation(rotation)
	s.MarkDirty()
}

// GetRotation returns the local rotation angle (in radians) of the node.
func (s *Node2D) GetRotation() float64 {
	return s.localTransform.GetRotation()
}

// SetScale adjusts the local scale factors along the X and Y axes and marks the node as dirty.
func (s *Node2D) SetScale(x, y float64) {
	s.localTransform.Scale(x, y)
	s.MarkDirty()
}

// MarkDirty flags this node and its children to recalculate their world transforms on the next query.
func (s *Node2D) MarkDirty() {
	if s.isDirty {
		return
	}
	s.isDirty = true
	s.Node.MarkDirty()
}

// GetWorldTransform returns the cached transform navigating from root to this node (root * ... * parent * self).
// It avoids iterating and reallocating if the hierarchy is currently clean.
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

	// Apply self local transform: combined = parent_world + local
	world.Concat(b.localTransform)

	// Cache
	b.worldTransform = world
	b.isDirty = false

	return b.worldTransform
}

// GetWorldPosition extracts and returns the absolute 2D position in the world from the calculated world transform.
func (b *Node2D) GetWorldPosition() math2D.Vector2D {
	worldTransform := b.GetWorldTransform()
	return worldTransform.GetPosition()
}

