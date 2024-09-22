package gena

import "math/cmplx"

type V2 complex128

func P[X, Y interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}](x X, y Y) V2 {
	return V2(complex(float64(x), float64(y)))
}

// real(v)
func (v V2) X() float64 {
	return real(v)
}

// imag(v)
func (v V2) Y() float64 {
	return imag(v)
}

func (v V2) XY() (x, y float64) {
	return v.X(), v.Y()
}

func Polar(r, theta float64) V2 {
	angle := V2(cmplx.Exp(complex(0, theta)))
	return angle * Coeff(r)
}

func Rotation(angle float64) V2 {
	return Polar(1, angle)
}

func (v V2) Magnitude() float64 {
	return cmplx.Abs(complex128(v))
}

func Coeff(x float64) V2 {
	return V2(complex(x, 0))
}

func Diag[N interface{ ~float64 | ~int }](x N) V2 {
	xx := float64(x)
	return P(xx, xx)
}

func Mul2(v, w V2) V2 {
	return V2(complex(v.X()*w.X(), v.Y()*w.Y()))
}

func Div2(v, w V2) V2 {
	return V2(complex(real(v)/real(w), imag(v)/imag(w)))
}

func (v V2) SetMag(m float64) V2 {
	return v * Coeff(m/v.Magnitude())
}

func (v V2) Normalized() V2 {
	return v.SetMag(1)
}

// ToPolar converts points from cartesian coordinates to polar coordinates
func (v V2) ToPolar() V2 {
	return V2(complex(cmplx.Abs(complex128(v)), cmplx.Phase(complex128(v))))
}

// ToCartesian converts points from polar coordinates to cartesian coordinates
func (v V2) ToCartesian() V2 {
	return V2(cmplx.Exp(complex(0, imag(v)))) * Coeff(v.X())
}

// ToPixel converts cartesian coordinates to actual pixels coordinates in an image
func ToPixel(v, axis, size V2) V2 {
	return Mul2(size/Coeff(2), Div2(v, axis)+Diag(1))
}

// Dist returns the Euclidean distance between two point on 2D dimension.
func Dist(p1, p2 V2) float64 {
	return (p1 - p2).Magnitude()
}

func Dot(p1, p2 V2) float64 {
	return real(p1 * V2(cmplx.Conj(complex128(p2))))
}

func LerpV2(a, b V2, t float64) V2 {
	return a + (b-a)*Coeff(t)
}
