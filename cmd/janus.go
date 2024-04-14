package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Janus would draw an image with multiple circles split at its center with random noise in the horizontal direction.
// TODO not finished.
func Janus(im *image.RGBA, colorSchema []color.RGBA, fg color.RGBA, decay float64) {
	const r = 0.3

	dc := NewContextFromRGBA(im)
	dc.SetColor(fg)

	s := 220.0
	for _, clr := range colorSchema {
		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(Size(im) / 2))

			theta := RandomF64(PI/4, 3*PI/4)
			x1, y1 := Cos(theta)*r, Sin(theta)*r
			x2, y2 := -x1, -y1

			noise := RandomF64(-Abs(y1), Abs(y1))
			y1 += noise
			y2 += noise

			s *= 0.836
			dc.TransformAdd(Scale(complex(s, s)))
			dc.DrawArc(complex(x1, y1), 1.0, PI*3/2+theta, PI*5/2+theta)
			dc.SetColor(clr)
			dc.Fill()
			dc.DrawArc(complex(x2, y2), 1.0, PI/2+theta, PI*3/2+theta)
			dc.SetColor(clr)
			dc.Fill()
		})
	}
}
