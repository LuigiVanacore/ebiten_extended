package ebiten_extended



type SceneNode interface {
	AddChildren(child SceneNode)
	GetChildren() []SceneNode
	AttachParent(node SceneNode)
	GetParent() SceneNode
	DetachChild(node SceneNode) bool
}
