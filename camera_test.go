package ludum

import (
	"math"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestCameraGetScreenCoords(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(0, 0)
	cam.SetZoom(1.0)

	sx, sy := cam.GetScreenCoords(10, 20)
	if !floatEqual(sx, 10, 0.01) || !floatEqual(sy, 20, 0.01) {
		t.Errorf("GetScreenCoords(10,20) with cam at (0,0): expected (10,20), got (%.2f,%.2f)", sx, sy)
	}

	cam.SetPosition(100, 50)
	sx, sy = cam.GetScreenCoords(100, 50)
	if !floatEqual(sx, 0, 0.01) || !floatEqual(sy, 0, 0.01) {
		t.Errorf("GetScreenCoords(100,50) with cam at (100,50): expected (0,0), got (%.2f,%.2f)", sx, sy)
	}
}

func TestCameraGetWorldCoords(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(0, 0)

	wx, wy := cam.GetWorldCoords(10, 20)
	if !floatEqual(wx, 10, 0.01) || !floatEqual(wy, 20, 0.01) {
		t.Errorf("GetWorldCoords(10,20) with cam at (0,0): expected (10,20), got (%.2f,%.2f)", wx, wy)
	}

	cam.SetPosition(100, 50)
	wx, wy = cam.GetWorldCoords(0, 0)
	if !floatEqual(wx, 100, 0.01) || !floatEqual(wy, 50, 0.01) {
		t.Errorf("GetWorldCoords(0,0) with cam at (100,50): expected (100,50), got (%.2f,%.2f)", wx, wy)
	}
}

func TestCameraApplyRelativeTranslation(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(50, 25)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 80) // object at world (100, 80)
	cam.ApplyRelativeTranslation(op, 0, 0)

	// After ApplyRelativeTranslation: Translate(-50, -25)
	// So final translation = (100-50, 80-25) = (50, 55)
	ex, ey := op.GeoM.Element(0, 2), op.GeoM.Element(1, 2)
	if !floatEqual(ex, 50, 0.01) || !floatEqual(ey, 55, 0.01) {
		t.Errorf("ApplyRelativeTranslation: expected (50, 55), got (%.2f, %.2f)", ex, ey)
	}
}

func TestCameraApplyRelativeTranslation_WithOffsets(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(50, 25)

	op := &ebiten.DrawImageOptions{}
	cam.ApplyRelativeTranslation(op, 100, 80)

	ex, ey := op.GeoM.Element(0, 2), op.GeoM.Element(1, 2)
	if !floatEqual(ex, 50, 0.01) || !floatEqual(ey, 55, 0.01) {
		t.Errorf("ApplyRelativeTranslation with offsets: expected (50, 55), got (%.2f, %.2f)", ex, ey)
	}
}

func TestCameraFollow(t *testing.T) {
	cam := NewCamera(320, 240)
	target := NewNode2D("target")
	target.SetPosition(30, 40)

	cam.SetFollow(target)
	cam.Update()

	pos := cam.GetPosition()
	if !floatEqual(pos.X(), 30, 0.01) || !floatEqual(pos.Y(), 40, 0.01) {
		t.Errorf("camera follow expected (30,40), got (%.2f,%.2f)", pos.X(), pos.Y())
	}
}

func TestCameraZoomAffectsCoords(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(0, 0)
	cam.SetZoom(2.0)

	// World (10, 20) with zoom 2 -> screen (20, 40)
	sx, sy := cam.GetScreenCoords(10, 20)
	if !floatEqual(sx, 20, 0.01) || !floatEqual(sy, 40, 0.01) {
		t.Errorf("GetScreenCoords with zoom 2: expected (20,40), got (%.2f,%.2f)", sx, sy)
	}

	// Screen (20, 40) with zoom 2 -> world (10, 20)
	wx, wy := cam.GetWorldCoords(20, 40)
	if !floatEqual(wx, 10, 0.01) || !floatEqual(wy, 20, 0.01) {
		t.Errorf("GetWorldCoords with zoom 2: expected (10,20), got (%.2f,%.2f)", wx, wy)
	}
}

func TestCameraRoundtrip(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(75, 33)
	cam.SetZoom(1.5)

	wx, wy := 100.0, 200.0
	sx, sy := cam.GetScreenCoords(wx, wy)
	rx, ry := cam.GetWorldCoords(sx, sy)

	if !floatEqual(rx, wx, 0.01) || !floatEqual(ry, wy, 0.01) {
		t.Errorf("Roundtrip: world (%.2f,%.2f) -> screen (%.2f,%.2f) -> world (%.2f,%.2f)",
			wx, wy, sx, sy, rx, ry)
	}
}

func TestCameraRotationRoundtrip(t *testing.T) {
	cam := NewCamera(320, 240)
	cam.SetPosition(0, 0)
	cam.SetRotation(math.Pi / 4)

	wx, wy := 10.0, 0.0
	sx, sy := cam.GetScreenCoords(wx, wy)
	rx, ry := cam.GetWorldCoords(sx, sy)

	if !floatEqual(rx, wx, 0.1) || !floatEqual(ry, wy, 0.1) {
		t.Errorf("Rotation roundtrip: expected (%.2f,%.2f), got (%.2f,%.2f)", wx, wy, rx, ry)
	}
}
