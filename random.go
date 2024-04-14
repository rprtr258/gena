package gena

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// rand()*(max-min) + min
func RandomIntN[I constraints.Integer](min, max I) I {
	return I(rand.Intn(int(max)-int(min)) + int(min))
}

// rand()*(max-min) + min
func RandomF64(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

// RandomGaussian returns a gaussian random float64 number.
func RandomGaussian(mean, std float64) float64 {
	return rand.NormFloat64()*std + mean
}

// (rand(), rand())
func RandomV2() V2 {
	return complex(rand.Float64(), rand.Float64())
}
