package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// SliderOrientation defines the slider direction.
type SliderOrientation int

const (
	SliderHorizontal SliderOrientation = iota
	SliderVertical
)

// SliderNode is an interactive slider with a draggable thumb.
// Value ranges from min to max (default 0 to 1).
type SliderNode struct {
	PanelNode

	inputManager *input.InputManager
	value        float64
	min, max     float64

	trackColor   color.Color
	thumbColor   color.Color
	thumbWidth   float64
	orientation  SliderOrientation

	dragging bool
	OnChange func(value float64)
}

// NewSliderNode creates a new horizontal slider.
func NewSliderNode(name string, width, height float64, im *input.InputManager) *SliderNode {
	p := NewPanelNode(name, width, height)
	p.SetBackgroundColor(color.RGBA{40, 40, 40, 255})

	thumbW := height * 1.5
	if thumbW > width*0.2 {
		thumbW = width * 0.2
	}

	return &SliderNode{
		PanelNode:    *p,
		inputManager: im,
		value:        0.5,
		min:          0,
		max:          1,
		trackColor:   color.RGBA{60, 60, 60, 255},
		thumbColor:   color.RGBA{100, 140, 200, 255},
		thumbWidth:   thumbW,
		orientation:  SliderHorizontal,
	}
}

// SetRange sets the slider's value range. GetValue returns values in [min, max].
func (s *SliderNode) SetRange(min, max float64) {
	if min >= max {
		return
	}
	s.min, s.max = min, max
	// Clamp current value
	if s.value < s.min {
		s.value = s.min
	}
	if s.value > s.max {
		s.value = s.max
	}
}

// SetOrientation sets the slider direction (horizontal or vertical).
func (s *SliderNode) SetOrientation(o SliderOrientation) {
	s.orientation = o
}

// SetValue updates the slider value (clamped to min..max).
func (s *SliderNode) SetValue(val float64) {
	if val < s.min {
		val = s.min
	}
	if val > s.max {
		val = s.max
	}
	if s.value != val {
		s.value = val
		if s.OnChange != nil {
			s.OnChange(s.value)
		}
	}
}

// GetValue returns the current slider value in [min, max].
func (s *SliderNode) GetValue() float64 {
	return s.value
}

// norm returns value normalized to 0..1 for internal use.
func (s *SliderNode) norm() float64 {
	r := s.max - s.min
	if r <= 0 {
		return 0
	}
	return (s.value - s.min) / r
}

// setFromNorm sets value from normalized 0..1.
func (s *SliderNode) setFromNorm(t float64) {
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	s.SetValue(s.min + t*(s.max-s.min))
}

// Update handles mouse drag and click to update the value.
func (s *SliderNode) Update() {
	if s.inputManager == nil {
		return
	}

	mouseX := s.inputManager.GetCursorPos().X()
	mouseY := s.inputManager.GetCursorPos().Y()
	pos := s.GetWorldPosition()
	scale := s.GetWorldScale()
	trackW := s.width * scale.X()
	trackH := s.height * scale.Y()

	minX := pos.X()

	isOverTrack := mouseX >= pos.X() && mouseX <= pos.X()+trackW &&
		mouseY >= pos.Y() && mouseY <= pos.Y()+trackH

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	var t float64
	trackLen := trackW - s.thumbWidth*scale.X()
	if s.orientation == SliderVertical {
		trackLen = trackH - s.thumbWidth*scale.Y()
	}
	if trackLen <= 0 {
		trackLen = 1
	}

	if s.dragging {
		if pressed {
			if s.orientation == SliderHorizontal {
				t = (mouseX - minX) / trackLen
			} else {
				t = (mouseY - pos.Y()) / trackLen
			}
			s.setFromNorm(t)
		} else {
			s.dragging = false
		}
	} else if pressed && isOverTrack {
		s.dragging = true
		if s.orientation == SliderHorizontal {
			t = (mouseX - minX) / trackLen
		} else {
			t = (mouseY - pos.Y()) / trackLen
		}
		s.setFromNorm(t)
	}
}

// Draw renders the track and thumb.
func (s *SliderNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.PanelNode.Draw(target, op)

	worldPos := s.GetWorldPosition()
	scale := s.GetWorldScale()
	trackW := s.width * scale.X()
	trackH := s.height * scale.Y()
	var thumbW, thumbH float64
	if s.orientation == SliderHorizontal {
		thumbW = s.thumbWidth * scale.X()
		thumbH = trackH * 1.2
		if thumbH > trackH*2 {
			thumbH = trackH * 2
		}
	} else {
		thumbW = trackW * 1.2
		if thumbW > trackW*2 {
			thumbW = trackW * 2
		}
		thumbH = s.thumbWidth * scale.Y()
	}

	if s.trackColor != nil {
		if s.orientation == SliderHorizontal {
			cy := worldPos.Y() + (trackH-thumbH)/2 + thumbH/2
			vector.StrokeLine(target, float32(worldPos.X()), float32(cy), float32(worldPos.X()+trackW), float32(cy), 2, s.trackColor, true)
		} else {
			cx := worldPos.X() + trackW/2
			vector.StrokeLine(target, float32(cx), float32(worldPos.Y()), float32(cx), float32(worldPos.Y()+trackH), 2, s.trackColor, true)
		}
	}

	if s.thumbColor != nil {
		n := s.norm()
		if s.orientation == SliderHorizontal {
			thumbX := worldPos.X() + (trackW-thumbW)*n
			thumbY := worldPos.Y() + (trackH-thumbH)/2
			vector.DrawFilledRect(target, float32(thumbX), float32(thumbY), float32(thumbW), float32(thumbH), s.thumbColor, true)
		} else {
			thumbX := worldPos.X() + (trackW-thumbW)/2
			thumbY := worldPos.Y() + (trackH-thumbH)*n
			vector.DrawFilledRect(target, float32(thumbX), float32(thumbY), float32(thumbW), float32(thumbH), s.thumbColor, true)
		}
	}
}
