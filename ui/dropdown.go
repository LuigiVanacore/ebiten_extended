package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// DropdownNode is a select dropdown with a trigger button and expandable list.
// SetItems to populate options; GetSelectedIndex/GetSelectedText for current selection.
type DropdownNode struct {
	PanelNode

	InputManager  *input.InputManager
	items         []string
	selectedIndex int
	open          bool

	itemHeight     float64
	font           text.Face
	textColor      color.Color
	itemBgColor    color.Color
	itemHoverColor color.Color

	OnSelectionChanged func(index int, text string)
}

// NewDropdownNode creates a dropdown with the given trigger size.
func NewDropdownNode(name string, width, height float64, font text.Face, im *input.InputManager) *DropdownNode {
	p := NewPanelNode(name, width, height)
	p.SetBackgroundColor(color.RGBA{80, 80, 80, 255})
	return &DropdownNode{
		PanelNode:      *p,
		InputManager:   im,
		selectedIndex:  -1,
		itemHeight:     24,
		font:           font,
		textColor:      color.White,
		itemBgColor:    color.RGBA{60, 60, 60, 255},
		itemHoverColor: color.RGBA{90, 90, 90, 255},
	}
}

// SetItems sets the dropdown options. Clears selection if index is invalid.
func (d *DropdownNode) SetItems(items []string) {
	d.items = items
	if d.selectedIndex >= len(d.items) {
		d.selectedIndex = -1
	}
}

// GetItems returns the current options.
func (d *DropdownNode) GetItems() []string {
	return d.items
}

// SetSelectedIndex sets the selection (-1 for none).
func (d *DropdownNode) SetSelectedIndex(i int) {
	if i < -1 || i >= len(d.items) {
		return
	}
	d.selectedIndex = i
}

// GetSelectedIndex returns the selected index or -1.
func (d *DropdownNode) GetSelectedIndex() int {
	return d.selectedIndex
}

// GetSelectedText returns the selected item text or empty string.
func (d *DropdownNode) GetSelectedText() string {
	if d.selectedIndex < 0 || d.selectedIndex >= len(d.items) {
		return ""
	}
	return d.items[d.selectedIndex]
}

// SetItemHeight sets the height of each list item.
func (d *DropdownNode) SetItemHeight(h float64) {
	d.itemHeight = h
}

// SetItemColors sets background and hover colors for list items.
func (d *DropdownNode) SetItemColors(bg, hover color.Color) {
	d.itemBgColor = bg
	d.itemHoverColor = hover
}

// Update handles click to toggle/open and select.
func (d *DropdownNode) Update() {
	if d.InputManager == nil {
		return
	}
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}
	cx := d.InputManager.GetCursorPos().X()
	cy := d.InputManager.GetCursorPos().Y()
	pos := d.GetWorldPosition()
	scale := d.GetWorldScale()
	sw, sh := scale.X(), scale.Y()

	triggerMinX := pos.X()
	triggerMaxX := pos.X() + d.width*sw
	triggerMinY := pos.Y()
	triggerMaxY := pos.Y() + d.height*sh

	inTrigger := cx >= triggerMinX && cx <= triggerMaxX && cy >= triggerMinY && cy <= triggerMaxY

	if d.open {
		listH := float64(len(d.items)) * d.itemHeight * sh
		listMinX := triggerMinX
		listMaxX := triggerMaxX
		listMinY := triggerMaxY
		listMaxY := triggerMaxY + listH
		inList := cx >= listMinX && cx <= listMaxX && cy >= listMinY && cy <= listMaxY

		if inList {
			idx := int((cy - listMinY) / (d.itemHeight * sh))
			if idx >= 0 && idx < len(d.items) {
				d.selectedIndex = idx
				d.open = false
				if d.OnSelectionChanged != nil {
					d.OnSelectionChanged(idx, d.items[idx])
				}
			}
		} else if !inTrigger {
			d.open = false
		} else {
			d.open = false
		}
	} else {
		if inTrigger {
			d.open = true
		}
	}
}

// Draw renders the trigger and optional list.
func (d *DropdownNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	d.PanelNode.Draw(target, op)
	pos := d.GetWorldPosition()
	scale := d.GetWorldScale()
	padX, padY := 6.0*scale.X(), 4.0*scale.Y()

	display := d.GetSelectedText()
	if display == "" && len(d.items) > 0 {
		display = "Select..."
	}
	if d.font != nil && display != "" {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(pos.X()+padX, pos.Y()+padY)
		opts.ColorScale.ScaleWithColor(d.textColor)
		text.Draw(target, display, d.font, opts)
	}

	if !d.open {
		return
	}

	listY := pos.Y() + d.height*scale.Y()
	cx, cy := 0.0, 0.0
	if d.InputManager != nil {
		cx = d.InputManager.GetCursorPos().X()
		cy = d.InputManager.GetCursorPos().Y()
	}
	for i, item := range d.items {
		itemY := listY + float64(i)*d.itemHeight*scale.Y()
		itemH := d.itemHeight * scale.Y()
		itemColor := d.itemBgColor
		if cx >= pos.X() && cx <= pos.X()+d.width*scale.X() &&
			cy >= itemY && cy <= itemY+itemH {
			itemColor = d.itemHoverColor
		}
		vector.DrawFilledRect(target,
			float32(pos.X()), float32(itemY),
			float32(d.width*scale.X()), float32(itemH),
			itemColor, true)
		if d.font != nil {
			opts := &text.DrawOptions{}
			opts.GeoM.Translate(pos.X()+padX, itemY+padY)
			opts.ColorScale.ScaleWithColor(d.textColor)
			text.Draw(target, item, d.font, opts)
		}
	}
}

// Ensure DropdownNode is an Updatable so World.Update calls it.
var _ ludum.Updatable = (*DropdownNode)(nil)
