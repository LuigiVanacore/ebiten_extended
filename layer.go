package ebiten_extended



type Layer struct {
	id int
	priority int
	rootScene SceneNode
}

func NewLayer(id int, priority int) *Layer {
	return &Layer{ id: id, priority: priority, rootScene: NewBaseNode("root")}
}


func (l *Layer) GetId() int {
	return l.id
}

func (l *Layer) SetId(id int) {
	l.id = id
}

func (l *Layer) SetPriority(priority int) {
	l.priority = priority
}

func (l *Layer) GetPriority() int {
	return l.priority
}

func (l *Layer) AddNode( node SceneNode) {
	l.rootScene.AddChildren(node)
}

func (l *Layer) GetRootScene() SceneNode {
	return l.rootScene
}

