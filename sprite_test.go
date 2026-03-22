package ludum

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestSpriteSetTextureRect(t *testing.T) {
	img := ebiten.NewImage(64, 64)
	defer img.Deallocate()

	sprite := NewSprite("test", img, 0, false)
	if sprite == nil {
		t.Fatal("NewSprite returned nil")
	}

	// Set a sub-rect
	sprite.SetTextureRect(0, 0, 32, 32)
	rect := sprite.GetTextureRect()
	if rect.GetSize().X() != 32 || rect.GetSize().Y() != 32 {
		t.Errorf("SetTextureRect: got size %v, want 32x32", rect.GetSize())
	}
}

func TestSpriteClone(t *testing.T) {
	img := ebiten.NewImage(32, 32)
	defer img.Deallocate()

	sprite := NewSprite("orig", img, 0, true)
	sprite.SetPosition(10, 20)
	sprite.SetVisible(true)

	clone := sprite.Clone()
	if clone == nil {
		t.Fatal("Clone returned nil")
	}
	if clone.GetTexture() != sprite.GetTexture() {
		t.Error("Clone should share texture")
	}
	pos := clone.GetPosition()
	if pos.X() != 10 || pos.Y() != 20 {
		t.Errorf("Clone position: got (%v, %v), want (10, 20)", pos.X(), pos.Y())
	}
	if !clone.GetVisible() {
		t.Error("Clone should be visible")
	}
}

func TestSpriteGetWorldBounds(t *testing.T) {
	img := ebiten.NewImage(100, 50)
	defer img.Deallocate()

	sprite := NewSprite("bounds", img, 0, true)
	sprite.SetPosition(200, 100)
	sprite.SetScale(1, 1)

	bounds := sprite.GetWorldBounds()
	// With pivot at center (50, 25), position at (200,100) means corners at roughly 150,75 to 250,125
	if bounds.GetSize().X() <= 0 || bounds.GetSize().Y() <= 0 {
		t.Errorf("GetWorldBounds: invalid size %v", bounds.GetSize())
	}
}

func TestSpriteFlip(t *testing.T) {
	img := ebiten.NewImage(16, 16)
	defer img.Deallocate()

	sprite := NewSprite("flip", img, 0, true)
	sprite.SetFlipX(true)
	sprite.SetFlipY(true)

	if !sprite.GetFlipX() || !sprite.GetFlipY() {
		t.Error("SetFlip did not set flip state")
	}

	sprite.SetFlip(false, false)
	if sprite.GetFlipX() || sprite.GetFlipY() {
		t.Error("SetFlip(false, false) did not clear flip")
	}
}

func TestSpriteVisible(t *testing.T) {
	img := ebiten.NewImage(8, 8)
	defer img.Deallocate()

	sprite := NewSprite("vis", img, 0, false)
	if !sprite.GetVisible() {
		t.Error("New sprite should be visible by default")
	}

	sprite.SetVisible(false)
	if sprite.GetVisible() {
		t.Error("SetVisible(false) did not take effect")
	}
}

func TestSpriteGetSize(t *testing.T) {
	img := ebiten.NewImage(64, 32)
	defer img.Deallocate()

	sprite := NewSprite("size", img, 0, false)
	sz := sprite.GetSize()
	if sz.X() != 64 || sz.Y() != 32 {
		t.Errorf("GetSize: got %v, want (64, 32)", sz)
	}
}
