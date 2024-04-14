package main

import (
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

func openfill() {
	dc := NewContext(complex(1000, 1000))
	for j := range Range(10) {
		for i := range Range(10) {
			v := complex(float64(i), float64(j))*100 + 50
			a1 := rand.Float64() * 2 * math.Pi
			a2 := a1 + rand.Float64()*math.Pi + math.Pi/2
			dc.DrawArc(v, 40, a1, a2)
			// dc.ClosePath()
		}
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
