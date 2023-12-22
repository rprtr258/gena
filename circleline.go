package gena

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

// CircleLine draws a cirle line image.
func CircleLine(c *image.RGBA, lineWidth float64, lineColor color.RGBA, step float64, lineNum int, radius float64, axis V2) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(lineWidth)
	dc.SetColor(lineColor)
	var points []V2
	for theta := -math.Pi; theta <= math.Pi; theta += step {
		points = append(points, ToPixel(Polar(radius, theta), axis, Size(c)))
	}

	for range lineNum {
		p1 := points[rand.Intn(len(points))]
		dc.MoveToV2(p1)
		p2 := points[rand.Intn(len(points))]
		dc.LineToV2(p2)
		dc.Stroke()
	}
}
