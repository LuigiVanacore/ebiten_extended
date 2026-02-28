package ebiten_extended

import (
	"fmt"
	"math"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)



type World struct {
	rootScene SceneNode
	layers    []*Layer
	camera    *Camera
	postUpdate func() // called after updateNode; e.g. set to collision.CollisionManager().CheckCollision to avoid import cycle
}

func NewWorld() *World {
	w, h := ebiten.WindowSize()
	return &World{layers: make([]*Layer, 0), rootScene: &Node{id: 0, name: "root", parent: nil}, camera: NewCamera(uint(w), uint(h))}
}

func (world *World) Camera() *Camera {
	return world.camera
}

func (world *World) AddLayer(layer *Layer) {
	world.layers = append(world.layers, layer)
	world.buildScene()
}

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

func (world *World) checkLayerId(layerID int) error {
	if layerID < MinLayerID {
		return fmt.Errorf("invalid layer id %d: must be >= %d", layerID, MinLayerID)
	}
	return nil
}



func (world *World) buildScene() {
	layer := world.layers[len(world.layers)-1]
	world.rootScene.AddChildren(layer.GetRootScene())
}

// SetPostUpdate sets a callback run after each Update (e.g. collision.CollisionManager().CheckCollision).
func (world *World) SetPostUpdate(f func()) {
	world.postUpdate = f
}

func (world *World) Update() {
	world.updateNode(world.rootScene)
	if world.postUpdate != nil {
		world.postUpdate()
	}
}

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

func (world *World) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	world.camera.surface.Clear()
	world.DrawNode(world.rootScene,target, op)
	world.camera.Draw(target)
}

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

func updateTransform(entity transform.Transformable, parent_geoM ebiten.GeoM) ebiten.GeoM {
	updated_GeoM := ebiten.GeoM{}
	transform := entity.GetTransform()
	position := transform.GetPosition()
	pivot := transform.GetPivot()
	rotation := transform.GetRotation()

	updated_GeoM.Translate(-pivot.X(), -pivot.Y())
	updated_GeoM.Rotate(float64(rotation%360) * 2 * math.Pi / 360)
	updated_GeoM.Translate(pivot.X(), pivot.Y())

	parent_geoM.Translate(position.X(), position.Y())

	updated_GeoM.Concat(parent_geoM)

	return updated_GeoM
}

