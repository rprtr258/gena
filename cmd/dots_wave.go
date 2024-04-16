package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a dots wave images.
//   - n: The number of dots wave in the image.
func DotsWave(dc *Context, colorSchema []color.RGBA, n int) {
	noise := NewPerlinNoiseDeprecated()

	im := dc.Image()
	dc.SetColor(Black)
	dc.Clear()
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

func dotswave() *image.RGBA {
	dc := NewContext(Diag(500))
	DotsWave(dc, []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}, 300)
	return dc.Image()
}
