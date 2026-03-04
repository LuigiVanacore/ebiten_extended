package ebiten_extended

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ShapeType int

const (
	RECT ShapeType = iota
	CiRCLE
	LINE
)

const CIRCLE ShapeType = CiRCLE

type DrawnShape struct {
	Node2D
	shapeType         ShapeType
	size              math2D.Vector2D
	radius            float32
	color             color.Color
	lineFrom          math2D.Vector2D
	lineTo            math2D.Vector2D
	isAntialiasActive bool
	layer             int
}


func NewDrawnCircle(name string, position math2D.Vector2D, radius float32, color color.Color, isAntialiasActive bool, layer int) *DrawnShape {
	circle := &DrawnShape{
		Node2D:            *NewNode2D(name),
		shapeType:         CiRCLE,
		radius:            radius,
		color:             color,
		isAntialiasActive: isAntialiasActive,
		layer:             layer,
	}
	circle.SetPosition(position.X(), position.Y())
	return circle
}

func NewDrawnRectangle(name string, position math2D.Vector2D, size math2D.Vector2D, color color.Color, isAntialiasActive bool, layer int) *DrawnShape {
	rect := &DrawnShape{
		Node2D:            *NewNode2D(name),
		shapeType:         RECT,
		size:              size,
		color:             color,
		isAntialiasActive: isAntialiasActive,
		layer:             layer,
	}
	rect.SetPosition(position.X(), position.Y())
	return rect
}

// NewDrawnLine creates a line shape from lineFrom to lineTo, relative to the node position.
func NewDrawnLine(name string, position, lineFrom, lineTo math2D.Vector2D, color color.Color, isAntialiasActive bool, layer int) *DrawnShape {
	line := &DrawnShape{
		Node2D:            *NewNode2D(name),
		shapeType:         LINE,
		color:             color,
		lineFrom:          lineFrom,
		lineTo:            lineTo,
		isAntialiasActive: isAntialiasActive,
		layer:             layer,
	}
	line.SetPosition(position.X(), position.Y())
	return line
}

func (d *DrawnShape) GetLayer() int {
	return d.layer
}


func (d *DrawnShape) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	x := float32(op.GeoM.Element(0, 2))
	y := float32(op.GeoM.Element(1, 2))
	switch d.shapeType {
	case CiRCLE:
		vector.DrawFilledCircle(target, x, y, d.radius, d.color, d.isAntialiasActive)
	case RECT:
		vector.DrawFilledRect(
			target,
			x-float32(d.size.X())/2,
			y-float32(d.size.Y())/2,
			float32(d.size.X()),
			float32(d.size.Y()),
			d.color,
			d.isAntialiasActive,
		)
	case LINE:
		vector.StrokeLine(
			target,
			x+float32(d.lineFrom.X()),
			y+float32(d.lineFrom.Y()),
			x+float32(d.lineTo.X()),
			y+float32(d.lineTo.Y()),
			1,
			d.color,
			d.isAntialiasActive,
		)
	}
}
