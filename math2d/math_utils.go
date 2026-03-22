package math2D

import "math"

// FLOAT_PI is the constant π for angle conversions.
const FLOAT_PI float64 = math.Pi

// Min returns the smaller of x and y.
func Min(x, y float64) float64 {
	if x <= y {
		return x
	}
	return y
}

// Max returns the larger of x and y.
func Max(x, y float64) float64 {
	if x >= y {
		return x
	}
	return y
}

// Distance returns the Euclidean distance between (x, y) and (x2, y2).
func Distance(x, y, x2, y2 float64) float64 {
	dx := x - x2
	dy := y - y2
	ds := (dx * dx) + (dy * dy)
	return math.Sqrt(math.Abs(ds))
}

// DegreesToRadian converts an angle from degrees to radians.
func DegreesToRadian(degree float64) float64 {
	return degree * FLOAT_PI / 180.0
}

// RadianToDegrees converts an angle from radians to degrees.
func RadianToDegrees(radian float64) float64 {
	return radian * 180.0 / FLOAT_PI
}
