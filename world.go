package ebiten_extended

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	root *Node
	uiRoot *Node
	camera    *Camera
	layers  *Layers
}

func NewWorld() *World {
	w, h := ebiten.WindowSize()
	return &World{ root: NewNode("root_world"), uiRoot: NewNode("ui_root"), camera: NewCamera(uint(w), uint(h)), layers: NewLayers()}
}

func (world *World) Camera() *Camera {
	return world.camera
}



// func (world *World) searchLayer(layerId int) (*Layer, error) {
// 	err := world.checkLayerId(layerId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, layer := range world.layers {
// 		if layer.GetId() == layerId {
// 			return layer, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("error: layer not found %d", layerId)
// }

// func (world *World) SetLayerPriority(layerId int, priority int) {
// 	layer, err := world.searchLayer(layerId)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	layer.SetPriority(priority)
// }

// func (world *World) checkLayerId(layerID int) error {
// 	if layerID < 2 {
// 		return fmt.Errorf("Invalid Layer: the layer id is %d but must be >= 2 to be valid", layerID)
// 	}
// 	return nil
// }

// func (world *World) buildScene() {
// 	layer := world.layers[len(world.layers)-1]
// 	world.rootScene.AddChild(layer.GetRootScene())
// }

func (world *World) AddNode( node SceneNode) {
	world.root.AddChild(node)
}

func (world *World) AddUINode(node SceneNode) {
	world.uiRoot.AddChild(node)
}

func (world *World) Update() {
	world.updateNode(world.root)
	world.updateNode(world.uiRoot)
	world.camera.Update()
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
	world.DrawNode(world.root, target, op)
	world.layers.DrawLayers()
	world.camera.Draw(target)
	world.DrawUINode(world.uiRoot, target, op)
}

// func (world *World) DrawNode(node SceneNode,target *ebiten.Image, op *ebiten.DrawImageOptions) {
// 	if node == nil  {
// 		return
// 	}
// 	// Draw to screen and zoom
// 	parent_op := *op
// 	if entity, ok := node.(transform.Transformable); ok {
// 	    parent_op.GeoM = updateTransform(entity, parent_op.GeoM)

// 		if entity, ok := node.(Drawable); ok {
// 			world.layers.AddNodeToLayer(entity.GetLayer(), entity, target, parent_op)
// 		}
// 	}
// 	for _, child := range node.GetChildren() {
// 		world.DrawNode(child, target, &parent_op)
// 	}
// }

func (world *World) DrawNode(node SceneNode,target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if node == nil  {
		return
	}
	// Draw to screen and zoom
	parent_op := *op
	if entity, ok := node.(transform.Transformable); ok {
	    parent_op.GeoM = updateTransform(entity, parent_op.GeoM)

		if entity, ok := node.(Drawable); ok {
			f := func() { 
				entity.Draw(world.camera.surface, world.camera.GetRelativeTranslation(&parent_op, 0, 0))
			}
			world.layers.AddNodeToLayerF(entity.GetLayer(), entity, f)
		}
	}
	for _, child := range node.GetChildren() {
		world.DrawNode(child, target, &parent_op)
	}
}

func (world *World) DrawUINode(node SceneNode, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if node == nil  {
		return
	}
	// Draw to screen and zoom
	parent_op := *op
	if entity, ok := node.(transform.Transformable); ok {
	    parent_op.GeoM = updateTransform(entity, parent_op.GeoM)

		if entity, ok := node.(Drawable); ok {
			entity.Draw(target, &parent_op)
		}
	}
	for _, child := range node.GetChildren() {
		world.DrawUINode(child, target, &parent_op)
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

	updated_GeoM.Translate(position.X(), position.Y())
	updated_GeoM.Concat(parent_geoM)

	return updated_GeoM
}
