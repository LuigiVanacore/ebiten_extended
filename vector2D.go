package ebiten_extended

import (
	"fmt"
	"math"
)

// Struct for manage 2D Vector math
type Vector2D struct {
	X float64
	Y float64
}

func NewVector2D(x, y float64) Vector2D {
	return Vector2D{x, y}
}

func ZeroVector2d() Vector2D {
	return Vector2D{0, 0}
}

func UnitVector2D() Vector2D {
	return Vector2D{1, 1}
}

func (v *Vector2D) SetPosition(x, y float64) {
	v.X = x
	v.Y = y
}

func (v *Vector2D) Copy() Vector2D {
	return Vector2D{v.X, v.Y}
}

func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.DotProduct(*v))
}

func (v *Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v *Vector2D) Subtract(v2 Vector2D) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vector2D) Negate() Vector2D {
	return Vector2D{-v.X, -v.Y}
}

func (v *Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{v.X * v2.X, v.Y * v2.Y}
}

func (v *Vector2D) Divide(v2 Vector2D) Vector2D {
	return Vector2D{v.X / v2.X, v.Y / v2.Y}
}

func (v *Vector2D) DotProduct(v2 Vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2D) MultiplyScalar(s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

func (v Vector2D) DivideScalar(s float64) Vector2D {
	return Vector2D{v.X / s, v.Y / s}
}

func (v *Vector2D) RotateVector(degrees float64) Vector2D {
	var radian = DegreesToRadian(degrees)
	var sine = math.Sin(radian)
	var cosine = math.Cos(radian)

	return Vector2D{v.X*cosine - v.Y*sine, v.X*sine + v.Y*cosine}
}

func (v *Vector2D) RotateVector90() Vector2D {
	return Vector2D{-v.Y, v.X}
}

func (v *Vector2D) RotateVector180() Vector2D {
	return v.Negate()
}

func (v *Vector2D) RotateVector270() Vector2D {
	return Vector2D{v.Y, -v.X}
}

func (v *Vector2D) EncloseAngle(v2 Vector2D) float64 {
	var uA = UnitVector2D()
	var uB = UnitVector2D()
	var dp = uA.DotProduct(uB)
	return RadianToDegrees(math.Acos(dp))
}

func (v *Vector2D) ProjectVector(v2 Vector2D) Vector2D {
	var d = v2.DotProduct(v2)
	if d != 0 {
		return v2.MultiplyScalar(v.DotProduct(v2) / d)
	}
	return v2
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}
