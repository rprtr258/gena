package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// CircleLine draws a cirle line image.
func CircleLine(
	dc *Context,
	lineWidth float64,
	lineColor color.RGBA,
	step float64,
	lineNum int,
	radius float64,
	axis V2,
) {
	im := dc.Image()

	var points []V2
	for theta := -PI; theta <= PI; theta += step {
		points = append(points, ToPixel(Polar(radius, theta), axis, Size(im)))
	}

	dc.SetColor(Tan)
	dc.Clear()

	dc.SetLineWidth(lineWidth)
	dc.SetColor(lineColor)
	for range Range(lineNum) {
		dc.MoveTo(RandomItem(points))
		dc.LineTo(RandomItem(points))
		dc.Stroke()
	}
}

func circleLine() *image.RGBA {
	dc := NewContext(Diag(600))
	CircleLine(dc, 1, LightPink, 0.02, 600, 1.5, complex(2, 2))
	return dc.Image()
}
