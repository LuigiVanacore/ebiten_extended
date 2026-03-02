package ebiten_extended

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

// World represents the main game world, managing the scene graph, layers, and camera.
// It handles updating and drawing all nodes within the game.
type World struct {
	rootScene  SceneNode
	layers     []*Layer
	camera     *Camera
	postUpdate func() // called after updateNode; e.g. set to collision.CollisionManager().CheckCollision to avoid import cycle
}

// NewWorld creates and initializes a new World instance.
// It sets up the root scene node, initializes the layer slice, and creates a camera
// matching the current window size.
func NewWorld() *World {
	w, h := ebiten.WindowSize()
	return &World{layers: make([]*Layer, 0), rootScene: &Node{id: 0, name: "root", parent: nil}, camera: NewCamera(uint(w), uint(h))}
}

// Camera returns the camera associated with the world.
func (world *World) Camera() *Camera {
	return world.camera
}

// AddLayer adds a new Layer to the world's layer list and rebuilds the scene graph.
func (world *World) AddLayer(layer *Layer) {
	world.layers = append(world.layers, layer)
	world.buildScene()
}

// searchLayer looks for a layer by its ID and returns it, or an error if not found.
func (world *World) searchLayer(layerId int) (*Layer, error) {
	err := world.checkLayerId(layerId)
	if err != nil {
		return nil, err
	}
	for _, layer := range world.layers {
		if layer.GetId() == layerId {
			return layer, nil
		}
	}
	return nil, fmt.Errorf("error: layer not found %d", layerId)
}

// SetLayerPriority searches for a layer by its ID and updates its priority (rendering order/Z-index).
func (world *World) SetLayerPriority(layerId int, priority int) {
	layer, err := world.searchLayer(layerId)
	if err != nil {
		fmt.Println(err)
		return
	}
	layer.SetPriority(priority)
}

// GetOrCreateDefaultLayer returns a layer with MinLayerID, creating it if needed.
func (world *World) GetOrCreateDefaultLayer() *Layer {
	for _, l := range world.layers {
		if l.GetId() == MinLayerID {
			return l
		}
	}
	layer := NewLayer(MinLayerID, 0, "default")
	world.AddLayer(layer)
	return layer
}

// AddNodeToDefaultLayer adds the node to the default layer of the current world.
func (world *World) AddNodeToDefaultLayer(node SceneNode) {
	world.GetOrCreateDefaultLayer().AddNode(node)
}

// MinLayerID is the minimum valid layer id (0 and 1 are reserved).
const MinLayerID = 2

// checkLayerId validates a layer ID, ensuring it is greater than or equal to MinLayerID.
func (world *World) checkLayerId(layerID int) error {
	if layerID < MinLayerID {
		return fmt.Errorf("invalid layer id %d: must be >= %d", layerID, MinLayerID)
	}
	return nil
}

// buildScene attaches the most recently added layer's root scene to the world's root scene.
func (world *World) buildScene() {
	layer := world.layers[len(world.layers)-1]
	world.rootScene.AddChildren(layer.GetRootScene())
}

// SetPostUpdate sets a callback run after each Update (e.g. collision.CollisionManager().CheckCollision).
func (world *World) SetPostUpdate(f func()) {
	world.postUpdate = f
}

// Update progresses the game state by one tick.
// It recursively updates all nodes in the scene graph and runs the postUpdate callback if set.
func (world *World) Update() {
	world.updateNode(world.rootScene)
	if world.postUpdate != nil {
		world.postUpdate()
	}
}

// updateNode recursively calls Update on the given node and its children,
// provided they implement the Updatable interface.
func (world *World) updateNode(node SceneNode) {
	if node == nil {
		return
	}
	for _, child := range node.GetChildren() {
		world.updateNode(child)
	}
	if entity, ok := node.(Updatable); ok {
		entity.Update()
	}
}

// Draw renders the world onto the target image.
// It clears the camera surface, draws the scene graph, and then draws the camera.
func (world *World) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	world.camera.surface.Clear()
	world.DrawNode(world.rootScene, target, op)
	world.camera.Draw(target)
}

// DrawNode recursively draws the given node and its children.
// It applies the transform matrix of each node to position, rotate, and scale it correctly.
func (world *World) DrawNode(node SceneNode, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if entity, ok := node.(transform.Transformable); ok {
		op.GeoM = updateTransform(entity, op.GeoM)
		transform := entity.GetTransform()
		position := transform.GetPosition()

		if drawable, ok := node.(Drawable); ok {
			childOp := *op
			world.camera.ApplyRelativeTranslation(&childOp, position.X(), position.Y())
			drawable.Draw(world.camera.GetSurface(), &childOp)
		}
	}
	parentGeoM := op.GeoM
	for _, child := range node.GetChildren() {
		op.GeoM = parentGeoM
		world.DrawNode(child, target, op)
	}
}

// updateTransform calculates the updated geometric matrix (GeoM) for an entity
// based on its transform (position, rotation, pivot) and its parent's transform matrix.
func updateTransform(entity transform.Transformable, parent_geoM ebiten.GeoM) ebiten.GeoM {
	updated_GeoM := ebiten.GeoM{}
	transform := entity.GetTransform()
	position := transform.GetPosition()
	pivot := transform.GetPivot()
	rotation := transform.GetRotation()

	updated_GeoM.Translate(-pivot.X(), -pivot.Y())
	updated_GeoM.Rotate(rotation)
	updated_GeoM.Translate(pivot.X(), pivot.Y())

	parent_geoM.Translate(position.X(), position.Y())

	updated_GeoM.Concat(parent_geoM)

	return updated_GeoM
}
