package main

import (
	"image/color"
	"math"

	"github.com/rprtr258/gena"
)

func mystar() {
	const n = 4

	dc := gena.NewContext(500, 500)

	dc.SetColor(color.Black)
	dc.Clear()

	points := PolygonAt(n, 250+250i, 100)
	for range gena.Range(10) {
		dc.RelativeTo(250+250i, func(dc *gena.Context) {
			dc.Rotate(math.Pi / 9)
		})

		for i := range gena.Range(n) {
			dc.LineToV2(points[i])
		}
		dc.LineToV2(points[0])

		dc.SetColor(color.White)
		dc.SetLineWidth(0.5)
		dc.Stroke()
	}

	gena.SavePNG("mystar.png", dc.Image())
}
