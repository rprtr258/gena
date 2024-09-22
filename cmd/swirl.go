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
	return P(
		Sin(s.a*p.Y())-Cos(s.b*p.X()),
		Sin(s.c*p.X())-Cos(s.d*p.Y()),
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
func Swirl(dc *Context, fg, bg color.RGBA, a, b, c, d float64, axis V2, iters int) {
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
	im := dc.Image()
	dc.SetColor(bg)
	dc.Clear()
	dc.SetLineWidth(3)
	start := P(1.0, 1.0)

	for range Range(s.iters) {
		next := s.swirlTransform(start)
		xy := ToPixel(next, s.axis, Size(im))
		im.Set(int(xy.X()), int(xy.Y()), s.fg)
		start = next
	}

	s.removeNoisy(im)
	s.removeNoisy(im)
}

func swirl() *image.RGBA {
	dc := NewContext(Diag(1600))
	Swirl(
		dc,
		color.RGBA{113, 3, 0, 140}, Azure,
		0.970, -1.899, 1.381, -1.506,
		2.4+2.4i,
		4000000,
	)
	return dc.Image()
}
