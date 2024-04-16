package gena

import (
	"image"
	"image/color"

	"github.com/golang/freetype/raster"
)

type Pattern2D func(f V2) color.Color

func PatternSolid(c color.Color) Pattern2D {
	return func(V2) color.Color {
		return c
	}
}

type RepeatOp int

const (
	RepeatBoth RepeatOp = iota
	RepeatX
	RepeatY
	RepeatNone
)

func PatternSurface(im image.Image, op RepeatOp) Pattern2D {
	return func(f V2) color.Color {
		b := im.Bounds()
		x, y := int(X(f)), int(Y(f))
		switch op {
		case RepeatX:
			if y >= b.Dy() {
				return color.Transparent
			}
		case RepeatY:
			if x >= b.Dx() {
				return color.Transparent
			}
		case RepeatNone:
			if x >= b.Dx() || y >= b.Dy() {
				return color.Transparent
			}
		}
		return im.At(x%b.Dx()+b.Min.X, y%b.Dy()+b.Min.Y)
	}
}

func newPatternPainter(im *image.RGBA, mask *image.Alpha, pattern Pattern2D) raster.Painter {
	return raster.PainterFunc(func(ss []raster.Span, done bool) {
		b := im.Bounds()
		for _, s := range ss {
			if s.Y < b.Min.Y {
				continue
			}
			if s.Y >= b.Max.Y {
				return
			}
			if s.X0 < b.Min.X {
				s.X0 = b.Min.X
			}
			if s.X1 > b.Max.X {
				s.X1 = b.Max.X
			}
			if s.X0 >= s.X1 {
				continue
			}
			const m = 1<<16 - 1
			y := s.Y - im.Rect.Min.Y
			x0 := s.X0 - im.Rect.Min.X
			// RGBAPainter.Paint() in $GOPATH/src/github.com/golang/freetype/raster/paint.go
			i0 := (s.Y-im.Rect.Min.Y)*im.Stride + (s.X0-im.Rect.Min.X)*4
			i1 := i0 + (s.X1-s.X0)*4
			for i, x := i0, x0; i < i1; i, x = i+4, x+1 {
				ma := s.Alpha
				if mask != nil {
					ma = ma * uint32(mask.AlphaAt(x, y).A) / 255
					if ma == 0 {
						continue
					}
				}
				c := pattern(complex(float64(x), float64(y)))
				cr, cg, cb, ca := c.RGBA()
				dr := uint32(im.Pix[i+0])
				dg := uint32(im.Pix[i+1])
				db := uint32(im.Pix[i+2])
				da := uint32(im.Pix[i+3])
				a := (m - (ca * ma / m)) * 0x101
				im.Pix[i+0] = uint8((dr*a + cr*ma) / m >> 8)
				im.Pix[i+1] = uint8((dg*a + cg*ma) / m >> 8)
				im.Pix[i+2] = uint8((db*a + cb*ma) / m >> 8)
				im.Pix[i+3] = uint8((da*a + ca*ma) / m >> 8)
			}
		}
	})
}
