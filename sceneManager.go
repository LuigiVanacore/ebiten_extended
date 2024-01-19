package ebiten_extended

import (
	"fmt"

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
	sceneManager := &sceneManager{layers: make([]*Layer, 0)}
	sceneManager.incrementNextIdVal()
	return sceneManager
}

func (sceneManager *sceneManager) initSceneManager() {
	sceneManager.AddLayer(NewLayer(UI_LAYER, UI_LAYER))
	sceneManager.AddLayer(NewLayer(DEFAULT_LAYER, DEFAULT_LAYER))
	sceneManager.buildScene()
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
}

func (sceneManager *sceneManager) AddEntity(entity any, name string, layerId int) *SceneNode {
	sceneNode := NewSceneNode(entity, name)
	layer, error := sceneManager.searchLayer(layerId)
	if error != nil {
		fmt.Println(error)
		fmt.Printf("layer %d not found, added node %s to defaul layer", layerId, name)
		sceneManager.layers[DEFAULT_LAYER].AddNode(sceneNode)
		return sceneNode
	}
	layer.AddNode(sceneNode)
	return sceneNode
}

func (sceneManager *sceneManager) AddEntityToDefaultLayer(entity any, name string) *SceneNode{
	sceneNode := NewSceneNode(entity, name)
	
	sceneManager.layers[DEFAULT_LAYER].AddNode(sceneNode)

	return sceneNode
}

func (sceneManager *sceneManager) AddSceneNodeToDefaultLayer(sceneNode *SceneNode) {
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
	for _, layer := range sceneManager.layers {
		sceneManager.rootScene.Addhildren(layer.GetRootScene())
	}
}

func (sceneManager *sceneManager) Update() {
	sceneManager.rootScene.Update()
}

func (sceneManager *sceneManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	sceneManager.rootScene.Draw(target, op)
}