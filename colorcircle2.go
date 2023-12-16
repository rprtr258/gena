package gena

import (
	"math"
	"math/rand"
)

type colorCircle2 struct {
	circleNum int
}

// NewColorCircle2 returns a colorCircle2 object.
func NewColorCircle2(circleNum int) *colorCircle2 {
	return &colorCircle2{
		circleNum: circleNum,
	}
}

// Generative draws a color circle image.
func (cc *colorCircle2) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())

	for i := 0; i < cc.circleNum; i++ {
		v := Mul2(RandomV2(), c.Size())

		r1 := RandomFloat64(50.0, float64(c.Width())/4)
		r2 := RandomFloat64(10.0, float64(c.Width())/3)

		cc.circle(ctex, c, v, r1, r2)
		if rand.Float64() < 0.3 {
			col := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
			ctex.SetColor(col)
			ctex.DrawCircleV2(v, r1/2.0)
			ctex.Fill()
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
