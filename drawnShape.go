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
	layer 		   int
}


func NewDrawnCircle(name string, position math2D.Vector2D, radius float32, color color.Color, isAntialiasActive bool, layer int) *DrawnShape {
	circle := &DrawnShape{
		Node2D: *NewNode2D(name),
		shapeType: CiRCLE,
		radius: radius,
		color: color,
		isAntialiasActive: isAntialiasActive,
		layer: layer,
	}
	circle.SetPosition(position)
	return circle
}

func NewDrawnRectangle(name string, position math2D.Vector2D, size math2D.Vector2D, color color.Color, isAntialiasActive bool, layer int) *DrawnShape {
	rect := &DrawnShape{
		Node2D: *NewNode2D(name),
		shapeType: RECT,
		size: size,
		color: color,
		isAntialiasActive: isAntialiasActive,
		layer: layer,
	}
	rect.SetPosition(position)
	return rect
}

func (d *DrawnShape) GetLayer() int {
	return d.layer
}


func (d *DrawnShape) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	switch d.shapeType {
		case CiRCLE:
			vector.DrawFilledCircle(target, float32(op.GeoM.Element(0, 2)), 
											 float32(op.GeoM.Element(1, 2)),
											 d.radius, 
											 d.color, 
											 d.isAntialiasActive)
		
		case RECT:
			vector.DrawFilledRect(target, float32(op.GeoM.Element(0, 2) - d.GetPosition().X() - d.size.X()/2), 
											 float32(op.GeoM.Element(1, 2) - d.GetPosition().Y() - d.size.Y()/2),
				float32(d.size.X()),
				float32(d.size.Y()),
				d.color,
				d.isAntialiasActive)
	}
}
