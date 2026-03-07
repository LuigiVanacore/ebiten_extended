package ebiten_extended

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// recordDrawable captures the GeoM translation when Draw is called.
type recordDrawable struct {
	Node2D
	layer     int
	recordedX float64
	recordedY float64
	drew      bool
}

func (r *recordDrawable) GetLayer() int { return r.layer }
func (r *recordDrawable) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	r.recordedX = op.GeoM.Element(0, 2)
	r.recordedY = op.GeoM.Element(1, 2)
	r.drew = true
}

func TestUpdateTransformSiblings(t *testing.T) {
	// Test that sibling nodes get correct independent transforms (parentGeoM not mutated).
	// Parent at (10, 20), child1 at (5, 0), child2 at (0, 7).
	// Expected world positions: child1=(15,20), child2=(10,27).
	world := NewWorld()
	world.Camera().SetPosition(0, 0)

	parent := NewNode2D("parent")
	parent.SetPosition(10, 20)

	child1 := &recordDrawable{Node2D: *NewNode2D("child1"), layer: 0}
	child1.SetPosition(5, 0)

	child2 := &recordDrawable{Node2D: *NewNode2D("child2"), layer: 0}
	child2.SetPosition(0, 7)

	parent.AddChildren(child1)
	parent.AddChildren(child2)

	world.AddNodeToLayer(parent, DefaultLayerIndex)

	target := ebiten.NewImage(320, 240)
	op := &ebiten.DrawImageOptions{}
	world.Draw(target, op)

	if !child1.drew || !child2.drew {
		t.Fatal("expected both children to be drawn")
	}

	// Camera at (0,0) so world coords = screen coords after ApplyRelativeTranslation
	tol := 0.001
	if !floatEqual(child1.recordedX, 15, tol) || !floatEqual(child1.recordedY, 20, tol) {
		t.Errorf("child1 expected position (15, 20), got (%.2f, %.2f)", child1.recordedX, child1.recordedY)
	}
	if !floatEqual(child2.recordedX, 10, tol) || !floatEqual(child2.recordedY, 27, tol) {
		t.Errorf("child2 expected position (10, 27), got (%.2f, %.2f)", child2.recordedX, child2.recordedY)
	}
}

func floatEqual(a, b, tol float64) bool {
	if a < b-tol || a > b+tol {
		return false
	}
	return true
}

func TestLayersByPriority(t *testing.T) {
	world := NewWorld()

	// Layer index = draw order (0 first, 2 last)
	drawOrder := make([]int, 0, 3)
	recA := &orderRecorderDrawable{Node2D: *NewNode2D("a"), layer: 0, id: 100, order: &drawOrder}
	recB := &orderRecorderDrawable{Node2D: *NewNode2D("b"), layer: 0, id: 0, order: &drawOrder}
	recC := &orderRecorderDrawable{Node2D: *NewNode2D("c"), layer: 0, id: 50, order: &drawOrder}

	world.AddNodeToLayer(recB, 0) // drawn first
	world.AddNodeToLayer(recC, 1) // drawn second
	world.AddNodeToLayer(recA, 2) // drawn last

	target := ebiten.NewImage(320, 240)
	world.Draw(target, &ebiten.DrawImageOptions{})

	// Lower priority drawn first: B(0), C(50), A(100)
	expected := []int{0, 50, 100}
	if len(drawOrder) != len(expected) {
		t.Fatalf("expected %d draws, got %d", len(expected), len(drawOrder))
	}
	for i := range expected {
		if drawOrder[i] != expected[i] {
			t.Errorf("draw order: expected %v, got %v", expected, drawOrder)
			break
		}
	}
}

type orderRecorderDrawable struct {
	Node2D
	layer int
	id    int
	order *[]int
}

func (o *orderRecorderDrawable) GetLayer() int { return o.layer }
func (o *orderRecorderDrawable) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	*o.order = append(*o.order, o.id)
}

func TestWorldRemoveNode(t *testing.T) {
	world := NewWorld()
	node := NewNode2D("test")
	world.AddNodeToDefaultLayer(node)

	if !world.RemoveNode(node) {
		t.Error("RemoveNode should return true when node has parent")
	}
	if node.GetParent() != nil {
		t.Error("node parent should be nil after remove")
	}
	if world.RemoveNode(node) {
		t.Error("RemoveNode should return false when node has no parent")
	}
}

func TestWorldClearLayer(t *testing.T) {
	world := NewWorld()
	a := NewNode2D("a")
	b := NewNode2D("b")
	world.AddNodeToLayer(a, 0)
	world.AddNodeToLayer(b, 0)

	world.ClearLayer(0)

	if a.GetParent() != nil || b.GetParent() != nil {
		t.Error("nodes should have nil parent after ClearLayer")
	}
	// Clearing invalid index should not panic
	world.ClearLayer(-1)
	world.ClearLayer(100)
}

func TestDrawableGetLayerSiblingOrder(t *testing.T) {
	// Siblings with different GetLayer should be drawn in ascending order (lower first).
	world := NewWorld()
	drawOrder := make([]int, 0, 3)
	mk := func(id, layer int) *orderRecorderDrawable {
		return &orderRecorderDrawable{Node2D: *NewNode2D(""), layer: layer, id: id, order: &drawOrder}
	}
	parent := NewNode2D("parent")
	parent.AddChildren(mk(100, 2)) // drawn last (highest GetLayer)
	parent.AddChildren(mk(0, 0))   // drawn first (lowest GetLayer)
	parent.AddChildren(mk(50, 1))  // drawn second

	world.AddNodeToLayer(parent, DefaultLayerIndex)

	target := ebiten.NewImage(320, 240)
	world.Draw(target, &ebiten.DrawImageOptions{})

	expected := []int{0, 50, 100}
	if len(drawOrder) != len(expected) {
		t.Fatalf("expected %d draws, got %d", len(expected), len(drawOrder))
	}
	for i := range expected {
		if drawOrder[i] != expected[i] {
			t.Errorf("GetLayer sibling order: expected %v, got %v", expected, drawOrder)
			break
		}
	}
}
