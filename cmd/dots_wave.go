package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

// Generative draws a dots wave images.
//   - dotsN: The number of dots wave in the image.
func DotsWave(c *image.RGBA, colorSchema []color.RGBA, dotsN int) {
	dc := NewContextForRGBA(c)
	noise := NewPerlinNoiseDeprecated()
	for range Range(dotsN) {
		v := Mul2(complex(
			RandomFloat64(-0.1, 1.1),
			RandomFloat64(-0.1, 1.1),
		), Size(c))

		num := RandomFloat64(100, 1000)
		r := rand.Float64() * float64(c.Bounds().Dx()) * 0.15 * rand.Float64()
		ind := RandomFloat64(1, 8)

		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(v))
			dc.TransformAdd(Rotate(float64(rand.Intn(8)) * math.Pi / 4))
			rand.Shuffle(len(colorSchema), func(i, j int) {
				colorSchema[i], colorSchema[j] = colorSchema[j], colorSchema[i]
			})
			for j := 0.0; j < num; j += ind {
				s := float64(c.Bounds().Dx()) * 0.15 * RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, rand.Float64()))))))
				ci := int(float64(len(colorSchema)) * noise.Noise3_1(j*0.01, X(v), Y(v)))
				dc.SetColor(colorSchema[ci])
				dc.DrawCircle(complex(j, r*math.Sin(j*0.05)), s*2/3)
				dc.Fill()
			}
		})
	}
}
