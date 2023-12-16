package gena

import "math/cmplx"

// GenFunc defines a func type used by julia set.
type GenFunc func(complex128) complex128

// Generative draws a julia set.
// Julia is to draw a Julia Set. [Julia Set](https://en.wikipedia.org/wiki/Julia_set) is a math concept. You can define your own formula in this package.
//   - fn: The custom julia set function.
//   - maxz: The maximum modulus length of a complex number.
//   - axis: The range for the X-Y coordination used to mapping the julia set number to the real pixel of the image. These should be positive. It only indicates the first quadrant range.
func Julia(c Canvas, fn GenFunc, maxz float64, axis V2, iters int) {
	n := uint8(min(len(c.ColorSchema), 255))

	for i := 0; i <= c.Width; i++ {
		for k := 0; k <= c.Height; k++ {
			z := complex(
				float64(i)/float64(c.Width)*2.0*Y(axis)-Y(axis),
				float64(k)/float64(c.Height)*2.0*Y(axis)-Y(axis),
			)

			var nit int
			for nit = 0; cmplx.Abs(z) <= maxz && nit < iters; nit++ {
				z = fn(z)
			}

			c.Img().Set(i, k, c.ColorSchema[uint8(nit*255/iters)%n])
		}
	}
}
