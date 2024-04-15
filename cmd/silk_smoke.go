package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// SilkSmoke draws a silk smoke image.
func SilkSmoke(
	im *image.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	alpha int,
	maxCircle, maxStepsPerCircle int,
	minSteps, maxSteps, minRadius, maxRadius float64,
) {
	dc := NewContextFromRGBA(im)

	cn := RandomInt(maxCircle) + int(maxCircle/3)
	circles := newCircleSlice(cn, Size(im), minSteps, maxSteps, minRadius, maxRadius)

	for range Range(maxStepsPerCircle) {
		dc.SetColor(ColorRGBA255(color.RGBA{}, 5))
		dc.DrawRectangle(0, Size(im))
		dc.Fill()

		for _, c1 := range circles {
			for _, c2 := range circles {
				if c1 == c2 || Dist(c1.center, c2.center) > c1.radius+c2.radius {
					continue
				}

				dc.SetColor(ColorRGBA255(lineColor, alpha))
				dc.SetLineWidth(lineWidth)

				dc.LineTo(c1.center)
				dc.LineTo(c2.center)
				dc.LineTo((c1.center + c2.center) / 2)
				dc.LineTo(c1.center)

				dc.Stroke()
			}
		}

		circles = circleSliceUpdate(circles, Size(im))
	}
}
