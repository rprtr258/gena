package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// CircleLoop draws a Circle Loop images.
func CircleLoop(
	dc *Context,
	lineWidth float64,
	lineColor color.RGBA,
	alpha int,
	radius float64,
	iters int,
) {
	dc.SetColor(Black)
	dc.Clear()

	dc.TransformAdd(Translate(Size(dc.Image()) / 2))
	for i := range Range(iters) {
		theta := PI / 2 * float64(i)
		dc.Stack(func(dc *Context) {
			v := P(
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

func circleLoop() *image.RGBA {
	dc := NewContext(Diag(500))
	CircleLoop(dc, 1, Orange, 30, 100, 1000)
	return dc.Image()
}
