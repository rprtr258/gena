package main

import . "github.com/rprtr258/gena"

func sine() {
	const W = 1200
	const H = 60
	dc := NewContext(complex(W, H))
	// dc.SetHexColor("#FFFFFF")
	// dc.Clear()
	dc.TransformAdd(Translate(complex(0.95, 0.75)))
	dc.TransformAdd(Scale(complex(W, H) / 2))
	for i, a := range RangeF64(0, 2*PI*8, W) {
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
