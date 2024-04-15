package main

import . "github.com/rprtr258/gena"

func sine() {
	const W = 1200
	const H = 60
	dc := NewContext(complex(W, H))
	dc.SetColor(White)
	dc.Clear()

	for i, a := range RangeF64(0, 8*PI*2, W) {
		dc.LineTo(complex(float64(i), (Sin(a)+1)/2*H))
	}
	dc.ClosePath()

	dc.SetColor(ColorHex("#3E606F"))
	dc.FillPreserve()

	dc.SetColor(ColorHex("#19344180"))
	dc.SetLineWidth(8)
	dc.Stroke()

	SavePNG("sine.png", dc.Image())
}
