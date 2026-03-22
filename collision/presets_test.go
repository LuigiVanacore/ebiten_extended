package collision

import (
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/utils"
)

func TestPresetMasks(t *testing.T) {
	playerMask := NewPresetMask(LayerPlayer, LayerWorld, LayerPickup)
	if !playerMask.GetIdentity().Has(LayerPlayer) {
		t.Error("player identity should have LayerPlayer")
	}
	if !playerMask.GetMask().Has(LayerWorld) {
		t.Error("player mask should have LayerWorld")
	}
	worldMask := NewPresetMask(LayerWorld, LayerPlayer, LayerEnemy)
	if !playerMask.IsCollidable(worldMask) {
		t.Error("player should collide with world")
	}
	if !worldMask.IsCollidable(playerMask) {
		t.Error("world should collide with player")
	}
}

func TestPresetVars(t *testing.T) {
	playerCol, _ := NewCollider("player", NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 10)), MaskPlayer)
	worldCol, _ := NewCollider("world", NewCollisionRect(math2d.NewRectangle(math2d.ZeroVector2D(), math2d.NewVector2D(100, 100))), MaskWorld)
	if !playerCol.CanCollideWith(worldCol) {
		t.Error("player preset should collide with world preset")
	}
	// Pickup doesn't collide with enemy
	pickupCol, _ := NewCollider("pickup", NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 5)), MaskPickup)
	enemyCol, _ := NewCollider("enemy", NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 10)), MaskEnemy)
	if pickupCol.CanCollideWith(enemyCol) {
		t.Error("pickup should not collide with enemy")
	}
	// Use utils.ByteSet for custom mask compatible with preset
	custom := NewCollisionMask(utils.ByteSet(LayerPlayer), utils.ByteSet(LayerWorld))
	if !custom.IsCollidable(worldCol.GetCollisionMask()) {
		t.Error("custom player mask should collide with world")
	}
}
