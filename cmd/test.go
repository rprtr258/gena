package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

func test() *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, 500, 500))

	dc := NewContextFromRGBA(dest)
	dc.Stack(func(dc *Context) {
		dc.TransformAdd(Translate(complex(500/2, 500/2)))
		dc.TransformAdd(Rotate(40))
		dc.SetColor(color.RGBA{0xFF, 0x00, 0x00, 255})
		for i := range Range(361) {
			theta := Radians(float64(i))
			p := complex(
				Cos(theta)-Pow(Sin(theta), 2)/Sqrt(2.0),
				Cos(theta)*Sin(theta),
			)

			alpha := Radians(float64(i + 1))

			p1 := complex(
				Cos(alpha)-Pow(Sin(alpha), 2)/Sqrt(2.0),
				Cos(alpha)*Sin(alpha),
			)

			dc.DrawLine(p*100, p1*100)
			dc.Stroke()
		}
	})
	return dc.Image()
}
