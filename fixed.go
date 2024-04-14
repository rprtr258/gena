package gena

import (
	"math"

	"golang.org/x/image/math/fixed"
)

func Fixed(v V2) fixed.Point26_6 {
	return fixed.Point26_6{X: fix(X(v)), Y: fix(Y(v))}
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func unfix(x fixed.Int26_6) float64 {
	const shift, mask = 6, 1<<6 - 1
	if x >= 0 {
		return float64(x>>shift) + float64(x&mask)/64
	}

	x = -x
	if x >= 0 {
		return -(float64(x>>shift) + float64(x&mask)/64)
	}

	return 0
}
