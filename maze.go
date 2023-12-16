package gena

import (
	"image/color"
	"math/rand"
)

// Maze draws a random maze image.
func Maze(c Canvas, lineWidth float64, lineColor color.RGBA, step int) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetColor(lineColor)
	ctex.SetLineWidth(lineWidth)
	for x := 0; x < c.Width; x += step {
		for y := 0; y < c.Height; y += step {
			v := complex(float64(x), float64(y))
			if rand.Float32() > 0.5 {
				ctex.DrawLine(
					v,
					v+complex(float64(step), float64(step)),
				)
			} else {
				ctex.DrawLine(
					v+complex(float64(step), 0),
					v+complex(0, float64(step)),
				)
			}
			ctex.Stroke()
		}
	}
}
