package main

import (
	"image"
	"image/color"
	"math"

	. "github.com/rprtr258/gena"
)

// Generative draws a point ribbon image.
// TODO: make the point as parameters.
func PointRibbon(c *image.RGBA, lineWidth, r float64, iters int) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(lineWidth)

	var t float64
	dt := 0.0001
	for range Range(iters) {
		delta := 2.0*r*math.Cos(4.0*dt*t) + r*math.Cos(t)
		dc.SetRGBA255(color.RGBA{
			uint8(delta),
			uint8(2*r*math.Sin(t) - r*math.Cos(3*dt*t)),
			100,
			0,
		}, 10)
		dc.DrawPoint(complex(
			2*math.Sin(2*t*dt)+math.Cos(t*dt),
			2*math.Sin(t*dt)-math.Sin(5*t),
		)*Coeff(r)+Size(c)/2, 1.0)
		dc.Stroke()
		t += 0.01
		dt += 0.1
	}
}
