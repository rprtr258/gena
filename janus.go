package gena

import (
	"image/color"
	"math"
)

// Janus would draw an image with multiple circles split at its center with random noise in the horizontal direction.
// TODO not finished.
func Janus(c Canvas, colorSchema []color.RGBA, fg color.RGBA, n int, decay float64) {
	dc := NewContextForRGBA(c.Img())
	dc.SetColor(fg)
	s := 220.0
	r := 0.3

	for i := 0; i < n; i++ {
		// k := rand.Intn(len(colorSchema()))
		k := i
		dc.Stack(func(ctx *Context) {
			dc.Translate(c.Size() / 2)

			// theta += rand.Float64()*math.Pi/2
			theta := RandomFloat64(math.Pi/4, 3*math.Pi/4)
			x1, y1 := math.Cos(theta)*r, math.Sin(theta)*r
			x2, y2 := -x1, -y1

			noise := RandomFloat64(-math.Abs(y1), math.Abs(y1))
			y1 += noise
			y2 += noise

			// r -= r*j.decay
			s *= 0.836
			dc.Scale(complex(s, s))
			// r *= 0.836
			dc.DrawArc(complex(x1, y1), 1.0, math.Pi*3/2+theta, math.Pi*5/2+theta)
			dc.SetColor(colorSchema[k])
			dc.Fill()
			dc.DrawArc(complex(x2, y2), 1.0, math.Pi/2+theta, math.Pi*3/2+theta)
			dc.SetColor(colorSchema[k])
			dc.Fill()
		})
	}
}
