package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

// RandomShape would draw images with random shapes.
// The whole image would rotate with some degree.
//   - shapeNum: It indicates how many shapes you want to draw.
func RandomShape(c *image.RGBA, colorSchema []color.RGBA, shapeNum int) {
	dc := NewContextForRGBA(c)

	dc.Translate(Size(c) / 2)
	dc.Rotate(RandomFloat64(-1, 1) * math.Pi * 0.25)
	dc.Translate(-Size(c) / 2)

	for range Range(shapeNum) {
		v := Mul2(complex(
			RandomGaussian(0.5, 0.2),
			RandomGaussian(0.5, 0.2),
		), Size(c))

		w := RandomFloat64(0, float64(c.Bounds().Dx())/3)*RandomFloat64(0, rand.Float64()) + 5.0
		h := w + RandomFloat64(-1, 1)*3.0

		rnd := rand.Intn(4)
		theta := math.Pi * 2.0 * float64(rand.Intn(4)) / 4

		dc.Stack(func(ctx *Context) {
			dc.Translate(v)
			dc.Rotate(theta)
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			switch rnd {
			case 0:
				dc.DrawCircle(0, w/2)
			case 1:
				dc.DrawRectangle(0, complex(w/2, w/2))
			case 2:
				if rand.Float64() < 0.5 {
					dc.DrawEllipse(0, complex(w/2, h/2))
				} else {
					dc.DrawRectangle(0, complex(w, h))
				}
			case 3:
				dc.DrawRectangle(0, complex(w*2, RandomFloat64(2, 10)))
			}
			dc.Fill()
		})
	}
}
