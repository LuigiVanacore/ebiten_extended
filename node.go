package ebiten_extended



type Node struct {
	id        uint64
	name      string
	children  []SceneNode
	parent    SceneNode
}

func NewNode(name string) *Node {
	id := SceneManager().GetNextIdVal()
	return &Node{id: id, name: name}
}


//	if Aer, ok := s.entity.(GetSprite); ok {
//		//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
//	for _, child := range s.children {
//		child.Draw(target, localOp)
//	}
//}

func (n *Node) GetID() uint64 {
	return n.id
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) SetName(name string) {
	n.name = name
}

func (s *Node) AddChildren(child SceneNode) {
	child.AttachParent(s)

	s.children = append(s.children, child)
}

func (s *Node) DetachChild(node SceneNode) bool {
	for i, child := range s.children {
		if child == node {
			s.children[i] = s.children[len(s.children)-1]
			s.children = s.children[:len(s.children)-1]
			return true
		}
	}
	return false
}

func (s *Node) AttachParent(node SceneNode) {
	s.parent = node
}

func (b *Node) GetParent() SceneNode {
	return b.parent
}

func (s *Node) GetChildren() []SceneNode {
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

func (s *Node) Delete() {
	s.parent.DetachChild(s)
}
