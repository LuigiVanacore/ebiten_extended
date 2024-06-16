package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type Node2D struct {
	Node
	transform transform.Transform
}

func NewNode2D(name string) *Node2D {
	return &Node2D{ Node: *NewNode(name)}
}

func (s *Node2D) GetTransform() *transform.Transform {
	return &s.transform
}

func (s *Node2D) SetTransform(transform transform.Transform) {
	s.transform = transform
}

func (s *Node2D) SetPosition(x, y float64) {
	s.transform.SetPosition(x, y)
}

func (s *Node2D) GetPosition() math2D.Vector2D {
	return s.transform.GetPosition()
}

func (s *Node2D) SetRotation(rotation int) {
	s.transform.SetRotation(rotation)
}

func (s *Node2D) GetRotation() int {
	return s.transform.GetRotation()
}

func (s *Node2D) SetScale(x, y float64) {
	s.transform.Scale(x, y)
}


func (b *Node2D) GetWorldTransform() transform.Transform {
	rootTransform := transform.Transform{}
	for node := b.GetParent(); node != nil; node = node.GetParent() {
		if entity, ok := node.(transform.Transformable); ok {
			rootTransform.Concat(*entity.GetTransform())
		}
	}
	return rootTransform
}

func (b *Node2D) GetWorldPosition() math2D.Vector2D {
	worldTransform := b.GetWorldTransform()
	return worldTransform.GetPosition()
}