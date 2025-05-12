package ebiten_extended

type SceneNode interface {
	AddChild(child SceneNode)
	GetChildren() []SceneNode
	AttachParent(Nodable SceneNode)
	GetParent() SceneNode
	DetachChild(Nodable SceneNode) bool
}
