package gena

import (
	"image/color"
	"math"
)

// CircleLoop draws a Circle Loop images.
func CircleLoop(c Canvas, lineWidth float64, lineColor color.RGBA, alpha int, radius float64, iters int) {
	ctex := NewContextForRGBA(c.Img())

	r := radius
	var theta float64 = 0
	for i := 0; i < iters; i++ {
		ctex.Stack(func(dc *Context) {
			ctex.Translate(c.Size() / 2)
			v := Mul(complex(
				math.Cos(Radians(theta)),
				math.Sin(Radians(theta*2)),
			), radius)

			ctex.SetLineWidth(lineWidth)
			ctex.SetColor(lineColor)
			ctex.SetRGBA255(lineColor, alpha)
			ctex.DrawCircleV2(v, r/2)
			ctex.Stroke()
		})
		r += math.Sin(theta * 1.5)
		theta += math.Pi / 2
	}
}
