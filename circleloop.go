package gena

import "math"

type circleLoop struct {
	radius float64
	iters  int
}

func NewCircleLoop(radius float64, iters int) *circleLoop {
	return &circleLoop{
		radius: radius,
		iters:  iters,
	}
}

// Generative draws a Circle Loop images.
func (cl *circleLoop) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())

	r := cl.radius
	var theta float64 = 0
	for i := 0; i < cl.iters; i++ {
		ctex.Push()
		ctex.Translate(c.Size() / 2)
		v := Mul(complex(
			math.Cos(Radians(theta)),
			math.Sin(Radians(theta*2)),
		), cl.radius)

		ctex.SetLineWidth(c.LineWidth)
		ctex.SetColor(c.LineColor)
		ctex.SetRGBA255(c.LineColor, c.Alpha)
		ctex.DrawCircleV2(v, r/2)
		ctex.Stroke()
		ctex.Pop()
		r += math.Cos((theta))*math.Sin((theta/2)) + math.Sin((theta))*math.Cos((theta/2))
		theta += math.Pi / 2
	}
}
