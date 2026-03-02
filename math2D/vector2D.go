package math2D

import (
	"fmt"
	"math"
)

// Vector2D represents a 2D vector with X and Y components.
// Use NewVector2D, ZeroVector2D, or OneVector2D to construct.
type Vector2D struct {
	x float64
	y float64
}

// NewVector2D returns a vector with the given x and y components.
func NewVector2D(x, y float64) Vector2D {
	return Vector2D{x, y}
}

// X returns the x component of the vector.
func (v Vector2D) X() float64 {
	return v.x
}

// Y returns the y component of the vector.
func (v Vector2D) Y() float64 {
	return v.y
}

// SetX sets the x component.
func (v *Vector2D) SetX(x float64) {
	v.x = x
}

// SetY sets the y component.
func (v *Vector2D) SetY(y float64) {
	v.y = y
}

// ZeroVector2D returns the zero vector (0, 0).
func ZeroVector2D() Vector2D {
	return Vector2D{0, 0}
}

 
// Normalize returns a unit-length vector in the same direction; returns v unchanged if length is zero.
func (v Vector2D) Normalize() Vector2D {
	length := v.Length()
	if length != 0 {
		return v.DivideScalar(length)
	}
	return v
}

// OneVector2D returns the vector (1, 1).
func OneVector2D() Vector2D {
	return Vector2D{1, 1}
}

// Magnitude returns the length of the vector (same as Length).
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vector2D) SetPosition(x, y float64) {
	v.x = x
	v.y = y
}

func (v *Vector2D) Translate(x, y float64) {
	v.x += x
	v.y += y
}

func (v Vector2D) Clone() Vector2D {
	return Vector2D{v.x, v.y}
}

func (v Vector2D) Length() float64 {
	return math.Sqrt(DotProduct(v,v))
}

func AddVectors(v1, v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

func SubtractVectors(v1, v2 Vector2D) Vector2D{
	return Vector2D { x: v1.x - v2.x, y: v1.y - v2.y }
}

func (v *Vector2D) SetToZero() {
	v.x = 0
	v.y = 0
}

func (v Vector2D) IsZero() bool {
	return v.x == 0 && v.y == 0
}

func (v Vector2D) Negate() Vector2D {
	return Vector2D{-v.x, -v.y}
}

func (v *Vector2D) MultiplyVector(v1 Vector2D) Vector2D {
	return Vector2D{v1.x * v.x, v1.y * v.y}
}

func (v *Vector2D) DivideVectors(v1 Vector2D) Vector2D {
	return Vector2D{v1.x / v.x, v1.y / v.y}
}

// DotProduct returns the dot product of two vectors.
func DotProduct(v1, v2 Vector2D) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

func (v *Vector2D) MultiplyScalar( s  float64) Vector2D {
	return Vector2D{v.x * s, v.y * s}
}

func(v *Vector2D)   DivideScalar( s float64) Vector2D {
	return Vector2D{v.x / s, v.y / s}
}

func (v Vector2D) RotateVector(degrees float64) Vector2D {
	var radian = DegreesToRadian(degrees)
	var sine = math.Sin(radian)
	var cosine = math.Cos(radian)

	return Vector2D{v.x*cosine - v.y*sine, v.x*sine + v.y*cosine}
}

func (v Vector2D) RotateVector90() Vector2D {
	return Vector2D{-v.y, v.x}
}

func (v Vector2D) RotateVector180() Vector2D {
	return v.Negate()
}

func (v Vector2D) RotateVector270() Vector2D {
	return Vector2D{v.y, -v.x}
}

func (v Vector2D) EncloseAngle(v2 Vector2D) float64 {
	var uA = OneVector2D()
	var uB = OneVector2D()
	var dp = DotProduct(uA, uB)
	return DegreesToRadian(math.Acos(dp))
}

func (v Vector2D) ProjectVector(v2 Vector2D) Vector2D {
	var d = DotProduct(v2, v2)
	if d != 0 {
		return v2.MultiplyScalar( DotProduct(v, v2) / d)
	}
	return v2
}


func (v Vector2D) IsParallel(v2 Vector2D) bool {
	na := v.RotateVector90()
	return !(0 == v.x && 0 == v.y ) && !(0 == v2.x && 0 == v2.y) && 0 == DotProduct(na, v2)
}

func (v Vector2D) IsEqual(v2 Vector2D) bool {	
	return v.x == v2.x && v.y == v2.y
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%v:%v", v.x, v.y)
}
