package ebiten_extended


import "sync/atomic"

var globalNodeID uint64

type Node struct {
	id        uint64
	name      string
	children  []SceneNode
	parent    SceneNode
}

func NewNode(name string) *Node {
	id := atomic.AddUint64(&globalNodeID, 1)
	return &Node{id: id, name: name}
}

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

func (s *Node) Delete() {
	s.parent.DetachChild(s)
}
