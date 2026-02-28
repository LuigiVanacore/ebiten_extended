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

func (s *Node2D) GetTransform() transform.Transform {
	return s.transform
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


// GetWorldTransform returns the transform from root to this node (root * ... * parent * self).
func (b *Node2D) GetWorldTransform() transform.Transform {
	var chain []transform.Transform
	for node := SceneNode(b); node != nil; node = node.GetParent() {
		if entity, ok := node.(transform.Transformable); ok {
			chain = append(chain, entity.GetTransform())
		}
	}
	// Apply from root (last) to self (first)
	world := transform.Transform{}
	for i := len(chain) - 1; i >= 0; i-- {
		world.Concat(chain[i])
	}
	return world
}

func (b *Node2D) GetWorldPosition() math2D.Vector2D {
	worldTransform := b.GetWorldTransform()
	return worldTransform.GetPosition()
}