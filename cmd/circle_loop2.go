package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

type circleLoop2 struct {
	noise       PerlinNoise
	colorSchema []color.RGBA
}

// CircleLoop2 draws a circle composed by colored circles.
//   - depth: Control the number of circles.
func CircleLoop2(im *image.RGBA, colorSchema []color.RGBA, depth int) {
	dc := NewContextFromRGBA(im)
	dc.TransformAdd(Translate(Size(im) / 2))
	cl := &circleLoop2{
		noise:       NewPerlinNoiseDeprecated(),
		colorSchema: colorSchema,
	}
	cl.recursionDraw(dc, im, float64(im.Bounds().Dx()), depth)
}

func (cl *circleLoop2) recursionDraw(dc *Context, im *image.RGBA, x float64, depth int) {
	if depth <= 0 {
		return
	}

	H := Y(Size(im))

	lw := 1.0
	if Random() >= 0.8 {
		lw = RandomF64(1.0, RandomF64(1, 3))
	}
	dc.SetLineWidth(lw)

	noise := cl.noise.Noise3_1(x*0.02+123.234, (1-x)*0.02, 345.4123)
	noise = Sqrt(noise)
	a2 := Remap(noise, 0.15, 0.85, 0.1, 0.6)

	p := complex(
		H*Pow(x/H, a2),
		H*(Pow(1-x/H, a2)-
			RandomF64(0, RandomF64(0.18, RandomF64(0.18, 0.7)))),
	) * 0.39

	dc.SetColor(RandomItem(cl.colorSchema))

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

	cl.recursionDraw(dc, im, 1*x/4.0, depth-1)
	cl.recursionDraw(dc, im, 2*x/4.0, depth-1)
	cl.recursionDraw(dc, im, 3*x/4.0, depth-1)
}
