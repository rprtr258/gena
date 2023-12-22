package gena

import (
	"image/color"
	"math"
	"math/rand"
)

// Generative draws a color circle images.
func ColorCircle(c Canvas, colorSchema []color.RGBA, circleNum int) {
	dc := NewContextForRGBA(c.Img())

	for i := 0; i < circleNum; i++ {
		v := Mul2(complex(
			RandomFloat64(-0.1, 1.1),
			RandomFloat64(-0.1, 1.1),
		), c.Size())
		s := RandomFloat64(0, RandomFloat64(0, float64(c.Width/2))) + 10

		rnd := rand.Intn(3)
		if rnd == 2 {
			rnd = rand.Intn(3)
		}
		switch rnd {
		case 0:
			n := RandomRangeInt(4, 30)
			cs := RandomFloat64(2, 8)
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			dc.Stack(func(ctx *Context) {
				dc.Translate(v)
				for a := 0.0; a < math.Pi*2.0; a += math.Pi * 2.0 / float64(n) {
					dc.DrawEllipse(Polar(s, a)/2, complex(cs/2, cs/2))
					dc.Fill()
				}
			})
		case 1:
			dc.SetLineWidth(RandomFloat64(0, 1))
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			dc.DrawCircleV2(v, RandomFloat64(0, s)/2)
			dc.Stroke()
		case 2:
			cl := colorSchema[rand.Intn(len(colorSchema))]
			dc.SetLineWidth(1.0)
			sx := s * RandomFloat64(0.1, 0.55)
			for j := 0.0001; j < sx; j++ {
				dd := s + j*2.0
				alpha := Constrain(int(255*sx/j), 0, 255)

				// alpha := RandomRangeInt(30, 150)
				cl.A = uint8(alpha)
				dc.SetColor(cl)

				for i := 0; i < 200; i++ {
					theta := RandomFloat64(0, math.Pi*2)
					dc.DrawPoint(v+Polar(dd*0.3, theta), 0.6)
					dc.Stroke()
				}
			}
		}
	}
}
