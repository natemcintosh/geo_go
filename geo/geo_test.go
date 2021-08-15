package gogeo

import (
	"math"
	"testing"
)

func TestPointAngle(t *testing.T) {
	testCases := []struct {
		desc string
		in   Point
		out  float64
	}{
		{
			desc: "Point on the x-axis",
			in:   Point{0, 0},
			out:  0,
		},
		{
			desc: "Point on the y-axis",
			in:   Point{0, 1},
			out:  math.Pi / 2,
		},
		{
			desc: "Point negative x-axis",
			in:   Point{-1, 0},
			out:  math.Pi,
		},
		{
			desc: "Point negative y-axis",
			in:   Point{0, -1},
			out:  -math.Pi / 2,
		},
		{
			desc: "45 deg angle",
			in:   Point{1, 1},
			out:  math.Pi / 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Angle(); got != tC.out {
				t.Errorf("Angle() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointAngle(b *testing.B) {
	benchmarks := []struct {
		desc string
		in   Point
	}{
		{"positive x-axis", Point{1, 0}},
		{"positive y-axis", Point{0, 1}},
		{"negative y-axis", Point{-1, 0}},
		{"random point 1", Point{3.4, -2.3}},
		{"random point 2", Point{100.2, 7.6}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.Angle()
			}

		})
	}
}

func TestPointRotate(t *testing.T) {
	testCases := []struct {
		desc  string
		in    Point
		angle float64
		out   Point
	}{
		{
			desc:  "rotate from x-axis by 45 deg",
			in:    Point{1, 0},
			angle: math.Pi / 4,
			out:   Point{math.Cos(math.Pi / 4), math.Sin(math.Pi / 4)},
		},
		{
			desc:  "rotate from x-axis 90 deg",
			in:    Point{1, 0},
			angle: math.Pi / 2,
			out:   Point{0, 1},
		},
		{
			desc:  "rotate from x-axis 180 deg",
			in:    Point{1, 0},
			angle: math.Pi,
			out:   Point{-1, 0},
		},
		{
			desc:  "rotate from x-axis 270 deg",
			in:    Point{1, 0},
			angle: math.Pi * 3 / 2,
			out:   Point{0, -1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Rotate(tC.angle); !got.AlmostEquals(tC.out) {
				t.Errorf("Rotate() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointRotate(b *testing.B) {
	benchmarks := []struct {
		desc  string
		in    Point
		angle float64
	}{
		{"Rotate x-axis with no angle", Point{1, 0}, 0},
		{"Rotate x-axis with 45 deg angle", Point{1, 0}, math.Pi / 4},
		{"Rotate random point by 90 deg angle", Point{3.4, -2.3}, math.Pi / 2},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.Rotate(bm.angle)
			}

		})
	}
}

func TestPointsXIntercept(t *testing.T) {
	testCases := []struct {
		desc string
		p    Point
		q    Point
		out  float64
	}{
		{
			desc: "Two points on the y-axis",
			p:    Point{0, -1},
			q:    Point{0, 1},
			out:  0,
		},
		{
			desc: "Two points vertically stacked",
			p:    Point{1, -1},
			q:    Point{1, 1},
			out:  1,
		},
		{
			desc: "Two points forming a line passing through the origin",
			p:    Point{1, 1},
			q:    Point{-1, -1},
			out:  0,
		},
		{
			desc: "Two points forming a line passing through 1.0",
			p:    Point{0, -1},
			q:    Point{2, 1},
			out:  1,
		},
		{
			desc: "Two points forming a line passing through 11",
			p:    Point{3, 4},
			q:    Point{5, 3},
			out:  11,
		},
		{
			desc: "A horizontal line",
			p:    Point{0, 1},
			q:    Point{10, 1},
			out:  math.Inf(1),
		},
		{
			desc: "Another horizontal line",
			p:    Point{0, 1},
			q:    Point{-10, 1},
			out:  math.Inf(1),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.p.XIntercept(tC.q); got != tC.out {
				t.Errorf("XIntercept() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointXIntercept(b *testing.B) {
	benchmarks := []struct {
		desc string
		p    Point
		q    Point
	}{
		{"Intercept of two vertical lines", Point{1, 0}, Point{1, 1}},
		{"Intercept of two horizontal lines", Point{0, 1}, Point{1, 1}},
		{"Intercept of two diagonal lines", Point{1, 1}, Point{3, 3}},
		{"Intercept of two lines with same slope", Point{1, 1}, Point{2, 2}},
		{"Intercept of two lines with different slope", Point{1, 1}, Point{2, 3}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.p.XIntercept(bm.q)
			}

		})
	}
}

func TestPointMagnitude(t *testing.T) {
	testCases := []struct {
		desc string
		in   Point
		out  float64
	}{
		{
			desc: "Should have magnitude of 1",
			in:   Point{1, 0},
			out:  1,
		},
		{
			desc: "Should have magnitude of 2",
			in:   Point{2, 0},
			out:  2,
		},
		{
			desc: "Should have magnitude of sqrt(2)",
			in:   Point{1, 1},
			out:  math.Sqrt(2),
		},
		{
			desc: "Should have magnitude of sqrt(8)",
			in:   Point{2, 2},
			out:  math.Sqrt(8),
		},
		{
			desc: "A 3-4-5 triangle",
			in:   Point{3, 4},
			out:  5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Magnitude(); got != tC.out {
				t.Errorf("Magnitude() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointMagnitude(b *testing.B) {
	benchmarks := []struct {
		desc string
		p    Point
	}{
		{"Point with magnitude 1", Point{1, 0}},
		{"Point with magnitude 2", Point{2, 0}},
		{"Point with magnitude sqrt(2)", Point{1, 1}},
		{"Point with magnitude 5", Point{3, 4}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.p.Magnitude()
			}

		})
	}
}

func TestPointNormalize(t *testing.T) {
	testCases := []struct {
		desc string
		in   Point
		out  Point
	}{
		{
			desc: "No change because already magnitude 1 along x-axis",
			in:   Point{1, 0},
			out:  Point{1, 0},
		},
		{
			desc: "No change because already magnitude 1 along y-axis",
			in:   Point{0, 1},
			out:  Point{0, 1},
		},
		{
			desc: "Should be Point with values of sqrt(2)/2",
			in:   Point{1, 1},
			out:  Point{math.Sqrt(2) / 2, math.Sqrt(2) / 2},
		},
		{
			desc: "A more complicated example",
			in:   Point{4, 5},
			out:  Point{4 / 6.4031242374328485, 5 / 6.4031242374328485},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Normalize(); !got.AlmostEquals(tC.out) {
				t.Errorf("Normalize() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointNormalize(b *testing.B) {
	benchmarks := []struct {
		desc string
		p    Point
	}{
		{"Point with magnitude 1", Point{1, 0}},
		{"Point with magnitude 2", Point{2, 0}},
		{"Point with magnitude sqrt(2)", Point{1, 1}},
		{"Point with magnitude 5", Point{3, 4}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.p.Normalize()
			}

		})
	}
}

func TestPointDotProduct(t *testing.T) {
	testCases := []struct {
		desc string
		p1   Point
		p2   Point
		out  float64
	}{
		{
			desc: "perpendicular vectors have dot product of 0",
			p1:   Point{1, 0},
			p2:   Point{0, 1},
			out:  0,
		},
		{
			desc: "identical vectors have dot product of the same magnitude",
			p1:   Point{1, 0},
			p2:   Point{1, 0},
			out:  1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.p1.DotProduct(tC.p2); got != tC.out {
				t.Errorf("DotProduct() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointDotProduct(b *testing.B) {
	benchmarks := []struct {
		desc string
		p    Point
		q    Point
	}{
		{"Point with magnitude 1", Point{1, 0}, Point{1, 0}},
		{"Point with magnitude 2", Point{2, 0}, Point{2, 0}},
		{"Two Points with random numbers", Point{3.4, -2.3}, Point{100.2, 7.6}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.p.Magnitude()
			}

		})
	}
}

func TestLineSegmentAdd(t *testing.T) {
	testCases := []struct {
		desc string
		l    LineSegment
		p    Point
		out  LineSegment
	}{
		{
			desc: "x-axis line plus nothing",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{0, 0},
			out:  LineSegment{Point{0, 0}, Point{1, 0}},
		},
		{
			desc: "x-axis line plus Point{1, 1}",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{1, 1},
			out:  LineSegment{Point{1, 1}, Point{2, 1}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.l.Plus(tC.p); !got.Equals(tC.out) {
				t.Errorf("Add() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestLineSegmentSubtract(t *testing.T) {
	testCases := []struct {
		desc string
		l    LineSegment
		p    Point
		out  LineSegment
	}{
		{
			desc: "x-axis line plus nothing",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{0, 0},
			out:  LineSegment{Point{0, 0}, Point{1, 0}},
		},
		{
			desc: "x-axis line plus Point{1, 1}",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{1, 1},
			out:  LineSegment{Point{-1, -1}, Point{0, -1}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.l.Minus(tC.p); !got.Equals(tC.out) {
				t.Errorf("Add() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestLineSegmentAngle(t *testing.T) {
	testCases := []struct {
		desc string
		in   LineSegment
		out  float64
	}{
		{
			desc: "Horizontal Line",
			in:   LineSegment{Point{0, 0}, Point{1, 0}},
			out:  0,
		},
		{
			desc: "Vertical Line",
			in:   LineSegment{Point{0, 0}, Point{0, 1}},
			out:  math.Pi / 2,
		},
		{
			desc: "45 deg angle",
			in:   LineSegment{Point{1, 1}, Point{2, 2}},
			out:  math.Pi / 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Angle(); got != tC.out {
				t.Errorf("Angle() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkLineSegmentAngle(b *testing.B) {
	benchmarks := []struct {
		desc string
		in   LineSegment
	}{
		{"positive x-axis", LineSegment{Point{1, 0}, Point{2, 0}}},
		{"positive y-axis", LineSegment{Point{0, 1}, Point{0, 2}}},
		{"negative y-axis", LineSegment{Point{-1, 0}, Point{-2, 0}}},
		{"random point 1", LineSegment{Point{3.4, -2.3}, Point{100.2, 7.6}}},
		{"random point 2", LineSegment{Point{23.554, 3990.2}, Point{0, 5.45345}}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.Angle()
			}

		})
	}
}

func TestLineSegmentRotateAboutOrigin(t *testing.T) {
	testCases := []struct {
		desc  string
		l     LineSegment
		angle float64
		out   LineSegment
	}{
		{
			desc:  "x-axis line with no angle",
			l:     LineSegment{Point{0, 0}, Point{1, 0}},
			angle: 0,
			out:   LineSegment{Point{0, 0}, Point{1, 0}},
		},
		{
			desc:  "y-axis line with -90 deg angle",
			l:     LineSegment{Point{0, 0}, Point{0, 1}},
			angle: -math.Pi / 2,
			out:   LineSegment{Point{0, 0}, Point{1, 0}},
		},
		{
			desc:  "x-axis line by 90 deg",
			l:     LineSegment{Point{1, 0}, Point{2, 0}},
			angle: math.Pi / 2,
			out:   LineSegment{Point{0, 1}, Point{0, 2}},
		},
		{
			desc:  "line at 45 deg rotated by 180 deg",
			l:     LineSegment{Point{1, 1}, Point{2, 2}},
			angle: math.Pi,
			out:   LineSegment{Point{-1, -1}, Point{-2, -2}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.l.RotateAboutOrigin(tC.angle); !got.AlmostEquals(tC.out) {
				t.Errorf("RotateAboutOrigin() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkLineSegmentRotateAboutOrigin(b *testing.B) {
	benchmarks := []struct {
		desc  string
		l1    LineSegment
		angle float64
	}{
		{"Rotate x-axis with no angle", LineSegment{Point{1, 0}, Point{2, 0}}, 0},
		{"Rotate x-axis with 45 deg angle", LineSegment{Point{1, 0}, Point{2, 0}}, math.Pi / 4},
		{"Rotate y-axis with no angle", LineSegment{Point{0, 1}, Point{0, 2}}, 0},
		{"Rotate line at 45 deg angle by 90 deg", LineSegment{Point{0, 0}, Point{1, 1}}, math.Pi / 2},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.l1.RotateAboutOrigin(bm.angle)
			}

		})
	}
}

func TestOpenIntervalIsEmpty(t *testing.T) {
	testCases := []struct {
		desc string
		in   OpenInterval
		out  bool
	}{
		{
			desc: "both NaN",
			in:   OpenInterval{math.NaN(), math.NaN()},
			out:  true,
		},
		{
			desc: "first NaN",
			in:   OpenInterval{math.NaN(), 1},
			out:  true,
		},
		{
			desc: "second NaN",
			in:   OpenInterval{1, math.NaN()},
			out:  true,
		},
		{
			desc: "both are regular numbers",
			in:   OpenInterval{1, 2},
			out:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.IsEmpty(); got != tC.out {
				t.Errorf("IsEmpty() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkOpenIntervalIsEmpty(b *testing.B) {
	benchmarks := []struct {
		desc string
		in   OpenInterval
	}{
		{"empty", OpenInterval{math.NaN(), math.NaN()}},
		{"first NaN", OpenInterval{math.NaN(), 1}},
		{"second NaN", OpenInterval{1, math.NaN()}},
		{"both are regular numbers", OpenInterval{1, 2}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.IsEmpty()
			}

		})
	}
}

func TestOpenIntervalIntersection(t *testing.T) {
	testCases := []struct {
		desc string
		o1   OpenInterval
		o2   OpenInterval
		out  OpenInterval
	}{
		{
			desc: "No overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{3, 4},
			out:  OpenInterval{math.NaN(), math.NaN()},
		},
		{
			desc: "Some overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{1.5, 2.5},
			out:  OpenInterval{1.5, 2},
		},
		{
			desc: "Complete overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{1, 2},
			out:  OpenInterval{1, 2},
		},
		{
			desc: "Single number overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{2, 3},
			out:  OpenInterval{2, 2},
		},
		{
			desc: "Some more overlap",
			o1:   OpenInterval{-10, -5},
			o2:   OpenInterval{-7.3, -2},
			out:  OpenInterval{-7.3, -5},
		},
		{
			desc: "Both NaN",
			o1:   OpenInterval{math.NaN(), math.NaN()},
			o2:   OpenInterval{math.NaN(), math.NaN()},
			out:  OpenInterval{math.NaN(), math.NaN()},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.o1.Intersection(tC.o2); !got.Equals(tC.out) {
				t.Errorf("Intersection() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkOpenIntervalIntersection(b *testing.B) {
	benchmarks := []struct {
		desc string
		o1   OpenInterval
		o2   OpenInterval
	}{
		{"No overlap", OpenInterval{1, 2}, OpenInterval{3, 4}},
		{"Some overlap", OpenInterval{1, 2}, OpenInterval{1.5, 2.5}},
		{"Complete overlap", OpenInterval{1, 2}, OpenInterval{1, 2}},
		{"Single number overlap", OpenInterval{1, 2}, OpenInterval{2, 3}},
		{"Some more overlap", OpenInterval{-10, -5}, OpenInterval{-7.3, -2}},
		{"Both NaN", OpenInterval{math.NaN(), math.NaN()}, OpenInterval{math.NaN(), math.NaN()}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.o1.Intersection(bm.o2)
			}

		})
	}
}

func TestLineSegmentIntersects(t *testing.T) {
	testCases := []struct {
		desc string
		l1   LineSegment
		l2   LineSegment
		out  bool
	}{
		{
			desc: "Two segments definitely cross",
			l1:   LineSegment{Point{0, 0}, Point{1, 1}},
			l2:   LineSegment{Point{1, 0}, Point{0, 1}},
			out:  true,
		},
		{
			desc: "Two segments definitely don't cross",
			l1:   LineSegment{Point{0, 0}, Point{1, 1}},
			l2:   LineSegment{Point{-10, -10}, Point{-20, -20}},
			out:  false,
		},
		{
			desc: "They meet at a one end",
			l1:   LineSegment{Point{0, 0}, Point{0, 1}},
			l2:   LineSegment{Point{1, 1}, Point{0, 1}},
			out:  true,
		},
		{
			desc: "They almost meet",
			l1:   LineSegment{Point{0, 0}, Point{1, 1}},
			l2:   LineSegment{Point{1.1, 1.1}, Point{1.2, 1.2}},
			out:  false,
		},
		{
			desc: "They overlap along the line y = x",
			l1:   LineSegment{Point{0, 0}, Point{1, 1}},
			l2:   LineSegment{Point{0.9, 0.9}, Point{1.1, 1.1}},
			out:  true,
		},
		{
			desc: "They overlap along the line y = 0",
			l1:   LineSegment{Point{0, 0}, Point{1, 0}},
			l2:   LineSegment{Point{0.9, 0}, Point{1.1, 0}},
			out:  true,
		},
		{
			desc: "They overlap along the line x = 0",
			l1:   LineSegment{Point{0, 0}, Point{0, 1}},
			l2:   LineSegment{Point{0, 0.9}, Point{0, 1.1}},
			out:  true,
		},
		{
			desc: "They cross at (0.5, 0)",
			l1:   LineSegment{Point{0, 0}, Point{1, 0}},
			l2:   LineSegment{Point{0.5, 1}, Point{0.5, -1}},
			out:  true,
		},
		{
			desc: "They cross at (1, 1)",
			l1:   LineSegment{Point{0, 0}, Point{2, 2}},
			l2:   LineSegment{Point{1, 0}, Point{0, 1}},
			out:  true,
		},
		{
			desc: "One crosses the end of the other",
			l1:   LineSegment{Point{0, 0}, Point{1, 0}},
			l2:   LineSegment{Point{1, 1}, Point{1, -1}},
			out:  true,
		},
		{
			desc: "The two segments don't cross",
			l1:   LineSegment{Point{100, 100}, Point{200, 200}},
			l2:   LineSegment{Point{-100, -100}, Point{-200, -200}},
			out:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.l1.Intersects(tC.l2); got != tC.out {
				t.Errorf("Intersects() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkLineSegmentsIntersects(b *testing.B) {
	benchmarks := []struct {
		desc string
		l1   LineSegment
		l2   LineSegment
	}{
		{
			"Two that definitely overlap",
			LineSegment{Point{0, 0}, Point{1, 1}},
			LineSegment{Point{1, 0}, Point{0, 1}},
		},
		{
			"Two that definitely don't overlap",
			LineSegment{Point{0, 0}, Point{1, 1}},
			LineSegment{Point{2, 0}, Point{3, 1}},
		},
		{
			"Two that overlap on one point",
			LineSegment{Point{0, 0}, Point{0, 1}},
			LineSegment{Point{1, 1}, Point{0, 1}},
		},
		{
			"They overlap along a section",
			LineSegment{Point{0, 0}, Point{1, 1}},
			LineSegment{Point{0.9, 0.9}, Point{1.1, 1.1}},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.l1.Intersects(bm.l2)
			}

		})
	}
}

func TestTriangleEquals(t *testing.T) {
	testCases := []struct {
		desc string
		t1   Triangle
		t2   Triangle
		out  bool
	}{
		{
			desc: "Two equal triangles",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			out:  true,
		},
		{
			desc: "They are the same but one is rotated",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0, 1}, Point{0, 0}, Point{1, 0}},
			out:  true,
		},
		{
			desc: "They are the same but one is rotated further",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{1, 0}, Point{0, 1}, Point{0, 0}},
			out:  true,
		},
		{
			desc: "They are not the same",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0, 1}, Point{5, 3}, Point{1, 0}},
			out:  false,
		},
		{
			desc: "They are almost the same",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0.00001, 0.00001}, Point{1, 0}, Point{0, 1}},
			out:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.t1.Equals(tC.t2); got != tC.out {
				t.Errorf("Equals() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkTriangleEquals(b *testing.B) {
	benchmarks := []struct {
		desc string
		t1   Triangle
		t2   Triangle
	}{
		{
			desc: "Two equal triangles",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
		},
		{
			desc: "They are the same but one is rotated",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0, 1}, Point{0, 0}, Point{1, 0}},
		},
		{
			desc: "They are the same but one is rotated further",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{1, 0}, Point{0, 1}, Point{0, 0}},
		},
		{
			desc: "The are equal, but on the final permutation",
			t1:   Triangle{Point{0, 0}, Point{1, 0}, Point{0, 1}},
			t2:   Triangle{Point{0, 1}, Point{1, 0}, Point{0, 0}},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.t1.Equals(bm.t2)
			}

		})
	}
}
