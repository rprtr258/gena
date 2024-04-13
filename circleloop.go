package gena

import (
	"image"
	"image/color"
	"math"
)

// CircleLoop draws a Circle Loop images.
func CircleLoop(c *image.RGBA, lineWidth float64, lineColor color.RGBA, alpha int, radius float64, iters int) {
	FillBackground(c, Black)
	dc := NewContextForRGBA(c)
	for i := range Range(iters) {
		theta := math.Pi / 2 * float64(i)
		dc.Stack(func(dc *Context) {
			dc.Translate(Size(c) / 2)
			v := Mul(complex(
				math.Cos(Radians(theta)),
				math.Sin(Radians(theta*2)),
			), radius)

			dc.SetLineWidth(lineWidth)
			dc.SetColor(lineColor)
			dc.SetRGBA255(lineColor, alpha)
			dc.DrawCircleV2(v, radius+math.Sin(theta*1.5)*float64(i)/2)
			dc.Stroke()
		})
	}
}
