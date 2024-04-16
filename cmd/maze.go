package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Maze draws a random maze image.
func Maze(dc *Context, lineWidth float64, lineColor color.RGBA, step int) {
	dc.SetColor(Azure)
	dc.Clear()

	dc.SetColor(lineColor)
	dc.SetLineWidth(lineWidth)
	for x := 0; x < dc.Image().Rect.Dx(); x += step {
		for y := 0; y < dc.Image().Rect.Dy(); y += step {
			v := complex(float64(x), float64(y))
			if Random() > 0.5 {
				dc.DrawLine(
					v,
					v+complex(float64(step), float64(step)),
				)
			} else {
				dc.DrawLine(
					v+complex(float64(step), 0),
					v+complex(0, float64(step)),
				)
			}
			dc.Stroke()
		}
	}
}

func maze() *image.RGBA {
	dc := NewContext(Diag(600))
	Maze(dc, 3, Orange, 20)
	return dc.Image()
}
