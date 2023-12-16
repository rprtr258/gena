package gena

import "math/rand"

// rand()*(max-min) + min
func RandomRangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// rand()*(max-min) + min
func RandomFloat64(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

// RandomGaussian returns a gaussian random float64 number.
func RandomGaussian(mean, std float64) float64 {
	return rand.NormFloat64()*std + mean
}

// rand()+rand()*i
func RandomV2() V2 {
	return complex(rand.Float64(), rand.Float64())
}
