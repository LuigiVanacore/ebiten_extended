package ebiten_extended

// Layer represents a logical, ordered grouping of scene nodes within the world, rendered by priority.
type Layer struct {
	id        int
	priority  int
	name string
	rootScene SceneNode
}

// NewLayer generates a new Layer assigned with a unique ID, sorting priority, and name.
func NewLayer(id int, priority int, name string) *Layer {
	return &Layer{id: id, priority: priority, name: name, rootScene: NewNode("root_"+name)}
}

// GetId returns the integer identifier for this layer.
func (l *Layer) GetId() int {
	return l.id
}

// SetId updates the numeric identifier of this layer.
func (l *Layer) SetId(id int) {
	l.id = id
}

// SetPriority alters the rendering Z-order position of this layer relative to others.
func (l *Layer) SetPriority(priority int) {
	l.priority = priority
}

// GetPriority resolves the current rendering sequence value of this layer.
func (l *Layer) GetPriority() int {
	return l.priority
}

// AddNode binds a SceneNode component to this layer's unified root scene.
func (l *Layer) AddNode(node SceneNode) {
	l.rootScene.AddChildren(node)
}

// GetRootScene retrieves the baseline underlying node serving as root for layer geometries.
func (l *Layer) GetRootScene() SceneNode {
	return l.rootScene
}
