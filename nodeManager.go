package ebiten_extended

var nodeManagerInstance *nodeManager

func NodeManager() *nodeManager {
	if nodeManagerInstance == nil {
		nodeManagerInstance = newNodeManager()
	}

	return nodeManagerInstance
}

type nodeManager struct {
	nextIdVal uint64
}

func newNodeManager() *nodeManager {
	nodeManager := &nodeManager{}
	nodeManager.incrementNextIdVal()
	return nodeManager
}

func (nodeManager *nodeManager) GetNextIdVal() uint64 {
	nextIdVal := nodeManager.nextIdVal
	nodeManager.incrementNextIdVal()
	return nextIdVal
}

func (nodeManager *nodeManager) setNextIdVal(nextIdVal uint64) *nodeManager {
	nodeManager.nextIdVal = nextIdVal
	return nodeManager
}

func (nodeManager *nodeManager) incrementNextIdVal() {
	nodeManager.setNextIdVal(nodeManager.nextIdVal + 1)
}
