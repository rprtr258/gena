package gena

import (
	"image/color"
	"math"
)

// Generative draws a point ribbon image.
// TODO: make the point as parameters.
func PointRibbon(c Canvas, lineWidth float64, r float64, iters int) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(lineWidth)

	var t float64
	dt := 0.0001
	for i := 0; i < iters; i++ {
		delta := 2.0*r*math.Cos(4.0*dt*t) + r*math.Cos(t)
		ctex.SetRGBA255(color.RGBA{
			uint8(delta),
			uint8(2*r*math.Sin(t) - r*math.Cos(3*dt*t)),
			100,
			0},
			10,
		)
		ctex.DrawPoint(Mul(complex(
			2*math.Sin(2*t*dt)+math.Cos(t*dt),
			2*math.Sin(t*dt)-math.Sin(5*t),
		), r)+c.Size()/2, 1.0)
		ctex.Stroke()
		t += 0.01
		dt += 0.1
	}
}
