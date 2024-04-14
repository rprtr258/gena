package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	. "github.com/rprtr258/gena"
)

// ColorMapping maps some parameters to color space.
type ColorMapping = func(float64, float64, float64) color.RGBA

func domainWarp(
	c *image.RGBA,
	noise PerlinNoise,
	scale float64,
	scale2 float64,
	offset V2,
	fn ColorMapping,
) {
	for h := 0; h < c.Bounds().Dy(); h++ {
		for w := 0; w < c.Bounds().Dx(); w++ {
			v := complex(float64(w), float64(h)) * Coeff(scale)
			q := noise.NoiseV2x2(v+offset, complex(5.2, 1.3))
			r := noise.NoiseV2x2(v+q*Coeff(scale2)+complex(1.7, 9.2), complex(6.4, -6.4))
			r1, m1, m2 := noise.NoiseV2(q+r*Coeff(scale2)), Magnitude(q), Magnitude(r)
			color := fn(r1, m1, m2)
			c.Set(w, h, color)
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
//   - cmap: A function to mapping the `noise` to color.
func DomainWarp(
	c *image.RGBA,
	scale, scale2 float64,
	offset V2,
	cmap ColorMapping,
	// Use these parameters to create images with time lapse.
	offsetStep V2,
	// How many images would be created in this generation.
	numImages int,
	// The imagPath for generative images.
	imgPath string,
) {
	noise := NewPerlinNoiseDeprecated()

	if numImages == 0 && imgPath == "" {
		domainWarp(c, noise, scale, scale2, offset, cmap)
		return
	}

	if numImages > 0 && imgPath == "" {
		log.Fatal("Missing the parameters numImages or imgPath")
	}

	for i := range Range(numImages) {
		imgfile := fmt.Sprintf("%v/domainwrap%03d.PNG", imgPath, i)
		offset += offsetStep * Coeff(float64(i))
		domainWarp(c, noise, scale, scale2, offset, cmap)
		SavePNG(imgfile, c)
	}
}
