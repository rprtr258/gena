package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Maze draws a random maze image.
func Maze(im *image.RGBA, lineWidth float64, lineColor color.RGBA, step int) {
	dc := NewContextFromRGBA(im)
	dc.SetColor(lineColor)
	dc.SetLineWidth(lineWidth)
	for x := 0; x < im.Bounds().Dx(); x += step {
		for y := 0; y < im.Bounds().Dy(); y += step {
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
