package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestArea2D_OnBodyEntered(t *testing.T) {
	manager := NewCollisionManager()
	areaShape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(100, 100))}
	area := NewArea2D(areaShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	area.SetPosition(100, 100)

	collShape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20))}
	coll := NewCollider(collShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	coll.SetPosition(120, 120)

	manager.AddCollider(coll)
	manager.AddParticipant(area)

	entered := false
	area.OnBodyEntered.Connect(nil, func(ev Area2DBodyEvent) {
		entered = true
	})

	manager.CheckCollision()

	if !entered {
		t.Error("OnBodyEntered should fire when body overlaps Area2D")
	}
}

func TestArea2D_OnBodyExited(t *testing.T) {
	manager := NewCollisionManager()
	areaShape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(100, 100))}
	area := NewArea2D(areaShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	area.SetPosition(100, 100)

	collShape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20))}
	coll := NewCollider(collShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	coll.SetPosition(120, 120)

	manager.AddCollider(coll)
	manager.AddParticipant(area)

	exited := false
	area.OnBodyExited.Connect(nil, func(ev Area2DBodyEvent) {
		exited = true
	})

	manager.CheckCollision() // Enter
	coll.SetPosition(500, 500)
	manager.CheckCollision() // Exit

	if !exited {
		t.Error("OnBodyExited should fire when body leaves Area2D")
	}
}

func TestArea2D_NoOverlapNoEvent(t *testing.T) {
	manager := NewCollisionManager()
	areaShape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(50, 50))}
	area := NewArea2D(areaShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	area.SetPosition(0, 0)

	collShape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(10, 10))}
	coll := NewCollider(collShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	coll.SetPosition(500, 500)

	manager.AddCollider(coll)
	manager.AddParticipant(area)

	entered := false
	area.OnBodyEntered.Connect(nil, func(ev Area2DBodyEvent) {
		entered = true
	})

	manager.CheckCollision()

	if entered {
		t.Error("OnBodyEntered should not fire when bodies do not overlap")
	}
}
