package gena

import (
	"image/color"
	"math"
)

type janus struct {
	fg    color.RGBA
	n     int
	decay float64
}

// NewJanus returns a janus object
func NewJanus(fg color.RGBA, n int, decay float64) *janus {
	return &janus{
		fg:    fg,
		n:     n,
		decay: decay,
	}
}

// Generative draws a janus image
// TODO not finished.
func (j *janus) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetColor(j.fg)
	s := 220.0
	r := 0.3

	for i := 0; i < j.n; i++ {
		// k := rand.Intn(len(c.ColorSchema()))
		k := i
		ctex.Push()
		ctex.Translate(c.Size() / 2)

		// theta += rand.Float64()*math.Pi/2
		theta := RandomFloat64(math.Pi/4, 3*math.Pi/4)
		x1, y1 := math.Cos(theta)*r, math.Sin(theta)*r
		x2, y2 := -x1, -y1

		noise := RandomFloat64(-math.Abs(y1), math.Abs(y1))
		y1 += noise
		y2 += noise

		// r -= r*j.decay
		s *= 0.836
		ctex.Scale(complex(s, s))
		// r *= 0.836
		ctex.DrawArc(complex(x1, y1), 1.0, math.Pi*3/2+theta, math.Pi*5/2+theta)
		ctex.SetColor(c.ColorSchema[k])
		ctex.Fill()
		ctex.DrawArc(complex(x2, y2), 1.0, math.Pi/2+theta, math.Pi*3/2+theta)
		ctex.SetColor(c.ColorSchema[k])
		ctex.Fill()
		ctex.Pop()
	}
}
