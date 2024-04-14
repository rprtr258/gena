package main

import . "github.com/rprtr258/gena"

func openfill() {
	dc := NewContext(complex(1000, 1000))
	for _, f := range RangeV2_2(10, 10) {
		a1 := PI * Random() * 2
		a2 := PI * (Random() + 0.5)
		dc.DrawArc(f*100+50, 40, a1, a1+a2)
		// dc.ClosePath()
	}
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.FillPreserve()
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.SetLineWidth(8)
	dc.StrokePreserve()
	dc.SetColor(ColorRGB(1, 0, 0))
	dc.SetLineWidth(4)
	dc.StrokePreserve()
	SavePNG("openfill.png", dc.Image())
}
