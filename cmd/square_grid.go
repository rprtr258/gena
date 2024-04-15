package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a grid squares image.
func GirdSquares(im *image.RGBA, colorSchema []color.RGBA, step, rectSize int, decay float64, iters int) {
	dc := NewContextFromRGBA(im)

	for x := 0; x < im.Bounds().Dx(); x += step {
		for y := 0; y < im.Bounds().Dy(); y += step {
			cl := RandomItem(colorSchema)

			v0 := complex(float64(x), float64(y))
			s := float64(rectSize)

			theta := RandomInt(360) + 1
			for i := range Range(iters) {
				dc.Stack(func(ctx *Context) {
					dc.TransformAdd(Translate(v0 + Diag(step)/2))
					dc.TransformAdd(Rotate(Radians(float64(theta * i))))

					dc.TransformAdd(Scale(complex(s, s)))

					dc.LineTo(complex(-0.5, 0.5))
					dc.LineTo(complex(0.5, 0.5))
					dc.LineTo(complex(0.5, -0.5))
					dc.LineTo(complex(-0.5, -0.5))
					dc.LineTo(complex(-0.5, 0.5))

					dc.SetLineWidth(3)
					dc.SetColor(Tomato)
					dc.StrokePreserve()
					dc.SetColor(ColorRGBA255(cl, 255))
					dc.Fill()
				})
				s -= decay * float64(rectSize)
			}
		}
	}
}
