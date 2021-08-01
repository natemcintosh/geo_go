// Package gogeo provides simple Point and Line Segment types. It also provides
// functionality for rotating Points and Line Segments, and checking if Line Segments
// intersect.
package gogeo

import "math"

// Point is a point in 2D space. It can also be thought of as a vector from the origin
// to the point.
type Point struct {
	X float64
	Y float64
}

// Angle is the angle of a Point in radians from the positive x-axis.
func (p Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

// Plus adds two points, interpreting the points as vectors.
func (p Point) Plus(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Minus subtracts two points, interpreting the points as vectors.
func (p Point) Minus(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// LineSegment is a line segment in 2D space. It is defined by two Points.
type LineSegment struct {
	P1 Point
	P2 Point
}
