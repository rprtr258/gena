package main

import (
	"image"
	"math/cmplx"

	. "github.com/rprtr258/gena"
)

// GenFunc defines a func type used by julia set
type GenFunc func(V2) V2

// Julia draws Julia Set. [Julia Set](https://en.wikipedia.org/wiki/Julia_set) is a math concept. You can define your own formula in this package.
//   - fn: The custom julia set function.
//   - maxz: The maximum modulus length of a complex number.
//   - axis: The range for the X-Y coordination used to mapping the julia set number to the real pixel of the image. These should be positive. It only indicates the first quadrant range.
func Julia(
	dc *Context,
	palette Pattern1D,
	fn GenFunc,
	maxz float64,
	axis V2,
	iters int,
) {
	dc.SetColor(Azure)
	dc.Clear()
	im := dc.Image()
	for x := 0; x <= im.Bounds().Dx(); x++ {
		for y := 0; y <= im.Bounds().Dy(); y++ {
			f := complex(float64(x), float64(y))
			z := Div2(f, Size(im))*Coeff(2.0*Y(axis)) - Diag(Y(axis))

			var nit int
			for nit = 0; cmplx.Abs(z) <= maxz && nit < iters; nit++ {
				z = fn(z)
			}

			im.Set(x, y, palette(float64(nit)/float64(iters)))
		}
	}
}

func julia() *image.RGBA {
	im := NewContext(Diag(500))
	Julia(im, Viridis, func(z complex128) complex128 {
		return z*z + complex(-0.1, 0.651)
	}, 40, 1.5+1.5i, 800)
	return im.Image()
}
