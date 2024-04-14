package gena

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
	var p [perlinSize + 1]float64
	for i := range p {
		p[i] = Random()
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

// Noise1_1 returns a float noise number on one dimension.
func (p *PerlinNoise) Noise1_1(x float64) float64 {
	return p.noise(x, 0, 0)
}

// Noise2_1 returns a float noise number on two dimensions.
func (p *PerlinNoise) Noise2_1(x, y float64) float64 {
	return p.noise(x, y, 0)
}

func (p *PerlinNoise) NoiseV2_1(v V2) float64 {
	return p.noise(X(v), Y(v), 0)
}

func (p *PerlinNoise) Noise2_V2(x, y float64) V2 {
	return complex(p.Noise1_1(x), p.Noise1_1(y))
}

func (p *PerlinNoise) Noise2V2_V2(v, d V2) V2 {
	return complex(p.NoiseV2_1(v), p.NoiseV2_1(v+d))
}

// Noise3_1 returns a float noise number on three dimensions.
func (p *PerlinNoise) Noise3_1(x, y, z float64) float64 {
	return p.noise(x, y, z)
}

func scaledCosin(x float64) float64 {
	return (1 - Cos(x*PI)) / 2
}

func (p *PerlinNoise) noise(x, y, z float64) float64 {
	x, y, z = Abs(x), Abs(y), Abs(z)
	xi, yi, zi := int(Floor(x)), int(Floor(y)), int(Floor(z))
	xf, yf, zf := x-float64(xi), y-float64(yi), z-float64(zi)

	var rxf, ryf, n1, n2, n3 float64
	var r float64
	var ampl float64 = 0.5
	for range Range(perlinOctaves) {
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
