package gena

import (
	"image/color"
	"math"
)

type swirl struct {
	fg, bg color.RGBA
	// These are some math parameters.
	// http://paulbourke.net/fractals/peterdejong/
	a, b, c, d float64
	axis       V2
	iters      int
}

func NewSwirl(fg, bg color.RGBA, a, b, c, d float64, axis V2, iters int) *swirl {
	return &swirl{
		fg:    fg,
		bg:    bg,
		a:     a,
		b:     b,
		c:     c,
		d:     d,
		axis:  axis,
		iters: iters,
	}
}

// Generative draws a swirl image.
func (s *swirl) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(c.LineWidth)
	start := complex(1.0, 1.0)

	for i := 0; i < s.iters; i++ {
		next := s.swirlTransform(start)
		xy := ToPixel(next, s.axis, c.Size())
		c.Img().Set(int(X(xy)), int(Y(xy)), s.fg)
		start = next
	}

	s.removeNoisy(c)
	s.removeNoisy(c)
}

func (s *swirl) swirlTransform(p V2) V2 {
	return complex(
		math.Sin(s.a*Y(p))-math.Cos(s.b*X(p)),
		math.Sin(s.c*X(p))-math.Cos(s.d*Y(p)),
	)
}

func (s *swirl) removeNoisy(c Canvas) {
	for i := 1; i < c.Width()-1; i++ {
		for j := 1; j < c.Height()-1; j++ {
			if c.Img().At(i, j) == s.bg {
				continue
			}

			var n int
			if c.Img().At(i+1, j) == s.bg {
				n += 1
			}
			if c.Img().At(i+1, j+1) == s.bg {
				n += 1
			}
			if c.Img().At(i, j+1) == s.bg {
				n += 1
			}

			if c.Img().At(i-1, j) == s.bg {
				n += 1
			}

			if c.Img().At(i-1, j+1) == s.bg {
				n += 1
			}

			if c.Img().At(i-1, j-1) == s.bg {
				n += 1
			}

			if c.Img().At(i+1, j-1) == s.bg {
				n += 1
			}

			if c.Img().At(i, j-1) == s.bg {
				n += 1
			}

			if n > 5 {
				c.Img().Set(i, j, s.bg)
			}

		}
	}
}
