package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// CircleLine draws a cirle line image.
func CircleLine(
	im *image.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	step float64,
	lineNum int,
	radius float64,
	axis V2,
) {
	dc := NewContextFromRGBA(im)
	dc.SetLineWidth(lineWidth)
	dc.SetColor(lineColor)

	var points []V2
	for theta := -PI; theta <= PI; theta += step {
		points = append(points, ToPixel(Polar(radius, theta), axis, Size(im)))
	}

	for range Range(lineNum) {
		dc.MoveTo(RandomItem(points))
		dc.LineTo(RandomItem(points))
		dc.Stroke()
	}
}
