package gena

import (
	"fmt"
	"image/color"
	"log"
)

// ColorMapping maps some parameters to color space.
type ColorMapping func(float64, float64, float64) color.RGBA

type domainWrap struct {
	noise  PerlinNoise
	scale  float64
	scale2 float64
	offset V2
	fn     ColorMapping
	// How many images would be created in this generation.
	numImages int
	// Use these parameters to create images with time lapse.
	offsetStep V2
	// The imagPath for generative images.
	imgPath string
}

// Generative draws a domain warp image.
// Reference: https://www.iquilezles.org/www/articles/warp/warp.htm
func DomainWrap(
	c Canvas,
	scale, scale2 float64,
	offset V2,
	cmap ColorMapping,
	step V2,
	n int,
	path string,
) {
	d := &domainWrap{
		scale:      scale,
		scale2:     scale2,
		offset:     offset,
		noise:      NewPerlinNoiseDeprecated(),
		fn:         cmap,
		offsetStep: step,
		numImages:  n,
		imgPath:    path,
	}

	if d.numImages == 0 && d.imgPath == "" {
		d.generative(c)
		return
	}

	if d.numImages > 0 && d.imgPath == "" {
		log.Fatal("Missing the parameters numImages or imgPath")
	}

	for i := 0; i < d.numImages; i++ {
		imgfile := fmt.Sprintf("%v/domainwrap%03d.PNG", d.imgPath, i)
		d.offset += Mul(d.offsetStep, float64(i))
		d.generative(c)
		c.ToPNG(imgfile)
	}
}

func (d *domainWrap) generative(c Canvas) {
	for h := 0.0; h < float64(c.Height()); h += 1.0 {
		for w := 0.0; w < float64(c.Width()); w += 1.0 {
			v := Mul(complex(w, h), d.scale)
			q := d.noise.NoiseV2x2(v+d.offset, complex(5.2, 1.3))
			r := d.noise.NoiseV2x2(v+Mul(q, d.scale2)+complex(1.7, 9.2), complex(6.4, -6.4))
			r1, m1, m2 := d.noise.NoiseV2(q+Mul(r, d.scale2)), Magnitude(q), Magnitude(r)
			color := d.fn(r1, m1, m2)
			c.Img().Set(int(w), int(h), color)
		}
	}
}
