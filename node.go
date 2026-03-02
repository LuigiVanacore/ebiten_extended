package ebiten_extended

import "sync/atomic"

var globalNodeID uint64

// Node represents an independent object or entity in a scene graph.
// It tracks its unique ID, an identifiable name, a parent, and multiple child elements.
type Node struct {
	id       uint64
	name     string
	children []SceneNode
	parent   SceneNode
	id       uint64
	name     string
	children []SceneNode
	parent   SceneNode
}

// NewNode creates and initializes a new Node with a generated unique ID and given name.
func NewNode(name string) *Node {
	id := atomic.AddUint64(&globalNodeID, 1)
	return &Node{id: id, name: name}
}

// GetID returns the unique 64-bit integer identifier of the node.
func (n *Node) GetID() uint64 {
	return n.id
}

// GetName returns the semantic string name of the node.
func (n *Node) GetName() string {
	return n.name
}

// SetName replaces the string name identifier of the node.
func (n *Node) SetName(name string) {
	n.name = name
}

// AddChildren attaches a child SceneNode downstream, establishing a parent-child relationship.
func (s *Node) AddChildren(child SceneNode) {
	child.AttachParent(s)
	s.children = append(s.children, child)
}

// DetachChild locates and unbinds a specific child SceneNode from this node,
// returning true on success or false if it isn't listed as a child.
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

// AttachParent establishes an upward connection configuring this node's parent SceneNode.
func (s *Node) AttachParent(node SceneNode) {
	s.parent = node
}

// GetParent provides the current parent SceneNode of this node.
func (b *Node) GetParent() SceneNode {
	return b.parent
}

// GetChildren fetches all child elements bound beneath this node.
func (s *Node) GetChildren() []SceneNode {
	return s.children
}

// Delete cleanly removes this operational node entirely from its active parent.
func (s *Node) Delete() {
	s.parent.DetachChild(s)
}

// MarkDirty triggers a recursive flag through this node and its children notifying them of updates required.
func (s *Node) MarkDirty() {
	for _, child := range s.children {
		child.MarkDirty()
	}
}
