package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ludum/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// CheckboxNode is a ButtonNode that maintains a Checked boolean state.
// It draws a graphical tick/indicator when turned on.
type CheckboxNode struct {
	ButtonNode
	Checked   bool
	TickColor color.Color
	TickImage *ebiten.Image
	OnToggle  func(bool) // Event fired when checkbox is toggled
}

// NewCheckboxNode creates a new geometric Checkbox.
func NewCheckboxNode(name string, size float64, im *input.InputManager) *CheckboxNode {
	// A checkbox is usually square
	btn := NewButtonNode(name, size, size, im)

	c := &CheckboxNode{
		ButtonNode: *btn,
		Checked:    false,
		TickColor:  color.RGBA{255, 255, 255, 255}, // White tick block
	}

	// Override the Button's OnClick to support the toggle mechanic
	c.ButtonNode.OnClick = func() {
		c.Checked = !c.Checked
		if c.OnToggle != nil {
			c.OnToggle(c.Checked)
		}
	}

	return c
}

// Draw renders the checkbox button surface, followed by an interior check/tick block if true.
// By default it draws a centered square as the "Tick", but an Image overrides this easily.
func (c *CheckboxNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	c.ButtonNode.Draw(target, op) // Draw base interactive button background

	if c.Checked {
		worldPos := c.GetWorldPosition()
		scale := c.GetWorldScale()

		if c.TickImage != nil {
			// Draw Tick Image
			bounds := c.TickImage.Bounds()
			imgW, imgH := float64(bounds.Dx()), float64(bounds.Dy())
			if imgW > 0 && imgH > 0 {
				drawOp := &ebiten.DrawImageOptions{}
				scaleX := (c.width * scale.X()) / imgW
				scaleY := (c.height * scale.Y()) / imgH
				drawOp.GeoM.Scale(scaleX, scaleY)
				drawOp.GeoM.Translate(worldPos.X(), worldPos.Y())

				if op != nil {
					drawOp.ColorScale = op.ColorScale
				}

				target.DrawImage(c.TickImage, drawOp)
			}
		} else if c.TickColor != nil {
			// Draw generic vector Tick (A smaller rect centered inside the box)
			paddingRatio := 0.25 // leaving 25% margin
			tickX := worldPos.X() + (c.width * scale.X() * paddingRatio)
			tickY := worldPos.Y() + (c.height * scale.Y() * paddingRatio)
			tickW := (c.width * scale.X() * (1 - paddingRatio*2))
			tickH := (c.height * scale.Y() * (1 - paddingRatio*2))

			vector.DrawFilledRect(
				target,
				float32(tickX),
				float32(tickY),
				float32(tickW),
				float32(tickH),
				c.TickColor,
				true,
			)
		}
	}
}
