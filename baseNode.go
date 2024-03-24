package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type BaseNode struct {
	transform transform.Transform
	id        uint64
	name      string
	children  []SceneNode
	parent    SceneNode
}

func NewBaseNode(name string) *BaseNode {
	id := SceneManager().GetNextIdVal()
	return &BaseNode{id: id, name: name}
}

func (s *BaseNode) GetTransform() *transform.Transform {
	return &s.transform
}

func (s *BaseNode) SetTransform(transform transform.Transform) {
	s.transform = transform
}

func (s *BaseNode) SetPosition(x, y float64) {
	s.transform.SetPosition(x, y)
}

func (s *BaseNode) GetPosition() math2D.Vector2D {
	return s.transform.GetPosition()
}

func (s *BaseNode) SetRotation(rotation int) {
	s.transform.SetRotation(rotation)
}

func (s *BaseNode) GetRotation() int {
	return s.transform.GetRotation()
}

func (s *BaseNode) SetScale(x, y float64) {
	s.transform.Scale(x, y)
}

//	if Aer, ok := s.entity.(GetSprite); ok {
//		//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
//	for _, child := range s.children {
//		child.Draw(target, localOp)
//	}
//}

func (s *BaseNode) AddChildren(child SceneNode) {
	child.AttachParent(s)

	s.children = append(s.children, child)
}



func (s *BaseNode) DetachChild(node SceneNode) bool {
	for i, child := range s.children {
		if child == node {
			s.children[i] = s.children[len(s.children)-1]
			s.children = s.children[:len(s.children)-1]
			return true
		}
	}
	return false
}

func (s *BaseNode) AttachParent(node SceneNode) {
	s.parent = node
}

func (b *BaseNode) GetParent() SceneNode {
	return b.parent
}

func (s *BaseNode) GetChildren() []SceneNode {
	return s.children
}

func (b *BaseNode) GetWorldTransform() transform.Transform {
	rootTransform := transform.Transform{}
	for node := b.GetParent(); node != nil; node = node.GetParent()  {
		if entity, ok := node.(transform.Transformable); ok {
			rootTransform.Concat(*entity.GetTransform())
		}
	}
	return rootTransform
}

func (b *BaseNode) GetWorldPosition() math2D.Vector2D {
	worldTransform := b.GetWorldTransform()
	return worldTransform.GetPosition()
}

//func (s *SceneNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
//	localOp := &ebiten.DrawImageOptions{}
//	localOp.GeoM = s.updateGeoM(localOp.GeoM)
//	if Aer, ok := s.entity.(GetSprite); ok {
//		//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
//	for _, child := range s.children {
//		child.Draw(target, localOp)
//	}
//}

//func (s *SceneNode) GetWorldPosition() math2d.Vector2D {
//	transform := s.GetWorldTransform()
//	x, y := transform.Apply(0, 0)
//	return math2d.Vector2D{X: x, Y: y}
//}
//
//func (s *SceneNode) GetWorldTransform() ebiten.GeoM {
//	transform := ebiten.GeoM{}
//
//	for node := Node(s); node != nil; node = s.parent {
//		getTransform := node.GetTransform()
//		transform.Concat(getTransform)
//	}
//
//	return transform
//}


func (s *BaseNode) Delete() {
	s.parent.DetachChild(s)
}