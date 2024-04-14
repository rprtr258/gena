package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

// PixelHole draws a hole with colored dots.
//   - n: number of points in each iteration
func PixelHole(im *image.RGBA, colorSchema []color.RGBA, n, iters int) {
	dc := NewContextFromRGBA(im)
	noise := NewPerlinNoiseDeprecated()
	for i := range Range(iters) {
		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(Size(im) / 2))
			dc.SetLineWidth(2.0)
			fc := float64(i)
			c1 := int(fc/100.0) % len(colorSchema)
			c2 := (int(fc/100.0) + 1) % len(colorSchema)
			ratio := fc/100 - math.Floor(fc/100)
			cl := ColorLerp(colorSchema[c1], colorSchema[c2], ratio)
			for j := 0.0; j < float64(n); j += 1.0 {
				dc.Stack(func(ctx *Context) {
					dc.SetColor(cl)
					dc.TransformAdd(Rotate(fc/(50+10*math.Log(fc+1)) + j/20))
					dd := fc/(5+j) + fc/5 + math.Sin(j)*50
					dc.TransformAdd(Translate(complex(RandomFloat64(dd/2, dd), 0)))
					x := noise.Noise2_1(fc/50+j/50, 5000)*float64(im.Bounds().Dx())/10 + rand.Float64()*float64(im.Bounds().Dx())/20
					y := noise.Noise2_1(fc/50+j/50, 10000)*float64(im.Bounds().Dy())/10 + rand.Float64()*float64(im.Bounds().Dy())/20

					rr := RandomFloat64(1.0, 6-math.Log(fc+1)/10)
					dc.DrawCircle(complex(x, y), rr)
					dc.Fill()
				})
			}
		})
	}
}
