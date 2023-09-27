package math2D

import "math"

//function and constat used for Math 2D

const FLOAT_PI float64 = 3.14159265358979323846

func Min(x float64, y float64) float64 {
	if x == 0 || x == y {
		return x
	}
	if x < y {
		return x
	}
	return y
}

func Max(x float64, y float64) float64 {
	if x == 0 || x == y {
		return y
	}
	if x > y {
		return x
	}
	return y
}

func Distance(x, y, x2, y2 float64) float64 {
	dx := x - x2
	dy := y - y2
	ds := (dx * dx) + (dy * dy)
	return math.Sqrt(math.Abs(ds))
}

func DegreesToRadian(degree float64) float64 {
	return degree * FLOAT_PI / 180.0
}

func RadianToDegrees(radian float64) float64 {
	return radian * 180.0 / FLOAT_PI
}
