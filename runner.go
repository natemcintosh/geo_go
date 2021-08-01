package main

import (
	"fmt"

	gogeo "example.com/gogeo/geo"
)

func main() {
	p1 := gogeo.Point{X: 0, Y: 0}
	p2 := gogeo.Point{X: 1, Y: 1}
	p3 := gogeo.Point{X: 1, Y: 0}
	p4 := gogeo.Point{X: 2, Y: 1}
	p5 := gogeo.Point{X: 0.5, Y: 0}
	p6 := gogeo.Point{X: 0.5, Y: 1}
	l1 := gogeo.LineSegment{P1: p1, P2: p2}
	l2 := gogeo.LineSegment{P1: p3, P2: p4}
	l3 := gogeo.LineSegment{P1: p5, P2: p6}
	fmt.Println(l1.Angle())
	fmt.Println(l2.Angle())
	fmt.Println(l3.Angle())
}
