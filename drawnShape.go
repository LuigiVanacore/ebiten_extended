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

type DrawnShape struct {
	Node2D
	shapeType         ShapeType
	size              math2D.Vector2D
	radius            float32
	color 			  color.Color
	isAntialiasActive bool
}


func NewDrawnCircle(name string, position math2D.Vector2D, radius float32, color color.Color, isAntialiasActive bool) *DrawnShape {
	circle := &DrawnShape{
		Node2D: *NewNode2D(name),
		shapeType: CiRCLE,
		radius: radius,
		color: color,
		isAntialiasActive: isAntialiasActive,
	}
	circle.SetPosition(position)
	return circle
}


func (d *DrawnShape) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	switch d.shapeType {
		case CiRCLE:
			vector.DrawFilledCircle(target, float32(d.GetPosition().X()), 
											float32(d.GetPosition().Y()),
											 d.radius, 
											 d.color, 
											 d.isAntialiasActive)
	}
}
