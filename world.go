package ebiten_extended

import (
	"fmt"
	"math"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	rootScene Node
	layers    Layers
	camera    *Camera
	nextIdVal uint64
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

func (world *World) checkLayerId(layerID int) error {
	if layerID < 2 {
		return fmt.Errorf("Invalid Layer: the layer id is %d but must be >= 2 to be valid", layerID)
	}
	return nil
}

func (world *World) buildScene() {
	layer := world.layers[len(world.layers)-1]
	world.rootScene.AddChild(layer.GetRootScene())
}

func (world *World) Update() {
	world.updateNode(world.rootScene)
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
	world.DrawNode(world.rootScene, target, op)
	world.camera.Draw(target)
}

func (world *World) DrawNode(node Node, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	//playerOps := &ebiten.DrawImageOptions{}
	//playerOps = cam.GetTranslation(playerOps, PlayerX, PlayerY)
	//cam.DrawImage(player, playerOps)

	// Draw to screen and zoom
	if entity, ok := node.(transform.Transformable); ok {
		op.GeoM = updateTransform(entity, op.GeoM)
		transform := entity.GetTransform()
		position := transform.GetPosition()

		if entity, ok := node.(Drawable); ok {
			entity.Draw(world.camera.GetSurface(), world.camera.GetRelativeTranslation(op, position.X(), position.Y()))
		}
	}
	for _, child := range node.GetChildren() {
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
