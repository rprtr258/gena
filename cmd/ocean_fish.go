package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

type oceanFish struct {
	lineNum int
	fishNum int
}

// OceanFish draws an ocean and some fishes in the center.
//   - lineNum: The number of the line used to simulate the ocean wave.
//   - fishNum: The number of fish.
func OceanFish(c *image.RGBA, colorSchema []color.RGBA, lineNum, fishNum int) {
	o := &oceanFish{
		lineNum: lineNum,
		fishNum: fishNum,
	}

	dc := NewContextFromRGBA(c)

	o.drawlines(dc, c, colorSchema)

	for i := range Range(o.fishNum) {
		dc.Stack(func(ctx *Context) {
			dc.Stack(func(ctx *Context) {
				theta := float64(360*i) / float64(o.fishNum)
				r := float64(c.Bounds().Dx()) / 4.0
				dc.TransformAdd(Translate(Mul2(Size(c)/2, Polar(r, Radians(theta)))))
				dc.TransformAdd(Rotate(Radians(theta + 90)))
				o.drawfish(dc, 0, float64(c.Bounds().Dx())/10)
			})
			dc.Clip()
			o.drawlines(dc, c, colorSchema)
		})
		dc.ClearPath()
		dc.ResetClip()
	}
}

func (o *oceanFish) drawlines(ctx *Context, c *image.RGBA, colorSchema []color.RGBA) {
	for range Range(o.lineNum) {
		cl := colorSchema[rand.Intn(len(colorSchema))]
		ctx.SetColor(cl)
		ctx.SetLineWidth(RandomFloat64(3, 20))
		y := rand.Float64() * float64(c.Bounds().Dy())
		ctx.DrawLine(
			complex(0, y+RandomFloat64(-50, 50)),
			complex(float64(c.Bounds().Dx()), y+RandomFloat64(-50, 50)),
		)
		ctx.Stroke()
	}
}

func fishPt(r, theta float64) V2 {
	return complex(
		r*math.Cos(theta)-r*math.Pow(math.Sin(theta), 2)/math.Sqrt(2),
		r*math.Cos(theta)*math.Sin(theta),
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
