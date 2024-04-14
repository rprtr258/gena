package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

// Generative draws a color circle images.
func ColorCircle(im *image.RGBA, colorSchema []color.RGBA, n int) {
	dc := NewContextFromRGBA(im)

	for range Range(n) {
		v := Mul2(complex(
			RandomF64(-0.1, 1.1),
			RandomF64(-0.1, 1.1),
		), Size(im))
		s := RandomF64(0, RandomF64(0, float64(im.Bounds().Dx()/2))) + 10

		rnd := rand.Intn(3)
		if rnd == 2 {
			rnd = rand.Intn(3)
		}
		switch rnd {
		case 0:
			n := RandomIntN(4, 30)
			cs := RandomF64(2, 8)
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			dc.Stack(func(ctx *Context) {
				dc.TransformAdd(Translate(v))
				for a := 0.0; a < math.Pi*2.0; a += math.Pi * 2.0 / float64(n) {
					dc.DrawEllipse(Polar(s, a)/2, complex(cs/2, cs/2))
					dc.Fill()
				}
			})
		case 1:
			dc.SetLineWidth(RandomF64(0, 1))
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			dc.DrawCircle(v, RandomF64(0, s)/2)
			dc.Stroke()
		case 2:
			cl := colorSchema[rand.Intn(len(colorSchema))]
			dc.SetLineWidth(1.0)
			sx := s * RandomF64(0.1, 0.55)
			for j := 0.0001; j < sx; j++ {
				dd := s + j*2.0
				alpha := Clamp(int(255*sx/j), 0, 255)

				cl.A = uint8(alpha)
				dc.SetColor(cl)

				for range Range(200) {
					theta := RandomF64(0, math.Pi*2)
					dc.DrawPoint(v+Polar(dd*0.3, theta), 0.6)
					dc.Stroke()
				}
			}
		}
	}
}
