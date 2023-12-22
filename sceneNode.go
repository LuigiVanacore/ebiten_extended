package ebiten_extended

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneNode struct {
	id uint64
	name string
	entity any
	children []*SceneNode
	parent   *SceneNode
}

func NewSceneNode(entity any, name string) *SceneNode {
	id := SceneManager().GetNextIdVal()
	return &SceneNode{entity: entity, id: id, name: name}
}

func NewRootSceneNode() *SceneNode {
	id := SceneManager().GetNextIdVal()
	return &SceneNode{entity: nil, id: id, name: "Root" }
}


func (s *SceneNode) Update() {
	s.UpdateChildren()
	if entity, ok := s.entity.(Updatable); ok {
		entity.Update()
	}
}

func (s *SceneNode) UpdateChildren() {
	for _, child := range s.children {
		child.Update()
	}
}

func (s *SceneNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {

	if entity, ok := s.entity.(Drawable); ok {
		op.GeoM = updateTransform(entity, op.GeoM )
		entity.Draw(target, op)
	}
	for _, child := range s.children {
				child.Draw(target, op)
	}
}

func updateTransform(entity transform.Transformable, parent_geoM ebiten.GeoM) ebiten.GeoM {
		rotation_geoM := ebiten.GeoM{}
		transform := entity.GetTransform()
		position := transform.GetPosition()
		pivot := transform.GetPivot()
		rotation := transform.GetRotation()

		rotation_geoM.Translate(-pivot.X(), -pivot.Y())
		rotation_geoM.Rotate(float64(rotation%360) * 2 * math.Pi / 360)
		rotation_geoM.Translate(pivot.X(), pivot.Y())

		parent_geoM.Translate(position.X(), position.Y())

		rotation_geoM.Concat(parent_geoM)
	
	return rotation_geoM
}
	
//	if Aer, ok := s.entity.(GetSprite); ok {
//		//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
//	for _, child := range s.children {
//		child.Draw(target, localOp)
//	}
//}

func (s *SceneNode) Addhildren(child *SceneNode) {
	child.AttachParent(s)

	s.children = append(s.children, child)
}

func (s *SceneNode) AddChildrenEntity(entity any, name string) {
	sceneNode := NewSceneNode(entity, name)

	s.Addhildren(sceneNode)
}

func (s *SceneNode) DetachChild(node *SceneNode) bool {
	for i, child := range s.children {
		if child == node {
			s.children[i] = s.children[len(s.children)-1]
			s.children = s.children[:len(s.children)-1]
			return true
		}
	}
	return false
}

func (s *SceneNode) AttachParent(node *SceneNode) {
	s.parent = node
}

func (s *SceneNode) GetChildren() []*SceneNode {
	return s.children
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
//
//func (s *SceneNode) Translate(x, y float64) {
//	s.transform.Translate(x, y)
//}

func (s *SceneNode) Delete() {
	s.parent.DetachChild(s)
}