package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

// Generative draws a grid squares image
func GirdSquares(dc *Context, palette Pattern1D, step, rectSize int, decay float64, iters int) {
	im := dc.Image()

	for x := 0; x < im.Bounds().Dx(); x += step {
		for y := 0; y < im.Bounds().Dy(); y += step {
			cl := palette.Random()

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
					dc.SetColor(cl)
					dc.Fill()
				})
				s -= decay * float64(rectSize)
			}
		}
	}
}

func gridsquares() *image.RGBA {
	dc := NewContext(Diag(600))
	GirdSquares(dc, DarkPink, 24, 10, 0.2, 20)
	return dc.Image()
}
