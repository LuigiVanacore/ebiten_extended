package ludum

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// TextNode represents a visual 2D scene graph node dedicated to drawing scalable geometry-based text phrases.
type TextNode struct {
	Node2D
	message     string
	color       color.Color
	font        text.Face
	layer       int
	maxWidth    float64 // if > 0, wrap text to fit; 0 = no wrap
	drawOpts    text.DrawOptions
	cachedLines []string // lines after word wrap, invalidated when message/maxWidth/font change
}

// NewTextNode instantiates a display entity resolving specific text output using the assigned typography face format.
func NewTextNode(name string, message string, font text.Face, c color.Color) *TextNode {
	label := &TextNode{message: message, Node2D: *NewNode2D(name), font: font, color: c}
	return label
}

// SetMessage dynamically overrides the actively drawn text string maintained by this graph element.
func (l *TextNode) SetMessage(message string) {
	if l.message != message {
		l.message = message
		l.cachedLines = nil
	}
}

// SetMaxWidth enables word wrap when > 0. Text is wrapped to fit within the given width in pixels.
// Set to 0 to disable wrapping.
func (l *TextNode) SetMaxWidth(w float64) {
	if l.maxWidth != w {
		l.maxWidth = w
		l.cachedLines = nil
	}
}

// GetMaxWidth returns the current max width for wrapping (0 = disabled).
func (l *TextNode) GetMaxWidth() float64 {
	return l.maxWidth
}

// GetMessage returns the current text.
func (l *TextNode) GetMessage() string {
	return l.message
}

// SetFont sets the font face used for drawing.
func (l *TextNode) SetFont(face text.Face) {
	l.font = face
	l.cachedLines = nil
}

// GetFont returns the font face used for drawing.
func (l *TextNode) GetFont() text.Face {
	return l.font
}

// SetColor sets the text color.
func (l *TextNode) SetColor(c color.Color) {
	l.color = c
}

// GetColor returns the text color.
func (l *TextNode) GetColor() color.Color {
	return l.color
}

// GetLayer returns the render layer of this text node.
func (l *TextNode) GetLayer() int {
	return l.layer
}

// SetLayer sets the render layer of this text node.
func (l *TextNode) SetLayer(layer int) {
	l.layer = layer
}

// GetWidth returns the width of the rendered text. Implements ui.SizeProvider for layout.
// Returns the maximum line width when maxWidth is set, or the measured width of the longest line.
func (l *TextNode) GetWidth() float64 {
	lines := l.getLines()
	if len(lines) == 0 || l.font == nil {
		return 0
	}
	var maxW float64
	for _, ln := range lines {
		w, _ := text.Measure(ln, l.font, 0)
		if w > maxW {
			maxW = w
		}
	}
	if l.maxWidth > 0 && maxW > l.maxWidth {
		return l.maxWidth
	}
	return maxW
}

// GetHeight returns the height of the rendered text. Implements ui.SizeProvider for layout.
// Returns lineCount * lineHeight.
func (l *TextNode) GetHeight() float64 {
	lines := l.getLines()
	if len(lines) == 0 || l.font == nil {
		return 0
	}
	_, lineHeight := text.Measure("M", l.font, 0)
	return float64(len(lines)) * lineHeight
}

func (l *TextNode) getLines() []string {
	if l.font == nil {
		return nil
	}
	if l.cachedLines != nil {
		return l.cachedLines
	}
	if l.maxWidth <= 0 {
		if l.message == "" {
			return nil
		}
		l.cachedLines = []string{l.message}
		return l.cachedLines
	}
	l.cachedLines = l.wrapText(l.message, l.font, l.maxWidth)
	return l.cachedLines
}

func (l *TextNode) wrapText(msg string, face text.Face, maxW float64) []string {
	if msg == "" {
		return nil
	}
	words := strings.Fields(msg)
	if len(words) == 0 {
		return []string{msg}
	}
	var lines []string
	var line strings.Builder
	line.WriteString(words[0])
	lineW, _ := text.Measure(words[0], face, 0)

	for i := 1; i < len(words); i++ {
		w := words[i]
		addW, _ := text.Measure(" "+w, face, 0)
		if lineW+addW <= maxW {
			line.WriteString(" ")
			line.WriteString(w)
			lineW += addW
		} else {
			lines = append(lines, line.String())
			line.Reset()
			line.WriteString(w)
			lineW, _ = text.Measure(w, face, 0)
		}
	}
	if line.Len() > 0 {
		lines = append(lines, line.String())
	}
	return lines
}

// Draw translates the text formatting onto the assigned canvas plane mapping it through inherent node translation attributes.
func (l *TextNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	lines := l.getLines()
	if len(lines) == 0 || l.font == nil {
		return
	}

	l.drawOpts = text.DrawOptions{}
	if op != nil {
		l.drawOpts.GeoM = op.GeoM
	} else {
		worldPos := l.GetWorldPosition()
		l.drawOpts.GeoM.Translate(worldPos.X(), worldPos.Y())
	}
	l.drawOpts.ColorScale.ScaleWithColor(l.color)

	_, lineHeight := text.Measure("M", l.font, 0)
	y := 0.0
	for _, ln := range lines {
		lineOpts := l.drawOpts
		lineOpts.GeoM.Translate(0, y)
		text.Draw(target, ln, l.font, &lineOpts)
		y += lineHeight
	}
}
