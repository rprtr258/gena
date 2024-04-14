package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

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

	var lw float64
	if rand.Float64() < 0.8 {
		lw = 1
	} else {
		lw = RandomFloat64(1.0, RandomFloat64(1, 3))
	}
	dc.SetLineWidth(lw)

	noise := cl.noise.Noise3_1(x*0.02+123.234, (1-x)*0.02, 345.4123)
	noise = math.Pow(noise, 0.5)
	a2 := Remap(noise, 0.15, 0.85, 0.1, 0.6)

	px := float64(im.Bounds().Dy()) * math.Pow(x/float64(im.Bounds().Dy()), a2)
	py := float64(im.Bounds().Dy()) * (math.Pow(1-x/float64(im.Bounds().Dy()), a2) -
		RandomFloat64(0, RandomFloat64(0.18, RandomFloat64(0.18, 0.7))))

	dc.SetColor(cl.colorSchema[rand.Intn(len(cl.colorSchema))])

	nCircles := RandomRangeInt(1, 6)
	if rand.Float64() < 0.03 {
		nCircles = RandomRangeInt(8, 10)
	}

	r := math.Pow(rand.Float64(), 2) * 50

	if rand.Float64() < 0.7 {
		for i := range Range(nCircles) {
			dc.DrawCircle(complex(px, py)*0.39, rand.Float64()*float64(i)*r/float64(nCircles))
			dc.Stroke()
		}
	} else {
		for i := range Range(nCircles) {
			dc.DrawCircle(complex(px, py)*0.39, float64(i)*r/float64(nCircles))
			dc.Stroke()
		}
	}

	dc.TransformAdd(Rotate(x / float64(im.Bounds().Dy()) * 0.2))

	cl.recursionDraw(dc, im, 1*x/4.0, depth-1)
	cl.recursionDraw(dc, im, 2*x/4.0, depth-1)
	cl.recursionDraw(dc, im, 3*x/4.0, depth-1)
}
