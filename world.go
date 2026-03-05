package ebiten_extended

import (
	"slices"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

// DefaultLayerIndex is the layer index used by AddNodeToDefaultLayer.
const DefaultLayerIndex = 0

// World represents the main game world, managing the scene graph, layers, and camera.
// It handles updating and drawing all nodes within the game.
type World struct {
	rootScene  SceneNode
	layerRoots []SceneNode // layerRoots[i] = root node for layer index i
	drawLayers *Layers
	camera     *Camera
	postUpdate func() // called after updateNode; e.g. set to collision.CollisionManager().CheckCollision to avoid import cycle
}

// NewWorld creates and initializes a new World instance.
func NewWorld() *World {
	w, h := ebiten.WindowSize()
	return &World{
		rootScene:  NewNode("root_world"),
		layerRoots: make([]SceneNode, 0),
		drawLayers: NewLayers(),
		camera:     NewCamera(uint(w), uint(h)),
	}
}

// Camera returns the camera associated with the world.
func (world *World) Camera() *Camera {
	return world.camera
}

// AddNodeToLayer adds a node to the layer at the given index.
// Lower layer indices are drawn first (background). The index is created if it doesn't exist.
func (world *World) AddNodeToLayer(node SceneNode, layerIndex int) {
	if layerIndex < 0 {
		return
	}
	for layerIndex >= len(world.layerRoots) {
		root := NewNode("layer_root")
		world.layerRoots = append(world.layerRoots, root)
		world.rootScene.AddChildren(root)
	}
	world.layerRoots[layerIndex].AddChildren(node)
}

// AddNodeToDefaultLayer adds the node to the default layer (index 0).
func (world *World) AddNodeToDefaultLayer(node SceneNode) {
	world.AddNodeToLayer(node, DefaultLayerIndex)
}

// SetPostUpdate sets a callback run after each Update (e.g. collision.CollisionManager().CheckCollision).
func (world *World) SetPostUpdate(f func()) {
	world.postUpdate = f
}

// Update progresses the game state by one tick.
// It recursively updates all nodes in the scene graph and runs the postUpdate callback if set.
func (world *World) Update() {
	world.camera.Update()
	world.updateNode(world.rootScene)
	if world.postUpdate != nil {
		world.postUpdate()
	}
}

// updateNode recursively calls Update on the given node and its children (pre-order:
// parent updates before its children so children can read up-to-date parent state).
func (world *World) updateNode(node SceneNode) {
	if node == nil {
		return
	}
	if entity, ok := node.(Updatable); ok {
		entity.Update()
	}
	for _, child := range node.GetChildren() {
		world.updateNode(child)
	}
}

// queueNodeToLayers traverses the node subtree and queues draw callbacks to drawLayers.
func (world *World) queueNodeToLayers(node SceneNode, parentGeoM ebiten.GeoM, layerIndex int) {
	if node == nil {
		return
	}
	op := ebiten.DrawImageOptions{}
	if entity, ok := node.(transform.Transformable); ok {
		op.GeoM = updateTransform(entity, parentGeoM)
	}
	children := node.GetChildren()
	// Sort descending: we push to a stack (LIFO), so higher GetLayer first => drawn last (on top).
	// slices.SortFunc avoids the closure allocation of sort.Slice.
	slices.SortFunc(children, func(a, b SceneNode) int {
		la, lb := 0, 0
		if d, ok := a.(Drawable); ok {
			la = d.GetLayer()
		}
		if d, ok := b.(Drawable); ok {
			lb = d.GetLayer()
		}
		return lb - la
	})
	for _, child := range children {
		world.queueNodeToLayers(child, op.GeoM, layerIndex)
	}
	if _, ok := node.(transform.Transformable); ok {
		if drawable, ok := node.(Drawable); ok {
			childOp := op
			world.camera.ApplyRelativeTranslation(&childOp, 0, 0)
			_ = world.drawLayers.AddNodeToLayer(layerIndex, drawable, world.camera.GetSurface(), childOp)
		}
	}
}

// Draw renders the world onto the target image.
// It clears the camera surface, queues all nodes to layers by traverse order, draws layers, and then draws the camera.
func (world *World) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	world.camera.surface.Clear()
	baseGeoM := ebiten.GeoM{}
	for i := range world.layerRoots {
		world.queueNodeToLayers(world.layerRoots[i], baseGeoM, i)
	}
	world.drawLayers.DrawLayers()
	world.camera.DrawWithOptions(target, op)
}

// updateTransform calculates the updated geometric matrix (GeoM) for an entity
// based on its transform (position, rotation, scale, pivot) and its parent's transform matrix.
// parentGeoM is not mutated so siblings are positioned correctly.
func updateTransform(entity transform.Transformable, parentGeoM ebiten.GeoM) ebiten.GeoM {
	updated := ebiten.GeoM{}
	tr := entity.GetTransform()
	position := tr.GetPosition()
	pivot := tr.GetPivot()
	rotation := tr.GetRotation()
	scale := tr.GetScale()

	updated.Translate(-pivot.X(), -pivot.Y())
	updated.Scale(scale.X(), scale.Y())
	updated.Rotate(rotation)
	updated.Translate(pivot.X(), pivot.Y())
	updated.Translate(position.X(), position.Y())
	updated.Concat(parentGeoM)

	return updated
}
