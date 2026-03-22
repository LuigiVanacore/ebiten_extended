package ludum

import (
	"image/color"
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestLayersAddNodeToLayerF_InvalidIndex(t *testing.T) {
	layers := NewLayers()

	err := layers.AddNodeToLayerF(-1, func() {})
	if err != ErrInvalidLayerIndex {
		t.Errorf("expected ErrInvalidLayerIndex, got %v", err)
	}
}

func TestLayersAddNodeToLayer_InvalidIndex(t *testing.T) {
	layers := NewLayers()
	rect := NewDrawnRectangle("r", math2d.ZeroVector2D(), math2d.ZeroVector2D(), color.White, false, 0)
	target := ebiten.NewImage(1, 1)
	op := ebiten.DrawImageOptions{}

	err := layers.AddNodeToLayer(-1, rect, target, op)
	if err != ErrInvalidLayerIndex {
		t.Errorf("expected ErrInvalidLayerIndex, got %v", err)
	}
}

func TestLayersAddNodeToLayerF_ValidIndex(t *testing.T) {
	layers := NewLayers()

	err := layers.AddNodeToLayerF(0, func() {})
	if err != nil {
		t.Errorf("expected nil error for valid index, got %v", err)
	}
}
