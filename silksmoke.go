package gena

import (
	"image/color"
	"math/rand"
)

// SilkSmoke draws a silk smoke image.
func SilkSmoke(c Canvas, lineWidth float64, lineColor color.RGBA, alpha int, maxCircle, maxStepsPerCircle int, minSteps, maxSteps, minRadius, maxRadius float64, isRandColor bool) {
	ctex := NewContextForRGBA(c.Img())

	cn := rand.Intn(maxCircle) + int(maxCircle/3)
	circles := newCircleSlice(cn, c.Width, c.Height, minSteps, maxSteps, minRadius, maxRadius)

	for i := 0; i < maxStepsPerCircle; i++ {
		ctex.SetRGBA255(color.RGBA{}, 5)
		ctex.DrawRectangle(0, c.Size())
		ctex.Fill()

		for _, c1 := range circles {
			for _, c2 := range circles {

				cl := lineColor
				if isRandColor {
					cl = c.ColorSchema[rand.Intn(len(c.ColorSchema))]
				}

				if c1 == c2 {
					continue
				}

				if Dist(c1.pos, c2.pos) <= c1.radius+c2.radius {
					cc := (c1.pos + c2.pos) / 2

					ctex.SetRGBA255(cl, alpha)
					ctex.SetLineWidth(lineWidth)

					ctex.LineToV2(c1.pos)
					ctex.LineToV2(c2.pos)
					ctex.LineToV2(cc)
					ctex.LineToV2(c1.pos)

					ctex.Stroke()
				}
			}
		}

		circles = circleSliceUpdate(circles, c.Width, c.Height)
	}
}
