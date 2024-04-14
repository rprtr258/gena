package main

import (
	"image"
	"image/color"
	"image/draw"

	. "github.com/rprtr258/gena"
)

// SolarFlare draws a solar flare images
func SolarFlare(c image.Image, lineColor color.Color) {
	noise := NewPerlinNoiseDeprecated()
	const offsetInc = 0.006
	const inc = 1.0
	const m = 1.005

	var offset V2
	for r := 1.0; r < 200; {
		for range Range(10) {
			nPoints := min(int(2*PI*r), 500)

			img := image.NewRGBA(image.Rect(0, 0, c.Bounds().Dx(), c.Bounds().Dy()))
			draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)
			dc := NewContextFromRGBA(img)

			dc.Stack(func(ctx *Context) {
				dc.TransformAdd(Translate(Size(c) / 2))
				dc.SetLineWidth(1.0)
				dc.SetColor(lineColor)
				for j := 0; j < nPoints+1; j += 1 {
					a := float64(j) / float64(nPoints) * PI * 2
					p := Polar(1, a)
					dc.LineTo(p * Coeff(noise.NoiseV2_1(offset+p*Coeff(inc))*r))
				}
				dc.Stroke()
			})

			c = Blend(img, c, Add)
			offset += Diag(offsetInc)
			r *= m
		}
	}
}
