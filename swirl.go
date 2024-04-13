package gena

import (
	"image"
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

func (s *swirl) swirlTransform(p V2) V2 {
	return complex(
		math.Sin(s.a*Y(p))-math.Cos(s.b*X(p)),
		math.Sin(s.c*X(p))-math.Cos(s.d*Y(p)),
	)
}

func (s *swirl) removeNoisy(c *image.RGBA) {
	for i := 1; i < c.Bounds().Dx()-1; i++ {
		for j := 1; j < c.Bounds().Dy()-1; j++ {
			if c.At(i, j) == s.bg {
				continue
			}

			var n int
			if c.At(i+1, j) == s.bg {
				n += 1
			}
			if c.At(i+1, j+1) == s.bg {
				n += 1
			}
			if c.At(i, j+1) == s.bg {
				n += 1
			}

			if c.At(i-1, j) == s.bg {
				n += 1
			}

			if c.At(i-1, j+1) == s.bg {
				n += 1
			}

			if c.At(i-1, j-1) == s.bg {
				n += 1
			}

			if c.At(i+1, j-1) == s.bg {
				n += 1
			}

			if c.At(i, j-1) == s.bg {
				n += 1
			}

			if n > 5 {
				c.Set(i, j, s.bg)
			}
		}
	}
}

// Swirl draws a swirl image.
func Swirl(img *image.RGBA, fg, bg color.RGBA, a, b, c, d float64, axis V2, iters int) {
	s := &swirl{
		fg:    fg,
		bg:    bg,
		a:     a,
		b:     b,
		c:     c,
		d:     d,
		axis:  axis,
		iters: iters,
	}
	dc := NewContextForRGBA(img)
	dc.SetLineWidth(3)
	start := complex(1.0, 1.0)

	for range Range(s.iters) {
		next := s.swirlTransform(start)
		xy := ToPixel(next, s.axis, Size(img))
		img.Set(int(X(xy)), int(Y(xy)), s.fg)
		start = next
	}

	s.removeNoisy(img)
	s.removeNoisy(img)
}
