package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/input"
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

	// Optional text label integrated within the button
	label *ebiten_extended.TextNode

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

// SetText initializes or updates the inside label of the button.
func (b *ButtonNode) SetText(textStr string, face text.Face, c color.Color) {
	if b.label == nil {
		b.label = ebiten_extended.NewTextNode(b.GetName()+"_label", textStr, face, c)
		// Embed the label visually in the center relative to the button
		// Assuming origin top-left, we place it with some padding for now
		b.label.SetPosition(10, b.height/2-10)
		b.AddChildren(b.label)
	} else {
		b.label.SetMessage(textStr)
		b.label.SetFont(face)
		b.label.SetColor(c)
	}
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
