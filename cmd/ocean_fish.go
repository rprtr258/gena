package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

func drawOceanFishLines(dc *Context, size V2, colorSchema []color.RGBA, n int) {
	for range Range(n) {
		dc.SetColor(RandomItem(colorSchema))
		dc.SetLineWidth(RandomF64(9, 20))
		y := RandomF64(0, Y(size))
		dc.DrawLine(
			complex(0, y+RandomF64(-50, 50)),
			complex(X(size), y+RandomF64(-50, 50)),
		)
		dc.Stroke()
	}
}

func fishPt(r, theta float64) V2 {
	return complex(
		r*Cos(theta)-r*Pow(Sin(theta), 2)/Sqrt(2.0),
		r*Cos(theta)*Sin(theta),
	)
}

func drawfish(dc *Context, r float64) {
	dc.TransformAdd(Rotate(PI * 3 / 2))

	for _, a := range RangeF64(0, PI*2, 360) {
		dc.LineTo(fishPt(r, a))
	}
	dc.ClosePath()
}

// OceanFish draws an ocean and some fishes in the center.
//   - lineNum: The number of the line used to simulate the ocean wave.
//   - fishNum: The number of fish.
func OceanFish(im *image.RGBA, colorSchema []color.RGBA, lineNum, fishNum int) {
	W := X(Size(im))

	dc := NewContextFromRGBA(im)
	dc.SetColor(Black)
	dc.Clear()

	drawOceanFishLines(dc, Size(im), colorSchema, lineNum)

	dc.TransformAdd(Translate(Size(im) / 2))
	for i := range Range(fishNum) {
		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(complex(W/4, 0)))
			drawfish(dc, W/10)
			dc.Clip()

			dc.TransformAdd(Rotate(-float64(i) * PI * 2 / float64(fishNum)))
			dc.TransformAdd(Rotate(PI / 2))
			dc.TransformAdd(Translate(-complex(W/4, 0)))
			drawOceanFishLines(dc, Size(im), colorSchema, lineNum)
			dc.TransformAdd(Rotate(float64(i) * PI * 2 / float64(fishNum)))
		})
		dc.ClearPath()
		dc.ResetClip()
		dc.TransformAdd(Rotate(PI * 2 / float64(fishNum)))
	}
}
