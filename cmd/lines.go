package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func lines() *image.RGBA {
	const W = 1024
	const H = 1024
	dc := NewContext(P(W, H))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Clear()
	for range Range(1000) {
		x1 := Random() * W
		y1 := Random() * H
		x2 := Random() * W
		y2 := Random() * H
		r := Random()
		g := Random()
		b := Random()
		a := Random()*0.5 + 0.5
		w := Random()*4 + 1
		dc.SetColor(ColorRGBA(r, g, b, a))
		dc.SetLineWidth(w)
		dc.DrawLine(P(x1, y1), P(x2, y2))
		dc.Stroke()
	}
	return dc.Image()
}
