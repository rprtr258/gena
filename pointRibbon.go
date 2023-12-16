package gena

import (
	"image/color"
	"math"
)

type pointRibbon struct {
	r     float64
	iters int
}

// NewPointRibbon returns a pointRibbon object.
func NewPointRibbon(r float64, iters int) *pointRibbon {
	return &pointRibbon{
		r:     r,
		iters: iters,
	}
}

// Generative draws a point ribbon image.
// TODO: make the point as parameters.
func (s *pointRibbon) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(c.LineWidth)

	var t float64
	dt := 0.0001
	for i := 0; i < s.iters; i++ {
		delta := 2.0*s.r*math.Cos(4.0*dt*t) + s.r*math.Cos(t)
		ctex.SetRGBA255(color.RGBA{
			uint8(delta),
			uint8(2*s.r*math.Sin(t) - s.r*math.Cos(3*dt*t)),
			100,
			0},
			10,
		)
		ctex.DrawPoint(Mul(complex(
			2*math.Sin(2*t*dt)+math.Cos(t*dt),
			2*math.Sin(t*dt)-math.Sin(5*t),
		), s.r)+c.Size()/2, 1.0)
		ctex.Stroke()
		t += 0.01
		dt += 0.1
	}
}
