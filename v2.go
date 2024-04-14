package gena

import "math/cmplx"

type V2 = complex128

// real(v)
func X(v V2) float64 {
	return real(v)
}

// imag(v)
func Y(v V2) float64 {
	return imag(v)
}

func Polar(r, theta float64) V2 {
	return cmplx.Exp(complex(0, theta)) * complex(r, 0)
}

func Rotation(angle float64) V2 {
	return Polar(1, angle)
}

func Magnitude(v V2) float64 {
	return cmplx.Abs(v)
}

func Coeff(x float64) V2 {
	return complex(x, 0)
}

func Diag(x float64) V2 {
	return complex(x, x)
}

func Mul2(v, w V2) V2 {
	return complex(X(v)*X(w), Y(v)*Y(w))
}

func Div2(v, w V2) V2 {
	return complex(real(v)/real(w), imag(v)/imag(w))
}

func SetMag(v V2, m float64) V2 {
	return v * Coeff(m/Magnitude(v))
}

func Normalized(v V2) V2 {
	return SetMag(v, 1)
}

func Plus(v V2, z float64) V2 {
	return v + Diag(z)
}

func Sub(v V2, z float64) V2 {
	return v - Diag(z)
}

// ToPolar converts points from cartesian coordinates to polar coordinates
func ToPolar(v V2) V2 {
	return complex(cmplx.Abs(v), cmplx.Phase(v))
}

// ToCartesian converts points from polar coordinates to cartesian coordinates
func ToCartesian(v V2) V2 {
	return cmplx.Exp(complex(0, imag(v))) * Coeff(X(v))
}

// ToPixel converts cartesian coordinates to actual pixels coordinates in an image
func ToPixel(v, axis, size V2) V2 {
	return Mul2(size/Coeff(2), Plus(Div2(v, axis), 1))
}

// Dist returns the Euclidean distance between two point on 2D dimension.
func Dist(p1, p2 V2) float64 {
	return Magnitude(p1 - p2)
}

func Dot(p1, p2 V2) float64 {
	return real(p1 * cmplx.Conj(p2))
}
