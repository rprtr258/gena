package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func linewidth() *image.RGBA {
	dc := NewContext(P(1000, 1000))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(ColorRGB(0, 0, 0))
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(P(x+50, 0), P(x-50, 1000))
		dc.SetLineWidth(w)
		dc.Stroke()
		w += 0.1
	}
	return dc.Image()
}
