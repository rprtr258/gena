package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

type colorCircle2 struct {
	circleNum int
}

// ColorCircle2 is version 2 of ColorCircle.
// It still draws the circle and point cloud.
//   - circleNum: The number of the circle in this drawing.
func ColorCircle2(c *image.RGBA, colorSchema []color.RGBA, circleNum int) {
	cc := &colorCircle2{
		circleNum: circleNum,
	}

	dc := NewContextForRGBA(c)
	for range Range(cc.circleNum) {
		v := Mul2(RandomV2(), Size(c))

		r1 := RandomFloat64(50.0, float64(c.Bounds().Dx())/4)
		r2 := RandomFloat64(10.0, float64(c.Bounds().Dx())/3)

		cc.circle(dc, colorSchema, v, r1, r2)
		if rand.Float64() < 0.3 {
			col := colorSchema[rand.Intn(len(colorSchema))]
			dc.SetColor(col)
			dc.DrawCircle(v, r1/2.0)
			dc.Fill()
		}
	}
}

func (cc *colorCircle2) circle(dc *Context, colorSchema []color.RGBA, v V2, d, dx float64) {
	col := colorSchema[rand.Intn(len(colorSchema))]
	for j := 0.0; j < dx; j += 1.0 {
		dd := d + j*2.0
		alpha := min(int((dx/j)*255.0), 255)
		col.A = uint8(alpha)
		dc.SetColor(col)
		for range Range(150) {
			theta := RandomFloat64(0, math.Pi*2)
			dc.DrawPoint(v+Polar(dd/2, theta), 0.51)
			dc.Fill()
		}
	}
}
