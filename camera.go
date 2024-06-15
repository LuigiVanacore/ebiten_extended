package ebiten_extended

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	transform.Transform
	width     uint
	height    uint
	zoom 	float64
	surface   *ebiten.Image
}

// func (c *Camera) SetZoom(zoom float64) *Camera {
// 	c.Scale = zoom
// 	if c.Scale <= 0.01 {
// 		c.Scale = 0.01
// 	}
// 	return c
// }



// // GetScreenCoords converts world coords into screen coords
// func (c *Camera) GetScreenCoords(x, y float64) (float64, float64) {
// 	w, h := c.Width, c.Height

// 	x, y = x-c.X, y-c.Y
// 	x, y = x*c.Scale, y*c.Scale

// 	// Translate to screen center
// 	return x + float64(w)/2, y + float64(h)/2
// }

// // GetWorldCoords converts screen coords into world coords
// func (c *Camera) GetWorldCoords(x, y float64) (float64, float64) {
// 	w, h := c.Width, c.Height

// 	x, y = x-float64(w)/2, y-float64(h)/2

// 	// Translate the coordinates
// 	x += float64(w) / 2
// 	y += float64(h) / 2

// 	return x, y
// }

// // Center returns the center point of the camera, based on its Width and Height.
// func (c *Camera) Center() (float64, float64) {
// 	return float64(c.Width) * 0.5, float64(c.Height) * 0.5
// }

// // GetCursorCoords converts cursor/screen coords into world coords
// func (c *Camera) GetCursorCoords() (float64, float64) {
// 	cx, cy := ebiten.CursorPosition()
// 	return c.GetWorldCoords(float64(cx), float64(cy))
// }

// // WorldMatrix modifies the `ops` parameter to be world relative.
// func (c *Camera) WorldMatrix(ops *ebiten.DrawImageOptions) {
// 	centerX, centerY := c.Center()
// 	ops.GeoM.Scale(c.Scale, c.Scale)
// 	ops.GeoM.Translate(centerX, centerY)
// 	ops.GeoM.Translate(-c.X*c.Scale, -c.Y*c.Scale)
// }

// // SetPosition looks at a position
// func (c *Camera) SetPosition(x, y float64) *Camera {
// 	c.X = x
// 	c.Y = y
// 	return c
// }




 // Zoom *= the current zoom
 func (c *Camera) Zoom(mul float64) *Camera {
 	c.zoom *= mul
 	if c.zoom <= 0.01 {
 		c.zoom = 0.01
 	}
 	c.Resize(c.width, c.height)
 	return c
 }

// // SetZoom sets the zoom
// func (c *Camera) SetZoom(zoom float64) *Camera {
// 	c.Scale = zoom
// 	if c.Scale <= 0.01 {
// 		c.Scale = 0.01
// 	}
// 	c.Resize(c.Width, c.Height)
// 	return c
// }

 // Resize resizes the camera Surface
 func (c *Camera) Resize(w, h uint) *Camera {
 	c.width = w
 	c.height = h
 	newW := int(float64(w) * 1.0 / c.zoom)
 	newH := int(float64(h) * 1.0 / c.zoom)
 	if newW <= 16384 && newH <= 16384 {
 		c.surface.Dispose()
 		c.surface = ebiten.NewImage(newW, newH)
 	}
 	return c
 }

// // GetTranslation alters the provided *ebiten.DrawImageOptions' translation based on the given x,y offset and the
// // camera's position
// func (c *Camera) GetTranslation(ops *ebiten.DrawImageOptions, x, y float64) *ebiten.DrawImageOptions {
// 	w, h := c.Surface.Size()
// 	ops.GeoM.Translate(float64(w)/2, float64(h)/2)
// 	ops.GeoM.Translate(-c.X+x, -c.Y+y)
// 	return ops
// }

// // GetRotation alters the provided *ebiten.DrawImageOptions' rotation using the provided originX and originY args
// func (c *Camera) GetRotation(ops *ebiten.DrawImageOptions, rot, originX, originY float64) *ebiten.DrawImageOptions {
// 	ops.GeoM.Translate(originX, originY)
// 	ops.GeoM.Rotate(rot)
// 	ops.GeoM.Translate(-originX, -originY)
// 	return ops
// }

// // GetScale alters the provided *ebiten.DrawImageOptions' scale
// func (c *Camera) GetScale(ops *ebiten.DrawImageOptions, scaleX, scaleY float64) *ebiten.DrawImageOptions {
// 	ops.GeoM.Scale(scaleX, scaleY)
// 	return ops
// }

// // GetSkew alters the provided *ebiten.DrawImageOptions' skew
// func (c *Camera) GetSkew(ops *ebiten.DrawImageOptions, skewX, skewY float64) *ebiten.DrawImageOptions {
// 	ops.GeoM.Skew(skewX, skewY)
// 	return ops
// }

// draws the camera's surface to the screen and applies zoom
 func (c *Camera) Draw(screen *ebiten.Image) {
 	op := &ebiten.DrawImageOptions{}
 	w := c.surface.Bounds().Size().X
	h := c.surface.Bounds().Size().Y
 	cx := float64(w) / 2.0
 	cy := float64(h) / 2.0

 	op.GeoM.Translate(-cx, -cy)
 	op.GeoM.Scale(c.zoom, c.zoom)
 	op.GeoM.Rotate(float64(c.GetRotation()))
 	op.GeoM.Translate(cx*c.zoom, cy*c.zoom)

 	screen.DrawImage(c.surface, op)
 }

 // GetScreenCoords converts world coords into screen coords
 func (c *Camera) GetScreenCoords(x, y float64) (float64, float64) {
 	w, h := c.width, c.height
 	co := math.Cos(float64(c.GetRotation()))
 	si := math.Sin(float64(c.GetRotation()))
	camera_position := c.GetPosition()
	camera_x := camera_position.X()
	camera_y := camera_position.Y()
 	x, y = x-camera_x, y-camera_y
 	x, y = co*x-si*y, si*x+co*y

	return x*c.zoom + float64(w)/2, y*c.zoom + float64(h)/2
 }

 // GetWorldCoords converts screen coords into world coords
 func (c *Camera) GetWorldCoords(x, y float64) (float64, float64) {
 	w, h := c.width, c.height
 	cos := math.Cos(-float64(c.GetRotation()))
 	sin := math.Sin(-float64(c.GetRotation()))

 	x, y = (x-float64(w)/2)/c.zoom, (y-float64(h)/2)/c.zoom
 	x, y = cos*x-sin*y, sin*x+cos*y
	camera_position := c.GetPosition()
 	return x + camera_position.X(), y + camera_position.Y()
 }

 // GetCursorCoords converts cursor/screen coords into world coords
 func (c *Camera) GetCursorCoords() (float64, float64) {
 	cx, cy := ebiten.CursorPosition()
 	return c.GetWorldCoords(float64(cx), float64(cy))
 }