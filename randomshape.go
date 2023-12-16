package gena

import (
	"math"
	"math/rand"
)

// RandomShape would draw images with random shapes.
// The whole image would rotate with some degree.
//   - shapeNum: It indicates how many shapes you want to draw.
func RandomShape(c Canvas, shapeNum int) {
	dc := NewContextForRGBA(c.Img())

	dc.Translate(c.Size() / 2)
	dc.Rotate(RandomFloat64(-1, 1) * math.Pi * 0.25)
	dc.Translate(-c.Size() / 2)

	for i := 0; i < shapeNum; i++ {
		v := Mul2(complex(
			RandomGaussian(0.5, 0.2),
			RandomGaussian(0.5, 0.2),
		), c.Size())

		w := RandomFloat64(0, float64(c.Width)/3)*RandomFloat64(0, rand.Float64()) + 5.0
		h := w + RandomFloat64(-1, 1)*3.0

		rnd := rand.Intn(4)
		theta := math.Pi * 2.0 * float64(rand.Intn(4)) / 4

		dc.Stack(func(ctx *Context) {
			dc.Translate(v)
			dc.Rotate(theta)
			dc.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
			switch rnd {
			case 0:
				dc.DrawCircle(0, 0, w/2)
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
