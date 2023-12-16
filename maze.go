package gena

import (
	"math/rand"
)

type maze struct {
	step int
}

// NewMaze returns a maze generator.
func NewMaze(step int) *maze {
	return &maze{
		step: step,
	}
}

// Generative draws a random maze image.
func (m *maze) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetColor(c.LineColor)
	ctex.SetLineWidth(c.LineWidth)
	for x := 0; x < c.Width(); x += m.step {
		for y := 0; y < c.Height(); y += m.step {
			v := complex(float64(x), float64(y))
			if rand.Float32() > 0.5 {
				ctex.DrawLine(
					v,
					v+complex(float64(m.step), float64(m.step)),
				)
			} else {
				ctex.DrawLine(
					v+complex(float64(m.step), 0),
					v+complex(0, float64(m.step)),
				)
			}
			ctex.Stroke()
		}
	}
}
