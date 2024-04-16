package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// SolarFlare draws a solar flare images
// TODO: looks very different from original
func SolarFlare(dc *Context, lineColor color.NRGBA) {
	const offsetInc = 0.006
	const inc = 1.0
	const m = 1.009

	noise := NewPerlinNoiseDeprecated()
	im := dc.Image()
	dc.SetColor(Black)
	dc.Clear()

	dc.TransformAdd(Translate(Size(im) / 2))
	dc.SetLineWidth(1)

	offset := 0.0
	for r := 1.0; r < 200; {
		lineColor.A = uint8(Remap(Pow(Remap(r, 1, 200, 0, 1), 2.5), 0, 1, 0, 255))
		dc.SetColor(lineColor)
		for range Range(10) {
			n := max(int(2*PI*r), 500)
			for _, a := range RangeF64(0, PI*2, n) {
				coeff := noise.NoiseV2_1(Diag(offset) + Polar(inc, a))
				dc.LineTo(Polar(coeff*r, a))
			}
			dc.Stroke()

			offset += offsetInc
			r *= m
		}
	}
}

func solarFlare() *image.RGBA {
	dc := NewContext(Diag(500))
	SolarFlare(dc, color.NRGBA{255, 64, 8, 255})
	return dc.Image()
}
