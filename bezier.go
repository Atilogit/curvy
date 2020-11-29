package curvy

import (
	"gonum.org/v1/plot/tools/bezier"
	"gonum.org/v1/plot/vg"
)

type Bezier struct {
	curve      bezier.Curve
	arcLengths []float64
}

// Point returns the point at t along the curve, where 0 ≤ t ≤ 1.
func (b Bezier) Point(t float64) vg.Point {
	return b.curve.Point(t)
}

// Curve returns a slice of vg.Point, p, filled with points along the Bézier curve described by b.
// If the length of p is less than 2, the curve points are undefined. The length of p is not
// altered by the call.
func (b Bezier) Curve(p []vg.Point) []vg.Point {
	return b.curve.Curve(p)
}

// Point returns the point at exactly u percent along the length of the curve, where 0 ≤ u ≤ 1.
func (b Bezier) PointConstantSpacing(u float64) vg.Point {
	return b.curve.Point(scaleT(u, b.arcLengths))
}

// Curve returns a slice of vg.Point, p, filled with equally spaced points along the Bézier curve described by b.
// If the length of p is less than 2, the curve points are undefined. The length of p is not
// altered by the call.
func (b Bezier) CurveConstantSpacing(p []vg.Point) []vg.Point {
	for i, nf := 0, float64(len(p)-1); i < len(p); i++ {
		p[i] = b.PointConstantSpacing(float64(i) / nf)
	}
	return p
}

// NewBezier returns a Bézier curve initialized with the control points in cp.
func NewBezier(cp ...vg.Point) Bezier {
	curve := bezier.New(cp...)

	return Bezier{curve: curve, arcLengths: calculateArcLengths(curve, 100)}
}
