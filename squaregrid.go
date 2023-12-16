package gena

import "math/rand"

// Generative draws a grid squares image.
func GirdSquares(c Canvas, step, rectSize int, decay float64, iters int) {
	dc := NewContextForRGBA(c.Img())

	for x := 0; x < c.Width; x += step {
		for y := 0; y < c.Height; y += step {
			cl := c.ColorSchema[rand.Intn(len(c.ColorSchema))]

			v0 := complex(float64(x), float64(y))
			s := float64(rectSize)

			theta := rand.Intn(360) + 1
			for i := 0; i < iters; i++ {
				dc.Stack(func(ctx *Context) {
					dc.Translate(Plus(v0, float64(step)/2))
					dc.Rotate(Radians(float64(theta * i)))

					dc.Scale(complex(s, s))

					dc.LineTo(-0.5, 0.5)
					dc.LineTo(0.5, 0.5)
					dc.LineTo(0.5, -0.5)
					dc.LineTo(-0.5, -0.5)
					dc.LineTo(-0.5, 0.5)

					dc.SetLineWidth(3)
					dc.SetColor(Tomato)
					dc.StrokePreserve()
					dc.SetRGBA255(cl, 255)
					dc.Fill()
				})
				s -= decay * float64(rectSize)
			}
		}
	}
}
