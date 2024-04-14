package main

import (
	"image/color"

	. "github.com/rprtr258/gena"
)

func mystar() {
	const n = 4

	dc := NewContext(complex(500, 500))

	dc.SetColor(color.Black)
	dc.Clear()

	points := PolygonAt(n, 250+250i, 100)
	for range Range(10) {
		dc.TransformAdd(Translate(complex(250, 250)))
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
