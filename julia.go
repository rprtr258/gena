package gena

import "math/cmplx"

// GenFunc defines a func type used by julia set.
type GenFunc func(complex128) complex128

type julia struct {
	fn    GenFunc
	maxz  float64
	axis  V2
	iters int
}

func NewJulia(formula GenFunc, maxz float64, axis V2, iters int) *julia {
	return &julia{
		fn:    formula,
		maxz:  maxz,
		axis:  axis,
		iters: iters,
	}
}

// Generative draws a julia set.
func (j *julia) Generative(c Canvas) {
	n := uint8(min(len(c.ColorSchema), 255))

	for i := 0; i <= c.Width(); i++ {
		for k := 0; k <= c.Height(); k++ {
			z := complex(
				float64(i)/float64(c.Width())*2.0*Y(j.axis)-Y(j.axis),
				float64(k)/float64(c.Height())*2.0*Y(j.axis)-Y(j.axis),
			)

			var nit int
			for nit = 0; cmplx.Abs(z) <= j.maxz && nit < j.iters; nit++ {
				z = j.fn(z)
			}

			c.Img().Set(i, k, c.ColorSchema[uint8(nit*255/j.iters)%n])
		}
	}
}
