package ebiten_extended


 
var nodeManagerinstance *nodeManager

func NodeManager() *nodeManager {
	if nodeManagerinstance == nil {
		nodeManagerinstance = newNodeManager()
	}

	return nodeManagerinstance
}

type nodeManager struct {
	nextIdVal uint64
}


func newNodeManager() *nodeManager {
	nodeManager := &nodeManager{}
	nodeManager.incrementNextIdVal()
	return nodeManager
}

func (nodemanager *nodeManager) GetNextIdVal() uint64 {
	nextIdVal := nodemanager.nextIdVal
	nodemanager.incrementNextIdVal()
	return nextIdVal
}

func (nodemanager *nodeManager) setNextIdVal(nextIdVal uint64) *nodeManager {
	nodemanager.nextIdVal = nextIdVal
	return nodemanager
}

func (nodeManager *nodeManager) incrementNextIdVal() {
	nodeManager.setNextIdVal(nodeManager.nextIdVal+1)
}