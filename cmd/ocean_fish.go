package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

type oceanFish struct {
	lineNum int
	fishNum int
}

// OceanFish draws an ocean and some fishes in the center.
//   - lineNum: The number of the line used to simulate the ocean wave.
//   - fishNum: The number of fish.
func OceanFish(im *image.RGBA, colorSchema []color.RGBA, lineNum, fishNum int) {
	o := &oceanFish{
		lineNum: lineNum,
		fishNum: fishNum,
	}

	dc := NewContextFromRGBA(im)

	o.drawlines(dc, im, colorSchema)

	for i := range Range(o.fishNum) {
		dc.Stack(func(ctx *Context) {
			dc.Stack(func(ctx *Context) {
				theta := float64(360*i) / float64(o.fishNum)
				r := float64(im.Bounds().Dx()) / 4.0
				dc.TransformAdd(Translate(Mul2(Size(im)/2, Polar(r, Radians(theta)))))
				dc.TransformAdd(Rotate(Radians(theta + 90)))
				o.drawfish(dc, 0, float64(im.Bounds().Dx())/10)
			})
			dc.Clip()
			o.drawlines(dc, im, colorSchema)
		})
		dc.ClearPath()
		dc.ResetClip()
	}
}

func (o *oceanFish) drawlines(ctx *Context, im *image.RGBA, colorSchema []color.RGBA) {
	for range Range(o.lineNum) {
		cl := RandomItem(colorSchema)
		ctx.SetColor(cl)
		ctx.SetLineWidth(RandomF64(3, 20))
		y := Random() * float64(im.Bounds().Dy())
		ctx.DrawLine(
			complex(0, y+RandomF64(-50, 50)),
			complex(float64(im.Bounds().Dx()), y+RandomF64(-50, 50)),
		)
		ctx.Stroke()
	}
}

func fishPt(r, theta float64) V2 {
	return complex(
		r*Cos(theta)-r*Pow(Sin(theta), 2)/Sqrt(2.0),
		r*Cos(theta)*Sin(theta),
	)
}

func (o *oceanFish) drawfish(dc *Context, v V2, r float64) {
	dc.Stack(func(ctx *Context) {
		dc.TransformAdd(Translate(v))
		dc.TransformAdd(Rotate(Radians(180)))

		dc.MoveTo(fishPt(r, 0))
		for i := 1; i < 361; i++ {
			dc.LineTo(fishPt(r, Radians(float64(i))))
		}
		dc.ClosePath()
	})
}
