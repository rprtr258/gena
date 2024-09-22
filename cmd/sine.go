package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func sine() *image.RGBA {
	const W = 1200
	const H = 60
	dc := NewContext(P(W, H))
	dc.SetColor(White)
	dc.Clear()

	for i, a := range RangeF64(0, 8*PI*2, W) {
		dc.LineTo(P(i, (Sin(a)+1)/2*H))
	}
	dc.ClosePath()

	dc.SetColor(ColorHex("#3E606F"))
	dc.FillPreserve()

	dc.SetColor(ColorHex("#19344180"))
	dc.SetLineWidth(8)
	dc.Stroke()

	return dc.Image()
}
