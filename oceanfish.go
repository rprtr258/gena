package gena

import (
	"image"
	"image/color"
	"math"
	"math/rand"
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

	dc := NewContextForRGBA(c)

	o.drawlines(dc, c, colorSchema)

	for i := range o.fishNum {
		dc.Stack(func(ctx *Context) {
			dc.Stack(func(ctx *Context) {
				theta := float64(360*i) / float64(o.fishNum)
				r := float64(c.Bounds().Dx()) / 4.0
				dc.Translate(Mul2(Size(c)/2, Polar(r, Radians(theta))))
				dc.Rotate(Radians(theta + 90))
				o.drawfish(dc, c, 0, float64(c.Bounds().Dx())/10)
			})
			dc.Clip()
			o.drawlines(dc, c, colorSchema)
		})
		dc.ClearPath()
		dc.ResetClip()
	}
}

func (o *oceanFish) drawlines(ctx *Context, c *image.RGBA, colorSchema []color.RGBA) {
	for range o.lineNum {
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

func (o *oceanFish) drawfish(dc *Context, c *image.RGBA, v V2, r float64) {
	dc.Stack(func(ctx *Context) {
		dc.Translate(v)
		dc.Rotate(Radians(180))

		dc.MoveToV2(fishPt(r, 0))
		for i := 1; i < 361; i++ {
			dc.LineToV2(fishPt(r, Radians(float64(i))))
		}
		dc.ClosePath()
	})
}
