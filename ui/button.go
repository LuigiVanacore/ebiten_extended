package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ButtonState represents the interactive status of the button.
type ButtonState int

const (
	ButtonStateIdle ButtonState = iota
	ButtonStateHover
	ButtonStatePressed
)

// ButtonNode is an interactive UI Panel that responds to mouse input (Hover and Click).
type ButtonNode struct {
	PanelNode

	inputManager *input.InputManager
	state        ButtonState
	focused      bool

	// Optional text label integrated within the button
	label *ludum.TextNode

	// Configurable stylings for different states
	IdleColor    color.Color
	HoverColor   color.Color
	PressedColor color.Color

	// Optional image styling for different states
	IdleImage    *ebiten.Image
	HoverImage   *ebiten.Image
	PressedImage *ebiten.Image

	// Callbacks
	OnClick        func()
	OnMouseEnter   func()
	OnMouseExit    func()
	OnMousePressed func()
}

// NewButtonNode creates a new interactive generic button.
func NewButtonNode(name string, width, height float64, im *input.InputManager) *ButtonNode {
	p := NewPanelNode(name, width, height)
	btn := &ButtonNode{
		PanelNode:    *p,
		inputManager: im,
		state:        ButtonStateIdle,
		IdleColor:    color.RGBA{100, 100, 100, 255}, // Default gray
		HoverColor:   color.RGBA{130, 130, 130, 255}, // Light gray
		PressedColor: color.RGBA{70, 70, 70, 255},    // Dark gray
	}
	btn.SetBackgroundColor(btn.IdleColor)
	return btn
}

// SetSize updates the button dimensions and re-centers the label if present.
func (b *ButtonNode) SetSize(w, h float64) {
	b.PanelNode.SetSize(w, h)
	b.centerLabel()
}

// SetText initializes or updates the inside label of the button, centered horizontally and vertically.
func (b *ButtonNode) SetText(textStr string, face text.Face, c color.Color) {
	if b.label == nil {
		b.label = ludum.NewTextNode(b.GetName()+"_label", textStr, face, c)
		b.AddChildren(b.label)
	} else {
		b.label.SetMessage(textStr)
		b.label.SetFont(face)
		b.label.SetColor(c)
	}
	// Center the label in the button
	b.centerLabel()
}

// centerLabel positions the label in the center of the button.
func (b *ButtonNode) centerLabel() {
	if b.label == nil || b.label.GetFont() == nil {
		return
	}
	w, h := text.Measure(b.label.GetMessage(), b.label.GetFont(), 0)
	padX := (b.width - w) / 2
	padY := (b.height - h) / 2
	b.label.SetPosition(padX, padY)
}

// Update reads input to process hover and click events.
func (b *ButtonNode) Update() {
	if b.inputManager == nil {
		return
	}

	mouseX := b.inputManager.GetCursorPos().X()
	mouseY := b.inputManager.GetCursorPos().Y()

	isHovering := b.ContainsPoint(mouseX, mouseY)
	isClicking := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Determine new state based on intersection
	newState := ButtonStateIdle
	if isHovering {
		if isClicking {
			newState = ButtonStatePressed
		} else {
			newState = ButtonStateHover
		}
	}

	// State transitioning logic & Callbacks
	if b.state != newState {
		// Exit Events
		if b.state == ButtonStateHover && newState == ButtonStateIdle {
			if b.OnMouseExit != nil {
				b.OnMouseExit()
			}
		}

		// Enter Events
		if b.state == ButtonStateIdle && newState == ButtonStateHover {
			if b.OnMouseEnter != nil {
				b.OnMouseEnter()
			}
		} else if b.state == ButtonStateHover && newState == ButtonStatePressed {
			if b.OnMousePressed != nil {
				b.OnMousePressed()
			}
		} else if b.state == ButtonStatePressed && newState == ButtonStateHover {
			// Button released inside the boundings -> Registered as a full "Click"
			if b.OnClick != nil {
				b.OnClick()
			}
		}

		b.state = newState
		b.updateVisuals()
	}
}

// SetFocused implements Focusable.
func (b *ButtonNode) SetFocused(focused bool) {
	if b.focused == focused {
		return
	}
	b.focused = focused
	if focused {
		b.state = ButtonStateHover
	} else {
		b.state = ButtonStateIdle
	}
	b.updateVisuals()
}

// IsFocused implements Focusable.
func (b *ButtonNode) IsFocused() bool {
	return b.focused
}

// OnFocusAction implements Focusable.
func (b *ButtonNode) OnFocusAction() {
	// Temporarily simulate press visually
	b.state = ButtonStatePressed
	b.updateVisuals()
	if b.OnClick != nil {
		b.OnClick()
	}
	// Return to hover (focused) visualization
	b.state = ButtonStateHover
	b.updateVisuals()
}

// updateVisuals changes the background color or image responding to the current state.
func (b *ButtonNode) updateVisuals() {
	switch b.state {
	case ButtonStateIdle:
		b.SetBackgroundColor(b.IdleColor)
		b.SetBackgroundImage(b.IdleImage)
	case ButtonStateHover:
		b.SetBackgroundColor(b.HoverColor)
		b.SetBackgroundImage(b.HoverImage)
	case ButtonStatePressed:
		b.SetBackgroundColor(b.PressedColor)
		b.SetBackgroundImage(b.PressedImage)
	}
}
