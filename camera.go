package ludum

import (
	"image/color"
	"math"

	"github.com/LuigiVanacore/ludum/input"
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

// Camera represents a 2D camera that provides the viewpoint for the scene.
// It inherits from Node2D and renders everything within its view onto a surface.
type Camera struct {
	Node2D
	width           uint
	height          uint
	zoom            float64
	surface         *ebiten.Image
	nodeToFollow    transform.Transformable
	FollowSmoothing float64

	shakeIntensity float64
	shakeDuration  float64
	shakeRemaining float64
	shakeOffsetX   float64
	shakeOffsetY   float64
	shakePhase     float64

	// bounds constrain camera position when set (boundsSet true). World rect: minX, minY, maxX, maxY.
	boundsSet                                      bool
	boundsMinX, boundsMinY, boundsMaxX, boundsMaxY float64
}

// NewCamera creates and initializes a new Camera with the specified width and height.
func NewCamera(w, h uint) *Camera {
	c := &Camera{
		width:  w,
		height: h,
		zoom:   1.0,
	}
	c.surface = ebiten.NewImage(int(w), int(h))
	return c
}

// GetZoom returns the current zoom level of the camera.
func (c *Camera) GetZoom() float64 {
	return c.zoom
}

// SetZoom sets the zoom level of the camera (clamped to a minimum of 0.01)
// and dynamically resizes its rendering surface to accommodate the zoom.
func (c *Camera) SetZoom(zoom float64) *Camera {
	c.zoom = zoom
	if c.zoom <= 0.01 {
		c.zoom = 0.01
	}
	c.Resize(c.width, c.height)
	return c
}

// GetSurface returns the image surface onto which this camera captures the scene.
func (c *Camera) GetSurface() *ebiten.Image {
	return c.surface
}

// Resize changes the base dimensions of the camera and reallocates its
// surface memory if the new dimensions, considering zoom, don't exceed max limits.
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

// Deallocate releases the resources associated with the camera's surface.
func (c *Camera) Deallocate() {
	c.surface.Deallocate()
}

// Fill clears the camera's surface with a predefined background color.
func (c *Camera) Fill(color color.Color) {
	c.surface.Fill(color)
}

// ApplyRelativeTranslation applies camera-relative translation to op in place.
// World (0,0) maps to the top-left of the camera surface; camera position acts as view offset.
// Shake offset is applied when Shake is active.
func (c *Camera) ApplyRelativeTranslation(op *ebiten.DrawImageOptions, x, y float64) {
	px := c.GetPosition().X() + c.shakeOffsetX
	py := c.GetPosition().Y() + c.shakeOffsetY
	op.GeoM.Translate(x-px, y-py)
}

// GetRelativeTranslation applies camera-relative translation to op in place and returns op.
func (c *Camera) GetRelativeTranslation(op *ebiten.DrawImageOptions, x, y float64) *ebiten.DrawImageOptions {
	c.ApplyRelativeTranslation(op, x, y)
	return op
}

// GetRelativeRotation calculates and applies camera-relative rotation around a specific origin point, returning op.
func (c *Camera) GetRelativeRotation(ops *ebiten.DrawImageOptions, rotation, originX, originY float64) *ebiten.DrawImageOptions {
	ops.GeoM.Translate(originX, originY)
	ops.GeoM.Rotate(rotation)
	ops.GeoM.Translate(-originX, -originY)
	return ops
}

// GetRelativeScale applies camera-relative scaling directly to the ops matrix.
func (c *Camera) GetRelativeScale(ops *ebiten.DrawImageOptions, scaleX, scaleY float64) *ebiten.DrawImageOptions {
	ops.GeoM.Scale(scaleX, scaleY)
	return ops
}

// GetRelativeSkew applies camera-relative skewing directly to the ops matrix.
func (c *Camera) GetRelativeSkew(ops *ebiten.DrawImageOptions, skewX, skewY float64) *ebiten.DrawImageOptions {
	ops.GeoM.Skew(skewX, skewY)
	return ops
}

// DrawImage handles drawing a given image directly onto the camera's surface using provided options.
func (c *Camera) DrawImage(image *ebiten.Image, ops *ebiten.DrawImageOptions) {
	c.surface.DrawImage(image, ops)
}

// Draw transfers the camera's rendered scene surface onto the final screen, applying overall rotation and zoom.
func (c *Camera) Draw(screen *ebiten.Image) {
	c.DrawWithOptions(screen, &ebiten.DrawImageOptions{})
}

// DrawWithOptions transfers the camera surface onto screen using base options, then camera transform.
func (c *Camera) DrawWithOptions(screen *ebiten.Image, baseOp *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	if baseOp != nil {
		op.GeoM = baseOp.GeoM
		op.ColorScale = baseOp.ColorScale
		op.CompositeMode = baseOp.CompositeMode
		op.Filter = baseOp.Filter
		op.DisableMipmaps = baseOp.DisableMipmaps
		op.Blend = baseOp.Blend
	}
	size := c.surface.Bounds().Size()
	cx := float64(size.X) / 2.0
	cy := float64(size.Y) / 2.0

	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Scale(c.zoom, c.zoom)
	op.GeoM.Rotate(float64(c.GetRotation()))
	op.GeoM.Translate(cx*c.zoom, cy*c.zoom)

	screen.DrawImage(c.surface, op)
}

// SetFollow sets a target transformable to follow. Pass nil to disable following.
func (c *Camera) SetFollow(node transform.Transformable) {
	c.nodeToFollow = node
}

// SetBounds constrains the camera position to a world rectangle (minX, minY, maxX, maxY).
// Pass minX >= maxX to disable bounds.
func (c *Camera) SetBounds(minX, minY, maxX, maxY float64) {
	if minX >= maxX || minY >= maxY {
		c.boundsSet = false
		return
	}
	c.boundsSet = true
	c.boundsMinX, c.boundsMinY = minX, minY
	c.boundsMaxX, c.boundsMaxY = maxX, maxY
}

// GetBounds returns the current bounds and whether they are active.
func (c *Camera) GetBounds() (minX, minY, maxX, maxY float64, active bool) {
	return c.boundsMinX, c.boundsMinY, c.boundsMaxX, c.boundsMaxY, c.boundsSet
}

// GetVisibleWorldRect returns the axis-aligned world rectangle currently visible (for culling).
func (c *Camera) GetVisibleWorldRect() math2d.Rectangle {
	px := c.GetPosition().X()
	py := c.GetPosition().Y()
	w := float64(c.width) / c.zoom
	h := float64(c.height) / c.zoom
	return math2d.NewRectangle(
		math2d.NewVector2D(px, py),
		math2d.NewVector2D(w, h),
	)
}

// applyBounds clamps the camera position to bounds if set.
func (c *Camera) applyBounds() {
	if !c.boundsSet {
		return
	}
	viewW := float64(c.width) / c.zoom
	viewH := float64(c.height) / c.zoom
	// Camera position is top-left of view; clamp so view stays within bounds
	minCamX := c.boundsMinX
	minCamY := c.boundsMinY
	maxCamX := c.boundsMaxX - viewW
	maxCamY := c.boundsMaxY - viewH
	if maxCamX < minCamX {
		maxCamX = minCamX
	}
	if maxCamY < minCamY {
		maxCamY = minCamY
	}
	px, py := c.GetPosition().X(), c.GetPosition().Y()
	if px < minCamX {
		px = minCamX
	} else if px > maxCamX {
		px = maxCamX
	}
	if py < minCamY {
		py = minCamY
	} else if py > maxCamY {
		py = maxCamY
	}
	c.SetPosition(px, py)
}

// Shake adds a screen shake effect. intensity is the max pixel offset; duration is in seconds.
// Call repeatedly to extend or intensify the shake.
func (c *Camera) Shake(intensity, duration float64) {
	if intensity > c.shakeIntensity || duration > c.shakeRemaining {
		c.shakeIntensity = intensity
		c.shakeDuration = duration
		if c.shakeRemaining < duration {
			c.shakeRemaining = duration
		}
	}
}

// Update syncs camera position with the followed node (if any) and decays shake.
// Uses fixed 60 FPS timestep (1/60 s) for shake decay.
func (c *Camera) Update() {
	const tick = 1.0 / 60.0
	if c.shakeRemaining > 0 {
		c.shakeRemaining -= tick
		if c.shakeRemaining < 0 {
			c.shakeRemaining = 0
			c.shakeOffsetX, c.shakeOffsetY = 0, 0
		} else {
			frac := c.shakeRemaining / c.shakeDuration
			c.shakePhase += tick * 60
			c.shakeOffsetX = c.shakeIntensity * frac * math.Sin(c.shakePhase*1.7)
			c.shakeOffsetY = c.shakeIntensity * frac * math.Sin(c.shakePhase*2.3+1)
		}
	}
	if c.nodeToFollow != nil {
		worldTransform := c.nodeToFollow.GetWorldTransform()
		target := (&worldTransform).GetPosition()

		if c.FollowSmoothing <= 0 {
			c.SetPosition(target.X(), target.Y())
		} else {
			t := c.FollowSmoothing
			if t > 1 {
				t = 1
			}
			current := c.GetPosition()
			x := current.X() + (target.X()-current.X())*t
			y := current.Y() + (target.Y()-current.Y())*t
			c.SetPosition(x, y)
		}
	}
	c.applyBounds()
}

// GetScreenCoords converts world coordinates (x, y) into camera surface coordinates (top-left origin).
func (c *Camera) GetScreenCoords(x, y float64) (float64, float64) {
	co := math.Cos(float64(c.GetRotation()))
	si := math.Sin(float64(c.GetRotation()))

	px := c.GetPosition().X() + c.shakeOffsetX
	py := c.GetPosition().Y() + c.shakeOffsetY
	x, y = x-px, y-py
	x, y = co*x-si*y, si*x+co*y

	return x * c.zoom, y * c.zoom
}

// GetWorldCoords converts screen coordinates (e.g., from a mouse, top-left origin) into world coordinates.
func (c *Camera) GetWorldCoords(x, y float64) (float64, float64) {
	co := math.Cos(-float64(c.GetRotation()))
	si := math.Sin(-float64(c.GetRotation()))

	x, y = x/c.zoom, y/c.zoom
	x, y = co*x-si*y, si*x+co*y

	px := c.GetPosition().X() + c.shakeOffsetX
	py := c.GetPosition().Y() + c.shakeOffsetY
	return x + px, y + py
}

// GetCursorCoords utilizes the provided InputManager to retrieve mouse cursor pos naturally mapped into world coordinates.
func (c *Camera) GetCursorCoords(inputMgr *input.InputManager) (float64, float64) {
	cursor_position := inputMgr.GetCursorPos()
	return c.GetWorldCoords(cursor_position.X(), cursor_position.Y())
}
