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

// Equals tests if two Points are the same
func (p Point) Equals(q Point) bool {
	return (p.X == q.X) && (p.Y == q.Y)
}

const float64EqualityThreshold = 1e-9

func (p Point) AlmostEquals(q Point) bool {
	x_absolute_diff := math.Abs(p.X - q.X)
	y_absolute_diff := math.Abs(p.Y - q.Y)
	return (x_absolute_diff < float64EqualityThreshold) && (y_absolute_diff < float64EqualityThreshold)
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

// Times multiplies a Point by a scalar `f`.
func (p Point) Times(f float64) Point {
	return Point{p.X * f, p.Y * f}
}

// Divide divides a Point by a scalar `f`.
func (p Point) Divide(f float64) Point {
	return Point{p.X / f, p.Y / f}
}

// Rotate rotates a LineSegment by the given angle in radians.
func (p Point) Rotate(angle float64) Point {
	s := math.Sin(angle)
	c := math.Cos(angle)
	return Point{
		X: c*p.X - s*p.Y,
		Y: s*p.X + c*p.Y,
	}
}

// LineSegment is a line segment in 2D space. It is defined by two Points.
type LineSegment struct {
	P1 Point
	P2 Point
}

// Plus adds the x and y components of a Point to a LineSegment.
func (l LineSegment) Plus(p Point) LineSegment {
	return LineSegment{l.P1.Plus(p), l.P2.Plus(p)}
}

// Minus subtracts the x and y components of a Point to a LineSegment.
func (l LineSegment) Minus(p Point) LineSegment {
	return LineSegment{l.P1.Minus(p), l.P2.Minus(p)}
}

// Equals tests if two LineSegments are equal.
func (l LineSegment) Equals(m LineSegment) bool {
	return (l.P1.X == m.P1.X) && (l.P1.Y == m.P1.Y) && (l.P2.X == m.P2.X) && (l.P2.Y == m.P2.Y)
}

// Angle calculates the angle of a LineSegment in radians from where it intersects the positive x-axis.
func (l LineSegment) Angle() float64 {
	return math.Atan2(l.P2.Y-l.P1.Y, l.P2.X-l.P1.X)
}

// RotateAboutOrigin rotates a LineSegment by the given angle in radians about the origin.
func (l LineSegment) RotateAboutOrigin(angle float64) LineSegment {
	return LineSegment{l.P1.Rotate(angle), l.P2.Rotate(angle)}
}

// OpenInterval represents the open interval [a, b].
type OpenInterval struct {
	Lower float64
	Upper float64
}

func (o OpenInterval) Equals(p OpenInterval) bool {
	// Check if the lower bound is NaN on both, or are equal.
	if (math.IsNaN(o.Lower) && math.IsNaN(p.Lower)) || (o.Lower == p.Lower) {
		// Check if the upper bound is NaN on both, or are equal.
		if (math.IsNaN(o.Upper) && math.IsNaN(p.Upper)) || (o.Upper == p.Upper) {
			return true
		}
	}
	return false
}

// Intersection calculates the overlap of two OpenIntervals. If there is no overlap, it
// returns an OpenInterval with NaN values
func (o OpenInterval) Intersection(p OpenInterval) OpenInterval {
	if (o.Upper < p.Lower) || (p.Upper < o.Lower) {
		return OpenInterval{math.NaN(), math.NaN()}
	}
	q_start := math.Max(o.Lower, p.Lower)
	q_end := math.Min(o.Upper, p.Upper)
	return OpenInterval{q_start, q_end}
}

// IsEmpty tests if an OpenInterval is empty. An OpenInterval is assumed empty if either
// bound is NaN.
func (o OpenInterval) IsEmpty() bool {
	return math.IsNaN(o.Lower) || math.IsNaN(o.Upper)
}
