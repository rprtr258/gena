package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func star(n int) *image.RGBA {
	dc := NewContext(complex(1024, 1024))
	dc.SetColor(ColorHex("#fff"))
	dc.Clear()
	dc.TransformAdd(Translate(complex(512, 512)))
	dc.TransformAdd(Scale(Diag(400)))
	points := Polygon(n)
	for i := range Range(n + 1) {
		dc.LineTo(points[(i*2)%n])
	}
	dc.SetColor(ColorRGBA(0, 0.5, 0, 1))
	dc.SetFillRule(FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetColor(ColorRGBA(0, 1, 0, 0.5))
	dc.SetLineWidth(16)
	dc.Stroke()
	return dc.Image()
}
