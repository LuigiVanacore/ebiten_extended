package ebiten_extended





var sceneManagerinstance *sceneManager

func SceneManager() *sceneManager {
	if sceneManagerinstance == nil {
		sceneManagerinstance = newSceneManager()
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


// func (sceneManager *sceneManager) AddScene(scene *Scene) {
// 	sceneManager.scenes = append(sceneManager.scenes, scene)
// }



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

// func (sceneManager *sceneManager) Update() {
// 	sceneManager.
// }


// func (sceneManager *sceneManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
// sceneManager.}
