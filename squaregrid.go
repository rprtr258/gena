package gena

import (
	"math/rand"
)

// Generative draws a grid squares image.
func GirdSquares(c Canvas, step, rectSize int, decay float64, iters int) {
	ctex := NewContextForRGBA(c.Img())

	for x := 0; x < c.Width(); x += step {
		for y := 0; y < c.Height(); y += step {
			cl := c.ColorSchema[rand.Intn(len(c.ColorSchema))]

			v0 := complex(float64(x), float64(y))
			s := float64(rectSize)

			theta := rand.Intn(360) + 1
			for i := 0; i < iters; i++ {
				ctex.Push()

				ctex.Translate(Plus(v0, float64(step)/2))
				ctex.Rotate(Radians(float64(theta * i)))

				ctex.Scale(complex(s, s))

				ctex.LineTo(-0.5, 0.5)
				ctex.LineTo(0.5, 0.5)
				ctex.LineTo(0.5, -0.5)
				ctex.LineTo(-0.5, -0.5)
				ctex.LineTo(-0.5, 0.5)

				ctex.SetLineWidth(c.LineWidth)
				ctex.SetColor(c.LineColor)
				ctex.StrokePreserve()
				ctex.SetRGBA255(cl, c.Alpha)
				ctex.Fill()
				ctex.Pop()
				s -= decay * float64(rectSize)
			}
		}
	}
}
