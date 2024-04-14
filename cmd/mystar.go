package main

import (
	"image/color"
	"math"

	. "github.com/rprtr258/gena"
)

func mystar() {
	const n = 4

	dc := NewContext(complex(500, 500))

	dc.SetColor(color.Black)
	dc.Clear()

	points := PolygonAt(n, 250+250i, 100)
	for range Range(10) {
		dc.RelativeTo(250+250i, func(dc *Context) {
			dc.TransformAdd(Rotate(math.Pi / 9))
		})

		for i := range Range(n) {
			dc.LineTo(points[i])
		}
		dc.LineTo(points[0])

		dc.SetColor(color.White)
		dc.SetLineWidth(0.5)
		dc.Stroke()
	}

	SavePNG("mystar.png", dc.Image())
}
