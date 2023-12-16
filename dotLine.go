package gena

import (
	"image/color"
	"math/rand"
)

// TODO: figure out what the meaning of these parameters.
type dotLine struct {
	n         int
	ras       float64
	canv      float64
	randColor bool
	iters     int
}

func NewDotLine(n int, ras, canv float64, randColor bool, iters int) *dotLine {
	return &dotLine{
		n:         n,
		ras:       ras,
		canv:      canv,
		randColor: randColor,
		iters:     iters,
	}
}

func (d *dotLine) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())

	ctex.SetLineWidth(c.LineWidth)
	var dir []int = []int{-1, 1}
	for i := 0; i < d.iters; i++ {
		old := complex(
			float64(rand.Intn(d.n-1)),
			float64(rand.Intn(d.n-1)),
		)

		n := rand.Intn(7)
		if d.randColor {
			ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
		} else {
			ctex.SetRGBA255(color.RGBA{
				uint8(RandomRangeInt(222, 255)),
				uint8(RandomRangeInt(20, 222)),
				0,
				0},
				255)
		}
		for j := 0; j < n; j++ {
			new := old + complex(float64(dir[rand.Intn(2)]), float64(dir[rand.Intn(2)]))
			if Dist(new, complex(float64(d.n/2), float64(d.n/2))) > float64(d.n/2-10) {
				new = old
			}
			if X(new) == X(old) && rand.Intn(6) > 4 {
				ctex.DrawEllipse(Plus(Mul(old, d.ras), d.canv), complex(c.LineWidth, c.LineWidth))
				ctex.Fill()
				continue
			}
			ctex.DrawLine(
				Plus(Mul(old, d.ras), d.canv),
				Plus(Mul(new, d.ras), d.canv),
			)
			old = new

			ctex.Stroke()
		}
	}
}
