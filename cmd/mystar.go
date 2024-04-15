package main

import (
	"image/color"

	. "github.com/rprtr258/gena"
)

func mystar() {
	const n = 4

	dc := NewContext(Diag(500))
	dc.SetColor(color.Black)
	dc.Clear()
	dc.TransformAdd(Translate(Diag(250)))

	points := PolygonAt(n, 0, 100)
	for range Range(10) {
		dc.TransformAdd(Rotate(PI / 9))

		for _, point := range points {
			dc.LineTo(point)
		}
		dc.LineTo(points[0])

		dc.SetColor(color.White)
		dc.SetLineWidth(0.5)
		dc.Stroke()
	}

	SavePNG("mystar.png", dc.Image())
}
