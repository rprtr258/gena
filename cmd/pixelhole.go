package main

import (
	"image"
	"image/color"
	"math"

	. "github.com/rprtr258/gena"
)

// PixelHole draws a hole with colored dots.
//   - n: number of points in each iteration
func PixelHole(dc *Context, colorSchema []color.RGBA, n, iters int) {
	noise := NewPerlinNoiseDeprecated()

	im := dc.Image()
	dc.SetColor(Black)
	dc.Clear()
	for _, fc := range RangeF64(0, float64(iters), iters) {
		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(Size(im) / 2))
			dc.SetLineWidth(2.0)
			c1 := int(fc/100.0) % len(colorSchema)
			c2 := (int(fc/100.0) + 1) % len(colorSchema)
			ratio := fc/100 - Floor(fc/100)
			cl := ColorLerp(colorSchema[c1], colorSchema[c2], ratio)
			for _, j := range RangeStepF64(0, float64(n), 1.0) {
				dc.Stack(func(ctx *Context) {
					dc.SetColor(cl)
					dc.TransformAdd(Rotate(fc/(50+10*math.Log(fc+1)) + j/20))
					dd := fc/(5+j) + fc/5 + Sin(j)*50
					dc.TransformAdd(Translate(complex(RandomF64(dd/2, dd), 0)))
					x := noise.Noise2_1(fc/50+j/50, 5000)*float64(im.Bounds().Dx())/10 + Random()*float64(im.Bounds().Dx())/20
					y := noise.Noise2_1(fc/50+j/50, 10000)*float64(im.Bounds().Dy())/10 + Random()*float64(im.Bounds().Dy())/20

					rr := RandomF64(1.0, 6-math.Log(fc+1)/10)
					dc.DrawCircle(complex(x, y), rr)
					dc.Fill()
				})
			}
		})
	}
}

func pixelhole() *image.RGBA {
	dc := NewContext(Diag(800))
	PixelHole(dc, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 60, 1200)
	return dc.Image()
}
