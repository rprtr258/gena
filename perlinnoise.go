package gena

import (
	"math"
	"math/rand"
)

const (
	perlinYwrapb     = 4
	perlinYwrap      = 1 << perlinYwrapb
	perlinZwrapb     = 8
	perlinZwrap      = 1 << perlinZwrapb
	perlinSize       = 4095
	perlinOctaves    = 4
	perlinAmpFalloff = 0.5
)

// PerlinNoise is used to generate perlin noise.
type PerlinNoise struct {
	perlin [perlinSize + 1]float64
}

func NewPerlinNoiseDeprecated() PerlinNoise {
	var p [4096]float64
	for i := 0; i < len(p); i++ {
		p[i] = rand.Float64()
	}
	return PerlinNoise{
		perlin: p,
	}
}

// NewPerlinNoise returns a PerlinNoise struct.
func NewPerlinNoise(perlin [perlinSize + 1]float64) PerlinNoise {
	return PerlinNoise{
		perlin: perlin,
	}
}

// Noise1D returns a float noise number on one dimension.
func (p *PerlinNoise) Noise1D(x float64) float64 {
	return p.noise(x, 0, 0)
}

// Noise2D returns a float noise number on two dimensions.
func (p *PerlinNoise) Noise2D(x, y float64) float64 {
	return p.noise(x, y, 0)
}

func (p *PerlinNoise) NoiseV2(v V2) float64 {
	return p.noise(X(v), Y(v), 0)
}

func (p *PerlinNoise) NoiseV2x2(v, d V2) V2 {
	return complex(p.NoiseV2(v), p.NoiseV2(v+d))
}

// Noise3D returns a float noise number on three dimensions.
func (p *PerlinNoise) Noise3D(x, y, z float64) float64 {
	return p.noise(x, y, z)
}

func scaledCosin(x float64) float64 {
	return (1 - math.Cos(x*math.Pi)) / 2
}

func (p *PerlinNoise) noise(x, y, z float64) float64 {
	x, y, z = math.Abs(x), math.Abs(y), math.Abs(z)
	xi, yi, zi := int(math.Floor(x)), int(math.Floor(y)), int(math.Floor(z))
	xf, yf, zf := x-float64(xi), y-float64(yi), z-float64(zi)

	var rxf, ryf, n1, n2, n3 float64
	var r float64
	var ampl float64 = 0.5
	for o := 0; o < perlinOctaves; o++ {
		of := xi + (yi << perlinYwrapb) + (zi << perlinZwrapb)

		rxf = scaledCosin(xf)
		ryf = scaledCosin(yf)

		n1 = p.perlin[of&perlinSize]
		n1 += rxf * (p.perlin[(of+1)&perlinSize] - n1)
		n2 = p.perlin[(of+perlinYwrap)&perlinSize]
		n2 += rxf * (p.perlin[(of+perlinYwrap+1)&perlinSize] - n2)
		n1 += ryf * (n2 - n1)

		of += perlinZwrap
		n2 = p.perlin[of&perlinSize]
		n2 += rxf * (p.perlin[(of+1)&perlinSize] - n2)
		n3 = p.perlin[(of+perlinYwrap)&perlinSize]
		n3 += rxf * (p.perlin[(of+perlinYwrap+1)&perlinSize] - n3)
		n2 += ryf * (n3 - n2)

		n1 += scaledCosin(zf) * (n2 - n1)

		r += n1 * ampl
		ampl *= perlinAmpFalloff
		xi <<= 1
		xf *= 2
		yi <<= 1
		yf *= 2
		zi <<= 1
		zf *= 2

		if xf >= 1.0 {
			xi += 1
			xf -= 1
		}
		if yf >= 1.0 {
			yi += 1
			yf -= 1
		}
		if zf >= 1.0 {
			zi += 1
			zf -= 1
		}
	}
	return r
}
