package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// CircleLoop draws a Circle Loop images.
func CircleLoop(
	im *image.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	alpha int,
	radius float64,
	iters int,
) {
	FillBackground(im, Black)
	dc := NewContextFromRGBA(im)
	dc.TransformAdd(Translate(Size(im) / 2))
	for i := range Range(iters) {
		theta := PI / 2 * float64(i)
		dc.Stack(func(dc *Context) {
			v := complex(
				Cos(Radians(theta)),
				Sin(Radians(theta*2)),
			) * Coeff(radius)

			dc.SetLineWidth(lineWidth)
			dc.SetColor(ColorRGBA255(lineColor, alpha))
			dc.DrawCircle(v, (radius+Sin(theta*1.5))/2)
			dc.Stroke()
		})
	}
}
