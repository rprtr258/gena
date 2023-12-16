package gena

import (
	"math"
	"math/rand"
)

type colorCircle struct {
	circleNum int
}

func NewColorCircle(circleNum int) *colorCircle {
	return &colorCircle{
		circleNum: circleNum,
	}
}

// Generative draws a color circle images.
func (cc *colorCircle) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())

	for i := 0; i < cc.circleNum; i++ {
		rnd := rand.Intn(3)
		v := Mul2(complex(
			RandomFloat64(-0.1, 1.1),
			RandomFloat64(-0.1, 1.1),
		), c.Size())
		s := RandomFloat64(0, RandomFloat64(0, float64(c.Width()/2))) + 10
		if rnd == 2 {
			rnd = rand.Intn(3)
		}
		switch rnd {
		case 0:
			cc.drawCircleV1(ctex, c, v, s)
		case 1:
			ctex.SetLineWidth(RandomFloat64(0, 1))
			ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
			ctex.DrawCircleV2(v, RandomFloat64(0, s)/2)
			ctex.Stroke()
		case 2:
			cc.drawCircleV2(ctex, c, v, s)
		}
	}
}

func (cc *colorCircle) drawCircleV1(ctex *Context, c Canvas, v V2, s float64) {
	n := RandomRangeInt(4, 30)
	cs := RandomFloat64(2, 8)
	ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
	ctex.Push()
	ctex.Translate(v)
	for a := 0.0; a < math.Pi*2.0; a += math.Pi * 2.0 / float64(n) {
		ctex.DrawEllipse(Polar(s, a)/2, complex(cs/2, cs/2))
		ctex.Fill()
	}
	ctex.Pop()
}

func (cc *colorCircle) drawCircleV2(ctex *Context, c Canvas, v V2, s float64) {
	cl := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
	ctex.SetLineWidth(1.0)
	sx := s * RandomFloat64(0.1, 0.55)
	for j := 0.0001; j < sx; j++ {
		dd := s + j*2.0
		alpha := Constrain(int(255*sx/j), 0, 255)

		// alpha := RandomRangeInt(30, 150)
		cl.A = uint8(alpha)
		ctex.SetColor(cl)

		for i := 0; i < 200; i++ {
			theta := RandomFloat64(0, math.Pi*2)
			ctex.DrawPoint(v+Polar(dd*0.3, theta), 0.6)
			ctex.Stroke()
		}
	}
}
