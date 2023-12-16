package gena

import (
	"image/color"
	"math/rand"
)

type spiralSquare struct {
	squareNum int
	decay     float64
	rectSide  float64
	fg        color.RGBA
	randColor bool
}

// NewSpiralSquare returns a spiralSquare object.
func NewSpiralSquare(
	squareNum int,
	rectSide, decay float64,
	fg color.RGBA, randColor bool,
) *spiralSquare {
	return &spiralSquare{
		squareNum: squareNum,
		decay:     decay,
		rectSide:  rectSide,
		fg:        fg, randColor: randColor,
	}
}

// Generative draws a spiral square images.
func (s *spiralSquare) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())

	sl := s.rectSide
	theta := rand.Intn(360) + 1
	for i := 0; i < s.squareNum; i++ {
		ctex.Push()
		ctex.Translate(c.Size() / 2)
		ctex.Rotate(Radians(float64(theta * (i + 1))))

		ctex.Scale(complex(sl, sl))

		ctex.LineTo(-0.5, 0.5)
		ctex.LineTo(0.5, 0.5)
		ctex.LineTo(0.5, -0.5)
		ctex.LineTo(-0.5, -0.5)
		ctex.LineTo(-0.5, 0.5)

		ctex.SetLineWidth(c.LineWidth)
		ctex.SetColor(c.LineColor)
		ctex.StrokePreserve()

		if s.randColor {
			ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
		} else {
			ctex.SetColor(s.fg)
		}
		ctex.Fill()
		ctex.Pop()
		sl -= s.decay * s.rectSide

		if sl < 0 {
			return
		}
	}
}
