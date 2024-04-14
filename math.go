package gena

import (
	"cmp"
	"math"

	"golang.org/x/exp/constraints"
)

const PI = math.Pi

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Abs[T interface {
	constraints.Float | constraints.Integer
}](x T) T {
	return T(math.Abs(float64(x)))
}

func Floor[T interface {
	constraints.Float | constraints.Integer
}](x T) T {
	return T(math.Floor(float64(x)))
}

func Mod[T interface {
	constraints.Float | constraints.Signed
}](x, y T) T {
	t := math.Mod(float64(x), float64(y))
	if t < 0 {
		t += float64(y)
	}
	return T(t)
}

func Pow[T interface {
	constraints.Float | constraints.Integer
}](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Sqrt[T interface {
	constraints.Float | constraints.Integer
}](x T) T {
	return T(math.Sqrt(float64(x)))
}

// max(min(x, h), l)
func Clamp[T cmp.Ordered](x, low, high T) T {
	return max(min(x, high), low)
}

// (x-l1)/(h1-l1)*(h2-l2) + l2
func Remap(x, low1, high1, low2, high2 float64) float64 {
	return (x-low1)/(high1-low1)*(high2-low2) + low2
}

// x*PI/180
func Radians(degrees float64) float64 {
	return degrees * PI / 180
}

// x*180/pi
func Degrees(radians float64) float64 {
	return radians * 180 / PI
}

func Lerp(c1, c2, coeff float64) float64 {
	return c1*(1-coeff) + c2*coeff
}
