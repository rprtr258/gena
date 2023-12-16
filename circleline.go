package gena

import (
	"image/color"
	"math"
	"math/rand"
)

// CircleLine draws a cirle line image.
func CircleLine(c Canvas, lineWidth float64, lineColor color.RGBA, step float64, lineNum int, radius float64, axis V2) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(lineWidth)
	ctex.SetColor(lineColor)
	var points []V2
	for theta := -math.Pi; theta <= math.Pi; theta += step {
		points = append(points, ToPixel(Polar(radius, theta), axis, c.Size()))
	}

	for i := 0; i < lineNum; i++ {
		p1 := points[rand.Intn(len(points))]
		ctex.MoveToV2(p1)
		p2 := points[rand.Intn(len(points))]
		ctex.LineToV2(p2)
		ctex.Stroke()
	}
}
