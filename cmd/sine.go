package main

import (
	"math"

	. "github.com/rprtr258/gena"
)

func sine() {
	const W = 1200
	const H = 60
	dc := NewContext(complex(W, H))
	// dc.SetHexColor("#FFFFFF")
	// dc.Clear()
	dc.TransformAdd(Translate(complex(0.95, 0.75)))
	dc.TransformAdd(Scale(complex(W, H) / 2))
	for i := range Range(W) {
		a := float64(i) * 2 * math.Pi / W * 8
		dc.LineTo(complex(float64(i), (math.Sin(a)+1)/2*H))
	}
	dc.ClosePath()
	dc.SetColor(ColorHex("#3E606F"))
	dc.FillPreserve()
	dc.SetColor(ColorHex("#19344180"))
	dc.SetLineWidth(8)
	dc.Stroke()
	SavePNG("sine.png", dc.Image())
}
