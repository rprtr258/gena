package gena

import (
	"math/rand"
)

type silkSky struct {
	circleNum int
	sunRadius float64
}

// NewSilkSky returns a silkSky object.
func NewSilkSky(circleNum int, sunRadius float64) *silkSky {
	return &silkSky{
		circleNum: circleNum,
		sunRadius: sunRadius,
	}
}

// Generative draws a silk sky image.
func (s *silkSky) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	xm := float64(rand.Intn(c.Width()/5)) + float64(c.Width()*4/5-c.Width()/5)
	ym := float64(rand.Intn(c.Height()/5)) + float64(c.Height()*4/5-c.Height()/5)

	mh := s.circleNum*2 + 2
	ms := s.circleNum*2 + 50
	mv := 100

	for i := 0; i < s.circleNum; i++ {
		for j := 0; j < s.circleNum; j++ {
			hsv := HSV{
				H: s.circleNum + j,
				S: i + 50,
				V: 70,
			}
			rgba := hsv.ToRGB(mh, ms, mv)
			n := Div(Mul2(Plus(complex(float64(i), float64(j)), 0.5), c.Size()), float64(s.circleNum))
			ctex.SetRGBA255(rgba, c.Alpha)
			r := Dist(n, complex(xm, ym))
			ctex.DrawCircleV2(n, r-s.sunRadius/2)
			ctex.Fill()
		}
	}
}
