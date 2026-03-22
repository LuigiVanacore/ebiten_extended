package ludum

// SceneNode defines the interface for an element within the game's hierarchical scene graph.
type SceneNode interface {
	// GetID retrieves the unique node identifier.
	GetID() uint64
	// GetName retrieves the node name.
	GetName() string
	// AddChildren attaches a child node to this node.
	AddChildren(child SceneNode)
	// GetChildren retrieves all child nodes assigned to this node.
	GetChildren() []SceneNode
	// AttachParent establishes an upward connection to another SceneNode.
	AttachParent(node SceneNode)
	// GetParent retrieves the current parent of this node.
	GetParent() SceneNode
	// DetachChild removes a specific child node, returning true if successful.
	DetachChild(node SceneNode) bool
	// MarkDirty flags this node and its descendants for updates.
	MarkDirty()
}
