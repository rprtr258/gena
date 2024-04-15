package gena

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func RandomInt[I constraints.Integer](max I) I {
	return I(rand.Intn(int(max)))
}

// rand()*(max-min) + min
func RandomIntN[I constraints.Integer](min, max I) I {
	return I(RandomInt(int(max)-int(min)) + int(min))
}

func Random() float64 {
	return rand.Float64()
}

// rand()*(max-min) + min
func RandomF64(min, max float64) float64 {
	return Random()*(max-min) + min
}

func RandomGaussianNormal() float64 {
	return rand.NormFloat64()
}

// RandomGaussian returns a gaussian random float64 number
func RandomGaussian(mean, std float64) float64 {
	return RandomGaussianNormal()*std + mean
}

// (rand(), rand())
func RandomV2() V2 {
	return complex(Random(), Random())
}

func RandomV2N(min, max V2) V2 {
	return Mul2(RandomV2(), (max-min)) + min
}

func RandomItem[T any](items []T) T {
	return items[RandomInt(len(items))]
}

func RandomWeighted[K comparable](items map[K]float64) K {
	sum := 0.0
	for _, p := range items {
		sum += p
	}

	r := Random() * sum
	for k, p := range items {
		r -= p
		if r <= 0 {
			return k
		}
	}
	for k := range items {
		return k
	}
	panic("empty map")
}

func Shuffle[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}
