package ebiten_extended

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

func TestNewNode2D(t *testing.T) {
	node := NewNode2D("TestNode")
	if node == nil {
		t.Fatal("Expected NewNode2D to return a non-nil value")
	}
	if node.name != "TestNode" {
		t.Errorf("Expected node name to be 'TestNode', got '%s'", node.name)
	}
}

func TestSetAndGetTransform(t *testing.T) {
	node := NewNode2D("TestNode")
	newTransform := transform.Transform{}
	node.SetTransform(newTransform)

	if node.GetTransform() != newTransform {
		t.Error("SetTransform or GetTransform did not work as expected")
	}
}

func TestGetWorldTransform(t *testing.T) {
	root := NewNode2D("Root")
	child := NewNode2D("Child")
	root.AddChild(child)

	rootTransform := transform.Transform{}
	root.SetTransform(rootTransform)

	worldTransform := child.GetWorldTransform()
	if worldTransform != rootTransform {
		t.Error("GetWorldTransform did not return the expected transform")
	}
}

func TestGetWorldPosition(t *testing.T) {
	root := NewNode2D("Root")
	child := NewNode2D("Child")
	root.AddChild(child)

	rootTransform := transform.Transform{}
	rootTransform.SetPosition(math2D.NewVector2D(10, 20))
	root.SetTransform(rootTransform)

	worldPosition := child.GetWorldPosition()
	expectedPosition := math2D.NewVector2D(10, 20)
	if worldPosition != expectedPosition {
		t.Errorf("Expected world position to be %v, got %v", expectedPosition, worldPosition)
	}
}