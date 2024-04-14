package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/rprtr258/gena"
)

// PixelHole draws a hole with colored dots.
//   - dotN: The number of point in each iteration.
func PixelHole(c *image.RGBA, colorSchema []color.RGBA, dotN, iters int) {
	dc := gena.NewContextForRGBA(c)
	noise := gena.NewPerlinNoiseDeprecated()
	for i := range gena.Range(iters) {
		dc.Stack(func(ctx *gena.Context) {
			dc.Translate(gena.Size(c) / 2)
			dc.SetLineWidth(2.0)
			fc := float64(i)
			c1 := int(fc/100.0) % len(colorSchema)
			c2 := (int(fc/100.0) + 1) % len(colorSchema)
			ratio := fc/100 - math.Floor(fc/100)
			cl := gena.LerpColor(colorSchema[c1], colorSchema[c2], ratio)
			for j := 0.0; j < float64(dotN); j += 1.0 {
				dc.Stack(func(ctx *gena.Context) {
					dc.SetColor(cl)
					dc.Rotate(fc/(50+10*math.Log(fc+1)) + j/20)
					dd := fc/(5+j) + fc/5 + math.Sin(j)*50
					dc.Translate(complex(gena.RandomFloat64(dd/2, dd), 0))
					x := noise.Noise2_1(fc/50+j/50, 5000)*float64(c.Bounds().Dx())/10 + rand.Float64()*float64(c.Bounds().Dx())/20
					y := noise.Noise2_1(fc/50+j/50, 10000)*float64(c.Bounds().Dy())/10 + rand.Float64()*float64(c.Bounds().Dy())/20

					rr := gena.RandomFloat64(1.0, 6-math.Log(fc+1)/10)
					dc.DrawCircle(complex(x, y), rr)
					dc.Fill()
				})
			}
		})
	}
}
