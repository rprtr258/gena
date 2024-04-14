package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

type swirlOpts struct {
	fg, bg color.RGBA
	// These are some math parameters.
	// http://paulbourke.net/fractals/peterdejong/
	a, b, c, d float64
	axis       V2
	iters      int
}

func (s *swirlOpts) swirlTransform(p V2) V2 {
	return complex(
		Sin(s.a*Y(p))-Cos(s.b*X(p)),
		Sin(s.c*X(p))-Cos(s.d*Y(p)),
	)
}

func (s *swirlOpts) removeNoisy(im *image.RGBA) {
	for i := 1; i < im.Bounds().Dx()-1; i++ {
		for j := 1; j < im.Bounds().Dy()-1; j++ {
			if im.At(i, j) == s.bg {
				continue
			}

			var n int
			if im.At(i+1, j) == s.bg {
				n += 1
			}
			if im.At(i+1, j+1) == s.bg {
				n += 1
			}
			if im.At(i, j+1) == s.bg {
				n += 1
			}

			if im.At(i-1, j) == s.bg {
				n += 1
			}

			if im.At(i-1, j+1) == s.bg {
				n += 1
			}

			if im.At(i-1, j-1) == s.bg {
				n += 1
			}

			if im.At(i+1, j-1) == s.bg {
				n += 1
			}

			if im.At(i, j-1) == s.bg {
				n += 1
			}

			if n > 5 {
				im.Set(i, j, s.bg)
			}
		}
	}
}

// Swirl draws a swirl image.
func Swirl(img *image.RGBA, fg, bg color.RGBA, a, b, c, d float64, axis V2, iters int) {
	s := &swirlOpts{
		fg:    fg,
		bg:    bg,
		a:     a,
		b:     b,
		c:     c,
		d:     d,
		axis:  axis,
		iters: iters,
	}
	dc := NewContextFromRGBA(img)
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
