package ebiten_extended

import (
	"fmt"
	"math"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UI_LAYER      = 0
	DEFAULT_LAYER = 1
)

var sceneManagerinstance *sceneManager

func SceneManager() *sceneManager {
	if sceneManagerinstance == nil {
		sceneManagerinstance = newSceneManager()
		sceneManagerinstance.initSceneManager()
	}

	return sceneManagerinstance
}

type sceneManager struct {
	nextIdVal uint64
	rootScene SceneNode
	layers    []*Layer
}

func newSceneManager() *sceneManager {
	sceneManager := &sceneManager{layers: make([]*Layer, 0), rootScene: &Node{id: 0, name: "root", parent: nil}}
	sceneManager.incrementNextIdVal()
	return sceneManager
}

func (sceneManager *sceneManager) initSceneManager() {
	sceneManager.AddLayer(NewLayer(UI_LAYER, UI_LAYER, "UI"))
	sceneManager.AddLayer(NewLayer(DEFAULT_LAYER, DEFAULT_LAYER, "default"))
}

func (sceneManager *sceneManager) GetNextIdVal() uint64 {
	nextIdVal := sceneManager.nextIdVal
	sceneManager.incrementNextIdVal()
	return nextIdVal
}

func (sceneManager *sceneManager) setNextIdVal(nextIdVal uint64) *sceneManager {
	sceneManager.nextIdVal = nextIdVal
	return sceneManager
}

func (sceneManager *sceneManager) incrementNextIdVal() {
	sceneManager.setNextIdVal(sceneManager.nextIdVal + 1)
}

func (sceneManager *sceneManager) AddLayer(layer *Layer) {
	sceneManager.layers = append(sceneManager.layers, layer)
	sceneManager.buildScene()
}

// func (sceneManager *sceneManager) AddEntity(entity any, name string, layerId int) *SceneNode {
// 	sceneNode := NewSceneNode(entity, name)
// 	layer, error := sceneManager.searchLayer(layerId)
// 	if error != nil {
// 		fmt.Println(error)
// 		fmt.Printf("layer %d not found, added node %s to defaul layer", layerId, name)
// 		sceneManager.layers[DEFAULT_LAYER].AddNode(sceneNode)
// 		return sceneNode
// 	}
// 	layer.AddNode(sceneNode)
// 	return sceneNode
// }

// func (sceneManager *sceneManager) AddNodeToDefaultLayer(entity any, name string) SceneNode{
// 	sceneNode := NewSceneNode(entity, name)

// 	sceneManager.layers[DEFAULT_LAYER].AddNode(sceneNode)

// 	return sceneNode
// }

func (sceneManager *sceneManager) AddSceneNodeToDefaultLayer(sceneNode SceneNode) {
	sceneManager.layers[DEFAULT_LAYER].AddNode(sceneNode)
}

func (sceneManager *sceneManager) searchLayer(layerId int) (*Layer, error) {
	err := checkLayerId(layerId)
	if err != nil {
		return nil, err
	}
	for _, layer := range sceneManager.layers {
		if layer.GetId() == layerId {
			return layer, nil
		}
	}
	return nil, fmt.Errorf("error: layer not found %d", layerId)
}

func (sceneManager *sceneManager) SetLayerPriority(layerId int, priority int) {

	layer, err := sceneManager.searchLayer(layerId)

	if err != nil {
		fmt.Println(err)
		return
	}
	layer.SetPriority(priority)
}

func checkLayerId(layerID int) error {
	if layerID < 2 {
		return fmt.Errorf("Invalid Layer: the layer id is %d but must be >= 2 to be valid", layerID)
	}
	return nil
}

func (sceneManager *sceneManager) buildScene() {
	layer := sceneManager.layers[len(sceneManager.layers)-1]
	sceneManager.rootScene.AddChildren(layer.GetRootScene())
	
}

func (sceneManager *sceneManager) Update() {
	sceneManager.updateNode(sceneManager.rootScene)
}

func (sceneManager *sceneManager) updateNode(node SceneNode) {
	if node == nil {
		return
	}
	for _, child := range node.GetChildren() {
		sceneManager.updateNode(child)

	}
	if entity, ok := node.(Updatable); ok {
		entity.Update()
	}
}

func (sceneManager *sceneManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	sceneManager.DrawNode(sceneManager.rootScene, target, op)
}

func (sceneManager *sceneManager) DrawNode(node SceneNode, target *ebiten.Image, op *ebiten.DrawImageOptions) {

	if entity, ok := node.(Drawable); ok {
		op.GeoM = updateTransform(entity, op.GeoM)
		entity.Draw(target, op)
	}
	for _, child := range node.GetChildren() {
		sceneManager.DrawNode(child, target, op)
	}
}

func updateTransform(entity transform.Transformable, parent_geoM ebiten.GeoM) ebiten.GeoM {
	rotation_geoM := ebiten.GeoM{}
	transform := entity.GetTransform()
	position := transform.GetPosition()
	pivot := transform.GetPivot()
	rotation := transform.GetRotation()

	rotation_geoM.Translate(-pivot.X(), -pivot.Y())
	rotation_geoM.Rotate(float64(rotation%360) * 2 * math.Pi / 360)
	rotation_geoM.Translate(pivot.X(), pivot.Y())

	parent_geoM.Translate(position.X(), position.Y())

	rotation_geoM.Concat(parent_geoM)

	return rotation_geoM
}
