// Package math2d provides 2D geometry and vector math for games and collision detection.
//
// Types:
//   - Vector2D: 2D vector with X, Y; operations include Normalize, DotProduct, RotateVector, ProjectVector.
//   - Circle: center and radius.
//   - Rectangle: position (top-left) and size (width, height); GetCenter for center point.
//   - Segment: line segment from start to end point; ProjectSegment for projection onto an axis.
//   - Line: infinite line with base point and direction vector.
//   - Range: 1D interval (min, max); OverlappingRanges, SortRange, RangeHull.
//   - OrientedRectangle: rectangle with center, half-extents, and rotation (for SAT collision).
//
// Utilities (mathUtils): Min, Max, Distance, DegreesToRadian, RadianToDegrees, FLOAT_PI.
package math2d
