package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	. "github.com/rprtr258/gena"
)

// SolarFlare draws a solar flare images.
func SolarFlare(c image.Image, lineColor color.Color) {
	var xOffset, yOffset float64
	const offsetInc = 0.006
	const inc = 1.0
	const m = 1.005
	noise := NewPerlinNoiseDeprecated()

	for r := 1.0; r < 200; {
		for range Range(10) {
			nPoints := int(2 * math.Pi * r)
			nPoints = min(nPoints, 500)

			img := image.NewRGBA(image.Rect(0, 0, c.Bounds().Dx(), c.Bounds().Dy()))
			draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)
			dc := NewContextFromRGBA(img)

			dc.Stack(func(ctx *Context) {
				dc.TransformAdd(Translate(Size(c) / 2))
				dc.SetLineWidth(1.0)
				dc.SetColor(lineColor)
				for j := 0; j < nPoints+1; j += 1 {
					a := float64(j) / float64(nPoints) * math.Pi * 2
					px := math.Cos(a)
					py := math.Sin(a)
					n := noise.Noise2_1(xOffset+px*inc, yOffset+py*inc) * r
					px *= n
					py *= n
					dc.LineTo(complex(px, py))
				}
				dc.Stroke()
			})

			c = Blend(img, c, Add)
			xOffset += offsetInc
			yOffset += offsetInc
			r *= m
		}
	}
}
