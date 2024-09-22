package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Janus would draw an image with multiple circles split at its center with random noise in the horizontal direction.
// TODO not finished.
func Janus(dc *Context, colorSchema []color.RGBA, fg color.RGBA, decay float64) {
	const r = 0.3

	im := dc.Image()
	dc.SetColor(Black)
	dc.Clear()
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
			dc.TransformAdd(Scale(P(s, s)))
			dc.DrawArc(P(x1, y1), 1.0, PI*3/2+theta, PI*5/2+theta)
			dc.SetColor(clr)
			dc.Fill()
			dc.DrawArc(P(x2, y2), 1.0, PI/2+theta, PI*3/2+theta)
			dc.SetColor(clr)
			dc.Fill()
		})
	}
}

func janus() *image.RGBA {
	DarkRed := []color.RGBA{
		{48, 19, 21, 255},
		{71, 15, 16, 255},
		{92, 10, 12, 255},
		{115, 5, 4, 255},
		{139, 0, 3, 255},
		{163, 32, 1, 255},
		{185, 66, 0, 255},
		{208, 98, 1, 255},
		{231, 133, 0, 255},
		{254, 165, 0, 255},
	}

	dc := NewContext(Diag(500))
	Janus(dc, DarkRed, LightPink, 0.2)
	return dc.Image()
}
