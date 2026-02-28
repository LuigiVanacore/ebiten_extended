package ebiten_extended

import (
	"image/color"
	"math"

	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Node2D
	width     uint
	height    uint
	zoom 	float64
	surface   *ebiten.Image
}


func NewCamera(w, h uint) *Camera {
	c := &Camera{
		width:  w,
		height: h,
		zoom: 1.0,
	}
	c.surface = ebiten.NewImage(int(w), int(h))
	return c
}

func (c *Camera) GetZoom() float64 {
	return c.zoom
}


func (c *Camera) SetZoom(zoom float64) *Camera {
	c.zoom = zoom
	if c.zoom <= 0.01 {
		c.zoom = 0.01
	}
	c.Resize(c.width, c.height)
	return c
}

func (c *Camera) GetSurface() *ebiten.Image {
	return c.surface
}

func (c *Camera) Resize(w, h uint) *Camera {
	c.width = w
	c.height = h
	newW := int(float64(w) * 1.0 / c.zoom)
	newH := int(float64(h) * 1.0 / c.zoom)

	if newW <= 16384 && newH <= 16384 {
		c.surface.Deallocate()
		c.surface = ebiten.NewImage(newW, newH)
	}
	return c
}


func (c *Camera) Deallocate() {
	c.surface.Deallocate()
}

func (c *Camera) Fill(color color.Color) {
	c.surface.Fill(color)
}



// ApplyRelativeTranslation applies camera-relative translation to op in place
// (center of surface then offset by camera position). Use this to avoid allocations.
func (c *Camera) ApplyRelativeTranslation(op *ebiten.DrawImageOptions, x, y float64) {
	size := c.surface.Bounds().Size()
	op.GeoM.Translate(float64(size.X)/2, float64(size.Y)/2)
	op.GeoM.Translate(-c.GetPosition().X(), -c.GetPosition().Y())
}

// GetRelativeTranslation applies camera-relative translation to op in place and returns op.
func (c *Camera) GetRelativeTranslation(op *ebiten.DrawImageOptions, x, y float64) *ebiten.DrawImageOptions {
	c.ApplyRelativeTranslation(op, x, y)
	return op
}

func (c *Camera) GetRelativeRotation(ops *ebiten.DrawImageOptions, rotation, originX, originY float64) *ebiten.DrawImageOptions {
	ops.GeoM.Translate(originX, originY)
	ops.GeoM.Rotate(rotation)
	ops.GeoM.Translate(-originX, -originY)
	return ops
}

func (c *Camera) GetRelativeScale(ops *ebiten.DrawImageOptions, scaleX, scaleY float64) *ebiten.DrawImageOptions {
	ops.GeoM.Scale(scaleX, scaleY)
	return ops
}

func (c *Camera) GetRelativeSkew(ops *ebiten.DrawImageOptions, skewX, skewY float64) *ebiten.DrawImageOptions {
	ops.GeoM.Skew(skewX, skewY)
	return ops
}

func (c * Camera) DrawImage(image *ebiten.Image, ops *ebiten.DrawImageOptions) {
	c.surface.DrawImage(image, ops)
}

func (c *Camera) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	size := c.surface.Bounds().Size()
	cx := float64(size.X) / 2.0
	cy := float64(size.Y) / 2.0

	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Scale(c.zoom, c.zoom)
	op.GeoM.Rotate(float64(c.GetRotation()))
	op.GeoM.Translate(cx*c.zoom, cy*c.zoom)

	screen.DrawImage(c.surface, op)
}

func (c *Camera) GetScreenCoords(x, y float64) (float64, float64) {
	w, h := c.width, c.height
	co := math.Cos(float64(c.GetRotation()))
	si := math.Sin(float64(c.GetRotation()))

	x, y = x-c.GetPosition().X(), y-c.GetPosition().Y()
	x, y = co*x-si*y, si*x+co*y

	return x*c.zoom + float64(w)/2, y*c.zoom + float64(h)/2
}

func (c *Camera) GetWorldCoords(x, y float64) (float64, float64) {
	w, h := c.width, c.height
	co := math.Cos(-float64(c.GetRotation()))
	si := math.Sin(-float64(c.GetRotation()))

	x, y = (x-float64(w)/2)/c.zoom, (y-float64(h)/2)/c.zoom
	x, y = co*x-si*y, si*x+co*y

	return x + c.GetPosition().X(), y + c.GetPosition().Y()
}

func (c *Camera) GetCursorCoords(inputMgr *input.InputManager) (float64, float64) {
	cursor_position := inputMgr.GetCursorPos()
	return c.GetWorldCoords(cursor_position.X(), cursor_position.Y())
}