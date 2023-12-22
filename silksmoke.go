package gena

import (
	"image"
	"image/color"
	"math/rand"
)

// SilkSmoke draws a silk smoke image.
func SilkSmoke(
	c *image.RGBA,
	colorSchema []color.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	alpha int,
	maxCircle, maxStepsPerCircle int,
	minSteps, maxSteps, minRadius, maxRadius float64,
	isRandColor bool,
) {
	dc := NewContextForRGBA(c)

	cn := rand.Intn(maxCircle) + int(maxCircle/3)
	circles := newCircleSlice(cn, c.Bounds().Dx(), c.Bounds().Dy(), minSteps, maxSteps, minRadius, maxRadius)

	for range maxStepsPerCircle {
		dc.SetRGBA255(color.RGBA{}, 5)
		dc.DrawRectangle(0, Size(c))
		dc.Fill()

		for _, c1 := range circles {
			for _, c2 := range circles {

				cl := lineColor
				if isRandColor {
					cl = colorSchema[rand.Intn(len(colorSchema))]
				}

				if c1 == c2 {
					continue
				}

				if Dist(c1.pos, c2.pos) <= c1.radius+c2.radius {
					dc.SetRGBA255(cl, alpha)
					dc.SetLineWidth(lineWidth)

					dc.LineToV2(c1.pos)
					dc.LineToV2(c2.pos)
					dc.LineToV2((c1.pos + c2.pos) / 2)
					dc.LineToV2(c1.pos)

					dc.Stroke()
				}
			}
		}

		circles = circleSliceUpdate(circles, c.Bounds())
	}
}
