package main

import (
	"math"

	"github.com/rprtr258/gena"
)

func PolygonAt(n int, p gena.V2, r float64) []gena.V2 {
	result := make([]gena.V2, n)
	for i := range result {
		result[i] = p + gena.Polar(r, math.Pi*(float64(i)*2/float64(n)-0.5))
	}
	return result
}

func star(n int) {
	points := PolygonAt(n, 512+512i, 400)
	dc := gena.NewContext(1024, 1024)
	dc.SetColor(gena.ColorHex("#fff"))
	dc.Clear()
	for i := range gena.Range(n + 1) {
		dc.LineTo(points[(i*2)%n])
	}
	dc.SetColor(gena.ColorRGBA(0, 0.5, 0, 1))
	dc.SetFillRule(gena.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetColor(gena.ColorRGBA(0, 1, 0, 0.5))
	dc.SetLineWidth(16)
	dc.Stroke()
	gena.SavePNG("star.png", dc.Image())
}
