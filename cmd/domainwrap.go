package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Pattern3D maps some parameters to color space.
type Pattern3D = func(float64, float64, float64) color.RGBA

var cmap Pattern3D = func(r, m1, m2 float64) color.RGBA {
	return color.RGBA{
		R: uint8(Clamp(m1*200*r, 0, 255)),
		G: uint8(Clamp(r*200, 0, 255)),
		B: uint8(Clamp(m2*255*r, 70, 255)),
		A: 255,
	}
}

func domainWarp(
	dc *Context,
	noise PerlinNoise,
	scale float64,
	scale2 float64,
	offset V2,
) {
	dc.SetColor(Black)
	dc.Clear()
	im := dc.Image()
	for h := 0; h < im.Bounds().Dy(); h++ {
		for w := 0; w < im.Bounds().Dx(); w++ {
			v := P(float64(w), float64(h)) * Coeff(scale)
			q := noise.Noise2V2_V2(v+offset, P(5.2, 1.3))
			r := noise.Noise2V2_V2(v+q*Coeff(scale2)+P(1.7, 9.2), P(6.4, -6.4))
			r1, m1, m2 := noise.NoiseV2_1(q+r*Coeff(scale2)), q.Magnitude(), r.Magnitude()
			color := cmap(r1, m1, m2)
			im.Set(w, h, color)
		}
	}
}

// Generative draws a domain warp image.
// Reference: https://www.iquilezles.org/www/articles/warp/warp.htm
// Warping, or domain distortion is a very common technique in computer graphics
// for generating procedural textures and geometry. It's often used to pinch an object,
// stretch it, twist it, bend it, make it thicker or apply any deformation you want.
//   - scale: Control the noise generator.
//   - offset: Control the noise generator.
func DomainWarp(
	scale, scale2 float64,
	offset V2,
	// Use these parameters to create images with time lapse.
	offsetStep V2,
	// How many images would be created in this generation.
	numImages int,
) []*image.RGBA {
	noise := NewPerlinNoiseDeprecated()

	res := make([]*image.RGBA, numImages)
	for i := range Range(numImages) {
		dc := NewContext(Diag(500))
		offset += offsetStep * Coeff(float64(i))
		domainWarp(dc, noise, scale, scale2, offset)
		res[i] = dc.Image()
	}
	return res
}

func domainwrap() *image.RGBA {
	dc := NewContext(Diag(500))
	domainWarp(dc, NewPerlinNoiseDeprecated(), 0.01, 4, P(4, 20))
	return dc.Image()
}
