package gena

import (
	"cmp"
	"math"
)

// max(min(x, h), l)
func Constrain[T cmp.Ordered](x, low, high T) T {
	return max(min(x, high), low)
}

// (x-l1)/(h1-l1)*(h2-l2) + l2
func Remap(x, low1, high1, low2, high2 float64) float64 {
	return (x-low1)/(high1-low1)*(high2-low2) + low2
}

// x*PI/180
func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// x*180/pi
func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// max(-x,x)
func Abs[T ~int32 | ~float64](x T) T {
	return max(-x, x)
}
