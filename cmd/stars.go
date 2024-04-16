package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func Polygon(n int) []V2 {
	result := make([]V2, n)
	for i := range result {
		result[i] = Rotation(PI * (float64(i)*2/float64(n) - 0.5))
	}
	return result
}

func stars() *image.RGBA {
	const W = 1200
	const H = 120
	const S = 100
	dc := NewContext(complex(W, H))
	dc.SetColor(ColorHex("#FFFFFF"))
	dc.Clear()
	n := 5
	points := Polygon(n)
	for x := S / 2; x < W; x += S {
		dc.Stack(func(ctx *Context) {
			s := Random()*S/4 + S/4
			dc.TransformAdd(Translate(complex(float64(x), H/2)))
			dc.TransformAdd(Rotate(Random() * 2 * PI))
			dc.TransformAdd(Scale(complex(s, s)))
			for i := range Range(n + 1) {
				dc.LineTo(points[(i*2)%n])
			}
			dc.SetLineWidth(10)
			dc.SetColor(ColorHex("#FFCC00"))
			dc.StrokePreserve()
			dc.SetColor(ColorHex("#FFE43A"))
			dc.Fill()
		})
	}
	return dc.Image()
}
