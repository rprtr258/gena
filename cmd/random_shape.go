package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// RandomShape would draw images with random shapes.
// The whole image would rotate with some degree.
//   - n: It indicates how many shapes you want to draw.
func RandomShape(im *image.RGBA, colorSchema []color.RGBA, n int) {
	dc := NewContextFromRGBA(im)

	dc.TransformAdd(Translate(Size(im) / 2))
	dc.TransformAdd(Rotate(RandomF64(-1, 1) * PI * 0.25))
	dc.TransformAdd(Translate(-Size(im) / 2))

	for range Range(n) {
		v := Mul2(complex(
			RandomGaussian(0.5, 0.2),
			RandomGaussian(0.5, 0.2),
		), Size(im))

		w := RandomF64(0, float64(im.Bounds().Dx())/3)*RandomF64(0, Random()) + 5.0
		h := w + RandomF64(-1, 1)*3.0

		rnd := RandomInt(4)
		theta := PI * 2.0 * float64(RandomInt(4)) / 4

		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(v))
			dc.TransformAdd(Rotate(theta))
			dc.SetColor(RandomItem(colorSchema))
			switch rnd {
			case 0:
				dc.DrawCircle(0, w/2)
			case 1:
				dc.DrawRectangle(0, complex(w/2, w/2))
			case 2:
				if Random() < 0.5 {
					dc.DrawEllipse(0, complex(w/2, h/2))
				} else {
					dc.DrawRectangle(0, complex(w, h))
				}
			case 3:
				dc.DrawRectangle(0, complex(w*2, RandomF64(2, 10)))
			}
			dc.Fill()
		})
	}
}
