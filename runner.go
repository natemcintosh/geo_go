package main

import (
	"fmt"

	gogeo "example.com/gogeo/geo"
)

func main() {
	// p1 := gogeo.Point{X: 0, Y: 0}
	// p2 := gogeo.Point{X: 1, Y: 1}
	// p3 := gogeo.Point{X: 1, Y: 0}
	// p4 := gogeo.Point{X: 2, Y: 1}
	// p5 := gogeo.Point{X: 0.5, Y: 0}
	// p6 := gogeo.Point{X: 0.5, Y: 1}
	// l1 := gogeo.LineSegment{P1: p1, P2: p2}
	// l2 := gogeo.LineSegment{P1: p3, P2: p4}
	// l3 := gogeo.LineSegment{P1: p5, P2: p6}
	// fmt.Println(l1.Intersects(l2))
	l1 := gogeo.LineSegment{
		P1: gogeo.Point{X: 0, Y: 0},
		P2: gogeo.Point{X: 1, Y: 1},
	}

	l2 := gogeo.LineSegment{
		P1: gogeo.Point{X: 0.9, Y: 0.9},
		P2: gogeo.Point{X: 1.1, Y: 1.1},
	}

	l1_translated := l1.Minus(l1.P1)
	l2_translated := l2.Minus(l1.P1)
	fmt.Printf("l1_translated: %v\n", l1_translated)
	fmt.Printf("l2_translated: %v\n", l2_translated)

	// Rotate all points so that segment 1 is aligned with the x-axis.
	angle_to_rotate_through := -l1_translated.Angle()
	l1_rotated := l1_translated.RotateAboutOrigin(angle_to_rotate_through)
	any_difference := gogeo.Point{X: 0, Y: l1_rotated.P2.Y}

	l2_rotated := l2_translated.RotateAboutOrigin(angle_to_rotate_through)
	fmt.Printf("l1_rotated: %v\n", l1_rotated)
	fmt.Printf("l2_rotated: %v\n", l2_rotated)

	fmt.Printf("any difference: %v\n", any_difference)
	l2_with_diff := l2_rotated.Minus(any_difference)
	fmt.Printf("l2_with_diff: %v\n", l2_with_diff)

	// Find the x-intercept of segment 2
	l2_x_intercept := l2_rotated.XIntercept()
	fmt.Printf("l2_x_intercept: %v\n", l2_x_intercept)
}
