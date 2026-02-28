package ebiten_extended

import (
	"testing"
)

func TestNodeIDGeneration(t *testing.T) {
	// The id should increment atomically, no singletons involved.
	node1 := NewNode("Node1")
	node2 := NewNode("Node2")

	if node1.GetID() == 0 {
		t.Errorf("Node ID should not be zero or uninitialized")
	}
	
	if node1.GetID() == node2.GetID() {
		t.Errorf("Nodes generated identical IDs: %d", node1.GetID())
	}

	if node2.GetID() <= node1.GetID() {
		t.Errorf("Node IDs are not incrementing sequentially")
	}
}
