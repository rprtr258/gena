package gena

import (
	"math"
	"math/rand"
)

type colorCircle2 struct {
	circleNum int
}

// ColorCircle2 is version 2 of ColorCircle.
// It still draws the circle and point cloud.
//   - circleNum: The number of the circle in this drawing.
func ColorCircle2(c Canvas, circleNum int) {
	cc := &colorCircle2{
		circleNum: circleNum,
	}

	dc := NewContextForRGBA(c.Img())
	for i := 0; i < cc.circleNum; i++ {
		v := Mul2(RandomV2(), c.Size())

		r1 := RandomFloat64(50.0, float64(c.Width)/4)
		r2 := RandomFloat64(10.0, float64(c.Width)/3)

		cc.circle(dc, c, v, r1, r2)
		if rand.Float64() < 0.3 {
			col := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
			dc.SetColor(col)
			dc.DrawCircleV2(v, r1/2.0)
			dc.Fill()
		}
	}
}

func (cc *colorCircle2) circle(ctex *Context, c Canvas, v V2, d, dx float64) {
	col := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
	for j := 0.0; j < dx; j += 1.0 {
		dd := d + j*2.0
		alpha := min(int((dx/j)*255.0), 255)
		col.A = uint8(alpha)
		ctex.SetColor(col)
		for i := 0; i < 150; i++ {
			theta := RandomFloat64(0, math.Pi*2)
			ctex.DrawPoint(v+Polar(dd/2, theta), 0.51)
			ctex.Fill()
		}
	}
}
