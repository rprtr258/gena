package gena

import (
	"math"
	"math/rand"
)

type oceanFish struct {
	lineNum int
	fishNum int
}

// OceanFish draws a ocean and fish images.
func OceanFish(c Canvas, lineNum, fishNum int) {
	o := &oceanFish{
		lineNum: lineNum,
		fishNum: fishNum,
	}

	ctex := NewContextForRGBA(c.Img())

	o.drawlines(ctex, c)

	for i := 0; i < o.fishNum; i++ {
		ctex.Push()

		theta := float64(360*i) / float64(o.fishNum)
		r := float64(c.Width()) / 4.0

		ctex.Push()
		ctex.Translate(Mul2(c.Size()/2, Polar(r, Radians(theta))))
		ctex.Rotate(Radians(theta + 90))
		o.drawfish(ctex, c, 0, float64(c.Width())/10)
		ctex.Pop()

		ctex.Clip()
		o.drawlines(ctex, c)
		ctex.Pop()
		ctex.ClearPath()
		ctex.ResetClip()
	}
}

func (o *oceanFish) drawlines(ctx *Context, c Canvas) {
	for i := 0; i < o.lineNum; i++ {
		cl := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
		ctx.SetColor(cl)
		ctx.SetLineWidth(RandomFloat64(3, 20))
		y := rand.Float64() * float64(c.Height())
		ctx.DrawLine(
			complex(0, y+RandomFloat64(-50, 50)),
			complex(float64(c.Width()), y+RandomFloat64(-50, 50)),
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

func (o *oceanFish) drawfish(ctex *Context, c Canvas, oo V2, r float64) {
	ctex.Push()
	ctex.Translate(oo)
	ctex.Rotate(Radians(180))

	ctex.MoveToV2(fishPt(r, 0))
	for i := 1; i < 361; i++ {
		ctex.LineToV2(fishPt(r, Radians(float64(i))))
	}
	ctex.ClosePath()
	ctex.Pop()
}
