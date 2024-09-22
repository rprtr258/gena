package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a point ribbon image.
// TODO: make the point as parameters.
func PointRibbon(dc *Context, lineWidth, r float64, iters int) {
	im := dc.Image()
	dc.SetColor(Lavender)
	dc.Clear()
	dc.SetLineWidth(lineWidth)

	var t float64
	dt := 0.0001
	for range Range(iters) {
		delta := 2.0*r*Cos(4.0*dt*t) + r*Cos(t)
		dc.SetColor(color.NRGBA{
			uint8(delta),
			uint8(2*r*Sin(t) - r*Cos(3*dt*t)),
			100,
			10,
		})
		dc.DrawPoint(P(
			2*Sin(2*t*dt)+Cos(t*dt),
			2*Sin(t*dt)-Sin(5*t),
		)*Coeff(r)+Size(im)/2, 1.0)
		dc.Stroke()
		t += 0.01
		dt += 0.1
	}
}

func pointRibbon() *image.RGBA {
	dc := NewContext(Diag(500))
	PointRibbon(dc, 2, 50, 150000)
	return dc.Image()
}
