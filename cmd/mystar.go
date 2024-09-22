package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

func mystar() *image.RGBA {
	const n = 4
	const m = 9

	dc := NewContext(Diag(500))
	dc.SetColor(color.Black)
	dc.Clear()
	dc.TransformAdd(Translate(Diag(250)))
	dc.TransformAdd(Scale(Diag(100)))

	points := Polygon(n)
	for range Range(m) {
		dc.TransformAdd(Rotate(PI / m))

		for _, point := range points {
			dc.LineTo(point)
		}
		dc.LineTo(points[0])

		dc.SetColor(color.White)
		dc.SetLineWidth(0.5)
		dc.Stroke()
	}

	return dc.Image()
}
