package curvy

import (
	"math"

	"gonum.org/v1/plot/tools/bezier"
	"gonum.org/v1/plot/vg"
)

func calculateArcLengths(curve bezier.Curve, samples int) []float64 {
	total := .0
	samples *= len(curve)
	arr := make([]float64, samples)

	arr[0] = 0
	last := curve.Point(0)
	for i := 1; i < samples; i++ {
		t := float64(i) / float64(samples-1)
		p := curve.Point(t)

		total += distance(p, last)
		last = p
		arr[i] = total
	}

	return arr
}

func distance(a, b vg.Point) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

func scaleT(u float64, arcLengths []float64) float64 {
	targetLen := u * arcLengths[len(arcLengths)-1]

	start, offset := 0, len(arcLengths)-1

	for offset > 0 {
		if arcLengths[start+offset/2] < targetLen {
			start += offset/2 + 1 // + 1 because only <
			offset = offset / 2
		} else { // >=
			offset = offset / 2
		}
	}

	if start == len(arcLengths)-1 {
		return 1
	}

	return (float64(start) + (targetLen-arcLengths[start])/(arcLengths[start+1]-targetLen)) / float64(len(arcLengths)-1)
}
