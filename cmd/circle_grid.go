package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// CircleGrid draws a circle grid image
func CircleGrid(im *image.RGBA, colorSchema []color.RGBA, lineWidth float64, n int) {
	dc := NewContextFromRGBA(im)
	W, H := XY(Size(im))

	// grid
	{
		const cells = 100
		cellSize := float64(im.Bounds().Dx()) / cells

		dc.SetColor(White)
		dc.SetLineWidth(0.6)
		for i := range Range(cells) {
			z := float64(i) * cellSize
			dc.DrawLine(complex(0, z), complex(W, z))
			dc.Stroke()
			dc.DrawLine(complex(z, 0), complex(z, H))
			dc.Stroke()
		}
	}

	// dc.TransformAdd(Translate(Size(im) / 2))
	// dc.TransformAdd(Scale(Diag(0.9)))
	// dc.TransformAdd(Translate(-Size(im) / 2))

	w := W / float64(n)

	for i := range Range(n) {
		for j := range Range(n) {
			f := complex(float64(i), float64(j))
			center := f*Coeff(w) + Diag(w/2)
			dc.Stack(func(dc *Context) {
				dc.TransformAdd(Translate(center))
				dc.SetColor(RandomItem(colorSchema))
				dc.DrawCircle(0, w/2*RandomF64(0.1, 0.5))
				dc.Fill()

				dc.TransformAdd(Rotate(float64(RandomInt(10))))
				dc.SetColor(RandomItem(colorSchema))
				dc.SetLineWidth(lineWidth)

				r := w / 2 * RandomF64(0.6, 0.95)
				switch RandomInt(4) {
				case 0:
					// just orbit
					dc.DrawCircle(0, r)
					dc.Stroke()
				case 1:
					// orbit
					dc.DrawCircle(0, r)
					dc.Stroke()
					// satellites
					n := RandomIntN(1, 4) * 2
					for range Range(n) {
						dc.TransformAdd(Rotate(PI * 2 / float64(n)))
						dc.DrawCircle(complex(r, 0), r*0.1)
						dc.Fill()
					}
				case 2:
					n := RandomIntN(8, 20)
					theta := PI * 0.5 * float64(RandomIntN(1, 5))
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
						dc.TransformAdd(Rotate(PI * 2 / float64(n)))
						dc.DrawLine(complex(r/2, 0), complex(r*2/3-r*0.05, 0))
						dc.Stroke()
					}
				}
			})
		}
	}
}
