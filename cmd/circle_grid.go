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

func (cg *circleGrid) grid(dc *Context, im *image.RGBA) {
	var segment int = 100
	w := float64(im.Bounds().Dx()) / float64(segment)

	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.SetLineWidth(0.6)
	for i := range Range(segment) {
		dc.DrawLine(complex(0, float64(i)*w), complex(float64(im.Bounds().Dx()), float64(i)*w))
		dc.Stroke()
	}

	for j := range Range(segment) {
		dc.DrawLine(complex(float64(j)*w, 0), complex(float64(j)*w, float64(im.Bounds().Dy())))
		dc.Stroke()
	}
}

// Generative draws a circle grid image.
func CircleGrid(im *image.RGBA, colorSchema []color.RGBA, lineWidth float64, circleNumMin, circleNumMax int) {
	cg := &circleGrid{
		circleNumMin: circleNumMin,
		circleNumMax: circleNumMax,
	}

	dc := NewContextFromRGBA(im)
	cg.grid(dc, im)
	dc.TransformAdd(Translate(Size(im) / 2))
	dc.TransformAdd(Scale(complex(0.9, 0.9)))
	dc.TransformAdd(Translate(-Size(im) / 2))

	seg := RandomIntN(cg.circleNumMin, cg.circleNumMax)
	w := float64(im.Bounds().Dx()) / float64(seg)

	for i := range Range(seg) {
		for j := range Range(seg) {
			v := complex(float64(i), float64(j))*Coeff(w) + Diag(w/2)
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			dc.DrawCircle(v, w/2*RandomF64(0.1, 0.5))
			dc.Fill()

			// draw
			r := w / 2 * RandomF64(0.6, 0.95)
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
					n := RandomIntN(1, 4) * 2
					dc.DrawCircle(0, r)
					dc.Stroke()
					for range Range(n) {
						dc.TransformAdd(Rotate(math.Pi * 2 / float64(n)))
						dc.DrawCircle(complex(r, 0), r*0.1)
						dc.Fill()
					}
				case 2:
					n := RandomIntN(8, 20)
					theta := math.Pi * 0.5 * float64(RandomIntN(1, 5))
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
					n := RandomIntN(5, 20)
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
