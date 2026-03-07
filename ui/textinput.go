package ui

import (
	"image/color"
	"unicode/utf8"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// TextInputNode is an editable text field. Click to focus, type to edit.
// Supports UTF-8 input, backspace, cursor, selection, and Ctrl+C/V/X copy/paste/cut.
type TextInputNode struct {
	PanelNode
	InputManager *input.InputManager
	text         string
	placeholder  string
	font         text.Face
	textColor    color.Color
	placeholderColor color.Color
	focused      bool
	cursorBlink  float64
	maxLength    int // 0 = unlimited
	cursorIndex  int // rune index; 0 = before first char
	selStart     int
	selEnd       int
	clipboard    string // internal buffer for copy/paste

	OnSubmit func(text string)
	OnChange func(text string)
}

// NewTextInputNode creates an editable text field.
func NewTextInputNode(name string, width, height float64, font text.Face, im *input.InputManager) *TextInputNode {
	p := NewPanelNode(name, width, height)
	p.SetBackgroundColor(color.RGBA{40, 40, 40, 255})
	return &TextInputNode{
		PanelNode:        *p,
		InputManager:     im,
		font:             font,
		textColor:        color.White,
		placeholderColor: color.RGBA{120, 120, 120, 255},
	}
}

// SetText sets the current text and resets cursor/selection to end.
func (t *TextInputNode) SetText(s string) {
	if t.text != s {
		t.text = s
		n := utf8.RuneCountInString(s)
		t.cursorIndex = n
		t.selStart, t.selEnd = n, n
		if t.OnChange != nil {
			t.OnChange(s)
		}
	}
}

// GetText returns the current text.
func (t *TextInputNode) GetText() string {
	return t.text
}

// SetPlaceholder sets the placeholder when text is empty.
func (t *TextInputNode) SetPlaceholder(s string) {
	t.placeholder = s
}

// SetMaxLength limits the number of runes (0 = unlimited).
func (t *TextInputNode) SetMaxLength(n int) {
	t.maxLength = n
}

// SetTextColor sets the color for typed text.
func (t *TextInputNode) SetTextColor(c color.Color) {
	t.textColor = c
}

// IsFocused returns whether the field has keyboard focus.
func (t *TextInputNode) IsFocused() bool {
	return t.focused
}

// Focus sets focus so the field receives keyboard input.
func (t *TextInputNode) Focus() {
	t.focused = true
	t.cursorBlink = 0
}

// Blur removes focus.
func (t *TextInputNode) Blur() {
	t.focused = false
}

// Update handles focus (click) and keyboard input.
func (t *TextInputNode) Update() {
	if t.InputManager == nil {
		return
	}
	cx := t.InputManager.GetCursorPos().X()
	cy := t.InputManager.GetCursorPos().Y()
	pos := t.GetWorldPosition()
	scale := t.GetWorldScale()
	minX, minY := pos.X(), pos.Y()
	maxX, maxY := pos.X()+t.width*scale.X(), pos.Y()+t.height*scale.Y()
	inBounds := cx >= minX && cx <= maxX && cy >= minY && cy <= maxY
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if inBounds {
			t.Focus()
			t.cursorFromX(cx - minX - 4*scale.X())
		} else {
			t.Blur()
		}
	}
	if !t.focused {
		return
	}
	t.cursorBlink += ebiten_extended.FIXED_DELTA
	if t.cursorBlink > 1 {
		t.cursorBlink -= 1
	}

	ctrlOrCmd := ebiten.IsKeyPressed(ebiten.KeyControl) || ebiten.IsKeyPressed(ebiten.KeyMetaLeft) || ebiten.IsKeyPressed(ebiten.KeyMetaRight)

	if ctrlOrCmd && inpututil.IsKeyJustPressed(ebiten.KeyC) {
		t.copySelection()
	}
	if ctrlOrCmd && inpututil.IsKeyJustPressed(ebiten.KeyX) {
		t.cutSelection()
	}
	if ctrlOrCmd && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		t.paste()
	}
	if ctrlOrCmd && inpututil.IsKeyJustPressed(ebiten.KeyA) {
		t.selStart, t.selEnd = 0, utf8.RuneCountInString(t.text)
	}

	shift := ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight)
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		t.moveCursor(-1, shift)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		t.moveCursor(1, shift)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyHome) {
		t.moveCursorTo(0, shift)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnd) {
		t.moveCursorTo(utf8.RuneCountInString(t.text), shift)
	}

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		t.deleteBackward()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDelete) {
		t.deleteForward()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
		if t.OnSubmit != nil {
			t.OnSubmit(t.text)
		}
	}
	for _, r := range ebiten.AppendInputChars(nil) {
		t.insertRune(r)
	}
}

func (t *TextInputNode) cursorFromX(offsetX float64) {
	if t.font == nil || offsetX <= 0 {
		t.cursorIndex = 0
		return
	}
	runes := []rune(t.text)
	best := 0
	bestDist := offsetX
	for i := 0; i <= len(runes); i++ {
		prefix := string(runes[:i])
		w, _ := text.Measure(prefix, t.font, 0)
		dist := offsetX - w
		if dist < 0 {
			dist = -dist
		}
		if dist < bestDist {
			bestDist = dist
			best = i
		}
	}
	t.cursorIndex = best
	t.selStart, t.selEnd = best, best
}

func (t *TextInputNode) moveCursor(delta int, extend bool) {
	n := utf8.RuneCountInString(t.text)
	newIdx := t.cursorIndex + delta
	if newIdx < 0 {
		newIdx = 0
	}
	if newIdx > n {
		newIdx = n
	}
	if extend {
		if delta > 0 {
			t.selStart = min(t.selStart, t.selEnd)
			t.selEnd = newIdx
		} else {
			t.selEnd = max(t.selStart, t.selEnd)
			t.selStart = newIdx
		}
	} else {
		t.selStart, t.selEnd = newIdx, newIdx
	}
	t.cursorIndex = newIdx
}

func (t *TextInputNode) moveCursorTo(idx int, extend bool) {
	n := utf8.RuneCountInString(t.text)
	if idx > n {
		idx = n
	}
	if idx < 0 {
		idx = 0
	}
	if extend {
		if idx >= max(t.selStart, t.selEnd) {
			t.selStart = min(t.selStart, t.selEnd)
			t.selEnd = idx
		} else {
			t.selEnd = max(t.selStart, t.selEnd)
			t.selStart = idx
		}
	} else {
		t.selStart, t.selEnd = idx, idx
	}
	t.cursorIndex = idx
}

func (t *TextInputNode) selectionBounds() (start, end int) {
	if t.selStart < t.selEnd {
		return t.selStart, t.selEnd
	}
	return t.selEnd, t.selStart
}

func (t *TextInputNode) hasSelection() bool {
	return t.selStart != t.selEnd
}

func (t *TextInputNode) getSelectedText() string {
	s, e := t.selectionBounds()
	if s >= e {
		return ""
	}
	runes := []rune(t.text)
	return string(runes[s:e])
}

func (t *TextInputNode) copySelection() {
	if t.hasSelection() {
		t.clipboard = t.getSelectedText()
	} else {
		t.clipboard = t.text
	}
}

func (t *TextInputNode) cutSelection() {
	t.copySelection()
	t.deleteSelection()
}

func (t *TextInputNode) paste() {
	if t.clipboard == "" {
		return
	}
	if t.hasSelection() {
		t.deleteSelection()
	}
	runes := []rune(t.text)
	before := string(runes[:t.cursorIndex])
	after := string(runes[t.cursorIndex:])
	newLen := utf8.RuneCountInString(before) + utf8.RuneCountInString(t.clipboard) + utf8.RuneCountInString(after)
	if t.maxLength > 0 && newLen > t.maxLength {
		room := t.maxLength - (utf8.RuneCountInString(before) + utf8.RuneCountInString(after))
		if room <= 0 {
			return
		}
		clipRunes := []rune(t.clipboard)
		if len(clipRunes) > room {
			t.clipboard = string(clipRunes[:room])
		}
	}
	t.text = before + t.clipboard + after
	t.cursorIndex = utf8.RuneCountInString(before) + utf8.RuneCountInString(t.clipboard)
	t.selStart, t.selEnd = t.cursorIndex, t.cursorIndex
	if t.OnChange != nil {
		t.OnChange(t.text)
	}
}

func (t *TextInputNode) deleteSelection() {
	s, e := t.selectionBounds()
	if s >= e {
		return
	}
	runes := []rune(t.text)
	t.text = string(runes[:s]) + string(runes[e:])
	t.cursorIndex = s
	t.selStart, t.selEnd = s, s
	if t.OnChange != nil {
		t.OnChange(t.text)
	}
}

func (t *TextInputNode) deleteBackward() {
	if t.hasSelection() {
		t.deleteSelection()
		return
	}
	if t.cursorIndex <= 0 {
		return
	}
	runes := []rune(t.text)
	prefix := string(runes[:t.cursorIndex-1])
	suffix := string(runes[t.cursorIndex:])
	t.text = prefix + suffix
	t.cursorIndex--
	t.selStart, t.selEnd = t.cursorIndex, t.cursorIndex
	if t.OnChange != nil {
		t.OnChange(t.text)
	}
}

func (t *TextInputNode) deleteForward() {
	if t.hasSelection() {
		t.deleteSelection()
		return
	}
	runes := []rune(t.text)
	if t.cursorIndex >= len(runes) {
		return
	}
	prefix := string(runes[:t.cursorIndex])
	suffix := string(runes[t.cursorIndex+1:])
	t.text = prefix + suffix
	t.selStart, t.selEnd = t.cursorIndex, t.cursorIndex
	if t.OnChange != nil {
		t.OnChange(t.text)
	}
}

func (t *TextInputNode) insertRune(r rune) {
	if t.hasSelection() {
		t.deleteSelection()
	}
	if t.maxLength > 0 && utf8.RuneCountInString(t.text) >= t.maxLength {
		return
	}
	runes := []rune(t.text)
	before := string(runes[:t.cursorIndex])
	after := string(runes[t.cursorIndex:])
	t.text = before + string(r) + after
	t.cursorIndex++
	t.selStart, t.selEnd = t.cursorIndex, t.cursorIndex
	if t.OnChange != nil {
		t.OnChange(t.text)
	}
}

// Draw renders the field, text, selection highlight, and cursor.
func (t *TextInputNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	t.PanelNode.Draw(target, op)
	pos := t.GetWorldPosition()
	scale := t.GetWorldScale()
	padX, padY := 4.0*scale.X(), 4.0*scale.Y()
	drawX, drawY := pos.X()+padX, pos.Y()+padY
	displayText := t.text
	displayColor := t.textColor
	if displayText == "" && t.placeholder != "" {
		displayText = t.placeholder
		displayColor = t.placeholderColor
	}
	if t.font != nil && displayText != "" {
		textOpts := &text.DrawOptions{}
		textOpts.GeoM.Translate(drawX, drawY)
		textOpts.ColorScale.ScaleWithColor(displayColor)
		text.Draw(target, displayText, t.font, textOpts)
	}
	if t.focused && t.font != nil {
		runes := []rune(t.text)
		ci := t.cursorIndex
		if ci > len(runes) {
			ci = len(runes)
		}
		if ci < 0 {
			ci = 0
		}
		cursorX := drawX
		if ci > 0 {
			prefix := string(runes[:ci])
			w, _ := text.Measure(prefix, t.font, 0)
			cursorX = drawX + w
		}
		selStart, selEnd := t.selectionBounds()
		if selStart < selEnd && displayText == t.text {
			startPrefix := string(runes[:selStart])
			endPrefix := string(runes[:selEnd])
			w0, _ := text.Measure(startPrefix, t.font, 0)
			w1, _ := text.Measure(endPrefix, t.font, 0)
			h, _ := text.Measure("M", t.font, 0)
			selColor := color.RGBA{80, 120, 200, 120}
			vector.DrawFilledRect(target, float32(drawX+w0), float32(drawY), float32(w1-w0), float32(h), selColor, true)
		}
		if int(t.cursorBlink*2)%2 == 0 {
			h, _ := text.Measure("M", t.font, 0)
			vector.StrokeLine(target, float32(cursorX), float32(drawY), float32(cursorX), float32(drawY+h), 2, t.textColor, true)
		}
	}
}
