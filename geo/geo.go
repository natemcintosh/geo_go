// Package gogeo provides simple Point and Line Segment types. It also provides
// functionality for rotating Points and Line Segments, and checking if Line Segments
// intersect.
package gogeo

import (
	"math"
)

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

// Rotate rotates a Point by the given angle in radians.
func (p Point) Rotate(angle float64) Point {
	s := math.Sin(angle)
	c := math.Cos(angle)
	return Point{
		X: c*p.X - s*p.Y,
		Y: s*p.X + c*p.Y,
	}
}

// XIntercept will calculate the x-intercept of an infinite line, as defined by the two
// points `p` and `q`. If the line is horizontal, returns +Inf.
func (p Point) XIntercept(q Point) float64 {
	i := p.X - (p.Y * (q.X - p.X) / (q.Y - p.Y))
	if math.IsInf(i, 0) {
		return math.Inf(1)
	} else {
		return i
	}
}

// Magnitude returns the 2-norm of a Point, interpreting the Point as a vector.
func (p Point) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Normalize will normalize a Point to unit magnitude.
func (p Point) Normalize() Point {
	return p.Divide(p.Magnitude())
}

// DotProduct is the dot product of two Points, intepreted as vectors.
func (p Point) DotProduct(q Point) float64 {
	return p.X*q.X + p.Y*q.Y
}

// LineSegment is a line segment in 2D space. It is defined by two Points.
type LineSegment struct {
	P1 Point
	P2 Point
}

// Equals tests if two LineSegments are equal.
func (l LineSegment) Equals(m LineSegment) bool {
	return (l.P1.X == m.P1.X) && (l.P1.Y == m.P1.Y) && (l.P2.X == m.P2.X) && (l.P2.Y == m.P2.Y)
}

func (l LineSegment) AlmostEquals(m LineSegment) bool {
	return l.P1.AlmostEquals(m.P1) && l.P2.AlmostEquals(m.P2)
}

// Plus adds the x and y components of a Point to a LineSegment.
func (l LineSegment) Plus(p Point) LineSegment {
	return LineSegment{l.P1.Plus(p), l.P2.Plus(p)}
}

// Minus subtracts the x and y components of a Point to a LineSegment.
func (l LineSegment) Minus(p Point) LineSegment {
	return LineSegment{l.P1.Minus(p), l.P2.Minus(p)}
}

// Angle calculates the angle of a LineSegment in radians from where it intersects the positive x-axis.
func (l LineSegment) Angle() float64 {
	return math.Atan2(l.P2.Y-l.P1.Y, l.P2.X-l.P1.X)
}

// RotateAboutOrigin rotates a LineSegment by the given angle in radians about the origin.
func (l LineSegment) RotateAboutOrigin(angle float64) LineSegment {
	return LineSegment{l.P1.Rotate(angle), l.P2.Rotate(angle)}
}

// sign returns +1 for positive, 0 for 0.0, and -1 for negative
func sign(x float64) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

// almost_zero will check if a number is almost equal to 0
func almost_zero(x float64) bool {
	return math.Abs(x) < float64EqualityThreshold
}

// sign_close_to_zero is very similar to `sign()`, but will check if `x` is almost equal
// to zero.
func sign_close_to_zero(x float64) int {
	if almost_zero(x) {
		return 0
	} else if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

// XIntercept will return the x-coordinate of the intersection of a LineSegment with the
// x-axis.
// Looking at the signs of the y-values of the vertices, there are the following cases:
// 1. both zero -> OpenInterval between x vertices
// 2. both negative -> OpenInterval of NaN -> NaN. I.e. nothing will match
// 3. both positive -> OpenInterval of NaN -> NaN. I.e. nothing will match
// 4. one zero, one negative -> OpenInterval of the one vertex on the x-axis
// 5. one zero, one positive -> OpenInterval of the one vertex on the x-axis
// 6. one negative, one positive -> OpenInterval of the intersection
func (l LineSegment) XIntercept() OpenInterval {
	// First make sure neither point is NaN. If so, return an empty OpenInterval.
	if math.IsNaN(l.P1.X) || math.IsNaN(l.P2.X) {
		return OpenInterval{math.NaN(), math.NaN()}
	}

	// Get the sign of the y points of the line
	sign_y1 := sign_close_to_zero(l.P1.Y)
	sign_y2 := sign_close_to_zero(l.P2.Y)
	sum_of_signs := float64(sign_y1 + sign_y2)

	if (sign_y1 == 0) && (sign_y2 == 0) {
		// 1) both zero -> OpenInterval between x vertices
		return OpenInterval{l.P1.X, l.P2.X}
	} else if math.Abs(sum_of_signs) == 2 {
		// 2 & 3) both points are above or below the x-axis, no intersection
		return OpenInterval{math.NaN(), math.NaN()}
	} else if sum_of_signs == -1 {
		// 4) one zero, one negative -> OpenInterval of the one vertex on the x-axis
		if sign_y1 < 0 { // p1 is below x-axis, p2 is on the x-axis
			return OpenInterval{l.P2.X, l.P2.X}
		} else { // p2 is below x-axis, p1 is on the x-axis
			return OpenInterval{l.P1.X, l.P1.X}
		}
	} else if sum_of_signs == 1 {
		// 5) one zero, one positive -> OpenInterval of the one vertex on the x-axis
		if sign_y2 > 0 { // p2 is above x-axis, p1 is on the x-axis
			return OpenInterval{l.P1.X, l.P1.X}
		} else { // p1 is above x-axis, p2 is on the x-axis
			return OpenInterval{l.P2.X, l.P2.X}
		}
	} else {
		// 6) one negative, one positive -> OpenInterval of the intersection
		// Get the x-intercept of the line
		x_intercept := l.P1.XIntercept(l.P2)
		if math.IsInf(x_intercept, 0) { // the line is horizontal
			// Make the OpenInterval with NaNs
			return OpenInterval{math.NaN(), math.NaN()}
		} else {
			return OpenInterval{x_intercept, x_intercept}
		}

	}
}

// Intersects will determine if two LineSegments intersect. They are said to intersect
// if any point on the segments, including the endpoints intersects.
func (l1 LineSegment) Intersects(l2 LineSegment) bool {
	// Pick a point on segment 1 and make it the origin. Move other points relative to it.
	l1_translated := l1.Minus(l1.P1)
	l2_translated := l2.Minus(l1.P1)

	// Rotate all points so that segment 1 is aligned with the x-axis.
	angle_to_rotate_through := -l1_translated.Angle()
	l1_rotated := l1_translated.RotateAboutOrigin(angle_to_rotate_through)
	l2_rotated := l2_translated.RotateAboutOrigin(angle_to_rotate_through)

	// Find the x-intercept of segment 2
	l2_x_intercept := l2_rotated.XIntercept()

	// Is it between the two points on segment 1?
	l1_x_intercept := OpenInterval{l1_rotated.P1.X, l1_rotated.P2.X}

	return !l1_x_intercept.Intersection(l2_x_intercept).IsEmpty()
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

// Triangle is made up of three Points.
type Triangle struct {
	P1 Point
	P2 Point
	P3 Point
}

// Equals compares all three Points of a Triangle. The points do not necessarily
// need to be in the same order. I.e. they can be in any permutation of the three
func (t Triangle) Equals(u Triangle) bool {
	return (t.P1.Equals(u.P1) && t.P2.Equals(u.P2) && t.P3.Equals(u.P3)) ||
		(t.P1.Equals(u.P1) && t.P2.Equals(u.P3) && t.P3.Equals(u.P2)) ||
		(t.P1.Equals(u.P2) && t.P2.Equals(u.P1) && t.P3.Equals(u.P3)) ||
		(t.P1.Equals(u.P2) && t.P2.Equals(u.P3) && t.P3.Equals(u.P1)) ||
		(t.P1.Equals(u.P3) && t.P2.Equals(u.P1) && t.P3.Equals(u.P2)) ||
		(t.P1.Equals(u.P3) && t.P2.Equals(u.P2) && t.P3.Equals(u.P1))

}

// Area is the area of a Triangle.
func (t Triangle) Area() float64 {
	return 0.5 * math.Abs(
		t.P1.X*(t.P2.Y-t.P3.Y)+
			t.P2.X*(t.P3.Y-t.P1.Y)+
			t.P3.X*(t.P1.Y-t.P2.Y))
}
