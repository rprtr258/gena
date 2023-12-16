package main

import (
	"math"

	"github.com/rprtr258/gena"
)

func PolygonAt(n int, p gena.V2, r float64) []gena.V2 {
	result := make([]gena.V2, n)
	for i := 0; i < n; i++ {
		result[i] = p + gena.Polar(r, math.Pi*(float64(i)*2/float64(n)-0.5))
	}
	return result
}

func star(n int) {
	points := PolygonAt(n, 512+512i, 400)
	dc := gena.NewContext(1024, 1024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		dc.LineToV2(points[(i*2)%n])
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gena.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()
	dc.SavePNG("star.png")
}
