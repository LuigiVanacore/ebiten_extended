package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestCollisionOrientedRect_Circle(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	circ := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 15))
	rect := NewCollisionOrientedRectFromSize(30, 20)
	colCirc, _ := NewCollider("c", circ, mask)
	colRect, _ := NewCollider("r", rect, mask)
	colCirc.SetPosition(0, 0)
	colRect.SetPosition(50, 0)
	colRect.SetRotation(0)
	tA := colCirc.GetWorldTransform()
	tB := colRect.GetWorldTransform()
	if ShapeCollides(circ, tA, rect, tB) {
		t.Error("circle at 0,0 should not hit oriented rect at 50,0")
	}
	colCirc.SetPosition(50, 0)
	tA = colCirc.GetWorldTransform()
	tB = colRect.GetWorldTransform()
	if !ShapeCollides(circ, tA, rect, tB) {
		t.Error("circle at 50,0 should hit rect center 50,0")
	}
}

func TestCollisionOrientedRect_OverlapPoint(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	rect := NewCollisionOrientedRectFromSize(25, 25)
	col, _ := NewCollider("r", rect, mask)
	col.SetPosition(100, 100)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	results := mgr.OverlapPoint(math2D.NewVector2D(100, 100))
	if len(results) != 1 {
		t.Errorf("point at center: got %d overlaps", len(results))
	}
}
