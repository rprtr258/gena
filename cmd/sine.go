package main

import (
	"math"

	"github.com/rprtr258/gena"
)

func sine() {
	const W = 1200
	const H = 60
	dc := gena.NewContext(complex(W, H))
	// dc.SetHexColor("#FFFFFF")
	// dc.Clear()
	dc.RelativeTo(complex(0.95, 0.75), func(dc *gena.Context) {
		dc.Scale(complex(W, H) / 2)
	})
	for i := range gena.Range(W) {
		a := float64(i) * 2 * math.Pi / W * 8
		dc.LineTo(complex(float64(i), (math.Sin(a)+1)/2*H))
	}
	dc.ClosePath()
	dc.SetColor(gena.ColorHex("#3E606F"))
	dc.FillPreserve()
	dc.SetColor(gena.ColorHex("#19344180"))
	dc.SetLineWidth(8)
	dc.Stroke()
	gena.SavePNG("sine.png", dc.Image())
}
