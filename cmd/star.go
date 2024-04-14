package main

import (
	"math"

	. "github.com/rprtr258/gena"
)

func PolygonAt(n int, p V2, r float64) []V2 {
	result := make([]V2, n)
	for i := range result {
		result[i] = p + Polar(r, math.Pi*(float64(i)*2/float64(n)-0.5))
	}
	return result
}

func star(n int) {
	points := PolygonAt(n, 512+512i, 400)
	dc := NewContext(complex(1024, 1024))
	dc.SetColor(ColorHex("#fff"))
	dc.Clear()
	for i := range Range(n + 1) {
		dc.LineTo(points[(i*2)%n])
	}
	dc.SetColor(ColorRGBA(0, 0.5, 0, 1))
	dc.SetFillRule(FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetColor(ColorRGBA(0, 1, 0, 0.5))
	dc.SetLineWidth(16)
	dc.Stroke()
	SavePNG("star.png", dc.Image())
}
