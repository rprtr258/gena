package main

import (
	"math"
	"math/rand"

	"github.com/rprtr258/gena"
)

func openfill() {
	dc := gena.NewContext(1000, 1000)
	for j := range 10 {
		for i := range 10 {
			v := complex(float64(i), float64(j))*100 + 50
			a1 := rand.Float64() * 2 * math.Pi
			a2 := a1 + rand.Float64()*math.Pi + math.Pi/2
			dc.DrawArc(v, 40, a1, a2)
			// dc.ClosePath()
		}
	}
	dc.SetRGB(0, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(8)
	dc.StrokePreserve()
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(4)
	dc.StrokePreserve()
	gena.SavePNG("openfill.png", dc.Image())
}
