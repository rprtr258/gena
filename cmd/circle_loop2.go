package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

func recursionDraw(
	dc *Context,
	im *image.RGBA,
	x float64, depth int,
	noise0 *PerlinNoise,
	palette Pattern1D,
) {
	if depth <= 0 {
		return
	}

	H := Size(im).Y()

	lw := 1.0
	if Random() >= 0.8 {
		lw = RandomF64(1.0, RandomF64(1, 3))
	}
	dc.SetLineWidth(lw)

	noise := noise0.Noise3_1(x*0.02+123.234, (1-x)*0.02, 345.4123)
	noise = Sqrt(noise)
	a2 := Remap(noise, 0.15, 0.85, 0.1, 0.6)

	p := P(
		H*Pow(x/H, a2),
		H*(Pow(1-x/H, a2)-
			RandomF64(0, RandomF64(0.18, RandomF64(0.18, 0.7)))),
	) * 0.39

	dc.SetColor(palette.Random())

	nCircles := RandomIntN(1, 6)
	if Random() < 0.03 {
		nCircles = RandomIntN(8, 10)
	}

	r := Pow(Random(), 2) * 50
	if Random() < 0.7 {
		for _, z := range RangeF64(0, r, nCircles) {
			dc.DrawCircle(p, Random()*z)
			dc.Stroke()
		}
	} else {
		for _, z := range RangeF64(0, r, nCircles) {
			dc.DrawCircle(p, z)
			dc.Stroke()
		}
	}

	dc.TransformAdd(Rotate(x / H * 0.2))

	recursionDraw(dc, im, 1*x/4.0, depth-1, noise0, palette)
	recursionDraw(dc, im, 2*x/4.0, depth-1, noise0, palette)
	recursionDraw(dc, im, 3*x/4.0, depth-1, noise0, palette)
}

// CircleLoop2 draws a circle composed by colored circles.
//   - depth: Control the number of circles.
func CircleLoop2(dc *Context, palette Pattern1D, depth int) {
	dc.SetColor(color.RGBA{8, 10, 20, 255})
	dc.Clear()

	dc.TransformAdd(Translate(Size(dc.Image()) / 2))
	noise := NewPerlinNoiseDeprecated()
	recursionDraw(dc, dc.Image(), float64(dc.Image().Bounds().Dx()), depth, &noise, palette)
}

func circleloop2() *image.RGBA {
	dc := NewContext(Diag(500))
	CircleLoop2(dc, NewPaletteFromRGBAs([]color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}...), 7)
	return dc.Image()
}
