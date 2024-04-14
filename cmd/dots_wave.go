package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a dots wave images.
//   - n: The number of dots wave in the image.
func DotsWave(im *image.RGBA, colorSchema []color.RGBA, n int) {
	dc := NewContextFromRGBA(im)
	noise := NewPerlinNoiseDeprecated()
	for range Range(n) {
		v := Mul2(complex(
			RandomF64(-0.1, 1.1),
			RandomF64(-0.1, 1.1),
		), Size(im))

		num := RandomF64(100, 1000)
		r := Random() * float64(im.Bounds().Dx()) * 0.15 * Random()
		ind := RandomF64(1, 8)

		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(v))
			dc.TransformAdd(Rotate(float64(RandomInt(8)) * PI / 4))
			Shuffle(colorSchema)
			for j := 0.0; j < num; j += ind {
				s := float64(im.Bounds().Dx()) * 0.15 * RandomF64(0, RandomF64(0, RandomF64(0, RandomF64(0, RandomF64(0, RandomF64(0, Random()))))))
				ci := int(float64(len(colorSchema)) * noise.Noise3_1(j*0.01, X(v), Y(v)))
				dc.SetColor(colorSchema[ci])
				dc.DrawCircle(complex(j, r*Sin(j*0.05)), s*2/3)
				dc.Fill()
			}
		})
	}
}
