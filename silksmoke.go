package gena

import (
	"image/color"
	"math/rand"
)

type sileSmoke struct {
	maxCircle         int
	maxStepsPerCircle int
	minSteps          float64
	maxSteps          float64
	minRadius         float64
	maxRadius         float64
	isRandColor       bool
}

func NewSilkSmoke(mc, msp int, minStep, maxStep, minRadius, maxRadius float64, isRandColor bool) *sileSmoke {
	return &sileSmoke{
		maxCircle:         mc,
		maxStepsPerCircle: msp,
		minSteps:          minStep,
		maxSteps:          maxStep,
		minRadius:         minRadius,
		maxRadius:         maxRadius,
		isRandColor:       isRandColor,
	}
}

// Generative draws a silk smoke image.
func (s *sileSmoke) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())

	cn := rand.Intn(s.maxCircle) + int(s.maxCircle/3)
	circles := newCircleSlice(cn, c.Width(), c.Height(), s.minSteps, s.maxSteps, s.minRadius, s.maxRadius)

	for i := 0; i < s.maxStepsPerCircle; i++ {
		ctex.SetRGBA255(color.RGBA{}, 5)
		ctex.DrawRectangle(0, c.Size())
		ctex.Fill()

		for _, c1 := range circles {
			for _, c2 := range circles {
				cl := c.LineColor
				if s.isRandColor {
					cl = c.ColorSchema[rand.Intn(len(c.ColorSchema))]
				}

				if c1 == c2 {
					continue
				}

				if Dist(c1.pos, c2.pos) <= c1.radius+c2.radius {
					cc := c1.pos + c2.pos/2
					ctex.SetRGBA255(cl, c.Alpha)
					ctex.SetLineWidth(c.LineWidth)
					ctex.LineToV2(c1.pos)
					ctex.LineToV2(c2.pos)
					ctex.LineToV2(cc)
					ctex.LineToV2(c1.pos)
					ctex.Stroke()
				}
			}
		}

		circles = circleSliceUpdate(circles, c.Width(), c.Height())
	}
}
