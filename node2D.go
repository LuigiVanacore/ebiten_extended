package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type Node2D struct {
	Node
	transform.Transform
}

func NewNode2D(name string) *Node2D {
	return &Node2D{Node: *NewNode(name)}
}

func (s *Node2D) GetTransform() transform.Transform {
	return s.Transform
}

func (s *Node2D) SetTransform(transform transform.Transform) {
	s.Transform = transform
}

func (b *Node2D) GetWorldTransform() transform.Transform {
	rootTransform := transform.Transform{}
	for node := b.GetParent(); node != nil; node = node.GetParent() {
		if entity, ok := node.(transform.Transformable); ok {
			rootTransform.Concat(entity.GetTransform())
		}
	}
	return rootTransform
}

func (b *Node2D) GetWorldPosition() math2D.Vector2D {
	worldTransform := b.GetWorldTransform()
	return worldTransform.GetPosition()
}
