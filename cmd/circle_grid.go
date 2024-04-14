package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

type circleGrid struct {
	circleNumMin, circleNumMax int
}

func (cg *circleGrid) grid(dc *Context, c *image.RGBA) {
	var segment int = 100
	w := float64(c.Bounds().Dx()) / float64(segment)

	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.SetLineWidth(0.6)
	for i := range Range(segment) {
		dc.DrawLine(complex(0, float64(i)*w), complex(float64(c.Bounds().Dx()), float64(i)*w))
		dc.Stroke()
	}

	for j := range Range(segment) {
		dc.DrawLine(complex(float64(j)*w, 0), complex(float64(j)*w, float64(c.Bounds().Dy())))
		dc.Stroke()
	}
}

// Generative draws a circle grid image.
func CircleGrid(c *image.RGBA, colorSchema []color.RGBA, lineWidth float64, circleNumMin, circleNumMax int) {
	cg := &circleGrid{
		circleNumMin: circleNumMin,
		circleNumMax: circleNumMax,
	}

	dc := NewContextFromRGBA(c)
	cg.grid(dc, c)
	dc.TransformAdd(Translate(Size(c) / 2))
	dc.TransformAdd(Scale(complex(0.9, 0.9)))
	dc.TransformAdd(Translate(-Size(c) / 2))

	seg := RandomRangeInt(cg.circleNumMin, cg.circleNumMax)
	w := float64(c.Bounds().Dx()) / float64(seg)

	for i := range Range(seg) {
		for j := range Range(seg) {
			v := complex(float64(i), float64(j))*Coeff(w) + Diag(w/2)
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			dc.DrawCircle(v, w/2*RandomFloat64(0.1, 0.5))
			dc.Fill()

			// draw
			r := w / 2 * RandomFloat64(0.6, 0.95)
			rnd := rand.Intn(4)
			col := colorSchema[rand.Intn(len(colorSchema))]
			dc.Stack(func(dc *Context) {
				dc.TransformAdd(Translate(v))
				dc.TransformAdd(Rotate(float64(rand.Intn(10))))
				dc.SetColor(col)
				dc.SetLineWidth(lineWidth)

				switch rnd {
				case 0:
					dc.DrawCircle(0, r)
					dc.Stroke()
				case 1:
					n := RandomRangeInt(1, 4) * 2
					dc.DrawCircle(0, r)
					dc.Stroke()
					for range Range(n) {
						dc.TransformAdd(Rotate(math.Pi * 2 / float64(n)))
						dc.DrawCircle(complex(r, 0), r*0.1)
						dc.Fill()
					}
				case 2:
					n := RandomRangeInt(8, 20)
					theta := math.Pi * 0.5 * float64(RandomRangeInt(1, 5))
					for i := range Range(n) {
						d := float64(i) / float64(n)
						if d > r*0.1 {
							d = r * 0.1
						}
						dc.TransformAdd(Rotate(theta / float64(n)))
						dc.DrawCircle(complex(r/2, 0), d*2)
						dc.Fill()
					}
				case 3:
					n := RandomRangeInt(5, 20)
					for range Range(n) {
						dc.TransformAdd(Rotate(math.Pi * 2 / float64(n)))
						dc.DrawLine(complex(r/2, 0), complex(r*2/3-r*0.05, 0))
						dc.Stroke()
					}
				}
			})
		}
	}
}
