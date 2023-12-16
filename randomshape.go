package gena

import (
	"math"
	"math/rand"
)

// RandomShape draws a random shape image.
func RandomShape(c Canvas, shapeNum int) {
	ctex := NewContextForRGBA(c.Img())

	ctex.Translate(c.Size() / 2)
	ctex.Rotate(RandomFloat64(-1, 1) * math.Pi * 0.25)
	ctex.Translate(-c.Size() / 2)

	for i := 0; i < shapeNum; i++ {
		v := Mul2(complex(
			RandomGaussian(0.5, 0.2),
			RandomGaussian(0.5, 0.2),
		), c.Size())

		w := RandomFloat64(0, float64(c.Width())/3)*RandomFloat64(0, rand.Float64()) + 5.0
		h := w + RandomFloat64(-1, 1)*3.0

		rnd := rand.Intn(4)
		theta := math.Pi * 2.0 * float64(rand.Intn(4)) / 4

		ctex.Push()
		ctex.Translate(v)
		ctex.Rotate(theta)
		ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
		switch rnd {
		case 0:
			ctex.DrawCircle(0, 0, w/2)
		case 1:
			ctex.DrawRectangle(0, complex(w/2, w/2))
		case 2:
			if rand.Float64() < 0.5 {
				ctex.DrawEllipse(0, complex(w/2, h/2))
			} else {
				ctex.DrawRectangle(0, complex(w, h))
			}
		case 3:
			ctex.DrawRectangle(0, complex(w*2, RandomFloat64(2, 10)))
		}
		ctex.Fill()
		ctex.Pop()
	}
}
