package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

func mystar() *image.RGBA {
	const numPoints = 4
	const rotations = 9
	size := Diag(500)
	const radius = 100

	dc := NewContext(size)
	dc.SetColor(color.Black)
	dc.Clear()
	dc.TransformAdd(Translate(size / 2))
	dc.TransformAdd(Scale(Diag(radius)))

	points := Polygon(numPoints)
	for range Range(rotations) {
		dc.TransformAdd(Rotate(PI / rotations))

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
